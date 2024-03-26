package odontologo

import (
	"database/sql"
	"final-go-back-III/internal/domain"
	"fmt"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db}
}

type OdontologoRepository interface {
	GetAll() ([]domain.Odontologo, error)
	GetByID(id int) (domain.Odontologo, error)
	Create(p domain.Odontologo) (domain.Odontologo, error)
	Update(id int, p domain.Odontologo) (domain.Odontologo, error)
	Delete(id int) error
	GetByMat(mat string) (domain.Odontologo, error)
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

func (r *MySQLRepository) GetByID(id int) (domain.Odontologo, error) {
	query := "SELECT id, nombre, apellido, matricula FROM odontologos WHERE id = ?"
	row := r.db.QueryRow(query, id)
	var o domain.Odontologo

	err := row.Scan(&o.ID, &o.Nombre, &o.Apellido, &o.Matricula)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Odontologo{}, fmt.Errorf("Odontologo with ID %d not found", id)
		}
		return domain.Odontologo{}, err
	}
	return o, nil
}

func (r *MySQLRepository) GetByMat(mat string) (domain.Odontologo, error) {
	query := "SELECT * FROM odontologos WHERE matricula = ?"
	row := r.db.QueryRow(query, mat)
	var o domain.Odontologo

	err := row.Scan(&o.ID, &o.Nombre, &o.Apellido, &o.Matricula)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Odontologo{}, fmt.Errorf("Odontologo with matricula %s not found", mat)
		}
		return domain.Odontologo{}, err
	}
	return o, nil
}

func (r *MySQLRepository) Create(o domain.Odontologo) (domain.Odontologo, error) {
	query := "INSERT INTO odontologos (nombre, apellido, matricula) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, o.Nombre, o.Apellido, o.Matricula)
	if err != nil {
		return domain.Odontologo{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Odontologo{}, err
	}
	o.ID = int(id)
	return o, nil
}

func (r *MySQLRepository) Update(id int, o domain.Odontologo) (domain.Odontologo, error) {
	query := "UPDATE odontologos SET nombre = ?, apellido = ?, matricula = ? WHERE id = ?"
	_, err := r.db.Exec(query, o.Nombre, o.Apellido, o.Matricula, id)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return o, nil
}

func (r *MySQLRepository) Delete(id int) error {
	query := "DELETE FROM odontologos WHERE id = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("Odontologo with ID %d not found", id)
	}
	return nil
}
