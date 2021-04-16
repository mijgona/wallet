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

	err = s.ExportToFile("../../data/accounts.txt")
	if err != nil {
		return
	}	
	err = s.ImportFromFile("../../data/accounts.txt")
	if err != nil {
		return
	}
	
}