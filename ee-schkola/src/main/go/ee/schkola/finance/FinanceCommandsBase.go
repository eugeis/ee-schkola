package finance

import (
    "ee/schkola/person"
    "encoding/json"
    "fmt"
    "github.com/eugeis/gee/enum"
    "github.com/looplab/eventhorizon"
    "gopkg.in/mgo.v2/bson"
    "time"
)
const (
     CreateCommand eventhorizon.CommandType = "Create"
     DeleteCommand eventhorizon.CommandType = "Delete"
     UpdateCommand eventhorizon.CommandType = "Update"
)


const (
     CreateCommand eventhorizon.CommandType = "Create"
     DeleteCommand eventhorizon.CommandType = "Delete"
     UpdateCommand eventhorizon.CommandType = "Update"
)


const (
     CreateCommand eventhorizon.CommandType = "Create"
     DeleteCommand eventhorizon.CommandType = "Delete"
     UpdateCommand eventhorizon.CommandType = "Update"
)


const (
     CreateCommand eventhorizon.CommandType = "Create"
     DeleteCommand eventhorizon.CommandType = "Delete"
     UpdateCommand eventhorizon.CommandType = "Update"
)




        
type Create struct {
    Purpose *ExpensePurpose `json:"purpose" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Profile *person.Profile `json:"profile" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Create) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Create) AggregateType() eventhorizon.AggregateType  { return ExpenseAggregateType }
func (o *Create) CommandType() eventhorizon.CommandType      { return CreateCommand }



        
type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Delete) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Delete) AggregateType() eventhorizon.AggregateType  { return ExpenseAggregateType }
func (o *Delete) CommandType() eventhorizon.CommandType      { return DeleteCommand }



        
type Update struct {
    Purpose *ExpensePurpose `json:"purpose" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Profile *person.Profile `json:"profile" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Update) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Update) AggregateType() eventhorizon.AggregateType  { return ExpenseAggregateType }
func (o *Update) CommandType() eventhorizon.CommandType      { return UpdateCommand }



        
type Create struct {
    Name string `json:"name" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Create) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Create) AggregateType() eventhorizon.AggregateType  { return ExpensePurposeAggregateType }
func (o *Create) CommandType() eventhorizon.CommandType      { return CreateCommand }



        
type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Delete) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Delete) AggregateType() eventhorizon.AggregateType  { return ExpensePurposeAggregateType }
func (o *Delete) CommandType() eventhorizon.CommandType      { return DeleteCommand }



        
type Update struct {
    Name string `json:"name" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Update) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Update) AggregateType() eventhorizon.AggregateType  { return ExpensePurposeAggregateType }
func (o *Update) CommandType() eventhorizon.CommandType      { return UpdateCommand }



        
type Create struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Kind *FeeKind `json:"kind" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Create) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Create) AggregateType() eventhorizon.AggregateType  { return FeeAggregateType }
func (o *Create) CommandType() eventhorizon.CommandType      { return CreateCommand }



        
type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Delete) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Delete) AggregateType() eventhorizon.AggregateType  { return FeeAggregateType }
func (o *Delete) CommandType() eventhorizon.CommandType      { return DeleteCommand }



        
type Update struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Kind *FeeKind `json:"kind" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Update) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Update) AggregateType() eventhorizon.AggregateType  { return FeeAggregateType }
func (o *Update) CommandType() eventhorizon.CommandType      { return UpdateCommand }



        
type Create struct {
    Name string `json:"name" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Create) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Create) AggregateType() eventhorizon.AggregateType  { return FeeKindAggregateType }
func (o *Create) CommandType() eventhorizon.CommandType      { return CreateCommand }



        
type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Delete) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Delete) AggregateType() eventhorizon.AggregateType  { return FeeKindAggregateType }
func (o *Delete) CommandType() eventhorizon.CommandType      { return DeleteCommand }



        
type Update struct {
    Name string `json:"name" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Update) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Update) AggregateType() eventhorizon.AggregateType  { return FeeKindAggregateType }
func (o *Update) CommandType() eventhorizon.CommandType      { return UpdateCommand }





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

func (o ExpenseCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *ExpenseCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := ExpenseCommandTypes().ParseExpenseCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ExpenseCommandType %q", lit.Name)
        }
	}
	return
}

func (o ExpenseCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *ExpenseCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := ExpenseCommandTypes().ParseExpenseCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ExpenseCommandType %q", lit)
        }
    }
    return
}

func (o *ExpenseCommandType) IsCreate() bool {
    return o == _expenseCommandTypes.Create()
}

func (o *ExpenseCommandType) IsDelete() bool {
    return o == _expenseCommandTypes.Delete()
}

func (o *ExpenseCommandType) IsUpdate() bool {
    return o == _expenseCommandTypes.Update()
}

type expenseCommandTypes struct {
	values []*ExpenseCommandType
    literals []enum.Literal
}

