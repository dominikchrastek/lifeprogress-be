package models

// WSourceCommon is struct that contains common stuff
// for another WSource structs
type WSourceCommon struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Type string `json:"type" db:"ws_type"`
}

// WSource is struct that contains WSourceCommon + CurrencyID
type WSource struct {
	WSourceCommon
	CurrencyID string `json:"currency_id" db:"currency_id"`
}

// WSourceC is struct that contains WSourceCommon + Currencies
type WSourceC struct {
	WSourceCommon
	Currencies []Currency `json:"currencies"`
}
