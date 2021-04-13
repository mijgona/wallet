package wallet

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/mijgona/wallet/pkg/types"
)

//Встраивание функцийй (embedding)
 type testService struct {
	 *Service
 }
 //Структура тестового аккаунт
 type testAccount struct {
	 phone 			types.Phone
	 balance		types.Money
	 payments [] struct {
		 amount		types.Money
		 category	types.PaymentCategory
	 }
 }

 //Данные тестового аккаунта
 var defaultTestAccount= testAccount{
 	phone:    "+992990999099",
 	balance:  10_000_00,
 	payments: []struct{
		 amount types.Money
		 category types.PaymentCategory
		 }{
			 {amount: 1_000_00, category: "auto" },
		 },
 }

 //Функция конструктор
 func newTestService () *testService {
	 return &testService{Service: &Service{}}
 }

 //Создаём аккаунт для тестирования
func (s *testService) addAccount(data testAccount) (*types.Account, []*types.Payment, error)  {
 //Регистрируем пользователя
 account, err := s.RegisterAccount(data.phone)
	if err !=nil {
		return nil, nil, fmt.Errorf("can`t register account, error=%v", err)
	}
//Пополняет его счет
err=s.Deposit(account.ID,data.balance)
	if err!=nil{
		return nil, nil, fmt.Errorf("can`t deposit account, error=%v", err)
	}

//выполняем платежи
//Можем сразу создать слайс нужной длины, поскольку знаем размер
payments := make([]*types.Payment, len(data.payments))
 for i, payment := range data.payments {
	 //Работаем через индекс
	 payments[i], err = s.Pay(account.ID, payment.amount, payment.category)
	 	if err!=nil{
		return nil, nil, fmt.Errorf("can`t make payment, error=%v", err)
	}
 }
return account, payments, nil
}

func TestService_FindPaymentByID_success(t *testing.T) {
	//Создаём сервис
	s:=newTestService()
	_,payments, err :=s.addAccount(defaultTestAccount)
	if err!=nil {
		t.Error(err)
		return
	}
	//Пробуем найти платёж
	payment :=payments[0]
	got, err := s.FindPaymentByID(payment.ID)
	if err!=nil {
		t.Errorf("can`t find payment by ID: err=%v",err)
		return
	}
	//Сравниваем платежи
	if !reflect.DeepEqual(payment, got){
		t.Errorf("can`t find payment by ID: wrong payment returned=%v",err)
		return
	}
}
func TestService_FindPaymentByID_fail(t *testing.T) {
	//Создаём сервис
	s:=newTestService()
	_, _, err :=s.addAccount(defaultTestAccount)
	if err!=nil {
		t.Error(err)
		return
	}
	//Пробуем найти несуществующий платёж
	_, err = s.FindPaymentByID(uuid.New().String())
	if err==nil {
		t.Error("can`t find payment by ID: must return error, returned nil")
		return
	}
	
	if err!=ErrPaymentNotFound {
		t.Errorf("can`t find payment by ID: must return ErrPaymentNotFound, returned: %v", err)
	}
}

func TestService_Reject_success(t *testing.T) {
	//Создаём сервис
	s:=newTestService()
	_, payments, err :=s.addAccount(defaultTestAccount)
	if err!=nil {
		t.Error(err)
		return
	}
	//Пробуем отменить платёж
	payment:=payments[0]
	err = s.Reject(payment.ID)
	if err!=nil {
		t.Errorf("Reject: err=%v",err)
		return
	}
	savedPayment, err :=s.FindPaymentByID(payment.ID)
	if err!=nil {
		t.Errorf("Reject: can`t find payment by ID, err=%v",err)
		return
	}

	if savedPayment.Status!=types.PaymentStatusFail{
		t.Errorf("Reject: status didn`t changed payment=%v",savedPayment)
		return		
	}
	savedAccount, err := s.FindAccountByID(payment.AccountID)
	if savedAccount.Balance!= defaultTestAccount.balance{
		t.Errorf("Reject: balance didn`t changed Account=%v",savedAccount)
		return
	}
	
}

func TestService_Reject_Fail(t *testing.T) {
	//Создаём сервис
	svc:=&Service{}
	_,err := svc.FindPaymentByID("payment.ID")
	if err==nil{
		t.Error("Должна быть ошибка платёж не найден")
		return
	}
	//пробуем отменить несуществующий платёж
	err=svc.Reject("payment.ID")
	if err==nil{
		t.Error("Должна быть ошибка платёж не найден")
		return
	}
}

func TestService_Repeat_success(t *testing.T) {
	//Создаём сервис
	s:=newTestService()
	_, payments, err :=s.addAccount(defaultTestAccount)
	if err!=nil {
		t.Error(err)
		return
	}
	//Пробуем повторить платёж
	pay:=payments[0]
	got, err := s.Repeat(pay.ID)
	if err!=nil {
		t.Errorf("Repeat: err=%v",err)
		return
	}
	//Сравниваем платежи
	if !(pay.AccountID==got.AccountID&&pay.Amount==got.Amount&&pay.Category==got.Category&&pay.Status==got.Status){
		t.Errorf("can`t find payment by ID: wrong payment returned=%v",err)
		return
	}
	
}