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
     SendEnabledConfirmationCommand eventhorizon.CommandType = "SendEnabledConfirmation"
     SendDisabledConfirmationCommand eventhorizon.CommandType = "SendDisabledConfirmation"
     LoginCommand eventhorizon.CommandType = "Login"
     SendCreatedConfirmationCommand eventhorizon.CommandType = "SendCreatedConfirmation"
     CreateCommand eventhorizon.CommandType = "Create"
     DeleteCommand eventhorizon.CommandType = "Delete"
     DisableCommand eventhorizon.CommandType = "Disable"
     EnableCommand eventhorizon.CommandType = "Enable"
     UpdateCommand eventhorizon.CommandType = "Update"
)




        
type SendEnabledConfirmation struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *SendEnabledConfirmation) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *SendEnabledConfirmation) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *SendEnabledConfirmation) CommandType() eventhorizon.CommandType      { return SendEnabledConfirmationCommand }



        
type SendDisabledConfirmation struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *SendDisabledConfirmation) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *SendDisabledConfirmation) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *SendDisabledConfirmation) CommandType() eventhorizon.CommandType      { return SendDisabledConfirmationCommand }



        
type Login struct {
    Username string `json:"username" eh:"optional"`
    Email string `json:"email" eh:"optional"`
    Password string `json:"password" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Login) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Login) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *Login) CommandType() eventhorizon.CommandType      { return LoginCommand }



        
type SendCreatedConfirmation struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *SendCreatedConfirmation) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *SendCreatedConfirmation) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *SendCreatedConfirmation) CommandType() eventhorizon.CommandType      { return SendCreatedConfirmationCommand }



        
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
func (o *Create) CommandType() eventhorizon.CommandType      { return CreateCommand }



        
type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Delete) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Delete) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *Delete) CommandType() eventhorizon.CommandType      { return DeleteCommand }



        
type Disable struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Disable) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Disable) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *Disable) CommandType() eventhorizon.CommandType      { return DisableCommand }



        
type Enable struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Enable) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Enable) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *Enable) CommandType() eventhorizon.CommandType      { return EnableCommand }



        
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
func (o *Update) CommandType() eventhorizon.CommandType      { return UpdateCommand }





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

func (o *AccountCommandType) IsSendEnabledConfirmation() bool {
    return o == _accountCommandTypes.SendEnabledConfirmation()
}

func (o *AccountCommandType) IsSendDisabledConfirmation() bool {
    return o == _accountCommandTypes.SendDisabledConfirmation()
}

func (o *AccountCommandType) IsLogin() bool {
    return o == _accountCommandTypes.Login()
}

func (o *AccountCommandType) IsSendCreatedConfirmation() bool {
    return o == _accountCommandTypes.SendCreatedConfirmation()
}

func (o *AccountCommandType) IsCreate() bool {
    return o == _accountCommandTypes.Create()
}

func (o *AccountCommandType) IsDelete() bool {
    return o == _accountCommandTypes.Delete()
}

func (o *AccountCommandType) IsDisable() bool {
    return o == _accountCommandTypes.Disable()
}

func (o *AccountCommandType) IsEnable() bool {
    return o == _accountCommandTypes.Enable()
}

func (o *AccountCommandType) IsUpdate() bool {
    return o == _accountCommandTypes.Update()
}

type accountCommandTypes struct {
	values []*AccountCommandType
    literals []enum.Literal
}

var _accountCommandTypes = &accountCommandTypes{values: []*AccountCommandType{
    {name: "SendEnabledConfirmation", ordinal: 0},
    {name: "SendDisabledConfirmation", ordinal: 1},
    {name: "Login", ordinal: 2},
    {name: "SendCreatedConfirmation", ordinal: 3},
    {name: "Create", ordinal: 4},
    {name: "Delete", ordinal: 5},
    {name: "Disable", ordinal: 6},
    {name: "Enable", ordinal: 7},
    {name: "Update", ordinal: 8}},
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

func (o *accountCommandTypes) SendEnabledConfirmation() *AccountCommandType {
    return _accountCommandTypes.values[0]
}

func (o *accountCommandTypes) SendDisabledConfirmation() *AccountCommandType {
    return _accountCommandTypes.values[1]
}

func (o *accountCommandTypes) Login() *AccountCommandType {
    return _accountCommandTypes.values[2]
}

func (o *accountCommandTypes) SendCreatedConfirmation() *AccountCommandType {
    return _accountCommandTypes.values[3]
}

func (o *accountCommandTypes) Create() *AccountCommandType {
    return _accountCommandTypes.values[4]
}

func (o *accountCommandTypes) Delete() *AccountCommandType {
    return _accountCommandTypes.values[5]
}

func (o *accountCommandTypes) Disable() *AccountCommandType {
    return _accountCommandTypes.values[6]
}

func (o *accountCommandTypes) Enable() *AccountCommandType {
    return _accountCommandTypes.values[7]
}

func (o *accountCommandTypes) Update() *AccountCommandType {
    return _accountCommandTypes.values[8]
}

func (o *accountCommandTypes) ParseAccountCommandType(name string) (ret *AccountCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*AccountCommandType), ok
	}
	return
}



