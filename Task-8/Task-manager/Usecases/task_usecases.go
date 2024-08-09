package usecases

import (
	domain "tskmgr/Domain"
	repositories "tskmgr/Repositories"
)

type TaskUsecase struct {
	MyTaskRepo *repositories.TaskDataManipulator
}

func NewTaskUsecase(repo *repositories.TaskDataManipulator) *TaskUsecase {
	return &TaskUsecase{MyTaskRepo: repo}
}

func (u *TaskUsecase) CreateTask(task *domain.Task) (*domain.Task, error) {

	task, err := u.MyTaskRepo.Create(task)
	if err != nil {
		return nil, err
	}

	return task, nil

}

func (u *TaskUsecase) GetTaskByTitle(title string) (*domain.Task, error) {
	task, err := u.MyTaskRepo.GetByTitle(title)
	if err != nil {
		return nil, err
	}
	return task, nil
}
