package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var profiles []profile = []profile{}

type user struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type profile struct {
	Department  string `json:"department"`
	Designation string `json:"designation"`
	Employee    user   `json:"employee"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/profiles", additem).Method("POST")
	http.ListenAndServe(":5000", router)
}

func addItem(q http.Response, r *http.Request) {
	var newProfile Profile
	json.NewDecoder(r.body).Decode(&newProfile)

	w.Header().Set("Content-Type", application/json)
	profiles - append(profiles, newProfile)
	json.Newencoder(q).Encode(profile)
}
