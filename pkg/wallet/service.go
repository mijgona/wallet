package wallet

import (
	"github.com/google/uuid"
	"errors"
	"github.com/mijgona/wallet/pkg/types"
	"log"
	"os"
	"strings"
	"strconv"
)

var (
	ErrPhoneRegistred       = errors.New("Phone already registred")
	ErrAmountMustBePositive = errors.New("amount must be greater than 0")
	ErrAccountNotFound      = errors.New("Account not found")
	ErrNotEnoughBalance     = errors.New("Not enagh balance")
	ErrPaymentNotFound      = errors.New("Payment Not Found")
	ErrFavoriteNotFound      = errors.New("Favorite Not Found")
)

//Service serv
type Service struct {
	nextAccountID int64
	accounts      	[]*types.Account
	payments      	[]*types.Payment
	favorites		[]*types.Favorite
}

//RegisterAccount method for registration new account
func (s *Service) RegisterAccount(phone types.Phone) (*types.Account, error) {
	for _, account := range s.accounts {
		if account.Phone == phone {
			return nil, ErrPhoneRegistred
		}

	}
	s.nextAccountID++
	account := &types.Account{
		ID:      s.nextAccountID,
		Phone:   phone,
		Balance: 0,
	}
	s.accounts = append(s.accounts, account)
	return account, nil
}

//Deposit method
func (s *Service) Deposit(accountID int64, amount types.Money) error {
	if amount <= 0 {
		return ErrAmountMustBePositive
	}
	var account *types.Account

	for _, acc := range s.accounts {
		if acc.ID == accountID {
			account = acc
			break
		}

	}

	if account == nil {
		return ErrAccountNotFound
	}
	//Зачисление средств пока не рассматриваем как платеж
	account.Balance += amount
	return nil
}

//Pay метод для регистрации платижа
func (s *Service) Pay(accountID int64, amount types.Money, category types.PaymentCategory) (*types.Payment, error) {
	if amount <= 0 {
		return nil, ErrAmountMustBePositive
	}

	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == accountID {
			account = acc
			break
		}
	}
	if account == nil {
		return nil, ErrAccountNotFound
	}
	if account.Balance < amount {
		return nil, ErrNotEnoughBalance
	}
	account.Balance -= amount
	paymentID := uuid.New().String()
	payment := &types.Payment{
		ID:        paymentID,
		AccountID: accountID,
		Amount:    amount,
		Category:  category,
		Status:    types.StatusInProgress,
	}
	s.payments = append(s.payments, payment)
	return payment, nil
}

// FindAccountByID ищем пользователя по ID
func (s *Service) FindAccountByID(accountID int64) (*types.Account, error) {
	for _, account := range s.accounts {
		if account.ID == accountID {
			return account, nil
		}
	}
	return nil, ErrAccountNotFound
}

// FindPaymentByID ищем платёж по ID
func (s *Service) FindPaymentByID(paymentID string) (*types.Payment, error) {
	for _, payment := range s.payments {
		if payment.ID == paymentID {
			return payment, nil
		}
	}
	return nil, ErrPaymentNotFound
}



func (s *Service) Reject(paymentID string) error {
	pay, err := s.FindPaymentByID(paymentID)
	if err != nil {
		return err
	}

	acc, err := s.FindAccountByID(pay.AccountID)
	if err != nil {
		return err
	}

	pay.Status = types.StatusFail
	acc.Balance += pay.Amount

	return nil
}
//Repeat add a payment to repeat it
func (s *Service) Repeat(paymentID string)(*types.Payment, error)  {
	pay, err := s.FindPaymentByID(paymentID)
	if err != nil {
		return nil,err
	}
	pay.ID=uuid.New().String()
	return pay, nil
}


//FavoritePayment add payment to favorites
func (s *Service) FavoritePayment(paymentID string, name string) (*types.Favorite, error)  {
	gotPayment, err := s.FindPaymentByID(paymentID)
	if err != nil {
		return nil,err
	}
	favorite := &types.Favorite{
		ID:        	gotPayment.ID,
		AccountID: 	gotPayment.AccountID,
		Name:		name,
		Amount:    	gotPayment.Amount,
		Category:  	gotPayment.Category,
	}
	s.favorites= append (s.favorites, favorite)
	return favorite, nil
}

func (s *Service) PayFromFavorite(favoriteID string)(*types.Payment,error)  {
	var favorite *types.Favorite

	for _, fav := range s.favorites {
		if fav.ID == favoriteID {
			favorite=fav
		}else{			
		err:= ErrFavoriteNotFound
		return nil, err
		}
	}
	payment:= &types.Payment{
		ID:				favorite.ID,
		AccountID:		favorite.AccountID,
		Amount:			favorite.Amount,
		Category:		favorite.Category,
		Status:			types.StatusInProgress,
	}
	payment, err := s.Repeat(payment.ID)
	if err != nil {
		return nil,err
	}
	return payment, nil
}

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

func (s *Service) ImportFromFile(path string) error  {
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
	content :=make([]byte,0)
	buf := make([]byte, 4)
	for {
	read, err := file.Read(buf)
		if err!= nil {
			break
		}
		content = append(content, buf[:read]...)
	}
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
