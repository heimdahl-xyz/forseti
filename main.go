package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"net/http"
)

type FTMessage struct {
	Timestamp    int64    `json:"timestamp"`
	FromAddress  string   `json:"from_address"`
	FromOwner    string   `json:"from_owner,omitempty"`
	ToAddress    string   `json:"to_address"`
	ToOwner      string   `json:"to_owner,omitempty"`
	Amount       *big.Int `json:"amount"`
	TokenAddress string   `json:"token_address"`
	Symbol       string   `json:"symbol"`
	Chain        string   `json:"chain"`
	Network      string   `json:"network"`
	TxHash       string   `json:"tx_hash"`
	Decimals     uint8    `json:"decimals"`
	Position     uint64   `json:"position"`
}

type CollectorRepository interface {
	SaveTransfer(ft *FTMessage) error
}

type DummyRepo struct {
}

func (d *DummyRepo) SaveTransfer(ft *FTMessage) error {
	log.Printf("saved message %+v", ft)
	return nil
}

type ForsetiServer struct {
	http *gin.Engine
	repo CollectorRepository
}

func NewForsetiServer() (*ForsetiServer, error) {
	r := gin.Default()

	s := &ForsetiServer{
		http: r,
		repo: &DummyRepo{},
	}

	r.POST("/v1/transfers", s.CollectTransfer)
	return s, nil
}

func (s *ForsetiServer) CollectTransfer(ctx *gin.Context) {
	var ft FTMessage

	if err := ctx.ShouldBindJSON(&ft); err != nil {
		log.Println("failed to parse transfer message %s", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.repo.SaveTransfer(&ft); err != nil {
		log.Println("failed to save transfer message %s", err)
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
