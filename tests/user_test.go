package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/Candoo/thornton-pickard-api/internal/handlers"
	"github.com/Candoo/thornton-pickard-api/internal/models"
)

// NOTE: setupTestDB must be accessible from this package (as it is in your camera_test.go)

func TestGetUsers(t *testing.T) {
	// Re-use the setup function from camera_test.go
	db := setupTestDB() 
	
	// Create two test users, now including FirstName and LastName
	user1 := models.User{
		Email: "alice@test.com",
		Password: "password123",
		FirstName: "Alice", // Added
		LastName: "Smith",   // Added
	}
	user2 := models.User{
		Email: "bob@test.com",
		Password: "password123",
		FirstName: "Bob", // Added
		LastName: "Jones",  // Added
	}
	db.Create(&user1)
	db.Create(&user2)


	// Setup router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	
	// Instantiate the handler
	userHandler := handlers.NewUserHandler(db)
	
	// NOTE: Since the route in main.go is protected, we must skip or mock Auth/Admin middleware 
	// for the test to reach the handler, or just test the handler directly.
	// For simplicity and matching the structure of other tests, we register without middleware here.
	router.GET("/users", userHandler.GetUsers) 

	// Make request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	// Assert Status
	assert.Equal(t, http.StatusOK, w.Code)
	
	// Assert Content
	var response []models.UserResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	
	assert.NoError(t, err)
	assert.Len(t, response, 2)
	
	// Assert newly added fields
	assert.Equal(t, "alice@test.com", response[0].Email)
	assert.Equal(t, "Alice", response[0].FirstName) 
	assert.Equal(t, "Smith", response[0].LastName) 

	assert.Equal(t, "bob@test.com", response[1].Email)
	assert.Equal(t, "Bob", response[1].FirstName)
	assert.Equal(t, "Jones", response[1].LastName)
	
	// Crucial check: Assert that the sensitive field 'Password' is NOT returned
	assert.NotContains(t, w.Body.String(), "password", "Response body should not contain the word 'password'")
}