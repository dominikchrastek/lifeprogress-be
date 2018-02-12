package models

type Weight struct {
	ID        string `json:"id" db:"id"`
	Value     int    `json:"value" db:"value"`
	Unit      string `json:"unit" db:"unit"`
	Timestamp string `json:"timestamp" db:"timestamp"`
}

type WeightConnectUser struct {
	ID       int    `json:"id" db:"id"`
	UserID   string `json:"user_id" db:"user_id"`
	WeightID string `json:"weight_id" db:"weight_id"`
}
