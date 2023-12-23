package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/samuelorlato/task-manager-api/internal/core/ports"
	"github.com/samuelorlato/task-manager-api/internal/handlers/dtos"
	"github.com/samuelorlato/task-manager-api/pkg/errors"
)

type HTTPHandler struct {
	engine            *gin.Engine
	taskUsecase       ports.TaskUsecase
	userUsecase       ports.UserUsecase
	encryptionService ports.EncryptionService
	errorHandler      *ErrorHandler
}

func NewHTTPHandler(engine *gin.Engine, taskUsecase ports.TaskUsecase, userUsecase ports.UserUsecase, encryptionService ports.EncryptionService, errorHandler *ErrorHandler) *HTTPHandler {
	return &HTTPHandler{
		engine:            engine,
		taskUsecase:       taskUsecase,
		userUsecase:       userUsecase,
		encryptionService: encryptionService,
		errorHandler:      errorHandler,
	}
}

func (h *HTTPHandler) SetRoutes() {
	h.engine.GET("/tasks", h.getTasks)
	h.engine.POST("/tasks", h.createTask)
	h.engine.GET("/tasks/:id", h.getTaskById)
	h.engine.PATCH("/tasks/:id", h.updateTask)
	h.engine.DELETE("/tasks/:id", h.deleteTask)
	h.engine.POST("/login", h.login)
	h.engine.POST("/register", h.register)
}

func (h *HTTPHandler) login(c *gin.Context) {
	var userDTO dtos.UserDTO

	bindErr := c.BindJSON(&userDTO)
	if bindErr != nil {
		err := errors.NewValidationError(bindErr)
		h.errorHandler.Handle(err, c)
		return
	}

	_, err := h.userUsecase.GetUser(userDTO.Email, userDTO.Password)
	if err != nil {
		h.errorHandler.Handle(err, c)
		return
	}
}

func (h *HTTPHandler) register(c *gin.Context) {
	var userDTO dtos.UserDTO

	bindErr := c.BindJSON(&userDTO)
	if bindErr != nil {
		err := errors.NewValidationError(bindErr)
		h.errorHandler.Handle(err, c)
		return
	}

	err := h.userUsecase.CreateUser(userDTO.Email, userDTO.Password)
	if err != nil {
		h.errorHandler.Handle(err, c)
		return
	}
}

func (h *HTTPHandler) getTasks(c *gin.Context) {
	tasks, err := h.taskUsecase.GetTasks()
	if err != nil {
		h.errorHandler.Handle(err, c)
		return
	}

	c.JSON(200, tasks)
}

func (h *HTTPHandler) createTask(c *gin.Context) {
	var createTaskDTO dtos.CreateTaskDTO

	bindErr := c.BindJSON(&createTaskDTO)
	if bindErr != nil {
		err := errors.NewValidationError(bindErr)
		h.errorHandler.Handle(err, c)
		return
	}

	err := h.taskUsecase.CreateTask(createTaskDTO.Title, &createTaskDTO.Description, createTaskDTO.ToDate, &createTaskDTO.Tags)
	if err != nil {
		h.errorHandler.Handle(err, c)
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}

func (h *HTTPHandler) getTaskById(c *gin.Context) {
	id := c.Param("id")

	task, err := h.taskUsecase.GetTaskById(id)
	if err != nil {
		h.errorHandler.Handle(err, c)
		return
	}

	c.JSON(200, task)
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

	err := h.taskUsecase.UpdateTask(id, &updateTaskDTO.Title, &updateTaskDTO.Description, &updateTaskDTO.ToDate, &updateTaskDTO.Completed, &updateTaskDTO.Tags)
	if err != nil {
		h.errorHandler.Handle(err, c)
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}

func (h *HTTPHandler) deleteTask(c *gin.Context) {
	id := c.Param("id")

	err := h.taskUsecase.DeleteTask(id)
	if err != nil {
		h.errorHandler.Handle(err, c)
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}
