package router

import (
	"bytes"
	"encoding/json"
	"gogolook-interview/controllers"
	"gogolook-interview/models"
	"gogolook-interview/repositorys"
	"gogolook-interview/requests"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	repositorys.DatabaseSetUp()
}

func TestGetTasks(t *testing.T) {
	db := repositorys.GetDb()

	getTask := models.Task{
		ID:     1,
		Name:   "default get task",
		Status: false,
	}
	db.Create(&getTask)
	writer := makeRequest("GET", "/tasks", nil)

	expectedStatus := http.StatusOK
	expectedTask := controllers.Task{
		ID:     getTask.ID,
		Name:   getTask.Name,
		Status: 0,
	}
	assert.Equal(t, expectedStatus, writer.Code)
	assert.Contains(t, expectedTask, writer.Body)
}

func TestCreateTask(t *testing.T) {
	request := requests.CreateTaskRequest{
		Name: "test create task",
	}

	writer := makeRequest("GET", "/tasks", request)

	expectedStatus := http.StatusOK
	expectedTask := controllers.Task{
		ID:     1,
		Name:   request.Name,
		Status: 0,
	}
	assert.Equal(t, expectedStatus, writer.Code)
	assert.Contains(t, expectedTask, writer.Body)
}

func TestUpdateTask(t *testing.T) {
	db := repositorys.GetDb()

	updateTask := models.Task{
		ID:     1,
		Name:   "default task",
		Status: false,
	}
	db.Create(&updateTask)

	request := requests.UpdateTaskRequest{
		Name: "test update task",
	}

	writer := makeRequest("GET", "/tasks/"+strconv.Itoa(updateTask.ID), request)

	expectedStatus := http.StatusOK
	expectedTask := controllers.Task{
		ID:     1,
		Name:   request.Name,
		Status: 0,
	}
	assert.Equal(t, expectedStatus, writer.Code)
	assert.Contains(t, expectedTask, writer.Body)
}

func TestDeleteTask(t *testing.T) {
	db := repositorys.GetDb()

	deleteTask := models.Task{
		ID:     1,
		Name:   "default task",
		Status: false,
	}
	db.Create(&deleteTask)

	request := requests.DeleteTaskRequest{
		ID: deleteTask.ID,
	}

	writer := makeRequest("GET", "/tasks/"+strconv.Itoa(deleteTask.ID), request)

	expectedStatus := http.StatusOK
	assert.Equal(t, expectedStatus, writer.Code)
	assert.Nil(t, writer.Body)
	assert.Empty(t, db.Find(&deleteTask))
}

func makeRequest(method, url string, body interface{}) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

	writer := httptest.NewRecorder()

	Api().ServeHTTP(writer, request)
	return writer
}
