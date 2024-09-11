package models

type Transaction struct {
	ID        int     `json:"id"`
	UserID    int     `json:"user_id"`
	Amount    float64 `json:"amount"`
	Category  string  `json:"category"`
	Type      string  `json:"type"`
	CreatedAt string  `json:"created_at"`
}
