package database

import (
	"context"
	"fmt"
	"github.com/RaimonxDev/inventory-with-go/settings"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {
	// Syntax connection: user/pass/host/nameDatabase
	// parse time nos ayuda a trabajar con fechas mas facil
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		s.DB.User,
		s.DB.Password,
		s.DB.Host,
		s.DB.Port,
		s.DB.Name,
	)
	return sqlx.ConnectContext(ctx, "mysql", connectionString)
}
