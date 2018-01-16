package auth

import (
    "ee/schkola/person"
    "ee/schkola/shared"
    "encoding/json"
    "fmt"
    "github.com/eugeis/gee/enum"
    "github.com/looplab/eventhorizon"
    "gopkg.in/mgo.v2/bson"
)
const (
     SendAccountEnabledConfirmationCommand eventhorizon.CommandType = "SendAccountEnabledConfirmation"
     SendAccountDisabledConfirmationCommand eventhorizon.CommandType = "SendAccountDisabledConfirmation"
     LoginAccountCommand eventhorizon.CommandType = "LoginAccount"
     SendAccountCreatedConfirmationCommand eventhorizon.CommandType = "SendAccountCreatedConfirmation"
     CreateAccountCommand eventhorizon.CommandType = "CreateAccount"
     DeleteAccountCommand eventhorizon.CommandType = "DeleteAccount"
     DisableAccountCommand eventhorizon.CommandType = "DisableAccount"
     EnableAccountCommand eventhorizon.CommandType = "EnableAccount"
     UpdateAccountCommand eventhorizon.CommandType = "UpdateAccount"
)




        
type SendEnabledConfirmation struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *SendEnabledConfirmation) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *SendEnabledConfirmation) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *SendEnabledConfirmation) CommandType() eventhorizon.CommandType      { return SendAccountEnabledConfirmationCommand }



        
type SendDisabledConfirmation struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *SendDisabledConfirmation) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *SendDisabledConfirmation) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *SendDisabledConfirmation) CommandType() eventhorizon.CommandType      { return SendAccountDisabledConfirmationCommand }



        
type Login struct {
    Username string `json:"username" eh:"optional"`
    Email string `json:"email" eh:"optional"`
    Password string `json:"password" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Login) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Login) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *Login) CommandType() eventhorizon.CommandType      { return LoginAccountCommand }



        
type SendCreatedConfirmation struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *SendCreatedConfirmation) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *SendCreatedConfirmation) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *SendCreatedConfirmation) CommandType() eventhorizon.CommandType      { return SendAccountCreatedConfirmationCommand }



        
type Create struct {
    Name *shared.PersonName `json:"name" eh:"optional"`
    Username string `json:"username" eh:"optional"`
    Password string `json:"password" eh:"optional"`
    Email string `json:"email" eh:"optional"`
    Disabled bool `json:"disabled" eh:"optional"`
    Roles []string `json:"roles" eh:"optional"`
    Profile *person.Profile `json:"profile" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *Create) AddToRoles(item string) string {
    o.Roles = append(o.Roles, item)
    return item
}
func (o *Create) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Create) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *Create) CommandType() eventhorizon.CommandType      { return CreateAccountCommand }



        
type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Delete) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Delete) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *Delete) CommandType() eventhorizon.CommandType      { return DeleteAccountCommand }



        
type Disable struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Disable) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Disable) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *Disable) CommandType() eventhorizon.CommandType      { return DisableAccountCommand }



        
type Enable struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Enable) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Enable) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *Enable) CommandType() eventhorizon.CommandType      { return EnableAccountCommand }



        
type Update struct {
    Name *shared.PersonName `json:"name" eh:"optional"`
    Username string `json:"username" eh:"optional"`
    Password string `json:"password" eh:"optional"`
    Email string `json:"email" eh:"optional"`
    Disabled bool `json:"disabled" eh:"optional"`
    Roles []string `json:"roles" eh:"optional"`
    Profile *person.Profile `json:"profile" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *Update) AddToRoles(item string) string {
    o.Roles = append(o.Roles, item)
    return item
}
func (o *Update) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Update) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *Update) CommandType() eventhorizon.CommandType      { return UpdateAccountCommand }





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

func (o AccountCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *AccountCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := AccountCommandTypes().ParseAccountCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid AccountCommandType %q", lit.Name)
        }
	}
	return
}

func (o AccountCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *AccountCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := AccountCommandTypes().ParseAccountCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid AccountCommandType %q", lit)
        }
    }
    return
}

func (o *AccountCommandType) IsSendAccountEnabledConfirmation() bool {
    return o == _accountCommandTypes.SendAccountEnabledConfirmation()
}

func (o *AccountCommandType) IsSendAccountDisabledConfirmation() bool {
    return o == _accountCommandTypes.SendAccountDisabledConfirmation()
}

func (o *AccountCommandType) IsLoginAccount() bool {
    return o == _accountCommandTypes.LoginAccount()
}

func (o *AccountCommandType) IsSendAccountCreatedConfirmation() bool {
    return o == _accountCommandTypes.SendAccountCreatedConfirmation()
}

func (o *AccountCommandType) IsCreateAccount() bool {
    return o == _accountCommandTypes.CreateAccount()
}

func (o *AccountCommandType) IsDeleteAccount() bool {
    return o == _accountCommandTypes.DeleteAccount()
}

func (o *AccountCommandType) IsDisableAccount() bool {
    return o == _accountCommandTypes.DisableAccount()
}

func (o *AccountCommandType) IsEnableAccount() bool {
    return o == _accountCommandTypes.EnableAccount()
}

func (o *AccountCommandType) IsUpdateAccount() bool {
    return o == _accountCommandTypes.UpdateAccount()
}

type accountCommandTypes struct {
	values []*AccountCommandType
    literals []enum.Literal
}

var _accountCommandTypes = &accountCommandTypes{values: []*AccountCommandType{
    {name: "SendAccountEnabledConfirmation", ordinal: 0},
    {name: "SendAccountDisabledConfirmation", ordinal: 1},
    {name: "LoginAccount", ordinal: 2},
    {name: "SendAccountCreatedConfirmation", ordinal: 3},
    {name: "CreateAccount", ordinal: 4},
    {name: "DeleteAccount", ordinal: 5},
    {name: "DisableAccount", ordinal: 6},
    {name: "EnableAccount", ordinal: 7},
    {name: "UpdateAccount", ordinal: 8}},
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

func (o *accountCommandTypes) SendAccountEnabledConfirmation() *AccountCommandType {
    return _accountCommandTypes.values[0]
}

func (o *accountCommandTypes) SendAccountDisabledConfirmation() *AccountCommandType {
    return _accountCommandTypes.values[1]
}

func (o *accountCommandTypes) LoginAccount() *AccountCommandType {
    return _accountCommandTypes.values[2]
}

func (o *accountCommandTypes) SendAccountCreatedConfirmation() *AccountCommandType {
    return _accountCommandTypes.values[3]
}

func (o *accountCommandTypes) CreateAccount() *AccountCommandType {
    return _accountCommandTypes.values[4]
}

func (o *accountCommandTypes) DeleteAccount() *AccountCommandType {
    return _accountCommandTypes.values[5]
}

func (o *accountCommandTypes) DisableAccount() *AccountCommandType {
    return _accountCommandTypes.values[6]
}

func (o *accountCommandTypes) EnableAccount() *AccountCommandType {
    return _accountCommandTypes.values[7]
}

func (o *accountCommandTypes) UpdateAccount() *AccountCommandType {
    return _accountCommandTypes.values[8]
}

func (o *accountCommandTypes) ParseAccountCommandType(name string) (ret *AccountCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*AccountCommandType), ok
	}
	return
}



