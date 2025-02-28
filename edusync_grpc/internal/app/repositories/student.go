package repositories

import (
	"context"
	"fmt"

	"github.com/meles-z/edusysnc_grpc/internal/app/models"
	"github.com/meles-z/edusysnc_grpc/proto/student"
	"gorm.io/gorm"
)

type StudentRepository interface {
	CreateStudent(ctx context.Context, req *student.CreateStudentRequest) (*student.CreateStudentResponse, error)
	GetAllStudent(ctx context.Context, req *student.GetAllStudentsRequest) (*student.GetAllStudentsResponse, error)
	WithTrx(*gorm.DB) StudentRepository
}

type studentRepo struct {
	DB *gorm.DB
}

func (s *studentRepo) WithTrx(handleDb *gorm.DB) StudentRepository {
	if handleDb == nil {
		fmt.Println("Transaction of database is not found")
		return s
	}
	s.DB = handleDb
	return s
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepo{DB: db}
}
func (repo *studentRepo) CreateStudent(ctx context.Context, req *student.CreateStudentRequest) (*student.CreateStudentResponse, error) {
	fmt.Println("Create students")
	stud := req.GetStudent()

	data := models.Student{
		Name:        stud.GetName(),
		Email:       stud.GetEmail(),
		PhoneNumber: stud.GetPhoneNumber(),
	}
	if err := repo.DB.Create(&data).Error; err != nil {
		return nil, fmt.Errorf("error to created student:%s", err.Error())
	}

	return &student.CreateStudentResponse{
		Student: &student.Student{
			Name:        stud.GetName(),
			Email:       stud.GetEmail(),
			PhoneNumber: stud.GetPhoneNumber(),
		},
	}, nil
}
func (repo studentRepo) GetAllStudent(ctx context.Context, req *student.GetAllStudentsRequest) (*student.GetAllStudentsResponse, error) {
	fmt.Println("Returns all students")
	stud := []*student.Student{}
	if err := repo.DB.Find(&stud).Error; err != nil {
		return nil, fmt.Errorf("error: record not found:%s", err.Error())
	}
	return &student.GetAllStudentsResponse{
		Students: stud,
	}, nil
}
