package main

import (
	"course-bwastartup-backend/auth"
	"course-bwastartup-backend/handler"
	"course-bwastartup-backend/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	dsn := "root:F!rentia2818@tcp(localhost:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyfQ.NWDPizOkCqxnse2hDsHVG-KGPdWU50QDPRoEvGphXaQ")
	if err != nil {
		fmt.Println("error")
		fmt.Println("error")
		fmt.Println("error")
	}

	if token.Valid {
		fmt.Println("valid")
		fmt.Println("valid")
		fmt.Println("valid")
	} else {
		fmt.Println("invalid")
		fmt.Println("invalid")
		fmt.Println("invalid")
	}

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()
}
