package handlers

import (
	"github.com/a-h/templ"
	"github.com/farkmi/pocketbase-templ-starter/internal/httperrors"
	"github.com/farkmi/pocketbase-templ-starter/web/templates/components"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func WrappedTemplHandler(e *core.RequestEvent, h *templ.ComponentHandler) error {
	return apis.WrapStdHandler(h)(e)
}

func ErrorHandler(e *core.RequestEvent, err *httperrors.HTTPError) error {
	templHandler := templ.Handler(components.ErrorPage(
		err.Code,
		err.Title,
		err.Type,
		err.Detail,
	))

	return apis.WrapStdHandler(templHandler)(e)
}
