package models

type WSRecordCommon struct {
	ID        string  `json:"id" db:"id"`
	Name      string  `json:"name" db:"name" binding:"required"`
	Value     float32 `json:"value" db:"value" binding:"required"`
	Timestamp string  `json:"timestamp" db:"timestamp"`
}

type WSRecordPost struct {
	WSRecordCommon
	WSourceID  string `json:"ws_id" db:"ws_id" binding:"required"`
	CurrencyID string `json:"currency_id" db:"currency_id" binding:"required"`
	UserID     string `json:"user_id" db:"user_id" binding:"required"`
}

type WSRecordGet struct {
	WSRecordCommon
	Currency string `json:"currency" db:"currency_name"`
}
