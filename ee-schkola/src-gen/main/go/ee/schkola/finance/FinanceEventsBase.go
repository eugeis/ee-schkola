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
    Id  string
}


type ExpenseUpdated struct {
    Purpose  *ExpensePurpose
    Amount  float64
    Profile  *person.Profile
    Date  *time.Time
    *schkola.SchkolaBase
}


type ExpensePurposeCreated struct {
    Name  string
    Description  string
    *schkola.SchkolaBase
}


type ExpensePurposeDeleted struct {
    Id  string
}


type ExpensePurposeUpdated struct {
    Name  string
    Description  string
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
    Id  string
}


type FeeUpdated struct {
    Student  *person.Profile
    Amount  float64
    Kind  *FeeKind
    Date  *time.Time
    *schkola.SchkolaBase
}


type FeeKindCreated struct {
    Name  string
    Amount  float64
    Description  string
    *schkola.SchkolaBase
}


type FeeKindDeleted struct {
    Id  string
}


type FeeKindUpdated struct {
    Name  string
    Amount  float64
    Description  string
    *schkola.SchkolaBase
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
    switch name {
      case "CreatedExpense":
        ret = o.CreatedExpense()
      case "DeletedExpense":
        ret = o.DeletedExpense()
      case "UpdatedExpense":
        ret = o.UpdatedExpense()
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
    switch name {
      case "CreatedExpensePurpose":
        ret = o.CreatedExpensePurpose()
      case "DeletedExpensePurpose":
        ret = o.DeletedExpensePurpose()
      case "UpdatedExpensePurpose":
        ret = o.UpdatedExpensePurpose()
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
    switch name {
      case "CreatedFee":
        ret = o.CreatedFee()
      case "DeletedFee":
        ret = o.DeletedFee()
      case "UpdatedFee":
        ret = o.UpdatedFee()
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
    switch name {
      case "CreatedFeeKind":
        ret = o.CreatedFeeKind()
      case "DeletedFeeKind":
        ret = o.DeletedFeeKind()
      case "UpdatedFeeKind":
        ret = o.UpdatedFeeKind()
    }
    return
}



