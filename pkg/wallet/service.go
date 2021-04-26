package wallet

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/mijgona/wallet/pkg/types"
)

var ErrPhoneRegistered = errors.New("Phone already Registered")
var ErrAmountMustBePositive= errors.New("Amount must be greater than 0")
var ErrAccountNotFound= errors.New("Account not found")
var ErrNotEnoughBalance = errors.New("Balance not enough")
var ErrPaymentNotFound = errors.New("Payment Not Found")
var ErrFavoriteNotFound = errors.New("Favorite Not Found")


type Service struct{
	nextAccountID	int64
	accounts 		[]*types.Account
	payments 		[]*types.Payment
	favorites		[]*types.Favorite
}

//RegisterAccount регистрирует аккаут
func(s *Service)  RegisterAccount(phone types.Phone) (*types.Account, error) {
	for _, account := range s.accounts {
		if account.Phone==phone{
			return nil, ErrPhoneRegistered
		}
	}
	s.nextAccountID++
	account :=	&types.Account{
		ID:      s.nextAccountID,
		Phone:   phone,
		Balance: 0,
	}
	s.accounts=append(s.accounts, account)
	return account, nil
}

//Deposit пополняет счет аккаунта
func (s *Service) Deposit(accountID int64, amount types.Money) error {
	if amount<= 0{
		return ErrAmountMustBePositive
	}
	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID==accountID{
			account=acc
			break
		}		
	}
	if account==nil{
		return ErrAccountNotFound
	}

	account.Balance+=amount
	return nil
}

//Pay производит платеж
func (s *Service) Pay(accountID int64, amount types.Money, category types.PaymentCategory) (*types.Payment, error) {
	if amount<=0{
		return nil, ErrAmountMustBePositive
	}

	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID==accountID{
			account =acc
			break
		}			
	}
	if account==nil{
		return nil, ErrAccountNotFound
	}
	if account.Balance<amount{
		return nil, ErrNotEnoughBalance
	}
	account.Balance-=amount
	paymentID:=uuid.New().String()
	payment:=&types.Payment{
		ID:       	paymentID,
		AccountID:	accountID,
		Amount:   	amount,
		Category: 	category,
		Status:   	types.PaymentStatusInProgress,
	}

	s.payments = append(s.payments, payment)
	return payment, nil
}

//FindAccountByID ишет аккаут по заданнаму ID
func (s *Service) FindAccountByID(accountID int64) (*types.Account, error)  {
	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID==accountID{
			account=acc
			break
		}		
	}
	if account==nil{
		return nil, ErrAccountNotFound
	}
	return account,nil
}

//FindPaymentByID ишет платёж по ID
func (s *Service) FindPaymentByID(paymentID string) (*types.Payment, error)  {
	var payment *types.Payment
	for _, pay := range s.payments {
		if pay.ID==paymentID{
			payment=pay
		}
	}

	if payment==nil{
		return nil, ErrPaymentNotFound
	}
	return payment,nil
}

//Reject отменяет заданный платёж, если его статус INPROGRESS
func (s *Service) Reject(paymentID string) error  {
	payment, err :=s.FindPaymentByID(paymentID)
	if err!= nil{
		return err
	}

	account, err := s.FindAccountByID(payment.AccountID)
	if err!= nil{
		return err
	}

	payment.Status=types.PaymentStatusFail
	account.Balance+=payment.Amount
	return nil
}

//Repeat Повторяет заданный платёж
func (s *Service) Repeat(paymentID string) (*types.Payment, error)  {
	//Находим платёж
	payment, err :=s.FindPaymentByID(paymentID)
	if err!= nil{
		return nil,err
	}

	//Добавляем платёж повторно в сервис
	newPayment,err :=s.Pay(payment.AccountID,payment.Amount,payment.Category)
	if err!= nil{
		return nil,err
	}
	
	return newPayment, nil
}

//FavoritePayment создаёт избранное из конкретного платежа
func (s *Service) FavoritePayment(paymentID string, name string) (*types.Favorite, error) {
	//Находим платёж
	payment, err :=s.FindPaymentByID(paymentID)
	if err!= nil{
		return nil,err
	}

	//Создаём из платежа избранное
	favorite:=&types.Favorite{
		ID:        uuid.New().String(),
		AccountID: payment.AccountID,
		Name:      name,
		Amount:    payment.Amount,
		Category:  payment.Category,
	}

	//Добавляем его в сервис
	s.favorites = append(s.favorites, favorite)
	return favorite,nil
}

//FindFavoriteByID Находит избранное по ID
func (s *Service) FindFavoriteByID(favoriteID string) (*types.Favorite, error) {
	var favorite *types.Favorite
	for _, fav := range s.favorites {
		if fav.ID==favoriteID{
			favorite=fav
		}
	}

	if favorite==nil{
		return nil, ErrFavoriteNotFound
	}
	return favorite,nil
}

//PayFromFavorite Совершает платёж из конкретного избранного
func (s *Service) PayFromFavorite(favoriteID string) (*types.Payment, error) {
	//Находим избранное
	favorite, err :=s.FindFavoriteByID(favoriteID)
	if err!= nil{
		return nil,err
	}

	//совершаем платёж
	payment,err:=s.Pay(favorite.AccountID, favorite.Amount,favorite.Category)
	return payment,err
}

