package auth

import (
    "ee/schkola"
    "ee/schkola/person"
    "github.com/looplab/eventhorizon"
)
type Account struct {
    Name *schkola.PersonName `json:"name" eh:"optional"`
    Username string `json:"username" eh:"optional"`
    Password string `json:"password" eh:"optional"`
    Email string `json:"email" eh:"optional"`
    Disabled bool `json:"disabled" eh:"optional"`
    Roles []string `json:"roles" eh:"optional"`
    Profile *person.Profile `json:"profile" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func NewAccount() (ret *Account) {
    ret = &Account{}
    return
}

func (o *Account) AddToRoles(item string) string {
    o.Roles = append(o.Roles, item)
    return item
}




type UserCredentials struct {
    Username string `json:"username" eh:"optional"`
    Password string `json:"password" eh:"optional"`
}

func NewUserCredentials() (ret *UserCredentials) {
    ret = &UserCredentials{}
    return
}







