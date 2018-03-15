package models

// Weight struct represents weight
type Weight struct {
	ID        string  `json:"id" db:"id"`
	Value     float32 `json:"value" db:"value" binding:"required"`
	Unit      string  `json:"unit" db:"unit" binding:"required"`
	Timestamp string  `json:"timestamp" db:"timestamp"`
}
