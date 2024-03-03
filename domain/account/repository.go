package account

type Repository interface {
	GetAccountByUsername(string) (*Account, error)
	GetAccountByID(string) (*Account, error)
}
