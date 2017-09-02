package finance

import (
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "github.com/looplab/eventhorizon"
    "time"
)
const (
     ExpenseCreatedEvent eventhorizon.EventType = "ExpenseCreated"
     ExpenseDeletedEvent eventhorizon.EventType = "ExpenseDeleted"
     ExpenseUpdatedEvent eventhorizon.EventType = "ExpenseUpdated"
)


const (
     ExpensePurposeCreatedEvent eventhorizon.EventType = "ExpensePurposeCreated"
     ExpensePurposeDeletedEvent eventhorizon.EventType = "ExpensePurposeDeleted"
     ExpensePurposeUpdatedEvent eventhorizon.EventType = "ExpensePurposeUpdated"
)


const (
     FeeCreatedEvent eventhorizon.EventType = "FeeCreated"
     FeeDeletedEvent eventhorizon.EventType = "FeeDeleted"
     FeeUpdatedEvent eventhorizon.EventType = "FeeUpdated"
)


const (
     FeeKindCreatedEvent eventhorizon.EventType = "FeeKindCreated"
     FeeKindDeletedEvent eventhorizon.EventType = "FeeKindDeleted"
     FeeKindUpdatedEvent eventhorizon.EventType = "FeeKindUpdated"
)




type ExpenseCreated struct {
    Purpose *ExpensePurpose`eh:"optional"`
    Amount float64`eh:"optional"`
    Profile *person.Profile`eh:"optional"`
    Date *time.Time`eh:"optional"`
    Id eventhorizon.UUID`eh:"optional"`
}


type ExpenseDeleted struct {
    Id eventhorizon.UUID
}


type ExpenseUpdated struct {
    Purpose *ExpensePurpose`eh:"optional"`
    Amount float64`eh:"optional"`
    Profile *person.Profile`eh:"optional"`
    Date *time.Time`eh:"optional"`
    Id eventhorizon.UUID`eh:"optional"`
}


type ExpensePurposeCreated struct {
    Name string`eh:"optional"`
    Description string`eh:"optional"`
    Id eventhorizon.UUID`eh:"optional"`
}


type ExpensePurposeDeleted struct {
    Id eventhorizon.UUID
}


type ExpensePurposeUpdated struct {
    Name string`eh:"optional"`
    Description string`eh:"optional"`
    Id eventhorizon.UUID`eh:"optional"`
}


type FeeCreated struct {
    Student *person.Profile`eh:"optional"`
    Amount float64`eh:"optional"`
    Kind *FeeKind`eh:"optional"`
    Date *time.Time`eh:"optional"`
    Id eventhorizon.UUID`eh:"optional"`
}


type FeeDeleted struct {
    Id eventhorizon.UUID
}


type FeeUpdated struct {
    Student *person.Profile`eh:"optional"`
    Amount float64`eh:"optional"`
    Kind *FeeKind`eh:"optional"`
    Date *time.Time`eh:"optional"`
    Id eventhorizon.UUID`eh:"optional"`
}


type FeeKindCreated struct {
    Name string`eh:"optional"`
    Amount float64`eh:"optional"`
    Description string`eh:"optional"`
    Id eventhorizon.UUID`eh:"optional"`
}


type FeeKindDeleted struct {
    Id eventhorizon.UUID
}


type FeeKindUpdated struct {
    Name string`eh:"optional"`
    Amount float64`eh:"optional"`
    Description string`eh:"optional"`
    Id eventhorizon.UUID`eh:"optional"`
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



