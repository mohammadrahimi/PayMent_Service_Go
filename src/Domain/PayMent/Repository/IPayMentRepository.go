package repository

import (
	queries "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain.Contract/Queries/PayMent/ResultPayMent"
	aggregate "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent"
)

type IPayMentRepository interface {
	Create(payment *aggregate.PayMent) ( status string, err error)
	FindByPayMentId(id string) (queries.ResultPayMentQuery, error)
	FindByOrderId(id string) (queries.ResultPayMentQuery, error)
	FindByUserId(id string) (queries.ResultPayMentQuery, error)
	FindAll() ([]queries.ResultPayMentQuery, error)
	 
}