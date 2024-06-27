package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	r := gin.Default()
	r.GET("data", func(c *gin.Context) {
		time.Sleep(time.Second * 30)
		c.JSON(http.StatusOK, gin.H{"status": 1000})
	})
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("err:%s\n", err.Error())
	}
}

func TestDefaultClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*35)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://127.0.0.1:8080/api/v1/proxy/baidu", nil)
	if err != nil {
		log.Fatalf("New Request Err:%s\n", err.Error())
	}
	res, err1 := http.DefaultClient.Do(req)
	if err1 != nil {
		log.Fatalf("Do Request Err:%s\n", err1.Error())
	}
	data, err2 := io.ReadAll(res.Body)
	if err2 != nil {
		log.Fatalf("Read data Err:%s\n", err2.Error())
	}
	t.Logf("get data %s\n", data)
}

func TestClient(t *testing.T) {
	client := http.Client{Timeout: time.Second * 5}
	res, err := client.Get("http://127.0.0.1:8080/api/v1/proxy/baidu")
	if err != nil {
		log.Fatalf("get data err:%s\n", err.Error())
	}
	data, err2 := io.ReadAll(res.Body)
	if err2 != nil {
		log.Fatalf("Read data Err:%s\n", err2.Error())
	}
	t.Logf("get data %s\n", data)
}
