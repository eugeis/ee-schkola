package finance

import (
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "time"
)

type ExpenseCreated struct {
    Purpose  *ExpensePurpose
    Amount  float64
    Profile  *person.Profile
    Date  *time.Time
}

func NewExpenseCreated(purpose *ExpensePurpose, amount float64, profile *person.Profile, date *time.Time) (ret *ExpenseCreated, err error) {
    ret = &ExpenseCreated{
        Purpose : purpose,
        Amount : amount,
        Profile : profile,
        Date : date,
    }
    return
}



type ExpenseDeleted struct {
    Id  string
}

func NewExpenseDeleted(id string) (ret *ExpenseDeleted, err error) {
    ret = &ExpenseDeleted{
        Id : id,
    }
    return
}



type ExpenseUpdated struct {
    Purpose  *ExpensePurpose
    Amount  float64
    Profile  *person.Profile
    Date  *time.Time
}

func NewExpenseUpdated(purpose *ExpensePurpose, amount float64, profile *person.Profile, date *time.Time) (ret *ExpenseUpdated, err error) {
    ret = &ExpenseUpdated{
        Purpose : purpose,
        Amount : amount,
        Profile : profile,
        Date : date,
    }
    return
}



type ExpensePurposeCreated struct {
    Name  string
    Description  string
}

func NewExpensePurposeCreated(name string, description string) (ret *ExpensePurposeCreated, err error) {
    ret = &ExpensePurposeCreated{
        Name : name,
        Description : description,
    }
    return
}



type ExpensePurposeDeleted struct {
    Id  string
}

func NewExpensePurposeDeleted(id string) (ret *ExpensePurposeDeleted, err error) {
    ret = &ExpensePurposeDeleted{
        Id : id,
    }
    return
}



type ExpensePurposeUpdated struct {
    Name  string
    Description  string
}

func NewExpensePurposeUpdated(name string, description string) (ret *ExpensePurposeUpdated, err error) {
    ret = &ExpensePurposeUpdated{
        Name : name,
        Description : description,
    }
    return
}



type FeeCreated struct {
    Student  *person.Profile
    Amount  float64
    Kind  *FeeKind
    Date  *time.Time
}

func NewFeeCreated(student *person.Profile, amount float64, kind *FeeKind, date *time.Time) (ret *FeeCreated, err error) {
    ret = &FeeCreated{
        Student : student,
        Amount : amount,
        Kind : kind,
        Date : date,
    }
    return
}



type FeeDeleted struct {
    Id  string
}

func NewFeeDeleted(id string) (ret *FeeDeleted, err error) {
    ret = &FeeDeleted{
        Id : id,
    }
    return
}



type FeeUpdated struct {
    Student  *person.Profile
    Amount  float64
    Kind  *FeeKind
    Date  *time.Time
}

func NewFeeUpdated(student *person.Profile, amount float64, kind *FeeKind, date *time.Time) (ret *FeeUpdated, err error) {
    ret = &FeeUpdated{
        Student : student,
        Amount : amount,
        Kind : kind,
        Date : date,
    }
    return
}



type FeeKindCreated struct {
    Name  string
    Amount  float64
    Description  string
}

func NewFeeKindCreated(name string, amount float64, description string) (ret *FeeKindCreated, err error) {
    ret = &FeeKindCreated{
        Name : name,
        Amount : amount,
        Description : description,
    }
    return
}



type FeeKindDeleted struct {
    Id  string
}

func NewFeeKindDeleted(id string) (ret *FeeKindDeleted, err error) {
    ret = &FeeKindDeleted{
        Id : id,
    }
    return
}



type FeeKindUpdated struct {
    Name  string
    Amount  float64
    Description  string
}

func NewFeeKindUpdated(name string, amount float64, description string) (ret *FeeKindUpdated, err error) {
    ret = &FeeKindUpdated{
        Name : name,
        Amount : amount,
        Description : description,
    }
    return
}




type ExpenseEventType struct {
	name  string
	ordinal int
}

func (o *ExpenseEventType) Name() string {
    return o.name
}

func (o *ExpenseEventType) Ordinal() int {
    return o.ordinal
}

func (o *ExpenseEventType) IsCreatedExpense() bool {
    return o == _expenseEventTypes.CreatedExpense()
}

func (o *ExpenseEventType) IsDeletedExpense() bool {
    return o == _expenseEventTypes.DeletedExpense()
}

func (o *ExpenseEventType) IsUpdatedExpense() bool {
    return o == _expenseEventTypes.UpdatedExpense()
}

type expenseEventTypes struct {
	values []*ExpenseEventType
    literals []enum.Literal
}

var _expenseEventTypes = &expenseEventTypes{values: []*ExpenseEventType{
    {name: "createdExpense", ordinal: 0},
    {name: "deletedExpense", ordinal: 1},
    {name: "updatedExpense", ordinal: 2}},
}

func ExpenseEventTypes() *expenseEventTypes {
	return _expenseEventTypes
}

func (o *expenseEventTypes) Values() []*ExpenseEventType {
	return o.values
}

func (o *expenseEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *expenseEventTypes) CreatedExpense() *ExpenseEventType {
    return _expenseEventTypes.values[0]
}

func (o *expenseEventTypes) DeletedExpense() *ExpenseEventType {
    return _expenseEventTypes.values[1]
}

func (o *expenseEventTypes) UpdatedExpense() *ExpenseEventType {
    return _expenseEventTypes.values[2]
}

