package dbactions

import (
	"github.com/jmoiron/sqlx"
)

const queryInsertNewUser = "INSERT INTO users (name) VALUES ($1)"

type UserCreateRepository struct {
	db *sqlx.DB
}

func NewCreateUserRepository(db *sqlx.DB) *UserCreateRepository {
	return &UserCreateRepository{db: db}
}

func (cur *UserCreateRepository) UserCreate(name string) error {
	_, err := cur.db.Exec(queryInsertNewUser, name)
	if err != nil {
		return err
	}

	return nil
}
