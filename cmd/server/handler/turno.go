package handler

import (
	"final-go-back-III/internal/domain"
	"final-go-back-III/internal/turnos"
	"final-go-back-III/pkg/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type turnoHandler struct {
	s turnos.Service
}

func NewTurnoHandler(s turnos.Service) *turnoHandler {
	return &turnoHandler{
		s: s,
	}
}

// GetAll godoc
// @Summary Get all turnos
// @Description Retrieve all turnos
// @Tags Turno
// @Produce json
// @Success 200 {array} domain.Turno "OK"
// @Router /turnos/all [get]
func (h *turnoHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		turnos, err := h.s.GetAll()
		if err != nil {
			panic(err)
		}
		web.Success(c, http.StatusOK, turnos)
	}
}

// GetByID godoc
// @Summary Get a turno by ID
// @Description Retrieve a turno by ID
// @Tags Turno
// @Produce json
// @Param id path int true "Turno ID"
// @Success 200 {object} domain.Turno "OK"
// @Router /turnos/{id} [get]
func (h *turnoHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}
		turno, err := h.s.GetByID(id)
		if err != nil {
			panic(err)
		}
		web.Success(c, http.StatusOK, turno)
	}
}

// GetByPacienteDNI godoc
// @Summary Get turnos by paciente's DNI
// @Description Retrieve turnos by paciente's DNI
// @Tags Turno
// @Produce json
// @Param dni query string true "Paciente DNI"
// @Success 200 {object} interface{} "OK"
// @Router /turnos/byDNI [get]
func (h *turnoHandler) GetByPacienteDNI() gin.HandlerFunc {
	return func(c *gin.Context) {
		dni := c.Query("dni")
		if dni == "" {
			panic("DNI parameter is required")
		}

		turnoDetail, err := h.s.GetByPacienteDNI(dni)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, turnoDetail)
	}
}

// Create godoc
// @Summary Create a new turno
// @Description Create a new turno
// @Tags Turno
// @Accept json
// @Produce json
// @Param turno body domain.Turno true "Turno object"
// @Param token header string true "TOKEN"
// @Success 201 {object} domain.Turno "Created"
// @Router /turnos/create [post]
func (h *turnoHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var t domain.Turno
		if err := c.BindJSON(&t); err != nil {
			panic(err)
		}
		createdTurno, err := h.s.Create(t)
		if err != nil {
			panic(err)
		}
		web.Success(c, http.StatusCreated, createdTurno)
	}
}

// CreateByDniAndMatricula godoc
// @Summary Create a new turno by Paciente's DNI and Odontologo's Matricula
// @Description Create a new turno by Paciente's DNI and Odontologo's Matricula
// @Tags Turno
// @Accept json
// @Produce json
// @Param turno body turnos.CreateTurnoData true "Turno data"
// @Param token query string true "TOKEN"
// @Success 201 {object} domain.Turno "Created"
// @Router /turnos/createByDniAndMatricula [post]
func (h *turnoHandler) CreateByDniAndMatricula() gin.HandlerFunc {
	return func(c *gin.Context) {
		var t turnos.CreateTurnoData
		if err := c.BindJSON(&t); err != nil {
			panic(err)
		}
		createdTurno, err := h.s.CreateByDniAndMatricula(t)
		if err != nil {
			panic(err)
		}
		web.Success(c, http.StatusCreated, createdTurno)
	}
}

// Update godoc
// @Summary Update an existing turno
// @Description Update an existing turno
// @Tags Turno
// @Accept json
// @Param id path int true "Turno ID"
// @Param turno body domain.Turno true "Turno object"
// @Param token header string true "TOKEN"
// @Success 200 {object} domain.Turno "Updated"
// @Router /turnos/update/{id} [put]
func (h *turnoHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}
		var p domain.Turno
		if err := c.BindJSON(&p); err != nil {
			panic(err)
		}
		updatedTurno, err := h.s.Update(id, p)
		if err != nil {
			panic(err)
		}
		web.Success(c, http.StatusOK, updatedTurno)
	}
}

// UpdateField godoc
// @Summary Update a specific field of a turno
// @Description Update a specific field of a turno by ID
// @Tags Turno
// @Accept json
// @Param id path int true "Turno ID"
// @Param token query string true "Authentication Token"
// @Param field_name body string true "Name of the field to update"
// @Param value body string true "New value for the field"
// @Param token header string true "TOKEN"
// @Success 200 {string} string "Field updated successfully"
// @Router /turnos/{id}/update-field [put]
func (h *turnoHandler) UpdateField() gin.HandlerFunc {
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
// @Summary Delete a turno
// @Description Delete a turno by ID
// @Tags Turno
// @Param id path int true "Turno ID"
// @Param token header string true "TOKEN"
// @Success 200 {string} string "Turno deleted successfully"
// @Router /turnos/{id} [delete]
func (h *turnoHandler) Delete() gin.HandlerFunc {
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
		web.Success(c, http.StatusOK, "Turno deleted successfully")
	}
}
