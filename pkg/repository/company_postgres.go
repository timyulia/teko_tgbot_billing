package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CompanyPostgres struct {
	db *sqlx.DB
}

func NewCompanyPostgres(db *sqlx.DB) *CompanyPostgres {
	return &CompanyPostgres{db: db}
}

func (r *CompanyPostgres) AddCompany(id int) (int, error) {

	query := fmt.Sprintf("INSERT INTO %s (id) VALUES ($1)", companyTable)

	_, err := r.db.Exec(query, id)

	return 0, err
}

func (r *CompanyPostgres) ValidateCompany(id int, chatId int64) (bool, error) {
	query := fmt.Sprintf("SELECT count(1) from %s WHERE id = $1 ", companyTable)
	var count int
	err := r.db.Get(&count, query, id)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	if count > 0 {
		query := fmt.Sprintf("UPDATE %s SET company_id=$1, condition='validated' where chat_id=$2", userTable)

		_, err := r.db.Exec(query, id, chatId)

		if err != nil {
			fmt.Println(err.Error())
			return false, err

		}
		return true, nil
	}
	//fmt.Println("<1")
	return false, nil
}
