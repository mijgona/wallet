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
type Service struct{
	nextAccountID	int64
	accounts 		[]*types.Account
	payments 		[]*types.Payment
}

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