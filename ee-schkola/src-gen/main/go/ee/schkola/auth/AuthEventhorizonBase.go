package auth

import (
    "github.com/looplab/eventhorizon"
)
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
    
    aggregateType := eventhorizon.AggregateType(AuthAggregateTypes().Account)
    for _, command := range AccountCommandTypes().Values() {
        handler.SetAggregate(aggregateType, eventhorizon.CommandType(command.Name()))
    }
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










type AuthAggregateType struct {
	name  string
	ordinal int
}

func (o *AuthAggregateType) Name() string {
    return o.name
}

func (o *AuthAggregateType) Ordinal() int {
    return o.ordinal
}

func (o *AuthAggregateType) IsAccount() bool {
    return o == _authAggregateTypes.Account()
}

type authAggregateTypes struct {
	values []*AuthAggregateType
}

var _authAggregateTypes = &authAggregateTypes{values: []*AuthAggregateType{
    {name: "Account", ordinal: 0}},
}

func AuthAggregateTypes() *authAggregateTypes {
	return _authAggregateTypes
}

func (o *authAggregateTypes) Values() []*AuthAggregateType {
	return o.values
}

func (o *authAggregateTypes) Account() *AuthAggregateType {
    return _authAggregateTypes.values[0]
}

func (o *authAggregateTypes) ParseAuthAggregateType(name string) (ret *AuthAggregateType, ok bool) {
    switch name {
      case o.Account().Name():
        ret = o.Account()
    }
    return
}



