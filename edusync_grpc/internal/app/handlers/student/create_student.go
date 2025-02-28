package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/meles-z/edusysnc_grpc/internal/app/models"
	"github.com/meles-z/edusysnc_grpc/internal/app/services"
	"github.com/meles-z/edusysnc_grpc/proto/student"
	"gorm.io/gorm"
)

type FailureResponse struct {
	Reason string `json:"reason"`
}

func CreateStudent(studentService services.IStudentService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var stud models.Student
		if err := c.Bind(&stud); err != nil {
			return c.JSON(http.StatusUnprocessableEntity, &FailureResponse{
				Reason: "Invalid request/JSON",
			})
		}
		// Get transaction from context
		tx, ok := c.Get("db_trx").(*gorm.DB)
		if !ok || tx == nil {
			return c.JSON(http.StatusInternalServerError, &FailureResponse{
				Reason: "Database transaction not initialized",
			})
		}
		data := &student.Student{
			Name:        stud.Name,
			Email:       stud.Email,
			PhoneNumber: stud.PhoneNumber,
		}
		studentSvc := studentService.WithTrx(tx)
		studentRecord, err := studentSvc.CreateStudent(c.Request().Context(), &student.CreateStudentRequest{
			Student: data,
		})
		if err != nil {
			fmt.Printf("CreateEvent Error: %v\n", err)
			return c.JSON(http.StatusInternalServerError, &FailureResponse{
				Reason: err.Error(),
			})
		}
		return c.JSON(http.StatusCreated, studentRecord)
	}
}
