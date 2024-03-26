package turnos

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

type Repository interface {
	GetAll() ([]domain.Turno, error)
	GetByID(id int) (domain.Turno, error)
	Create(t domain.Turno) (domain.Turno, error)
	Update(id int, p domain.Turno) (domain.Turno, error)
	UpdateField(id int, fieldName string, value interface{}) error
	Delete(id int) error
}

func (r *MySQLRepository) GetAll() ([]domain.Turno, error) {
	rows, err := r.db.Query("SELECT * FROM turnos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var turnos []domain.Turno
	for rows.Next() {
		var t domain.Turno
		err := rows.Scan(&t.Id, &t.PacienteId, &t.OdontologoId, &t.Fecha, &t.Descripcion)
		if err != nil {
			return nil, err
		}
		turnos = append(turnos, t)
	}
	return turnos, nil
}

func (r *MySQLRepository) GetByID(id int) (domain.Turno, error) {
	query := "SELECT * FROM turnos WHERE id = ?"
	row := r.db.QueryRow(query, id)
	var t domain.Turno

	err := row.Scan(&t.Id, &t.PacienteId, &t.OdontologoId, &t.Fecha, &t.Descripcion)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Turno{}, fmt.Errorf("Turno with ID %d not found", id)
		}
		return domain.Turno{}, err
	}
	return t, nil
}

func (r *MySQLRepository) Create(t domain.Turno) (domain.Turno, error) {
	query := "INSERT INTO turnos (id_paciente, id_odontologo, fecha, descripcion) VALUES (?, ?, ?, ?)"
	result, err := r.db.Exec(query, t.PacienteId, t.OdontologoId, t.Fecha, t.Descripcion)
	if err != nil {
		return domain.Turno{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return domain.Turno{}, err
	}
	t.Id = int(id)
	return t, nil
}

func (r *MySQLRepository) Update(id int, t domain.Turno) (domain.Turno, error) {
	query := "UPDATE turnos SET id_paciente = ?, id_odontologo = ?, fecha = ?, descripcion = ? WHERE id = ?"
	_, err := r.db.Exec(query, t.PacienteId, t.OdontologoId, t.Fecha, t.Descripcion, id)
	if err != nil {
		return domain.Turno{}, err
	}
	return t, nil
}

func (r *MySQLRepository) UpdateField(id int, fieldName string, value interface{}) error {
	query := fmt.Sprintf("UPDATE turnos SET %s = ? WHERE id = ?", fieldName)
	_, err := r.db.Exec(query, value, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *MySQLRepository) Delete(id int) error {
	query := "DELETE FROM turnos WHERE id = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("Turno with ID %d not found", id)
	}
	return nil
}
