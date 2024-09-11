package api

import (
	"MVPFinanceApp/api/handlers"
	"MVPFinanceApp/api/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", handlers.Login).Methods("POST")

	// Protected routes
	router.Use(middleware.AuthMiddleware)

	router.HandleFunc("/transactions", handlers.CreateTransaction).Methods("POST")
	router.HandleFunc("/transactions", handlers.GetTransactions).Methods("GET")

	return router
}
