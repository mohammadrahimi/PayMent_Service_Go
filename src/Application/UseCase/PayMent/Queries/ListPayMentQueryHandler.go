package useCase

import (
	 
	repository "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent/Repository"
	cqrs "github.com/mohammadrahimi/PayMent_Service_Go/src/Framework.Core/Bus"
)

type ListPayMentQueryHandler struct {
	repository repository.IPayMentRepository
}

func NewListPayMentQueryHandler(repository repository.IPayMentRepository) *ListPayMentQueryHandler {
	return &ListPayMentQueryHandler{
		repository: repository,
	}
}

func (h *ListPayMentQueryHandler) Handle(query cqrs.Query)  (cqrs.QueryResult,error)  {  
  
			 
			  queryResult,err:= h.repository.FindAll() 
			  if(err != nil){
				  return   nil,err
			  }
			  return queryResult,nil
			 
}