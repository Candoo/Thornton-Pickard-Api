package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/Candoo/thornton-pickard-api/internal/models"
)

// Define a struct to hold the database dependency
type UserHandler struct {
	DB *gorm.DB
}

// NewUserHandler creates a new handler instance
func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

// GetUsers retrieves all users
// NOTE: This endpoint should typically be protected and requires administrator privileges.
// @Summary List all users
// @Description Get a list of all users
// @Tags users
// @Produce json
// @Success 200 {array} models.UserResponse
// @Router /users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	var users []models.User

	// 1. Fetch all users from the database. GORM automatically excludes soft-deleted records.
	if err := h.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	// 2. Convert to safe response struct (excluding the password field)
	userResponses := make([]models.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = user.ToUserResponse()
	}

	c.JSON(http.StatusOK, userResponses)
}