package hooks

import (
	"errors"

	"github.com/farkmi/pocketbase-templ-starter/internal/server/config"
	"github.com/pocketbase/pocketbase/core"
)

func EnforceCollectionsImmutable(e *core.RecordEvent) error {
	isLocked := config.GetServerConfig().SetCollectionsImmutable
	if isLocked {
		return errors.New("collection is immutable")
	}

	return e.Next()
}
