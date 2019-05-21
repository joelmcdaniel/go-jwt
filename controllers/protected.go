package controllers

import (
	"database/sql"
	"jwt-auth-restapi/utils"
	"net/http"
)

// Controller ...
type Controller struct{}

// ProtectedEndpoint ...
func (c Controller) ProtectedEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ResponseJSON(w, "Yes")
	}
}
