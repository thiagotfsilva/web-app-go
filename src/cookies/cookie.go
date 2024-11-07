package cookies

import (
	"net/http"
	"web-app-go/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func Config() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func Save(w http.ResponseWriter, id, token string) error {
	data := map[string]string{
		"id":    id,
		"token": token,
	}

	dataEncode, err := s.Encode("data", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    dataEncode,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}
