package handler

import (
	"final-go-back-III/internal/domain"
	"final-go-back-III/internal/odontologo"
	"final-go-back-III/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type odontologoHandler struct {
	s odontologo.Service
}

func NewOdontologoHandler(s odontologo.Service) *odontologoHandler {
	return &odontologoHandler{
		s: s,
	}
}

func (h *odontologoHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		odontologos, err := h.s.GetAll()
		if err != nil {
			panic(err)
		}
		web.Success(c, 200, odontologos)
	}
}

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
