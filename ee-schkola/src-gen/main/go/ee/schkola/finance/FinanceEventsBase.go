package finance

import (
    "ee/schkola"
    "ee/schkola/person"
    "time"
)
type ExpenseCreated struct {
    Purpose  *ExpensePurpose
    Amount  float64
    Profile  *person.Profile
    Date  *time.Time
    *schkola.SchkolaBase
}


type ExpenseDeleted struct {
    Id  string
}


type ExpenseUpdated struct {
    Purpose  *ExpensePurpose
    Amount  float64
    Profile  *person.Profile
    Date  *time.Time
    *schkola.SchkolaBase
}


type ExpensePurposeCreated struct {
    Name  string
    Description  string
    *schkola.SchkolaBase
}


type ExpensePurposeDeleted struct {
    Id  string
}


type ExpensePurposeUpdated struct {
    Name  string
    Description  string
    *schkola.SchkolaBase
}


type FeeCreated struct {
    Student  *person.Profile
    Amount  float64
    Kind  *FeeKind
    Date  *time.Time
    *schkola.SchkolaBase
}


type FeeDeleted struct {
    Id  string
}


type FeeUpdated struct {
    Student  *person.Profile
    Amount  float64
    Kind  *FeeKind
    Date  *time.Time
    *schkola.SchkolaBase
}


type FeeKindCreated struct {
    Name  string
    Amount  float64
    Description  string
    *schkola.SchkolaBase
}


type FeeKindDeleted struct {
    Id  string
}


type FeeKindUpdated struct {
    Name  string
    Amount  float64
    Description  string
    *schkola.SchkolaBase
}





