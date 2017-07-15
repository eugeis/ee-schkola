package auth

import (
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "time"
)

type CreatedAccount struct {
    Username  string
    Password  string
    Email  string
    Disabled  bool
    LastLoginAt  *time.Time
    Profile  *person.Profile
}

func NewCreatedAccount(username string, password string, email string, disabled bool, lastLoginAt *time.Time, profile *person.Profile) (ret *CreatedAccount, err error) {
    ret = &CreatedAccount{
        Username : username,
        Password : password,
        Email : email,
        Disabled : disabled,
        LastLoginAt : lastLoginAt,
        Profile : profile,
    }
    return
}



type DeletedAccount struct {
    Id  string
}

func NewDeletedAccount(id string) (ret *DeletedAccount, err error) {
    ret = &DeletedAccount{
        Id : id,
    }
    return
}



type UpdatedAccount struct {
    Username  string
    Password  string
    Email  string
    Disabled  bool
    LastLoginAt  *time.Time
    Profile  *person.Profile
}

func NewUpdatedAccount(username string, password string, email string, disabled bool, lastLoginAt *time.Time, profile *person.Profile) (ret *UpdatedAccount, err error) {
    ret = &UpdatedAccount{
        Username : username,
        Password : password,
        Email : email,
        Disabled : disabled,
        LastLoginAt : lastLoginAt,
        Profile : profile,
    }
    return
}




type AccountAggregateEventType struct {
	name  string
	ordinal int
}

func (o *AccountAggregateEventType) Name() string {
    return o.name
}

func (o *AccountAggregateEventType) Ordinal() int {
    return o.ordinal
}

func (o *AccountAggregateEventType) IsAccountCreated() bool {
    return o == _accountAggregateEventTypes.AccountCreated()
}

func (o *AccountAggregateEventType) IsAccountDeleted() bool {
    return o == _accountAggregateEventTypes.AccountDeleted()
}

func (o *AccountAggregateEventType) IsAccountUpdated() bool {
    return o == _accountAggregateEventTypes.AccountUpdated()
}

type accountAggregateEventTypes struct {
	values []*AccountAggregateEventType
    literals []enum.Literal
}

var _accountAggregateEventTypes = &accountAggregateEventTypes{values: []*AccountAggregateEventType{
    {name: "AccountCreated", ordinal: 0},
    {name: "AccountDeleted", ordinal: 1},
    {name: "AccountUpdated", ordinal: 2}},
}

func AccountAggregateEventTypes() *accountAggregateEventTypes {
	return _accountAggregateEventTypes
}

func (o *accountAggregateEventTypes) Values() []*AccountAggregateEventType {
	return o.values
}

func (o *accountAggregateEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *accountAggregateEventTypes) AccountCreated() *AccountAggregateEventType {
    return _accountAggregateEventTypes.values[0]
}

func (o *accountAggregateEventTypes) AccountDeleted() *AccountAggregateEventType {
    return _accountAggregateEventTypes.values[1]
}

func (o *accountAggregateEventTypes) AccountUpdated() *AccountAggregateEventType {
    return _accountAggregateEventTypes.values[2]
}

func (o *accountAggregateEventTypes) ParseAccountAggregateEventType(name string) (ret *AccountAggregateEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*AccountAggregateEventType), ok
	}
	return
}



