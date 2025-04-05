package controller

import (
	"database/sql"
	"gin-app/db"
	"gin-app/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBanner(c *gin.Context) {
	var banner models.Banner

	if err := c.ShouldBindJSON(&banner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `INSERT INTO banners (name, src, link)
                     VALUES (?, ?, ?)`
	result, err := db.GetDB().Exec(sqlStatement, banner.Name, banner.Src, banner.Link)
	if err != nil {
		log.Println("Error creating banner:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	banner.ID = int(id)
	c.JSON(http.StatusCreated, banner)
}

func GetAllBanners(c *gin.Context) {
	sqlStatement := `SELECT id, name, src, link FROM banners`
	rows, err := db.GetDB().Query(sqlStatement)
	if err != nil {
		log.Println("Error fetching banners:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch banners"})
		return
	}
	defer rows.Close()

	var banners []models.Banner
	for rows.Next() {
		var banner models.Banner
		if err := rows.Scan(&banner.ID, &banner.Name, &banner.Src, &banner.Link); err != nil {
			log.Println("Error scanning row:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read banner data"})
			return
		}
		banners = append(banners, banner)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch banners"})
		return
	}

	c.JSON(http.StatusOK, banners)
}

func GetBanner(c *gin.Context) {
	id := c.Param("id")

	var banner models.Banner
	sqlStatement := `SELECT id, name, src, link FROM banners WHERE id = ?`
	err := db.GetDB().QueryRow(sqlStatement, id).Scan(&banner.ID, &banner.Name, &banner.Src, &banner.Link)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "banner not found"})
		} else {
			log.Println("Error fetching banner:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch banner"})
		}
		return
	}

	c.JSON(http.StatusOK, banner)
}

func UpdateBanner(c *gin.Context) {
	id := c.Param("id")
	var banner models.Banner

	if err := c.ShouldBindJSON(&banner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `UPDATE banners SET name = ?, src = ?, link = ? WHERE id = ?`
	_, err := db.GetDB().Exec(sqlStatement, banner.Name, banner.Src, banner.Link, id)
	if err != nil {
		log.Println("Error updating banner:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update banner"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Banner updated successfully"})
}

func DeleteBanner(c *gin.Context) {
	id := c.Param("id")

	sqlStatement := `DELETE FROM banners WHERE id = ?`
	_, err := db.GetDB().Exec(sqlStatement, id)
	if err != nil {
		log.Println("Error deleting banner:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete banner"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "banner deleted successfully"})
}
