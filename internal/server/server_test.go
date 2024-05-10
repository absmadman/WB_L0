package server

import (
	"WB_L0/pkg/cache"
	"WB_L0/pkg/config"
	"WB_L0/pkg/db"
	nats "WB_L0/pkg/nats"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

var h *Handler
var s *nats.Streaming

func TestMain(m *testing.M) {
	log.Println("START TESTING")
	cfg := config.Config{
		"postgres",
		"postgres",
		"postgres",
		"receiver",
		"test-cluster",
		256,
	}

	db := db.NewDatabase(&cfg)
	cache := cache.NewCache(&cfg)
	s, err := nats.NewStreaming(&cfg, db, cache)
	if err != nil {
		log.Fatalln("Nats: error connect nats streaming")
	}

	h = NewHandler(db, gin.Default(), cache, &cfg)
	h.InitCache(&cfg)

	err = db.CreateDatabase()
	if err != nil {
		log.Fatalln("Database: error creating table")
	}
	s.Sub()

	h.router.GET("/hello", h.hello)
	h.router.GET("/", h.getDataById)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: h.router,
	}
	go func() {
		srv.ListenAndServe()
	}()
	time.Sleep(time.Millisecond * 200)
	res := m.Run()
	time.Sleep(time.Millisecond * 200)
	srv.Close()
	if res == 0 {
		fmt.Println("TESTS PASSED SUCCESSFULLY")
	} else {
		fmt.Println("TESTS PASSED UNSUCCESSFULLY")
	}
}

func TestGetPositive(t *testing.T) {
	testConn, _ := stan.Connect("test-cluster", "publisher")
	data, _ := os.ReadFile("../../data/model.json")

	testConn.Publish("post", data)

	time.Sleep(time.Millisecond * 100)

	item, _ := h.database.SelectById(1)

	if item.Id != "1" || len(item.Message) == 0 {
		t.Error("Incorrect")
	} else {
		fmt.Println("successfully selected")
	}
	testConn.Close()
}

func TestGetNegative(t *testing.T) {
	testConn, _ := stan.Connect("test-cluster", "publisher")
	data, _ := os.ReadFile("../../data/model.json")

	testConn.Publish("post", data)

	time.Sleep(time.Millisecond * 100)

	_, err := h.database.SelectById(0)

	if err == nil {
		t.Error("incorrect")
	} else {
		fmt.Println("successfully selected")
	}
	testConn.Close()
}

func TestInsertPositive(t *testing.T) {
	testConn, _ := stan.Connect("test-cluster", "publisher")
	data, _ := os.ReadFile("../../data/model.json")

	before := h.database.GetTotal()
	cacheBefore := h.cache.Cache.Len()
	fmt.Println("rows before: ", before)
	fmt.Println("in cache before: ", before)
	testConn.Publish("post", data)

	time.Sleep(time.Millisecond * 100)

	after := h.database.GetTotal()
	cacheAfter := h.cache.Cache.Len()
	fmt.Println("rows after: ", after)
	fmt.Println("in cache after: ", cacheAfter)
	if before+1 != after || cacheBefore+1 != cacheAfter {
		fmt.Println(before, after, cacheBefore, cacheAfter)
		t.Error("incorrect")
	}
	testConn.Close()
}

func TestInsertNegative(t *testing.T) {
	testConn, _ := stan.Connect("test-cluster", "publisher")
	data, _ := os.ReadFile("../../data/invalid_model.json")

	before := h.database.GetTotal()
	cacheBefore := h.cache.Cache.Len()
	fmt.Println("rows before: ", before)
	fmt.Println("in cache before: ", before)
	testConn.Publish("post", data)

	time.Sleep(time.Millisecond * 100)

	after := h.database.GetTotal()
	cacheAfter := h.cache.Cache.Len()
	fmt.Println("rows after: ", after)
	fmt.Println("in cache after: ", cacheAfter)
	if before != after || cacheBefore != cacheAfter {
		fmt.Println(before, after, cacheBefore, cacheAfter)
		t.Error("incorrect")
	}
	testConn.Close()
}
