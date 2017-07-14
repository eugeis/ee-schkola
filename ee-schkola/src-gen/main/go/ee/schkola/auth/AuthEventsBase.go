package auth

import (
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "time"
)

type AccountCreated struct {
    Username  string
    Password  string
    Email  string
    Disabled  bool
    LastLoginAt  *time.Time
    Profile  *person.Profile
}

func NewAccountCreated(username string, password string, email string, disabled bool, lastLoginAt *time.Time, profile *person.Profile) (ret *AccountCreated, err error) {
    ret = &AccountCreated{
        Username : username,
        Password : password,
        Email : email,
        Disabled : disabled,
        LastLoginAt : lastLoginAt,
        Profile : profile,
    }
    return
}



type AccountDeleted struct {
    Id  string
}

func NewAccountDeleted(id string) (ret *AccountDeleted, err error) {
    ret = &AccountDeleted{
        Id : id,
    }
    return
}



type AccountUpdated struct {
    Username  string
    Password  string
    Email  string
    Disabled  bool
    LastLoginAt  *time.Time
    Profile  *person.Profile
}

func NewAccountUpdated(username string, password string, email string, disabled bool, lastLoginAt *time.Time, profile *person.Profile) (ret *AccountUpdated, err error) {
    ret = &AccountUpdated{
        Username : username,
        Password : password,
        Email : email,
        Disabled : disabled,
        LastLoginAt : lastLoginAt,
        Profile : profile,
    }
    return
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

func (o *AccountEventType) IsCreatedAccount() bool {
    return o == _accountEventTypes.CreatedAccount()
}

func (o *AccountEventType) IsDeletedAccount() bool {
    return o == _accountEventTypes.DeletedAccount()
}

func (o *AccountEventType) IsUpdatedAccount() bool {
    return o == _accountEventTypes.UpdatedAccount()
}

type accountEventTypes struct {
	values []*AccountEventType
    literals []enum.Literal
}

var _accountEventTypes = &accountEventTypes{values: []*AccountEventType{
    {name: "createdAccount", ordinal: 0},
    {name: "deletedAccount", ordinal: 1},
    {name: "updatedAccount", ordinal: 2}},
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

func (o *accountEventTypes) CreatedAccount() *AccountEventType {
    return _accountEventTypes.values[0]
}

func (o *accountEventTypes) DeletedAccount() *AccountEventType {
    return _accountEventTypes.values[1]
}

func (o *accountEventTypes) UpdatedAccount() *AccountEventType {
    return _accountEventTypes.values[2]
}

func (o *accountEventTypes) ParseAccountEventType(name string) (ret *AccountEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*AccountEventType), ok
	}
	return
}



