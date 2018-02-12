package models

// Life stuff

// Unit weight unit
type Unit struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// Asset types
type Asset struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
