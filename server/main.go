package main

import (
	"basketball-league-server/handler"
	"basketball-league-server/repository/impl"
	"basketball-league-server/service"
	"database/sql"
	_ "github.com/godror/godror"
	"github.com/gorilla/handlers"
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

func startServer(teamHandler *handler.TeamHandler, pickHandler *handler.PickHandler, userHandler *handler.UserHandler,
	recruitHandler *handler.RecruitHandler, playerHandler *handler.PlayerHandler, employeeHandler *handler.EmployeeHandler,
	authenticationHandler *handler.AuthenticationHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/tim", teamHandler.GetAll).Methods("GET")
	router.HandleFunc("/tim/{id}", teamHandler.GetByID).Methods("GET")

	router.HandleFunc("/Pick", pickHandler.GetAll).Methods("GET")
	router.HandleFunc("/Pick/{id}", pickHandler.GetByID).Methods("GET")
	router.HandleFunc("/Pick/{teamId}", pickHandler.GetAllByTeamID).Methods("GET")
	router.HandleFunc("/Pick/{year}", pickHandler.GetAllByYear).Methods("GET")

	router.HandleFunc("/korisnik", userHandler.GetAll).Methods("GET")
	router.HandleFunc("/korisnik/{id}", userHandler.GetByID).Methods("GET")

	router.HandleFunc("/Recruit", recruitHandler.GetAll).Methods("GET")
	router.HandleFunc("/Recruit/{id}", recruitHandler.GetByID).Methods("GET")
	router.HandleFunc("/Recruit", recruitHandler.Create).Methods("POST")
	router.HandleFunc("/Recruit", recruitHandler.Update).Methods("PUT")

	router.HandleFunc("/igrac", playerHandler.GetAll).Methods("GET")
	router.HandleFunc("/igrac/{id}", playerHandler.GetByID).Methods("GET")
	router.HandleFunc("/igrac/{teamId}", playerHandler.GetAllByTeamID).Methods("GET")

	router.HandleFunc("/zaposleni", employeeHandler.GetAll).Methods("GET")
	router.HandleFunc("/zaposleni/{id}", employeeHandler.GetByID).Methods("GET")

	router.HandleFunc("/login", authenticationHandler.LogIn).Methods("POST")

	corsAllowedOrigins := handlers.AllowedOrigins([]string{"*"})
	corsAllowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	corsAllowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	log.Println("Server starting on :8081")
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(corsAllowedOrigins, corsAllowedMethods, corsAllowedHeaders)(router)))

}

func main() {
	db := initDB()
	defer db.Close()

	teamRepository := impl.NewTeamRepository(&db)
	teamService := service.NewTeamService(teamRepository)
	teamHandler := handler.NewTeamHandler(teamService)

	pickRepository := impl.NewPickRepository(&db)
	pickService := service.NewPickService(pickRepository)
	pickHandler := handler.NewPickHandler(pickService)

	userRepository := impl.NewUserRepository(&db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	recruitRepository := impl.NewRecruitRepository(&db)
	recruitService := service.NewRecruitService(recruitRepository)
	recruitHandler := handler.NewRecruitHandler(recruitService)

	playerRepository := impl.NewPlayerRepository(&db)
	playerService := service.NewPlayerService(playerRepository)
	playerHandler := handler.NewPlayerHandler(playerService)

	employeeRepository := impl.NewEmployeeRepository(&db)
	employeeService := service.NewEmployeeService(employeeRepository)
	employeeHandler := handler.NewEmployeeHandler(employeeService)

	authenticationHandler := handler.NewAuthenticationHandler(userService)

	startServer(teamHandler, pickHandler, userHandler, recruitHandler, playerHandler, employeeHandler, authenticationHandler)
}
