package finance

import (
    "ee/schkola/person"
    "github.com/looplab/eventhorizon"
    "time"
)
type Expense struct {
    Purpose *ExpensePurpose `json:"purpose" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Profile *person.Profile `json:"profile" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func NewExpense() (ret *Expense) {
    ret = &Expense{}
    return
}


type ExpensePurpose struct {
    Name string `json:"name" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func NewExpensePurpose() (ret *ExpensePurpose) {
    ret = &ExpensePurpose{}
    return
}


type Fee struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Kind *FeeKind `json:"kind" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func NewFee() (ret *Fee) {
    ret = &Fee{}
    return
}


type FeeKind struct {
    Name string `json:"name" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func NewFeeKind() (ret *FeeKind) {
    ret = &FeeKind{}
    return
}









