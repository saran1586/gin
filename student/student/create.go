package main

import(
"net/http" 
"github.com/gin-gonic/gin"
"gorm.io/driver/postgres"
"gorm.io/gorm"
)

type Student struct{
ID int  `gorm:"primaryKey"  json:"id"`
Name string `json:"name"`
Salary float64 `json:"salary"`
}

func main(){


name := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
db,err := gorm.Open(postgres.Open(name),&gorm.Config{})

if err!=nil{
panic(err)
}
db.AutoMigrate(&Student{})


server := gin.Default()
server.POST("/test",func(a *gin.Context){
var student Student
if err:=a.BindJSON(&student);err!=nil{
a.JSON(http.StatusBadRequest,gin.H{

"message" : "Invalid",

})
return
}

db.Create(&student)

if err := db.Create(&student).Error; err != nil {
	a.JSON(http.StatusInternalServerError, gin.H{
		"message": "Failed to create student",
	})
	return




a.JSON(http.StatusCreated,gin.H{
"message" : "student is created ",
"data" :student,                      
}) 
})

server.Run()
}
