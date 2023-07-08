package main

import (
	"github.com/joho/godotenv"
	"gorilla-mux-postgres-sqlx-sample/app"
	"gorilla-mux-postgres-sqlx-sample/logger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file!!")
	}
	app.Start()
	//router := http.NewServeMux()
	//router.HandleFunc("/students", getAllStudents)
	//router.HandleFunc("/studentsXml", getAllStudentsXml)
	//
	//http.ListenAndServe(":8080", router)

}

//func getAllStudents(w http.ResponseWriter, r *http.Request) {
//
//	students := []Student{
//		{1, "Sushil", 30},
//		{2, "John", 30},
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(200)
//	if r.Header.Get("Content-Type") == "application/xml" {
//		xml.NewEncoder(w).Encode(students)
//	} else {
//		json.NewEncoder(w).Encode(students)
//	}
//
//}
//func getAllStudentsXml(w http.ResponseWriter, r *http.Request) {
//
//	students := []Student{
//		{1, "Sushil", 30},
//		{2, "John", 30},
//	}
//
//	w.Header().Set("Content-Type", "application/xml")
//	w.WriteHeader(200)
//	xml.NewEncoder(w).Encode(students)
//}
//
//type Student struct {
//	Id   int    `json:"id" xml:"id"`
//	Name string `json:"name" xml:"name"`
//	Age  int    `json:"age" xml:"age"`
//}
