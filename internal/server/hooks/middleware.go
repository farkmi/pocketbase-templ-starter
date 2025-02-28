package hooks

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/farkmi/pocketbase-templ-starter/internal/constants"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

func APIKeyMiddleware(e *core.RequestEvent) error {
	log := e.App.Logger()

	headerValue := strings.ToLower(e.Request.Header.Get(constants.APIKeyHeader))
	if headerValue == "" {
		log.Debug("x-api-key header value is missing!")
		return e.BadRequestError("x-api-key header value is missing!", nil)
	}

	apiKey, err := e.App.FindFirstRecordByFilter("apiKeys",
		"token ~ {:token}",
		dbx.Params{
			"token": headerValue,
		},
	)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Error("Failed to get current API key", "error", err)
			return err
		}
		log.Debug("API key not found", "token", headerValue)
		return e.UnauthorizedError("API key not valid", nil)
	}

	// should not be possible
	if headerValue != apiKey.GetString("token") {
		log.Error("API key was loaded but does not match the header value", "headerValue", headerValue, "apiKey", apiKey.GetString("token"))
		return e.UnauthorizedError("API key not valid", nil)
	}

	return e.Next()
}
