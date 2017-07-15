package auth

import (
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "time"
)

type AccountCreate struct {
    Username  string
    Password  string
    Email  string
    Disabled  bool
    LastLoginAt  *time.Time
    Profile  *person.Profile
}

func NewAccountCreate(username string, password string, email string, disabled bool, lastLoginAt *time.Time, profile *person.Profile) (ret *AccountCreate, err error) {
    ret = &AccountCreate{
        Username : username,
        Password : password,
        Email : email,
        Disabled : disabled,
        LastLoginAt : lastLoginAt,
        Profile : profile,
    }
    return
}



type AccountDelete struct {
    Id  string
}

func NewAccountDelete(id string) (ret *AccountDelete, err error) {
    ret = &AccountDelete{
        Id : id,
    }
    return
}



type AccountUpdate struct {
    Username  string
    Password  string
    Email  string
    Disabled  bool
    LastLoginAt  *time.Time
    Profile  *person.Profile
}

func NewAccountUpdate(username string, password string, email string, disabled bool, lastLoginAt *time.Time, profile *person.Profile) (ret *AccountUpdate, err error) {
    ret = &AccountUpdate{
        Username : username,
        Password : password,
        Email : email,
        Disabled : disabled,
        LastLoginAt : lastLoginAt,
        Profile : profile,
    }
    return
}



type AccountEnable struct {
}



type AccountDisable struct {
}



type AccountRegister struct {
    Username  string
    Email  string
    Password  string
}

func NewAccountRegister(username string, email string, password string) (ret *AccountRegister, err error) {
    ret = &AccountRegister{
        Username : username,
        Email : email,
        Password : password,
    }
    return
}




type AccountAggregateCommandType struct {
	name  string
	ordinal int
}

func (o *AccountAggregateCommandType) Name() string {
    return o.name
}

func (o *AccountAggregateCommandType) Ordinal() int {
    return o.ordinal
}

func (o *AccountAggregateCommandType) IsCreateAccount() bool {
    return o == _accountAggregateCommandTypes.CreateAccount()
}

func (o *AccountAggregateCommandType) IsDeleteAccount() bool {
    return o == _accountAggregateCommandTypes.DeleteAccount()
}

func (o *AccountAggregateCommandType) IsUpdateAccount() bool {
    return o == _accountAggregateCommandTypes.UpdateAccount()
}

func (o *AccountAggregateCommandType) IsEnableAccount() bool {
    return o == _accountAggregateCommandTypes.EnableAccount()
}

func (o *AccountAggregateCommandType) IsDisableAccount() bool {
    return o == _accountAggregateCommandTypes.DisableAccount()
}

func (o *AccountAggregateCommandType) IsRegisterAccount() bool {
    return o == _accountAggregateCommandTypes.RegisterAccount()
}

type accountAggregateCommandTypes struct {
	values []*AccountAggregateCommandType
    literals []enum.Literal
}

var _accountAggregateCommandTypes = &accountAggregateCommandTypes{values: []*AccountAggregateCommandType{
    {name: "createAccount", ordinal: 0},
    {name: "deleteAccount", ordinal: 1},
    {name: "updateAccount", ordinal: 2},
    {name: "enableAccount", ordinal: 3},
    {name: "disableAccount", ordinal: 4},
    {name: "registerAccount", ordinal: 5}},
}

func AccountAggregateCommandTypes() *accountAggregateCommandTypes {
	return _accountAggregateCommandTypes
}

func (o *accountAggregateCommandTypes) Values() []*AccountAggregateCommandType {
	return o.values
}

func (o *accountAggregateCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *accountAggregateCommandTypes) CreateAccount() *AccountAggregateCommandType {
    return _accountAggregateCommandTypes.values[0]
}

func (o *accountAggregateCommandTypes) DeleteAccount() *AccountAggregateCommandType {
    return _accountAggregateCommandTypes.values[1]
}

func (o *accountAggregateCommandTypes) UpdateAccount() *AccountAggregateCommandType {
    return _accountAggregateCommandTypes.values[2]
}

func (o *accountAggregateCommandTypes) EnableAccount() *AccountAggregateCommandType {
    return _accountAggregateCommandTypes.values[3]
}

func (o *accountAggregateCommandTypes) DisableAccount() *AccountAggregateCommandType {
    return _accountAggregateCommandTypes.values[4]
}

func (o *accountAggregateCommandTypes) RegisterAccount() *AccountAggregateCommandType {
    return _accountAggregateCommandTypes.values[5]
}

func (o *accountAggregateCommandTypes) ParseAccountAggregateCommandType(name string) (ret *AccountAggregateCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*AccountAggregateCommandType), ok
	}
	return
}



