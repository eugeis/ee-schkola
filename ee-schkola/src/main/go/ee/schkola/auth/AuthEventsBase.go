package auth

import (
    "ee/schkola"
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "github.com/looplab/eventhorizon"
)
const (
     AccountCreatedEvent eventhorizon.EventType = "AccountCreated"
     AccountDeletedEvent eventhorizon.EventType = "AccountDeleted"
     AccountLoggedEvent eventhorizon.EventType = "AccountLogged"
     AccountUpdatedEvent eventhorizon.EventType = "AccountUpdated"
     AccountEnabledEvent eventhorizon.EventType = "AccountEnabled"
     AccountDisabledEvent eventhorizon.EventType = "AccountDisabled"
)




type AccountCreated struct {
    Name *schkola.PersonName`eh:"optional"`
    Username string`eh:"optional"`
    Password string`eh:"optional"`
    Email string`eh:"optional"`
    Disabled bool`eh:"optional"`
    Roles []string`eh:"optional"`
    Profile *person.Profile`eh:"optional"`
    Id eventhorizon.UUID
}

func (o *AccountCreated) AddToRoles(item string) string {
    o.Roles = append(o.Roles, item)
    return item
}


type AccountDeleted struct {
    Id eventhorizon.UUID
}


type AccountLogged struct {
    Username string`eh:"optional"`
    Email string`eh:"optional"`
    Password string`eh:"optional"`
    Id eventhorizon.UUID
}


type AccountUpdated struct {
    Name *schkola.PersonName`eh:"optional"`
    Username string`eh:"optional"`
    Password string`eh:"optional"`
    Email string`eh:"optional"`
    Disabled bool`eh:"optional"`
    Roles []string`eh:"optional"`
    Profile *person.Profile`eh:"optional"`
    Id eventhorizon.UUID
}

func (o *AccountUpdated) AddToRoles(item string) string {
    o.Roles = append(o.Roles, item)
    return item
}


type AccountEnabled struct {
    Id eventhorizon.UUID
}


type AccountDisabled struct {
    Id eventhorizon.UUID
}




type AccountEventType struct {
	name  string
	ordinal int
}

func (o *AccountEventType) Name() string {
    return o.name
}

func (o *AccountEventType) Ordinal() int {
    return o.ordinal
}

func (o *AccountEventType) IsAccountCreated() bool {
    return o == _accountEventTypes.AccountCreated()
}

func (o *AccountEventType) IsAccountDeleted() bool {
    return o == _accountEventTypes.AccountDeleted()
}

func (o *AccountEventType) IsAccountLogged() bool {
    return o == _accountEventTypes.AccountLogged()
}

func (o *AccountEventType) IsAccountUpdated() bool {
    return o == _accountEventTypes.AccountUpdated()
}

func (o *AccountEventType) IsAccountEnabled() bool {
    return o == _accountEventTypes.AccountEnabled()
}

func (o *AccountEventType) IsAccountDisabled() bool {
    return o == _accountEventTypes.AccountDisabled()
}

type accountEventTypes struct {
	values []*AccountEventType
    literals []enum.Literal
}

var _accountEventTypes = &accountEventTypes{values: []*AccountEventType{
    {name: "AccountCreated", ordinal: 0},
    {name: "AccountDeleted", ordinal: 1},
    {name: "AccountLogged", ordinal: 2},
    {name: "AccountUpdated", ordinal: 3},
    {name: "AccountEnabled", ordinal: 4},
    {name: "AccountDisabled", ordinal: 5}},
}

func AccountEventTypes() *accountEventTypes {
	return _accountEventTypes
}

func (o *accountEventTypes) Values() []*AccountEventType {
	return o.values
}

func (o *accountEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *accountEventTypes) AccountCreated() *AccountEventType {
    return _accountEventTypes.values[0]
}

func (o *accountEventTypes) AccountDeleted() *AccountEventType {
    return _accountEventTypes.values[1]
}

func (o *accountEventTypes) AccountLogged() *AccountEventType {
    return _accountEventTypes.values[2]
}

func (o *accountEventTypes) AccountUpdated() *AccountEventType {
    return _accountEventTypes.values[3]
}

func (o *accountEventTypes) AccountEnabled() *AccountEventType {
    return _accountEventTypes.values[4]
}

func (o *accountEventTypes) AccountDisabled() *AccountEventType {
    return _accountEventTypes.values[5]
}

func (o *accountEventTypes) ParseAccountEventType(name string) (ret *AccountEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*AccountEventType), ok
	}
	return
}



