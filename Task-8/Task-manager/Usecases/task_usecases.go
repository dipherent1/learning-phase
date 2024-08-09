package usecases

import (
	"errors"
	domain "tskmgr/Domain"
	repositories "tskmgr/Repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

// get all tasks
func (u *TaskUsecase) GetAllTasks() ([]domain.Task, error) {
	tasks, err := u.MyTaskRepo.GetAllTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// get user tasks
func (u *TaskUsecase) GetUserTasks(userid primitive.ObjectID) ([]domain.Task, error) {
	tasks, err := u.MyTaskRepo.GetUserTasks(userid)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// validate and update task
func (u *TaskUsecase) UpdateTask(userrole string, userid primitive.ObjectID, title string, newtask *domain.Task) (*domain.Task, error) {

	//validate if user is admin or if task user id is the same as the user id
	task, err := u.MyTaskRepo.GetByTitle(title)
	if err!=nil{
		return nil,err
	}
	
	if userrole != "admin" && task.UserId != userid {
		return nil, errors.New("you are not authorized to update this task")

	}
	newtask.Id = task.Id
	newtask.UserId = task.UserId
	
	task, err = u.MyTaskRepo.UpdateTask(title, newtask)
	if err != nil {
		return nil, err
	}
	return task, nil
}

//validate and delete task
func (u *TaskUsecase) DeleteTask(userrole string, userid primitive.ObjectID, title string) error {
	//validate if user is admin or if task user id is the same as the user id
	task, err := u.MyTaskRepo.GetByTitle(title)
	if err!=nil{
		return err
	}
	if userrole != "admin" && task.UserId != userid {
		return errors.New("you are not authorized to delete this task")

	}

	err = u.MyTaskRepo.DeleteTask(title)
	if err != nil {
		return err
	}
	return nil
}
