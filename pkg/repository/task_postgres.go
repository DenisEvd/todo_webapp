package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todo_webapp"
)

type TaskPostgres struct {
	db *sqlx.DB
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (r *TaskPostgres) Create(userId int, task todo_webapp.Task) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	var id int
	createTaskQuery := fmt.Sprintf("INSERT INTO %s (tag, title, description, deadline) VALUES ($1, $2, $3, $4) RETURNING id", taskTable)
	row := tx.QueryRow(createTaskQuery, task.Tag, task.Title, task.Description, task.Deadline)
	if err := row.Scan(&id); err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	createTasksUsersQuery := fmt.Sprintf("INSERT INTO %s (user_id, task_id) VALUES ($1, $2)", taskUserTable)
	_, err = tx.Exec(createTasksUsersQuery, userId, id)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
