package builder

type BankBuilder struct {
	bank Bank
}

func NewBankBuilder(bankAccount Bank) BankBuilder {
	return BankBuilder{bank: bankAccount}
}

func (b *BankBuilder) SetAccountNumber(accountNo string) Bank  {
	b.bank.AccountNumber = accountNo
	return b.bank
}

func (b *BankBuilder) SetAccountName(name string) Bank  {
	b.bank.AccountName = name
	return b.bank
}

func (b *BankBuilder) SetBranch(branch string) Bank  {
	b.bank.Branch = branch
	return b.bank
}

func (b *BankBuilder) SetBalance(balance float64) Bank  {
	b.bank.Balance = balance
	return b.bank
}

func (b *BankBuilder) Build() Bank  {
	dataBank := Bank{
		AccountNumber: b.bank.AccountNumber,
		AccountName:   b.bank.AccountName,
		Branch:        b.bank.Branch,
		Balance:       b.bank.Balance,
	}

	return dataBank
}
