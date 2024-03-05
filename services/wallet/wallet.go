package wallet

import (
	"GemiApp/domain/transaction"
)

type WalletService struct {
	repo transaction.Repository
}

var ()

func NewWalletService(r transaction.Repository) *WalletService {
	return &WalletService{
		repo: r,
	}
}

func (s *WalletService) GetMyTransactions(username string) []transaction.Transaction {
	return s.repo.GetAccountTransactions(username)
}
