package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"task_udevs/api/models"
	"task_udevs/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type orderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) storage.OrderRepoI {
	return &orderRepo{
		db: db,
	}
}

func (r *orderRepo) Create(ctx context.Context, req models.CreateOrder) (resp models.Order, err error) {

	var (
		orderId = uuid.NewString()
		query   = `
			INSERT INTO "orders"(
				"id",
				"user_id",
				"curier_id",
				"total_price",
				"status",
				"delivery_address",
				"payment_method"
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		`
	)
	_, err = r.db.Exec(ctx,
		query,
		orderId,
		req.UserId,
		req.CurierId,
		req.TotalPrice,
		req.Status,
	)
	if err != nil {
		return models.Order{}, err
	}

	var (
		historyId    = uuid.NewString()
		queryhistory = `
			INSERT INTO "history_user"(
				"id",
				"user_id",
				"order_id"
			) VALUES ($1, $2, $3)
		`
	)
	_, err = r.db.Exec(ctx,
		queryhistory,
		historyId,
		req.UserId,
		orderId,
	)
	if err != nil {
		return models.Order{}, err
	}

	return r.GetByID(ctx, models.OrderPrimaryKey{Id: orderId})

}

func (r *orderRepo) GetByID(ctx context.Context, req models.OrderPrimaryKey) (models.Order, error) {
	var (
		query = `
			SELECT
				"id",
				"user_id",
				"curier_id",
				"total_price",
				"status",
				"delivery_address",
				"payment_method",
				"created_at",
				"updated_at"
			FROM "orders"
			WHERE "id"= $1
		`
		id               sql.NullString
		user_id          sql.NullString
		curier_id        sql.NullString
		total_price      sql.NullString
		status           sql.NullString
		delivery_address sql.NullString
		payment_method   sql.NullString
		created_at       sql.NullString
		updated_at       sql.NullString
	)

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&user_id,
		&curier_id,
		&total_price,
		&status,
		&delivery_address,
		&payment_method,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return models.Order{}, err
	}

	return models.Order{
		Id:              id.String,
		UserId:          user_id.String,
		CurierId:        curier_id.String,
		TotalPrice:      total_price.String,
		Status:          status.String,
		DeliveryAddress: delivery_address.String,
		PaymentMethod:   payment_method.String,
		CreatedAt:       created_at.String,
		UpdatedAt:       updated_at.String,
	}, nil
}

func (r *orderRepo) GetList(ctx context.Context, req models.GetListOrderRequest) (models.GetListOrderResponse, error) {
	var (
		resp   models.GetListOrderResponse
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		sort   = " ORDER BY created_at DESC"
	)

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query := `
		SELECT
			COUNT(*) OVER(),
			"id",
			"user_id",
			"curier_id",
			"total_price",
			"status",
			"delivery_address",
			"payment_method",
			"created_at",
			"updated_at"
		FROM "orders"
		WHERE "status" = 'pending'
	`

	query += where + sort + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return models.GetListOrderResponse{}, err
	}

	for rows.Next() {
		var (
			order            models.Order
			id               sql.NullString
			user_id          sql.NullString
			curier_id        sql.NullString
			total_price      sql.NullString
			status           sql.NullString
			delivery_address sql.NullString
			payment_method   sql.NullString
			created_at       sql.NullString
			updated_at       sql.NullString
		)
		err = rows.Scan(
			&resp.Count,
			&id,
			&user_id,
			&curier_id,
			&total_price,
			&status,
			&delivery_address,
			&payment_method,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return models.GetListOrderResponse{}, err
		}

		order = models.Order{
			Id:              id.String,
			UserId:          user_id.String,
			CurierId:        curier_id.String,
			TotalPrice:      total_price.String,
			Status:          status.String,
			DeliveryAddress: delivery_address.String,
			PaymentMethod:   payment_method.String,
			CreatedAt:       created_at.String,
			UpdatedAt:       updated_at.String,
		}

		resp.Orders = append(resp.Orders, order)
	}

	return resp, nil
}

func (r *orderRepo) Update(ctx context.Context, req models.UpdateOrder) (int64, error) {
	var (
		query = `
			UPDATE "orders"
				SET
					"status" = $2,
					"updated_at" = NOW()
			WHERE "id" = $1
		`
	)

	rowsAffected, err := r.db.Exec(ctx,
		query,
		req.Status,
	)

	if err != nil {
		return 0, err
	}
	return rowsAffected.RowsAffected(), nil
}
