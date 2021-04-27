package wallet

import (
	"fmt"
	"log"
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
	 favorites [] struct {
		 name 		string		
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
		favorites: []struct{
			name 		string
		}{
			{
				name: "audi",
			},
		},
 }

 //Функция конструктор
 func newTestService () *testService {
	 return &testService{Service: &Service{}}
 }

 //Создаём аккаунт для тестирования
func (s *testService) addAccount(data testAccount) (*types.Account, []*types.Payment, []*types.Favorite, error)  {
 //Регистрируем пользователя
 account, err := s.RegisterAccount(data.phone)
	if err !=nil {
		return nil, nil, nil, fmt.Errorf("can`t register account, error=%v", err)
	}
//Пополняет его счет
err=s.Deposit(account.ID,data.balance)
	if err!=nil{
		return nil, nil, nil, fmt.Errorf("can`t deposit account, error=%v", err)
	}

//выполняем платежи
//Можем сразу создать слайс нужной длины, поскольку знаем размер
payments := make([]*types.Payment, len(data.payments))
 for i, payment := range data.payments {
	 //Работаем через индекс
	 payments[i], err = s.Pay(account.ID, payment.amount, payment.category)
	 	if err!=nil{
		return nil, nil, nil, fmt.Errorf("can`t make payment, error=%v", err)
	}
 }

 //делаем из платежа избранное
favorites:=make([]*types.Favorite, len(data.favorites))
 for i, favorite := range data.favorites {
	 //Работаем через индекс
	 favorites[i], err = s.FavoritePayment(payments[0].ID, favorite.name)
	 	if err!=nil{
		return nil, nil, nil, fmt.Errorf("can`t make favorite, error=%v", err)
	}
 }
return account, payments, favorites, nil
}

func TestService_FindPaymentByID_success(t *testing.T) {
	//Создаём сервис
	s:=newTestService()
	_,payments,_, err :=s.addAccount(defaultTestAccount)
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
	_, _, _, err :=s.addAccount(defaultTestAccount)
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
	_, payments, _, err :=s.addAccount(defaultTestAccount)
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
	_, payments,_, err :=s.addAccount(defaultTestAccount)
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

func TestService_FavoritePayment_success(t *testing.T) {
	//Создаём сервис
	s:=newTestService()
	_, _,favorites, err :=s.addAccount(defaultTestAccount)
	if err!=nil {
		t.Error(err)
		return
	}
	//совершаем платёж из избранного
	fav:=favorites[0]
	_, err=s.PayFromFavorite(fav.ID)
	if err!=nil {
		t.Error(err)
		return
	}
}


func TestService_Export_success(t *testing.T) {
//Создаём сервис
s:=newTestService()

//импортируем
_, _,_, err :=s.addAccount(defaultTestAccount)
err=s.Import("../../data")
if err != nil {
	t.Errorf("Невозможно выполнить импорт, ошибка=%v",err)
}	
		
err = s.Export("../../data")
if err != nil {
	t.Errorf("Невозможно выполнить экспорт, ошибка=%v",err)
}
}
func TestService_Import_success(t *testing.T) {
//Создаём сервис
s:=newTestService()

//импортируем
err:=s.Import("../../data")
if err != nil {
	t.Errorf("Невозможно выполнить импорт, ошибка=%v",err)
}	
}

func TestService_History_success(t *testing.T) {
	//Создаём сервис
	s:=newTestService()	
	_, _,_, err :=s.addAccount(defaultTestAccount)
	err=s.Import("../../data")
	if err != nil {
		t.Errorf("Невозможно выполнить импорт, ошибка=%v",err)
	}	
	payment, err :=s.ExportAccountHistory(1)
	if err != nil {
		t.Errorf("Невозможно выполнить Экспорт', ошибка=%v",err)
	}
	
	err=s.HistoryToFiles(payment, "../../data/payments",3)
	if err != nil {
		t.Errorf("Невозможно записать историю в файл, ошибка=%v",err)
	}
	
	}

	func TestService_SumPayments_success(t *testing.T) {
		s := newTestService()
		dir:="../../data"
		s.Import(dir)
		sum:= s.SumPayments(5)
		log.Print(sum)
		
	}

	func BenchmarkSumPayments(b *testing.B) {
		s := newTestService()
		dir:="../../data"
		s.Import(dir)
		want := types.Money(100000)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result :=s.SumPayments(55)
			b.StopTimer()
			if result!=want {
				b.Fatalf("invalid result, got %v, want %v", result, want)
			}
			b.StartTimer()			
		}
	}

	func BenchmarkFilterPayments(b *testing.B) {
		s := newTestService()
		dir:="../../data"
		s.Import(dir)
		want,_ := s.FindPaymentsByID(1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result,err :=s.FilterPayments(1,56)
			if err != nil {
				b.Fatalf("Невозможно записать историю в файл, ошибка=%v",err)
			}
			b.StopTimer()
				//Сравниваем платежи
				l:=equal(want, result)
			if !l{
				b.Fatalf("can`t find payments: wrong payment returned=%v want=%v",result,want)
				return
			}
			b.StartTimer()			
		}
	}

// Equal проверяет, что a и b содержат одинаковые элементы.
// nil аргумент эквивалентен пустому срезу.
func equal(a, b []types.Payment) bool {
    if len(a) != len(b) {
        return false
    }
		
    for _, v := range a {
		exist:=false
		for _, d := range b {	
			if v == d {
				exist=true
				break
			}
		}
		if exist {
			exist=!exist
		}else{
			return false		
	}
    }
    return true
}