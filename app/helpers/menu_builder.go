package helpers

type MenuItem struct {
	Name string
	Link string
}

func MenuBuilder() []MenuItem {
	return []MenuItem{
		{Name: "LOCKED", Link: "#"},
		{Name: "LOCKED", Link: "#"},
		{Name: "LOCKED", Link: "#"},
		{Name: "SEP", Link: "#"},
		{Name: "Change Password", Link: "#"},
		{Name: "Transactions", Link: "#"},
		{Name: "History", Link: "#"},
		{Name: "SEP", Link: "#"},
	}
}
