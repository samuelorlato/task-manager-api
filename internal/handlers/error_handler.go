package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/samuelorlato/task-manager-api/pkg/errors"
)

type ErrorHandler struct{}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

func (e *ErrorHandler) Handle(err *errors.HTTPError, c *gin.Context) {
	c.JSON(err.StatusCode, gin.H{
		"error":       err.Error(),
		"description": err.Description,
	})
}
