package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	r := gin.Default()
	r.GET("data", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": 1000})
	})
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("err:%s\n", err.Error())
	}
}
