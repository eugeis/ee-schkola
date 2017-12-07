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
     SendEnabledConfirmationEvent eventhorizon.EventType = "SendEnabledConfirmation"
     SendDisabledConfirmationEvent eventhorizon.EventType = "SendDisabledConfirmation"
     LoginEvent eventhorizon.EventType = "Login"
     SendCreatedConfirmationEvent eventhorizon.EventType = "SendCreatedConfirmation"
     CreateEvent eventhorizon.EventType = "Create"
     CreatedEvent eventhorizon.EventType = "Created"
     DeleteEvent eventhorizon.EventType = "Delete"
     DeletedEvent eventhorizon.EventType = "Deleted"
     LoggedEvent eventhorizon.EventType = "Logged"
     SendCreatedConfirmationedEvent eventhorizon.EventType = "SendCreatedConfirmationed"
     SendEnabledConfirmationedEvent eventhorizon.EventType = "SendEnabledConfirmationed"
     SendDisabledConfirmationedEvent eventhorizon.EventType = "SendDisabledConfirmationed"
     DisableEvent eventhorizon.EventType = "Disable"
     EnableEvent eventhorizon.EventType = "Enable"
     UpdateEvent eventhorizon.EventType = "Update"
     UpdatedEvent eventhorizon.EventType = "Updated"
     EnabledEvent eventhorizon.EventType = "Enabled"
     DisabledEvent eventhorizon.EventType = "Disabled"
)




type SendEnabledConfirmation struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type SendDisabledConfirmation struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Login struct {
    Username string `json:"username" eh:"optional"`
    Email string `json:"email" eh:"optional"`
    Password string `json:"password" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type SendCreatedConfirmation struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


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


type Created struct {
    Name *shared.PersonName `json:"name" eh:"optional"`
    Username string `json:"username" eh:"optional"`
    Password string `json:"password" eh:"optional"`
    Email string `json:"email" eh:"optional"`
    Disabled bool `json:"disabled" eh:"optional"`
    Roles []string `json:"roles" eh:"optional"`
    Profile *person.Profile `json:"profile" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *Created) AddToRoles(item string) string {
    o.Roles = append(o.Roles, item)
    return item
}


type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Deleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Logged struct {
    Username string `json:"username" eh:"optional"`
    Email string `json:"email" eh:"optional"`
    Password string `json:"password" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type SendCreatedConfirmationed struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type SendEnabledConfirmationed struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type SendDisabledConfirmationed struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Disable struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Enable struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


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


type Updated struct {
    Name *shared.PersonName `json:"name" eh:"optional"`
    Username string `json:"username" eh:"optional"`
    Password string `json:"password" eh:"optional"`
    Email string `json:"email" eh:"optional"`
    Disabled bool `json:"disabled" eh:"optional"`
    Roles []string `json:"roles" eh:"optional"`
    Profile *person.Profile `json:"profile" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *Updated) AddToRoles(item string) string {
    o.Roles = append(o.Roles, item)
    return item
}


