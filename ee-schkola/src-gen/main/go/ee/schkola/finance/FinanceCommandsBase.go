package finance

import (
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "time"
)

type ExpenseCreate struct {
    purpose *ExpensePurpose
    amount float64
    profile *person.Profile
    date *time.Time
}

func NewExpenseCreate(purpose *ExpensePurpose, amount float64, profile *person.Profile, date *time.Time) (ret *ExpenseCreate, err error) {
    ret = &ExpenseCreate{
        purpose: purpose,
        amount: amount,
        profile: profile,
        date: date,
    }
    
    return
}



type ExpenseDelete struct {
    id string
}

func NewExpenseDelete(id string) (ret *ExpenseDelete, err error) {
    ret = &ExpenseDelete{
        id: id,
    }
    
    return
}



type ExpenseUpdate struct {
    purpose *ExpensePurpose
    amount float64
    profile *person.Profile
    date *time.Time
}

func NewExpenseUpdate(purpose *ExpensePurpose, amount float64, profile *person.Profile, date *time.Time) (ret *ExpenseUpdate, err error) {
    ret = &ExpenseUpdate{
        purpose: purpose,
        amount: amount,
        profile: profile,
        date: date,
    }
    
    return
}



type ExpensePurposeCreate struct {
    name string
    description string
}

func NewExpensePurposeCreate(name string, description string) (ret *ExpensePurposeCreate, err error) {
    ret = &ExpensePurposeCreate{
        name: name,
        description: description,
    }
    
    return
}



type ExpensePurposeDelete struct {
    id string
}

func NewExpensePurposeDelete(id string) (ret *ExpensePurposeDelete, err error) {
    ret = &ExpensePurposeDelete{
        id: id,
    }
    
    return
}



type ExpensePurposeUpdate struct {
    name string
    description string
}

func NewExpensePurposeUpdate(name string, description string) (ret *ExpensePurposeUpdate, err error) {
    ret = &ExpensePurposeUpdate{
        name: name,
        description: description,
    }
    
    return
}



type FeeCreate struct {
    student *person.Profile
    amount float64
    kind *FeeKind
    date *time.Time
}

func NewFeeCreate(student *person.Profile, amount float64, kind *FeeKind, date *time.Time) (ret *FeeCreate, err error) {
    ret = &FeeCreate{
        student: student,
        amount: amount,
        kind: kind,
        date: date,
    }
    
    return
}



type FeeDelete struct {
    id string
}

func NewFeeDelete(id string) (ret *FeeDelete, err error) {
    ret = &FeeDelete{
        id: id,
    }
    
    return
}



type FeeUpdate struct {
    student *person.Profile
    amount float64
    kind *FeeKind
    date *time.Time
}

func NewFeeUpdate(student *person.Profile, amount float64, kind *FeeKind, date *time.Time) (ret *FeeUpdate, err error) {
    ret = &FeeUpdate{
        student: student,
        amount: amount,
        kind: kind,
        date: date,
    }
    
    return
}



type FeeKindCreate struct {
    name string
    amount float64
    description string
}

func NewFeeKindCreate(name string, amount float64, description string) (ret *FeeKindCreate, err error) {
    ret = &FeeKindCreate{
        name: name,
        amount: amount,
        description: description,
    }
    
    return
}



type FeeKindDelete struct {
    id string
}

func NewFeeKindDelete(id string) (ret *FeeKindDelete, err error) {
    ret = &FeeKindDelete{
        id: id,
    }
    
    return
}



type FeeKindUpdate struct {
    name string
    amount float64
    description string
}

func NewFeeKindUpdate(name string, amount float64, description string) (ret *FeeKindUpdate, err error) {
    ret = &FeeKindUpdate{
        name: name,
        amount: amount,
        description: description,
    }
    
    return
}




type ExpenseCommandType struct {
	name  string
	ordinal int
}

func (o *ExpenseCommandType) Name() string {
    return o.name
}

func (o *ExpenseCommandType) Ordinal() int {
    return o.ordinal
}

func (o *ExpenseCommandType) IsCreateExpense() bool {
    return o == _expenseCommandTypes.CreateExpense()
}

func (o *ExpenseCommandType) IsDeleteExpense() bool {
    return o == _expenseCommandTypes.DeleteExpense()
}

func (o *ExpenseCommandType) IsUpdateExpense() bool {
    return o == _expenseCommandTypes.UpdateExpense()
}

type expenseCommandTypes struct {
	values []*ExpenseCommandType
    literals []enum.Literal
}

var _expenseCommandTypes = &expenseCommandTypes{values: []*ExpenseCommandType{
    {name: "createExpense", ordinal: 0},
    {name: "deleteExpense", ordinal: 1},
    {name: "updateExpense", ordinal: 2}},
}

func ExpenseCommandTypes() *expenseCommandTypes {
	return _expenseCommandTypes
}

func (o *expenseCommandTypes) Values() []*ExpenseCommandType {
	return o.values
}

func (o *expenseCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *expenseCommandTypes) CreateExpense() *ExpenseCommandType {
    return _expenseCommandTypes.values[0]
}

func (o *expenseCommandTypes) DeleteExpense() *ExpenseCommandType {
    return _expenseCommandTypes.values[1]
}

func (o *expenseCommandTypes) UpdateExpense() *ExpenseCommandType {
    return _expenseCommandTypes.values[2]
}

func (o *expenseCommandTypes) ParseExpenseCommandType(name string) (ret *ExpenseCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*ExpenseCommandType), ok
	}
	return
}


