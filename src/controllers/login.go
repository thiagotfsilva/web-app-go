package controllers

import (
	"net/http"
	"web-app-go/src/utils"
)

func LoadLoginView(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "login.html", nil)
}
