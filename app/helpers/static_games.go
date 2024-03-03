package helpers

type Game struct {
	ID       string
	ImageURI string
	Players  int
	Wins     int
}

func StaticGames() []Game {
	return []Game{
		{ID: "001", ImageURI: "/assets/images/ludo.jpg", Players: 4, Wins: 30},
		{ID: "002", ImageURI: "/assets/images/ludo.jpg", Players: 8, Wins: 80},
		{ID: "003", ImageURI: "/assets/images/ludo.jpg", Players: 16, Wins: 160},
	}
}