type ExpensePurposeCommandType struct {
	name  string
	ordinal int
}

func (o *ExpensePurposeCommandType) Name() string {
    return o.name
}

func (o *ExpensePurposeCommandType) Ordinal() int {
    return o.ordinal
}

func (o *ExpensePurposeCommandType) IsCreateExpensePurpose() bool {
    return o == _expensePurposeCommandTypes.CreateExpensePurpose()
}

func (o *ExpensePurposeCommandType) IsDeleteExpensePurpose() bool {
    return o == _expensePurposeCommandTypes.DeleteExpensePurpose()
}

func (o *ExpensePurposeCommandType) IsUpdateExpensePurpose() bool {
    return o == _expensePurposeCommandTypes.UpdateExpensePurpose()
}

type expensePurposeCommandTypes struct {
	values []*ExpensePurposeCommandType
    literals []enum.Literal
}

var _expensePurposeCommandTypes = &expensePurposeCommandTypes{values: []*ExpensePurposeCommandType{
    {name: "createExpensePurpose", ordinal: 0},
    {name: "deleteExpensePurpose", ordinal: 1},
    {name: "updateExpensePurpose", ordinal: 2}},
}

func ExpensePurposeCommandTypes() *expensePurposeCommandTypes {
	return _expensePurposeCommandTypes
}

func (o *expensePurposeCommandTypes) Values() []*ExpensePurposeCommandType {
	return o.values
}

func (o *expensePurposeCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *expensePurposeCommandTypes) CreateExpensePurpose() *ExpensePurposeCommandType {
    return _expensePurposeCommandTypes.values[0]
}

func (o *expensePurposeCommandTypes) DeleteExpensePurpose() *ExpensePurposeCommandType {
    return _expensePurposeCommandTypes.values[1]
}

func (o *expensePurposeCommandTypes) UpdateExpensePurpose() *ExpensePurposeCommandType {
    return _expensePurposeCommandTypes.values[2]
}

func (o *expensePurposeCommandTypes) ParseExpensePurposeCommandType(name string) (ret *ExpensePurposeCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*ExpensePurposeCommandType), ok
	}
	return
}


type FeeCommandType struct {
	name  string
	ordinal int
}

func (o *FeeCommandType) Name() string {
    return o.name
}

func (o *FeeCommandType) Ordinal() int {
    return o.ordinal
}

func (o *FeeCommandType) IsCreateFee() bool {
    return o == _feeCommandTypes.CreateFee()
}

func (o *FeeCommandType) IsDeleteFee() bool {
    return o == _feeCommandTypes.DeleteFee()
}

func (o *FeeCommandType) IsUpdateFee() bool {
    return o == _feeCommandTypes.UpdateFee()
}

type feeCommandTypes struct {
	values []*FeeCommandType
    literals []enum.Literal
}

var _feeCommandTypes = &feeCommandTypes{values: []*FeeCommandType{
    {name: "createFee", ordinal: 0},
    {name: "deleteFee", ordinal: 1},
    {name: "updateFee", ordinal: 2}},
}

func FeeCommandTypes() *feeCommandTypes {
	return _feeCommandTypes
}

func (o *feeCommandTypes) Values() []*FeeCommandType {
	return o.values
}

func (o *feeCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *feeCommandTypes) CreateFee() *FeeCommandType {
    return _feeCommandTypes.values[0]
}

func (o *feeCommandTypes) DeleteFee() *FeeCommandType {
    return _feeCommandTypes.values[1]
}

func (o *feeCommandTypes) UpdateFee() *FeeCommandType {
    return _feeCommandTypes.values[2]
}

func (o *feeCommandTypes) ParseFeeCommandType(name string) (ret *FeeCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*FeeCommandType), ok
	}
	return
}


type FeeKindCommandType struct {
	name  string
	ordinal int
}

func (o *FeeKindCommandType) Name() string {
    return o.name
}

func (o *FeeKindCommandType) Ordinal() int {
    return o.ordinal
}

func (o *FeeKindCommandType) IsCreateFeeKind() bool {
    return o == _feeKindCommandTypes.CreateFeeKind()
}

func (o *FeeKindCommandType) IsDeleteFeeKind() bool {
    return o == _feeKindCommandTypes.DeleteFeeKind()
}

func (o *FeeKindCommandType) IsUpdateFeeKind() bool {
    return o == _feeKindCommandTypes.UpdateFeeKind()
}

type feeKindCommandTypes struct {
	values []*FeeKindCommandType
    literals []enum.Literal
}

var _feeKindCommandTypes = &feeKindCommandTypes{values: []*FeeKindCommandType{
    {name: "createFeeKind", ordinal: 0},
    {name: "deleteFeeKind", ordinal: 1},
    {name: "updateFeeKind", ordinal: 2}},
}

func FeeKindCommandTypes() *feeKindCommandTypes {
	return _feeKindCommandTypes
}

func (o *feeKindCommandTypes) Values() []*FeeKindCommandType {
	return o.values
}

func (o *feeKindCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *feeKindCommandTypes) CreateFeeKind() *FeeKindCommandType {
    return _feeKindCommandTypes.values[0]
}

func (o *feeKindCommandTypes) DeleteFeeKind() *FeeKindCommandType {
    return _feeKindCommandTypes.values[1]
}

func (o *feeKindCommandTypes) UpdateFeeKind() *FeeKindCommandType {
    return _feeKindCommandTypes.values[2]
}

func (o *feeKindCommandTypes) ParseFeeKindCommandType(name string) (ret *FeeKindCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*FeeKindCommandType), ok
	}
	return
}



