package main

import (
	"context"
	"fmt"

	serviceProvider "github.com/Artenso/goods-storage/internal/app/service_provider"
)

func main() {
	ctx := context.Background()
	app, err := serviceProvider.NewApp(ctx)
	if err != nil {
		fmt.Printf("failed to create app: %s", err.Error())
		return
	}

	err = app.Run(ctx)
	if err != nil {
		fmt.Printf("failed to run app: %s", err.Error())
	}
}
