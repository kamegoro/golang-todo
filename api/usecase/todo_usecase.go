package usecase

import (
	"example.com/mod/domain"
)

type todoUsecase struct {
	todoRepo domain.TodoRepository
}

func NewTodoUsecase(tr domain.TodoRepository) domain.TodoUsecase {
	return &todoUsecase{
		todoRepo: tr,
	}
}

func (t *todoUsecase) AllGet() ([]domain.Todo, error) {
	todos, err := t.todoRepo.AllGet()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *todoUsecase) StatusUpdate(id int) error {
	err := t.todoRepo.StatusUpdate(id)
	if err != nil {
		return err
	}
	return nil
}

func (t *todoUsecase) Store(todo domain.Todo) error {
	err := t.todoRepo.Store(todo)
	if err != nil {
		return err
	}
	return nil
}

func (t *todoUsecase) Delete(id int) error {
	err := t.todoRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
