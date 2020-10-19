package wallet

import (
	"github.com/google/uuid"
	"github.com/mijgona/wallet/pkg/types"
	"fmt"
	"testing"
	"reflect"
	
)
func TestService_FindAccountByID_Success(t *testing.T) {
	s := newTestService()
	accounts, _,_, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
		}
	//try to find payment
	account := accounts
	got, err := s.FindAccountByID(account.ID)
	if err != nil {
		t.Errorf("Can't find payment By ID, Error: %v", err)
		return
	}

	if !reflect.DeepEqual(account, got){
		t.Errorf("wrong payment returned, Error: %v", err)
		return
	}
}


func TestService_FindAccountByID_Fail(t *testing.T) {
	s := newTestService()
	_, _, _, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = s.FindAccountByID(2)
	if err == nil {
		fmt.Println("must be error")
		return
	}
}

func TestService_FindPaymenByID_success(t *testing.T) {
	s := newTestService()
	_, payments,_, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	//try to find payment
	payment := payments[0]
	got, err := s.FindPaymentByID(payment.ID)
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
	s := newTestService()
	_, _, _, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	//try to find payment
	_, err = s.FindPaymentByID(uuid.New().String())
	if err == nil {
		t.Errorf("Can't find payment By ID, Error: %v", err)
		return
	}

	if err!= ErrPaymentNotFound{
		t.Errorf("must return error, returned: %v", err)
	}
}
func TestService_Reject_success(t *testing.T) {
	s := newTestService()
	_, payments,_, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	//try to reject payment
	payment := payments[0]
	err = s.Reject(payment.ID)
	if err != nil {
		t.Errorf("Reject(), Error: %v", err)
		return
	}

	savedPayment, err := s.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("Reject() can't find payment by id, Error: %v", err)
		return
	}
	if savedPayment.Status!= types.StatusFail{
		t.Errorf("Reject() status didn't changed, Error: %v", err)
		return
	}
	savedAccount, err := s.FindAccountByID(payment.AccountID)
	if err != nil {
		t.Errorf("Reject() can't find account by id, Error: %v", err)
		return
	}
	if savedAccount.Balance!=defaultTestAccount.balance {
		t.Errorf("Reject() balance didn't changed, account = %v", savedAccount)
	}
}

func TestService_Reject_Fail(t *testing.T) {
	s := newTestService()
	_, _, _, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	//try to reject payment
	err = s.Reject(uuid.New().String())
	if err == nil {
		t.Errorf("Reject(), Error: %v", err)
		return
	}

	if err!= ErrPaymentNotFound{
		t.Errorf("must return error, returned: %v", err)
	}

}



func TestService_Repeat_success(t *testing.T) {
	s := newTestService()
	_, payments, _, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	//try to find payment
	payment := payments[0]	
	got, err := s.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("Can't find payment By ID, Error: %v", err)
		return
	}

	if !reflect.DeepEqual(payment, got){
		t.Errorf("wrong payment returned, Error: %v", err)
		return
	}

	_, err = s.Repeat(payment.ID)
	if err != nil {
		t.Errorf("Cant Repeat payment, Error:  %v", err)
		return
	}

}




func TestService_Repeat_Fail(t *testing.T) {
	s := newTestService()
	_, _, _, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	//try to find payment
	_, err = s.FindPaymentByID(uuid.New().String())
	if err == nil {
		t.Errorf("Can't find payment By ID, Error: %v", err)
		return
	}

	if err!= ErrPaymentNotFound{
		t.Errorf("must return error, returned: %v", err)
	}
}

func TestService_PayFromFavorite_success(t *testing.T) {
	s := newTestService()
	_, _, favorites, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	//try to Pay From Favorite
	_, err = s.PayFromFavorite(favorites[0].ID)
	if err != nil {
		t.Errorf("PayFromFavorite(), Error: %v", err)
		return
	}

}


func TestService_PayFromFavorite_fail(t *testing.T) {
	s := newTestService()
	_, _, _, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = s.PayFromFavorite(uuid.New().String())
	if err == nil {
		t.Error("PayFromFavorite() must be Error")
		return
	}
}


type testService struct{
	*Service
}

func newTestService() *testService{
	return &testService{Service:&Service{}}
}

type testAccount struct{
	phone		types.Phone
	balance		types.Money
	payments 	[]struct{
		amount 		types.Money
		category	types.PaymentCategory
	}
}

var defaultTestAccount = testAccount{
	phone:		"+992900300900",
	balance:	10_000_00,
	payments:	[]struct {
		amount	types.Money
		category	types.PaymentCategory
	}{
		{amount:	1_000_00,	category: "auto"},
	},
}

func (s *testService) addAccount(data testAccount) (*types.Account, []*types.Payment, []*types.Favorite, error)  {
	//reg new account
	account, err:=s.RegisterAccount(data.phone)
	if err != nil{
		return nil, nil, nil, fmt.Errorf("can't register account, error: %v", err)
	}

	//deposit to account
	err = s.Deposit(account.ID, data.balance)
	if err != nil{
		return nil, nil,nil, fmt.Errorf("can't deposit to account, error: %v", err)
	}

	//Make payments
	payments := make([]*types.Payment, len(data.payments))
	for i, payment := range data.payments {
		payments[i],err=s.Pay(account.ID, payment.amount, payment.category)
		if err != nil {
			return nil, nil,nil, fmt.Errorf("can't make payment, error: %v", err)
		}
	}

	//make favorites
	var favorites []*types.Favorite
	var fav *types.Favorite
	for _, payment := range payments {	
	fav, err=s.FavoritePayment(payment.ID, "Favorite") 
	if err != nil {
		return nil, nil,nil, fmt.Errorf("can't make favorite, error: %v", err)
	}
	favorites=append(favorites,fav)
}
	return account, payments, favorites, nil
}