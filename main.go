package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/heimdahl-xyz/forseti/repositories"
	"github.com/heimdahl-xyz/forseti/types"
	"log"
	"net/http"
	"os"
)

type ForsetiServer struct {
	http *gin.Engine
	repo types.Processor
}

func NewForsetiServer() (*ForsetiServer, error) {
	r := gin.Default()

	db := os.Getenv("FORSETI_DB")
	pgdb, err := repositories.NewPostgresRepository(db)
	if err != nil {
		return nil, fmt.Errorf("failed to setup db %s", err)
	}

	s := &ForsetiServer{
		http: r,
		repo: pgdb,
	}

	r.GET("/v1/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"ok": true})
	})
	r.POST("/v1/transfers", s.CollectTransfer)
	return s, nil
}

func (s *ForsetiServer) CollectTransfer(ctx *gin.Context) {
	var ft types.FTMessage

	if err := ctx.ShouldBindJSON(&ft); err != nil {
		log.Printf("failed to parse transfer message %s", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.repo.ProcessTransfer(&ft); err != nil {
		log.Printf("failed to save transfer message %s", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"ok": true})
}

func (s *ForsetiServer) Start() error {
	return s.http.Run(":9009")
}

func main() {
	s, err := NewForsetiServer()
	if err != nil {
		log.Fatal(err)
	}

	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
