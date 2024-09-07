package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"task_udevs/api/models"
	"task_udevs/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type historyUserRepo struct {
	db *pgxpool.Pool
}

func NewHistoryUserRepo(db *pgxpool.Pool) storage.HistoryUserRepoI {
	return &historyUserRepo{
		db: db,
	}
}

func (r *historyUserRepo) GetByID(ctx context.Context, req models.HistoryUserPrimaryKey) (models.HistoryUser, error) {
	var (
		query = `
			SELECT
				hu."id",
				p."product_id",
				p."product_name",
				p."product_img",
				hu."created_at",
				hu."updated_at"
			FROM "history_user"
			JOIN "products" as p ON p.id = hu.product_id
			WHERE "id"= $1
		`
		id           sql.NullString
		product_id   sql.NullString
		product_name sql.NullString
		product_img  sql.NullString
		created_at   sql.NullString
		updated_at   sql.NullString
	)

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&product_id,
		&product_name,
		&product_img,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return models.HistoryUser{}, err
	}

	return models.HistoryUser{
		Id:          id.String,
		ProductId:   product_id.String,
		ProductName: product_name.String,
		ProductImg:  product_img.String,
		CreatedAt:   created_at.String,
		UpdatedAt:   updated_at.String,
	}, nil
}

func (r *historyUserRepo) GetList(ctx context.Context, req models.GetHistoryUserListRequest) (models.GetHistoryUserListResponse, error) {
	var (
		resp   models.GetHistoryUserListResponse
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

	if len(req.Search) > 0 {
		where += " AND p.product_name ILIKE" + " '%" + req.Search + "%'"
	}

	query := `
		SELECT
			COUNT(hu.*) OVER(),
			hu."id",
			hu."product_id",
			p."name",
			p."product_img",
			hu."created_at",
			hu."updated_at"
		FROM "history_user" as hu
		JOIN "products" as p ON p.id = hu.product_id

	`

	query += where + sort + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return models.GetHistoryUserListResponse{}, err
	}

	for rows.Next() {
		var (
			history      models.HistoryUser
			id           sql.NullString
			product_id   sql.NullString
			product_name sql.NullString
			product_img  sql.NullString
			created_at   sql.NullString
			updated_at   sql.NullString
		)
		err = rows.Scan(
			&resp.Count,
			&id,
			&product_id,
			&product_name,
			&product_img,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return models.GetHistoryUserListResponse{}, err
		}

		history = models.HistoryUser{
			Id:          id.String,
			ProductId:   product_id.String,
			ProductName: product_name.String,
			ProductImg:  product_img.String,
			CreatedAt:   created_at.String,
			UpdatedAt:   updated_at.String,
		}

		resp.Histories = append(resp.Histories, history)
	}

	return resp, nil
}
