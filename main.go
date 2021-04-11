//go:generate goversioninfo
package main

import (
	"fmt"

	"github.com/Artic-Dev/ArticStyleApi-GO/db"
	"github.com/Artic-Dev/ArticStyleApi-GO/handlers"
)

func main() {
	_, err := db.FirtConnection()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	handlers.Controller()
}
