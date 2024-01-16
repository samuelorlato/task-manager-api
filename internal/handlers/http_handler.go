package handlers

import (
	"os"
	"time"

	baseErrors "errors"

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
	authService       ports.AuthService
	errorHandler      *ErrorHandler
}

func NewHTTPHandler(engine *gin.Engine, taskUsecase ports.TaskUsecase, userUsecase ports.UserUsecase, encryptionService ports.EncryptionService, authService ports.AuthService, errorHandler *ErrorHandler) *HTTPHandler {
	return &HTTPHandler{
		engine:            engine,
		taskUsecase:       taskUsecase,
		userUsecase:       userUsecase,
		encryptionService: encryptionService,
		authService:       authService,
		errorHandler:      errorHandler,
	}
}

func (h *HTTPHandler) SetRoutes() {
	h.engine.GET("/tasks", h.authenticateMiddleware, h.getTasks)
	h.engine.POST("/tasks", h.authenticateMiddleware, h.createTask)
	h.engine.GET("/tasks/:id", h.authenticateMiddleware, h.getTaskById)
	h.engine.PATCH("/tasks/:id", h.authenticateMiddleware, h.updateTask)
	h.engine.DELETE("/tasks/:id", h.authenticateMiddleware, h.deleteTask)
	h.engine.DELETE("/user", h.authenticateMiddleware, h.deleteUser)
	h.engine.PATCH("/user", h.authenticateMiddleware, h.updateUser)
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

	userModel, err := h.userUsecase.GetUser(userDTO.Email, userDTO.Password)
	if err != nil {
		h.errorHandler.Handle(err, c)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	secret := os.Getenv("JWT_SECRET")
	token, authErr := h.authService.GenerateToken(userModel.Email, &expirationTime, secret)
	if authErr != nil {
		err := errors.NewJWTGenerationError(authErr)
		h.errorHandler.Handle(err, c)
		return
	}

	c.JSON(200, gin.H{"token": token})
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

	c.JSON(200, gin.H{"ok": "use /login to authenticate"})
}

func (h *HTTPHandler) authenticateMiddleware(c *gin.Context) {
	var email string
	var tokenErr error

	if c.Request.Header["Authorization"] == nil {
		baseError := baseErrors.New("Authorization header not found")
		err := errors.NewAuthorizationError(baseError)
		h.errorHandler.Handle(err, c)
		c.Abort()
		return
	} else {
		secret := os.Getenv("JWT_SECRET")
		email, tokenErr = h.authService.ValidateToken(c.Request.Header["Authorization"][0], secret)
		if tokenErr != nil {
			err := errors.NewAuthorizationError(tokenErr)
			h.errorHandler.Handle(err, c)
			c.Abort()
			return
		}
	}

	c.Set("loggedAs", email)
	c.Next()
}

func (h *HTTPHandler) getLoggedEmail(c *gin.Context) (string, error) {
	email, exists := c.Get("loggedAs")
	if !exists {
		baseError := baseErrors.New("Email not found in context, apparently not logged")
		return "", baseError
	}

	return email.(string), nil
}

func (h *HTTPHandler) getTasks(c *gin.Context) {
	email, authenticationError := h.getLoggedEmail(c)
	if authenticationError != nil {
		err := errors.NewAuthorizationError(authenticationError)
		h.errorHandler.Handle(err, c)
		return
	}

	tasks, err := h.taskUsecase.GetTasks(email)
	if err != nil {
		h.errorHandler.Handle(err, c)
		return
	}

	var tasksDTO []dtos.TaskDTO
	for _, task := range tasks {
		taskDTO := dtos.TaskDTO{
			Id:          task.Id.String(),
			Title:       task.Title,
			Description: task.Description,
			ToDate:      task.ToDate.String(),
			Completed:   *task.Completed,
			Tags:        task.Tags,
		}

		tasksDTO = append(tasksDTO, taskDTO)
	}

	c.JSON(200, tasksDTO)
}

func (h *HTTPHandler) createTask(c *gin.Context) {
	email, authenticationError := h.getLoggedEmail(c)
	if authenticationError != nil {
		err := errors.NewAuthorizationError(authenticationError)
		h.errorHandler.Handle(err, c)
		return
	}

	var createTaskDTO dtos.CreateTaskDTO

	bindErr := c.BindJSON(&createTaskDTO)
	if bindErr != nil {
		err := errors.NewValidationError(bindErr)
		h.errorHandler.Handle(err, c)
		return
	}

	id, err := h.taskUsecase.CreateTask(email, createTaskDTO.Title, &createTaskDTO.Description, createTaskDTO.ToDate, &createTaskDTO.Tags)
	if err != nil {
		h.errorHandler.Handle(err, c)
		return
	}

	c.JSON(200, gin.H{"createdTaskId": id})
}

func (h *HTTPHandler) getTaskById(c *gin.Context) {
	email, authenticationError := h.getLoggedEmail(c)
	if authenticationError != nil {
		err := errors.NewAuthorizationError(authenticationError)
		h.errorHandler.Handle(err, c)
		return
	}

	id := c.Param("id")

	task, err := h.taskUsecase.GetTaskById(email, id)
	if err != nil {
		h.errorHandler.Handle(err, c)
		return
	}

	taskDTO := dtos.TaskDTO{
		Id:          task.Id.String(),
		Title:       task.Title,
		Description: task.Description,
		ToDate:      task.ToDate.String(),
		Completed:   *task.Completed,
		Tags:        task.Tags,
	}

	c.JSON(200, taskDTO)
}

func (h *HTTPHandler) updateTask(c *gin.Context) {
	email, authenticationError := h.getLoggedEmail(c)
	if authenticationError != nil {
		err := errors.NewAuthorizationError(authenticationError)
		h.errorHandler.Handle(err, c)
		return
	}

	var updateTaskDTO dtos.UpdateTaskDTO

	bindErr := c.BindJSON(&updateTaskDTO)
	if bindErr != nil {
		err := errors.NewValidationError(bindErr)
		h.errorHandler.Handle(err, c)
		return
	}

	id := c.Param("id")

	err := h.taskUsecase.UpdateTask(email, id, updateTaskDTO.Title, updateTaskDTO.Description, updateTaskDTO.ToDate, updateTaskDTO.Completed, updateTaskDTO.Tags)
	if err != nil {
		h.errorHandler.Handle(err, c)
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}

func (h *HTTPHandler) deleteTask(c *gin.Context) {
	email, authenticationError := h.getLoggedEmail(c)
	if authenticationError != nil {
		err := errors.NewAuthorizationError(authenticationError)
		h.errorHandler.Handle(err, c)
		return
	}

	id := c.Param("id")

	err := h.taskUsecase.DeleteTask(email, id)
	if err != nil {
		h.errorHandler.Handle(err, c)
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}

func (h *HTTPHandler) updateUser(c *gin.Context) {
	email, authenticationError := h.getLoggedEmail(c)
	if authenticationError != nil {
		err := errors.NewAuthorizationError(authenticationError)
		h.errorHandler.Handle(err, c)
		return
	}

	var updateUserDTO dtos.UpdateUserDTO

	bindErr := c.BindJSON(&updateUserDTO)
	if bindErr != nil {
		err := errors.NewValidationError(bindErr)
		h.errorHandler.Handle(err, c)
		return
	}

	err := h.userUsecase.UpdateUser(email, updateUserDTO.Email, updateUserDTO.Password)
	if err != nil {
		h.errorHandler.Handle(err, c)
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}

func (h *HTTPHandler) deleteUser(c *gin.Context) {
	email, authenticationError := h.getLoggedEmail(c)
	if authenticationError != nil {
		err := errors.NewAuthorizationError(authenticationError)
		h.errorHandler.Handle(err, c)
		return
	}

	err := h.userUsecase.DeleteUser(email)
	if err != nil {
		h.errorHandler.Handle(err, c)
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}
