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
    Disabled  bool
    LastLoginAt  *time.Time
    Profile  *person.Profile
    *schkola.SchkolaBase
}

func NewAccount(username string, password string, email string, disabled bool, lastLoginAt *time.Time, profile *person.Profile, 
                SchkolaBase *schkola.SchkolaBase) (ret *Account, err error) {
    ret = &Account{
        Username : username,
        Password : password,
        Email : email,
        Disabled : disabled,
        LastLoginAt : lastLoginAt,
        Profile : profile,
        SchkolaBase: SchkolaBase,
    }
    return
}









