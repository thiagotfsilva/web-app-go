package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"web-app-go/src/config"
	"web-app-go/src/models"
	"web-app-go/src/request"
	"web-app-go/src/response"
	"web-app-go/src/utils"

	"github.com/gorilla/mux"
)

func CreatePublication(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publication, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})
	if err != nil {
		response.JSON(
			w,
			http.StatusBadRequest,
			response.ErroResponse{Erro: err.Error()},
		)
	}

	url := fmt.Sprintf("%s/publications", config.ApiUrl)
	res, err := request.HandlerRequestAuthenticate(
		r,
		http.MethodPost,
		url,
		bytes.NewBuffer(publication),
	)
	if err != nil {
		response.JSON(
			w,
			http.StatusInternalServerError,
			response.ErroResponse{Erro: err.Error()},
		)
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		response.HandleStatusCode(w, res)
	}

	response.JSON(w, res.StatusCode, nil)
}

func LikePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationId, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		response.JSON(
			w,
			http.StatusBadRequest,
			response.ErroResponse{Erro: err.Error()},
		)
		return
	}

	url := fmt.Sprintf("%s/publications/%d/like", config.ApiUrl, publicationId)

	res, err := request.HandlerRequestAuthenticate(
		r,
		http.MethodPost,
		url,
		nil,
	)
	if err != nil {
		response.JSON(
			w,
			http.StatusInternalServerError,
			response.ErroResponse{Erro: err.Error()},
		)
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		response.HandleStatusCode(w, res)
		return
	}

	response.JSON(w, res.StatusCode, nil)

}

func DislikePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationId, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		response.JSON(
			w,
			http.StatusBadRequest,
			response.ErroResponse{Erro: err.Error()},
		)
		return
	}

	url := fmt.Sprintf("%s/publications/%d/dislike", config.ApiUrl, publicationId)

	res, err := request.HandlerRequestAuthenticate(
		r,
		http.MethodPost,
		url,
		nil,
	)
	if err != nil {
		response.JSON(
			w,
			http.StatusInternalServerError,
			response.ErroResponse{Erro: err.Error()},
		)
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		response.HandleStatusCode(w, res)
		return
	}

	response.JSON(w, res.StatusCode, nil)
}

func LoadEditPage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationId, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroResponse{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d", config.ApiUrl, publicationId)
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

	var publication models.Publication
	if err = json.NewDecoder(res.Body).Decode(&publication); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErroResponse{Erro: err.Error()})
		return
	}

	utils.ExecTemplate(w, "edit-publication.html", publication)
}

func EditPublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationId, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroResponse{Erro: err.Error()})
		return
	}

	r.ParseForm()
	publication, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroResponse{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d", config.ApiUrl, publicationId)
	res, err := request.HandlerRequestAuthenticate(
		r,
		http.MethodPut,
		url,
		bytes.NewBuffer(publication),
	)
	if err != nil {
		response.JSON(
			w,
			http.StatusInternalServerError,
			response.ErroResponse{Erro: err.Error()},
		)
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		response.HandleStatusCode(w, res)
	}

	response.JSON(w, res.StatusCode, nil)
}

func DeletePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationId, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroResponse{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d", config.ApiUrl, publicationId)
	res, err := request.HandlerRequestAuthenticate(
		r,
		http.MethodDelete,
		url,
		nil,
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
