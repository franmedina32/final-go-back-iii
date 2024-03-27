package handler

import (
	"final-go-back-III/internal/domain"
	"final-go-back-III/internal/paciente"
	"final-go-back-III/pkg/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type pacienteHandler struct {
	s paciente.Service
}

func NewPacienteHandler(s paciente.Service) *pacienteHandler {
	return &pacienteHandler{
		s: s,
	}
}

// GetAll godoc
// @Summary Get all patients
// @Description Retrieve all patients
// @Tags Paciente
// @Produce json
// @Success 200 {array} domain.Paciente "OK"
// @Router /pacientes/all [get]
func (h *pacienteHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		pacientes, err := h.s.GetAll()
		if err != nil {
			panic(err)
		}
		web.Success(c, http.StatusOK, pacientes)
	}
}

// GetByID godoc
// @Summary Get a patient by ID
// @Description Retrieve a patient by ID
// @Tags Paciente
// @Produce json
// @Param id path int true "Patient ID"
// @Success 200 {object} domain.Paciente "OK"
// @Router /pacientes/{id} [get]
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

// Create godoc
// @Summary Create a new patient
// @Description Create a new patient
// @Tags Paciente
// @Accept json
// @Produce json
// @Param patient body domain.Paciente true "Patient object"
// @Param token header string true "TOKEN"
// @Success 201 {object} domain.Paciente "Created"
// @Router /pacientes/create [post]
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

// Update godoc
// @Summary Update an existing patient
// @Description Update an existing patient
// @Tags Paciente
// @Accept json
// @Produce json
// @Param id path int true "Patient ID"
// @Param patient body domain.Paciente true "Patient object"
// @Param token header string true "TOKEN"
// @Success 200 {object} domain.Paciente "Updated"
// @Router /pacientes/update/{id} [put]
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

// UpdateField godoc
// @Summary Update a specific field of a patient
// @Description Update a specific field of a patient
// @Tags Paciente
// @Accept json
// @Produce json
// @Param id path int true "Patient ID"
// @Param field_name body string true "Field name"
// @Param value body string true "New value"
// @Param token header string true "TOKEN"
// @Success 200 {string} string "Field updated successfully"
// @Router /pacientes/updateField/{id} [patch]
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

// Delete godoc
// @Summary Delete a patient
// @Description Delete a patient by ID
// @Tags Paciente
// @Param id path int true "Patient ID"
// @Param token header string true "TOKEN"
// @Success 200 {string} string "Paciente deleted successfully"
// @Router /pacientes/delete/{id} [delete]
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
