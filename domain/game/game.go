package game

type GameMode struct {
	ID        string
	NbPlayers int
	NbRounds  int
}

type Game struct {
	ID          string
	Name        string
	Description string
	Status      bool
	BgImg       string
	Modes       []GameMode
}
