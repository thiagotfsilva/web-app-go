package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"web-app-go/src/config"
	"web-app-go/src/cookies"
	"web-app-go/src/router"
	"web-app-go/src/utils"

	"github.com/gorilla/securecookie"
)

func init() {
	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(hashKey)

	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(blockKey)
}

func main() {
	config.Load()
	cookies.Config()
	utils.LoadTemplates()
	r := router.Router()
	fmt.Printf("Rodando na porta %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
