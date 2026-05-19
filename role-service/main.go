package main 

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type user struct {
	Email string `"json:email"`
	Password string `json:"password"`
	Role string `json:"role"`
}

var users = []user {
	{Email:"admin@admin.com", Password:"admin@123", Role:"admin"},
	{Email:"operator@operator.com", Password:"operator@123", Role:"operator"},
}



func main() {
	router := gin.Default()

	router.GET("/getUsers", getAllUsers)
	router.POST("/addUser", createUser)
	router.Run("localhost:8080")
}


func getAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}


func createUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return 
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)

}