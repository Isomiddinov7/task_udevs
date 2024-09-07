package models

type HistoryUser struct {
	Id          string `json:"id"`
	ProductId   string `json:"product_id"`
	OrderId     string `json:"order_id"`
	ProductName string `json:"product_name"`
	ProductImg  string `json:"product_img"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateHistoryUser struct {
	ProductId   string `json:"product_id"`
	OrderId     string `json:"order_id"`
	ProductName string `json:"product_name"`
	ProductImg  string `json:"product_img"`
}

type HistoryUserPrimaryKey struct {
	Id string `json:"id"`
}

type GetHistoryUserListRequest struct {
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
	Search string `json:"search"`
}

type GetHistoryUserListResponse struct {
	Count     int64         `json:"count"`
	Histories []HistoryUser `json:"histories"`
}