//ExportToFile экспортирует аккаунты в файл
func (s *Service) ExportToFile(path string) error  {
	file, err :=os.Create(path)
	if err!=nil {
		return err
	}
	defer func(){
		err := file.Close()
		if err!=nil {
			log.Print(err)
		}
	}()
	for _, account := range s.accounts {
		acc := strconv.FormatInt(account.ID,10)+";"+string(account.Phone)+";"+strconv.Itoa(int(account.Balance))+"|"
		_, err = file.Write([]byte(acc))
		if err!= nil {
			log.Print(err)
			return err
		}
	}


	return nil
}
//ImportFromFile импортирует аккаунты из файла
func (s *Service) ImportFromFile(path string) error  {
//открываем файл
	file, err :=os.Open(path)
	if err!=nil {
		return err
	}
	defer func(){
		err := file.Close()
		if err!=nil {
			log.Print(err)
		}
	}()
//Читаем содержимое
	content :=make([]byte,0)
	buf := make([]byte, 4)
	for {
	read, err := file.Read(buf)
		if err!= nil {
			break
		}
		content = append(content, buf[:read]...)
	}
//сортируем содержимое полей и обновляем сервис
	all:= strings.Split(string(content), "|")
	var phone string
	var id int64
	var balance int64
	acc:=all
	for _, str := range all {
		if str!=""{
		log.Println(str)
		acc = strings.Split(str,";")
		id,err=strconv.ParseInt(acc[0], 10, 64)
		phone =acc[1]
		balance,err =strconv.ParseInt(acc[2], 10, 64)
	
		account := &types.Account{
		ID:      id,
		Phone:   types.Phone(phone),
		Balance: types.Money(balance),
		}
		s.accounts = append(s.accounts, account)
		}
	}
	return nil
}

//Export экспортирует все данные в файл
func (s *Service) Export(dir string) error {
	
	accDir := dir+"/accounts.dump"
	accDir, err:=filepath.Abs(accDir)
	if err != nil {
		return err
	}
	
	err= s.exportAccounts(accDir)
	if err != nil {
		return err
	}

//Записываем платежи
	payDir:=dir+"/payments.dump"
	payDir, err=filepath.Abs(payDir)
	if err!=nil {
		return err
	}
	err = s.exportPayments(payDir)
	if err != nil {
		return err
	}

//записываем избранные платежи

	favDir:= dir+"/favorites.dump"
	favDir, err=filepath.Abs(favDir)
	if err!=nil {
		return err
	}	
 	err= s.exportFavorites(favDir)	
	if err != nil {
		return err
	}

	return nil
}

//Import импортирует все данные из файла
func (s *Service) Import(dir string) error {
	pathToAccount :=dir + "/accounts.dump"
	pathToAccount,err:=filepath.Abs(pathToAccount)
	if err != nil {
		return err
	}
	err=s.importAccount(pathToAccount)

	pathToPayment:= dir + "/payments.dump"
	pathToPayment, err =filepath.Abs(pathToPayment)
	if err != nil {
		return err
	}
	err=s.importPayments(pathToPayment)
	

	pathToFavorite := dir +"/favorites.dump"
	pathToFavorite, err= filepath.Abs(pathToFavorite)
	if err != nil {
		return err
	}
	err=s.importFavorites(pathToFavorite)

	return nil
}

func (s *Service) ExportAccountHistory(accountID int64) ([]types.Payment, error) {
	if s.payments!=nil{
		var found []types.Payment
		for _, payment := range s.payments {
			if payment.AccountID==accountID{
				found=append(found,*payment)
			}
		}
		return found,nil
		}
		return nil, ErrAccountNotFound
}
	

func (s *Service) HistoryToFiles(payments []types.Payment, dir string, records int) error{
	if payments!=nil{
		var pay string
		nextFile :=records
		currentFile:=1
		fileName:=1
		for _, payment := range payments {
			if len(payments)<=records {
				pathToPayment := dir+"/payments.dump"
				var pay string
				for _, payment := range payments {
					pay += payment.ID+";"+strconv.Itoa(int(payment.AccountID))+";"+strconv.Itoa(int(payment.Amount)) +";" + string(payment.Category) +";" + string(payment.Status)+"\n"
				}					
				err :=ioutil.WriteFile(pathToPayment,[]byte(pay), 0666)
				if err!=nil {
					return err
				}
				break	
			}
			if nextFile>=currentFile {					
				if currentFile==nextFile||payments[len(payments)-1]==payment{
					nextFile+=records
					currentFile++					
					pay +=payment.ID+";"+strconv.Itoa(int(payment.AccountID))+";"+strconv.Itoa(int(payment.Amount)) +";" + string(payment.Category) +";" + string(payment.Status)+"\n"
				} else {				
					currentFile++					
					pay +=payment.ID+";"+strconv.Itoa(int(payment.AccountID))+";"+strconv.Itoa(int(payment.Amount)) +";" + string(payment.Category) +";" + string(payment.Status)+"\n"
					continue
				}
			}
			pathToPayment := dir+"/payments"+strconv.Itoa(fileName)+".dump"	
			fileName++		
			err :=ioutil.WriteFile(pathToPayment,[]byte(pay), 0666)
			if err!=nil {
				return err
			}
			pay=""
		}
		return nil
	}
	
	return ErrPaymentNotFound
}
