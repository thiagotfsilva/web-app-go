package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"web-app-go/src/config"
	"web-app-go/src/cookies"
	"web-app-go/src/models"
	"web-app-go/src/request"
	"web-app-go/src/response"
	"web-app-go/src/utils"

	"github.com/gorilla/mux"
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

	url := fmt.Sprintf("%s/users", config.ApiUrl)
	res, err := http.Post(
		url,
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

func LoadUserFindView(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	url := fmt.Sprintf("%s/users?user=%s", config.ApiUrl, nameOrNick)
	res, err := request.HandlerRequestAuthenticate(
		r,
		http.MethodGet,
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
	}

	var users []models.User
	if err = json.NewDecoder(res.Body).Decode(&users); err != nil {
		response.JSON(
			w,
			http.StatusUnprocessableEntity,
			response.ErroResponse{Erro: err.Error()},
		)
		return
	}

	utils.ExecTemplate(w, "users.html", users)
}

func LoadUserPageView(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroResponse{Erro: err.Error()})
		return
	}

	cookie, _ := cookies.ReadCookies(r)
	userIdLoged, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if userId == userIdLoged {
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}

	user, err := models.FindUser(userId, r)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroResponse{Erro: err.Error()})
		return
	}

	utils.ExecTemplate(w, "user.html", struct {
		User        models.User
		UserIdLoged uint64
	}{
		User:        user,
		UserIdLoged: userIdLoged,
	})
}

//localhost:5000/users/{userId}/unfollow
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroResponse{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/unfollow", config.ApiUrl, userId)
	res, err := request.HandlerRequestAuthenticate(r, http.MethodPost, url, nil)
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

func FollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroResponse{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/follow", config.ApiUrl, userId)
	res, err := request.HandlerRequestAuthenticate(r, http.MethodPost, url, nil)
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

func LoadUserProfileView(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.ReadCookies(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	user, err := models.FindUser(userId, r)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroResponse{Erro: err.Error()})
		return
	}

	utils.ExecTemplate(w, "profile.html", user)
}

func LoadFormEditUserView(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.ReadCookies(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	channel := make(chan models.User)
	go models.FindUserData(channel, userId, r)
	user := <-channel

	if user.ID == 0 {
		response.JSON(w, http.StatusInternalServerError, response.ErroResponse{Erro: "Erro ao buscar usuário"})
		return
	}

	utils.ExecTemplate(w, "edit-user.html", user)
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, err := json.Marshal(map[string]string{
		"name":  r.FormValue("name"),
		"nick":  r.FormValue("nick"),
		"email": r.FormValue("email"),
	})
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroResponse{Erro: err.Error()})
		return
	}

	cookie, _ := cookies.ReadCookies(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)
	url := fmt.Sprintf("%s/users/%d", config.ApiUrl, userId)

	res, err := request.HandlerRequestAuthenticate(
		r,
		http.MethodPut,
		url,
		bytes.NewBuffer(user),
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
		return
	}

	response.JSON(w, res.StatusCode, nil)
}
