package repository

import (
	"github.com/jmoiron/sqlx"
	"todo_webapp"
)

type Authorization interface {
	CreateUser(user todo_webapp.User) (int, error)
	GetUser(username, passwords string) (todo_webapp.User, error)
}

type Task interface {
	Create(userId int, task todo_webapp.Task) (int, error)
}

type Repository struct {
	Authorization
	Task
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Task:          NewTaskPostgres(db),
	}
}
