package auth

import (
    "ee/schkola"
    "ee/schkola/person"
    "time"
)
type Account struct {
    Username string
    Password string
    Email string
    Disabled bool
    LastLoginAt *time.Time
    Profile *person.Profile
    *schkola.SchkolaBase
}

func NewAccount() (ret *Account) {
    ret = &Account{
        SchkolaBase: schkola.NewSchkolaBase(),
    }
    return
}









