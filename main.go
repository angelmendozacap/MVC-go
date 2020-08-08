package main

import (
	"log"
	"net/http"

	"github.com/angelmendozacap/mvc/controllers"
)

func main() {
	http.HandleFunc("/posts", controllers.GetPosts)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
