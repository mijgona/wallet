package wallet

import (
	"fmt"
	"testing"
	
)




func TestService_FindAccountByID_Success(t *testing.T) {
	svc := &Service{}
	account, err := svc.RegisterAccount("+992907306999")
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(account)

	account, err = svc.FindAccountByID(account.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
}


func TestService_FindAccountByID_Fail(t *testing.T) {
	svc := &Service{}
	_, err := svc.RegisterAccount("+992907306999")
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(account)
	_, err = svc.FindAccountByID(2)
	if err == nil {
		fmt.Println("must be error")
		return
	}
}