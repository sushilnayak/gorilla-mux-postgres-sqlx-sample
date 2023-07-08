package model

import "gorilla-mux-postgres-sqlx-sample/dto"

type Student struct {
	Id   int
	Name string
	Age  int
}

func (student Student) ToDto() dto.StudentResponse {
	return dto.StudentResponse{
		Id:   student.Id,
		Name: student.Name,
		Age:  student.Age,
	}
}

//type StudentResponse struct {
//	Id   int    `json:"id"`
//	Name string `json:"name"`
//	Age  int    `json:"age"`
//}
