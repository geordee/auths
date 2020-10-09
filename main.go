package main

import (
	"log"
	"net/http"

	api "github.com/geordee/auths/api"
	config "github.com/geordee/auths/config"
)

func main() {
	config.InitializeDB()
	http.HandleFunc("/users/", api.Users)
	http.HandleFunc("/", api.Index)
	log.Fatal(http.ListenAndServe(":9095", nil))
}
