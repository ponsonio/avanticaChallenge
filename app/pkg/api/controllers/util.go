package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SetResponseWithError(c *gin.Context, err error) {
	switch err.(type) {
	//here we should set the adecuate http code/response according to the error
	default:
		//not a good idea to show unidentified error info, this is only an exercise
		SetResponse(c, http.StatusInternalServerError, err.Error())
		log.Println("Unexpected error %v", err)
	}
}

// SetResponseWithBody ...
func SetResponseWithBody(c *gin.Context, code int, body interface{}) {
	c.JSON(code, body)
}

// SetResponse by default
func SetResponse(c *gin.Context, code int, message interface{}) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
