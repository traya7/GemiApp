package helpers

type Game struct {
	ID       string
	Name     string
	ImageURI string
	Status   bool
}

func StaticGames() []Game {
	return []Game{
		{ID: "ludo_classic", Name: "Classic Ludo", ImageURI: "/assets/images/ludo.jpg", Status: true},
		{ID: "#", Name: "Super Ludo", ImageURI: "/assets/images/ludo.jpg", Status: false},
		{ID: "#", Name: "Noufi", ImageURI: "/assets/images/ludo.jpg", Status: false},
		{ID: "#", Name: "Classic UNO", ImageURI: "/assets/images/ludo.jpg", Status: false},
	}
}
