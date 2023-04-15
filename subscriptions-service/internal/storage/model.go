package storage

import "time"

type Ticker struct {
	TickerId           string    `json:"ticker_id"`
	CompanyName        string    `json:"company_name"`
	CompanyLocale      string    `json:"company_locale"`
	CurrencyName       string    `json:"currency_name"`
	CompanyDescription string    `json:"company_description"`
	HomepageUrl        string    `json:"homepage_url"`
	PhoneNumber        string    `json:"phone_number"`
	TotalEmployees     int       `json:"total_employees"`
	CompanyState       string    `json:"company_state"`
	CompanyCity        string    `json:"company_city"`
	CompanyAddress     string    `json:"company_address"`
	CompanyPostalCode  string    `json:"company_postal_code"`
	TickerCik          string    `json:"ticker_cik"`
	Active             bool      `json:"active"`
	CreatedAt          time.Time `json:"created_at"`
	ExternalUpdatedAt  time.Time `json:"external_updated_at"`
}

type Stock struct {
	StockId       string    `json:"stock_id"`
	TickerId      string    `json:"ticker_id"`
	OpenPrice     float64   `json:"open_price"`
	ClosePrice    float64   `json:"close_price"`
	HighestPrice  float64   `json:"highest_price"`
	LowestPrice   float64   `json:"lowest_price"`
	TradingVolume float64   `json:"trading_volume"`
	StockedAt     time.Time `json:"stocked_time"`
	CreatedAt     time.Time `json:"created_at"`
}
