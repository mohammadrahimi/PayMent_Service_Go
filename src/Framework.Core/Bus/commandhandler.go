package cqrs


type CommandHandler interface {
	Handle(Command) error
}
 
