package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Handler struct {
	router *gin.Engine
}

func NewHandler(engine *gin.Engine) *Handler {
	return &Handler{
		router: engine,
	}
}

func (h *Handler) hello(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, `message: Hello World!`)
}

func (h *Handler) Serve() {
	h.router.GET("/v1/hello", h.hello)

	err := h.router.Run("localhost:8080")

	if err != nil {
		log.Println("Error running server")
	}

}
