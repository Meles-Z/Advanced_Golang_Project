package app

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	handlers "github.com/meles-z/edusysnc_grpc/internal/app/handlers/student"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, server Server) {
	apiMerchant := e.Group("/student")
	apiMerchant.Use(DBTransactionMidleware(server.DB))
	apiMerchant.POST("/create", handlers.CreateStudent(server.studentService))
}

func DBTransactionMidleware(db *gorm.DB) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderServer, "Echo/4.13.3")
			if db == nil {
				log.Fatalln("Db not found")
			}
			txHandle := db.Begin()
			defer func() {
				if r := recover(); r != nil {
					txHandle.Rollback()
				}

			}()
			c.Set("db_trx", txHandle)
			result := next(c)

			if StatusInList(c.Response().Status, []int{http.StatusOK, http.StatusCreated}) {
				if err := txHandle.Commit().Error; err != nil {
					return result
				}
				return result
			} else {
				txHandle.Rollback()
				return result
			}
		}
	}
}

// func AdminAuthMiddleware(c configs.Configration) echo.MiddlewareFunc {
// 	signingKey := c.Auth.Secret

// }
func StatusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return false
}