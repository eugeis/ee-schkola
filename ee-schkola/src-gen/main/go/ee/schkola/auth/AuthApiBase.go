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
    Trace  *schkola.Trace
}

func NewAccount(@@EMPTY@@ string, username string, password string, email string, disabled bool, lastLoginAt *time.Time, 
                profile *person.Profile) (ret Account, err error) {
    ret = Account{
        @@EMPTY@@ : @@EMPTY@@,
        Username : username,
        Password : password,
        Email : email,
        Disabled : disabled,
        LastLoginAt : lastLoginAt,
        Profile : profile,
    }
    return
}



