package main

import (
	"final-go-back-III/cmd/server/handler"
	"final-go-back-III/internal/odontologo"
	"final-go-back-III/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	//if err := godotenv.Load("../.././.env"); err != nil {
	//	panic("Error loading .env file: " + err.Error())
	//}

	storage := db.StorageDB

	repo := odontologo.NewMySQLRepository(storage)
	service := odontologo.NewService(repo)
	odontologoHandler := handler.NewOdontologoHandler(service)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	odontologos := r.Group("/odontologos")
	{
		odontologos.GET("/all", odontologoHandler.GetAll())
	}
	r.Run(":8080")
}
