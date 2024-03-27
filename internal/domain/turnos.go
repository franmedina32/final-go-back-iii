package domain

import "encoding/json"

type Turno struct {
	Id           int    `json:"id" example:"1"`
	PacienteId   int    `json:"id_paciente" example:"1"`
	OdontologoId int    `json:"id_odontologo" example:"1"`
	Fecha        string `json:"fecha" example:"2006-01-02 15:04:05"`
	Descripcion  string `json:"descripcion" example:"string"`
}

func (t *Turno) UnmarshalJSON(data []byte) error {
	var aux struct {
		Id           int    `json:"id" example:"1"`
		PacienteId   int    `json:"id_paciente" example:"1"`
		OdontologoId int    `json:"id_odontologo" example:"1"`
		Fecha        string `json:"fecha" example:"2006-01-02 15:04:05"`
		Descripcion  string `json:"descripcion" example:"string"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	t.Id = aux.Id
	t.PacienteId = aux.PacienteId
	t.OdontologoId = aux.OdontologoId
	t.Fecha = aux.Fecha
	t.Descripcion = aux.Descripcion

	return nil
}
