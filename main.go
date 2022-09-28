package main

import (
	"context"
	"fmt"
	"github.com/RaimonxDev/inventory-with-go/internal/database"
	"github.com/RaimonxDev/inventory-with-go/internal/model"
	"github.com/RaimonxDev/inventory-with-go/settings"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New),
		fx.Invoke(
			configureLyfeCycle),
	)
	app.Run()
}

func configureLyfeCycle(lc fx.Lifecycle, db *gorm.DB) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Starting application")
			return db.AutoMigrate(
				&model.User{},
				&model.Product{},
				&model.Role{},
				&model.UserRole{})
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping application")
			return database.Close()
		},
	})
}
