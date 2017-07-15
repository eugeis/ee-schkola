package auth

import (
    "ee/schkola"
    "ee/schkola/person"
    "time"
)

type Account struct {
    Username  string
    Password  string
    Email  string
    disabled bool
    lastLoginAt *time.Time
    profile *person.Profile
    *schkola.SchkolaBase
}

func NewAccount(username string, password string, email string, disabled bool, lastLoginAt *time.Time, profile *person.Profile, 
                SchkolaBase *schkola.SchkolaBase) (ret *Account) {
    ret = &Account{
        Username : username,
        Password : password,
        Email : email,
        disabled: disabled,
        lastLoginAt: lastLoginAt,
        profile: profile,
        SchkolaBase: SchkolaBase,
    }
    
    return
}









