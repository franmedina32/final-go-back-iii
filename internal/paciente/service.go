package paciente

import "final-go-back-III/internal/domain"

type Service interface {
	GetAll() ([]domain.Paciente, error)
	GetByID(id int) (domain.Paciente, error)
	Create(p domain.Paciente) (domain.Paciente, error)
	Update(id int, p domain.Paciente) (domain.Paciente, error)
	UpdateField(id int, fieldName string, value interface{}) error
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s service) GetAll() ([]domain.Paciente, error) {
	pacientes, err := s.r.GetAll()
	if err != nil {
		panic(err)
	}
	return pacientes, nil
}

func (s service) GetByID(id int) (domain.Paciente, error) {
	paciente, err := s.r.GetByID(id)
	if err != nil {
		return domain.Paciente{}, err
	}
	return paciente, nil
}

func (s service) Create(p domain.Paciente) (domain.Paciente, error) {
	createdPaciente, err := s.r.Create(p)
	if err != nil {
		return domain.Paciente{}, err
	}
	return createdPaciente, nil
}

func (s service) Update(id int, p domain.Paciente) (domain.Paciente, error) {
	updatedPaciente, err := s.r.Update(id, p)
	if err != nil {
		return domain.Paciente{}, err
	}
	return updatedPaciente, nil
}

func (s service) UpdateField(id int, fieldName string, value interface{}) error {
	err := s.r.UpdateField(id, fieldName, value)
	if err != nil {
		return err
	}
	return nil
}

func (s service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
