package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Imangali2002/trello-app/view"
)

func (app *App) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := app.userRepo.GetAllUsers()

	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, users)
}

func (app *App) getUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/users/"))

	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := app.userRepo.GetUserByID(uint(id))

	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, user)
}

func (app *App) createUser(w http.ResponseWriter, r *http.Request) {
	var user view.UserCreate

	// Parse the request body and decode it into the user struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	// Create the user using the repository
	err = app.userRepo.RegisterUser(user)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, nil)
}

func (app *App) userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if strings.HasPrefix(r.URL.Path, "/users/") {
			app.getUserByID(w, r)
		} else {
			app.getUsers(w, r)
		}
	case http.MethodPost:
		app.createUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
