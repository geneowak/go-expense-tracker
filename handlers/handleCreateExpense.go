package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/geneowak/go-expense-tracker/internal/database"
)

type createExpenseRequest struct {
	ItemName     string `json:"item_name" validate:"required,alphanum,min=2"`
	CategoryName string `json:"category_name" validate:"required,alphanum,min=2"`
	Quantity     string `json:"quantity" validate:"required,number,min=1"`
	UnitCost     string `json:"unit_cost" validate:"required,number,min=1"`
}

func (cfg *ApiConfig) handleCreateExpense(w http.ResponseWriter, r *http.Request, authUser database.User) {
	var req createExpenseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if err := cfg.Validate.Struct(req); err != nil {
		handleValidationErrors(w, err)
		return
	}

	// we've already validated that these are numbers so we don't expect errors
	quantity, _ := strconv.Atoi(req.Quantity)
	unitCost, _ := strconv.Atoi(req.UnitCost)

	expense, err := cfg.DB.CreateExpense(r.Context(), database.CreateExpenseParams{
		ItemName:     req.ItemName,
		CategoryName: req.CategoryName,
		Quantity:     int32(quantity),
		UnitCost:     int32(unitCost),
		UserID:       authUser.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating expense", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, expense)
}
