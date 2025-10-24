package topic2

import (
	"errors"
	"gorm.io/gorm"
)

type Account struct {
	ID      int `gorm:"primary_key"`
	Balance float32
}

type Transaction struct {
	ID            int `gorm:"primary_key"`
	Amount        float32
	FromAccountID int
	ToAccountID   int
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&Account{}, &Transaction{})

	db.Create(&Account{ID: 1, Balance: 1000.0})
	db.Create(&Account{ID: 2, Balance: 1000.0})

	db.Transaction(func(tx *gorm.DB) error {
		var accountA = Account{}
		tx.Where("id = ?", 1).Find(&accountA)
		if accountA.Balance < 100.0 {
			return errors.New("A账户余额不足！")
		}
		accountA.Balance -= 100.0
		tx.Save(&accountA)

		var accountB = Account{}
		tx.Where("id = ?", 2).Find(&accountB)
		accountB.Balance += 100.0
		tx.Save(&accountB)
		tx.Create(&Transaction{
			Amount:        100.0,
			FromAccountID: accountA.ID,
			ToAccountID:   accountB.ID,
		})
		return nil
	})
}
