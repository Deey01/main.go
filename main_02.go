package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Employee API Methods
	router.GET("/employee", GETEmployee)
	router.POST("/employee", POSTEmployee) 
	router.PUT("/employee", PUTEmployee)
	router.DELETE("/employee", DELETEEmployee)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func GETEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee GET Method!",
	})
}

func POSTEmployee(c *gin.Context) {  // Corrected the function name
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee POST Method!",
	})
}

func PUTEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee PUT Method!",
	})
}

func DELETEEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee DELETE Method!",
	})
}
