package main

import (
	"GemiApp/handler"
	"GemiApp/services/auth"
	"GemiApp/services/game"
	"GemiApp/services/wallet"
	"log"
	"net/http"
)

type Config struct {
	Addr string

	AuthUri   string
	WalletUri string
	GameUri   string
}

func run(cfg Config) error {
	// INIT SERVICES
	as := auth.NewAuthService(cfg.AuthUri)
	gs := game.NewGameService(cfg.GameUri)
	ws := wallet.NewWalletService()

	// INIT Handler
	handler := handler.NewRouter(handler.Params{
		AuthService:   as,
		GameService:   gs,
		WalletService: ws,
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
		Addr:      ":8000",
		AuthUri:   "http://api.traya7.com/v1/auth",
		WalletUri: "http://api.traya7.com/v1/wallet",

		GameUri: "http://traya7.com/gameservice",
		//GameUri: "http://192.168.229.117:8001/gameservice",
	}

	if err := run(cfg); err != nil {
		log.Fatal("Server Shutdown cause: ", err)
	}
}
