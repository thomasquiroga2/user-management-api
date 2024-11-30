package main

import (
	"log"
	"net/http"
	"user-management-api/routes"
)

func main() {
	r := routes.SetupRoutes()

	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
