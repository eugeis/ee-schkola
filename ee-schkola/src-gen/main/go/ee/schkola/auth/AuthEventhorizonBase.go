package auth

import (
    "github.com/looplab/eventhorizon"
)
type AccountAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func (o *AccountAggregateInitializer) setup() string {
    panic("Not implemented yet.")
}

func (o *AccountAggregateInitializer) registerCommands() string {
    panic("Not implemented yet.")
}


type AuthEventhorizonInitializer struct {
    Store  *eventhorizon.EventStore
    EventBus  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    CommandBus  *eventhorizon.CommandBus
}

func (o *AuthEventhorizonInitializer) setup() string {
    panic("Not implemented yet.")
}

func (o *AuthEventhorizonInitializer) registerCommands() string {
    panic("Not implemented yet.")
}


type AccountAggregate struct {
    *eventhorizon.AggregateBase
    *Account
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
      case "Account":
        ret = o.Account()
    }
    return
}



