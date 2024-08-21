package useCase

import (
	"fmt"

	commands "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain.Contract/Commands/PayMent/Create"
	aggregate "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent"
	repository "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent/Repository"
	cqrs "github.com/mohammadrahimi/PayMent_Service_Go/src/Framework.Core/Bus"
)
 
type CreatePayMentCommandHandler struct {
	  repository  repository.IPayMentRepository
}

func NewCreatePayMentCommandHandler(repository  repository.IPayMentRepository) *CreatePayMentCommandHandler{
	return &CreatePayMentCommandHandler{
		repository: repository,
	}
}
 
func (h *CreatePayMentCommandHandler) Handle(command  cqrs.Command) error {
  
	switch c := command.(type) {

			case *commands.CreatePayMentCommand:{
				 
				  payment,err:= aggregate.NewPayMent(c)
				  if(err != nil){
					  return err
				  }
				 
				status,err:= h.repository.Create(payment)
				if(err != nil){
					return err
				}

				fmt.Println("  status = " +   status  )    

			}
    }
	
	return nil
}