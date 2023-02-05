package requests

type CreateTaskRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTaskRequest struct {
	ID     int    `json:"id" uri:"id" binding:"required"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

type DeleteTaskRequest struct {
	ID int `uri:"id" binding:"required"`
}
