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
    Id  
}


type ExpenseUpdated struct {
    Purpose  *ExpensePurpose
    Amount  float64
    Profile  *person.Profile
    Date  *time.Time
    *schkola.SchkolaBase
}


type ExpensePurposeCreated struct {
    Name  
    Description  
    *schkola.SchkolaBase
}


type ExpensePurposeDeleted struct {
    Id  
}


type ExpensePurposeUpdated struct {
    Name  
    Description  
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
    Id  
}


type FeeUpdated struct {
    Student  *person.Profile
    Amount  float64
    Kind  *FeeKind
    Date  *time.Time
    *schkola.SchkolaBase
}


type FeeKindCreated struct {
    Name  
    Amount  float64
    Description  
    *schkola.SchkolaBase
}


type FeeKindDeleted struct {
    Id  
}


type FeeKindUpdated struct {
    Name  
    Amount  float64
    Description  
    *schkola.SchkolaBase
}





