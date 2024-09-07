package storage

import (
	"context"
	"task_udevs/api/models"
)

type StorageI interface {
	User() UserRepoI
	Curier() CurierRepoI
	Product() ProductRepoI
	HistoryUser() HistoryUserRepoI
	AdditionProduct() AdditionProductRepoI
	Order() OrderRepoI
	Cart() CartRepoI
	HistoryCurier() HistoryCurierRepoI
}

type UserRepoI interface {
	Auth(ctx context.Context, req models.UserAuthRequest) (resp models.UserAuthResponse, err error)
	DeserializeUser(ctx context.Context, req models.GetUserById) (err error)
}

type CurierRepoI interface {
	Auth(ctx context.Context, req models.CurierAuthRequest) (resp models.CurierAuthResponse, err error)
	DeserializeCurier(ctx context.Context, req models.GetCurierById) (err error)
}

type ProductRepoI interface {
	Create(ctx context.Context, req models.CreateProduct) (resp models.Product, err error)
	GetByID(ctx context.Context, req models.ProductPrimaryKey) (resp models.Product, err error)
	GetList(ctx context.Context, req models.GetProductListRequest) (resp models.GetProductListResponse, err error)
	Update(ctx context.Context, req models.UpdateProduct) (int64, error)
	Delete(ctx context.Context, req models.ProductPrimaryKey) error
}

type HistoryUserRepoI interface {
	GetByID(ctx context.Context, req models.HistoryUserPrimaryKey) (resp models.HistoryUser, err error)
	GetList(ctx context.Context, req models.GetHistoryUserListRequest) (resp models.GetHistoryUserListResponse, err error)
}

type AdditionProductRepoI interface {
	Create(ctx context.Context, req models.CreateAdditionProduct) (resp string, err error)
	GetByID(ctx context.Context, req models.GetAdditionProductById) (resp models.GetAdditionProductByIdResponse, err error)
}

type OrderRepoI interface {
	Create(ctx context.Context, req models.CreateOrder) (resp models.Order, err error)
	GetByID(ctx context.Context, req models.OrderPrimaryKey) (resp models.Order, err error)
	GetList(ctx context.Context, req models.GetListOrderRequest) (resp models.GetListOrderResponse, err error)
	Update(ctx context.Context, req models.UpdateOrder) (int64, error)
}

type CartRepoI interface {
	Create(ctx context.Context, req models.CreateCart) (err error)
	GetByID(ctx context.Context, req models.CartPrimaryKey) (resp models.Cart, err error)
	Delete(ctx context.Context, req models.CartPrimaryKey) (err error)
}

type HistoryCurierRepoI interface {
	Create(ctx context.Context, req models.CreateHistoryCurier) (err error)
	GetByID(ctx context.Context, req models.HistoryCurierPrimaryKey) (resp models.HistoryCurier, err error)
	GetList(ctx context.Context, req models.GetListHistoryCurierRequest) (resp models.GetListHistoryCurierResponse, err error)
}
