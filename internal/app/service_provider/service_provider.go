package service_provider

import (
	"context"
	"log"

	impl "github.com/Artenso/goods-storage/internal/app/goods_storage"
	goodsPg "github.com/Artenso/goods-storage/internal/repository/goods_pg"
	goodsStorage "github.com/Artenso/goods-storage/internal/service/goods_storage"
	"github.com/jackc/pgx/v5"
)

// serviceProvider di-container
type serviceProvider struct {
	dbConn         *pgx.Conn
	goodsRepo      *goodsPg.Repository
	service        *goodsStorage.Service
	implementation *impl.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) getDbConn(ctx context.Context) *pgx.Conn {
	if s.dbConn == nil {
		dbDSN := "postgres://postgres:postgres@localhost:5432/goods_storage"
		conn, err := pgx.Connect(ctx, dbDSN)
		if err != nil {
			log.Fatalf("failed to init db connection: %s", err.Error())
		}

		s.dbConn = conn
	}

	return s.dbConn
}

func (s *serviceProvider) getGoodsRepo(ctx context.Context) *goodsPg.Repository {
	if s.goodsRepo == nil {
		s.goodsRepo = goodsPg.New(s.getDbConn(ctx))
	}

	return s.goodsRepo
}

func (s *serviceProvider) getService(ctx context.Context) *goodsStorage.Service {
	if s.service == nil {
		s.service = goodsStorage.New(s.getGoodsRepo(ctx))
	}

	return s.service
}

func (s *serviceProvider) getGoodsStorage(ctx context.Context) *impl.Implementation {
	if s.implementation == nil {
		s.implementation = impl.NewGoodsStorage(s.getService(ctx))
	}

	return s.implementation
}
