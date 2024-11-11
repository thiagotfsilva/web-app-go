package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web-app-go/src/config"
	"web-app-go/src/models"
	"web-app-go/src/request"
	"web-app-go/src/response"
	"web-app-go/src/utils"
)

func LoadHomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publications", config.ApiUrl)
	res, err := request.HandlerRequestAuthenticate(r, http.MethodGet, url, nil)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroResponse{Erro: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		response.HandleStatusCode(w, res)
		return
	}

	var publications []models.Publication
	if err = json.NewDecoder(res.Body).Decode(&publications); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErroResponse{Erro: err.Error()})
		return
	}

	utils.ExecTemplate(w, "home.html", publications)
}