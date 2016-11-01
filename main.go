package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// FindOrganization should return a single Organization found by the ID in the route
func FindOrganization(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	db := DB{}
	db.init()
	db.initSchema()

	handlers := Handlers{db.DB}

	router := httprouter.New()

	router.GET("/organizations", handlers.organizationsIndex)
	router.POST("/organizations", handlers.organizationCreate)
	router.GET("/organizations/:id", handlers.organizationByID)

	port := ":8080"
	log.Printf("Serving on port: %v", port)
	log.Fatal(http.ListenAndServe(port, router))
}
