package finance

import (
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "time"
)

type CreatedExpense struct {
    Purpose  *ExpensePurpose
    Amount  float64
    Profile  *person.Profile
    Date  *time.Time
}

func NewCreatedExpense(purpose *ExpensePurpose, amount float64, profile *person.Profile, date *time.Time) (ret *CreatedExpense, err error) {
    ret = &CreatedExpense{
        Purpose : purpose,
        Amount : amount,
        Profile : profile,
        Date : date,
    }
    return
}



type DeletedExpense struct {
    Id  string
}

func NewDeletedExpense(id string) (ret *DeletedExpense, err error) {
    ret = &DeletedExpense{
        Id : id,
    }
    return
}



type UpdatedExpense struct {
    Purpose  *ExpensePurpose
    Amount  float64
    Profile  *person.Profile
    Date  *time.Time
}

func NewUpdatedExpense(purpose *ExpensePurpose, amount float64, profile *person.Profile, date *time.Time) (ret *UpdatedExpense, err error) {
    ret = &UpdatedExpense{
        Purpose : purpose,
        Amount : amount,
        Profile : profile,
        Date : date,
    }
    return
}



type CreatedExpensePurpose struct {
    Name  string
    Description  string
}

func NewCreatedExpensePurpose(name string, description string) (ret *CreatedExpensePurpose, err error) {
    ret = &CreatedExpensePurpose{
        Name : name,
        Description : description,
    }
    return
}



type DeletedExpensePurpose struct {
    Id  string
}

func NewDeletedExpensePurpose(id string) (ret *DeletedExpensePurpose, err error) {
    ret = &DeletedExpensePurpose{
        Id : id,
    }
    return
}



type UpdatedExpensePurpose struct {
    Name  string
    Description  string
}

func NewUpdatedExpensePurpose(name string, description string) (ret *UpdatedExpensePurpose, err error) {
    ret = &UpdatedExpensePurpose{
        Name : name,
        Description : description,
    }
    return
}



type CreatedFee struct {
    Student  *person.Profile
    Amount  float64
    Kind  *FeeKind
    Date  *time.Time
}

func NewCreatedFee(student *person.Profile, amount float64, kind *FeeKind, date *time.Time) (ret *CreatedFee, err error) {
    ret = &CreatedFee{
        Student : student,
        Amount : amount,
        Kind : kind,
        Date : date,
    }
    return
}



type DeletedFee struct {
    Id  string
}

func NewDeletedFee(id string) (ret *DeletedFee, err error) {
    ret = &DeletedFee{
        Id : id,
    }
    return
}



type UpdatedFee struct {
    Student  *person.Profile
    Amount  float64
    Kind  *FeeKind
    Date  *time.Time
}

func NewUpdatedFee(student *person.Profile, amount float64, kind *FeeKind, date *time.Time) (ret *UpdatedFee, err error) {
    ret = &UpdatedFee{
        Student : student,
        Amount : amount,
        Kind : kind,
        Date : date,
    }
    return
}



type CreatedFeeKind struct {
    Name  string
    Amount  float64
    Description  string
}

func NewCreatedFeeKind(name string, amount float64, description string) (ret *CreatedFeeKind, err error) {
    ret = &CreatedFeeKind{
        Name : name,
        Amount : amount,
        Description : description,
    }
    return
}



type DeletedFeeKind struct {
    Id  string
}

func NewDeletedFeeKind(id string) (ret *DeletedFeeKind, err error) {
    ret = &DeletedFeeKind{
        Id : id,
    }
    return
}



type UpdatedFeeKind struct {
    Name  string
    Amount  float64
    Description  string
}

func NewUpdatedFeeKind(name string, amount float64, description string) (ret *UpdatedFeeKind, err error) {
    ret = &UpdatedFeeKind{
        Name : name,
        Amount : amount,
        Description : description,
    }
    return
}




type ExpenseAggregateEventType struct {
	name  string
	ordinal int
}

func (o *ExpenseAggregateEventType) Name() string {
    return o.name
}

func (o *ExpenseAggregateEventType) Ordinal() int {
    return o.ordinal
}

