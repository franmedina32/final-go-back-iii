package handler

import (
	"final-go-back-III/internal/domain"
	"final-go-back-III/internal/odontologo"
	"final-go-back-III/pkg/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type odontologoHandler struct {
	s odontologo.Service
}

func NewOdontologoHandler(s odontologo.Service) *odontologoHandler {
	return &odontologoHandler{
		s: s,
	}
}

// GetAll godoc
// @Summary Get all odontologos
// @Description Retrieve all odontologos
// @Tags Odontologo
// @Produce json
// @Success 200 {array} domain.Odontologo "OK"
// @Router /odontologos/all [get]
func (h *odontologoHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		odontologos, err := h.s.GetAll()
		if err != nil {
			panic(err)
		}
		web.Success(c, 200, odontologos)
	}
}

// GetByID godoc
// @Summary Get an odontologo by ID
// @Description Retrieve an odontologo by ID
// @Tags Odontologo
// @Produce json
// @Param id path int true "Odontologo ID"
// @Success 200 {object} domain.Odontologo "OK"
// @Router /odontologos/{id} [get]
func (h *odontologoHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}
		odontologo, err := h.s.GetByID(id)
		if err != nil {
			panic(err)
		}
		web.Success(c, http.StatusOK, odontologo)
	}
}

// Create godoc
// @Summary Create a new odontologo
// @Description Create a new odontologo
// @Tags Odontologo
// @Accept json
// @Produce json
// @Param odontologo body domain.Odontologo true "Odontologo object"
// @Param token header string true "TOKEN"
// @Success 201 {object} domain.Odontologo "Created"
// @Router /odontologos/create [post]
func (h *odontologoHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var p domain.Odontologo
		if err := c.BindJSON(&p); err != nil {
			panic(err)
		}

		createdOdontologo, err := h.s.Create(p)
		if err != nil {
			panic(err)
		}
		web.Success(c, http.StatusCreated, createdOdontologo)
	}
}

// Update godoc
// @Summary Update an existing odontologo
// @Description Update an existing odontologo
// @Tags Odontologo
// @Accept json
// @Produce json
// @Param id path int true "Odontologo ID"
// @Param odontologo body domain.Odontologo true "Odontologo object"
// @Param token header string true "TOKEN"
// @Success 200 {object} domain.Odontologo "Updated"
// @Router /odontologos/update/{id} [put]
func (h *odontologoHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}
		var p domain.Odontologo
		if err := c.BindJSON(&p); err != nil {
			panic(err)
		}
		updatedOdontologo, err := h.s.Update(id, p)
		if err != nil {
			panic(err)
		}
		web.Success(c, http.StatusOK, updatedOdontologo)
	}
}

// UpdateField godoc
// @Summary Update a specific field of an odontologo
// @Description Update a specific field of an odontologo
// @Tags Odontologo
// @Accept json
// @Produce json
// @Param id path int true "Odontologo ID"
// @Param body body domain.UpdateFieldRequest true "Update Field Request"
// @Param token header string true "TOKEN"
// @Success 200 {string} string "Field updated successfully"
// @Router /odontologos/update-field/{id} [patch]
func (h *odontologoHandler) UpdateField() gin.HandlerFunc {
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
// @Summary Delete an odontologo
// @Description Delete an odontologo by ID
// @Tags Odontologo
// @Param id path int true "Odontologo ID"
// @Param token header string true "TOKEN"
// @Success 200 {string} string "Odontologo deleted successfully"
// @Router /odontologos/delete/{id} [delete]
func (h *odontologoHandler) Delete() gin.HandlerFunc {
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
		web.Success(c, http.StatusOK, "Odontologo deleted successfully")
	}
}
