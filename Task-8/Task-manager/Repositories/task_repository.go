package repositories

import (
	"context"
	domain "tskmgr/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskDataManipulator struct {
	Data *mongo.Collection
}

func NewTaskDataManipulator (coll *mongo.Collection) *TaskDataManipulator{
	return &TaskDataManipulator{ Data: coll}
}

func (repo *TaskDataManipulator) Create (task *domain.Task) (*domain.Task,error){
	task.Id = primitive.NewObjectID()
	_,err:= repo.Data.InsertOne(context.TODO(),task)
	
	if err != nil{
		return nil,err
	}

	return task,nil
}

func (repo *TaskDataManipulator) GetByTitle(title string) (*domain.Task,error) {
	filter := bson.M{"title":title}
	task := &domain.Task{}
	err := repo.Data.FindOne(context.TODO(),filter).Decode(task)

	if err != nil{
		return nil,err
	}

	return task,err
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


