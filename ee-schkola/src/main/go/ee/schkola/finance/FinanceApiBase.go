package finance

import (
    "ee/schkola"
    "ee/schkola/person"
    "time"
)
type Expense struct {
    Purpose *ExpensePurpose
    Amount float64
    Profile *person.Profile
    Date *time.Time
    *schkola.SchkolaBase
}

func NewExpense() (ret *Expense) {
    SchkolaBase:= schkola.NewSchkolaBase()
    ret = &Expense{
        SchkolaBase: SchkolaBase,
    }
    return
}


type ExpensePurpose struct {
    Name string
    Description string
    *schkola.SchkolaBase
}

func NewExpensePurpose() (ret *ExpensePurpose) {
    SchkolaBase:= schkola.NewSchkolaBase()
    ret = &ExpensePurpose{
        SchkolaBase: SchkolaBase,
    }
    return
}


type Fee struct {
    Student *person.Profile
    Amount float64
    Kind *FeeKind
    Date *time.Time
    *schkola.SchkolaBase
}

func NewFee() (ret *Fee) {
    SchkolaBase:= schkola.NewSchkolaBase()
    ret = &Fee{
        SchkolaBase: SchkolaBase,
    }
    return
}


type FeeKind struct {
    Name string
    Amount float64
    Description string
    *schkola.SchkolaBase
}

func NewFeeKind() (ret *FeeKind) {
    SchkolaBase:= schkola.NewSchkolaBase()
    ret = &FeeKind{
        SchkolaBase: SchkolaBase,
    }
    return
}









