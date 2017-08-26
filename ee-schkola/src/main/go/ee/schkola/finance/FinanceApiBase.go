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
    schkolaBase := schkola.NewSchkolaBase()
    ret = &Expense{
        SchkolaBase: schkolaBase,
    }
    return
}


type ExpensePurpose struct {
    Name string
    Description string
    *schkola.SchkolaBase
}

func NewExpensePurpose() (ret *ExpensePurpose) {
    schkolaBase := schkola.NewSchkolaBase()
    ret = &ExpensePurpose{
        SchkolaBase: schkolaBase,
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
    schkolaBase := schkola.NewSchkolaBase()
    ret = &Fee{
        SchkolaBase: schkolaBase,
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
    schkolaBase := schkola.NewSchkolaBase()
    ret = &FeeKind{
        SchkolaBase: schkolaBase,
    }
    return
}









