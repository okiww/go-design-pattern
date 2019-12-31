package main

/*
	- Implement Builder Pattern in Go
	- Author : Okky Muhamad Budiman
	- Email : budimanokky93@gmail.com
*/

import (
	"fmt"
	"okky/go-design-pattern/Builder/builder"
)

/*
	- Assume we have an bank account for every users
*/

func main()  {
	user1 := builder.NewBankBuilder(builder.Bank{})
	user1.SetAccountNumber("asdasd")
	user1.SetAccountName("asdasdasd")
	user1.Build()


	user2 := builder.NewBankBuilder(builder.Bank{})
	user2.SetAccountNumber("asdasd")
	user2.SetBalance(750)
	user2.Build()

	fmt.Println(user1)
	fmt.Println(user2)
}
