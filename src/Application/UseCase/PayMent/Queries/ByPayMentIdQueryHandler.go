package useCase

import (
	"reflect"

	queries "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain.Contract/Queries/PayMent/ByIdPayMent"
	repository "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent/Repository"
	cqrs "github.com/mohammadrahimi/PayMent_Service_Go/src/Framework.Core/Bus"
)

type ByPayMentIdQueryHandler struct {
	repository repository.IPayMentRepository
}

func NewByPayMentIdQueryHandler(repository repository.IPayMentRepository) *ByPayMentIdQueryHandler {
	return &ByPayMentIdQueryHandler{
		repository: repository,
	}
}

func (h *ByPayMentIdQueryHandler) Handle(query cqrs.Query) (cqrs.QueryResult,error) {

			obj := reflect.ValueOf(query)
			byPayMentIdQuery := obj.Interface().(queries.ByIdPayMentQuery)
			
			queryResult,err:= h.repository.FindByPayMentId(byPayMentIdQuery.PayMentId) 
			if(err != nil){
				return   nil,err
			}
			return queryResult,nil
			
}