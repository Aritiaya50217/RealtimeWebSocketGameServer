package database

import (
	"errors"

	"gorm.io/gorm"
)

type Transaction struct {
	db *gorm.DB
}

func NewTransaction(db *gorm.DB) *Transaction {
	if db == nil {
		panic("transaction: db is nil")
	}
	return &Transaction{db: db}
}

func (t *Transaction) WithTransaction(fn func(tx *gorm.DB) error) error {
	if t.db == nil {
		return errors.New("database not initialized")
	}

	tx := t.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error

}
