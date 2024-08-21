package queries

import "time"

type ResultPayMentQuery struct {
	PayMentId       string
	CreatedAt       time.Time
	OrderId         string
	UserId          string
	Description     string
	TransactionCode int32
	TansactionState string
}

func NewResultPayMentQuery(payMentId string, createdAt time.Time, orderId string, 
	userId string, description string,transactionCode int32,tansactionState string) ResultPayMentQuery {
	return ResultPayMentQuery{
		PayMentId:      payMentId,
		CreatedAt:      createdAt,
		OrderId:        orderId,
		UserId:         userId,
		Description:    description,
		TransactionCode: transactionCode,
		TansactionState: tansactionState,
	}
}
