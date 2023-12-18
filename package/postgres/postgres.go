package postgres

import (
	"fmt"

	"github.com/Jhnvlglmlbrt/crud-gin/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.PG) (*gorm.DB, error) {
	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	return gorm.Open(postgres.Open(connection), &gorm.Config{})
}
