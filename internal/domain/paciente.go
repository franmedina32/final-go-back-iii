package domain

import (
	//"database/sql/driver"
	"encoding/json"
)

type Paciente struct {
	ID        int    `json:"id" example:"1"`
	Nombre    string `json:"nombre" example:"John"`
	Apellido  string `json:"apellido" example:"Doe"`
	Domicilio string `json:"domicilio" example:"123 Main St"`
	Dni       string `json:"dni" example:"12345678"`
	FechaAlta string `json:"fecha_alta" example:"2006-01-02 15:04:05"`
}

type PacienteTurno struct {
	ID        int    `json:"-"`
	Nombre    string `json:"nombre" example:"John"`
	Apellido  string `json:"apellido" example:"Doe"`
	Domicilio string `json:"domicilio" example:"123 Main St"`
	Dni       string `json:"dni" example:"12345678"`
	FechaAlta string `json:"fecha_alta" example:"2006-01-02 15:04:05"`
}

func (p *Paciente) UnmarshalJSON(data []byte) error {
	var aux struct {
		ID        int    `json:"id"`
		Nombre    string `json:"nombre"`
		Apellido  string `json:"apellido"`
		Domicilio string `json:"domicilio"`
		Dni       string `json:"dni"`
		// Change the type of FechaAlta to string
		FechaAlta string `json:"fecha_alta"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	p.ID = aux.ID
	p.Nombre = aux.Nombre
	p.Apellido = aux.Apellido
	p.Domicilio = aux.Domicilio
	p.Dni = aux.Dni
	p.FechaAlta = aux.FechaAlta

	return nil
}
