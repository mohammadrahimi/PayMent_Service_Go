package useCase

import (
	"reflect"

	queries "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain.Contract/Queries/PayMent/ByIdOrder"
	repository "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent/Repository"
	cqrs "github.com/mohammadrahimi/PayMent_Service_Go/src/Framework.Core/Bus"
)

type ByOrderIdQueryHandler struct {
	repository repository.IPayMentRepository
}

func NewByOrderIdQueryHandler(repository repository.IPayMentRepository) *ByOrderIdQueryHandler {
	return &ByOrderIdQueryHandler{
		repository: repository,
	}
}

func (h *ByOrderIdQueryHandler) Handle(query cqrs.Query) (cqrs.QueryResult,error) {

			obj := reflect.ValueOf(query)
			byOrderIdQuery := obj.Interface().(queries.ByIdOrderQuery)
			
			queryResult,err:= h.repository.FindByOrderId(byOrderIdQuery.OrderId) 
			if(err != nil){
				return   nil,err
			}
			return queryResult,nil
			
}