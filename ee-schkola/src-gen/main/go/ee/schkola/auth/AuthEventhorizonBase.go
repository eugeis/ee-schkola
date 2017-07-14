package auth

import (
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

type AccountAccountAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewAccountAccountAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *AccountAccountAggregateInitializer, err error) {
    ret = &AccountAccountAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *AccountAccountAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, AccountAggregateType, AccountCommandTypes().Literals())
}




const AccountAccountType eventhorizon.AggregateType = "AccountAccount"
type AccountAccount struct {
    *eventhorizon.AggregateBase
    *Account
}

func NewAccountAccount(AggregateBase *eventhorizon.AggregateBase, Entity *Account) (ret *AccountAccount, err error) {
    ret = &AccountAccount{
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











