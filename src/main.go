package main

import (
	"example.com/m/v2/src/router"
	"log"
	"net/http"
)

func main() {
	r := router.CreateRouter()
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}
