package main

import (
	"fmt"
	"log"
	"net/http"
	"web-app-go/src/config"
	"web-app-go/src/cookies"
	"web-app-go/src/router"
	"web-app-go/src/utils"
)

func main() {
	config.Load()
	cookies.Config()
	fmt.Println(config.Port)
	utils.LoadTemplates()
	r := router.Router()
	fmt.Printf("Rodando na porta %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
