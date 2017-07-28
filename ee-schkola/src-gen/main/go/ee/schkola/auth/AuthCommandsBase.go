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
    Id  eventhorizon.UUID`eh:"optional"`
    Username  string`eh:"optional"`
    Password  string`eh:"optional"`
    Email  string`eh:"optional"`
    Disabled  bool`eh:"optional"`
    LastLoginAt  *time.Time`eh:"optional"`
    Profile  *person.Profile`eh:"optional"`
}

func NewCreate() (ret *CreateAccount) {
    ret = &CreateAccount{}
    return
}
func (o *CreateAccount) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *CreateAccount) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *CreateAccount) CommandType() eventhorizon.CommandType      { return CreateAccountCommand }



        

type DeleteAccount struct {
    Id  eventhorizon.UUID
}

func NewDelete() (ret *DeleteAccount) {
    ret = &DeleteAccount{}
    return
}
func (o *DeleteAccount) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *DeleteAccount) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *DeleteAccount) CommandType() eventhorizon.CommandType      { return DeleteAccountCommand }



        

type UpdateAccount struct {
    Id  eventhorizon.UUID`eh:"optional"`
    Username  string`eh:"optional"`
    Password  string`eh:"optional"`
    Email  string`eh:"optional"`
    Disabled  bool`eh:"optional"`
    LastLoginAt  *time.Time`eh:"optional"`
    Profile  *person.Profile`eh:"optional"`
}

func NewUpdate() (ret *UpdateAccount) {
    ret = &UpdateAccount{}
    return
}
func (o *UpdateAccount) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *UpdateAccount) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *UpdateAccount) CommandType() eventhorizon.CommandType      { return UpdateAccountCommand }



        

type EnableAccount struct {
}

func NewEnable() (ret *EnableAccount) {
    ret = &EnableAccount{}
    return
}
func (o *EnableAccount) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *EnableAccount) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *EnableAccount) CommandType() eventhorizon.CommandType      { return EnableAccountCommand }



        

type DisableAccount struct {
}

func NewDisable() (ret *DisableAccount) {
    ret = &DisableAccount{}
    return
}
func (o *DisableAccount) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *DisableAccount) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *DisableAccount) CommandType() eventhorizon.CommandType      { return DisableAccountCommand }



        

type RegisterAccount struct {
    Username  string
    Email  string
    Password  string
}

func NewRegister() (ret *RegisterAccount) {
    ret = &RegisterAccount{}
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
    {name: "CreateAccount", ordinal: 0},
    {name: "DeleteAccount", ordinal: 1},
    {name: "UpdateAccount", ordinal: 2},
    {name: "EnableAccount", ordinal: 3},
    {name: "DisableAccount", ordinal: 4},
    {name: "RegisterAccount", ordinal: 5}},
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



