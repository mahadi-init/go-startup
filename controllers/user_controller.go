package controllers

import (
	"gin-app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// UserController handles user-related requests
type UserController struct {
	// In a real application, you would inject a service or repository here
	users  []models.User
	nextID int
}

// NewUserController creates a new user controller with sample data
func NewUserController() *UserController {
	// Initialize with some sample data
	return &UserController{
		users: []models.User{
			{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30},
			{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Age: 25},
		},
		nextID: 3,
	}
}

// GetUsers returns all users
func (uc *UserController) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, uc.users)
}

// GetUser returns a specific user by ID
func (uc *UserController) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	for _, user := range uc.users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// CreateUser creates a new user
func (uc *UserController) CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate user data
	if newUser.Name == "" || newUser.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name and email are required"})
		return
	}

	// Set the ID and add to users
	newUser.ID = uc.nextID
	uc.nextID++
	uc.users = append(uc.users, newUser)

	c.JSON(http.StatusCreated, newUser)
}

// UpdateUser updates an existing user
func (uc *UserController) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, user := range uc.users {
		if user.ID == id {
			// Preserve the ID
			updatedUser.ID = id
			uc.users[i] = updatedUser
			c.JSON(http.StatusOK, updatedUser)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// DeleteUser deletes a user by ID
func (uc *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	for i, user := range uc.users {
		if user.ID == id {
			// Remove the user from the slice
			uc.users = append(uc.users[:i], uc.users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
