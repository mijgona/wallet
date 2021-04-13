package wallet

import (
	"errors"

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

//Находит избранное по ID
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

//Совершает платёж из конкретного избранного
func (s *Service) PayFromFavorite(favoriteID string) (*types.Payment, error) {
	//Находим избранное
	favorite, err :=s.FindFavoriteByID(favoriteID)
	if err!= nil{
		return nil,err
	}

	//совершаем платёж
	payment,err:=s.Repeat(favorite.ID)
	return payment,err
}
