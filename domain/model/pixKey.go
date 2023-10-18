package model

import (
	uuid "github.com/satori/go.uuid"
	"github.com/asaskevich/govalidator"
	"time"
)

type PixKeyRepositoryInterface interface{
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(Key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
}


type PixKey struct {
	Base				`valid: "required"`
	Kind		string	`json: "kind" valid: "required"`
	Key			string	`json: "key" valid: "required"`
	AccountID	string	`json: "account_id" valid: "required"`
	Account		*Account`valid: "-"`
	Status		string	`json: "status" valid: "required"`
}


func(pixKey *PixKey) isValid() error{
	_, err := govalidator.ValidateStruct(pixKey)

	if pixKey.Kind != "email" && pixKey.Kind != "cpf"{
		return errors.New(text: "Invalid type of key")
	}
	
	if pixKey.Status != "active" && pixKey.Status != "inative"{
		return errors.New(text: "Invalid status")
	}

	if err != nil {
		return err
	}
	return nil
}



func NewPixKey(kind string, account *Account, key string) (*PixKey, error){
	pixKey := PixKey{
		Kind: kind,
		Key: key,
		Account: account,
		Status: "active",
	}

	pixKey.ID = uuid.NewV4().string()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}

	return &pixKey, nil

}