package finance

import (
    "ee/schkola"
    "ee/schkola/person"
    "time"
)

type Expense struct {
    purpose *ExpensePurpose
    amount float64
    profile *person.Profile
    date *time.Time
    *schkola.SchkolaBase
}

func NewExpense(purpose *ExpensePurpose, amount float64, profile *person.Profile, date *time.Time, SchkolaBase *schkola.SchkolaBase) (ret *Expense) {
    ret = &Expense{
        purpose: purpose,
        amount: amount,
        profile: profile,
        date: date,
        SchkolaBase: SchkolaBase,
    }
    
    return
}



type ExpensePurpose struct {
    name string
    description string
    *schkola.SchkolaBase
}

func NewExpensePurpose(name string, description string, SchkolaBase *schkola.SchkolaBase) (ret *ExpensePurpose) {
    ret = &ExpensePurpose{
        name: name,
        description: description,
        SchkolaBase: SchkolaBase,
    }
    
    return
}



type Fee struct {
    student *person.Profile
    amount float64
    kind *FeeKind
    date *time.Time
    *schkola.SchkolaBase
}

func NewFee(student *person.Profile, amount float64, kind *FeeKind, date *time.Time, SchkolaBase *schkola.SchkolaBase) (ret *Fee) {
    ret = &Fee{
        student: student,
        amount: amount,
        kind: kind,
        date: date,
        SchkolaBase: SchkolaBase,
    }
    
    return
}



type FeeKind struct {
    name string
    amount float64
    description string
    *schkola.SchkolaBase
}

func NewFeeKind(name string, amount float64, description string, SchkolaBase *schkola.SchkolaBase) (ret *FeeKind) {
    ret = &FeeKind{
        name: name,
        amount: amount,
        description: description,
        SchkolaBase: SchkolaBase,
    }
    
    return
}









