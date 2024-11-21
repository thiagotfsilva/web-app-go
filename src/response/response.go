package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErroResponse struct {
	Erro string `json:"erro"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}

}

// HandleStatusCode trata as requisições com status code 400 ou superior
func HandleStatusCode(w http.ResponseWriter, r *http.Response) {
	var erro ErroResponse
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}
