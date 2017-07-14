package finance

import (
    "ee/schkola/person"
    "time"
)
type CreateExpense struct {
    Purpose  *ExpensePurpose
    Amount  float64
    Profile  *person.Profile
    Date  *time.Time
}

func NewCreateExpense(purpose *ExpensePurpose, amount float64, profile *person.Profile, date *time.Time) (ret *CreateExpense, err error) {
    ret = &CreateExpense{
        Purpose : purpose,
        Amount : amount,
        Profile : profile,
        Date : date,
    }
    return
}


type DeleteExpense struct {
    Id  string
}

func NewDeleteExpense(id string) (ret *DeleteExpense, err error) {
    ret = &DeleteExpense{
        Id : id,
    }
    return
}


type UpdateExpense struct {
    Purpose  *ExpensePurpose
    Amount  float64
    Profile  *person.Profile
    Date  *time.Time
}

func NewUpdateExpense(purpose *ExpensePurpose, amount float64, profile *person.Profile, date *time.Time) (ret *UpdateExpense, err error) {
    ret = &UpdateExpense{
        Purpose : purpose,
        Amount : amount,
        Profile : profile,
        Date : date,
    }
    return
}


type CreateExpensePurpose struct {
    Name  string
    Description  string
}

func NewCreateExpensePurpose(name string, description string) (ret *CreateExpensePurpose, err error) {
    ret = &CreateExpensePurpose{
        Name : name,
        Description : description,
    }
    return
}


type DeleteExpensePurpose struct {
    Id  string
}

func NewDeleteExpensePurpose(id string) (ret *DeleteExpensePurpose, err error) {
    ret = &DeleteExpensePurpose{
        Id : id,
    }
    return
}


type UpdateExpensePurpose struct {
    Name  string
    Description  string
}

func NewUpdateExpensePurpose(name string, description string) (ret *UpdateExpensePurpose, err error) {
    ret = &UpdateExpensePurpose{
        Name : name,
        Description : description,
    }
    return
}


type CreateFee struct {
    Student  *person.Profile
    Amount  float64
    Kind  *FeeKind
    Date  *time.Time
}

func NewCreateFee(student *person.Profile, amount float64, kind *FeeKind, date *time.Time) (ret *CreateFee, err error) {
    ret = &CreateFee{
        Student : student,
        Amount : amount,
        Kind : kind,
        Date : date,
    }
    return
}


type DeleteFee struct {
    Id  string
}

func NewDeleteFee(id string) (ret *DeleteFee, err error) {
    ret = &DeleteFee{
        Id : id,
    }
    return
}


type UpdateFee struct {
    Student  *person.Profile
    Amount  float64
    Kind  *FeeKind
    Date  *time.Time
}

func NewUpdateFee(student *person.Profile, amount float64, kind *FeeKind, date *time.Time) (ret *UpdateFee, err error) {
    ret = &UpdateFee{
        Student : student,
        Amount : amount,
        Kind : kind,
        Date : date,
    }
    return
}


type CreateFeeKind struct {
    Name  string
    Amount  float64
    Description  string
}

func NewCreateFeeKind(name string, amount float64, description string) (ret *CreateFeeKind, err error) {
    ret = &CreateFeeKind{
        Name : name,
        Amount : amount,
        Description : description,
    }
    return
}


type DeleteFeeKind struct {
    Id  string
}

func NewDeleteFeeKind(id string) (ret *DeleteFeeKind, err error) {
    ret = &DeleteFeeKind{
        Id : id,
    }
    return
}


type UpdateFeeKind struct {
    Name  string
    Amount  float64
    Description  string
}

func NewUpdateFeeKind(name string, amount float64, description string) (ret *UpdateFeeKind, err error) {
    ret = &UpdateFeeKind{
        Name : name,
        Amount : amount,
        Description : description,
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
      case o.ExpenseCreate().Name():
        ret = o.ExpenseCreate()
      case o.ExpenseDelete().Name():
        ret = o.ExpenseDelete()
      case o.ExpenseUpdate().Name():
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
      case o.ExpensePurposeCreate().Name():
        ret = o.ExpensePurposeCreate()
      case o.ExpensePurposeDelete().Name():
        ret = o.ExpensePurposeDelete()
      case o.ExpensePurposeUpdate().Name():
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
      case o.FeeCreate().Name():
        ret = o.FeeCreate()
      case o.FeeDelete().Name():
        ret = o.FeeDelete()
      case o.FeeUpdate().Name():
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
      case o.FeeKindCreate().Name():
        ret = o.FeeKindCreate()
      case o.FeeKindDelete().Name():
        ret = o.FeeKindDelete()
      case o.FeeKindUpdate().Name():
        ret = o.FeeKindUpdate()
    }
    return
}



