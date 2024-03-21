package service_provider

import (
	"context"
	"fmt"
	"net"
	"net/http"

	desc "github.com/Artenso/goods-storage/pkg/goods_storage/github.com/Artenso/goods_storage/pkg/goods_storage"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = ":50051"
	httpPort = ":8000"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	group, groupCtx := errgroup.WithContext(ctx)

	group.Go(func() error {
		list, err := net.Listen("tcp", grpcPort)
		if err != nil {
			return fmt.Errorf("failed to mapping port: %s", err.Error())
		}

		if err := a.grpcServer.Serve(list); err != nil {
			return fmt.Errorf("failed to server: %s", err.Error())
		}

		return nil
	})

	group.Go(func() error {
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithInsecure()} // nolint: staticcheck

		err := desc.RegisterGoodsStorageHandlerFromEndpoint(groupCtx, mux, grpcPort, opts)
		if err != nil {
			return err
		}

		return http.ListenAndServe(httpPort, mux)
	})

	if err := group.Wait(); err != nil {
		return err
	}

	return nil
}

// initDeps initialize dependencies
func (a *App) initDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initServiceProvider,
		a.initGrpcServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGrpcServer(ctx context.Context) error {
	s := grpc.NewServer()
	desc.RegisterGoodsStorageServer(s, a.serviceProvider.getGoodsStorage(ctx))

	reflection.Register(s)

	a.grpcServer = s
	return nil
}
