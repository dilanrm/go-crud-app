package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func peopleHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
		case "GET":
			q := r.URL.Query().Get("id")
			if q == "" {
                handleGetPeople(w, r)
            } else {
                handleGetPerson(w, r)
            }
			// handleGetPeople(w, r)
		case "POST":
			handleCreatePerson(w, r)
		case "PUT":
			handleUpdatePerson(w, r)
		case "DELETE":
			handleDeletePerson(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleGetPerson(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
    if idStr == "" {
        http.Error(w, "ID is required", http.StatusBadRequest)
        return
    }
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
	person := store.GetPerson(id)

	if person.ID == 0 {
		person = Person{}
	}
	json.NewEncoder(w).Encode(person)
}

func handleGetPeople(w http.ResponseWriter, r *http.Request){
	people := store.GetAllPeople()
	json.NewEncoder(w).Encode(people)
}

func handleCreatePerson(w http.ResponseWriter, r *http.Request) {
    var p Person
    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    createdPerson := store.CreatePerson(p)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createdPerson)
}

func handleUpdatePerson(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    if idStr == "" {
        http.Error(w, "ID is required", http.StatusBadRequest)
        return
    }
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    var p Person
    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if store.UpdatePerson(id, p) {
        w.WriteHeader(http.StatusNoContent)
    } else {
        http.Error(w, "Person not found", http.StatusNotFound)
    }
}

func handleDeletePerson(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    if idStr == "" {
        http.Error(w, "ID is required", http.StatusBadRequest)
        return
    }
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    if store.DeletePerson(id) {
        w.WriteHeader(http.StatusNoContent)
    } else {
        http.Error(w, "Person not found", http.StatusNotFound)
    }
}