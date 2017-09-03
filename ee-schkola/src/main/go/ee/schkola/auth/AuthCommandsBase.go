package auth

import (
    "ee/schkola"
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "github.com/looplab/eventhorizon"
)
const (
     EnableAccountCommand eventhorizon.CommandType = "EnableAccount"
     DisableAccountCommand eventhorizon.CommandType = "DisableAccount"
     CreateAccountCommand eventhorizon.CommandType = "CreateAccount"
     DeleteAccountCommand eventhorizon.CommandType = "DeleteAccount"
     UpdateAccountCommand eventhorizon.CommandType = "UpdateAccount"
)




        
type EnableAccount struct {
    Id eventhorizon.UUID
}

func NewEnable() (ret *EnableAccount) {
    ret = &EnableAccount{}
    return
}
func (o *EnableAccount) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *EnableAccount) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *EnableAccount) CommandType() eventhorizon.CommandType      { return EnableAccountCommand }



        
type DisableAccount struct {
    Id eventhorizon.UUID
}

func NewDisable() (ret *DisableAccount) {
    ret = &DisableAccount{}
    return
}
func (o *DisableAccount) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *DisableAccount) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *DisableAccount) CommandType() eventhorizon.CommandType      { return DisableAccountCommand }



        
type CreateAccount struct {
    Name *schkola.PersonName`eh:"optional"`
    Username string`eh:"optional"`
    Password string`eh:"optional"`
    Email string`eh:"optional"`
    Profile *person.Profile`eh:"optional"`
    Id eventhorizon.UUID`eh:"optional"`
}
func (o *CreateAccount) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *CreateAccount) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *CreateAccount) CommandType() eventhorizon.CommandType      { return CreateAccountCommand }



        
type DeleteAccount struct {
    Id eventhorizon.UUID
}
func (o *DeleteAccount) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *DeleteAccount) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *DeleteAccount) CommandType() eventhorizon.CommandType      { return DeleteAccountCommand }



        
type UpdateAccount struct {
    Name *schkola.PersonName`eh:"optional"`
    Username string`eh:"optional"`
    Password string`eh:"optional"`
    Email string`eh:"optional"`
    Profile *person.Profile`eh:"optional"`
    Id eventhorizon.UUID`eh:"optional"`
}
func (o *UpdateAccount) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *UpdateAccount) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *UpdateAccount) CommandType() eventhorizon.CommandType      { return UpdateAccountCommand }





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

func (o *AccountCommandType) IsEnableAccount() bool {
    return o == _accountCommandTypes.EnableAccount()
}

func (o *AccountCommandType) IsDisableAccount() bool {
    return o == _accountCommandTypes.DisableAccount()
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

type accountCommandTypes struct {
	values []*AccountCommandType
    literals []enum.Literal
}

var _accountCommandTypes = &accountCommandTypes{values: []*AccountCommandType{
    {name: "EnableAccount", ordinal: 0},
    {name: "DisableAccount", ordinal: 1},
    {name: "CreateAccount", ordinal: 2},
    {name: "DeleteAccount", ordinal: 3},
    {name: "UpdateAccount", ordinal: 4}},
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

func (o *accountCommandTypes) EnableAccount() *AccountCommandType {
    return _accountCommandTypes.values[0]
}

func (o *accountCommandTypes) DisableAccount() *AccountCommandType {
    return _accountCommandTypes.values[1]
}

func (o *accountCommandTypes) CreateAccount() *AccountCommandType {
    return _accountCommandTypes.values[2]
}

func (o *accountCommandTypes) DeleteAccount() *AccountCommandType {
    return _accountCommandTypes.values[3]
}

func (o *accountCommandTypes) UpdateAccount() *AccountCommandType {
    return _accountCommandTypes.values[4]
}

func (o *accountCommandTypes) ParseAccountCommandType(name string) (ret *AccountCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*AccountCommandType), ok
	}
	return
}



