package finance

import (
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "time"
)

type ExpenseCreate struct {
    Purpose  *ExpensePurpose
    Amount  float64
    Profile  *person.Profile
    Date  *time.Time
}

func NewExpenseCreate(purpose *ExpensePurpose, amount float64, profile *person.Profile, date *time.Time) (ret *ExpenseCreate, err error) {
    ret = &ExpenseCreate{
        Purpose : purpose,
        Amount : amount,
        Profile : profile,
        Date : date,
    }
    return
}



type ExpenseDelete struct {
    Id  string
}

func NewExpenseDelete(id string) (ret *ExpenseDelete, err error) {
    ret = &ExpenseDelete{
        Id : id,
    }
    return
}



type ExpenseUpdate struct {
    Purpose  *ExpensePurpose
    Amount  float64
    Profile  *person.Profile
    Date  *time.Time
}

func NewExpenseUpdate(purpose *ExpensePurpose, amount float64, profile *person.Profile, date *time.Time) (ret *ExpenseUpdate, err error) {
    ret = &ExpenseUpdate{
        Purpose : purpose,
        Amount : amount,
        Profile : profile,
        Date : date,
    }
    return
}



type ExpensePurposeCreate struct {
    Name  string
    Description  string
}

func NewExpensePurposeCreate(name string, description string) (ret *ExpensePurposeCreate, err error) {
    ret = &ExpensePurposeCreate{
        Name : name,
        Description : description,
    }
    return
}



type ExpensePurposeDelete struct {
    Id  string
}

func NewExpensePurposeDelete(id string) (ret *ExpensePurposeDelete, err error) {
    ret = &ExpensePurposeDelete{
        Id : id,
    }
    return
}



type ExpensePurposeUpdate struct {
    Name  string
    Description  string
}

func NewExpensePurposeUpdate(name string, description string) (ret *ExpensePurposeUpdate, err error) {
    ret = &ExpensePurposeUpdate{
        Name : name,
        Description : description,
    }
    return
}



type FeeCreate struct {
    Student  *person.Profile
    Amount  float64
    Kind  *FeeKind
    Date  *time.Time
}

func NewFeeCreate(student *person.Profile, amount float64, kind *FeeKind, date *time.Time) (ret *FeeCreate, err error) {
    ret = &FeeCreate{
        Student : student,
        Amount : amount,
        Kind : kind,
        Date : date,
    }
    return
}



type FeeDelete struct {
    Id  string
}

func NewFeeDelete(id string) (ret *FeeDelete, err error) {
    ret = &FeeDelete{
        Id : id,
    }
    return
}



type FeeUpdate struct {
    Student  *person.Profile
    Amount  float64
    Kind  *FeeKind
    Date  *time.Time
}

func NewFeeUpdate(student *person.Profile, amount float64, kind *FeeKind, date *time.Time) (ret *FeeUpdate, err error) {
    ret = &FeeUpdate{
        Student : student,
        Amount : amount,
        Kind : kind,
        Date : date,
    }
    return
}



type FeeKindCreate struct {
    Name  string
    Amount  float64
    Description  string
}

func NewFeeKindCreate(name string, amount float64, description string) (ret *FeeKindCreate, err error) {
    ret = &FeeKindCreate{
        Name : name,
        Amount : amount,
        Description : description,
    }
    return
}



type FeeKindDelete struct {
    Id  string
}

func NewFeeKindDelete(id string) (ret *FeeKindDelete, err error) {
    ret = &FeeKindDelete{
        Id : id,
    }
    return
}



type FeeKindUpdate struct {
    Name  string
    Amount  float64
    Description  string
}

func NewFeeKindUpdate(name string, amount float64, description string) (ret *FeeKindUpdate, err error) {
    ret = &FeeKindUpdate{
        Name : name,
        Amount : amount,
        Description : description,
    }
    return
}




type ExpenseAggregateCommandType struct {
	name  string
	ordinal int
}

func (o *ExpenseAggregateCommandType) Name() string {
    return o.name
}

func (o *ExpenseAggregateCommandType) Ordinal() int {
    return o.ordinal
}

func (o *ExpenseAggregateCommandType) IsCreateExpense() bool {
    return o == _expenseAggregateCommandTypes.CreateExpense()
}

