package main

import (
	"net/http"

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

	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"

db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
if err != nil {
panic(err)
}

if err := db.AutoMigrate(&Student{}); err != nil {
panic(err)
}

server := gin.Default()

	
server.POST("/students", func(c *gin.Context) {

var student Student

if err := c.ShouldBindJSON(&student); err != nil {
c.JSON(http.StatusBadRequest, gin.H{
"message": "Invalid Request",
})
return
}

if err := db.Create(&student).Error; err != nil {
c.JSON(http.StatusInternalServerError, gin.H{
"message": "Failed to create student",
})
return
}

c.JSON(http.StatusCreated, gin.H{
"message": "Student Created Successfully",
"data":    student,
})

})
server.GET("/students", func(c *gin.Context) {

var students []Student

if err := db.Find(&students).Error; err != nil {
c.JSON(http.StatusInternalServerError, gin.H{
"message": "failed to fetch students",
})
return
}

c.JSON(http.StatusOK, gin.H{
"message": "students fetched successfully",
"data":    students,
})

})

server.Run(":8080")
}
