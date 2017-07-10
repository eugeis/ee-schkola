package auth

import (
    "ee/schkola"
    "ee/schkola/person"
    "time"
)
type AccountCommands struct {
	name  string
	ordinal int
}

func (o *AccountCommands) Name() string {
    return o.name
}

func (o *AccountCommands) Ordinal() int {
    return o.ordinal
}

func (o *AccountCommands) IsEnableAccount() bool {
    return o == _accountCommandss.EnableAccount()
}

func (o *AccountCommands) IsDisableAccount() bool {
    return o == _accountCommandss.DisableAccount()
}

func (o *AccountCommands) IsRegisterAccount() bool {
    return o == _accountCommandss.RegisterAccount()
}

type accountCommandss struct {
	values []*AccountCommands
}

var _accountCommandss = &accountCommandss{values: []*AccountCommands{
    {name: "enableAccount", ordinal: 0},
    {name: "disableAccount", ordinal: 1},
    {name: "registerAccount", ordinal: 2}},
}

func AccountCommandss() *accountCommandss {
	return _accountCommandss
}

func (o *accountCommandss) Values() []*AccountCommands {
	return o.values
}

func (o *accountCommandss) EnableAccount() *AccountCommands {
    return _accountCommandss.values[0]
}

func (o *accountCommandss) DisableAccount() *AccountCommands {
    return _accountCommandss.values[1]
}

func (o *accountCommandss) RegisterAccount() *AccountCommands {
    return _accountCommandss.values[2]
}

func (o *accountCommandss) ParseAccountCommands(name string) (ret *AccountCommands, ok bool) {
    switch name {
      case "EnableAccount":
        ret = o.EnableAccount()
      case "DisableAccount":
        ret = o.DisableAccount()
      case "RegisterAccount":
        ret = o.RegisterAccount()
    }
    return
}




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
        SchkolaBase : SchkolaBase,
    }
    return
}



