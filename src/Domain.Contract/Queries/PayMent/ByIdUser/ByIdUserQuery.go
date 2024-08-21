package queries
 

type ByIdUserQuery struct {
	UserId     string
}

func NewByIdUserQuery(userId string) ByIdUserQuery{
       return ByIdUserQuery{
		UserId: userId,
	   }
}