func (o *ExpenseAggregateEventType) IsExpenseCreated() bool {
    return o == _expenseAggregateEventTypes.ExpenseCreated()
}

func (o *ExpenseAggregateEventType) IsExpenseDeleted() bool {
    return o == _expenseAggregateEventTypes.ExpenseDeleted()
}

func (o *ExpenseAggregateEventType) IsExpenseUpdated() bool {
    return o == _expenseAggregateEventTypes.ExpenseUpdated()
}

type expenseAggregateEventTypes struct {
	values []*ExpenseAggregateEventType
    literals []enum.Literal
}

var _expenseAggregateEventTypes = &expenseAggregateEventTypes{values: []*ExpenseAggregateEventType{
    {name: "ExpenseCreated", ordinal: 0},
    {name: "ExpenseDeleted", ordinal: 1},
    {name: "ExpenseUpdated", ordinal: 2}},
}

func ExpenseAggregateEventTypes() *expenseAggregateEventTypes {
	return _expenseAggregateEventTypes
}

func (o *expenseAggregateEventTypes) Values() []*ExpenseAggregateEventType {
	return o.values
}

func (o *expenseAggregateEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *expenseAggregateEventTypes) ExpenseCreated() *ExpenseAggregateEventType {
    return _expenseAggregateEventTypes.values[0]
}

func (o *expenseAggregateEventTypes) ExpenseDeleted() *ExpenseAggregateEventType {
    return _expenseAggregateEventTypes.values[1]
}

func (o *expenseAggregateEventTypes) ExpenseUpdated() *ExpenseAggregateEventType {
    return _expenseAggregateEventTypes.values[2]
}

func (o *expenseAggregateEventTypes) ParseExpenseAggregateEventType(name string) (ret *ExpenseAggregateEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*ExpenseAggregateEventType), ok
	}
	return
}


type ExpensePurposeAggregateEventType struct {
	name  string
	ordinal int
}

func (o *ExpensePurposeAggregateEventType) Name() string {
    return o.name
}

func (o *ExpensePurposeAggregateEventType) Ordinal() int {
    return o.ordinal
}

func (o *ExpensePurposeAggregateEventType) IsExpensePurposeCreated() bool {
    return o == _expensePurposeAggregateEventTypes.ExpensePurposeCreated()
}

func (o *ExpensePurposeAggregateEventType) IsExpensePurposeDeleted() bool {
    return o == _expensePurposeAggregateEventTypes.ExpensePurposeDeleted()
}

func (o *ExpensePurposeAggregateEventType) IsExpensePurposeUpdated() bool {
    return o == _expensePurposeAggregateEventTypes.ExpensePurposeUpdated()
}

type expensePurposeAggregateEventTypes struct {
	values []*ExpensePurposeAggregateEventType
    literals []enum.Literal
}

var _expensePurposeAggregateEventTypes = &expensePurposeAggregateEventTypes{values: []*ExpensePurposeAggregateEventType{
    {name: "ExpensePurposeCreated", ordinal: 0},
    {name: "ExpensePurposeDeleted", ordinal: 1},
    {name: "ExpensePurposeUpdated", ordinal: 2}},
}

func ExpensePurposeAggregateEventTypes() *expensePurposeAggregateEventTypes {
	return _expensePurposeAggregateEventTypes
}

func (o *expensePurposeAggregateEventTypes) Values() []*ExpensePurposeAggregateEventType {
	return o.values
}

func (o *expensePurposeAggregateEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *expensePurposeAggregateEventTypes) ExpensePurposeCreated() *ExpensePurposeAggregateEventType {
    return _expensePurposeAggregateEventTypes.values[0]
}

func (o *expensePurposeAggregateEventTypes) ExpensePurposeDeleted() *ExpensePurposeAggregateEventType {
    return _expensePurposeAggregateEventTypes.values[1]
}

func (o *expensePurposeAggregateEventTypes) ExpensePurposeUpdated() *ExpensePurposeAggregateEventType {
    return _expensePurposeAggregateEventTypes.values[2]
}

func (o *expensePurposeAggregateEventTypes) ParseExpensePurposeAggregateEventType(name string) (ret *ExpensePurposeAggregateEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*ExpensePurposeAggregateEventType), ok
	}
	return
}


