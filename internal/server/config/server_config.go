package config

import (
	"github.com/farkmi/pocketbase-templ-starter/internal/util"
	"github.com/pocketbase/pocketbase/core"
)

var (
	defaultConfigInstance *ServerConfig
)

type ServerConfig struct {
	EmailTemplatesDir       string          // EmailTemplatesDir is the directory where email templates are stored
	SMTP                    core.SMTPConfig // SMTP is the configuration for the SMTP server. This overrides the app-level SMTP settings
	Meta                    core.MetaConfig // Meta is the configuration for the app metadata. This overrides the app-level Meta settings
	LockMetaSettings        bool            // LockMetaSettings is a flag to lock the Meta settings - useful for production or if setting through env
	LockSMTPSettings        bool            // LockSMTPSettings is a flag to lock the SMTP settings - useful for production or if setting through env
	UseBuiltinCron          bool            // UseBuiltinCron is a flag to control whether the commands defined in BindCronjobs() are bound on startup
	SetCollectionsImmutable bool            // SetCollectionsImmutable is a flag to lock the collections defined in ImmutableCollections
	ImmutableCollections    []string        // ImmutableCollections is a list of collections that can be locked
}

func GetServerConfig() *ServerConfig {
	if defaultConfigInstance == nil {
		SetConfigFromEnv()
	}

	return defaultConfigInstance
}

func SetConfigFromEnv() {
	conf := &ServerConfig{
		EmailTemplatesDir: util.GetEnv("EMAIL_TEMPLATES_DIR", "web/templates/email"),
		SMTP: core.SMTPConfig{
			Host:     util.GetEnv("SMTP_HOST", "mailhog"),
			Port:     util.GetEnvAsInt("SMTP_PORT", 1025),
			Username: util.GetEnv("SMTP_USERNAME", ""),
			Password: util.GetEnv("SMTP_PASSWORD", ""),

			Enabled:    util.GetEnvAsBool("SMTP_ENABLED", true),
			AuthMethod: util.GetEnv("SMTP_AUTH_METHOD", "PLAIN"),
		},
		Meta: core.MetaConfig{
			AppName:       util.GetEnv("META_APP_NAME", "PocketBase"),
			AppURL:        util.GetEnv("META_APP_URL", "https://example.com"),
			SenderName:    util.GetEnv("META_SENDER_NAME", "Test Sender"),
			SenderAddress: util.GetEnv("META_SENDER_ADDRESS", "support@example.com"),
		},
		LockMetaSettings:        util.GetEnvAsBool("LOCK_META_SETTINGS", false),
		LockSMTPSettings:        util.GetEnvAsBool("LOCK_SMTP_SETTINGS", true),
		UseBuiltinCron:          util.GetEnvAsBool("USE_BUILTIN_CRON", true),
		SetCollectionsImmutable: util.GetEnvAsBool("SET_COLLECTIONS_IMMUTABLE", false),
		ImmutableCollections:    util.GetEnvAsStringArrTrimmed("IMMUTABLE_COLLECTIONS", []string{"_superusers"}),
	}

	defaultConfigInstance = conf
}
