package postgres

import (
	"context"
	"fmt"
	"task_udevs/config"
	"task_udevs/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db               *pgxpool.Pool
	user             storage.UserRepoI
	curier           storage.CurierRepoI
	product          storage.ProductRepoI
	history_user     storage.HistoryUserRepoI
	addition_product storage.AdditionProductRepoI
	order            storage.OrderRepoI
	cart             storage.CartRepoI
	history_curier   storage.HistoryCurierRepoI
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {

	config, err := pgxpool.ParseConfig(
		fmt.Sprintf(
			"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
			cfg.PostgresHost,
			cfg.PostgresUser,
			cfg.PostgresDatabase,
			cfg.PostgresPassword,
			cfg.PostgresPort,
		),
	)

	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnection

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pgxpool,
	}, nil
}

func (s *Store) User() storage.UserRepoI {

	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}

	return s.user
}

func (s *Store) Curier() storage.CurierRepoI {

	if s.curier == nil {
		s.curier = NewCurierRepo(s.db)
	}

	return s.curier
}

func (s *Store) Product() storage.ProductRepoI {

	if s.product == nil {
		s.product = NewProductRepo(s.db)
	}

	return s.product
}

func (s *Store) HistoryUser() storage.HistoryUserRepoI {

	if s.history_user == nil {
		s.history_user = NewHistoryUserRepo(s.db)
	}

	return s.history_user
}

func (s *Store) AdditionProduct() storage.AdditionProductRepoI {

	if s.addition_product == nil {
		s.addition_product = NewAdditionProductRepo(s.db)
	}

	return s.addition_product
}

func (s *Store) Order() storage.OrderRepoI {

	if s.order == nil {
		s.order = NewOrderRepo(s.db)
	}

	return s.order
}

func (s *Store) Cart() storage.CartRepoI {

	if s.cart == nil {
		s.cart = NewCartRepo(s.db)
	}

	return s.cart
}

func (s *Store) HistoryCurier() storage.HistoryCurierRepoI {

	if s.history_curier == nil {
		s.history_curier = NewHistoryCurierRepo(s.db)
	}

	return s.history_curier
}
