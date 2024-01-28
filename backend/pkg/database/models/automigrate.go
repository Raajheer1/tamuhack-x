package models

import (
	"errors"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&Account{},
		&Flight{},
		&Seat{},
	)
	if err != nil {
		return errors.New("[Database] Migration Error: " + err.Error())
	}

	return nil
}
