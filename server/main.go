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
	authenticationHandler *handler.AuthenticationHandler, draftRightHandler *handler.DraftRightHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/team", teamHandler.GetAll).Methods("GET")
	router.HandleFunc("/team/{id}", teamHandler.GetByID).Methods("GET")

	router.HandleFunc("/pick", pickHandler.GetAll).Methods("GET")
	router.HandleFunc("/pick/{id}", pickHandler.GetByID).Methods("GET")
	router.HandleFunc("/pick/{teamId}", pickHandler.GetAllByTeamID).Methods("GET")
	router.HandleFunc("/pick/{year}", pickHandler.GetAllByYear).Methods("GET")

	router.HandleFunc("/user", userHandler.GetAll).Methods("GET")
	router.HandleFunc("/user/{id}", userHandler.GetByID).Methods("GET")

	router.HandleFunc("/recruit", recruitHandler.GetAll).Methods("GET")
	router.HandleFunc("/recruit/{id}", recruitHandler.GetByID).Methods("GET")
	router.HandleFunc("/recruit", recruitHandler.Create).Methods("POST")
	router.HandleFunc("/recruit", recruitHandler.Update).Methods("PUT")

	router.HandleFunc("/player", playerHandler.GetAll).Methods("GET")
	router.HandleFunc("/player/{id}", playerHandler.GetByID).Methods("GET")
	router.HandleFunc("/player/{teamId}", playerHandler.GetAllByTeamID).Methods("GET")

	router.HandleFunc("/employee", employeeHandler.GetAll).Methods("GET")
	router.HandleFunc("/employee/{id}", employeeHandler.GetByID).Methods("GET")

	router.HandleFunc("/login", authenticationHandler.LogIn).Methods("POST")

	router.HandleFunc("/draftRight", teamHandler.GetAll).Methods("GET")
	router.HandleFunc("/draftRight/{id}", teamHandler.GetByID).Methods("GET")

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

	draftRightRepository := impl.NewDraftRightRepository(&db)
	draftRightService := service.NewDraftRightService(draftRightRepository)
	draftRightHandler := handler.NewDraftRightHandler(draftRightService)

	startServer(teamHandler, pickHandler, userHandler, recruitHandler, playerHandler, employeeHandler,
		authenticationHandler, draftRightHandler)
}
