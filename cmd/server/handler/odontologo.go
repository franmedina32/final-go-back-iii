package handler

import (
	"final-go-back-III/internal/odontologo"
	"final-go-back-III/pkg/web"
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

func (h *odontologoHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		odontologos, err := h.s.GetAll()
		if err != nil {
			panic(err)
		}
		web.Success(c, 200, odontologos)
	}
}
