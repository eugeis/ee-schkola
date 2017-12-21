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
     SendEnabledAccountConfirmationEvent eventhorizon.EventType = "SendEnabledAccountConfirmation"
     SendDisabledAccountConfirmationEvent eventhorizon.EventType = "SendDisabledAccountConfirmation"
     AccountLoginEvent eventhorizon.EventType = "AccountLogin"
     SendCreatedAccountConfirmationEvent eventhorizon.EventType = "SendCreatedAccountConfirmation"
     AccountCreateEvent eventhorizon.EventType = "AccountCreate"
     AccountCreatedEvent eventhorizon.EventType = "AccountCreated"
     AccountDeleteEvent eventhorizon.EventType = "AccountDelete"
     AccountDeletedEvent eventhorizon.EventType = "AccountDeleted"
     AccountLoggedEvent eventhorizon.EventType = "AccountLogged"
     SendCreatedAccountConfirmationedEvent eventhorizon.EventType = "SendCreatedAccountConfirmationed"
     SendEnabledAccountConfirmationedEvent eventhorizon.EventType = "SendEnabledAccountConfirmationed"
     SendDisabledAccountConfirmationedEvent eventhorizon.EventType = "SendDisabledAccountConfirmationed"
     AccountDisableEvent eventhorizon.EventType = "AccountDisable"
     AccountEnableEvent eventhorizon.EventType = "AccountEnable"
     AccountUpdateEvent eventhorizon.EventType = "AccountUpdate"
     AccountUpdatedEvent eventhorizon.EventType = "AccountUpdated"
     AccountEnabledEvent eventhorizon.EventType = "AccountEnabled"
     AccountDisabledEvent eventhorizon.EventType = "AccountDisabled"
)




type SendAccountEnabledConfirmation struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type SendAccountDisabledConfirmation struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type LoginAccount struct {
    Username string `json:"username" eh:"optional"`
    Email string `json:"email" eh:"optional"`
    Password string `json:"password" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type SendAccountCreatedConfirmation struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type CreateAccount struct {
    Name *shared.PersonName `json:"name" eh:"optional"`
    Username string `json:"username" eh:"optional"`
    Password string `json:"password" eh:"optional"`
    Email string `json:"email" eh:"optional"`
    Disabled bool `json:"disabled" eh:"optional"`
    Roles []string `json:"roles" eh:"optional"`
    Profile *person.Profile `json:"profile" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *CreateAccount) AddToRoles(item string) string {
    o.Roles = append(o.Roles, item)
    return item
}


type AccountCreated struct {
    Name *shared.PersonName `json:"name" eh:"optional"`
    Username string `json:"username" eh:"optional"`
    Password string `json:"password" eh:"optional"`
    Email string `json:"email" eh:"optional"`
    Disabled bool `json:"disabled" eh:"optional"`
    Roles []string `json:"roles" eh:"optional"`
    Profile *person.Profile `json:"profile" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *AccountCreated) AddToRoles(item string) string {
    o.Roles = append(o.Roles, item)
    return item
}


type DeleteAccount struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type AccountDeleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type AccountLogged struct {
    Username string `json:"username" eh:"optional"`
    Email string `json:"email" eh:"optional"`
    Password string `json:"password" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type SendCreatedAccountConfirmationed struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type SendEnabledAccountConfirmationed struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type SendDisabledAccountConfirmationed struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type DisableAccount struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type EnableAccount struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type UpdateAccount struct {
    Name *shared.PersonName `json:"name" eh:"optional"`
    Username string `json:"username" eh:"optional"`
    Password string `json:"password" eh:"optional"`
    Email string `json:"email" eh:"optional"`
    Disabled bool `json:"disabled" eh:"optional"`
    Roles []string `json:"roles" eh:"optional"`
    Profile *person.Profile `json:"profile" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *UpdateAccount) AddToRoles(item string) string {
    o.Roles = append(o.Roles, item)
    return item
}


type AccountUpdated struct {
    Name *shared.PersonName `json:"name" eh:"optional"`
    Username string `json:"username" eh:"optional"`
    Password string `json:"password" eh:"optional"`
    Email string `json:"email" eh:"optional"`
    Disabled bool `json:"disabled" eh:"optional"`
    Roles []string `json:"roles" eh:"optional"`
    Profile *person.Profile `json:"profile" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *AccountUpdated) AddToRoles(item string) string {
    o.Roles = append(o.Roles, item)
    return item
}


type AccountEnabled struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type AccountDisabled struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
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

func (o AccountEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *AccountEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := AccountEventTypes().ParseAccountEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid AccountEventType %q", lit.Name)
        }
	}
	return
}

func (o AccountEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *AccountEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := AccountEventTypes().ParseAccountEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid AccountEventType %q", lit)
        }
    }
    return
}

func (o *AccountEventType) IsSendEnabledAccountConfirmation() bool {
    return o == _accountEventTypes.SendEnabledAccountConfirmation()
}

func (o *AccountEventType) IsSendDisabledAccountConfirmation() bool {
    return o == _accountEventTypes.SendDisabledAccountConfirmation()
}

