package models

type AdditionProduct struct {
	ProductId  string `json:"product_id"`
	OrderId    string `json:"order_id"`
	Thing      string `json:"thing"`
	ThingPrice string `json:"thing_price"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type CreateAdditionProduct struct {
	ProductId  string `json:"product_id"`
	Thing      string `json:"thing"`
	ThingPrice string `json:"thing_price"`
}

type GetAdditionProductById struct {
	ProductId string `json:"product_id"`
}

type GetAdditionProductByIdResponse struct {
	Count     int64             `json:"count"`
	Additions []AdditionProduct `json:"additions"`
}
