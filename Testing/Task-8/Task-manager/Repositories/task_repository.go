package repositories

import (
	"context"
	domain "tskmgr/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskDataManipulator struct {
	Data domain.Collection
}

func NewTaskDataManipulator(coll domain.Collection) *TaskDataManipulator {
	return &TaskDataManipulator{Data: coll}
}

func (repo *TaskDataManipulator) Create(task *domain.Task) (*domain.Task, error) {
	task.Id = primitive.NewObjectID()
	_, err := repo.Data.InsertOne(context.TODO(), task)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (repo *TaskDataManipulator) GetByTitle(title string) (*domain.Task, error) {
	filter := bson.M{"title": title}
	task := &domain.Task{}
	err := repo.Data.FindOne(context.TODO(), filter).Decode(task)

	if err != nil {
		return nil, err
	}

	return task, err
}

func (repo *TaskDataManipulator) GetAllTasks() ([]domain.Task, error) {
	tasks := []domain.Task{}

	// Query the database for all tasks.
	cursor, err := repo.Data.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	// Iterate over the cursor and decode each task into a Task struct.
	err = cursor.All(context.Background(), &tasks)
	return tasks, err
}

// Get all tasks for a specific user.
func (repo *TaskDataManipulator) GetUserTasks(userid primitive.ObjectID) ([]domain.Task, error) {
	tasks := []domain.Task{}

	// Query the database for all tasks that have the specified user ID.
	cursor, err := repo.Data.Find(context.Background(), bson.M{"userid": userid})
	if err != nil {
		return nil, err
	}

	// Iterate over the cursor and decode each task into a Task struct.
	err = cursor.All(context.Background(), &tasks)
	return tasks, err
}

// Update a task with the specified title.
func (repo *TaskDataManipulator) UpdateTask(title string, task *domain.Task) (*domain.Task, error) {
	filter := bson.M{"title": title}

	update := bson.M{"$set": bson.M{
		"Id":          task.Id,
		"userid":      task.UserId,
		"title":       task.Title,
		"description": task.Description,
		"status":      task.Status,
		"priority":    task.Priority,
	}}

	// Update the task in the database.
	_, err := repo.Data.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// delete task by title
func (repo *TaskDataManipulator) DeleteTask(title string) error {
	filter := bson.M{"title": title}
	_, err := repo.Data.DeleteOne(context.Background(), filter)
	return err
}
