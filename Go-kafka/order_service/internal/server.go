package internal

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/meles-z/go-kafa/order_service/configs"
	"github.com/meles-z/go-kafa/order_service/internal/app/db"
	"gorm.io/gorm"
)

type IServer interface {
	Start() error
}
type Server struct {
	DB  *gorm.DB
	Cfg configs.Config
}

func NewServer(cfg configs.Config) IServer {
	mainDb, err := db.InitDB(cfg.DB)
	if err != nil {
		log.Fatalf("DB INITIALIZE ERROR :%s", err.Error())
	}
	return &Server{
		DB:  mainDb,
		Cfg: cfg,
	}
}

func (s *Server) Start() error {
	e := echo.New()
	return e.Start(fmt.Sprintf(":%d", s.Cfg.Server.Port))
}
