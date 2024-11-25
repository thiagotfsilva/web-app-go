package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"web-app-go/src/config"
	"web-app-go/src/cookies"
	"web-app-go/src/models"
	"web-app-go/src/response"
	"web-app-go/src/utils"
)

func LoadLoginView(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.ReadCookies(r)
	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
	}
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

	url := fmt.Sprintf("%s/login", config.ApiUrl)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroResponse{Erro: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		response.HandleStatusCode(w, res)
		return
	}

	var AuthData models.Auth
	if err = json.NewDecoder(res.Body).Decode(&AuthData); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErroResponse{Erro: err.Error()})
		return
	}

	if err = cookies.SaveCookies(w, AuthData.Id, AuthData.Token); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErroResponse{Erro: err.Error()})
		return
	}
	response.JSON(w, http.StatusOK, nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookies.DeleteCookies(w)
	http.Redirect(w, r, "/login", http.StatusFound)
}
