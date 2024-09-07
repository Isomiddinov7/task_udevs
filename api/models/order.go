package models

type CreateOrder struct {
	UserId          string `json:"user_id"`
	CurierId        string `json:"curier_id"`
	ProductId       string `json:"product_id"`
	TotalPrice      string `json:"total_price"`
	Status          string `json:"status"`
	DeliveryAddress string `json:"delivery_address"`
	PaymentMethod   string `json:"payment_methon"`
}

type Order struct {
	Id              string `json:"id"`
	UserId          string `json:"user_id"`
	CurierId        string `json:"curier_id"`
	ProductId       string `json:"product_id"`
	TotalPrice      string `json:"total_price"`
	Status          string `json:"status"`
	DeliveryAddress string `json:"delivery_address"`
	PaymentMethod   string `json:"payment_methon"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type UpdateOrder struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

type OrderPrimaryKey struct {
	Id string `json:"id"`
}

type GetListOrderRequest struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type GetListOrderResponse struct {
	Count  int64   `json:"count"`
	Orders []Order `json:"orders"`
}
