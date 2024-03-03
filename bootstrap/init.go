package bootstrap

import (
	"fmt"
	"log"
)

func (app *App) Start() {
	fmt.Println("[api-coordinator] service started...")
	err := app.router.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
