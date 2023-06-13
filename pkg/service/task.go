package service

import (
	"todo_webapp"
	"todo_webapp/pkg/repository"
)

type TaskService struct {
	repo repository.Task
}

func NewTaskService(rep repository.Task) *TaskService {
	return &TaskService{repo: rep}
}

func (s *TaskService) Create(userId int, task todo_webapp.Task) (int, error) {
	return s.repo.Create(userId, task)
}
