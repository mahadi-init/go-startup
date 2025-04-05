package controller

import (
	"database/sql"
	"gin-app/db"
	"gin-app/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `INSERT INTO users (name, email, phone, password, age, gender)
                     VALUES (?, ?, ?, ?, ?, ?)`
	result, err := db.GetDB().Exec(sqlStatement, user.Name, user.Email, user.Phone, user.Password, user.Age, user.Gender)
	if err != nil {
		log.Println("Error creating user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	user.ID = int(id)
	c.JSON(http.StatusCreated, user)
}

func GetAllUsers(c *gin.Context) {
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

func GetUser(c *gin.Context) {
	id := c.Param("id")

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

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `UPDATE users SET name = ?, email = ?, phone = ?, password = ?, age = ?, gender = ? WHERE id = ?`
	_, err := db.GetDB().Exec(sqlStatement, user.Name, user.Email, user.Phone, user.Password, user.Age, user.Gender, id)
	if err != nil {
		log.Println("Error updating user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	sqlStatement := `DELETE FROM users WHERE id = ?`
	_, err := db.GetDB().Exec(sqlStatement, id)
	if err != nil {
		log.Println("Error deleting user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
