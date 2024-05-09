package server

import (
	"WB_L0/internal/entities"
	"WB_L0/pkg/cache"
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
	cache    *cache.Cache
}

func NewHandler(db *db.Database, engine *gin.Engine, cache *cache.Cache, cfg *config.Config) *Handler {
	return &Handler{
		router:   engine,
		database: db,
		cache:    cache,
	}
}

func (h *Handler) hello(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, `message: Hello World!`)
}

func (h *Handler) getDataById(ctx *gin.Context) {
	var item entities.Item
	queryId := ctx.Query("id")
	id, err := strconv.Atoi(queryId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid parameter"})
		return
	}
	msg, ok := h.cache.Cache.Get(id)
	if ok {
		item.Id = queryId
		item.Message = msg
	} else {
		item, err = h.database.SelectById(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "id is not exist"})
			return
		}
	}
	h.cache.Cache.Add(id, item.Message)
	ctx.IndentedJSON(http.StatusOK, gin.H{"id": item.Id, "message": item.Message})
}

func (h *Handler) InitCache(cfg *config.Config) {
	for i := 1; i < cfg.CacheSize; i++ {
		item, err := h.database.SelectById(i)
		if err != nil {
			return
		}
		id, _ := strconv.Atoi(item.Id)
		h.cache.Cache.Add(id, item.Message)
	}
}

func (h *Handler) Serve() {
	h.router.GET("/hello", h.hello)
	h.router.GET("/", h.getDataById)
	err := h.router.Run("localhost:8080")
	if err != nil {
		log.Println("Server: error running server")
	}
}
