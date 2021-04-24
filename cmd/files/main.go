package main

import (
	"github.com/mijgona/wallet/pkg/wallet"
	//"github.com/mijgona/wallet/pkg/types"
	"log"
)


func main() {
	var s wallet.Service
	_, err:=s.RegisterAccount("+992925996655")
	if err != nil{
		 log.Println(err)
	}
	_, err =s.RegisterAccount("+992900000000")
	if err != nil{
		 log.Println(err)
	}

	err=s.Deposit(1,10_000_00)
	if err!=nil{
		log.Printf("can`t deposit account, error= %v", err)
	}

	_, err = s.Pay(1, 200, "auto")


	err = s.Export("../../data")
	if err != nil {
		return
	}
	
}