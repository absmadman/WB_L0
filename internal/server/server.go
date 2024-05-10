package server

import (
	"WB_L0/internal/entities"
	"WB_L0/pkg/cache"
	"WB_L0/pkg/config"
	"WB_L0/pkg/db"
	"context"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"os/signal"
	"strconv"
	"syscall"
	"time"
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
			tmpl, _ := template.New("page").Parse(outOfRangePage)
			tmpl.Execute(ctx.Writer, nil)
			return
		}
	}
	h.cache.Cache.Add(id, item.Message)
	total := h.database.GetTotal()
	pageData := entities.NewPage(string(item.Message), id, total, id+1, id-1, total)
	tmpl, _ := template.New("page").Parse(page)
	tmpl.Execute(ctx.Writer, pageData)
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
	// для graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	h.router.GET("/hello", h.hello)
	h.router.GET("/", h.getDataById)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: h.router,
	}
	//graceful shutdown part
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	<-ctx.Done()
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	log.Println("Server exiting")
}
