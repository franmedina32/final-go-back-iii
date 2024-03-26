package handler

import (
	"final-go-back-III/internal/domain"
	"final-go-back-III/internal/turnos"
	"final-go-back-III/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type turnoHandler struct {
	s turnos.Service
}

func NewTurnoHandler(s turnos.Service) *turnoHandler {
	return &turnoHandler{
		s: s,
	}
}

func (h *turnoHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		turnos, err := h.s.GetAll()
		if err != nil {
			panic(err)
		}
		web.Success(c, http.StatusOK, turnos)
	}
}

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

func (h *turnoHandler) CreateByUserDoc() gin.HandlerFunc {
	return func(c *gin.Context) {
		var t turnos.CreateTurnoData
		if err := c.BindJSON(&t); err != nil {
			panic(err)
		}
		createdTurno, err := h.s.CreateByUserDoc(t)
		if err != nil {
			panic(err)
		}
		web.Success(c, http.StatusCreated, createdTurno)
	}
}

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

func (h *turnoHandler) createByUserDoc(dto turnos.CreateTurnoData) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p domain.Turno
		if err := c.BindJSON(&p); err != nil {
			panic(err)
		}
		createdTurno, err := h.s.CreateByUserDoc(dto)
		if err != nil {
			panic(err)
		}
		web.Success(c, http.StatusCreated, createdTurno)
	}
}
