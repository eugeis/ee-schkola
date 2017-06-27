package auth

import (
    "ee/schkola"
    "ee/schkola/person"
    "time"
)


type Login struct {
    Principal  string
    Password  string
    Person  *person.User
    Disabled  bool
    LastLoginAt  *time.Time
    Trace  *schkola.Trace
}

func NewLogin(principal string, password string, person *person.User, disabled bool, lastLoginAt *time.Time) (ret Login, err error) {
    ret = Login{
        Principal : principal,
        Password : password,
        Person : person,
        Disabled : disabled,
        LastLoginAt : lastLoginAt,
    }
    return
}



