package auth

import (
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

const AccountAggregateType eventhorizon.AggregateType = "AccountAggregate"

type AccountAggregateInitializer struct {
    *eh.AggregateInitializer
}

func (o *AccountAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountAggregateEventTypes().AccountCreated())
}

func (o *AccountAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountAggregateEventTypes().AccountDeleted())
}

func (o *AccountAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountAggregateEventTypes().AccountUpdated())
}

func NewAccountAggregateInitializer(
	store *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher,
	commandBus *eventhorizon.CommandBus) (ret *AccountAggregateInitializer) {
	ret = &AccountAggregateInitializer{
        AggregateInitializer: eh.NewAggregateInitializer(AccountAggregateType, AccountAggregateCommandTypes().Literals(),
		ChurchAggregateEventTypes().Literals(), store, eventBus, publisher, commandBus),
    }
	return
}

type AccountAggregate struct {
    *eventhorizon.AggregateBase
    *Account
}

func NewAccountAggregate(AggregateBase *eventhorizon.AggregateBase, Entity *Account) (ret *AccountAggregate, err error) {
    ret = &AccountAggregate{
        AggregateBase: AggregateBase,
        Account: Entity,
    }
    return
}



type AuthEventhorizonInitializer struct {
    Store  *eventhorizon.EventStore
    EventBus  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    CommandBus  *eventhorizon.CommandBus
}

func NewAuthEventhorizonInitializer(store *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                commandBus *eventhorizon.CommandBus) (ret *AuthEventhorizonInitializer, err error) {
    ret = &AuthEventhorizonInitializer{
        Store : store,
        EventBus : eventBus,
        Publisher : publisher,
        CommandBus : commandBus,
    }
    return
}











