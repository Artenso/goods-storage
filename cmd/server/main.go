package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"

	goodsStorage "github.com/Artenso/goods-storage/internal/app/goods_storage"
	goodsRepo "github.com/Artenso/goods-storage/internal/repository/goods"
	goodsService "github.com/Artenso/goods-storage/internal/service/goods_storage"
	desc "github.com/Artenso/goods-storage/pkg/goods_storage/github.com/Artenso/goods_storage/pkg/goods_storage"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	host       = "localhost"
	dbPort     = "54321"
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "user_service_api"
	sslMode    = "disable"
	grpcPort   = ":50051"
	httpPort   = ":8000"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := runGRPC(); err != nil {
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := runHTTP(); err != nil {
			return
		}
	}()

	wg.Wait()
}

func runGRPC() error {
	list, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return fmt.Errorf("failed to mapping port: %s", err.Error())
	}

	repository := goodsRepo.New()
	srv := goodsService.New(repository)
	goodsStorageAPI := goodsStorage.NewGoodsStorage(srv)

	s := grpc.NewServer()
	desc.RegisterGoodsStorageServer(s, goodsStorageAPI)

	reflection.Register(s)

	if err = s.Serve(list); err != nil {
		return fmt.Errorf("failed to server: %s", err.Error())
	}

	return nil
}

func runHTTP() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()} // nolint: staticcheck

	err := desc.RegisterGoodsStorageHandlerFromEndpoint(ctx, mux, grpcPort, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(httpPort, mux)
}
