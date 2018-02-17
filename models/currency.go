package models

// Currency struct represents currency
type Currency struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
