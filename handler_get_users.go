package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (apiConfig *apiConfig)handlerGetUsers(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Id uuid.UUID `json:"id"`
	}

	decoder := json.NewDecoder(r.Body)
	parms := parameters{}

	err := decoder.Decode(&parms)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	getUser, er := apiConfig.DB.GetUser(r.Context(), parms.Id)

	if er != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	respondWithJson(w, http.StatusOK, databaseUserToUser(getUser))

}
