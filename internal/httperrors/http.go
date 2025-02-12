package httperrors

import (
	"net/http"

	"github.com/pocketbase/pocketbase/core"
)

func ErrBadRequestInvalidEmail(e *core.RequestEvent) error {
	return e.String(http.StatusBadRequest, "Invalid email format")
}
