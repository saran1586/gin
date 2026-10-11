package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Student struct {
	ID     int     `gorm:"primaryKey" json:"id"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
}

func main() {

	// Connection string comes from DATABASE_URL; the fallback is a local development database.
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&Student{}); err != nil {
		panic(err)
	}

	server := gin.Default()

	server.DELETE("/students/:id", func(c *gin.Context) {

		id := c.Param("id")

		var student Student

		if err := db.First(&student, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "student not found",
			})
			return
		}

		if err := db.Delete(&student).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "delete failed",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Student Deleted Successfully",
		})
	})

	server.Run(":8080")
}
