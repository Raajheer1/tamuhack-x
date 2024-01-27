package models

import "gorm.io/gorm"

type Account struct {
	ID        string   `json:"id" gorm:"primaryKey"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	PNRs      []string `json:"pnrs" gorm:"foreignKey:AccountID"`
	Money     uint     `json:"money"`
}

func (a *Account) Create(db *gorm.DB) error {
	return db.Create(a).Error
}

func (a *Account) Get(db *gorm.DB) error {
	account := &Account{}
	err := db.First(account, a.ID).Error
	if err != nil {
		return err
	}

	return nil
}

func GetAllAccounts(db *gorm.DB) ([]*Account, error) {
	var accounts []*Account
	err := db.Find(&accounts).Error
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (a *Account) UpdateAccount(db *gorm.DB) error {
	return db.Save(a).Error
}

func DeleteAccount(db *gorm.DB, id string) error {
	return db.Delete(&Account{}, id).Error
}
