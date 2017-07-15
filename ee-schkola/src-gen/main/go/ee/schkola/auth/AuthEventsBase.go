package auth

import (
    "ee/schkola/person"
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/enum"
    "time"
)

type AccountCreated struct {
    Id  eventhorizon.UUID`eh:"optional"`
    Username  string`eh:"optional"`
    Password  string`eh:"optional"`
    Email  string`eh:"optional"`
    Disabled  bool`eh:"optional"`
    LastLoginAt  *time.Time`eh:"optional"`
    Profile  *person.Profile`eh:"optional"`
}



type AccountDeleted struct {
    Id  eventhorizon.UUID
}



type AccountUpdated struct {
    Id  eventhorizon.UUID`eh:"optional"`
    Username  string`eh:"optional"`
    Password  string`eh:"optional"`
    Email  string`eh:"optional"`
    Disabled  bool`eh:"optional"`
    LastLoginAt  *time.Time`eh:"optional"`
    Profile  *person.Profile`eh:"optional"`
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

func (o *AccountEventType) IsAccountUpdated() bool {
    return o == _accountEventTypes.AccountUpdated()
}

type accountEventTypes struct {
	values []*AccountEventType
    literals []enum.Literal
}

var _accountEventTypes = &accountEventTypes{values: []*AccountEventType{
    {name: "AccountCreated", ordinal: 0},
    {name: "AccountDeleted", ordinal: 1},
    {name: "AccountUpdated", ordinal: 2}},
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

func (o *accountEventTypes) AccountUpdated() *AccountEventType {
    return _accountEventTypes.values[2]
}

func (o *accountEventTypes) ParseAccountEventType(name string) (ret *AccountEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*AccountEventType), ok
	}
	return
}



