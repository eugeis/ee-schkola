package auth

import (
    "ee/schkola/person"
    "ee/schkola/shared"
    "github.com/looplab/eventhorizon"
)
        
type Account struct {
    Name *shared.PersonName `json:"name" eh:"optional"`
    Username string `json:"username" eh:"optional"`
    Password string `json:"password" eh:"optional"`
    Email string `json:"email" eh:"optional"`
    Disabled bool `json:"disabled" eh:"optional"`
    Roles []string `json:"roles" eh:"optional"`
    Profile *person.Profile `json:"profile" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
    *AccountConfirmation
}

func New@@EMPTY@@() (ret *Account) {
    accountConfirmation := New@@EMPTY@@()
    ret = &Account{
        AccountConfirmation: accountConfirmation,
    }
    return
}

func (o *Account) AddToRoles(item string) string {
    o.Roles = append(o.Roles, item)
    return item
}
func (o *Account) EntityID() eventhorizon.UUID { return o.Id }





type UserCredentials struct {
    Username string `json:"username" eh:"optional"`
    Password string `json:"password" eh:"optional"`
}

func New@@EMPTY@@() (ret *UserCredentials) {
    ret = &UserCredentials{}
    return
}







