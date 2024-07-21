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

func startServer(timHandler *handler.TimHandler, pikHandler *handler.PikHandler, korisnikHandler *handler.KorisnikHandler,
	regrutHandler *handler.RegrutHandler, igracHandler *handler.IgracHandler, zaposleniHandler *handler.ZaposleniHandler,
	authenticationHandler *handler.AuthenticationHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/tim", timHandler.GetAll).Methods("GET")
	router.HandleFunc("/tim/{id}", timHandler.GetByID).Methods("GET")

	router.HandleFunc("/pik", pikHandler.GetAll).Methods("GET")
	router.HandleFunc("/pik/{id}", pikHandler.GetByID).Methods("GET")
	router.HandleFunc("/pik/{teamId}", pikHandler.GetAllByTeamID).Methods("GET")
	router.HandleFunc("/pik/{year}", pikHandler.GetAllByYear).Methods("GET")

	router.HandleFunc("/korisnik", korisnikHandler.GetAll).Methods("GET")
	router.HandleFunc("/korisnik/{id}", korisnikHandler.GetByID).Methods("GET")

	router.HandleFunc("/regrut", regrutHandler.GetAll).Methods("GET")
	router.HandleFunc("/regrut/{id}", regrutHandler.GetByID).Methods("GET")
	router.HandleFunc("/regrut", regrutHandler.Create).Methods("POST")
	router.HandleFunc("/regrut", regrutHandler.Update).Methods("PUT")

	router.HandleFunc("/igrac", igracHandler.GetAll).Methods("GET")
	router.HandleFunc("/igrac/{id}", igracHandler.GetByID).Methods("GET")
	router.HandleFunc("/igrac/{teamId}", igracHandler.GetAllByTeamID).Methods("GET")

	router.HandleFunc("/zaposleni", zaposleniHandler.GetAll).Methods("GET")
	router.HandleFunc("/zaposleni/{id}", zaposleniHandler.GetByID).Methods("GET")

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

	timRepository := impl.NewTimRepository(&db)
	timService := service.NewTimService(timRepository)
	timHandler := handler.NewTimHandler(timService)

	pikRepository := impl.NewPikRepository(&db)
	pikService := service.NewPikService(pikRepository)
	pikHandler := handler.NewPikHandler(pikService)

	korisnikRepository := impl.NewKorisnikRepository(&db)
	korisnikService := service.NewKorisnikService(korisnikRepository)
	korisnikHandler := handler.NewKorisnikHandler(korisnikService)

	regrutRepository := impl.NewRegrutRepository(&db)
	regrutService := service.NewRegrutService(regrutRepository)
	regrutHandler := handler.NewRegrutHandler(regrutService)

	igracRepository := impl.NewIgracRepository(&db)
	igracService := service.NewIgracService(igracRepository)
	igracHandler := handler.NewIgracHandler(igracService)

	zaposleniRepository := impl.NewZaposleniRepository(&db)
	zaposleniService := service.NewZaposleniService(zaposleniRepository)
	zaposleniHandler := handler.NewZaposleniHandler(zaposleniService)

	authenticationHandler := handler.NewAuthenticationHandler(korisnikService)

	startServer(timHandler, pikHandler, korisnikHandler, regrutHandler, igracHandler, zaposleniHandler, authenticationHandler)
}
