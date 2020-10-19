package types

// Money представляет собой денежную сумму в минимальных единицах (центы, копейки дирамы и.т.д.)
type Money int64

//PaymentCategory катеория платежей
type PaymentCategory string

//PaymentStatus представляет собой статус платежа
type PaymentStatus string

//предопределленные статусы платежей
const (
	StatusOk         PaymentStatus = "OK"
	StatusFail       PaymentStatus = "FAIL"
	StatusInProgress PaymentStatus = "INPROGRESS"
)

//Payment информация о платежах
type Payment struct {
	ID        string
	AccountID int64
	Amount    Money
	Category  PaymentCategory
	Status    PaymentStatus
}

//Favorite payments
type Favorite struct {
	ID        	string
	AccountID 	int64
	Name		string
	Amount    	Money
	Category  	PaymentCategory
}

//Phone number
type Phone string

//Account information
type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}
