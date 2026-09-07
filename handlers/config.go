package handlers

import "github.com/geneowak/go-expense-tracker/internal/database"

type ApiConfig struct {
	DB        *database.Queries
	Platform  string
	JwtSecret string
}