var _expenseCommandTypes = &expenseCommandTypes{values: []*ExpenseCommandType{
    {name: "Create", ordinal: 0},
    {name: "Delete", ordinal: 1},
    {name: "Update", ordinal: 2}},
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

func (o *expenseCommandTypes) Create() *ExpenseCommandType {
    return _expenseCommandTypes.values[0]
}

func (o *expenseCommandTypes) Delete() *ExpenseCommandType {
    return _expenseCommandTypes.values[1]
}

func (o *expenseCommandTypes) Update() *ExpenseCommandType {
    return _expenseCommandTypes.values[2]
}

func (o *expenseCommandTypes) ParseExpenseCommandType(name string) (ret *ExpenseCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
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

func (o ExpensePurposeCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *ExpensePurposeCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := ExpensePurposeCommandTypes().ParseExpensePurposeCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ExpensePurposeCommandType %q", lit.Name)
        }
	}
	return
}

func (o ExpensePurposeCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *ExpensePurposeCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := ExpensePurposeCommandTypes().ParseExpensePurposeCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ExpensePurposeCommandType %q", lit)
        }
    }
    return
}

func (o *ExpensePurposeCommandType) IsCreate() bool {
    return o == _expensePurposeCommandTypes.Create()
}

func (o *ExpensePurposeCommandType) IsDelete() bool {
    return o == _expensePurposeCommandTypes.Delete()
}

func (o *ExpensePurposeCommandType) IsUpdate() bool {
    return o == _expensePurposeCommandTypes.Update()
}

type expensePurposeCommandTypes struct {
	values []*ExpensePurposeCommandType
    literals []enum.Literal
}

var _expensePurposeCommandTypes = &expensePurposeCommandTypes{values: []*ExpensePurposeCommandType{
    {name: "Create", ordinal: 0},
    {name: "Delete", ordinal: 1},
    {name: "Update", ordinal: 2}},
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

func (o *expensePurposeCommandTypes) Create() *ExpensePurposeCommandType {
    return _expensePurposeCommandTypes.values[0]
}

func (o *expensePurposeCommandTypes) Delete() *ExpensePurposeCommandType {
    return _expensePurposeCommandTypes.values[1]
}

func (o *expensePurposeCommandTypes) Update() *ExpensePurposeCommandType {
    return _expensePurposeCommandTypes.values[2]
}

func (o *expensePurposeCommandTypes) ParseExpensePurposeCommandType(name string) (ret *ExpensePurposeCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
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

func (o FeeCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *FeeCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := FeeCommandTypes().ParseFeeCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid FeeCommandType %q", lit.Name)
        }
	}
	return
}

func (o FeeCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *FeeCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := FeeCommandTypes().ParseFeeCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid FeeCommandType %q", lit)
        }
    }
    return
}

func (o *FeeCommandType) IsCreate() bool {
    return o == _feeCommandTypes.Create()
}

func (o *FeeCommandType) IsDelete() bool {
    return o == _feeCommandTypes.Delete()
}

func (o *FeeCommandType) IsUpdate() bool {
    return o == _feeCommandTypes.Update()
}

type feeCommandTypes struct {
	values []*FeeCommandType
    literals []enum.Literal
}

var _feeCommandTypes = &feeCommandTypes{values: []*FeeCommandType{
    {name: "Create", ordinal: 0},
    {name: "Delete", ordinal: 1},
    {name: "Update", ordinal: 2}},
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

func (o *feeCommandTypes) Create() *FeeCommandType {
    return _feeCommandTypes.values[0]
}

func (o *feeCommandTypes) Delete() *FeeCommandType {
    return _feeCommandTypes.values[1]
}

func (o *feeCommandTypes) Update() *FeeCommandType {
    return _feeCommandTypes.values[2]
}

func (o *feeCommandTypes) ParseFeeCommandType(name string) (ret *FeeCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
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

func (o FeeKindCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *FeeKindCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := FeeKindCommandTypes().ParseFeeKindCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid FeeKindCommandType %q", lit.Name)
        }
	}
	return
}

func (o FeeKindCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *FeeKindCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := FeeKindCommandTypes().ParseFeeKindCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid FeeKindCommandType %q", lit)
        }
    }
    return
}

func (o *FeeKindCommandType) IsCreate() bool {
    return o == _feeKindCommandTypes.Create()
}

func (o *FeeKindCommandType) IsDelete() bool {
    return o == _feeKindCommandTypes.Delete()
}

func (o *FeeKindCommandType) IsUpdate() bool {
    return o == _feeKindCommandTypes.Update()
}

type feeKindCommandTypes struct {
	values []*FeeKindCommandType
    literals []enum.Literal
}

var _feeKindCommandTypes = &feeKindCommandTypes{values: []*FeeKindCommandType{
    {name: "Create", ordinal: 0},
    {name: "Delete", ordinal: 1},
    {name: "Update", ordinal: 2}},
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

func (o *feeKindCommandTypes) Create() *FeeKindCommandType {
    return _feeKindCommandTypes.values[0]
}

func (o *feeKindCommandTypes) Delete() *FeeKindCommandType {
    return _feeKindCommandTypes.values[1]
}

func (o *feeKindCommandTypes) Update() *FeeKindCommandType {
    return _feeKindCommandTypes.values[2]
}

func (o *feeKindCommandTypes) ParseFeeKindCommandType(name string) (ret *FeeKindCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*FeeKindCommandType), ok
	}
	return
}



