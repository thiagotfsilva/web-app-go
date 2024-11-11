package request

import (
	"io"
	"net/http"
	"web-app-go/src/cookies"
)

// HandlerRequestAuthenticate é utilizado para colocar o token na requisição
func HandlerRequestAuthenticate(
	r *http.Request,
	method, url string,
	data io.Reader,
) (*http.Response, error) {
	request, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	cookie, _ := cookies.ReadCookies(r)
	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
