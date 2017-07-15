package auth

import (
    "ee/schkola/person"
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/enum"
    "time"
)
const (
     CreateAccountCommand eventhorizon.CommandType = "CreateAccount"
     DeleteAccountCommand eventhorizon.CommandType = "DeleteAccount"
     UpdateAccountCommand eventhorizon.CommandType = "UpdateAccount"
     EnableAccountCommand eventhorizon.CommandType = "EnableAccount"
     DisableAccountCommand eventhorizon.CommandType = "DisableAccount"
     RegisterAccountCommand eventhorizon.CommandType = "RegisterAccount"
)




        

type CreateAccount struct {
    Username  string
    Password  string
    Email  string
    Disabled  bool
    LastLoginAt  *time.Time
    Profile  *person.Profile
}

func NewCreateAccount(username string, password string, email string, disabled bool, lastLoginAt *time.Time, profile *person.Profile) (ret *CreateAccount, err error) {
    ret = &CreateAccount{
        Username : username,
        Password : password,
        Email : email,
        Disabled : disabled,
        LastLoginAt : lastLoginAt,
        Profile : profile,
    }
    
    return
}
func (o *CreateAccount) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *CreateAccount) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *CreateAccount) CommandType() eventhorizon.CommandType      { return CreateAccountCommand }



        

type DeleteAccount struct {
    Id  string
}

func NewDeleteAccount(id string) (ret *DeleteAccount, err error) {
    ret = &DeleteAccount{
        Id : id,
    }
    
    return
}
func (o *DeleteAccount) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *DeleteAccount) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *DeleteAccount) CommandType() eventhorizon.CommandType      { return DeleteAccountCommand }



        

type UpdateAccount struct {
    Username  string
    Password  string
    Email  string
    Disabled  bool
    LastLoginAt  *time.Time
    Profile  *person.Profile
}

func NewUpdateAccount(username string, password string, email string, disabled bool, lastLoginAt *time.Time, profile *person.Profile) (ret *UpdateAccount, err error) {
    ret = &UpdateAccount{
        Username : username,
        Password : password,
        Email : email,
        Disabled : disabled,
        LastLoginAt : lastLoginAt,
        Profile : profile,
    }
    
    return
}
func (o *UpdateAccount) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *UpdateAccount) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *UpdateAccount) CommandType() eventhorizon.CommandType      { return UpdateAccountCommand }



        

type EnableAccount struct {
}
func (o *EnableAccount) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *EnableAccount) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *EnableAccount) CommandType() eventhorizon.CommandType      { return EnableAccountCommand }



        

type DisableAccount struct {
}
func (o *DisableAccount) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *DisableAccount) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *DisableAccount) CommandType() eventhorizon.CommandType      { return DisableAccountCommand }



        

type RegisterAccount struct {
    Username  string
    Email  string
    Password  string
}

func NewRegisterAccount(username string, email string, password string) (ret *RegisterAccount, err error) {
    ret = &RegisterAccount{
        Username : username,
        Email : email,
        Password : password,
    }
    
    return
}
func (o *RegisterAccount) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *RegisterAccount) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *RegisterAccount) CommandType() eventhorizon.CommandType      { return RegisterAccountCommand }





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

func (o *AccountCommandType) IsAccountCreate() bool {
    return o == _accountCommandTypes.AccountCreate()
}

func (o *AccountCommandType) IsAccountDelete() bool {
    return o == _accountCommandTypes.AccountDelete()
}

func (o *AccountCommandType) IsAccountUpdate() bool {
    return o == _accountCommandTypes.AccountUpdate()
}

func (o *AccountCommandType) IsAccountEnable() bool {
    return o == _accountCommandTypes.AccountEnable()
}

func (o *AccountCommandType) IsAccountDisable() bool {
    return o == _accountCommandTypes.AccountDisable()
}

func (o *AccountCommandType) IsAccountRegister() bool {
    return o == _accountCommandTypes.AccountRegister()
}

type accountCommandTypes struct {
	values []*AccountCommandType
    literals []enum.Literal
}

var _accountCommandTypes = &accountCommandTypes{values: []*AccountCommandType{
    {name: "AccountCreate", ordinal: 0},
    {name: "AccountDelete", ordinal: 1},
    {name: "AccountUpdate", ordinal: 2},
    {name: "AccountEnable", ordinal: 3},
    {name: "AccountDisable", ordinal: 4},
    {name: "AccountRegister", ordinal: 5}},
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

func (o *accountCommandTypes) AccountCreate() *AccountCommandType {
    return _accountCommandTypes.values[0]
}

func (o *accountCommandTypes) AccountDelete() *AccountCommandType {
    return _accountCommandTypes.values[1]
}

func (o *accountCommandTypes) AccountUpdate() *AccountCommandType {
    return _accountCommandTypes.values[2]
}

func (o *accountCommandTypes) AccountEnable() *AccountCommandType {
    return _accountCommandTypes.values[3]
}

func (o *accountCommandTypes) AccountDisable() *AccountCommandType {
    return _accountCommandTypes.values[4]
}

func (o *accountCommandTypes) AccountRegister() *AccountCommandType {
    return _accountCommandTypes.values[5]
}

func (o *accountCommandTypes) ParseAccountCommandType(name string) (ret *AccountCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*AccountCommandType), ok
	}
	return
}



