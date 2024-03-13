package wallet

import ()

type WalletService struct {
}

var ()

func NewWalletService() *WalletService {
	return &WalletService{}
}

func (s *WalletService) GetMyTransactions(username string) []any {
	return []any{}
}
