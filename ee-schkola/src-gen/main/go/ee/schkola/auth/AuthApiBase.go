package auth

import (
    "ee/schkola"
    "ee/schkola/person"
    "time"
)
type Account struct {
    Username  
    Password  
    Email  
    Disabled  bool
    LastLoginAt  *time.Time
    Profile  *person.Profile
    *schkola.SchkolaBase
}

func NewAccount(username , password , email , disabled bool, lastLoginAt *time.Time, profile *person.Profile, 
                SchkolaBase *schkola.SchkolaBase) (ret *Account, err error) {
    ret = (id)
        
    ret.Username = username
    ret.Password = password
    ret.Email = email
    ret.Disabled = disabled
    ret.LastLoginAt = lastLoginAt
    ret.Profile = profile
    ret.SchkolaBase = SchkolaBase
    return
}









