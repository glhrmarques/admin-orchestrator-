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
	{Email:"admin@admin.com", Password:"123", Role:"admin"},
	{Email:"operator@operator.com", Password:"operator@123", Role:"operator"},
}



func main() {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	router.GET("/getUsers", getAllUsers)
	router.POST("/addUser", createUser)
	router.POST("/login", loginUser)
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

func loginUser (c *gin.Context){
	var credentials struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	for _, u := range users {
		if u.Email == credentials.Email && u.Password == credentials.Password {
			c.JSON(http.StatusOK, gin.H{"role": u.Role})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or passoword"})
}