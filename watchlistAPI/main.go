package main

import (
	"fmt"
	"log"
	"net/http"
	"watchlistAPI/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":9090", r))
	fmt.Println("Listening at port 9090...")
}
