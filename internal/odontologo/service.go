package odontologo

import "final-go-back-III/internal/domain"

type Service interface {
	GetAll() ([]domain.Odontologo, error)
	//GetByID(id int) (domain.Odontologo, error)
	//SearchPriceGt(price float64) []domain.Odontologo
	//ConsumerPrice(list []int) ([]domain.Odontologo, float64, error)
	//Create(p domain.Odontologo) (domain.Odontologo, error)
	//Update(id int, p domain.Odontologo) (domain.Odontologo, error)
	//Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s service) GetAll() ([]domain.Odontologo, error) {
	odontologos, err := s.r.GetAll()
	if err != nil {
		panic(err)
	}
	return odontologos, nil
}

//
//func (s service) GetByID(id int) (domain.Odontologo, error) {
//	TODO implement me
//panic("implement me")
//}
//
//func (s service) SearchPriceGt(price float64) []domain.Odontologo {
//	TODO implement me
//panic("implement me")
//}
//
//func (s service) ConsumerPrice(list []int) ([]domain.Odontologo, float64, error) {
//	TODO implement me
//panic("implement me")
//}
//
//func (s service) Create(p domain.Odontologo) (domain.Odontologo, error) {
//	TODO implement me
//panic("implement me")
//}
//
//func (s service) Update(id int, p domain.Odontologo) (domain.Odontologo, error) {
//	TODO implement me
//panic("implement me")
//}
//
//func (s service) Delete(id int) error {
//	TODO implement me
//panic("implement me")
//}
//
