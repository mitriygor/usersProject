package app

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/mitriygor/usersProject/domain"
	"github.com/mitriygor/usersProject/service"
	"github.com/mitriygor/usersProjectLib/logger"
	"log"
	"net/http"
	"os"
)

func sanityCheck() {
	envProps := []string{
		"SERVER_HOST",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_HOST",
		"DB_PORT",
		"DB_NAME",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logger.Fatal(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}

func Start() {
	sanityCheck()
	router := mux.NewRouter()

	dbClient := getDbClient()
	userRepositoryDb := domain.NewUserRepositoryDb(dbClient)
	ch := UserHandlers{service.NewUserService(userRepositoryDb)}

	router.
		HandleFunc("/users", ch.getAllUsers).
		Methods(http.MethodGet).
		Name("GetAllUsers")

	am := AuthMiddleware{domain.NewAuthRepository()}
	router.Use(am.authorizationHandler())
	address := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")
	logger.Info(fmt.Sprintf("Starting server on %s:%s", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func getDbClient() *sql.DB {
	client, err := sql.Open("postgres", "user=postgres password=postgres dbname=dbusers sslmode=disable")
	if err != nil {
		panic(err)
	}

	return client
}
