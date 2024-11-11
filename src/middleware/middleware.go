package middleware

import (
	"log"
	"net/http"
	"web-app-go/src/cookies"
)

func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

func Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := cookies.ReadCookies(r); err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		nextFunc(w, r)
	}
}
