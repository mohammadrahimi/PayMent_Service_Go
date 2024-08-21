package cqrs

 
type QueryHandler interface {
	Handle(Query) (QueryResult, error)
}

type QueryResult interface{}
