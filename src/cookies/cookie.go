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

func SaveCookies(w http.ResponseWriter, id, token string) error {
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

func ReadCookies(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("data")
	if err != nil {
		return nil, err
	}

	valuesCookies := make(map[string]string)
	if err = s.Decode("data", cookie.Value, &valuesCookies); err != nil {
		return nil, err
	}

	return valuesCookies, nil
}
