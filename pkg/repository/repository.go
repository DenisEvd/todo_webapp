package repository

import (
	"github.com/jmoiron/sqlx"
	"todo_webapp"
)

type Authorization interface {
	CreateUser(user todo_webapp.User) (int, error)
}

type Task interface {
}

type Repository struct {
	Authorization
	Task
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
