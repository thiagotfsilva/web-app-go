package controllers

import (
	"net/http"
	"web-app-go/src/utils"
)

func LoadUserRegisterView(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "register.html", nil)
}
