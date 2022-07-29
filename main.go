package main

import (
	"log"
	"net/http"
	"os"

	"github.com/c3duan/Crypto-Trade-API/api"
	"github.com/c3duan/Crypto-Trade-API/src/middleware/helper"
)

func main() {
	mainRoute := api.New()
	var addrs string = "0.0.0.0:8080"

	if pr := os.Getenv("PORT"); pr != "" {
		addrs = "0.0.0.0:" + pr
	}

	log.Println("App running on " + addrs)

	if err := http.ListenAndServe(addrs, helper.WithSlashTrimming(mainRoute)); err != nil {
		log.Fatal(err)
	}
}
