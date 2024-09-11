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

	router.HandleFunc("/transaction", handlers.GetTransaction).Methods("GET")
	router.HandleFunc("/transaction", handlers.CreateTransaction).Methods("POST")

	router.HandleFunc("/transactions", handlers.GetTransactions).Methods("GET")

	return router
}
