package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ConditionPostgres struct {
	db *sqlx.DB
}

func NewConditionPostgres(db *sqlx.DB) *ConditionPostgres {
	return &ConditionPostgres{db: db}
}

func (r *ConditionPostgres) CheckCond(id int64) string {
	query := fmt.Sprintf("SELECT condition from %s WHERE chat_id = $1 ", userTable)

	var cond string
	r.db.Get(&cond, query, id)
	return cond

}

func (r *ConditionPostgres) AddUser(id int64) error {
	query := fmt.Sprintf("INSERT INTO %s (chat_id) VALUES ($1)", userTable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *ConditionPostgres) UpdateCond(chatId int64, cond string) error {
	query := fmt.Sprintf("UPDATE %s SET condition=$1 where chat_id=$2", userTable)

	_, err := r.db.Exec(query, cond, chatId)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
