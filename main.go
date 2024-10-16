package main

import (
	"fmt"
	"log"
	"net/http"
	"web-app-go/src/router"
	"web-app-go/src/utils"
)

func main() {
	utils.LoadTemplates()
	r := router.Router()
	fmt.Println("Rodando Webapp!")
	log.Fatal(http.ListenAndServe(":3000", r))
}
