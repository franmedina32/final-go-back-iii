package main

import (
	"final-go-back-III/cmd/server/handler"
	"final-go-back-III/internal/odontologo"
	"final-go-back-III/internal/paciente"
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

	repoPaciente := paciente.NewMySQLRepository(storage)
	servicePaciente := paciente.NewService(repoPaciente)
	pacienteHandler := handler.NewPacienteHandler(servicePaciente)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	odontologos := r.Group("/odontologos")
	{
		odontologos.GET("/all", odontologoHandler.GetAll())
	}
	pacientes := r.Group("/pacientes")
	{
		pacientes.GET("/all", pacienteHandler.GetAll())
		pacientes.GET("/:id", pacienteHandler.GetByID())
		pacientes.POST("/create", pacienteHandler.Create())
		pacientes.PUT("/update/:id", pacienteHandler.Update())
		pacientes.PATCH("/updateField/:id", pacienteHandler.UpdateField())
		pacientes.DELETE("/delete/:id", pacienteHandler.Delete())
	}
	r.Run(":8080")
}
