package cmd

import (
	"database/sql"
	"errors"
	"log"

	"github.com/farkmi/pocketbase-templ-starter/internal/data"
	"github.com/farkmi/pocketbase-templ-starter/internal/util"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/spf13/cobra"
)

var SeedCmd = &cobra.Command{
	Use: "seed",
	Run: seedFunc,
}

func seedFunc(_ *cobra.Command, _ []string) {
	count, cols, err := ApplySeedData()
	if err != nil {
		log.Fatalf("Failed to seed data: %v", err)
	}
	log.Printf("Seeded %d records in %d collections", count, cols)
}

func ApplySeedData() (int, int, error) {
	log.Println("Seeding data...")

	appConfig := pocketbase.Config{
		DefaultDataDir: util.GetEnv("PB_DATA_DIR", "/app/pb_data"),
	}

	app := pocketbase.NewWithConfig(appConfig)

	if err := app.Bootstrap(); err != nil {
		return 0, 0, err
	}
	log := app.Logger()

	cols := 0
	count := 0

	err := app.RunInTransaction(func(txApp core.App) error {
		for col, seeds := range data.Seeds {
			collection, err := txApp.FindCollectionByNameOrId(col)
			if err != nil {
				log.Error("collection not found",
					"collection", col,
					"error", err,
				)
				return err
			}
			for _, seed := range seeds {
				id, ok := seed["id"]
				if !ok {
					log.Error("seed record missing id",
						"seed", seed,
					)
					return errors.New("seed record missing id")
				}

				record, err := txApp.FindRecordById(collection, id.(string))
				if err != nil && !errors.Is(err, sql.ErrNoRows) {
					log.Info("failed to find record by id", "error", err)
					return err
				}
				if record == nil {
					log.Info("creating new record", "seed", seed)
					record = core.NewRecord(collection)
					record.Set("id", id.(string))
				} else {
					log.Info("updating existing record", "seed", seed)
				}
				for key, value := range seed {
					if key == "id" {
						continue
					}
					record.SetRaw(key, value)
				}
				err = txApp.Save(record)
				if err != nil {
					log.Error("failed to save record", "error", err)
					return err
				}
				count++
			}
			cols++
		}

		return nil
	})
	if err != nil {
		log.Error("failed to seed data", "error", err)
		return 0, 0, err
	}

	return count, cols, nil
}
