package server

import (
	"io/fs"

	"github.com/farkmi/pocketbase-templ-starter/internal/cmd"
	"github.com/farkmi/pocketbase-templ-starter/internal/server/config"
	"github.com/farkmi/pocketbase-templ-starter/internal/server/hooks"
	"github.com/farkmi/pocketbase-templ-starter/public"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func BindRoutes(app core.App) error {
	// extract the embedded public filesystem
	public, err := fs.Sub(public.EmbeddedFS, ".")
	if err != nil {
		return err
	}

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// serves static files from the provided embedded public
		se.Router.GET("/{path...}", apis.Static(public, true))

		return se.Next()
	})

	return nil
}

func BindCollectionHooks(app core.App) error {
	envConfig := config.GetServerConfig()

	app.OnSettingsUpdateRequest().BindFunc(hooks.LockAppSettings)
	app.OnRecordCreate(envConfig.ImmutableCollections...).BindFunc(hooks.EnforceCollectionsImmutable)
	app.OnRecordDelete(envConfig.ImmutableCollections...).BindFunc(hooks.EnforceCollectionsImmutable)
	app.OnRecordUpdate(envConfig.ImmutableCollections...).BindFunc(hooks.EnforceCollectionsImmutable)

	return nil
}

func BindCommands(app *pocketbase.PocketBase, isGoRun bool) error {
	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Admin UI
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	app.RootCmd.AddCommand(cmd.SeedCmd)

	return nil
}

func BindCronjobs(_ *pocketbase.PocketBase) error {
	// log := app.Logger()
	appConfig := config.GetServerConfig()
	if !appConfig.UseBuiltinCron {
		return nil
	}

	// app.Cron().MustAdd("Sample cronjob", "0 8 * * *", func() {
	// 	_, _, err := cmd.ApplySeedData()
	// 	if err != nil {
	// 		log.Error("Failed to apply seed data", "error", err)
	// 	}
	// })

	return nil
}
