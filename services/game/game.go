package game

import (
	"GemiApp/services"
	"GemiApp/types"
	"errors"
)

var (
	ErrNotFoundGame = errors.New("not found game")
)

type GameService struct {
	Uri   string
	Games []types.Game
}

func NewGameService(uri string) *GameService {
	return &GameService{
		Uri: uri,
		Games: []types.Game{
			{
				ID:          "ludo_classic",
				Name:        "Ludo Classic",
				BgImg:       "/assets/images/ludo.jpg",
				Description: "Ludo is a strategy board game for two to four players, in which the players race their four tokens from start to finish according to the rolls of a single die. ",
				Status:      true,
				Modes: []types.GameMode{
					{ID: "0lcx10", NbPlayers: 2, NbRounds: 1},
					{ID: "0lc01x", NbPlayers: 4, NbRounds: 1},
					{ID: "0lc10x", NbPlayers: 8, NbRounds: 2},
					{ID: "0lc1xx", NbPlayers: 16, NbRounds: 2},
				},
			},
			{
				ID:          "ludo_super",
				Name:        "Ludo Super",
				BgImg:       "/assets/images/super_ludo.webp",
				Description: "comming soon",
				Status:      false,
				Modes:       []types.GameMode{},
			},
		},
	}
}

func (s *GameService) GetAllGames() []types.Game {
	return s.Games
}

func (s *GameService) GetGameDetails(game_id string) (*types.Game, error) {
	for _, game := range s.Games {
		if game.ID == game_id {
			return &game, nil
		}
	}
	return nil, ErrNotFoundGame
}

func (s *GameService) NewGameSession(user_id, game_id string) map[string]any {
	if user_id == "" || game_id == "" {
		return map[string]any{"status": false, "message": "internal error"}
	}

	to := s.Uri + "/api/newsession"
	body := map[string]any{
		"user_id": user_id,
		"game_id": game_id,
	}

	r, err := services.GameServiceRequest(to, body)
	if err != nil {
		return map[string]any{
			"status":  false,
			"message": err.Error(),
		}
	}

	r["status"] = true
	return r
}
func (s *GameService) GameUriFrom(id string) string {
	return s.Uri + "/game/ludo?id=" + id
}
