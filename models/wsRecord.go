package models

type WSRecordCommon struct {
	ID        string  `json:"id" db:"id"`
	Value     float32 `json:"value" db:"value" binding:"required"`
	Timestamp string  `json:"timestamp" db:"timestamp"`
}

type WSRecordPost struct {
	WSRecordCommon
	WSourceID  string `json:"wsourceId" db:"ws_id" binding:"required"`
	CurrencyID string `json:"currencyId" db:"currency_id" binding:"required"`
	UserID     string `json:"userId" db:"user_id" binding:"required"`
}

type WSRecordGet struct {
	WSRecordCommon
	Currency string `json:"currency" db:"currency_name"`
}
