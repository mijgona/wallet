package wallet

import (
	"testing")


func TestService_FindAccountByID_success(t *testing.T) {
	svc:=&Service{}
	account, err := svc.RegisterAccount("+992901900999")
	if err!=nil{
		t.Error(err)
		return
	}
	_, err = svc.FindAccountByID(account.ID)
	if err!=nil{
		t.Error(err)
		return
	}	
}
func TestService_FindAccountByID_fail(t *testing.T) {
	svc:=&Service{}
	
	_, err := svc.FindAccountByID(0)
	if err==nil{
			t.Errorf("Аккаунт не должен находиться")
		return
	}	
}