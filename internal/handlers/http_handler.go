package handlers

import (
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
	// TODO: implement
}

func (h *HTTPHandler) createTask(c *gin.Context) {
	var taskDTO dtos.CreateTaskDTO

	err := c.BindJSON(&taskDTO)
	if err != nil {
		// TODO: handle
	}

	err = h.usecase.CreateTask(taskDTO.Title, &taskDTO.Description, taskDTO.ToDate)
	if err != nil {
		// TODO: handle
	}
}

func (h *HTTPHandler) getTaskById(c *gin.Context) {
	// TODO: implement
}

func (h *HTTPHandler) updateTask(c *gin.Context) {
	// TODO: implement
}

func (h *HTTPHandler) deleteTask(c *gin.Context) {
	// TODO: implement
}
