package main

import (
	"basketball-league-server/handler"
	"basketball-league-server/repository/impl"
	"basketball-league-server/service"
	"database/sql"
	_ "github.com/godror/godror"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func initDB() sql.DB {
	dsn := `user="sys" password="sys123" connectString="localhost:1521/xe" sysdba=1`
	db, err := sql.Open("godror", dsn)
	if err != nil {
		log.Fatalf("Failed to open a connection: %v", err)
	}
	defer db.Close()
	return *db
}

func startServer(timHandler *handler.TimHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/tim", timHandler.GetAll).Methods("GET")
	router.HandleFunc("/tim/{id}", timHandler.GetByID).Methods("GET")

	log.Println("Server starting on :8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	db := initDB()

	timRepository := impl.NewTimRepository(&db)
	timService := service.NewTimService(timRepository)
	timHandler := handler.NewTimHandler(timService)

	startServer(timHandler)
}