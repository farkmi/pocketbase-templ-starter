package server

import (
	"github.com/farkmi/pocketbase-templ-starter/internal/mailer"
	"github.com/farkmi/pocketbase-templ-starter/internal/server/config"
	"github.com/farkmi/pocketbase-templ-starter/internal/server/hooks"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func InitConfigFromEnv(e *core.BootstrapEvent) error {
	config.SetConfigFromEnv()

	return e.Next()
}

func InitServer(app *pocketbase.PocketBase, isGoRun bool) error {
	log := app.Logger()

	app.OnBootstrap().BindFunc(InitConfigFromEnv)
	app.OnServe().BindFunc(hooks.OverrideAppSettingsFromEnv)

	if err := BindCommands(app, isGoRun); err != nil {
		log.Error("Failed to bind commands", "error", err)
	}

	if err := BindRoutes(app); err != nil {
		log.Error("Failed to bind routes", "error", err)
	}

	if err := BindCollectionHooks(app); err != nil {
		log.Error("Failed to bind collection hooks", "error", err)
	}

	if err := BindCronjobs(app); err != nil {
		log.Error("Failed to bind cronjobs", "error", err)
	}

	if err := mailer.Init(app); err != nil {
		log.Error("Failed to initialize mailer", "error", err)
	}

	return nil
}
