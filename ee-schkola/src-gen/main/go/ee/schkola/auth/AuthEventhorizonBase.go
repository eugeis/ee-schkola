package auth

import (
    "context"
    "errors"
    "fmt"
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

const AccountAggregateType eventhorizon.AggregateType = "AccountAggregate"

func NewAccountAggregate(id eventhorizon.UUID) (ret *AccountAggregate) {
    ret = &AccountAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(AccountAggregateType, id),
    }
	ret.CommandHandler = NewAccountAggregateCommandHandler(ret)
    return
}

func (o *AccountAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}


func NewAccountAggregateCommandHandler(aggregate *AccountAggregate) *AccountAggregateCommandHandler {
	return &AccountAggregateCommandHandler{
		aggregate: aggregate,
        handlers: make(map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *AccountAggregate) error),
    }
}

type AccountAggregateCommandHandler struct {
	aggregate *AccountAggregate
	handlers  map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *AccountAggregate) error
}

func (o *AccountAggregateCommandHandler) AddHandler(commandType eventhorizon.CommandType,
	handler func(cmd eventhorizon.Command, aggregate *AccountAggregate) error) {
	o.handlers[commandType] = handler
}

func (o *AccountAggregateCommandHandler) HandleCommand(ctx context.Context, cmd eventhorizon.Command) (err error) {
	if handler, ok := o.handlers[cmd.CommandType()]; ok {
		err = handler(cmd, o.aggregate)
	} else {
		err = errors.New(fmt.Sprintf("There is no handlers for command %v registered in the aggregate %v",
			cmd.CommandType(), cmd.AggregateType()))
	}
	return
}

type AccountAggregate struct {
    *eventhorizon.AggregateBase
    *Account
    eventhorizon.CommandHandler
}



func NewAccountAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *AccountAggregateInitializer) {
	ret = &AccountAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(AccountAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate { return NewAccountAggregate(id) },
        AccountCommandTypes().Literals(), AccountEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
    }
	return
}


func (o *AccountAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountEventTypes().AccountCreated())
}

func (o *AccountAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountEventTypes().AccountDeleted())
}

func (o *AccountAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountEventTypes().AccountUpdated())
}

type AccountAggregateInitializer struct {
    *eh.AggregateInitializer
}



func NewAuthEventhorizonInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *AuthEventhorizonInitializer) {
	ret = &AuthEventhorizonInitializer{eventStore: eventStore, eventBus: eventBus, eventPublisher: eventPublisher,
            commandBus: commandBus, 
    AccountAggregateInitializer: NewAccountAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus)}
	return
}

func (o *AuthEventhorizonInitializer) Setup() (err error) {
    
    if err = o.AccountAggregateInitializer.Setup(); err != nil {
        return
    }
    return
}

type AuthEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore
    eventBus eventhorizon.EventBus
    eventPublisher eventhorizon.EventPublisher
    commandBus eventhorizon.CommandBus
    AccountAggregateInitializer  *AccountAggregateInitializer
}









