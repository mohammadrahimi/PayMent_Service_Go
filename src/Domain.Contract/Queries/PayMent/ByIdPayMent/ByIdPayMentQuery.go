package queries

type ByIdPayMentQuery struct {
	PayMentId string
}

func NewByIdPayMentQuery(payMentId string) ByIdPayMentQuery {
	return ByIdPayMentQuery{
		PayMentId: payMentId,
	}
}
