package model

import (
	uuid "github.com/satori/go.uuid"
	"github.com/asaskevich/govalidator"
	"time"
)

const ( 
	TransactionPending		string = "pending",
	TransactionCompleted 	string = "completed"
	TransactionError 		string = "Error"
	TransactionConfirmed 	string = "confirmed"
)

type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

type Transaction struct {
	Transaction []Transaction
}

type Transaction struct{
	Base 							`valid: "required"`
	AccountFrom 		*Account	`valid: "-"`
	Amount				float64		`json:"amount" valid: "notnull"`
	PixKeyTo			*PixKey		`valid: "-"`
	Status				string		`json:"status" valid: "notnull"`
	Description			string		`json: "description" valid: "notnull"`
	CancelDescription 	string		`json: "cancel_description" valid: "-"`
}


func(t *Transaction) isValid() error{
	_, err := govalidator.ValidateStruct(t)

	if t.amount <= 0{
		return errors.New(text: "the amount must be greater than 0")
	}

	if t.Status != TransactionPending && t.Status != TransactionCompleted && t.Status != TransactionError {
		return errors.New(text: "invalid status for the transaction")
	}

	if t.PixKeyTo.AccountID == t.AccountFrom.ID {
		return errors.New("you can't send money to yourself")
	}

	if err != nil {
		return err
	}
	return nil
}


func NewTransaction(AccountFrom *Account, amount float64, pixKeyTo *PixKey, description string,) (*Transaction, error){
	transaction := Transaction{
		AccountFrom: accountFrom,
		Amount: amount,
		PixKeyTo: pixKeyTo,
		Status: TransactionPending,
		Description: description,
		Account: account,
		Status: "active",
	}

	transaction.ID = uuid.NewV4().string()
	transaction.CreatedAt = time.Now()

	err := transaction.isValid()
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (t *Transaction) Complete() error {
	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Transaction) Confirm() error {
	t.Status = TransactionConfirmed
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Transaction) Cancel(description string) error {
	t.Status = TransactionError
	t.UpdatedAt = time.Now()
	t.Description = description
	err := t.isValid()
	return err
}