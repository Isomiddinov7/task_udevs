package models

type CreateHistoryCurier struct {
	CurierId string `json:"curier_id"`
	OrderId  string `json:"order_id"`
}

type HistoryCurier struct {
	Id              string `json:"id"`
	CurierId        string `josn:"curier_id"`
	TotalPrice      string `json:"total_price"`
	Status          string `json:"status"`
	DeliveryAddress string `json:"delivery_address"`
	PaymentMethod   string `json:"payment_method"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type HistoryCurierPrimaryKey struct {
	Id string `json:"id"`
}

type GetListHistoryCurierRequest struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type GetListHistoryCurierResponse struct {
	Count     int64           `json:"count"`
	Histories []HistoryCurier `json:"histories"`
}
