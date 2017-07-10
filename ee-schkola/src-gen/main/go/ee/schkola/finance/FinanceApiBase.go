package finance

import (
    "ee/schkola"
    "ee/schkola/person"
    "time"
)
type ExpenseCommands struct {
	name  string
	ordinal int
}

func (o *ExpenseCommands) Name() string {
    return o.name
}

func (o *ExpenseCommands) Ordinal() int {
    return o.ordinal
}

func (o *ExpenseCommands) IsRegisterExpense() bool {
    return o == _expenseCommandss.RegisterExpense()
}

func (o *ExpenseCommands) IsDeleteExpense() bool {
    return o == _expenseCommandss.DeleteExpense()
}

func (o *ExpenseCommands) IsChangeExpense() bool {
    return o == _expenseCommandss.ChangeExpense()
}

type expenseCommandss struct {
	values []*ExpenseCommands
}

var _expenseCommandss = &expenseCommandss{values: []*ExpenseCommands{
    {name: "registerExpense", ordinal: 0},
    {name: "deleteExpense", ordinal: 1},
    {name: "changeExpense", ordinal: 2}},
}

func ExpenseCommandss() *expenseCommandss {
	return _expenseCommandss
}

func (o *expenseCommandss) Values() []*ExpenseCommands {
	return o.values
}

func (o *expenseCommandss) RegisterExpense() *ExpenseCommands {
    return _expenseCommandss.values[0]
}

func (o *expenseCommandss) DeleteExpense() *ExpenseCommands {
    return _expenseCommandss.values[1]
}

func (o *expenseCommandss) ChangeExpense() *ExpenseCommands {
    return _expenseCommandss.values[2]
}

func (o *expenseCommandss) ParseExpenseCommands(name string) (ret *ExpenseCommands, ok bool) {
    switch name {
      case "RegisterExpense":
        ret = o.RegisterExpense()
      case "DeleteExpense":
        ret = o.DeleteExpense()
      case "ChangeExpense":
        ret = o.ChangeExpense()
    }
    return
}


type ExpensePurposeCommands struct {
	name  string
	ordinal int
}

func (o *ExpensePurposeCommands) Name() string {
    return o.name
}

func (o *ExpensePurposeCommands) Ordinal() int {
    return o.ordinal
}

func (o *ExpensePurposeCommands) IsRegisterExpensePurpose() bool {
    return o == _expensePurposeCommandss.RegisterExpensePurpose()
}

func (o *ExpensePurposeCommands) IsDeleteExpensePurpose() bool {
    return o == _expensePurposeCommandss.DeleteExpensePurpose()
}

func (o *ExpensePurposeCommands) IsChangeExpensePurpose() bool {
    return o == _expensePurposeCommandss.ChangeExpensePurpose()
}

type expensePurposeCommandss struct {
	values []*ExpensePurposeCommands
}

var _expensePurposeCommandss = &expensePurposeCommandss{values: []*ExpensePurposeCommands{
    {name: "registerExpensePurpose", ordinal: 0},
    {name: "deleteExpensePurpose", ordinal: 1},
    {name: "changeExpensePurpose", ordinal: 2}},
}

func ExpensePurposeCommandss() *expensePurposeCommandss {
	return _expensePurposeCommandss
}

func (o *expensePurposeCommandss) Values() []*ExpensePurposeCommands {
	return o.values
}

func (o *expensePurposeCommandss) RegisterExpensePurpose() *ExpensePurposeCommands {
    return _expensePurposeCommandss.values[0]
}

func (o *expensePurposeCommandss) DeleteExpensePurpose() *ExpensePurposeCommands {
    return _expensePurposeCommandss.values[1]
}

func (o *expensePurposeCommandss) ChangeExpensePurpose() *ExpensePurposeCommands {
    return _expensePurposeCommandss.values[2]
}

func (o *expensePurposeCommandss) ParseExpensePurposeCommands(name string) (ret *ExpensePurposeCommands, ok bool) {
    switch name {
      case "RegisterExpensePurpose":
        ret = o.RegisterExpensePurpose()
      case "DeleteExpensePurpose":
        ret = o.DeleteExpensePurpose()
      case "ChangeExpensePurpose":
        ret = o.ChangeExpensePurpose()
    }
    return
}


type FeeCommands struct {
	name  string
	ordinal int
}

func (o *FeeCommands) Name() string {
    return o.name
}

func (o *FeeCommands) Ordinal() int {
    return o.ordinal
}

func (o *FeeCommands) IsRegisterFee() bool {
    return o == _feeCommandss.RegisterFee()
}

func (o *FeeCommands) IsDeleteFee() bool {
    return o == _feeCommandss.DeleteFee()
}

func (o *FeeCommands) IsChangeFee() bool {
    return o == _feeCommandss.ChangeFee()
}

type feeCommandss struct {
	values []*FeeCommands
}

