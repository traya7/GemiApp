package main

import (
	"GemiApp/app"
	"GemiApp/domain"
	"GemiApp/domain/account"
	"GemiApp/services/auth"
	"log"
	"net/http"
)

type Config struct {
	Addr     string
	MongoURI string
	MongoDB  string
}

func newServer() {

}

func run(cfg Config) error {
	// INIT DB CONNECTION
	mongodb, err := domain.NewMongoDB(cfg.MongoURI, cfg.MongoDB)
	if err != nil {
		return err
	}
	log.Println("- Mongo DB connected!")

	// INIT SERVICES
	as := auth.NewAuthService(account.NewMongoRepo(mongodb))

	// INIT Handler
	handler := app.NewRouter(app.Params{
		AuthService: as,
	})

	// START SERVER

	log.Println("- Http Server running!")
	if err := http.ListenAndServe(cfg.Addr, handler); err == nil {
		return err
	}
	return nil
}

func main() {
	cfg := Config{
		Addr:     ":8000",
		MongoURI: "mongodb://dbxadmin2:Aopj0R89Zp3J@203.161.44.242:27017/",
		MongoDB:  "gmetour",
	}

	if err := run(cfg); err != nil {
		log.Fatal("Server Shutdown cause: ", err)
	}
}
