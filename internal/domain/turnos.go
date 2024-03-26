package domain

type Turno struct {
	Id           int
	PacienteId   int
	OdontologoId int
	Fecha        CustomTime
	Descripcion  string
}
