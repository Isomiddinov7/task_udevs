package postgres

import (
	"context"
	"database/sql"
	"task_udevs/api/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type cartRepo struct {
	db *pgxpool.Pool
}

func NewCartRepo(db *pgxpool.Pool) *cartRepo {
	return &cartRepo{
		db: db,
	}
}

func (r *cartRepo) Create(ctx context.Context, req models.CreateCart) (err error) {

	var (
		cartId = uuid.NewString()
		query  = `
			INSERT INTO "basket_user"(
				"id",
				"order_id",
				"user_id",
				"product_id"
			) VALUES ($1, $2, $3, $4)
		`
	)
	_, err = r.db.Exec(ctx,
		query,
		cartId,
		req.OrderId,
		req.UserId,
		req.ProductId,
	)
	if err != nil {
		return err
	}
	return nil

}

func (r *cartRepo) GetByID(ctx context.Context, req models.CartPrimaryKey) (resp models.Cart, err error) {
	var (
		query = `
			SELECT
				bu."id",
				o."total_price",
				o."delivery_address",
				o."payment_method",
				p."name",
				p."product_img"
			FROM "basket_user" as bu
			JOIN "orders" as o ON o.id = bu.order_id
			JOIN "products" as p ON p.id = bu.product_id
			WHERE id = $1
		`

		id               sql.NullString
		total_price      sql.NullString
		delivery_address sql.NullString
		payment_method   sql.NullString
		name             sql.NullString
		product_img      sql.NullString
	)

	err = r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&total_price,
		&delivery_address,
		&payment_method,
		&name,
		&product_img,
	)
	if err != nil {
		return models.Cart{}, err
	}

	return models.Cart{
		Id:              id.String,
		TotalPrice:      total_price.String,
		DeliveryAddress: delivery_address.String,
		PaymentMethod:   payment_method.String,
		ProductName:     name.String,
		ProductImg:      product_img.String,
	}, nil
}

func (r *cartRepo) Delete(ctx context.Context, req models.CartPrimaryKey) error {
	_, err := r.db.Exec(ctx, `DELETE FROM "product_img" WHERE id = $1`, req.Id)
	return err
}
