package internal

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/meles-z/go-kafa/order_service/configs"
	"github.com/meles-z/go-kafa/order_service/internal/app/db"
	"github.com/meles-z/go-kafa/order_service/internal/app/kafka"
	repository "github.com/meles-z/go-kafa/order_service/internal/app/respository"
	"github.com/meles-z/go-kafa/order_service/internal/app/service"
	"gorm.io/gorm"
)

const (
	kafkaBroker = "localhost:9092"
	kafkaTopic  = "order_created"
)

type IServer interface {
	Start() error
}

type Server struct {
	DB           *gorm.DB
	Cfg          configs.Config
	orderService service.OrderUseCase
}

func NewServer(cfg configs.Config) IServer {
	mainDb, err := db.InitDB(cfg.DB)
	if err != nil {
		log.Fatalf("DB INITIALIZE ERROR :%s", err.Error())
	}

	orderRepo := repository.NewOrderRepository(mainDb)
	kafkaProducer := kafka.NewKafkaProducer(kafkaBroker, kafkaTopic)
	orderUseCase := service.NewOrderUseCase(orderRepo, kafkaProducer)

	return &Server{
		DB:           mainDb,
		Cfg:          cfg,
		orderService: *orderUseCase,
	}
}

func (s *Server) Start() error {
	e := echo.New()

	Router(e, s)

	port := fmt.Sprintf(":%d", s.Cfg.Server.Port)
	log.Printf("Server is running on port %s...", port)
	return e.Start(port)
}
