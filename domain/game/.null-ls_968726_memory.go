package game

type MemoRepo struct {
	Games []Game
}

func NewMemo() Repository {
	return &MemoRepo{
		Games: []Game{
			{
				ID:          "ludo_classic",
				Name:        "Ludo Classic",
				BgImg:       "/assets/images/ludo.jpg",
				Description: "Ludo is a strategy board game for two to four players, in which the players race their four tokens from start to finish according to the rolls of a single die. ",
				Status:      true,
				Modes: []GameMode{
					{ID: "0lcx10", NbPlayers: 2, NbRounds: 1},
					{ID: "0lc01x", NbPlayers: 4, NbRounds: 1},
					{ID: "0lc10x", NbPlayers: 8, NbRounds: 2},
					{ID: "0lc1xx", NbPlayers: 16, NbRounds: 2},
				},
			},
      {
        ID: "ludo_super",
        Name: "Ludo Super",
        BgImg: "/assets/images/ludo_super.jpg",
        Description: "comming soon",
        Status: false,
        Modes: []GameMode{},
      }
		},
	}
}

func (db *MemoRepo) GetGames() []Game {
	return db.Games
}
