package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"web-app-go/src/response"
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
		"nick":     r.FormValue("nick"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroResponse{Erro: err.Error()})
		return
	}

	res, err := http.Post(
		"http://localhost:5000/users",
		"application/json",
		bytes.NewBuffer(user),
	)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroResponse{Erro: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		response.HandleStatusCode(w, res)
		return
	}

	response.JSON(w, res.StatusCode, nil)
}
