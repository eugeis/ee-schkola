package finance

import (
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "time"
)

type CreatedExpense struct {
    purpose *ExpensePurpose
    amount float64
    profile *person.Profile
    date *time.Time
}

func NewCreatedExpense(purpose *ExpensePurpose, amount float64, profile *person.Profile, date *time.Time) (ret *CreatedExpense, err error) {
    ret = &CreatedExpense{
        purpose: purpose,
        amount: amount,
        profile: profile,
        date: date,
    }
    
    return
}



type DeletedExpense struct {
    id string
}

func NewDeletedExpense(id string) (ret *DeletedExpense, err error) {
    ret = &DeletedExpense{
        id: id,
    }
    
    return
}



type UpdatedExpense struct {
    purpose *ExpensePurpose
    amount float64
    profile *person.Profile
    date *time.Time
}

func NewUpdatedExpense(purpose *ExpensePurpose, amount float64, profile *person.Profile, date *time.Time) (ret *UpdatedExpense, err error) {
    ret = &UpdatedExpense{
        purpose: purpose,
        amount: amount,
        profile: profile,
        date: date,
    }
    
    return
}



type CreatedExpensePurpose struct {
    name string
    description string
}

func NewCreatedExpensePurpose(name string, description string) (ret *CreatedExpensePurpose, err error) {
    ret = &CreatedExpensePurpose{
        name: name,
        description: description,
    }
    
    return
}



type DeletedExpensePurpose struct {
    id string
}

func NewDeletedExpensePurpose(id string) (ret *DeletedExpensePurpose, err error) {
    ret = &DeletedExpensePurpose{
        id: id,
    }
    
    return
}



type UpdatedExpensePurpose struct {
    name string
    description string
}

func NewUpdatedExpensePurpose(name string, description string) (ret *UpdatedExpensePurpose, err error) {
    ret = &UpdatedExpensePurpose{
        name: name,
        description: description,
    }
    
    return
}



type CreatedFee struct {
    student *person.Profile
    amount float64
    kind *FeeKind
    date *time.Time
}

func NewCreatedFee(student *person.Profile, amount float64, kind *FeeKind, date *time.Time) (ret *CreatedFee, err error) {
    ret = &CreatedFee{
        student: student,
        amount: amount,
        kind: kind,
        date: date,
    }
    
    return
}



type DeletedFee struct {
    id string
}

func NewDeletedFee(id string) (ret *DeletedFee, err error) {
    ret = &DeletedFee{
        id: id,
    }
    
    return
}



type UpdatedFee struct {
    student *person.Profile
    amount float64
    kind *FeeKind
    date *time.Time
}

func NewUpdatedFee(student *person.Profile, amount float64, kind *FeeKind, date *time.Time) (ret *UpdatedFee, err error) {
    ret = &UpdatedFee{
        student: student,
        amount: amount,
        kind: kind,
        date: date,
    }
    
    return
}



type CreatedFeeKind struct {
    name string
    amount float64
    description string
}

func NewCreatedFeeKind(name string, amount float64, description string) (ret *CreatedFeeKind, err error) {
    ret = &CreatedFeeKind{
        name: name,
        amount: amount,
        description: description,
    }
    
    return
}



type DeletedFeeKind struct {
    id string
}

func NewDeletedFeeKind(id string) (ret *DeletedFeeKind, err error) {
    ret = &DeletedFeeKind{
        id: id,
    }
    
    return
}



type UpdatedFeeKind struct {
    name string
    amount float64
    description string
}

