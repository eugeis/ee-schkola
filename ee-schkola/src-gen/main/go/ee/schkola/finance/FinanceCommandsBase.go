package finance

import (
    "ee/schkola"
    "ee/schkola/person"
    "time"
)
type CreateExpense struct {
    Purpose  *ExpensePurpose
    Amount  float64
    Profile  *person.Profile
    Date  *time.Time
    *schkola.SchkolaBase
}


type DeleteExpense struct {
    Id  string
}


type UpdateExpense struct {
    Purpose  *ExpensePurpose
    Amount  float64
    Profile  *person.Profile
    Date  *time.Time
    *schkola.SchkolaBase
}


type CreateExpensePurpose struct {
    Name  string
    Description  string
    *schkola.SchkolaBase
}


type DeleteExpensePurpose struct {
    Id  string
}


type UpdateExpensePurpose struct {
    Name  string
    Description  string
    *schkola.SchkolaBase
}


type CreateFee struct {
    Student  *person.Profile
    Amount  float64
    Kind  *FeeKind
    Date  *time.Time
    *schkola.SchkolaBase
}


type DeleteFee struct {
    Id  string
}


type UpdateFee struct {
    Student  *person.Profile
    Amount  float64
    Kind  *FeeKind
    Date  *time.Time
    *schkola.SchkolaBase
}


type CreateFeeKind struct {
    Name  string
    Amount  float64
    Description  string
    *schkola.SchkolaBase
}


type DeleteFeeKind struct {
    Id  string
}


