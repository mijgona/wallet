package types

//Money представляет собой денежную сумму в минимальных единицах(центы, дирамы, копейки) и т.д.
type Money int64

//Category представляет собой категорию, в которой был совершён платеж (авто, аптеки, рестораны и т.д.)
type PaymentCategory string

//Status представляет собой статус платежа
type PaymentStatus string

//Предопределённые статусы платежей
const(
	PaymentStatusOk PaymentStatus="OK"
	PaymentStatusFail PaymentStatus="FAIL"
	PaymentStatusInProgress PaymentStatus="INPROGRESS"
)

//Payment представляет информацию о платеже
type Payment struct {
	ID 			string	
	AccountID	int64
	Amount		Money
	Category	PaymentCategory
	Status		PaymentStatus
}

//Phone хранит номер телефона
type Phone string

//Account предоставляет информацию о счёте пользователя
type Account struct{
	ID			int64
	Phone		Phone
	Balance		Money
}

type Favorite struct{
	ID			string
	AccountID	int64
	Name		string
	Amount		Money
	Category	PaymentCategory
}