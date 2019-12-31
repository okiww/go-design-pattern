package builder

type Bank struct {
	AccountNumber string
	AccountName   string
	Branch        string
	Balance       float64
}

func NewBank(builder BankBuilder) Bank  {
	return Bank{
		AccountNumber: builder.bank.AccountNumber,
		AccountName:   builder.bank.AccountName,
		Branch:        builder.bank.Branch,
		Balance:       builder.bank.Balance,
	}
}
