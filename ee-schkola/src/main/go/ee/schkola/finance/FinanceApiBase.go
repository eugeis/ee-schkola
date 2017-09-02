package finance

import (
    "ee/schkola/person"
    "github.com/looplab/eventhorizon"
    "time"
)
type Expense struct {
    Purpose *ExpensePurpose
    Amount float64
    Profile *person.Profile
    Date *time.Time
    Id eventhorizon.UUID
}

func NewExpense() (ret *Expense) {
    ret = &Expense{}
    return
}


type ExpensePurpose struct {
    Name string
    Description string
    Id eventhorizon.UUID
}

func NewExpensePurpose() (ret *ExpensePurpose) {
    ret = &ExpensePurpose{}
    return
}


type Fee struct {
    Student *person.Profile
    Amount float64
    Kind *FeeKind
    Date *time.Time
    Id eventhorizon.UUID
}

func NewFee() (ret *Fee) {
    ret = &Fee{}
    return
}


type FeeKind struct {
    Name string
    Amount float64
    Description string
    Id eventhorizon.UUID
}

func NewFeeKind() (ret *FeeKind) {
    ret = &FeeKind{}
    return
}









