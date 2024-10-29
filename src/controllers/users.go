package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"web-app-go/src/utils"
)

func LoadUserRegisterView(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "register.html", nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Pega o copor da requisição

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nickName": r.FormValue("nickName"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		log.Fatal(err)
	}

	response, err := http.Post(
		"http://localhost:5000/users",
		"application/json",
		bytes.NewBuffer(user),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
}
