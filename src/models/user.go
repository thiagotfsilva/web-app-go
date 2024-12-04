package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"web-app-go/src/config"
	"web-app-go/src/request"
)

type User struct {
	ID           uint64        `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Nick         string        `json:"nick,omitempty"`
	Email        string        `json:"email,omitempty"`
	CreatedAt    time.Time     `json:"createdAt,omitempty"`
	Followers    []User        `json:"followers"`
	Following    []User        `json:"following"`
	Publications []Publication `json:"publications"`
}

func FindUser(id uint64, r *http.Request) (User, error) {
	channelUser := make(chan User)
	channelFollowers := make(chan []User)
	channelFollowing := make(chan []User)
	channelPublications := make(chan []Publication)

	go FindUserData(channelUser, id, r)
	go FindFollowers(channelFollowers, id, r)
	go FindoFollowing(channelFollowing, id, r)
	go FindPublications(channelPublications, id, r)

	var (
		user         User
		followers    []User
		following    []User
		publications []Publication
	)

	for i := 0; i < 4; i++ {
		select {
		case userData := <-channelUser:
			if userData.ID == 0 {
				return User{}, errors.New("erro ao buscar um usuário")
			}

			user = userData

		case followersData := <-channelFollowers:
			if followersData == nil {
				return User{}, errors.New("erro ao buscar um seguidores")
			}

			followers = followersData

		case followingData := <-channelFollowing:
			if followingData == nil {
				return User{}, errors.New("erro ao buscar quem estou seguind")
			}

			following = followingData

		case publicationsData := <-channelPublications:
			if publicationsData == nil {
				return User{}, errors.New("erro ao buscar publicações")
			}

			publications = publicationsData
		}

	}

	user.Followers = followers
	user.Following = following
	user.Publications = publications

	return user, nil
}

func FindUserData(channel chan<- User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.ApiUrl, userId)
	res, err := request.HandlerRequestAuthenticate(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- User{}
		return
	}
	defer res.Body.Close()

	var user User
	if err = json.NewDecoder(res.Body).Decode(&user); err != nil {
		channel <- User{}
		return
	}

	channel <- user
}

// bucas seguidores
func FindFollowers(channel chan<- []User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.ApiUrl, userId)
	res, err := request.HandlerRequestAuthenticate(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer res.Body.Close()

	var followers []User
	if err = json.NewDecoder(res.Body).Decode(&followers); err != nil {
		channel <- nil
		return
	}

	channel <- followers

}

func FindoFollowing(channel chan<- []User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.ApiUrl, userId)
	res, err := request.HandlerRequestAuthenticate(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	res.Body.Close()

	var following []User
	if err = json.NewDecoder(res.Body).Decode(&following); err != nil {
		channel <- nil
		return
	}

	channel <- following
}

func FindPublications(channel chan<- []Publication, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/publications/%d", config.ApiUrl, userId)
	res, err := request.HandlerRequestAuthenticate(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer res.Body.Close()

	var publications []Publication
	if err = json.NewDecoder(res.Body).Decode(&publications); err != nil {
		channel <- nil
		return
	}

	channel <- publications

}
