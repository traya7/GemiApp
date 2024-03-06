package game

import (
	"GemiApp/domain/game"
	"GemiApp/types"
	"errors"
)

var (
	ErrNotFoundGame = errors.New("not found game")
)

type GameService struct {
	repo game.Repository
}

func NewGameService(r game.Repository) *GameService {
	return &GameService{
		repo: r,
	}
}

func (s *GameService) GetAllGames() []types.Game {
	var r []types.Game = []types.Game{}
	games := s.repo.GetGames()
	for _, v := range games {
		item := types.Game{
			ID:     v.ID,
			Name:   v.Name,
			Status: v.Status,
			BgImg:  v.BgImg,
		}
		for _, m := range v.Modes {
			item.Modes = append(item.Modes, types.GameMode{
				ID:        m.ID,
				NbPlayers: m.NbPlayers,
				NbRounds:  m.NbRounds,
			})
		}
		r = append(r, item)
	}
	return r
}

func (s *GameService) GetGameDetails(game_id string) (*types.Game, error) {
	games := s.repo.GetGames()
	for _, game := range games {
		if game.ID == game_id {
			r := types.Game{
				ID:          game.ID,
				Name:        game.Name,
				Description: game.Description,
				Status:      game.Status,
				BgImg:       game.BgImg,
			}
			for _, m := range game.Modes {
				r.Modes = append(r.Modes, types.GameMode{
					ID:        m.ID,
					NbPlayers: m.NbPlayers,
					NbRounds:  m.NbRounds,
				})
			}
			return &r, nil
		}
	}
	return nil, ErrNotFoundGame
}