func (o *AccountEventType) IsAccountLogin() bool {
    return o == _accountEventTypes.AccountLogin()
}

func (o *AccountEventType) IsSendCreatedAccountConfirmation() bool {
    return o == _accountEventTypes.SendCreatedAccountConfirmation()
}

func (o *AccountEventType) IsAccountCreate() bool {
    return o == _accountEventTypes.AccountCreate()
}

func (o *AccountEventType) IsAccountCreated() bool {
    return o == _accountEventTypes.AccountCreated()
}

func (o *AccountEventType) IsAccountDelete() bool {
    return o == _accountEventTypes.AccountDelete()
}

func (o *AccountEventType) IsAccountDeleted() bool {
    return o == _accountEventTypes.AccountDeleted()
}

func (o *AccountEventType) IsAccountLogged() bool {
    return o == _accountEventTypes.AccountLogged()
}

func (o *AccountEventType) IsSendCreatedAccountConfirmationed() bool {
    return o == _accountEventTypes.SendCreatedAccountConfirmationed()
}

func (o *AccountEventType) IsSendEnabledAccountConfirmationed() bool {
    return o == _accountEventTypes.SendEnabledAccountConfirmationed()
}

func (o *AccountEventType) IsSendDisabledAccountConfirmationed() bool {
    return o == _accountEventTypes.SendDisabledAccountConfirmationed()
}

func (o *AccountEventType) IsAccountDisable() bool {
    return o == _accountEventTypes.AccountDisable()
}

func (o *AccountEventType) IsAccountEnable() bool {
    return o == _accountEventTypes.AccountEnable()
}

func (o *AccountEventType) IsAccountUpdate() bool {
    return o == _accountEventTypes.AccountUpdate()
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
    {name: "SendEnabledAccountConfirmation", ordinal: 0},
    {name: "SendDisabledAccountConfirmation", ordinal: 1},
    {name: "AccountLogin", ordinal: 2},
    {name: "SendCreatedAccountConfirmation", ordinal: 3},
    {name: "AccountCreate", ordinal: 4},
    {name: "AccountCreated", ordinal: 5},
    {name: "AccountDelete", ordinal: 6},
    {name: "AccountDeleted", ordinal: 7},
    {name: "AccountLogged", ordinal: 8},
    {name: "SendCreatedAccountConfirmationed", ordinal: 9},
    {name: "SendEnabledAccountConfirmationed", ordinal: 10},
    {name: "SendDisabledAccountConfirmationed", ordinal: 11},
    {name: "AccountDisable", ordinal: 12},
    {name: "AccountEnable", ordinal: 13},
    {name: "AccountUpdate", ordinal: 14},
    {name: "AccountUpdated", ordinal: 15},
    {name: "AccountEnabled", ordinal: 16},
    {name: "AccountDisabled", ordinal: 17}},
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

func (o *accountEventTypes) SendEnabledAccountConfirmation() *AccountEventType {
    return _accountEventTypes.values[0]
}

func (o *accountEventTypes) SendDisabledAccountConfirmation() *AccountEventType {
    return _accountEventTypes.values[1]
}

func (o *accountEventTypes) AccountLogin() *AccountEventType {
    return _accountEventTypes.values[2]
}

func (o *accountEventTypes) SendCreatedAccountConfirmation() *AccountEventType {
    return _accountEventTypes.values[3]
}

func (o *accountEventTypes) AccountCreate() *AccountEventType {
    return _accountEventTypes.values[4]
}

func (o *accountEventTypes) AccountCreated() *AccountEventType {
    return _accountEventTypes.values[5]
}

func (o *accountEventTypes) AccountDelete() *AccountEventType {
    return _accountEventTypes.values[6]
}

func (o *accountEventTypes) AccountDeleted() *AccountEventType {
    return _accountEventTypes.values[7]
}

func (o *accountEventTypes) AccountLogged() *AccountEventType {
    return _accountEventTypes.values[8]
}

func (o *accountEventTypes) SendCreatedAccountConfirmationed() *AccountEventType {
    return _accountEventTypes.values[9]
}

func (o *accountEventTypes) SendEnabledAccountConfirmationed() *AccountEventType {
    return _accountEventTypes.values[10]
}

func (o *accountEventTypes) SendDisabledAccountConfirmationed() *AccountEventType {
    return _accountEventTypes.values[11]
}

func (o *accountEventTypes) AccountDisable() *AccountEventType {
    return _accountEventTypes.values[12]
}

func (o *accountEventTypes) AccountEnable() *AccountEventType {
    return _accountEventTypes.values[13]
}

func (o *accountEventTypes) AccountUpdate() *AccountEventType {
    return _accountEventTypes.values[14]
}

func (o *accountEventTypes) AccountUpdated() *AccountEventType {
    return _accountEventTypes.values[15]
}

func (o *accountEventTypes) AccountEnabled() *AccountEventType {
    return _accountEventTypes.values[16]
}

func (o *accountEventTypes) AccountDisabled() *AccountEventType {
    return _accountEventTypes.values[17]
}

func (o *accountEventTypes) ParseAccountEventType(name string) (ret *AccountEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*AccountEventType), ok
	}
	return
}



