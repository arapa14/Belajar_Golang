package handler

import (
	"net/http"
	"try2/model"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
	return
}

func GetUsers(c *gin.Context) {
	dummy := []model.User{
		{ID: 1, Name: "Arapa", Age: 18},
		{ID: 2, Name: "Papoy", Age: 17},
	}

	c.JSON(http.StatusOK, dummy)
	return 
}

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user)
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error:" : err.Error()})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"message:" : "User created", "data:" : user})
	return 
}
