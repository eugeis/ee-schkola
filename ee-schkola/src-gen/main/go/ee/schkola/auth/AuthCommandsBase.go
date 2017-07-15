package auth

import (
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "time"
)

type AccountCreate struct {
    username string
    password string
    email string
    disabled bool
    lastLoginAt *time.Time
    profile *person.Profile
}

func NewAccountCreate(username string, password string, email string, disabled bool, lastLoginAt *time.Time, profile *person.Profile) (ret *AccountCreate, err error) {
    ret = &AccountCreate{
        username: username,
        password: password,
        email: email,
        disabled: disabled,
        lastLoginAt: lastLoginAt,
        profile: profile,
    }
    
    return
}



type AccountDelete struct {
    id string
}

func NewAccountDelete(id string) (ret *AccountDelete, err error) {
    ret = &AccountDelete{
        id: id,
    }
    
    return
}



type AccountUpdate struct {
    username string
    password string
    email string
    disabled bool
    lastLoginAt *time.Time
    profile *person.Profile
}

func NewAccountUpdate(username string, password string, email string, disabled bool, lastLoginAt *time.Time, profile *person.Profile) (ret *AccountUpdate, err error) {
    ret = &AccountUpdate{
        username: username,
        password: password,
        email: email,
        disabled: disabled,
        lastLoginAt: lastLoginAt,
        profile: profile,
    }
    
    return
}



type AccountEnable struct {
}



type AccountDisable struct {
}



type AccountRegister struct {
    username string
    email string
    password string
}

func NewAccountRegister(username string, email string, password string) (ret *AccountRegister, err error) {
    ret = &AccountRegister{
        username: username,
        email: email,
        password: password,
    }
    
    return
}




type AccountCommandType struct {
	name  string
	ordinal int
}

func (o *AccountCommandType) Name() string {
    return o.name
}

func (o *AccountCommandType) Ordinal() int {
    return o.ordinal
}

func (o *AccountCommandType) IsCreateAccount() bool {
    return o == _accountCommandTypes.CreateAccount()
}

func (o *AccountCommandType) IsDeleteAccount() bool {
    return o == _accountCommandTypes.DeleteAccount()
}

func (o *AccountCommandType) IsUpdateAccount() bool {
    return o == _accountCommandTypes.UpdateAccount()
}

func (o *AccountCommandType) IsEnableAccount() bool {
    return o == _accountCommandTypes.EnableAccount()
}

func (o *AccountCommandType) IsDisableAccount() bool {
    return o == _accountCommandTypes.DisableAccount()
}

func (o *AccountCommandType) IsRegisterAccount() bool {
    return o == _accountCommandTypes.RegisterAccount()
}

type accountCommandTypes struct {
	values []*AccountCommandType
    literals []enum.Literal
}

var _accountCommandTypes = &accountCommandTypes{values: []*AccountCommandType{
    {name: "createAccount", ordinal: 0},
    {name: "deleteAccount", ordinal: 1},
    {name: "updateAccount", ordinal: 2},
    {name: "enableAccount", ordinal: 3},
    {name: "disableAccount", ordinal: 4},
    {name: "registerAccount", ordinal: 5}},
}

func AccountCommandTypes() *accountCommandTypes {
	return _accountCommandTypes
}

func (o *accountCommandTypes) Values() []*AccountCommandType {
	return o.values
}

func (o *accountCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *accountCommandTypes) CreateAccount() *AccountCommandType {
    return _accountCommandTypes.values[0]
}

func (o *accountCommandTypes) DeleteAccount() *AccountCommandType {
    return _accountCommandTypes.values[1]
}

func (o *accountCommandTypes) UpdateAccount() *AccountCommandType {
    return _accountCommandTypes.values[2]
}

func (o *accountCommandTypes) EnableAccount() *AccountCommandType {
    return _accountCommandTypes.values[3]
}

func (o *accountCommandTypes) DisableAccount() *AccountCommandType {
    return _accountCommandTypes.values[4]
}

func (o *accountCommandTypes) RegisterAccount() *AccountCommandType {
    return _accountCommandTypes.values[5]
}

func (o *accountCommandTypes) ParseAccountCommandType(name string) (ret *AccountCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*AccountCommandType), ok
	}
	return
}



