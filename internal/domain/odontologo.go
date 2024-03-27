package domain

type Odontologo struct {
	ID        int    `json:"id" example:"1"`
	Nombre    string `json:"nombre" example:"John"`
	Apellido  string `json:"apellido" example:"Doe"`
	Matricula string `json:"matricula" example:"12345678"`
}

type OdontologoTurno struct {
	ID        int    `json:"-"`
	Nombre    string `json:"nombre" example:"John"`
	Apellido  string `json:"apellido" example:"Doe"`
	Matricula string `json:"matricula" example:"12345678"`
}
