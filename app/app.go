package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gorilla-mux-postgres-sqlx-sample/logger"
	"gorilla-mux-postgres-sqlx-sample/repository"
	"gorilla-mux-postgres-sqlx-sample/service"
	"log"
	"net/http"
	"os"
	"time"
)

func Start() {
	router := mux.NewRouter()

	dbClient := getDbClient()
	studentRepository := repository.NewStudentRepositoryImpl(dbClient)
	studentService := service.NewStudentServiceImpl(studentRepository)
	sh := StudentHandlers{
		studentService,
	}

	router.HandleFunc("/students", sh.getAllStudents).Methods(http.MethodGet).Name("GetAllStudents")
	router.HandleFunc("/students/{student_id:[0-9]+}", sh.getStudentById).Methods(http.MethodGet).Name("GetStudentById")

	port := os.Getenv("PORT")

	logger.Info(fmt.Sprintf("Starting server on port %s ...", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbAddr, dbUser, dbPasswd, dbName, dbPort)

	client, err := sqlx.Open("postgres", dataSource)
	if err != nil {
		//panic(err)
	}

	client.SetConnMaxLifetime(time.Second * 30)
	client.SetMaxOpenConns(5)
	client.SetMaxOpenConns(10)

	return client
}
