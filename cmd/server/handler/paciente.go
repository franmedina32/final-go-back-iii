package handler

import (
	"final-go-back-III/internal/domain"
	"final-go-back-III/internal/paciente"
	"final-go-back-III/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type pacienteHandler struct {
	s paciente.Service
}

func NewPacienteHandler(s paciente.Service) *pacienteHandler {
	return &pacienteHandler{
		s: s,
	}
}

func (h *pacienteHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		pacientes, err := h.s.GetAll()
		if err != nil {
			panic(err)
		}
		web.Success(c, http.StatusOK, pacientes)
	}
}

func (h *pacienteHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}
		paciente, err := h.s.GetByID(id)
		if err != nil {
			panic(err)
		}
		web.Success(c, http.StatusOK, paciente)
	}
}

func (h *pacienteHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var p domain.Paciente
		if err := c.BindJSON(&p); err != nil {
			panic(err)
		}

		createdPaciente, err := h.s.Create(p)
		if err != nil {
			panic(err)
		}
		web.Success(c, http.StatusCreated, createdPaciente)
	}
}

func (h *pacienteHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}
		var p domain.Paciente
		if err := c.BindJSON(&p); err != nil {
			panic(err)
		}
		updatedPaciente, err := h.s.Update(id, p)
		if err != nil {
			panic(err)
		}
		web.Success(c, http.StatusOK, updatedPaciente)
	}
}

func (h *pacienteHandler) UpdateField() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}
		var requestBody struct {
			FieldName string      `json:"field_name"`
			Value     interface{} `json:"value"`
		}
		if err := c.BindJSON(&requestBody); err != nil {
			panic(err)
		}
		err = h.s.UpdateField(id, requestBody.FieldName, requestBody.Value)
		if err != nil {
			panic(err)
		}
		web.Success(c, http.StatusOK, "Field updated successfully")
	}
}

func (h *pacienteHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}
		err = h.s.Delete(id)
		if err != nil {
			panic(err)
		}
		web.Success(c, http.StatusOK, "Paciente deleted successfully")
	}
}