var _feeCommandss = &feeCommandss{values: []*FeeCommands{
    {name: "registerFee", ordinal: 0},
    {name: "deleteFee", ordinal: 1},
    {name: "changeFee", ordinal: 2}},
}

func FeeCommandss() *feeCommandss {
	return _feeCommandss
}

func (o *feeCommandss) Values() []*FeeCommands {
	return o.values
}

func (o *feeCommandss) RegisterFee() *FeeCommands {
    return _feeCommandss.values[0]
}

func (o *feeCommandss) DeleteFee() *FeeCommands {
    return _feeCommandss.values[1]
}

func (o *feeCommandss) ChangeFee() *FeeCommands {
    return _feeCommandss.values[2]
}

func (o *feeCommandss) ParseFeeCommands(name string) (ret *FeeCommands, ok bool) {
    switch name {
      case "RegisterFee":
        ret = o.RegisterFee()
      case "DeleteFee":
        ret = o.DeleteFee()
      case "ChangeFee":
        ret = o.ChangeFee()
    }
    return
}


type FeeKindCommands struct {
	name  string
	ordinal int
}

func (o *FeeKindCommands) Name() string {
    return o.name
}

func (o *FeeKindCommands) Ordinal() int {
    return o.ordinal
}

func (o *FeeKindCommands) IsRegisterFeeKind() bool {
    return o == _feeKindCommandss.RegisterFeeKind()
}

func (o *FeeKindCommands) IsDeleteFeeKind() bool {
    return o == _feeKindCommandss.DeleteFeeKind()
}

func (o *FeeKindCommands) IsChangeFeeKind() bool {
    return o == _feeKindCommandss.ChangeFeeKind()
}

type feeKindCommandss struct {
	values []*FeeKindCommands
}

var _feeKindCommandss = &feeKindCommandss{values: []*FeeKindCommands{
    {name: "registerFeeKind", ordinal: 0},
    {name: "deleteFeeKind", ordinal: 1},
    {name: "changeFeeKind", ordinal: 2}},
}

func FeeKindCommandss() *feeKindCommandss {
	return _feeKindCommandss
}

func (o *feeKindCommandss) Values() []*FeeKindCommands {
	return o.values
}

func (o *feeKindCommandss) RegisterFeeKind() *FeeKindCommands {
    return _feeKindCommandss.values[0]
}

func (o *feeKindCommandss) DeleteFeeKind() *FeeKindCommands {
    return _feeKindCommandss.values[1]
}

func (o *feeKindCommandss) ChangeFeeKind() *FeeKindCommands {
    return _feeKindCommandss.values[2]
}

func (o *feeKindCommandss) ParseFeeKindCommands(name string) (ret *FeeKindCommands, ok bool) {
    switch name {
      case "RegisterFeeKind":
        ret = o.RegisterFeeKind()
      case "DeleteFeeKind":
        ret = o.DeleteFeeKind()
      case "ChangeFeeKind":
        ret = o.ChangeFeeKind()
    }
    return
}




type Expense struct {
    Purpose  *ExpensePurpose
    Amount  float64
    Profile  *person.Profile
    Date  *time.Time
    *schkola.SchkolaBase
}

func NewExpense(purpose *ExpensePurpose, amount float64, profile *person.Profile, date *time.Time, SchkolaBase *schkola.SchkolaBase) (ret *Expense, err error) {
    ret = &Expense{
        Purpose : purpose,
        Amount : amount,
        Profile : profile,
        Date : date,
        SchkolaBase : SchkolaBase,
    }
    return
}


type ExpensePurpose struct {
    Name  string
    Description  string
    *schkola.SchkolaBase
}

func NewExpensePurpose(name string, description string, SchkolaBase *schkola.SchkolaBase) (ret *ExpensePurpose, err error) {
    ret = &ExpensePurpose{
        Name : name,
        Description : description,
        SchkolaBase : SchkolaBase,
    }
    return
}


type Fee struct {
    Student  *person.Profile
    Amount  float64
    Kind  *FeeKind
    Date  *time.Time
    *schkola.SchkolaBase
}

func NewFee(student *person.Profile, amount float64, kind *FeeKind, date *time.Time, SchkolaBase *schkola.SchkolaBase) (ret *Fee, err error) {
    ret = &Fee{
        Student : student,
        Amount : amount,
        Kind : kind,
        Date : date,
        SchkolaBase : SchkolaBase,
    }
    return
}


type FeeKind struct {
    Name  string
    Amount  float64
    Description  string
    *schkola.SchkolaBase
}

func NewFeeKind(name string, amount float64, description string, SchkolaBase *schkola.SchkolaBase) (ret *FeeKind, err error) {
    ret = &FeeKind{
        Name : name,
        Amount : amount,
        Description : description,
        SchkolaBase : SchkolaBase,
    }
    return
}



