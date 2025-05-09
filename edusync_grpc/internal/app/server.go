package app

import (
	"fmt"
	"log"
	"net"

	c "github.com/meles-z/edusysnc_grpc/configs"
	"github.com/meles-z/edusysnc_grpc/internal/app/db"
	"github.com/meles-z/edusysnc_grpc/internal/app/repositories"
	"github.com/meles-z/edusysnc_grpc/internal/app/services"
	"github.com/meles-z/edusysnc_grpc/proto/student"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type StudentServerImpl struct {
	student.UnimplementedStudentServiceServer
	studentService services.IStudentService
}

type IServer interface {
	Start() error
}
type Server struct {
	DB             *gorm.DB
	cfg            c.Confgration
	studentService services.IStudentService
	grpcServer     *grpc.Server
}

func NewServer(cfg c.Confgration) IServer {
	mainDb, err := db.InitDB(&cfg.DB)
	if err != nil {
		log.Fatalf("ERROR to initialize database:%s", err.Error())
	}
	grpcServer := grpc.NewServer()
	studentRepo := repositories.NewStudentRepository(mainDb)
	// service
	studentService, err := services.NewStudentService(studentRepo)
	if err != nil {
		log.Fatal("Error to init student services", err)
	}
	return &Server{
		DB:             mainDb,
		cfg:            cfg,
		studentService: studentService,
		grpcServer:     grpcServer,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.Server.GRPC_PORT))
	if err != nil {
		log.Fatalf("failed to listen on port %d: %v", s.cfg.Server.GRPC_PORT, err)
	}

	// Register gRPC Services
	student.RegisterStudentServiceServer(s.grpcServer, &StudentServerImpl{
		studentService: s.studentService,
	})

	log.Printf("gRPC server listening at %v", listener.Addr())

	// Start gRPC Server
	if err := s.grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}

	return nil
}
