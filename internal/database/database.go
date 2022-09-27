package database

import (
	"fmt"
	"github.com/RaimonxDev/inventory-with-go/settings"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

func newPostgresDB(s *settings.Settings) {
	once.Do(func() {
		var err error
		dsn := fmt.Sprintf(
			"host=%s "+
				"user=%s "+
				"password=%s "+
				"dbname=%s "+
				"port=%d "+
				"sslmode=disable",
			s.DB.Host,
			s.DB.User,
			s.DB.Password,
			s.DB.Name,
			s.DB.Port,
		)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("ERROR DATABASE")
		}

	})
}
