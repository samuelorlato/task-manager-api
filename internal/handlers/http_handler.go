package handlers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/samuelorlato/task-manager-api/internal/core/ports"
	"github.com/samuelorlato/task-manager-api/internal/handlers/dtos"
)

type HTTPHandler struct {
	engine  *gin.Engine
	usecase ports.TaskUsecase
}

func NewHTTPHandler(engine *gin.Engine, usecase ports.TaskUsecase) *HTTPHandler {
	return &HTTPHandler{
		engine:  engine,
		usecase: usecase,
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
		// TODO: handle
	}

	b, err := json.Marshal(tasks)
	if err != nil {
		// TODO: handle
	}

	c.JSON(200, string(b))
}

func (h *HTTPHandler) createTask(c *gin.Context) {
	var createTaskDTO dtos.CreateTaskDTO

	err := c.BindJSON(&createTaskDTO)
	if err != nil {
		// TODO: handle
	}

	err = h.usecase.CreateTask(createTaskDTO.Title, &createTaskDTO.Description, createTaskDTO.ToDate)
	if err != nil {
		// TODO: handle
	}

	c.JSON(200, nil)
}

func (h *HTTPHandler) getTaskById(c *gin.Context) {
	var getTaskByIdDTO dtos.GetTaskByIdDTO

	err := c.BindJSON(&getTaskByIdDTO)
	if err != nil {
		// TODO: handle
	}

	task, err := h.usecase.GetTaskById(getTaskByIdDTO.Id)
	if err != nil {
		// TODO: handle
	}

	b, err := json.Marshal(task)
	if err != nil {
		// TODO: handle
	}

	c.JSON(200, string(b))
}

func (h *HTTPHandler) updateTask(c *gin.Context) {
	var updateTaskDTO dtos.UpdateTaskDTO

	err := c.BindJSON(&updateTaskDTO)
	if err != nil {
		// TODO: handle
	}

	err = h.usecase.UpdateTask(updateTaskDTO.Id, &updateTaskDTO.Title, &updateTaskDTO.Description, &updateTaskDTO.ToDate, &updateTaskDTO.Completed)
	if err != nil {
		// TODO: handle
	}

	c.JSON(200, nil)
}

func (h *HTTPHandler) deleteTask(c *gin.Context) {
	var deleteTaskDTO dtos.DeleteTaskDTO

	err := c.BindJSON(&deleteTaskDTO)
	if err != nil {
		// TODO: handle
	}

	err = h.usecase.DeleteTask(deleteTaskDTO.Id)
	if err != nil {
		// TODO: handle
	}

	c.JSON(200, nil)
}
