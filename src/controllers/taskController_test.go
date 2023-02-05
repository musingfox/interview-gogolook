package controllers

import (
	"gogolook-interview/models"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetTasks(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetTasks(tt.args.c)
		})
	}
}

func TestCreateTask(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateTask(tt.args.c)
		})
	}
}

func TestUpdateTask(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateTask(tt.args.c)
		})
	}
}

func TestDeleteTask(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteTask(tt.args.c)
		})
	}
}

func Test_response(t *testing.T) {
	type args struct {
		c        *gin.Context
		status   int
		response interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response(tt.args.c, tt.args.status, tt.args.response)
		})
	}
}

func Test_transform(t *testing.T) {
	type args struct {
		model models.Task
	}
	tests := []struct {
		name string
		args args
		want Task
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transform(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transform() = %v, want %v", got, tt.want)
			}
		})
	}
}
