package repository

import (
	bills "accounting_teko"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type BillingPostgres struct {
	db *sqlx.DB
}

func NewBillingPostgres(db *sqlx.DB) *BillingPostgres {
	return &BillingPostgres{db: db}
}

func (r *BillingPostgres) Create(bill bills.Bill) (int, error) {

	return 0, nil
}

func (r *BillingPostgres) GetHistory(chatId int64) ([]bills.Bill, error) {
	query := fmt.Sprintf("SELECT company_id from %s WHERE chat_id = $1 ", userTable)
	var id int
	err := r.db.Get(&id, query, chatId)
	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}
	query = fmt.Sprintf("SELECT id, amount, description, email, company_id from %s WHERE company_id = $1 ORDER BY time DESC LIMIT 10 ", billTable)
	var res []bills.Bill
	err = r.db.Select(&res, query, id)
	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}

	return res, nil
}

func (r *BillingPostgres) GetTotal(chatId int64) (int, error) {
	query := fmt.Sprintf("SELECT COUNT(1) from %s b inner join %s u  on b.company_id=u.company_id  WHERE chat_id = $1  and time>'today'", billTable, userTable)
	var res int
	err := r.db.Get(&res, query, chatId)
	if err != nil {
		fmt.Printf(err.Error())
		return 0, err
	}
	if res == 0 {
		return 0, nil
	}
	query = fmt.Sprintf("SELECT SUM(amount) from %s b inner join %s u  on b.company_id=u.company_id  WHERE chat_id = $1  and time>'today'", billTable, userTable)

	err = r.db.Get(&res, query, chatId)
	if err != nil {
		fmt.Printf(err.Error())
		return 0, err
	}
	return res, nil
}

func (r *BillingPostgres) AddAmount(amount int, chatId int64) error {
	query := fmt.Sprintf("SELECT company_id from %s WHERE chat_id = $1 ", userTable)
	var id int
	err := r.db.Get(&id, query, chatId)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	tx, err := r.db.Begin()
	query = fmt.Sprintf("INSERT INTO %s (amount, company_id, time) VALUES ($1, $2, current_timestamp) RETURNING id",
		billTable)
	row := tx.QueryRow(query, amount, id)
	err = row.Scan(&id)
	if err != nil {
		tx.Rollback()
		return err
	}
	query = fmt.Sprintf("UPDATE %s SET current_bill_id=$1 where chat_id=$2", userTable)
	_, err = tx.Exec(query, id, chatId)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()

}

func (r *BillingPostgres) AddDescription(desc string, chatId int64, sub string) error {
	query := fmt.Sprintf("SELECT current_bill_id from %s WHERE chat_id = $1 ", userTable)
	var id int
	err := r.db.Get(&id, query, chatId)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	query = fmt.Sprintf("UPDATE  %s SET %s=$1 WHERE  id=$2",
		billTable, sub)
	_, err = r.db.Exec(query, desc, id)

	if err != nil {
		fmt.Printf(err.Error())
		return err
	}

	return nil
}
