package finance

import (
    "ee/schkola/person"
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/enum"
    "time"
)
const (
     CreateExpenseCommand eventhorizon.CommandType = "CreateExpense"
     DeleteExpenseCommand eventhorizon.CommandType = "DeleteExpense"
     UpdateExpenseCommand eventhorizon.CommandType = "UpdateExpense"
)


const (
     CreateExpensePurposeCommand eventhorizon.CommandType = "CreateExpensePurpose"
     DeleteExpensePurposeCommand eventhorizon.CommandType = "DeleteExpensePurpose"
     UpdateExpensePurposeCommand eventhorizon.CommandType = "UpdateExpensePurpose"
)


const (
     CreateFeeCommand eventhorizon.CommandType = "CreateFee"
     DeleteFeeCommand eventhorizon.CommandType = "DeleteFee"
     UpdateFeeCommand eventhorizon.CommandType = "UpdateFee"
)


const (
     CreateFeeKindCommand eventhorizon.CommandType = "CreateFeeKind"
     DeleteFeeKindCommand eventhorizon.CommandType = "DeleteFeeKind"
     UpdateFeeKindCommand eventhorizon.CommandType = "UpdateFeeKind"
)




        

type CreateExpense struct {
    Id  eventhorizon.UUID
    Purpose  *ExpensePurpose
    Amount  float64
    Profile  *person.Profile
    Date  *time.Time
}
func (o *CreateExpense) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *CreateExpense) AggregateType() eventhorizon.AggregateType  { return ExpenseAggregateType }
func (o *CreateExpense) CommandType() eventhorizon.CommandType      { return CreateExpenseCommand }



        

type DeleteExpense struct {
    Id  eventhorizon.UUID
}
func (o *DeleteExpense) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *DeleteExpense) AggregateType() eventhorizon.AggregateType  { return ExpenseAggregateType }
func (o *DeleteExpense) CommandType() eventhorizon.CommandType      { return DeleteExpenseCommand }



        

type UpdateExpense struct {
    Id  eventhorizon.UUID
    Purpose  *ExpensePurpose
    Amount  float64
    Profile  *person.Profile
    Date  *time.Time
}
func (o *UpdateExpense) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *UpdateExpense) AggregateType() eventhorizon.AggregateType  { return ExpenseAggregateType }
func (o *UpdateExpense) CommandType() eventhorizon.CommandType      { return UpdateExpenseCommand }



        

type CreateExpensePurpose struct {
    Id  eventhorizon.UUID
    Name  string
    Description  string
}
func (o *CreateExpensePurpose) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *CreateExpensePurpose) AggregateType() eventhorizon.AggregateType  { return ExpensePurposeAggregateType }
func (o *CreateExpensePurpose) CommandType() eventhorizon.CommandType      { return CreateExpensePurposeCommand }



        

type DeleteExpensePurpose struct {
    Id  eventhorizon.UUID
}
func (o *DeleteExpensePurpose) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *DeleteExpensePurpose) AggregateType() eventhorizon.AggregateType  { return ExpensePurposeAggregateType }
func (o *DeleteExpensePurpose) CommandType() eventhorizon.CommandType      { return DeleteExpensePurposeCommand }



        

type UpdateExpensePurpose struct {
    Id  eventhorizon.UUID
    Name  string
    Description  string
}
func (o *UpdateExpensePurpose) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *UpdateExpensePurpose) AggregateType() eventhorizon.AggregateType  { return ExpensePurposeAggregateType }
func (o *UpdateExpensePurpose) CommandType() eventhorizon.CommandType      { return UpdateExpensePurposeCommand }



        

type CreateFee struct {
    Id  eventhorizon.UUID
    Student  *person.Profile
    Amount  float64
    Kind  *FeeKind
    Date  *time.Time
}
func (o *CreateFee) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *CreateFee) AggregateType() eventhorizon.AggregateType  { return FeeAggregateType }
func (o *CreateFee) CommandType() eventhorizon.CommandType      { return CreateFeeCommand }



        

type DeleteFee struct {
    Id  eventhorizon.UUID
}
func (o *DeleteFee) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *DeleteFee) AggregateType() eventhorizon.AggregateType  { return FeeAggregateType }
func (o *DeleteFee) CommandType() eventhorizon.CommandType      { return DeleteFeeCommand }



        

type UpdateFee struct {
    Id  eventhorizon.UUID
    Student  *person.Profile
    Amount  float64
    Kind  *FeeKind
    Date  *time.Time
}
func (o *UpdateFee) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *UpdateFee) AggregateType() eventhorizon.AggregateType  { return FeeAggregateType }
func (o *UpdateFee) CommandType() eventhorizon.CommandType      { return UpdateFeeCommand }



        

type CreateFeeKind struct {
    Id  eventhorizon.UUID
    Name  string
    Amount  float64
    Description  string
}
func (o *CreateFeeKind) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *CreateFeeKind) AggregateType() eventhorizon.AggregateType  { return FeeKindAggregateType }
func (o *CreateFeeKind) CommandType() eventhorizon.CommandType      { return CreateFeeKindCommand }



        

type DeleteFeeKind struct {
    Id  eventhorizon.UUID
}
func (o *DeleteFeeKind) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *DeleteFeeKind) AggregateType() eventhorizon.AggregateType  { return FeeKindAggregateType }
func (o *DeleteFeeKind) CommandType() eventhorizon.CommandType      { return DeleteFeeKindCommand }



        

type UpdateFeeKind struct {
    Id  eventhorizon.UUID
    Name  string
    Amount  float64
    Description  string
}
func (o *UpdateFeeKind) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *UpdateFeeKind) AggregateType() eventhorizon.AggregateType  { return FeeKindAggregateType }
func (o *UpdateFeeKind) CommandType() eventhorizon.CommandType      { return UpdateFeeKindCommand }





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



