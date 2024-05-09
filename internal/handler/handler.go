package handler

import (
	"WB_L0/pkg/config"
	"WB_L0/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	router   *gin.Engine
	database *db.Database
}

func NewHandler(engine *gin.Engine, cfg *config.Config) *Handler {
	return &Handler{
		router:   engine,
		database: db.NewDatabase(cfg),
	}
}

func (h *Handler) hello(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, `message: Hello World!`)
}

func (h *Handler) getDataById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid parameter"})
		return
	}
	h.database.SelectById(id)
}

func (h *Handler) Serve() {
	h.router.GET("/hello", h.hello)
	h.router.GET("/")

	err := h.router.Run("localhost:8080")
	if err != nil {
		log.Println("Error running server")
	}
}
