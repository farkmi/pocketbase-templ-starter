package hooks

import (
	"github.com/farkmi/pocketbase-templ-starter/internal/server/config"
	"github.com/pocketbase/pocketbase/core"
)

func LockAppSettings(e *core.SettingsUpdateRequestEvent) error {
	lockMetaSettings := config.GetServerConfig().LockMetaSettings
	if lockMetaSettings && e.NewSettings.Meta != e.OldSettings.Meta {
		return e.BadRequestError("Cannot change Meta setting", nil)
	}

	lockSMTPSettings := config.GetServerConfig().LockSMTPSettings
	if lockSMTPSettings && e.NewSettings.SMTP != e.OldSettings.SMTP {
		return e.BadRequestError("Cannot change SMTP settings", nil)
	}

	return e.Next()
}

func OverrideAppSettingsFromEnv(se *core.ServeEvent) error {
	envConfig := config.GetServerConfig()

	appSettings := se.App.Settings()
	appSettings.SMTP = envConfig.SMTP
	appSettings.Meta = envConfig.Meta

	return se.Next()
}
