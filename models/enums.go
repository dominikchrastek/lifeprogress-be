package models

// Unit weight unit
type Unit struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
}

// WSType types
type WSType struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
}
