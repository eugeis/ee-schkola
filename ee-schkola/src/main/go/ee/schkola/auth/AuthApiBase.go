package auth

import (
    "ee/schkola"
    "ee/schkola/person"
    "github.com/looplab/eventhorizon"
)
type Account struct {
    Name *schkola.PersonName
    Username string`eh:"optional"`
    Password string`eh:"optional"`
    Email string`eh:"optional"`
    Disabled bool
    Roles []string
    Profile *person.Profile
    Id eventhorizon.UUID
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
    Username string
    Password string
}

func NewUserCredentials() (ret *UserCredentials) {
    ret = &UserCredentials{}
    return
}







