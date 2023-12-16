package handlers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/samuelorlato/task-manager-api/internal/core/ports"
	"github.com/samuelorlato/task-manager-api/internal/handlers/dtos"
	"github.com/samuelorlato/task-manager-api/pkg/errors"
)

type HTTPHandler struct {
	engine       *gin.Engine
	usecase      ports.TaskUsecase
	errorHandler *ErrorHandler
}

func NewHTTPHandler(engine *gin.Engine, usecase ports.TaskUsecase, errorHandler *ErrorHandler) *HTTPHandler {
	return &HTTPHandler{
		engine:       engine,
		usecase:      usecase,
		errorHandler: errorHandler,
	}
}

func (h *HTTPHandler) SetRoutes() {
	h.engine.GET("/tasks", h.getTasks)
	h.engine.POST("/tasks", h.createTask)
	h.engine.GET("/tasks/:id", h.getTaskById)
	h.engine.PATCH("/tasks/:id", h.updateTask)
	h.engine.DELETE("/tasks/:id", h.deleteTask)
}

func (h *HTTPHandler) getTasks(c *gin.Context) {
	tasks, err := h.usecase.GetTasks()
	if err != nil {
		h.errorHandler.Handle(err, c)
		return
	}

	b, marshalErr := json.Marshal(tasks)
	if err != nil {
		err := errors.NewGenericError(marshalErr)
		h.errorHandler.Handle(err, c)
		return
	}

	c.JSON(200, string(b))
}

func (h *HTTPHandler) createTask(c *gin.Context) {
	var createTaskDTO dtos.CreateTaskDTO

	bindErr := c.BindJSON(&createTaskDTO)
	if bindErr != nil {
		err := errors.NewValidationError(bindErr)
		h.errorHandler.Handle(err, c)
		return
	}

	err := h.usecase.CreateTask(createTaskDTO.Title, &createTaskDTO.Description, createTaskDTO.ToDate, &createTaskDTO.Tags)
	if err != nil {
		h.errorHandler.Handle(err, c)
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}

func (h *HTTPHandler) getTaskById(c *gin.Context) {
	id := c.Param("id")

	task, err := h.usecase.GetTaskById(id)
	if err != nil {
		h.errorHandler.Handle(err, c)
		return
	}

	b, marshalErr := json.Marshal(task)
	if err != nil {
		err := errors.NewGenericError(marshalErr)
		h.errorHandler.Handle(err, c)
		return
	}

	c.JSON(200, string(b))
}

func (h *HTTPHandler) updateTask(c *gin.Context) {
	var updateTaskDTO dtos.UpdateTaskDTO

	bindErr := c.BindJSON(&updateTaskDTO)
	if bindErr != nil {
		err := errors.NewValidationError(bindErr)
		h.errorHandler.Handle(err, c)
		return
	}

	id := c.Param("id")

	err := h.usecase.UpdateTask(id, &updateTaskDTO.Title, &updateTaskDTO.Description, &updateTaskDTO.ToDate, &updateTaskDTO.Completed, &updateTaskDTO.Tags)
	if err != nil {
		h.errorHandler.Handle(err, c)
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}

func (h *HTTPHandler) deleteTask(c *gin.Context) {
	id := c.Param("id")

	err := h.usecase.DeleteTask(id)
	if err != nil {
		h.errorHandler.Handle(err, c)
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}
