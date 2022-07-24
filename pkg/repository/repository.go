package repository

import (
	bills "accounting_teko"
	"github.com/jmoiron/sqlx"
)

type Company interface {
	AddCompany(id int) (int, error)
	ValidateCompany(id int, chatId int64) (bool, error)
}

type Billing interface {
	Create(bill bills.Bill) (int, error)
	GetHistory(chatId int64) ([]bills.Bill, error)
	GetTotal(chatId int64) (int, error)
	AddAmount(amount int, chatId int64) error
	AddDescription(desc string, chatId int64, sub string) error
}

type Condition interface {
	CheckCond(id int64) (string, error)
	AddUser(id int64) error
	UpdateCond(chatId int64, cond string) error
}

type Repository struct {
	Company
	Billing
	Condition
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Company:   NewCompanyPostgres(db),
		Billing:   NewBillingPostgres(db),
		Condition: NewConditionPostgres(db),
	}
}
