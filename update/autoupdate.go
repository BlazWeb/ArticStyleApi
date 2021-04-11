package update

import (
	"net/http"

	"github.com/inconshreveable/go-update"
)

func DoUpdate(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		return err
	}
	return err
}
