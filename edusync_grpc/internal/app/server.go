package app

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	c "github.com/meles-z/edusysnc_grpc/configs"
	"github.com/meles-z/edusysnc_grpc/internal/app/db"
	"gorm.io/gorm"
)

type IServer interface {
	Start() error
}
type Server struct {
	DB  *gorm.DB
	cfg c.Confgration
}

func NewServer(cfg c.Confgration) IServer {
	mainDb, err := db.InitDB(&cfg.DB)
	if err != nil {
		log.Fatalf("ERROR to initialize database:%s", err.Error())
	}
	return &Server{
		DB:  mainDb,
		cfg: cfg,
	}
}

func (s *Server) Start() error {
	e := echo.New()
	return e.Start(fmt.Sprintf("Server start on:%d", s.cfg.DB.Port))
}
