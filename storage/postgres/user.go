package postgres

import (
	"context"
	"database/sql"
	"task_udevs/api/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Auth(ctx context.Context, req models.UserAuthRequest) (resp models.UserAuthResponse, err error) {
	var (
		query = `
			SELECT 
				id
			FROM "users"
			WHERE "email" = $1 AND "password" = $2
		`

		id sql.NullString
	)
	err = r.db.QueryRow(ctx, query, req.Email, req.Password).Scan(
		&id,
	)
	if err != nil {
		return models.UserAuthResponse{}, err
	}

	return models.UserAuthResponse{
		Success: "Success",
		UserId:  id.String,
	}, nil
}

func (r *userRepo) DeserializeUser(ctx context.Context, req models.GetUserById) (err error) {
	var (
		query = `
			SELECT 
				id
			FROM "users"
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
