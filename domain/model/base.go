package model


import (
	"github.com/asaskevich/govalidator"
	"time"
)

func init(){
	govalidator.SetFieldsRequiredByDefault(value: true)
}
type Base struct {
	ID 			string		`json: "id"         valid: "uuid"` 
	CreatedAt 	time.Time	`json: "created-at" valid: "-"`
	UpdatedAt 	time.Time	`json: "updated-at" valid: "-"`
}