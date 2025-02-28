package services

import (
	"context"

	"github.com/meles-z/edusysnc_grpc/internal/app/repositories"
	"github.com/meles-z/edusysnc_grpc/proto/student"
	"gorm.io/gorm"
)

type IStudentService interface {
	CreateStudent(ctx context.Context, req *student.CreateStudentRequest) (*student.CreateStudentResponse, error)
	GetAllStudent(ctx context.Context, req *student.GetAllStudentsRequest) (*student.GetAllStudentsResponse, error)
	WithTrx(*gorm.DB) StudentService
}

type StudentService struct {
	studentRepo repositories.StudentRepository
}

func (s StudentService) WithTrx(db *gorm.DB) StudentService {
	s.studentRepo = s.studentRepo.WithTrx(db)
	return s
}

func NewStudentService(studRepo repositories.StudentRepository) (IStudentService, error) {
	return StudentService{
		studentRepo: studRepo,
	}, nil
}

func (srv StudentService) CreateStudent(ctx context.Context, req *student.CreateStudentRequest) (*student.CreateStudentResponse, error) {
	newStud, err := srv.studentRepo.CreateStudent(ctx, req)
	if err != nil {
		return nil, err
	}
	return newStud, nil
}

func (srv StudentService) GetAllStudent(ctx context.Context, req *student.GetAllStudentsRequest) (*student.GetAllStudentsResponse, error) {
	students, err := srv.studentRepo.GetAllStudent(ctx, req)
	if err != nil {
		return nil, err
	}
	return students, nil
}
