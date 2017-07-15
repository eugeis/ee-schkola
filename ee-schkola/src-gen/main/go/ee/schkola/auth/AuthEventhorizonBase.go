package auth

import (
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

type AccountAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewAccountAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *AccountAggregateInitializer, err error) {
    ret = &AccountAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *AccountAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, AccountAggregateType, AccountCommandTypes().Literals())
}




const AccountAggregateType eventhorizon.AggregateType = "AccountAggregate"
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











