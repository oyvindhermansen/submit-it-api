package forms

import "go.mongodb.org/mongo-driver/bson/primitive"

type Service interface {
	Create(Form) (Form, error)
	List(ListOptions) ([]Form, error)
	FindById(primitive.ObjectID) (*Form, error)
}

type serviceImpl struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &serviceImpl{repository: repository}
}

func (s *serviceImpl) Create(form Form) (Form, error) {
	createdForm, err := s.repository.Create(form)

	if err != nil {
		return Form{}, err
	}

	return createdForm, nil
}

func (s *serviceImpl) FindById(id primitive.ObjectID) (*Form, error) {
	form, err := s.repository.FindById(id)

	if err != nil {
		return nil, err
	}

	return form, nil
}

func (s *serviceImpl) List(options ListOptions) ([]Form, error) {
	forms, err := s.repository.List(options)

	if err != nil {
		return nil, err
	}

	return forms, nil
}
