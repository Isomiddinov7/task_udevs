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

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) storage.ProductRepoI {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(ctx context.Context, req models.CreateProduct) (resp models.Product, err error) {

	var (
		productid = uuid.NewString()
		query     = `
			INSERT INTO "products"(
				"id",
				"name",
				"comment",
				"price",
				"product_img"
			) VALUES ($1, $2, $3, $4, $5)
		`
	)
	_, err = r.db.Exec(ctx,
		query,
		productid,
		req.Name,
		req.Comment,
		req.Price,
		req.ProductImg,
	)
	if err != nil {
		return models.Product{}, err
	}

	return r.GetByID(ctx, models.ProductPrimaryKey{Id: productid})

}

func (r *productRepo) GetByID(ctx context.Context, req models.ProductPrimaryKey) (models.Product, error) {
	var (
		query = `
			SELECT
				"id",
				"name",
				"comment",
				"price",
				"product_img",
				"created_at",
				"updated_at"
			FROM "products"
			WHERE "id"= $1
		`
		id          sql.NullString
		name        sql.NullString
		comment     sql.NullString
		price       sql.NullString
		product_img sql.NullString
		created_at  sql.NullString
		updated_at  sql.NullString
	)

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&comment,
		&price,
		&product_img,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return models.Product{}, err
	}

	return models.Product{
		Id:         id.String,
		Name:       name.String,
		Comment:    comment.String,
		Price:      price.String,
		ProductImg: product_img.String,
		CreatedAt:  created_at.String,
		UpdatedAt:  updated_at.String,
	}, nil
}

func (r *productRepo) GetList(ctx context.Context, req models.GetProductListRequest) (models.GetProductListResponse, error) {
	var (
		resp   models.GetProductListResponse
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
		where += " AND name ILIKE" + " '%" + req.Search + "%'"
	}

	query := `
		SELECT
			COUNT(*) OVER(),
			"id",
			"name",
			"comment",
			"price",
			"product_img",
			"created_at",
			"updated_at"
		FROM "products"
	`

	query += where + sort + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return models.GetProductListResponse{}, err
	}

	for rows.Next() {
		var (
			product     models.Product
			id          sql.NullString
			name        sql.NullString
			comment     sql.NullString
			price       sql.NullString
			product_img sql.NullString
			created_at  sql.NullString
			updated_at  sql.NullString
		)
		err = rows.Scan(
			&resp.Count,
			&id,
			&name,
			&comment,
			&price,
			&product_img,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return models.GetProductListResponse{}, err
		}

		product = models.Product{
			Id:         id.String,
			Name:       name.String,
			Comment:    comment.String,
			Price:      price.String,
			ProductImg: product_img.String,
			CreatedAt:  created_at.String,
			UpdatedAt:  updated_at.String,
		}

		resp.Products = append(resp.Products, product)
	}

	return resp, nil
}

func (r *productRepo) Update(ctx context.Context, req models.UpdateProduct) (int64, error) {
	var (
		query = `
			UPDATE "products"
				SET
					"name" = $2,
					"comment" = $3,
					"price" = $4,
					"product_img" = $5,
					"updated_at" = NOW()
			WHERE "id" = $1
		`
	)

	rowsAffected, err := r.db.Exec(ctx,
		query,
		req.Id,
		req.Name,
		req.Comment,
		req.Price,
		req.ProductImg,
	)

	if err != nil {
		return 0, err
	}
	return rowsAffected.RowsAffected(), nil
}

func (r *productRepo) Delete(ctx context.Context, req models.ProductPrimaryKey) error {
	_, err := r.db.Exec(ctx, "DELETE FROM products WHERE id = $1", req.Id)
	return err
}
