package queries
 

type ByIdOrderQuery struct {
	OrderId     string
}

func NewByIdOrderQuery(orderId string) ByIdOrderQuery{
       return ByIdOrderQuery{
		   OrderId: orderId,
	   }
}

