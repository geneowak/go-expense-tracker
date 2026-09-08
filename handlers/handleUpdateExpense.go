package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/geneowak/go-expense-tracker/internal/database"
	"github.com/google/uuid"
)

type updateExpenseRequest struct {
	ItemName     string `json:"item_name" validate:"required,alphanum,min=2"`
	CategoryName string `json:"category_name" validate:"required,alphanum,min=2"`
	Quantity     string `json:"quantity" validate:"required,number,min=1"`
	UnitCost     string `json:"unit_cost" validate:"required,number,min=1"`
}

func (cfg *ApiConfig) handleUpdateExpense(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("expense_id")
	expenseId, err := uuid.Parse(idParam)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid expense id", err)
		return
	}

	var req updateExpenseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if err := cfg.Validate.Struct(req); err != nil {
		handleValidationErrors(w, err)
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
	// we've already validated that these are numbers so we don't expect errors
	quantity, _ := strconv.Atoi(req.Quantity)
	unitCost, _ := strconv.Atoi(req.UnitCost)

	updatedExpense, err := cfg.DB.UpdateExpense(r.Context(), database.UpdateExpenseParams{
		ItemName:     req.ItemName,
		CategoryName: req.CategoryName,
		Quantity:     int32(quantity),
		UnitCost:     int32(unitCost),
		ID:           expenseId,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error updating expense", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, updatedExpense)
}
