package routers

import (
	"fmt"
	"net/http"
)

type sendmessage struct {
	Message string `json:"message"`
	Status  bool   `json:"success"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Rest Api ArticStylses. Att: ArticDev")
}