func NewUpdatedFeeKind(name string, amount float64, description string) (ret *UpdatedFeeKind, err error) {
    ret = &UpdatedFeeKind{
        name: name,
        amount: amount,
        description: description,
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

func (o *ExpenseEventType) IsExpenseCreated() bool {
    return o == _expenseEventTypes.ExpenseCreated()
}

func (o *ExpenseEventType) IsExpenseDeleted() bool {
    return o == _expenseEventTypes.ExpenseDeleted()
}

func (o *ExpenseEventType) IsExpenseUpdated() bool {
    return o == _expenseEventTypes.ExpenseUpdated()
}

type expenseEventTypes struct {
	values []*ExpenseEventType
    literals []enum.Literal
}

var _expenseEventTypes = &expenseEventTypes{values: []*ExpenseEventType{
    {name: "ExpenseCreated", ordinal: 0},
    {name: "ExpenseDeleted", ordinal: 1},
    {name: "ExpenseUpdated", ordinal: 2}},
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

func (o *expenseEventTypes) ExpenseCreated() *ExpenseEventType {
    return _expenseEventTypes.values[0]
}

func (o *expenseEventTypes) ExpenseDeleted() *ExpenseEventType {
    return _expenseEventTypes.values[1]
}

func (o *expenseEventTypes) ExpenseUpdated() *ExpenseEventType {
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

func (o *ExpensePurposeEventType) IsExpensePurposeCreated() bool {
    return o == _expensePurposeEventTypes.ExpensePurposeCreated()
}

func (o *ExpensePurposeEventType) IsExpensePurposeDeleted() bool {
    return o == _expensePurposeEventTypes.ExpensePurposeDeleted()
}

func (o *ExpensePurposeEventType) IsExpensePurposeUpdated() bool {
    return o == _expensePurposeEventTypes.ExpensePurposeUpdated()
}

type expensePurposeEventTypes struct {
	values []*ExpensePurposeEventType
    literals []enum.Literal
}

var _expensePurposeEventTypes = &expensePurposeEventTypes{values: []*ExpensePurposeEventType{
    {name: "ExpensePurposeCreated", ordinal: 0},
    {name: "ExpensePurposeDeleted", ordinal: 1},
    {name: "ExpensePurposeUpdated", ordinal: 2}},
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

func (o *expensePurposeEventTypes) ExpensePurposeCreated() *ExpensePurposeEventType {
    return _expensePurposeEventTypes.values[0]
}

func (o *expensePurposeEventTypes) ExpensePurposeDeleted() *ExpensePurposeEventType {
    return _expensePurposeEventTypes.values[1]
}

func (o *expensePurposeEventTypes) ExpensePurposeUpdated() *ExpensePurposeEventType {
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

func (o *FeeEventType) IsFeeCreated() bool {
    return o == _feeEventTypes.FeeCreated()
}

func (o *FeeEventType) IsFeeDeleted() bool {
    return o == _feeEventTypes.FeeDeleted()
}

func (o *FeeEventType) IsFeeUpdated() bool {
    return o == _feeEventTypes.FeeUpdated()
}

type feeEventTypes struct {
	values []*FeeEventType
    literals []enum.Literal
}

var _feeEventTypes = &feeEventTypes{values: []*FeeEventType{
    {name: "FeeCreated", ordinal: 0},
    {name: "FeeDeleted", ordinal: 1},
    {name: "FeeUpdated", ordinal: 2}},
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

func (o *feeEventTypes) FeeCreated() *FeeEventType {
    return _feeEventTypes.values[0]
}

func (o *feeEventTypes) FeeDeleted() *FeeEventType {
    return _feeEventTypes.values[1]
}

func (o *feeEventTypes) FeeUpdated() *FeeEventType {
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

func (o *FeeKindEventType) IsFeeKindCreated() bool {
    return o == _feeKindEventTypes.FeeKindCreated()
}

func (o *FeeKindEventType) IsFeeKindDeleted() bool {
    return o == _feeKindEventTypes.FeeKindDeleted()
}

func (o *FeeKindEventType) IsFeeKindUpdated() bool {
    return o == _feeKindEventTypes.FeeKindUpdated()
}

type feeKindEventTypes struct {
	values []*FeeKindEventType
    literals []enum.Literal
}

var _feeKindEventTypes = &feeKindEventTypes{values: []*FeeKindEventType{
    {name: "FeeKindCreated", ordinal: 0},
    {name: "FeeKindDeleted", ordinal: 1},
    {name: "FeeKindUpdated", ordinal: 2}},
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

func (o *feeKindEventTypes) FeeKindCreated() *FeeKindEventType {
    return _feeKindEventTypes.values[0]
}

func (o *feeKindEventTypes) FeeKindDeleted() *FeeKindEventType {
    return _feeKindEventTypes.values[1]
}

func (o *feeKindEventTypes) FeeKindUpdated() *FeeKindEventType {
    return _feeKindEventTypes.values[2]
}

func (o *feeKindEventTypes) ParseFeeKindEventType(name string) (ret *FeeKindEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*FeeKindEventType), ok
	}
	return
}



