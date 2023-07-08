package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorilla-mux-postgres-sqlx-sample/service"
	"net/http"
)

type StudentHandlers struct {
	studentService service.StudentService
}

func (s *StudentHandlers) getAllStudents(w http.ResponseWriter, r *http.Request) {
	students, err := s.studentService.GetAllStudents()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, students)
	}
}

func (s *StudentHandlers) getStudentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["student_id"]

	student, err := s.studentService.GetStudentById(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, student)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		//panic(err)
	}
}
