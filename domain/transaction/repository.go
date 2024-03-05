package transaction

type Repository interface {
	GetAccountTransactions(user_id string) []Transaction
}
