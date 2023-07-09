package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	errs "gorilla-mux-postgres-sqlx-sample/error"
	"gorilla-mux-postgres-sqlx-sample/logger"
	"gorilla-mux-postgres-sqlx-sample/model"
)

type StudentRepository interface {
	FindAll() ([]model.Student, *errs.AppError)
	FindById(string) (*model.Student, *errs.AppError)
}

type StudentRepositoryImpl struct {
	client *sqlx.DB
}

func NewStudentRepositoryImpl(dbClient *sqlx.DB) StudentRepositoryImpl {
	return StudentRepositoryImpl{dbClient}
}

func (repository StudentRepositoryImpl) FindAll() ([]model.Student, *errs.AppError) {
	var err error
	var students = make([]model.Student, 0)
	err = repository.client.Select(&students, "select id, name, age from students")

	if err != nil {
		logger.Error("Error while querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return students, nil
}

func (repository StudentRepositoryImpl) FindById(id string) (*model.Student, *errs.AppError) {

	var student model.Student
	err := repository.client.Get(&student, "select id, name, age from students where id=$1", id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Student not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &student, nil
}
