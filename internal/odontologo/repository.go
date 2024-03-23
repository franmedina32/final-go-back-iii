package odontologo

import (
	"database/sql"
	"final-go-back-III/internal/domain"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db}
}

type Repository interface {
	GetAll() ([]domain.Odontologo, error)
	//GetByID(id int) (domain.Odontologo, error)
	//SearchPriceGt(price float64) []domain.Odontologo
	//ConsumerPrice(list []int) ([]domain.Odontologo, float64, error)
	//Create(p domain.Odontologo) (domain.Odontologo, error)
	//Update(id int, p domain.Odontologo) (domain.Odontologo, error)
	//Delete(id int) error
}

func (r *MySQLRepository) GetAll() ([]domain.Odontologo, error) {
	rows, err := r.db.Query("SELECT id, nombre, apellido, matricula FROM odontologos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var odontologos []domain.Odontologo
	for rows.Next() {
		var o domain.Odontologo
		err := rows.Scan(&o.ID, &o.Nombre, &o.Apellido, &o.Matricula)
		if err != nil {
			return nil, err
		}
		odontologos = append(odontologos, o)
	}
	return odontologos, nil
}
