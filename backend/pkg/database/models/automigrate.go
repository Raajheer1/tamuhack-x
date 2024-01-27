package models

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Account{}, &Flight{}, &Seat{})
	if err != nil {
		return err
	}

	return nil
}
