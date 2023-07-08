package service

import (
	"gorilla-mux-postgres-sqlx-sample/dto"
	errs "gorilla-mux-postgres-sqlx-sample/error"
	"gorilla-mux-postgres-sqlx-sample/repository"
)

type StudentService interface {
	GetAllStudents() ([]dto.StudentResponse, *errs.AppError)
	GetStudentById(string) (*dto.StudentResponse, *errs.AppError)
}

type StudentServiceImpl struct {
	repository repository.StudentRepository
}

func NewStudentServiceImpl(repository repository.StudentRepositoryImpl) StudentServiceImpl {
	return StudentServiceImpl{repository}
}

func (s StudentServiceImpl) GetAllStudents() ([]dto.StudentResponse, *errs.AppError) {
	students, err := s.repository.FindAll()

	if err != nil {
		return nil, err
	}
	response := make([]dto.StudentResponse, 0)
	for _, student := range students {
		response = append(response, student.ToDto())
	}
	return response, nil
}

func (s StudentServiceImpl) GetStudentById(id string) (*dto.StudentResponse, *errs.AppError) {
	student, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	response := student.ToDto()

	return &response, nil
}
