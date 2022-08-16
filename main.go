package main

import (
	"fmt"
	"practice/First/banking"
)

func main() {
	accout := banking.Account{Owner: "pukja", Balance: 1000}
	fmt.Println("hello GO!!")
	fmt.Println(accout)
}