func (o *ExpenseAggregateCommandType) IsDeleteExpense() bool {
    return o == _expenseAggregateCommandTypes.DeleteExpense()
}

func (o *ExpenseAggregateCommandType) IsUpdateExpense() bool {
    return o == _expenseAggregateCommandTypes.UpdateExpense()
}

type expenseAggregateCommandTypes struct {
	values []*ExpenseAggregateCommandType
    literals []enum.Literal
}

var _expenseAggregateCommandTypes = &expenseAggregateCommandTypes{values: []*ExpenseAggregateCommandType{
    {name: "createExpense", ordinal: 0},
    {name: "deleteExpense", ordinal: 1},
    {name: "updateExpense", ordinal: 2}},
}

func ExpenseAggregateCommandTypes() *expenseAggregateCommandTypes {
	return _expenseAggregateCommandTypes
}

func (o *expenseAggregateCommandTypes) Values() []*ExpenseAggregateCommandType {
	return o.values
}

func (o *expenseAggregateCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *expenseAggregateCommandTypes) CreateExpense() *ExpenseAggregateCommandType {
    return _expenseAggregateCommandTypes.values[0]
}

func (o *expenseAggregateCommandTypes) DeleteExpense() *ExpenseAggregateCommandType {
    return _expenseAggregateCommandTypes.values[1]
}

func (o *expenseAggregateCommandTypes) UpdateExpense() *ExpenseAggregateCommandType {
    return _expenseAggregateCommandTypes.values[2]
}

func (o *expenseAggregateCommandTypes) ParseExpenseAggregateCommandType(name string) (ret *ExpenseAggregateCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*ExpenseAggregateCommandType), ok
	}
	return
}


type ExpensePurposeAggregateCommandType struct {
	name  string
	ordinal int
}

func (o *ExpensePurposeAggregateCommandType) Name() string {
    return o.name
}

func (o *ExpensePurposeAggregateCommandType) Ordinal() int {
    return o.ordinal
}

func (o *ExpensePurposeAggregateCommandType) IsCreateExpensePurpose() bool {
    return o == _expensePurposeAggregateCommandTypes.CreateExpensePurpose()
}

func (o *ExpensePurposeAggregateCommandType) IsDeleteExpensePurpose() bool {
    return o == _expensePurposeAggregateCommandTypes.DeleteExpensePurpose()
}

func (o *ExpensePurposeAggregateCommandType) IsUpdateExpensePurpose() bool {
    return o == _expensePurposeAggregateCommandTypes.UpdateExpensePurpose()
}

type expensePurposeAggregateCommandTypes struct {
	values []*ExpensePurposeAggregateCommandType
    literals []enum.Literal
}

var _expensePurposeAggregateCommandTypes = &expensePurposeAggregateCommandTypes{values: []*ExpensePurposeAggregateCommandType{
    {name: "createExpensePurpose", ordinal: 0},
    {name: "deleteExpensePurpose", ordinal: 1},
    {name: "updateExpensePurpose", ordinal: 2}},
}

func ExpensePurposeAggregateCommandTypes() *expensePurposeAggregateCommandTypes {
	return _expensePurposeAggregateCommandTypes
}

func (o *expensePurposeAggregateCommandTypes) Values() []*ExpensePurposeAggregateCommandType {
	return o.values
}

func (o *expensePurposeAggregateCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *expensePurposeAggregateCommandTypes) CreateExpensePurpose() *ExpensePurposeAggregateCommandType {
    return _expensePurposeAggregateCommandTypes.values[0]
}

func (o *expensePurposeAggregateCommandTypes) DeleteExpensePurpose() *ExpensePurposeAggregateCommandType {
    return _expensePurposeAggregateCommandTypes.values[1]
}

func (o *expensePurposeAggregateCommandTypes) UpdateExpensePurpose() *ExpensePurposeAggregateCommandType {
    return _expensePurposeAggregateCommandTypes.values[2]
}

func (o *expensePurposeAggregateCommandTypes) ParseExpensePurposeAggregateCommandType(name string) (ret *ExpensePurposeAggregateCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*ExpensePurposeAggregateCommandType), ok
	}
	return
}


