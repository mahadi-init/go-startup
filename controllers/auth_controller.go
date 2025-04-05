package controller

import (
	"gin-app/db"
	"gin-app/models"
	"gin-app/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
		return
	}
	user.Password = string(hashedPassword)

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

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `SELECT id, name, email, phone, password, age, gender FROM users WHERE email = ?`
	row := db.GetDB().QueryRow(sqlStatement, user.Email)

	var storedUser models.User
	if err := row.Scan(&storedUser.ID, &storedUser.Name, &storedUser.Email, &storedUser.Phone, &storedUser.Password, &storedUser.Age, &storedUser.Gender); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Create JWT token
	token, err := utils.CreateJWTToken(storedUser)
	if err != nil {
		log.Println("Error creating JWT token:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  storedUser,
	})
}
