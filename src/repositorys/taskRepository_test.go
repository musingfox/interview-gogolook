package repositorys

import (
	"gogolook-interview/models"
	"gogolook-interview/requests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	DatabaseSetUp()
}

func TestCreateTask(t *testing.T) {
	request := requests.CreateTaskRequest{
		Name: "create task",
	}

	task, _ := CreateTask(request)

	assert.Equal(t, request.Name, task.Name)
}

func TestGetTasks(t *testing.T) {
	db := GetDb()

	getTask := models.Task{
		ID:     1,
		Name:   "default get task",
		Status: false,
	}
	db.Create(&getTask)

	tasks, _ := GetTasks()
	var expectedResult []models.Task
	expectedResult = append(expectedResult, getTask)

	assert.Equal(t, expectedResult, tasks)
}

func TestUpdateTasks(t *testing.T) {
	db := GetDb()

	originTask := models.Task{
		ID:     1,
		Name:   "default task",
		Status: false,
	}
	db.Create(&originTask)

	request := requests.UpdateTaskRequest{
		Name: "update task",
	}

	tasks, _ := UpdateTask(request)

	expectTask := models.Task{
		ID:     1,
		Name:   request.Name,
		Status: false,
	}

	assert.Equal(t, expectTask, tasks)
}

func TestDeleteTasks(t *testing.T) {
	db := GetDb()

	originTask := models.Task{
		ID:     1,
		Name:   "default task",
		Status: false,
	}
	db.Create(&originTask)

	request := requests.DeleteTaskRequest{
		ID: 1,
	}

	err := DeleteTask(request)

	assert.Nil(t, err)
	assert.Empty(t, db.First(&originTask))
}
