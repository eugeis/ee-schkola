package auth

import (
    "ee/schkola"
    "ee/schkola/person"
    "github.com/looplab/eventhorizon"
    "time"
)
type Account struct {
    Name *schkola.PersonName
    Username string`eh:"optional"`
    Password string`eh:"optional"`
    Email string`eh:"optional"`
    Disabled bool
    LastLoginAt *time.Time
    Profile *person.Profile
    Id eventhorizon.UUID
}

func NewAccount() (ret *Account) {
    ret = &Account{}
    return
}









