package handlers

import (
	"net/http"
)

func (cfg *ApiConfig) handleGetExpenses(w http.ResponseWriter, r *http.Request) {

	expenses, err := cfg.DB.GetExpenses(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating expense", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, expenses)
}