func (o *expenseEventTypes) ParseExpenseEventType(name string) (ret *ExpenseEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*ExpenseEventType), ok
	}
	return
}


type ExpensePurposeEventType struct {
	name  string
	ordinal int
}

func (o *ExpensePurposeEventType) Name() string {
    return o.name
}

func (o *ExpensePurposeEventType) Ordinal() int {
    return o.ordinal
}

func (o *ExpensePurposeEventType) IsCreatedExpensePurpose() bool {
    return o == _expensePurposeEventTypes.CreatedExpensePurpose()
}

func (o *ExpensePurposeEventType) IsDeletedExpensePurpose() bool {
    return o == _expensePurposeEventTypes.DeletedExpensePurpose()
}

func (o *ExpensePurposeEventType) IsUpdatedExpensePurpose() bool {
    return o == _expensePurposeEventTypes.UpdatedExpensePurpose()
}

type expensePurposeEventTypes struct {
	values []*ExpensePurposeEventType
    literals []enum.Literal
}

var _expensePurposeEventTypes = &expensePurposeEventTypes{values: []*ExpensePurposeEventType{
    {name: "createdExpensePurpose", ordinal: 0},
    {name: "deletedExpensePurpose", ordinal: 1},
    {name: "updatedExpensePurpose", ordinal: 2}},
}

func ExpensePurposeEventTypes() *expensePurposeEventTypes {
	return _expensePurposeEventTypes
}

func (o *expensePurposeEventTypes) Values() []*ExpensePurposeEventType {
	return o.values
}

func (o *expensePurposeEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *expensePurposeEventTypes) CreatedExpensePurpose() *ExpensePurposeEventType {
    return _expensePurposeEventTypes.values[0]
}

func (o *expensePurposeEventTypes) DeletedExpensePurpose() *ExpensePurposeEventType {
    return _expensePurposeEventTypes.values[1]
}

func (o *expensePurposeEventTypes) UpdatedExpensePurpose() *ExpensePurposeEventType {
    return _expensePurposeEventTypes.values[2]
}

func (o *expensePurposeEventTypes) ParseExpensePurposeEventType(name string) (ret *ExpensePurposeEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*ExpensePurposeEventType), ok
	}
	return
}


type FeeEventType struct {
	name  string
	ordinal int
}

func (o *FeeEventType) Name() string {
    return o.name
}

func (o *FeeEventType) Ordinal() int {
    return o.ordinal
}

func (o *FeeEventType) IsCreatedFee() bool {
    return o == _feeEventTypes.CreatedFee()
}

func (o *FeeEventType) IsDeletedFee() bool {
    return o == _feeEventTypes.DeletedFee()
}

func (o *FeeEventType) IsUpdatedFee() bool {
    return o == _feeEventTypes.UpdatedFee()
}

type feeEventTypes struct {
	values []*FeeEventType
    literals []enum.Literal
}

var _feeEventTypes = &feeEventTypes{values: []*FeeEventType{
    {name: "createdFee", ordinal: 0},
    {name: "deletedFee", ordinal: 1},
    {name: "updatedFee", ordinal: 2}},
}

func FeeEventTypes() *feeEventTypes {
	return _feeEventTypes
}

func (o *feeEventTypes) Values() []*FeeEventType {
	return o.values
}

func (o *feeEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *feeEventTypes) CreatedFee() *FeeEventType {
    return _feeEventTypes.values[0]
}

func (o *feeEventTypes) DeletedFee() *FeeEventType {
    return _feeEventTypes.values[1]
}

func (o *feeEventTypes) UpdatedFee() *FeeEventType {
    return _feeEventTypes.values[2]
}

func (o *feeEventTypes) ParseFeeEventType(name string) (ret *FeeEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*FeeEventType), ok
	}
	return
}


type FeeKindEventType struct {
	name  string
	ordinal int
}

func (o *FeeKindEventType) Name() string {
    return o.name
}

func (o *FeeKindEventType) Ordinal() int {
    return o.ordinal
}

func (o *FeeKindEventType) IsCreatedFeeKind() bool {
    return o == _feeKindEventTypes.CreatedFeeKind()
}

func (o *FeeKindEventType) IsDeletedFeeKind() bool {
    return o == _feeKindEventTypes.DeletedFeeKind()
}

func (o *FeeKindEventType) IsUpdatedFeeKind() bool {
    return o == _feeKindEventTypes.UpdatedFeeKind()
}

type feeKindEventTypes struct {
	values []*FeeKindEventType
    literals []enum.Literal
}

var _feeKindEventTypes = &feeKindEventTypes{values: []*FeeKindEventType{
    {name: "createdFeeKind", ordinal: 0},
    {name: "deletedFeeKind", ordinal: 1},
    {name: "updatedFeeKind", ordinal: 2}},
}

func FeeKindEventTypes() *feeKindEventTypes {
	return _feeKindEventTypes
}

func (o *feeKindEventTypes) Values() []*FeeKindEventType {
	return o.values
}

func (o *feeKindEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *feeKindEventTypes) CreatedFeeKind() *FeeKindEventType {
    return _feeKindEventTypes.values[0]
}

func (o *feeKindEventTypes) DeletedFeeKind() *FeeKindEventType {
    return _feeKindEventTypes.values[1]
}

func (o *feeKindEventTypes) UpdatedFeeKind() *FeeKindEventType {
    return _feeKindEventTypes.values[2]
}

func (o *feeKindEventTypes) ParseFeeKindEventType(name string) (ret *FeeKindEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*FeeKindEventType), ok
	}
	return
}



