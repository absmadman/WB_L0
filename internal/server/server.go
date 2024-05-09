package server

import (
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

const a = `"{
  "order_uid": "b563feb7b2b84b6test",
  "track_number": "WBILMTESTTRACK",
  "entry": "WBIL",
  "delivery": {
    "name": "Test Testov",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Ploshad Mira 15",
    "region": "Kraiot",
    "email": "test@gmail.com"
  },
  "payment": {
    "transaction": "b563feb7b2b84b6test",
    "request_id": "",
    "currency": "USD",
    "provider": "wbpay",
    "amount": 1817,
    "payment_dt": 1637907727,
    "bank": "alpha",
    "delivery_cost": 1500,
    "goods_total": 317,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}"`

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
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid parameter"})
		return
	}
	item, err := h.database.SelectById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "id is not exist"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"id": item.Id, "message": item.Message})
}

func (h *Handler) putData(ctx *gin.Context) {
	h.database.InsertDatabase(a)
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
	h.router.POST("/", h.putData)

	err := h.router.Run("localhost:8080")
	if err != nil {
		log.Println("Error running server")
	}
}
