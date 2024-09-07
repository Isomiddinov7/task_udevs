package models

type Product struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Comment    string `json:"comment"`
	Price      string `json:"price"`
	ProductImg string `json:"product_img"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type CreateProduct struct {
	Name       string `json:"name"`
	Comment    string `json:"comment"`
	Price      string `json:"price"`
	ProductImg string `json:"product_img"`
}

type UpdateProduct struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Comment    string `json:"comment"`
	Price      string `json:"price"`
	ProductImg string `json:"product_img"`
}

type ProductPrimaryKey struct {
	Id string `json:"id"`
}

type GetProductListRequest struct {
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
	Search string `json:"search"`
}

type GetProductListResponse struct {
	Count    int64     `json:"count"`
	Products []Product `json:"products"`
}
