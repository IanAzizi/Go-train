package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler struct
type UserHandler struct{}

// NewUserHandler سازنده
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// GetUser godoc
// @Summary Get all users
// @Description Get a list of users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users [get]
func (u *UserHandler) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data":   []string{"user1", "user2"},
		"status": "success",
	})
}
