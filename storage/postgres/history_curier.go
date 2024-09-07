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

type historyCurierRepo struct {
	db *pgxpool.Pool
}

func NewHistoryCurierRepo(db *pgxpool.Pool) storage.HistoryCurierRepoI {
	return &historyCurierRepo{
		db: db,
	}
}

func (r *historyCurierRepo) Create(ctx context.Context, req models.CreateHistoryCurier) (err error) {

	var (
		productid = uuid.NewString()
		query     = `
			INSERT INTO "history_curier"(
				"id",
				"curier_id",
				"order_id"
			) VALUES ($1, $2, $3)
		`
	)
	_, err = r.db.Exec(ctx,
		query,
		productid,
		req.CurierId,
		req.OrderId,
	)
	if err != nil {
		return err
	}

	return nil

}

func (r *historyCurierRepo) GetByID(ctx context.Context, req models.HistoryCurierPrimaryKey) (models.HistoryCurier, error) {
	var (
		query = `
			SELECT
				hc."id",
				hc."curier_id",
				o."total_price",
				o."status",
				o."delivery_address",
				o."payment_method",
				hc."created_at",
				hc."updated_at"
			FROM "history_curier" as hc
			JOIN "orders" as o ON o.id = hc.order_id
			WHERE "id"= $1 AND status = 'cancelled'
		`
		id               sql.NullString
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
		&curier_id,
		&total_price,
		&status,
		&delivery_address,
		&payment_method,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return models.HistoryCurier{}, err
	}

	return models.HistoryCurier{
		Id:              id.String,
		CurierId:        curier_id.String,
		TotalPrice:      total_price.String,
		Status:          status.String,
		DeliveryAddress: delivery_address.String,
		PaymentMethod:   payment_method.String,
		CreatedAt:       created_at.String,
		UpdatedAt:       updated_at.String,
	}, nil
}

func (r *historyCurierRepo) GetList(ctx context.Context, req models.GetListHistoryCurierRequest) (models.GetListHistoryCurierResponse, error) {
	var (
		resp   models.GetListHistoryCurierResponse
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
			hc."id",
			hc."curier_id",
			o."total_price",
			o."status",
			o."delivery_address",
			o."payment_method",
			hc."created_at",
			hc."updated_at"
		FROM "history_curier" as hc
		JOIN "orders" as o ON o.id = hc.order_id
	`

	query += where + sort + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return models.GetListHistoryCurierResponse{}, err
	}

	for rows.Next() {
		var (
			history_curier   models.HistoryCurier
			id               sql.NullString
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
			&curier_id,
			&total_price,
			&status,
			&delivery_address,
			&payment_method,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return models.GetListHistoryCurierResponse{}, err
		}

		history_curier = models.HistoryCurier{
			Id:              id.String,
			CurierId:        curier_id.String,
			TotalPrice:      total_price.String,
			Status:          status.String,
			DeliveryAddress: delivery_address.String,
			PaymentMethod:   payment_method.String,
			CreatedAt:       created_at.String,
			UpdatedAt:       updated_at.String,
		}

		resp.Histories = append(resp.Histories, history_curier)
	}

	return resp, nil
}
