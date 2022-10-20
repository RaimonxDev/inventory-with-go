package main

import (
	"context"
	"github.com/RaimonxDev/inventory-with-go/internal/database"
	"github.com/RaimonxDev/inventory-with-go/internal/repository"
	"github.com/RaimonxDev/inventory-with-go/internal/service"
	"github.com/RaimonxDev/inventory-with-go/settings"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New),
		fx.Invoke(
			func(ctx context.Context, serv service.Service) {
				err := serv.RegisterUser(
					ctx, "ramon@ramon.com", "ramon", "Andreina_877")
				if err != nil {
					panic(err)
				}
			},
		),
	)
	app.Run()
}
