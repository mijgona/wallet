package wallet

import (
	"github.com/google/uuid"
	"fmt"
	"testing"
	"reflect"
	
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

func TestService_FindPaymenByID_success(t *testing.T) {
	svc := &Service{}
	account, err := svc.RegisterAccount("+992907307999")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.Deposit(account.ID, 10000)
	if err != nil {
		fmt.Println(err)
		return
	}

	payment, err := svc.Pay(account.ID, 100, "food")
	if err != nil {
		fmt.Println(err)
		return
	}

	//try to find payment
	got, err := svc.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("Can't find payment By ID, Error: %v", err)
		return
	}

	
	if !reflect.DeepEqual(payment, got){
		t.Errorf("wrong payment returned, Error: %v", err)
		return
	}
}

func TestService_FindPaymentByID_Fail(t *testing.T) {
	svc := &Service{}
	account, err := svc.RegisterAccount("+992907307999")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.Deposit(account.ID, 10000)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = svc.Pay(account.ID, 100, "food")
	if err != nil {
		fmt.Println(err)
		return
	}
	//try to find payment
	_, err = svc.FindPaymentByID(uuid.New().String())
	if err == nil {
		t.Errorf("Can't find payment By ID, Error: %v", err)
		return
	}

	if err!= ErrPaymentNotFound{
		t.Errorf("must return error, returned: %v", err)
	}
}
func TestService_Reject_success(t *testing.T) {
	svc := &Service{}
	account, err := svc.RegisterAccount("+992907307999")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.Deposit(account.ID, 10000)
	if err != nil {
		fmt.Println(err)
		return
	}

	payment, err := svc.Pay(account.ID, 100, "food")
	if err != nil {
		fmt.Println(err)
		return
	}

	payment, err = svc.FindPaymentByID(payment.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.Reject(payment.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestService_Reject_Fail(t *testing.T) {
	svc := &Service{}
	account, err := svc.RegisterAccount("+992907306999")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.Deposit(account.ID, 10000)
	if err != nil {
		fmt.Println(err)
		return
	}

	payment, err := svc.Pay(account.ID, 100, "food")
	if err != nil {
		fmt.Println(err)
		return
	}

	payment, err = svc.FindPaymentByID(payment.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	account, err = svc.FindAccountByID(account.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.Reject(uuid.New().String())
	if err == nil {
		fmt.Println("Must be error")
		return
	}

}
