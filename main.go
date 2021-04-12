//go:generate goversioninfo
package main

import (
	"log"
	"os"
	"time"

	"github.com/Artic-Dev/ArticStyleApi-GO/db"
	"github.com/Artic-Dev/ArticStyleApi-GO/handlers"
	"github.com/Artic-Dev/ArticStyleApi-GO/update"
)

func main() {
	handlers.Ascii()
	update.ConfirmAndSelfUpdate()
	err := db.FirtConnection()
	if err != nil {
		log.Println(err.Error())
		time.Sleep(4 * time.Second)
		os.Exit(1)
	}
	handlers.Controller()
}
