package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bimonugraraga/user-log-golang/collection"
	"github.com/bimonugraraga/user-log-golang/services"
	"github.com/julienschmidt/httprouter"
)

func Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	newUser := map[string]interface{}{}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	isTrue := services.CreateNewUser(newUser)
	if isTrue {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("JANCUKS")
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	allUser, err := services.GetAllUser()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(allUser)
	}
}

func LoginUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var logUser collection.LoggedUser
	err := json.NewDecoder(r.Body).Decode(&logUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	logBoy, err := services.Login(logUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(logBoy)
	}
}
