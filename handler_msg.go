package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/smirankat/messaggio-test/internal/database"
)

func (apiCfg *apiConfig) handlerCreateMessage(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Message string `json:"message"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	message, err := apiCfg.DB.CreateMessage(r.Context(), database.CreateMessageParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Message:   params.Message,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create message: %v", err))
		return
	}

	respondWithJSON(w, 200, message)
}
