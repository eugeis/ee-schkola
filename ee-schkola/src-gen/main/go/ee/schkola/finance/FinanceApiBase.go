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
    ret = &Expense{
        Purpose : purpose,
        Amount : amount,
        Profile : profile,
        Date : date,
        SchkolaBase : SchkolaBase,
    }
    return
}


type ExpensePurpose struct {
    Name  string
    Description  string
    *schkola.SchkolaBase
}

func NewExpensePurpose(name string, description string, SchkolaBase *schkola.SchkolaBase) (ret *ExpensePurpose, err error) {
    ret = &ExpensePurpose{
        Name : name,
        Description : description,
        SchkolaBase : SchkolaBase,
    }
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
    ret = &Fee{
        Student : student,
        Amount : amount,
        Kind : kind,
        Date : date,
        SchkolaBase : SchkolaBase,
    }
    return
}


type FeeKind struct {
    Name  string
    Amount  float64
    Description  string
    *schkola.SchkolaBase
}

func NewFeeKind(name string, amount float64, description string, SchkolaBase *schkola.SchkolaBase) (ret *FeeKind, err error) {
    ret = &FeeKind{
        Name : name,
        Amount : amount,
        Description : description,
        SchkolaBase : SchkolaBase,
    }
    return
}



