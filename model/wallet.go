package model

type WalletResponse struct {
	Address  string
	Link     string
	Balances []WalletBalanceResponse `json:"balances"`
}

type WalletBalanceResponse struct {
	Asset  string
	Amount string
}
