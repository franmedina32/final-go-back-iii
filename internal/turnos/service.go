package turnos

import (
	"final-go-back-III/internal/domain"
	"final-go-back-III/internal/odontologo"
	"final-go-back-III/internal/paciente"
)

type Service interface {
	GetAll() ([]domain.Turno, error)
	GetByID(id int) (domain.Turno, error)
	Create(t domain.Turno) (domain.Turno, error)
	CreateByUserDoc(dto CreateTurnoData) (domain.Turno, error)
	Update(id int, p domain.Turno) (domain.Turno, error)
	UpdateField(id int, fieldName string, value interface{}) error
	Delete(id int) error
}

type service struct {
	turnoRepo      Repository
	pacienteRepo   paciente.PacienteRepository
	odontologoRepo odontologo.OdontologoRepository
}

type CreateTurnoData struct {
	PacienteDocNumber   string
	OdontologoMatricula string
	Descripcion         string
	Fecha               domain.CustomTime
}

func NewService(
	turnoRepo Repository,
	pacienteRepo paciente.PacienteRepository,
	odontologoRepo odontologo.OdontologoRepository,
) Service {
	return &service{turnoRepo, pacienteRepo, odontologoRepo}
}

func (s service) CreateByUserDoc(dto CreateTurnoData) (domain.Turno, error) {
	o, err := s.odontologoRepo.GetByMat(dto.OdontologoMatricula)
	p, err2 := s.pacienteRepo.GetByDoc(dto.PacienteDocNumber)
	if err != nil {
		return domain.Turno{}, err
	}
	if err2 != nil {
		return domain.Turno{}, err2
	}
	t := domain.Turno{
		OdontologoId: o.ID,
		PacienteId:   p.ID,
		Fecha:        dto.Fecha,
		Descripcion:  dto.Descripcion,
	}
	return s.turnoRepo.Create(t)
}

func (s service) GetAll() ([]domain.Turno, error) {
	turnos, err := s.turnoRepo.GetAll()
	if err != nil {
		panic(err)
	}
	return turnos, nil
}

func (s service) GetByID(id int) (domain.Turno, error) {
	turnos, err := s.turnoRepo.GetByID(id)
	if err != nil {
		return domain.Turno{}, err
	}
	return turnos, nil
}

func (s service) Create(t domain.Turno) (domain.Turno, error) {
	createdTurno, err := s.turnoRepo.Create(t)
	if err != nil {
		return domain.Turno{}, err
	}
	return createdTurno, nil
}

func (s service) Update(id int, t domain.Turno) (domain.Turno, error) {
	updatedTurno, err := s.turnoRepo.Update(id, t)
	if err != nil {
		return domain.Turno{}, err
	}
	return updatedTurno, nil
}

func (s service) UpdateField(id int, fieldName string, value interface{}) error {
	err := s.turnoRepo.UpdateField(id, fieldName, value)
	if err != nil {
		return err
	}
	return nil
}

func (s service) Delete(id int) error {
	err := s.turnoRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
