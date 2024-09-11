package handlers

import (
	"MVPFinanceApp/api/models"
	"MVPFinanceApp/db"
	"encoding/json"
	"net/http"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var t models.Transaction

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	sqlStatement := `INSERT INTO transactions (user_id, amount, category, type, created_at)
					 VALUES ($1, $2, $3, $4, NOW()) RETURNING id`

	err = db.DB.QueryRow(sqlStatement, t.UserID, t.Amount, t.Category, t.Type).Scan(&t.ID)
	if err != nil {
		http.Error(w, "Failed to create transaction", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	rows, err := db.DB.Query("SELECT id, amount, category, type, created_at FROM transactions WHERE user_id = $1", userID)
	if err != nil {
		http.Error(w, "Failed to fetch transactions", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var t models.Transaction
		if err := rows.Scan(&t.ID, &t.Amount, &t.Category, &t.Type, &t.CreatedAt); err != nil {
			http.Error(w, "Failed to parse transactions", http.StatusInternalServerError)
			return
		}
		transactions = append(transactions, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}
