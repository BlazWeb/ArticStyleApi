//go:generate goversioninfo
package main

import (
	"fmt"

	"github.com/Artic-Dev/ArticStyleApi-GO/db"
	"github.com/Artic-Dev/ArticStyleApi-GO/handlers"
	"github.com/Artic-Dev/ArticStyleApi-GO/update"
)

func main() {
	err := update.DoUpdate("https://github.com/Artic-Dev/ArticStyleApi.git")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = db.FirtConnection()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	handlers.Controller()
}
