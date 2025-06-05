package handler

import "github.com/gin-gonic/gin"

func testhandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}
