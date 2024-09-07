package postgres

import (
	"context"
	"database/sql"
	"task_udevs/api/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type curierRepo struct {
	db *pgxpool.Pool
}

func NewCurierRepo(db *pgxpool.Pool) *curierRepo {
	return &curierRepo{
		db: db,
	}
}

func (r *curierRepo) Auth(ctx context.Context, req models.CurierAuthRequest) (resp models.CurierAuthResponse, err error) {
	var (
		query = `
			SELECT 
				id
			FROM "curiers"
			WHERE "email" = $1 AND "password" = $2
		`

		id sql.NullString
	)
	err = r.db.QueryRow(ctx, query, req.Email, req.Password).Scan(
		&id,
	)
	if err != nil {
		return models.CurierAuthResponse{}, err
	}

	return models.CurierAuthResponse{
		Success:  "Success",
		CurierId: id.String,
	}, nil
}

func (r *curierRepo) DeserializeCurier(ctx context.Context, req models.GetCurierById) (err error) {
	var (
		query = `
			SELECT 
				id
			FROM "curiers"
			WHERE "id" = $1
		`

		id sql.NullString
	)
	err = r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
	)
	if err != nil {
		return err
	}

	return nil
}
