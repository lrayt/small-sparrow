package handler

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type HttpHandler struct {
	router *gin.Engine
	srv    *http.Server
}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{
		router: nil,
		srv:    nil,
	}
}

func (h *HttpHandler) Run() error {
	log.Println("------")
	h.srv = &http.Server{
		Addr:    ":8080",
		Handler: h.router,
	}
	if err := h.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (h HttpHandler) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return h.srv.Shutdown(ctx)
}
