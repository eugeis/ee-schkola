package auth

import (
    "ee/schkola"
    "ee/schkola/person"
    "time"
)

type Account struct {
    username string
    password string
    email string
    disabled bool
    lastLoginAt *time.Time
    profile *person.Profile
    *schkola.SchkolaBase
}

func NewAccount(username string, password string, email string, disabled bool, lastLoginAt *time.Time, profile *person.Profile, 
                SchkolaBase *schkola.SchkolaBase) (ret *Account, err error) {
    ret = &Account{
        username: username,
        password: password,
        email: email,
        disabled: disabled,
        lastLoginAt: lastLoginAt,
        profile: profile,
        SchkolaBase: SchkolaBase,
    }
    
    return
}









