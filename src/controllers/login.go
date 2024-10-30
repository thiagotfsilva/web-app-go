package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"web-app-go/src/response"
	"web-app-go/src/utils"
)

func LoadLoginView(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "login.html", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroResponse{Erro: err.Error()})
		return
	}

	res, err := http.Post(
		"http://localhost:5000/login",
		"application/json",
		bytes.NewBuffer(user),
	)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroResponse{Erro: err.Error()})
		return
	}

	token, _ := io.ReadAll(res.Body)

	fmt.Println(res.StatusCode, string(token))
}
