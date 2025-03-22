package controller

import (
	"database/sql"
	"gin-app/db"
	"gin-app/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser handles creating a new user.
func CreateUser(c *gin.Context) {
	var user models.User

	// Bind JSON request body to user model
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the user in the database using raw SQL
	sqlStatement := `INSERT INTO users (name, email, phone, password, age, gender)
                     VALUES (?, ?, ?, ?, ?, ?)`
	result, err := db.GetDB().Exec(sqlStatement, user.Name, user.Email, user.Phone, user.Password, user.Age, user.Gender)
	if err != nil {
		log.Println("Error creating user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return success response
	id, _ := result.LastInsertId()
	user.ID = int(id)
	c.JSON(http.StatusCreated, user)
}

func GetAllUsers(c *gin.Context) {
	// Query to fetch all users from the database
	sqlStatement := `SELECT id, name, email, phone, password, age, gender FROM users`
	rows, err := db.GetDB().Query(sqlStatement)
	if err != nil {
		log.Println("Error fetching users:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.Age, &user.Gender); err != nil {
			log.Println("Error scanning row:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read user data"})
			return
		}
		users = append(users, user)
	}

	// Check if there was an error iterating over the rows
	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	// Return success response with the list of users
	c.JSON(http.StatusOK, users)
}

// GetUser retrieves a single user by ID.
func GetUser(c *gin.Context) {
	id := c.Param("id")

	// Query to fetch the user from the database
	var user models.User
	sqlStatement := `SELECT id, name, email, phone, password, age, gender FROM users WHERE id = ?`
	err := db.GetDB().QueryRow(sqlStatement, id).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.Age, &user.Gender)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			log.Println("Error fetching user:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser updates an existing user by ID.
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	// Bind JSON request body to user model
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the user in the database
	sqlStatement := `UPDATE users SET name = ?, email = ?, phone = ?, password = ?, age = ?, gender = ? WHERE id = ?`
	_, err := db.GetDB().Exec(sqlStatement, user.Name, user.Email, user.Phone, user.Password, user.Age, user.Gender, id)
	if err != nil {
		log.Println("Error updating user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser deletes a user by ID.
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// Delete the user from the database
	sqlStatement := `DELETE FROM users WHERE id = ?`
	_, err := db.GetDB().Exec(sqlStatement, id)
	if err != nil {
		log.Println("Error deleting user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
