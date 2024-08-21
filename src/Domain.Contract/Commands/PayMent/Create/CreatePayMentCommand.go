package commands

import (
	"github.com/mohammadrahimi/PayMent_Service_Go/src/Domain.Contract/Dto/PayMent/DtoTransaction"
	  
)
 

type CreatePayMentCommand struct {
    OrderId        string
    UserId         string
    Description    string
    Transaction    DtoTransaction.DtoTransaction
}

func NewCreatePayMentCommand(orderId string,userId string,description string,dtoTransaction DtoTransaction.DtoTransaction) CreatePayMentCommand{
       return CreatePayMentCommand{
		OrderId: orderId,
		UserId: userId,
		Description: description,
		Transaction: dtoTransaction,
	   }
}