type UpdateFeeKind struct {
    Name  string
    Amount  float64
    Description  string
    *schkola.SchkolaBase
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

func (o *ExpenseCommandType) IsExpenseCreate() bool {
    return o == _expenseCommandTypes.ExpenseCreate()
}

func (o *ExpenseCommandType) IsExpenseDelete() bool {
    return o == _expenseCommandTypes.ExpenseDelete()
}

func (o *ExpenseCommandType) IsExpenseUpdate() bool {
    return o == _expenseCommandTypes.ExpenseUpdate()
}

type expenseCommandTypes struct {
	values []*ExpenseCommandType
}

var _expenseCommandTypes = &expenseCommandTypes{values: []*ExpenseCommandType{
    {name: "ExpenseCreate", ordinal: 0},
    {name: "ExpenseDelete", ordinal: 1},
    {name: "ExpenseUpdate", ordinal: 2}},
}

func ExpenseCommandTypes() *expenseCommandTypes {
	return _expenseCommandTypes
}

func (o *expenseCommandTypes) Values() []*ExpenseCommandType {
	return o.values
}

func (o *expenseCommandTypes) ExpenseCreate() *ExpenseCommandType {
    return _expenseCommandTypes.values[0]
}

func (o *expenseCommandTypes) ExpenseDelete() *ExpenseCommandType {
    return _expenseCommandTypes.values[1]
}

func (o *expenseCommandTypes) ExpenseUpdate() *ExpenseCommandType {
    return _expenseCommandTypes.values[2]
}

func (o *expenseCommandTypes) ParseExpenseCommandType(name string) (ret *ExpenseCommandType, ok bool) {
    switch name {
      case "ExpenseCreate":
        ret = o.ExpenseCreate()
      case "ExpenseDelete":
        ret = o.ExpenseDelete()
      case "ExpenseUpdate":
        ret = o.ExpenseUpdate()
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

func (o *ExpensePurposeCommandType) IsExpensePurposeCreate() bool {
    return o == _expensePurposeCommandTypes.ExpensePurposeCreate()
}

func (o *ExpensePurposeCommandType) IsExpensePurposeDelete() bool {
    return o == _expensePurposeCommandTypes.ExpensePurposeDelete()
}

func (o *ExpensePurposeCommandType) IsExpensePurposeUpdate() bool {
    return o == _expensePurposeCommandTypes.ExpensePurposeUpdate()
}

type expensePurposeCommandTypes struct {
	values []*ExpensePurposeCommandType
}

var _expensePurposeCommandTypes = &expensePurposeCommandTypes{values: []*ExpensePurposeCommandType{
    {name: "ExpensePurposeCreate", ordinal: 0},
    {name: "ExpensePurposeDelete", ordinal: 1},
    {name: "ExpensePurposeUpdate", ordinal: 2}},
}

func ExpensePurposeCommandTypes() *expensePurposeCommandTypes {
	return _expensePurposeCommandTypes
}

func (o *expensePurposeCommandTypes) Values() []*ExpensePurposeCommandType {
	return o.values
}

func (o *expensePurposeCommandTypes) ExpensePurposeCreate() *ExpensePurposeCommandType {
    return _expensePurposeCommandTypes.values[0]
}

func (o *expensePurposeCommandTypes) ExpensePurposeDelete() *ExpensePurposeCommandType {
    return _expensePurposeCommandTypes.values[1]
}

func (o *expensePurposeCommandTypes) ExpensePurposeUpdate() *ExpensePurposeCommandType {
    return _expensePurposeCommandTypes.values[2]
}

func (o *expensePurposeCommandTypes) ParseExpensePurposeCommandType(name string) (ret *ExpensePurposeCommandType, ok bool) {
    switch name {
      case "ExpensePurposeCreate":
        ret = o.ExpensePurposeCreate()
      case "ExpensePurposeDelete":
        ret = o.ExpensePurposeDelete()
      case "ExpensePurposeUpdate":
        ret = o.ExpensePurposeUpdate()
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

func (o *FeeCommandType) IsFeeCreate() bool {
    return o == _feeCommandTypes.FeeCreate()
}

func (o *FeeCommandType) IsFeeDelete() bool {
    return o == _feeCommandTypes.FeeDelete()
}

func (o *FeeCommandType) IsFeeUpdate() bool {
    return o == _feeCommandTypes.FeeUpdate()
}

type feeCommandTypes struct {
	values []*FeeCommandType
}

var _feeCommandTypes = &feeCommandTypes{values: []*FeeCommandType{
    {name: "FeeCreate", ordinal: 0},
    {name: "FeeDelete", ordinal: 1},
    {name: "FeeUpdate", ordinal: 2}},
}

func FeeCommandTypes() *feeCommandTypes {
	return _feeCommandTypes
}

func (o *feeCommandTypes) Values() []*FeeCommandType {
	return o.values
}

func (o *feeCommandTypes) FeeCreate() *FeeCommandType {
    return _feeCommandTypes.values[0]
}

func (o *feeCommandTypes) FeeDelete() *FeeCommandType {
    return _feeCommandTypes.values[1]
}

func (o *feeCommandTypes) FeeUpdate() *FeeCommandType {
    return _feeCommandTypes.values[2]
}

func (o *feeCommandTypes) ParseFeeCommandType(name string) (ret *FeeCommandType, ok bool) {
    switch name {
      case "FeeCreate":
        ret = o.FeeCreate()
      case "FeeDelete":
        ret = o.FeeDelete()
      case "FeeUpdate":
        ret = o.FeeUpdate()
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

func (o *FeeKindCommandType) IsFeeKindCreate() bool {
    return o == _feeKindCommandTypes.FeeKindCreate()
}

func (o *FeeKindCommandType) IsFeeKindDelete() bool {
    return o == _feeKindCommandTypes.FeeKindDelete()
}

func (o *FeeKindCommandType) IsFeeKindUpdate() bool {
    return o == _feeKindCommandTypes.FeeKindUpdate()
}

type feeKindCommandTypes struct {
	values []*FeeKindCommandType
}

var _feeKindCommandTypes = &feeKindCommandTypes{values: []*FeeKindCommandType{
    {name: "FeeKindCreate", ordinal: 0},
    {name: "FeeKindDelete", ordinal: 1},
    {name: "FeeKindUpdate", ordinal: 2}},
}

func FeeKindCommandTypes() *feeKindCommandTypes {
	return _feeKindCommandTypes
}

func (o *feeKindCommandTypes) Values() []*FeeKindCommandType {
	return o.values
}

func (o *feeKindCommandTypes) FeeKindCreate() *FeeKindCommandType {
    return _feeKindCommandTypes.values[0]
}

func (o *feeKindCommandTypes) FeeKindDelete() *FeeKindCommandType {
    return _feeKindCommandTypes.values[1]
}

func (o *feeKindCommandTypes) FeeKindUpdate() *FeeKindCommandType {
    return _feeKindCommandTypes.values[2]
}

func (o *feeKindCommandTypes) ParseFeeKindCommandType(name string) (ret *FeeKindCommandType, ok bool) {
    switch name {
      case "FeeKindCreate":
        ret = o.FeeKindCreate()
      case "FeeKindDelete":
        ret = o.FeeKindDelete()
      case "FeeKindUpdate":
        ret = o.FeeKindUpdate()
    }
    return
}



