package cqrs

import (
	"fmt"
	"reflect"
)

type ICommandBus interface {
	Send(Command) error
	RegisterHandler(CommandHandler, Command)
}

type CommandBus struct {
	handlers map[string]CommandHandler
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[string]CommandHandler),
	}
}

func (b *CommandBus) Send(command Command) error {

	typeName := reflect.TypeOf(command).Elem().Name()
	if handler, ok := b.handlers[typeName]; ok {
		return handler.Handle(command)
	}
	return fmt.Errorf("The command bus does not have a handler for commands of type: %s", typeName)

}

func (b *CommandBus) RegisterHandler(handler CommandHandler, command Command) {

	typeName := reflect.TypeOf(command).Elem().Name()
	b.handlers[typeName] = handler

}
