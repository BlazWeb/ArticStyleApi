package main

import (
	"net/http"

	"github.com/Artic-Dev/ArticStyleApi-GO/handlers"
	"github.com/inconshreveable/go-update"
)

func main() {
	handlers.Controller()
	doUpdate("https://github.com/Artic-Dev/ArticStyleApi-GO.git")
}

func doUpdate(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		// error handling
	}
	return err
}
