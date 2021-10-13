package main

import (
	"fmt"
	"log"
	"net/http"
	"posts-crud-golang/controllers"
)

func main() {
	r := controllers.Register()

	fmt.Println("Server is running on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
