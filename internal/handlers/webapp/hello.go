package webapp

import (
	"github.com/a-h/templ"
	"github.com/farkmi/pocketbase-templ-starter/internal/handlers"
	"github.com/farkmi/pocketbase-templ-starter/internal/httperrors"
	"github.com/farkmi/pocketbase-templ-starter/web/templates/components"
	"github.com/pocketbase/pocketbase/core"
)

func HandleHello(e *core.RequestEvent) error {
	name := e.Request.URL.Query().Get("name")
	if name == "" {
		return handlers.ErrorHandler(e, httperrors.ErrBadRequestNameRequired)
	}

	return handlers.WrappedTemplHandler(e, templ.Handler(components.Hello(name)))
}
