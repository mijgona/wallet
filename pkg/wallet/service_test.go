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

func TestService_Reject_success(t *testing.T) {
	svc:=&Service{}
	account, err := svc.RegisterAccount("+992901900999")
	if err!=nil{
		t.Error(err)
		return
	}
	err=svc.Deposit(account.ID,100)
	if err!=nil{
		t.Error(err)
		return
	}
	payment, err:=svc.Pay(account.ID,20,"auto")
	if err!=nil{
		t.Error(err)
		return
	}
	
	err=svc.Reject(payment.ID)
	if err!=nil{
		t.Error(err)
		return
	}
}

func TestService_Reject_Fail(t *testing.T) {
	svc:=&Service{}
	_,err := svc.FindPaymentByID("payment.ID")
	if err==nil{
		t.Error("Должна быть ошибка платёж не найден")
		return
	}
	err=svc.Reject("payment.ID")
	if err==nil{
		t.Error("Должна быть ошибка платёж не найден")
		return
	}
}