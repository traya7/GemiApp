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
		{Name: "Change Password", Link: "/user/resetpwd"},
		{Name: "Game History", Link: "/user/history"},
		{Name: "Transactions", Link: "/user/transactions"},
		{Name: "SEP", Link: "#"},
	}
}
