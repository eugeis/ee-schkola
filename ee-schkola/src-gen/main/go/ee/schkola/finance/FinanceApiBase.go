package finance

import (
    "ee/schkola/person"
    "time"
)


type Expense struct {
    Purpose  *ExpensePurpose
    Amount  float64
    Profile  *person.Profile
    Date  *time.Time
}

func NewExpense(@@EMPTY@@ string, purpose *ExpensePurpose, amount float64, profile *person.Profile, date *time.Time) (ret Expense, err error) {
    ret = Expense{
        @@EMPTY@@ : @@EMPTY@@,
        Purpose : purpose,
        Amount : amount,
        Profile : profile,
        Date : date,
    }
    return
}


type ExpensePurpose struct {
    Name  string
    Description  string
}

func NewExpensePurpose(@@EMPTY@@ string, name string, description string) (ret ExpensePurpose, err error) {
    ret = ExpensePurpose{
        @@EMPTY@@ : @@EMPTY@@,
        Name : name,
        Description : description,
    }
    return
}


type Fee struct {
    Student  *person.Profile
    Amount  float64
    Kind  *FeeKind
    Date  *time.Time
}

func NewFee(@@EMPTY@@ string, student *person.Profile, amount float64, kind *FeeKind, date *time.Time) (ret Fee, err error) {
    ret = Fee{
        @@EMPTY@@ : @@EMPTY@@,
        Student : student,
        Amount : amount,
        Kind : kind,
        Date : date,
    }
    return
}


type FeeKind struct {
    Name  string
    Amount  float64
    Description  string
}

func NewFeeKind(@@EMPTY@@ string, name string, amount float64, description string) (ret FeeKind, err error) {
    ret = FeeKind{
        @@EMPTY@@ : @@EMPTY@@,
        Name : name,
        Amount : amount,
        Description : description,
    }
    return
}



