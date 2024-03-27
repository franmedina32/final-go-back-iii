package odontologo

import "final-go-back-III/internal/domain"

type Service interface {
	GetAll() ([]domain.Odontologo, error)
	GetByID(id int) (domain.Odontologo, error)
	Create(p domain.Odontologo) (domain.Odontologo, error)
	Update(id int, p domain.Odontologo) (domain.Odontologo, error)
	UpdateField(id int, fieldName string, value interface{}) error
	Delete(id int) error
}

type service struct {
	r OdontologoRepository
}

func NewService(r OdontologoRepository) Service {
	return &service{r}
}

func (s service) GetAll() ([]domain.Odontologo, error) {
	odontologos, err := s.r.GetAll()
	if err != nil {
		panic(err)
	}
	return odontologos, nil
}

func (s service) GetByID(id int) (domain.Odontologo, error) {
	odontologo, err := s.r.GetByID(id)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return odontologo, nil
}

func (s service) Create(o domain.Odontologo) (domain.Odontologo, error) {
	createdOdontologo, err := s.r.Create(o)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return createdOdontologo, nil
}

func (s service) Update(id int, o domain.Odontologo) (domain.Odontologo, error) {
	updatedOdontologo, err := s.r.Update(id, o)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return updatedOdontologo, nil
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
