package service

type Service struct {
	Library
}

func NewService(library Library) *Service {
	return &Service{Library: library}
}

type Library interface {
	GetAll()
}