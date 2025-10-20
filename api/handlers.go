package api

import (
	"encoding/json"
	"net/http"

	"github.com/ducthinh993/renovate-gomodtidyall-test/shared"
	"github.com/gorilla/mux"
)

// Handler uses shared library
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := shared.NewResponse("ok", "API is healthy")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// SetupRoutes configures the API router
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/health", HealthHandler).Methods("GET")
	return router
}