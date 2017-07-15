package auth

import (
    "context"
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

const AccountAggregateType eventhorizon.AggregateType = "AccountAggregate"

func NewAccountAggregate(id eventhorizon.UUID) *AccountAggregate {
	return &AccountAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(AccountAggregateType, id),
	}
}

func (o *AccountAggregate) HandleCommand(ctx context.Context, cmd eventhorizon.Command) error {
    println("HandleCommand %v - %v", ctx, cmd)
    return nil
}

func (o *AccountAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent %v - %v", ctx, event)
    return nil
}

type AccountAggregate struct {
    *eventhorizon.AggregateBase
    *Account
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
    accountAggregateInitializer: NewAccountAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus)}
	return
}

func (o *AuthEventhorizonInitializer) Setup() (err error) {
    
    if err = o.accountAggregateInitializer.Setup(); err != nil {
        return
    }
    return
}

type AuthEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore
    eventBus eventhorizon.EventBus
    eventPublisher eventhorizon.EventPublisher
    commandBus eventhorizon.CommandBus
    accountAggregateInitializer *AccountAggregateInitializer
}









