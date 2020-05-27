package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"clumio.com/roger/picture-perfect/webapp/model"
	"clumio.com/roger/picture-perfect/webapp/viewmodel"
)

type standLocator struct {
	standMapTemplate *template.Template
}

func (sl standLocator) registerRoutes() {
	http.HandleFunc("/stand-locator", sl.handleStandLocator)
	http.HandleFunc("/api/stands", sl.handleAPIStands)
}

func (sl standLocator) handleStandLocator(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewStandLocator()
	w.Header().Add("Content-Type", "text/html")
	sl.standMapTemplate.Execute(w, vm)
}

func (sl standLocator) handleAPIStands(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var loc struct {
		ZipCode string `json:"zipCode"`
	}
	err := dec.Decode(&loc)
	if err != nil {
		log.Println(fmt.Errorf("Error retrieving location: %v", err))
		enc := json.NewEncoder(w)
		enc.Encode([]model.StandCoordinate{})
		return
	}
	log.Println("location:", loc)
	vm := model.GetStandCoords()
	enc := json.NewEncoder(w)
	w.Header().Add("Content-Type", "application/json")
	enc.Encode(vm)
}
