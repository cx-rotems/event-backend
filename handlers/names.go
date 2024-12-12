package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"

	"EventBackend/database"
	"EventBackend/models"

	"github.com/gorilla/mux"
)

// GET /api/names
func GetNames(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start GetNames") 
	rows, err := database.DB.Query("SELECT id, firstName, lastName, arrived FROM names")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var names []models.Name
	for rows.Next() {
		var name models.Name
		err := rows.Scan(&name.ID, &name.FirstName, &name.LastName, &name.Arrived)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		names = append(names, name)
	}
	fmt.Println(names) 
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(names)
}

// POST /api/names
func AddName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start AddName") 
	var name models.Name
	if err := json.NewDecoder(r.Body).Decode(&name); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("received firstName %s, lastName %s", name.FirstName, name.LastName) 
	_, err := database.DB.Exec("INSERT INTO names (firstName, lastName) VALUES (?, ?)", name.FirstName, name.LastName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("added firstName %s, lastName %s", name.FirstName, name.LastName) 
	w.WriteHeader(http.StatusCreated)
}

// PUT /api/names/:id
func UpdateArrivedStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var name models.Name
	if err := json.NewDecoder(r.Body).Decode(&name); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("UPDATE names SET arrived = ? WHERE id = ?", name.Arrived, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
