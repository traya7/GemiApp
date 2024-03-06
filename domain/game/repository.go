package game

type Repository interface {
	GetGames() []Game
}
