package wallet

import (
	"github.com/google/uuid"
	"errors"
	"github.com/mijgona/wallet/pkg/types"
)

var (
	ErrPhoneRegistred       = errors.New("Phone already registred")
	ErrAmountMustBePositive = errors.New("amount must be greater than 0")
	ErrAccountNotFound      = errors.New("Account not found")
	ErrNotEnoughBalance     = errors.New("Not enagh balance")
	ErrPaymentNotFound      = errors.New("Payment Not Found"))

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