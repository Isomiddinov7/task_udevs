package models

type CreateCart struct {
	OrderId   string `json:"order_id"`
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
}

type CartPrimaryKey struct {
	Id string `json:"id"`
}

type Cart struct {
	Id              string `json:"id"`
	TotalPrice      string `json:"total_price"`
	DeliveryAddress string `json:"delivery_address"`
	PaymentMethod   string `json:"payment_method"`
	ProductName     string `json:"product_name"`
	ProductImg      string `json:"product_img"`
}
