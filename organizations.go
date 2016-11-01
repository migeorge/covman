package main

import (
	"net/http"
	"time"

	"github.com/google/jsonapi"
	"github.com/julienschmidt/httprouter"
)

// Organization is a data model to be used with the database
type Organization struct {
	ID        uint       `jsonapi:"primary,organizations" gorm:"primary_key"`
	Name      string     `jsonapi:"attr,name"`
	CreatedAt time.Time  `jsonapi:"attr,createdAt"`
	UpdatedAt time.Time  `jsonapi:"attr,updatedAt"`
	DeletedAt *time.Time `sql:"index"`
}

func (h *Handlers) organizationsIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/vnd.api+json")

	organizations := []Organization{}
	h.DB.Find(&organizations)

	w.WriteHeader(http.StatusOK)

	organizationInterfaces := []interface{}{}
	for i := 0; i < len(organizations); i++ {
		organizationInterfaces = append(organizationInterfaces, &organizations[i])
	}

	err := jsonapi.MarshalManyPayload(w, organizationInterfaces)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handlers) organizationByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/vnd.api+json")

	organization := Organization{}
	h.DB.First(&organization, p.ByName("id"))

	if organization.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	err := jsonapi.MarshalOnePayload(w, &organization)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handlers) organizationCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/vnd.api+json")

	organization := Organization{}

	err := jsonapi.UnmarshalPayload(r.Body, &organization)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	status := h.DB.NewRecord(organization)
	if !status {
		http.Error(w, "Unable to process request", http.StatusUnprocessableEntity)
		return
	}

	h.DB.Create(&organization)

	w.WriteHeader(http.StatusCreated)

	err = jsonapi.MarshalOnePayload(w, &organization)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
