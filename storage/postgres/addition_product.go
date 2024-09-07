package postgres

import (
	"context"
	"database/sql"
	"task_udevs/api/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type additionProductRepo struct {
	db *pgxpool.Pool
}

func NewAdditionProductRepo(db *pgxpool.Pool) *additionProductRepo {
	return &additionProductRepo{
		db: db,
	}
}

func (r *additionProductRepo) Create(ctx context.Context, req models.CreateAdditionProduct) (resp string, err error) {
	var (
		query = `
			INSERT INTO "addition"(
				"product_id",
				"thing",
				"thing_price"
			) VALUES($1, $2, $3)
		`
	)

	_, err = r.db.Exec(ctx, query,
		req.ProductId,
		req.Thing,
		req.ThingPrice,
	)
	if err != nil {
		return "", err
	}
	return "created addition product", nil
}

func (r *additionProductRepo) GetByID(ctx context.Context, req models.GetAdditionProductById) (models.GetAdditionProductByIdResponse, error) {
	var (
		resp models.GetAdditionProductByIdResponse
	)

	query := `
		SELECT
			COUNT(*) OVER(),
			"product_id",
			"thing",
			"thing_price",
			"created_at",
			"updated_at"
		FROM "addition"
		WHERE "product_id" = $1
	`

	rows, err := r.db.Query(ctx, query, req.ProductId)
	if err != nil {
		return models.GetAdditionProductByIdResponse{}, err
	}

	for rows.Next() {
		var (
			addtion_product models.AdditionProduct
			product_id      sql.NullString
			thing           sql.NullString
			thing_price     sql.NullString
			created_at      sql.NullString
			updated_at      sql.NullString
		)
		err = rows.Scan(
			&resp.Count,
			&product_id,
			&thing,
			&thing_price,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return models.GetAdditionProductByIdResponse{}, err
		}

		addtion_product = models.AdditionProduct{
			ProductId:  product_id.String,
			Thing:      thing.String,
			ThingPrice: thing_price.String,
			CreatedAt:  created_at.String,
			UpdatedAt:  updated_at.String,
		}

		resp.Additions = append(resp.Additions, addtion_product)
	}

	return resp, nil
}
