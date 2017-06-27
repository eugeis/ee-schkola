package finance

import (
    "ee/schkola/person"
    "time"
)


type Expense struct {
    Purpose  *ExpensePurpose
    Amount  float64
    Person  *person.User
    Date  *time.Time
}

func NewExpense(purpose *ExpensePurpose, amount float64, person *person.User, date *time.Time) (ret Expense, err error) {
    ret = Expense{
        Purpose : purpose,
        Amount : amount,
        Person : person,
        Date : date,
    }
    return
}


type ExpensePurpose struct {
    Name  string
    Description  string
}

func NewExpensePurpose(name string, description string) (ret ExpensePurpose, err error) {
    ret = ExpensePurpose{
        Name : name,
        Description : description,
    }
    return
}


type Fee struct {
    Student  *person.User
    Amount  float64
    Kind  *FeeKind
    Date  *time.Time
}

func NewFee(student *person.User, amount float64, kind *FeeKind, date *time.Time) (ret Fee, err error) {
    ret = Fee{
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

func NewFeeKind(name string, amount float64, description string) (ret FeeKind, err error) {
    ret = FeeKind{
        Name : name,
        Amount : amount,
        Description : description,
    }
    return
}