type FeeAggregateEventType struct {
	name  string
	ordinal int
}

func (o *FeeAggregateEventType) Name() string {
    return o.name
}

func (o *FeeAggregateEventType) Ordinal() int {
    return o.ordinal
}

func (o *FeeAggregateEventType) IsFeeCreated() bool {
    return o == _feeAggregateEventTypes.FeeCreated()
}

func (o *FeeAggregateEventType) IsFeeDeleted() bool {
    return o == _feeAggregateEventTypes.FeeDeleted()
}

func (o *FeeAggregateEventType) IsFeeUpdated() bool {
    return o == _feeAggregateEventTypes.FeeUpdated()
}

type feeAggregateEventTypes struct {
	values []*FeeAggregateEventType
    literals []enum.Literal
}

var _feeAggregateEventTypes = &feeAggregateEventTypes{values: []*FeeAggregateEventType{
    {name: "FeeCreated", ordinal: 0},
    {name: "FeeDeleted", ordinal: 1},
    {name: "FeeUpdated", ordinal: 2}},
}

func FeeAggregateEventTypes() *feeAggregateEventTypes {
	return _feeAggregateEventTypes
}

func (o *feeAggregateEventTypes) Values() []*FeeAggregateEventType {
	return o.values
}

func (o *feeAggregateEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *feeAggregateEventTypes) FeeCreated() *FeeAggregateEventType {
    return _feeAggregateEventTypes.values[0]
}

func (o *feeAggregateEventTypes) FeeDeleted() *FeeAggregateEventType {
    return _feeAggregateEventTypes.values[1]
}

func (o *feeAggregateEventTypes) FeeUpdated() *FeeAggregateEventType {
    return _feeAggregateEventTypes.values[2]
}

func (o *feeAggregateEventTypes) ParseFeeAggregateEventType(name string) (ret *FeeAggregateEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*FeeAggregateEventType), ok
	}
	return
}


type FeeKindAggregateEventType struct {
	name  string
	ordinal int
}

func (o *FeeKindAggregateEventType) Name() string {
    return o.name
}

func (o *FeeKindAggregateEventType) Ordinal() int {
    return o.ordinal
}

func (o *FeeKindAggregateEventType) IsFeeKindCreated() bool {
    return o == _feeKindAggregateEventTypes.FeeKindCreated()
}

func (o *FeeKindAggregateEventType) IsFeeKindDeleted() bool {
    return o == _feeKindAggregateEventTypes.FeeKindDeleted()
}

func (o *FeeKindAggregateEventType) IsFeeKindUpdated() bool {
    return o == _feeKindAggregateEventTypes.FeeKindUpdated()
}

type feeKindAggregateEventTypes struct {
	values []*FeeKindAggregateEventType
    literals []enum.Literal
}

var _feeKindAggregateEventTypes = &feeKindAggregateEventTypes{values: []*FeeKindAggregateEventType{
    {name: "FeeKindCreated", ordinal: 0},
    {name: "FeeKindDeleted", ordinal: 1},
    {name: "FeeKindUpdated", ordinal: 2}},
}

func FeeKindAggregateEventTypes() *feeKindAggregateEventTypes {
	return _feeKindAggregateEventTypes
}

func (o *feeKindAggregateEventTypes) Values() []*FeeKindAggregateEventType {
	return o.values
}

func (o *feeKindAggregateEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *feeKindAggregateEventTypes) FeeKindCreated() *FeeKindAggregateEventType {
    return _feeKindAggregateEventTypes.values[0]
}

func (o *feeKindAggregateEventTypes) FeeKindDeleted() *FeeKindAggregateEventType {
    return _feeKindAggregateEventTypes.values[1]
}

func (o *feeKindAggregateEventTypes) FeeKindUpdated() *FeeKindAggregateEventType {
    return _feeKindAggregateEventTypes.values[2]
}

func (o *feeKindAggregateEventTypes) ParseFeeKindAggregateEventType(name string) (ret *FeeKindAggregateEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*FeeKindAggregateEventType), ok
	}
	return
}



