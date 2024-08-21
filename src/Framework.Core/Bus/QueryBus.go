package cqrs

import (
	"fmt"
	"reflect"
)

type IQueryBus interface {
	Send(Query) error
	RegisterHandler(QueryHandler, Query)
}

type QueryBus struct {
	handlers map[string]QueryHandler
}

func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[string]QueryHandler),
	}
}

func (b *QueryBus) Send(query Query) (QueryResult, error) {

    var r QueryResult

	var typeName string
	if t := reflect.TypeOf(query); t.Kind() == reflect.Ptr {
		typeName = t.Elem().Name()
    } else {
		typeName  = t.Name()
    }
	 
	if handler, ok := b.handlers[typeName]; ok {
		return handler.Handle(query)
	}
	return    r,fmt.Errorf("The Query bus does not have a handler for Queries of type: %s", typeName) 

}

func (b *QueryBus) RegisterHandler(handler QueryHandler, query Query) {

	var typeName string
	if t := reflect.TypeOf(query); t.Kind() == reflect.Ptr {
		typeName = t.Elem().Name()
    } else {
		typeName  = t.Name()
    }
  
	b.handlers[typeName] = handler

}
