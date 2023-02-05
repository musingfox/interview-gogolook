package repositorys

import (
	"fmt"
	"gogolook-interview/models"
	"gogolook-interview/requests"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DatabaseSetUp() {
	var err error
	db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Task{})
}

func CreateTask(request requests.CreateTaskRequest) (newTaskModel models.Task, err error) {
	newTaskModel = models.Task{
		Name:   request.Name,
		Status: false,
	}
	result := db.Create(&newTaskModel)

	if result.Error != nil {
		err = result.Error
		return
	}

	return
}

func GetTasks() (taskModels []models.Task, err error) {
	result := db.Find(&taskModels)

	if result.Error != nil {
		fmt.Println(result.Error)
		panic("get task failed")
	}

	return
}

func UpdateTask(request requests.UpdateTaskRequest) (taskModel models.Task, err error) {
	taskModel = models.Task{
		ID: request.ID,
	}
	result := db.First(&taskModel)

	if result.Error != nil {
		err = result.Error

		return
	}

	if request.Name != "" {
		taskModel.Name = request.Name
	}
	if request.Status == 0 || request.Status == 1 {
		taskModel.Status = request.Status == 1
	}

	db.Save(&taskModel)

	return
}

func DeleteTask(request requests.DeleteTaskRequest) (err error) {
	taskModel := models.Task{
		ID: request.ID,
	}

	result := db.Delete(&taskModel)

	if result.Error != nil {
		err = result.Error

		return
	}

	return
}
