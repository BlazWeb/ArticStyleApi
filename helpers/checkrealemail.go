package helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type RealEmail struct {
	Status string `json:"status"`
}

func CheckRealEmail(email string) (string, error) {
	var url = "https://isitarealemail.com/api/email/validate?email=" + url.QueryEscape(email)
	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var myJson RealEmail
	json.Unmarshal(body, &myJson)
	return myJson.Status, nil
}
