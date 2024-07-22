package dbactions

import (
	"github.com/jmoiron/sqlx"
)

const queryInsertNewUser = "INSERT INTO users (name) VALUES ($1)"

type CreateUserRepository struct {
	db *sqlx.DB
}

func NewCreateUserRepository(db *sqlx.DB) *CreateUserRepository {
	return &CreateUserRepository{db: db}
}

func (cur *CreateUserRepository) CreateNewUser(name string) error {
	_, err := cur.db.Exec(queryInsertNewUser, name)
	if err != nil {
		return err
	}

	return nil

}
