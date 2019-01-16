package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type Person struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name`
	LastName  string `json:"last_name`
	Avatar    string `json:avatar`
	Age       uint   `json:age`
}

func main() {

	db, err = gorm.Open("sqlite3", "./assets/db/data.db") //TODO : move using configuration file
	if err != nil {
		panic("Failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Person{})

	r := gin.Default()
	r.GET("/", Home)
	r.GET("/persons", GetPersons)
	r.GET("/persons/:id", GetPersonsById)
	r.POST("/persons", CreatePersons)
	r.PUT("/persons/:id", UpdatePersonsById)
	r.DELETE("/persons/:id", DeletePersonsById)
	// r.Run("localhost:3000")
	r.Run(":3000")
}

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"app":     "Persons Data",
		"version": 1,
	})
}

func GetPersons(c *gin.Context) {
	var people []Person
	if err := db.Find(&people).Error; err != nil {
		RenderResponseArray(c, 404, nil)
		fmt.Println(err)
	} else {
		RenderResponseArray(c, 200, people)
	}
}

func GetPersonsById(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		RenderResponse(c, 404, person)
		fmt.Println(err)
	} else {
		RenderResponse(c, 200, person)
	}
}

func CreatePersons(c *gin.Context) {
	var person Person
	c.BindJSON(&person)

	db.Create(&person)
	RenderResponse(c, 200, person)
}

func UpdatePersonsById(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		// RenderResponse(c, 404, person)
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&person)
	db.Save(&person)
	RenderResponse(c, 200, person)
}

func DeletePersonsById(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person
	data := db.Where("id = ?", id).Delete(&person)
	fmt.Println(data)
	RenderResponse(c, 200, person)
}

func RenderResponse(c *gin.Context, code int, person Person) {
	c.JSON(code, gin.H{
		"msg":  GetMessageByCode(code),
		"code": code,
		"data": person,
	})
}

func RenderResponseArray(c *gin.Context, code int, people []Person) {
	c.JSON(code, gin.H{
		"msg":  GetMessageByCode(code),
		"code": code,
		"data": people,
	})
}

func GetMessageByCode(code int) string {
	messageMap := map[int]string{
		200: "success",
		404: "not found",
		500: "internal server error",
	}
	return messageMap[code]
}
