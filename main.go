package main

import (
	"log"
	"os"
	"strings"

	"github.com/farkmi/pocketbase-templ-starter/internal/server"
	_ "github.com/farkmi/pocketbase-templ-starter/migrations" // TODO: uncomment as soon as you have ./migrations/*.go files
	"github.com/pocketbase/pocketbase"
)

func main() {

	app := pocketbase.New()

	// loosely check if it was executed using "go run"
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	if err := server.InitServer(app, isGoRun); err != nil {
		log.Fatal(err)
	}

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
