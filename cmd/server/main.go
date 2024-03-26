package main

import (
	"final-go-back-III/cmd/server/handler"
	"final-go-back-III/internal/odontologo"
	"final-go-back-III/internal/paciente"
	turnos2 "final-go-back-III/internal/turnos"
	"final-go-back-III/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	//if err := godotenv.Load("../.././.env"); err != nil {
	//	panic("Error loading .env file: " + err.Error())
	//}

	storage := db.StorageDB

	odontologoRepo := odontologo.NewMySQLRepository(storage)
	odontologoService := odontologo.NewService(odontologoRepo)
	odontologoHandler := handler.NewOdontologoHandler(odontologoService)

	pacienteRepo := paciente.NewMySQLRepository(storage)
	pacienteService := paciente.NewService(pacienteRepo)
	pacienteHandler := handler.NewPacienteHandler(pacienteService)

	turnoRepo := turnos2.NewMySQLRepository(storage)
	turnoService := turnos2.NewService(turnoRepo, pacienteRepo, odontologoRepo)
	turnoHandler := handler.NewTurnoHandler(turnoService)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	odontologos := r.Group("/odontologos")
	{
		odontologos.GET("/all", odontologoHandler.GetAll())
		odontologos.GET("/:id", odontologoHandler.GetByID())
		odontologos.POST("/create", odontologoHandler.Create())
		odontologos.PUT("/update/:id", odontologoHandler.Update())
		odontologos.PATCH("/updateField/:id", odontologoHandler.UpdateField())
		odontologos.DELETE("/delete/:id", odontologoHandler.Delete())
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
	turnos := r.Group("/turnos")
	{
		turnos.GET("/all", turnoHandler.GetAll())
		turnos.GET("/:id", turnoHandler.GetByID())
		turnos.GET("/dni", turnoHandler.GetByPacienteDNI())
		turnos.POST("/create", turnoHandler.Create())
		turnos.POST("/create/dniAndMatricula", turnoHandler.CreateByDniAndMatricula())
		turnos.PUT("/update/:id", turnoHandler.Update())
		turnos.PATCH("/updateField/:id", turnoHandler.UpdateField())
		turnos.DELETE("/delete/:id", turnoHandler.Delete())

	}
	r.Run(":8080")
}
