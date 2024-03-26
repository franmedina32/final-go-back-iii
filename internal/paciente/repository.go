package paciente

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

type PacienteRepository interface {
	GetAll() ([]domain.Paciente, error)
	GetByID(id int) (domain.Paciente, error)
	Create(p domain.Paciente) (domain.Paciente, error)
	Update(id int, p domain.Paciente) (domain.Paciente, error)
	UpdateField(id int, fieldName string, value interface{}) error
	Delete(id int) error
	GetByDoc(doc string) (domain.Paciente, error)
}

func (r *MySQLRepository) GetAll() ([]domain.Paciente, error) {
	rows, err := r.db.Query("SELECT id, nombre, apellido, domicilio, dni, fecha_alta FROM pacientes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pacientes []domain.Paciente
	for rows.Next() {
		var o domain.Paciente
		var fechaAltaCustom domain.CustomTime
		err := rows.Scan(&o.ID, &o.Nombre, &o.Apellido, &o.Domicilio, &o.Dni, &fechaAltaCustom)
		if err != nil {
			return nil, err
		}
		o.FechaAlta = fechaAltaCustom
		pacientes = append(pacientes, o)
	}
	return pacientes, nil
}

func (r *MySQLRepository) GetByID(id int) (domain.Paciente, error) {
	query := "SELECT * FROM pacientes WHERE id = ?"
	row := r.db.QueryRow(query, id)
	var paciente domain.Paciente

	err := row.Scan(&paciente.ID, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio, &paciente.Dni, &paciente.FechaAlta)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Paciente{}, fmt.Errorf("Paciente with id %d not found", id)
		}
		return domain.Paciente{}, err
	}
	return paciente, nil
}

func (r *MySQLRepository) GetByDoc(doc string) (domain.Paciente, error) {
	query := "SELECT * FROM pacientes WHERE dni = ?"
	row := r.db.QueryRow(query, doc)
	var paciente domain.Paciente

	err := row.Scan(&paciente.ID, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio, &paciente.Dni, &paciente.FechaAlta)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Paciente{}, fmt.Errorf("Paciente with DOC %s not found", doc)
		}
		return domain.Paciente{}, err
	}
	return paciente, nil
}

func (r *MySQLRepository) Create(p domain.Paciente) (domain.Paciente, error) {
	fechaAltaTime := p.FechaAlta.Time
	query := "INSERT INTO pacientes (nombre, apellido, domicilio, dni, fecha_alta) VALUES (?, ?, ?, ?, ?)"
	result, err := r.db.Exec(query, p.Nombre, p.Apellido, p.Domicilio, p.Dni, fechaAltaTime)
	if err != nil {
		return domain.Paciente{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Paciente{}, err
	}
	p.ID = int(id)
	return p, nil
}

func (r *MySQLRepository) Update(id int, p domain.Paciente) (domain.Paciente, error) {
	query := "UPDATE pacientes SET nombre = ?, apellido = ?, domicilio = ?, dni = ?, fecha_alta = ? WHERE id = ?"
	_, err := r.db.Exec(query, p.Nombre, p.Apellido, p.Domicilio, p.Dni, p.FechaAlta, id)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}

func (r *MySQLRepository) UpdateField(id int, fieldName string, value interface{}) error {
	query := fmt.Sprintf("UPDATE pacientes SET %s = ? WHERE id = ?", fieldName)
	_, err := r.db.Exec(query, value, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *MySQLRepository) Delete(id int) error {
	query := "DELETE FROM pacientes WHERE id = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("Paciente with ID %d not found", id)
	}
	return nil
}