type Enabled struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Disabled struct {
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

func (o *AccountEventType) IsSendEnabledConfirmation() bool {
    return o == _accountEventTypes.SendEnabledConfirmation()
}

func (o *AccountEventType) IsSendDisabledConfirmation() bool {
    return o == _accountEventTypes.SendDisabledConfirmation()
}

func (o *AccountEventType) IsLogin() bool {
    return o == _accountEventTypes.Login()
}

func (o *AccountEventType) IsSendCreatedConfirmation() bool {
    return o == _accountEventTypes.SendCreatedConfirmation()
}

func (o *AccountEventType) IsCreate() bool {
    return o == _accountEventTypes.Create()
}

func (o *AccountEventType) IsCreated() bool {
    return o == _accountEventTypes.Created()
}

func (o *AccountEventType) IsDelete() bool {
    return o == _accountEventTypes.Delete()
}

func (o *AccountEventType) IsDeleted() bool {
    return o == _accountEventTypes.Deleted()
}

func (o *AccountEventType) IsLogged() bool {
    return o == _accountEventTypes.Logged()
}

func (o *AccountEventType) IsSendCreatedConfirmationed() bool {
    return o == _accountEventTypes.SendCreatedConfirmationed()
}

func (o *AccountEventType) IsSendEnabledConfirmationed() bool {
    return o == _accountEventTypes.SendEnabledConfirmationed()
}

func (o *AccountEventType) IsSendDisabledConfirmationed() bool {
    return o == _accountEventTypes.SendDisabledConfirmationed()
}

func (o *AccountEventType) IsDisable() bool {
    return o == _accountEventTypes.Disable()
}

func (o *AccountEventType) IsEnable() bool {
    return o == _accountEventTypes.Enable()
}

func (o *AccountEventType) IsUpdate() bool {
    return o == _accountEventTypes.Update()
}

func (o *AccountEventType) IsUpdated() bool {
    return o == _accountEventTypes.Updated()
}

func (o *AccountEventType) IsEnabled() bool {
    return o == _accountEventTypes.Enabled()
}

func (o *AccountEventType) IsDisabled() bool {
    return o == _accountEventTypes.Disabled()
}

type accountEventTypes struct {
	values []*AccountEventType
    literals []enum.Literal
}

var _accountEventTypes = &accountEventTypes{values: []*AccountEventType{
    {name: "SendEnabledConfirmation", ordinal: 0},
    {name: "SendDisabledConfirmation", ordinal: 1},
    {name: "Login", ordinal: 2},
    {name: "SendCreatedConfirmation", ordinal: 3},
    {name: "Create", ordinal: 4},
    {name: "Created", ordinal: 5},
    {name: "Delete", ordinal: 6},
    {name: "Deleted", ordinal: 7},
    {name: "Logged", ordinal: 8},
    {name: "SendCreatedConfirmationed", ordinal: 9},
    {name: "SendEnabledConfirmationed", ordinal: 10},
    {name: "SendDisabledConfirmationed", ordinal: 11},
    {name: "Disable", ordinal: 12},
    {name: "Enable", ordinal: 13},
    {name: "Update", ordinal: 14},
    {name: "Updated", ordinal: 15},
    {name: "Enabled", ordinal: 16},
    {name: "Disabled", ordinal: 17}},
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

func (o *accountEventTypes) SendEnabledConfirmation() *AccountEventType {
    return _accountEventTypes.values[0]
}

func (o *accountEventTypes) SendDisabledConfirmation() *AccountEventType {
    return _accountEventTypes.values[1]
}

func (o *accountEventTypes) Login() *AccountEventType {
    return _accountEventTypes.values[2]
}

func (o *accountEventTypes) SendCreatedConfirmation() *AccountEventType {
    return _accountEventTypes.values[3]
}

func (o *accountEventTypes) Create() *AccountEventType {
    return _accountEventTypes.values[4]
}

func (o *accountEventTypes) Created() *AccountEventType {
    return _accountEventTypes.values[5]
}

func (o *accountEventTypes) Delete() *AccountEventType {
    return _accountEventTypes.values[6]
}

func (o *accountEventTypes) Deleted() *AccountEventType {
    return _accountEventTypes.values[7]
}

func (o *accountEventTypes) Logged() *AccountEventType {
    return _accountEventTypes.values[8]
}

func (o *accountEventTypes) SendCreatedConfirmationed() *AccountEventType {
    return _accountEventTypes.values[9]
}

func (o *accountEventTypes) SendEnabledConfirmationed() *AccountEventType {
    return _accountEventTypes.values[10]
}

func (o *accountEventTypes) SendDisabledConfirmationed() *AccountEventType {
    return _accountEventTypes.values[11]
}

func (o *accountEventTypes) Disable() *AccountEventType {
    return _accountEventTypes.values[12]
}

func (o *accountEventTypes) Enable() *AccountEventType {
    return _accountEventTypes.values[13]
}

func (o *accountEventTypes) Update() *AccountEventType {
    return _accountEventTypes.values[14]
}

func (o *accountEventTypes) Updated() *AccountEventType {
    return _accountEventTypes.values[15]
}

func (o *accountEventTypes) Enabled() *AccountEventType {
    return _accountEventTypes.values[16]
}

func (o *accountEventTypes) Disabled() *AccountEventType {
    return _accountEventTypes.values[17]
}

func (o *accountEventTypes) ParseAccountEventType(name string) (ret *AccountEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*AccountEventType), ok
	}
	return
}



