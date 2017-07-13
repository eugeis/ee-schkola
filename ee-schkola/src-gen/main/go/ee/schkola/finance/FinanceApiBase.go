package finance

import (
    "ee/schkola"
    "ee/schkola/person"
    "time"
)
type Expense struct {
    Purpose  *ExpensePurpose
    Amount  float64
    Profile  *person.Profile
    Date  *time.Time
    *schkola.SchkolaBase
}

func NewExpense(purpose *ExpensePurpose, amount float64, profile *person.Profile, date *time.Time, SchkolaBase *schkola.SchkolaBase) (ret *Expense, err error) {
    ret = (id)
        
    ret.Purpose = purpose
    ret.Amount = amount
    ret.Profile = profile
    ret.Date = date
    ret.SchkolaBase = SchkolaBase
    return
}


type ExpensePurpose struct {
    Name  string
    Description  string
    *schkola.SchkolaBase
}

func NewExpensePurpose(name string, description string, SchkolaBase *schkola.SchkolaBase) (ret *ExpensePurpose, err error) {
    ret = (id)
        
    ret.Name = name
    ret.Description = description
    ret.SchkolaBase = SchkolaBase
    return
}


type Fee struct {
    Student  *person.Profile
    Amount  float64
    Kind  *FeeKind
    Date  *time.Time
    *schkola.SchkolaBase
}

func NewFee(student *person.Profile, amount float64, kind *FeeKind, date *time.Time, SchkolaBase *schkola.SchkolaBase) (ret *Fee, err error) {
    ret = (id)
        
    ret.Student = student
    ret.Amount = amount
    ret.Kind = kind
    ret.Date = date
    ret.SchkolaBase = SchkolaBase
    return
}


type FeeKind struct {
    Name  string
    Amount  float64
    Description  string
    *schkola.SchkolaBase
}

func NewFeeKind(name string, amount float64, description string, SchkolaBase *schkola.SchkolaBase) (ret *FeeKind, err error) {
    ret = (id)
        
    ret.Name = name
    ret.Amount = amount
    ret.Description = description
    ret.SchkolaBase = SchkolaBase
    return
}





