package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/geneowak/go-expense-tracker/internal/database"
)

type DateFilter string

const (
	FilterPastWeek    DateFilter = "past_week"
	FilterPastMonth   DateFilter = "past_month"
	FilterPast3Months DateFilter = "past_3_months"
	FilterCustom      DateFilter = "custom"
)

type getExpensesRequest struct {
	Filter    DateFilter `json:"filter" validate:"omitempty,oneof=past_week past_month, past_3_months, custom"`
	StartDate *time.Time `json:"start_date" validate:"datetime,required_if=filter custom,omitempty"`
	EndDate   *time.Time `json:"end_date" validate:"datetime,required_if=filter custom,gtfield=start_date,omitempty"`
}

func (cfg *ApiConfig) handleGetExpenses(w http.ResponseWriter, r *http.Request) {

	startDate, endDate, err := getTimeRange(r)

	expenses, err := cfg.DB.GetExpenses(r.Context(), database.GetExpensesParams{
		StartDate: sql.NullTime{Time: startDate, Valid: true},
		EndDate:   sql.NullTime{Time: endDate, Valid: true},
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating expense", err)
		return
	}

	respondWithJSON(w, http.StatusOK, expenses)
}

func getTimeRange(r *http.Request) (time.Time, time.Time, error) {
	filter := DateFilter(r.URL.Query().Get("filter"))

	if filter == "" {
		filter = FilterPastWeek
	}

	now := time.Now().UTC()
	endDate := now
	var startDate time.Time

	switch filter {
	case FilterPastWeek:
		startDate = now.AddDate(0, 0, -7)
	case FilterPastMonth:
		startDate = now.AddDate(0, -1, 0)
	case FilterPast3Months:
		startDate = now.AddDate(0, -3, 0)
	case FilterCustom:
		rawStart := r.URL.Query().Get("start_date")
		rawEnde := r.URL.Query().Get("end_date")
		startDate = now.AddDate(0, -3, 0)
		// startDate = *req.StartDate
		// endDate = *req.EndDate
	default:
		// default to past week
		startDate = now.AddDate(0, 0, -7)
	}

	return startDate, endDate
}
