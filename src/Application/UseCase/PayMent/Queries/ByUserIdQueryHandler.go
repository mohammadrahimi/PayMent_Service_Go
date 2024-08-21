package useCase

import (
	"reflect"

	queries "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain.Contract/Queries/PayMent/ByIdUser"
	repository "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent/Repository"
	cqrs "github.com/mohammadrahimi/PayMent_Service_Go/src/Framework.Core/Bus"
)

type ByUserIdQueryHandler struct {
	repository repository.IPayMentRepository
}

func NewByUserIdQueryHandler(repository repository.IPayMentRepository) *ByUserIdQueryHandler {
	return &ByUserIdQueryHandler{
		repository: repository,
	}
}

func (h *ByUserIdQueryHandler) Handle(query cqrs.Query) (cqrs.QueryResult,error)  {  

	 
			obj := reflect.ValueOf(query)
			byUserIdQuery := obj.Interface().(queries.ByIdUserQuery)
	 
			queryResult,err:= h.repository.FindByUserId(byUserIdQuery.UserId) 
			if(err != nil){
				return   nil,err
			}
			return queryResult,nil
	 
}