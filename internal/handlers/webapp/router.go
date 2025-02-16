package webapp

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/router"
)

func InitWebAppRoutes(baseRoute *router.RouterGroup[*core.RequestEvent]) {
	baseRoute.GET("/hello", HandleHello)
}
