package internal

import (
	"github.com/tonytangdev/go-rest-api/internal/model"
	"gorm.io/gorm"
)

func InitializeDB(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}
	return nil
}
