package wallet

import (
	"strconv"
	"os"
)

func (s *Service) exportAccounts(accDir string) error {
	var acc string

//Записываем аккаунты
	file, err := os.Create(accDir)
	if err != nil {
		return err
	}
	defer func() {
		cerr := file.Close()
		if err == nil {
			err = cerr
		}
	}()
	for _, account := range s.accounts {
		acc=strconv.FormatInt(account.ID,10)+";"+string(account.Phone)+";"+strconv.Itoa(int(account.Balance))+"\n"
		_, err := file.Write([]byte(acc))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) exportPayments(payDir string) error {	
	if len(s.payments) == 0 {
		return nil
	}
	
	file, err := os.Create(payDir)
	if err != nil {
		return err
	}
	defer func() {
		cerr := file.Close()
		if err == nil {
			err = cerr
		}
	}()

	var pay string
	for _, payment := range s.payments {
		pay= string(payment.ID) + ";" +  strconv.FormatInt(payment.AccountID,10) + ";" + strconv.Itoa(int(payment.Amount)) + ";" +  string(payment.Category) + ";" +  string(payment.Status)+"\n"
		_, err := file.Write([]byte(pay))
		if err!=nil {
			return err
		}
	}
	
	return nil
}

func (s *Service) exportFavorites(favDir string) error {
	if len(s.favorites) == 0 {
		return nil
	}
	
	file, err := os.Create(favDir)
	if err != nil {
		return err
	}
	defer func() {
		cerr := file.Close()
		if err == nil {
			err = cerr
		}
	}()
	var fav string
	for _, favorite := range s.favorites {
		fav= string(favorite.ID) + ";" + strconv.FormatInt(favorite.AccountID,10) + ";" + string(favorite.Name) +  ";" + strconv.Itoa(int(favorite.Amount)) +  ";" + string(favorite.Category)+"\n"
		_, err := file.Write([]byte(fav))
		if err!=nil {
			return err
		}
	}
	return nil
}