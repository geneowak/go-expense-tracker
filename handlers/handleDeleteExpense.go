package handlers

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func (cfg *ApiConfig) handleDeleteExpense(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("expense_id")
	expenseId, err := uuid.Parse(idParam)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid expense id", err)
		return
	}

	// get the expense and make sure that the user owns it
	expense, err := cfg.DB.GetExpenseById(r.Context(), expenseId)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			respondWithError(w, http.StatusNotFound, "Expense not found", err)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch expense", err)
		return
	}

	userId, err := UserIdFromContext(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Invalid user id", err)
		return
	}

	if expense.UserID != userId {
		respondWithError(w, http.StatusForbidden, "You don't own this expense.", err)
		return
	}
	err = cfg.DB.DeleteExpense(r.Context(), expenseId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error deleting expense", err)
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
