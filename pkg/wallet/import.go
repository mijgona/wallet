package wallet

import (
	"strings"
	"bufio"
	"log"
	"io"
	"strconv"
	"os"
	"reflect"
	"github.com/mijgona/wallet/pkg/types"

)


func (s *Service) importAccount(pathToAccount string) error  {
	_, err := os.Stat(pathToAccount)
	if err != nil {
		return nil
	}
	//открываем файл
	srcAccount, err :=os.Open(pathToAccount)
	if err!=nil {
		return nil
	}
	defer func(){
		cerr := srcAccount.Close()
		if cerr!=nil {
			log.Print(err)
		}
	}()
//Читаем содержимое файла аккаутов
	reader:= bufio.NewReader(srcAccount)
	for {

			line, err :=reader.ReadString('\n')
			if err==io.EOF{
				log.Print(line)
				break
			}
			if err!=nil {
				return err
			}
			line=strings.ReplaceAll(line,"\n","")
			line=strings.ReplaceAll(line,"\r","")
			accLine := strings.Split(line,";")
			id,err:=strconv.ParseInt(accLine[0], 10, 64)
			if err!=nil {
				return err
			}
			phone :=accLine[1]
			balance,err :=strconv.ParseInt(accLine[2], 10, 64)
			if err!=nil {
				return err
			}
			account := &types.Account{
				ID:      id,
				Phone:   types.Phone(phone),
				Balance: types.Money(balance),
			}
			gotAccount,err:=s.FindAccountByID(account.ID)
			if !reflect.DeepEqual(account, gotAccount){
				s.accounts = append(s.accounts, account)
				s.nextAccountID= s.accounts[len(s.accounts)-1].ID
			}
		}
	return nil
}

func (s *Service) importPayments(pathToPayment string) error  {
_, err := os.Stat(pathToPayment)
if err != nil {
	return nil
}
//Читаем содержимое файла платежей
srcPayment, err :=os.Open(pathToPayment)
if err!=nil {
	return nil
}
defer func(){
	cerr := srcPayment.Close()
	if cerr!=nil {
		log.Print(err)
	}
}()
reader :=bufio.NewReader(srcPayment)
for {
	line, err :=reader.ReadString('\n')
	if err==io.EOF{
		break
	}
	if err!=nil {
		return err
	}	
	line=strings.ReplaceAll(line,"\n","")
	line=strings.ReplaceAll(line,"\r","")
	payLine := strings.Split(line,";")
	id:=payLine[0]
	accountID,err :=strconv.ParseInt(payLine[1],10,64)
	if err!=nil {
		return err
	}
	amount,err :=strconv.ParseInt(payLine[2], 10, 64)
	category := payLine[3]
	status := payLine[4]
	if err!=nil {
		return err
	}
	payment := &types.Payment{
		ID:      	id,
		AccountID:	accountID,
		Amount:  	types.Money(amount),
		Category:	types.PaymentCategory(category),
		Status:		types.PaymentStatus(status),
	}
	
	gotPayment,err:=s.FindPaymentByID(payment.ID)
	if !reflect.DeepEqual(payment, gotPayment){
		s.payments = append(s.payments, payment)
	}
}
return nil
}

func (s *Service) importFavorites(pathToFavorite string) error {
	_, err := os.Stat(pathToFavorite)
	if err != nil {
		return nil
	}
//Читаем содержимое файла избранных

srcFavorite, err :=os.Open(pathToFavorite)
if err!=nil {
	return nil
}
defer func(){
	cerr := srcFavorite.Close()
	if cerr!=nil {
		log.Print(err)
	}
}()
reader :=bufio.NewReader(srcFavorite)
for {
	line, err :=reader.ReadString('\n')		
	line=strings.ReplaceAll(line,"\n","")
	line=strings.ReplaceAll(line,"\r","")
	if err==io.EOF{
		break
	}
	favLine := strings.Split(line,";")
	id:=favLine[0]
	accountID,err :=strconv.ParseInt(favLine[1],10,64)
	if err!=nil {
		return err
	}
	name :=favLine[2]
	amount,err :=strconv.ParseInt(favLine[3], 10, 64)
	category := favLine[4]
	if err!=nil {
		return err
	}
	favorite := &types.Favorite{
		ID:      	id,
		AccountID:	accountID,
		Name:		name,
		Amount:  	types.Money(amount),
		Category:	types.PaymentCategory(category),
	}
			
	gotFavorite,err:=s.FindFavoriteByID(favorite.ID)
	if !reflect.DeepEqual(favorite, gotFavorite){
		s.favorites = append(s.favorites, favorite)
	}
}


return nil
}