type FeeAggregateCommandType struct {
	name  string
	ordinal int
}

func (o *FeeAggregateCommandType) Name() string {
    return o.name
}

func (o *FeeAggregateCommandType) Ordinal() int {
    return o.ordinal
}

func (o *FeeAggregateCommandType) IsCreateFee() bool {
    return o == _feeAggregateCommandTypes.CreateFee()
}

func (o *FeeAggregateCommandType) IsDeleteFee() bool {
    return o == _feeAggregateCommandTypes.DeleteFee()
}

func (o *FeeAggregateCommandType) IsUpdateFee() bool {
    return o == _feeAggregateCommandTypes.UpdateFee()
}

type feeAggregateCommandTypes struct {
	values []*FeeAggregateCommandType
    literals []enum.Literal
}

var _feeAggregateCommandTypes = &feeAggregateCommandTypes{values: []*FeeAggregateCommandType{
    {name: "createFee", ordinal: 0},
    {name: "deleteFee", ordinal: 1},
    {name: "updateFee", ordinal: 2}},
}

func FeeAggregateCommandTypes() *feeAggregateCommandTypes {
	return _feeAggregateCommandTypes
}

func (o *feeAggregateCommandTypes) Values() []*FeeAggregateCommandType {
	return o.values
}

func (o *feeAggregateCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *feeAggregateCommandTypes) CreateFee() *FeeAggregateCommandType {
    return _feeAggregateCommandTypes.values[0]
}

func (o *feeAggregateCommandTypes) DeleteFee() *FeeAggregateCommandType {
    return _feeAggregateCommandTypes.values[1]
}

func (o *feeAggregateCommandTypes) UpdateFee() *FeeAggregateCommandType {
    return _feeAggregateCommandTypes.values[2]
}

func (o *feeAggregateCommandTypes) ParseFeeAggregateCommandType(name string) (ret *FeeAggregateCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*FeeAggregateCommandType), ok
	}
	return
}


type FeeKindAggregateCommandType struct {
	name  string
	ordinal int
}

func (o *FeeKindAggregateCommandType) Name() string {
    return o.name
}

func (o *FeeKindAggregateCommandType) Ordinal() int {
    return o.ordinal
}

func (o *FeeKindAggregateCommandType) IsCreateFeeKind() bool {
    return o == _feeKindAggregateCommandTypes.CreateFeeKind()
}

func (o *FeeKindAggregateCommandType) IsDeleteFeeKind() bool {
    return o == _feeKindAggregateCommandTypes.DeleteFeeKind()
}

func (o *FeeKindAggregateCommandType) IsUpdateFeeKind() bool {
    return o == _feeKindAggregateCommandTypes.UpdateFeeKind()
}

type feeKindAggregateCommandTypes struct {
	values []*FeeKindAggregateCommandType
    literals []enum.Literal
}

var _feeKindAggregateCommandTypes = &feeKindAggregateCommandTypes{values: []*FeeKindAggregateCommandType{
    {name: "createFeeKind", ordinal: 0},
    {name: "deleteFeeKind", ordinal: 1},
    {name: "updateFeeKind", ordinal: 2}},
}

func FeeKindAggregateCommandTypes() *feeKindAggregateCommandTypes {
	return _feeKindAggregateCommandTypes
}

func (o *feeKindAggregateCommandTypes) Values() []*FeeKindAggregateCommandType {
	return o.values
}

func (o *feeKindAggregateCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *feeKindAggregateCommandTypes) CreateFeeKind() *FeeKindAggregateCommandType {
    return _feeKindAggregateCommandTypes.values[0]
}

func (o *feeKindAggregateCommandTypes) DeleteFeeKind() *FeeKindAggregateCommandType {
    return _feeKindAggregateCommandTypes.values[1]
}

func (o *feeKindAggregateCommandTypes) UpdateFeeKind() *FeeKindAggregateCommandType {
    return _feeKindAggregateCommandTypes.values[2]
}

func (o *feeKindAggregateCommandTypes) ParseFeeKindAggregateCommandType(name string) (ret *FeeKindAggregateCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*FeeKindAggregateCommandType), ok
	}
	return
}



