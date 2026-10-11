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

	server.PUT("/students/:id", func(c *gin.Context) {

		id := c.Param("id")

		var student Student

		if err := db.First(&student, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Student not found",
			})
			return
		}

		var input Student

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Request",
			})
			return
		}

		student.Name = input.Name
		student.Salary = input.Salary

		if err := db.Save(&student).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Update Failed",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Student Updated Successfully",
			"data":    student,
		})
	})

	server.Run(":8080")
}
