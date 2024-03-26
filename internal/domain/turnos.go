package domain

type Turno struct {
	Id           int        `json:"id"`
	PacienteId   int        `json:"id_paciente"`
	OdontologoId int        `json:"id_odontologo"`
	Fecha        CustomTime `json:"fecha"`
	Descripcion  string     `json:"descripcion"`
}
