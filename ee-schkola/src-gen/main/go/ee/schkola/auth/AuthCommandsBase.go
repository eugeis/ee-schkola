package auth


type EnableAccount struct {
}


type DisableAccount struct {
}


type RegisterAccount struct {
    Username  
    Email  
    Password  
}

func NewRegisterAccount(username , email , password ) (ret *RegisterAccount, err error) {
    ret = &RegisterAccount{
        Username : username,
        Email : email,
        Password : password,
    }
    return
}




type AccountCommandType struct {
	name  string
	ordinal int
}

func (o *AccountCommandType) Name() string {
    return o.name
}

func (o *AccountCommandType) Ordinal() int {
    return o.ordinal
}

func (o *AccountCommandType) IsAccountEnable() bool {
    return o == _accountCommandTypes.AccountEnable()
}

func (o *AccountCommandType) IsAccountDisable() bool {
    return o == _accountCommandTypes.AccountDisable()
}

func (o *AccountCommandType) IsAccountRegister() bool {
    return o == _accountCommandTypes.AccountRegister()
}

type accountCommandTypes struct {
	values []*AccountCommandType
}

var _accountCommandTypes = &accountCommandTypes{values: []*AccountCommandType{
    {name: "AccountEnable", ordinal: 0},
    {name: "AccountDisable", ordinal: 1},
    {name: "AccountRegister", ordinal: 2}},
}

func AccountCommandTypes() *accountCommandTypes {
	return _accountCommandTypes
}

func (o *accountCommandTypes) Values() []*AccountCommandType {
	return o.values
}

func (o *accountCommandTypes) AccountEnable() *AccountCommandType {
    return _accountCommandTypes.values[0]
}

func (o *accountCommandTypes) AccountDisable() *AccountCommandType {
    return _accountCommandTypes.values[1]
}

func (o *accountCommandTypes) AccountRegister() *AccountCommandType {
    return _accountCommandTypes.values[2]
}

func (o *accountCommandTypes) ParseAccountCommandType(name string) (ret *AccountCommandType, ok bool) {
    switch name {
      case "AccountEnable":
        ret = o.AccountEnable()
      case "AccountDisable":
        ret = o.AccountDisable()
      case "AccountRegister":
        ret = o.AccountRegister()
    }
    return
}



