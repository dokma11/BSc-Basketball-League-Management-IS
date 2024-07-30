package main

import (
	"basketball-league-server/handler"
	"basketball-league-server/repository/impl"
	"basketball-league-server/service"
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func initDB() sql.DB {
	dsn := `user="C##Dokma11" password="dokma11" connectString="localhost:1521/xe" sysdba=0`
	db, err := sql.Open("godror", dsn)
	if err != nil {
		log.Fatalf("Failed to open a connection: %v", err)
	}
	defer db.Close()
	fmt.Println("Connected to DB")
	return *db
}

func startServer(teamHandler *handler.TeamHandler, pickHandler *handler.PickHandler, userHandler *handler.UserHandler,
	recruitHandler *handler.RecruitHandler, playerHandler *handler.PlayerHandler, employeeHandler *handler.EmployeeHandler,
	authenticationHandler *handler.AuthenticationHandler, draftRightHandler *handler.DraftRightHandler,
	contractHandler *handler.ContractHandler, draftHandler *handler.DraftHandler, tradeProposalHandler *handler.TradeProposalHandler,
	tradeHandler *handler.TradeHandler, trainingHandler *handler.TrainingHandler, trainingRequestHandler *handler.TrainingRequestHandler,
	interviewHandler *handler.InterviewHandler, interviewRequestHandler *handler.InterviewRequestHandler, tradeSubjectHandler *handler.TradeSubjectHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/team", teamHandler.GetAll).Methods("GET")
	router.HandleFunc("/team/{id}", teamHandler.GetByID).Methods("GET")
	router.HandleFunc("/team-user/{userId}", teamHandler.GetByUserID).Methods("GET")
	router.HandleFunc("/team-trade-subject/{tradeSubjectId}", teamHandler.GetPlayerTradeDestination).Methods("GET")

	router.HandleFunc("/pick", pickHandler.GetAll).Methods("GET")
	router.HandleFunc("/pick/{id}", pickHandler.GetByID).Methods("GET")
	router.HandleFunc("/pick/team/{teamId}", pickHandler.GetAllByTeamID).Methods("GET")
	router.HandleFunc("/pick/year/{year}", pickHandler.GetAllByYear).Methods("GET")

	router.HandleFunc("/user", userHandler.GetAll).Methods("GET")
	router.HandleFunc("/user/{id}", userHandler.GetByID).Methods("GET")
	router.HandleFunc("/user", userHandler.Update).Methods("PUT")

	router.HandleFunc("/recruit", recruitHandler.GetAll).Methods("GET")
	router.HandleFunc("/recruit/{id}", recruitHandler.GetByID).Methods("GET")
	router.HandleFunc("/recruit", recruitHandler.Create).Methods("POST")
	router.HandleFunc("/recruit", recruitHandler.Update).Methods("PUT")

	router.HandleFunc("/player", playerHandler.GetAll).Methods("GET")
	router.HandleFunc("/player/{id}", playerHandler.GetByID).Methods("GET")
	router.HandleFunc("/player/team/{teamId}", playerHandler.GetAllByTeamID).Methods("GET")

	router.HandleFunc("/employee", employeeHandler.GetAll).Methods("GET")
	router.HandleFunc("/employee/{id}", employeeHandler.GetByID).Methods("GET")
	router.HandleFunc("/employee/team/{teamId}", employeeHandler.GetByTeamID).Methods("GET")

	router.HandleFunc("/login", authenticationHandler.LogIn).Methods("POST")

	router.HandleFunc("/draftRight", draftRightHandler.GetAll).Methods("GET")
	router.HandleFunc("/draftRight/{id}", draftRightHandler.GetByID).Methods("GET")
	router.HandleFunc("/draftRight/team/{teamId}", draftRightHandler.GetAllByTeamID).Methods("GET")

	router.HandleFunc("/contract", contractHandler.GetAll).Methods("GET")
	router.HandleFunc("/contract/{id}", contractHandler.GetByID).Methods("GET")

	router.HandleFunc("/draft", draftHandler.GetAll).Methods("GET")
	router.HandleFunc("/draft/{id}", draftHandler.GetByID).Methods("GET")

	router.HandleFunc("/tradeProposal", tradeProposalHandler.GetAll).Methods("GET")
	router.HandleFunc("/tradeProposal/{id}", tradeProposalHandler.GetByID).Methods("GET")
	router.HandleFunc("/tradeProposal-received/{managerId}", tradeProposalHandler.GetAllReceivedByManagerID).Methods("GET")
	router.HandleFunc("/tradeProposal-sent/{managerId}", tradeProposalHandler.GetAllSentByManagerID).Methods("GET")
	router.HandleFunc("/tradeProposal-latest", tradeProposalHandler.GetLatest).Methods("GET")
	router.HandleFunc("/tradeProposal", tradeProposalHandler.Create).Methods("POST")
	router.HandleFunc("/tradeProposal", tradeProposalHandler.Update).Methods("PUT")

	router.HandleFunc("/trade", tradeHandler.GetAll).Methods("GET")
	router.HandleFunc("/trade/{id}", tradeHandler.GetByID).Methods("GET")
	router.HandleFunc("/trade/{teamId}", tradeHandler.GetAllByTeamID).Methods("GET")
	router.HandleFunc("/trade", tradeHandler.Create).Methods("POST")

	router.HandleFunc("/training", trainingHandler.GetAll).Methods("GET")
	router.HandleFunc("/training/{id}", trainingHandler.GetByID).Methods("GET")
	router.HandleFunc("/training/{userId}", trainingHandler.GetAllByUserID).Methods("GET")
	router.HandleFunc("/training", trainingHandler.Create).Methods("POST")

	router.HandleFunc("/trainingRequest", trainingRequestHandler.GetAll).Methods("GET")
	router.HandleFunc("/trainingRequest/{id}", trainingRequestHandler.GetByID).Methods("GET")
	router.HandleFunc("/trainingRequest/sender/{userId}", trainingRequestHandler.GetAllBySenderID).Methods("GET")
	router.HandleFunc("/trainingRequest/receiver/{userId}", trainingRequestHandler.GetAllByReceiverID).Methods("GET")
	router.HandleFunc("/trainingRequest", trainingRequestHandler.Create).Methods("POST")
	router.HandleFunc("/trainingRequest", trainingRequestHandler.Update).Methods("PUT")

	router.HandleFunc("/interview", interviewHandler.GetAll).Methods("GET")
	router.HandleFunc("/interview/{id}", interviewHandler.GetByID).Methods("GET")
	router.HandleFunc("/interview/{userId}", interviewHandler.GetAllByUserID).Methods("GET")
	router.HandleFunc("/interview", interviewHandler.Create).Methods("POST")

	router.HandleFunc("/interviewRequest", interviewRequestHandler.GetAll).Methods("GET")
	router.HandleFunc("/interviewRequest/{id}", interviewRequestHandler.GetByID).Methods("GET")
	router.HandleFunc("/interviewRequest/sender/{userId}", interviewRequestHandler.GetAllBySenderID).Methods("GET")
	router.HandleFunc("/interviewRequest/receiver/{userId}", interviewRequestHandler.GetAllByReceiverID).Methods("GET")
	router.HandleFunc("/interviewRequest", interviewRequestHandler.Create).Methods("POST")
	router.HandleFunc("/interviewRequest", interviewRequestHandler.Update).Methods("PUT")

	router.HandleFunc("/tradeSubject", tradeSubjectHandler.GetAll).Methods("GET")
	router.HandleFunc("/tradeSubject/{id}", tradeSubjectHandler.GetByID).Methods("GET")
	router.HandleFunc("/tradeSubject-trade/{tradeId}", tradeSubjectHandler.GetAllByTradeID).Methods("GET")
	router.HandleFunc("/tradeSubject-details/{tradeId}", tradeSubjectHandler.GetDetailsForTradeProposal).Methods("GET")
	router.HandleFunc("/tradeSubject", tradeSubjectHandler.Create).Methods("POST")
	router.HandleFunc("/tradeSubject-commit-trade", tradeSubjectHandler.CommitTrade).Methods("POST")

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

	authenticationHandler := handler.NewAuthenticationHandler(userService, teamService)

	draftRightRepository := impl.NewDraftRightRepository(&db)
	draftRightService := service.NewDraftRightService(draftRightRepository)
	draftRightHandler := handler.NewDraftRightHandler(draftRightService)

	contractRepository := impl.NewContractRepository(&db)
	contractService := service.NewContractService(contractRepository)
	contractHandler := handler.NewContractHandler(contractService)

	draftRepository := impl.NewDraftRepository(&db)
	draftService := service.NewDraftService(draftRepository)
	draftHandler := handler.NewDraftHandler(draftService)

	tradeProposalRepository := impl.NewTradeProposalRepository(&db)
	tradeProposalService := service.NewTradeProposalService(tradeProposalRepository)
	tradeProposalHandler := handler.NewTradeProposalHandler(tradeProposalService, employeeService)

	tradeRepository := impl.NewTradeRepository(&db)
	tradeService := service.NewTradeService(tradeRepository)
	tradeHandler := handler.NewTradeHandler(tradeService)

	trainingRepository := impl.NewTrainingRepository(&db)
	trainingService := service.NewTrainingService(trainingRepository)
	trainingHandler := handler.NewTrainingHandler(trainingService)

	trainingRequestRepository := impl.NewTrainingRequestRepository(&db)
	trainingRequestService := service.NewTrainingRequestService(trainingRequestRepository)
	trainingRequestHandler := handler.NewTrainingRequestHandler(trainingRequestService)

	interviewRepository := impl.NewInterviewRepository(&db)
	interviewService := service.NewInterviewService(interviewRepository)
	interviewHandler := handler.NewInterviewHandler(interviewService)

	interviewRequestRepository := impl.NewInterviewRequestRepository(&db)
	interviewRequestService := service.NewInterviewRequestService(interviewRequestRepository)
	interviewRequestHandler := handler.NewInterviewRequestHandler(interviewRequestService)

	tradeSubjectRepository := impl.NewTradeSubjectRepository(&db)
	tradeSubjectService := service.NewTradeSubjectService(tradeSubjectRepository)
	tradeSubjectHandler := handler.NewTradeSubjectHandler(tradeSubjectService, tradeProposalService, teamService,
		pickService, draftRightService, employeeService, contractService, tradeService)

	startServer(teamHandler, pickHandler, userHandler, recruitHandler, playerHandler, employeeHandler,
		authenticationHandler, draftRightHandler, contractHandler, draftHandler, tradeProposalHandler, tradeHandler,
		trainingHandler, trainingRequestHandler, interviewHandler, interviewRequestHandler, tradeSubjectHandler)
}
