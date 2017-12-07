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
     CreateEvent eventhorizon.EventType = "Create"
     CreatedEvent eventhorizon.EventType = "Created"
     DeleteEvent eventhorizon.EventType = "Delete"
     DeletedEvent eventhorizon.EventType = "Deleted"
     UpdateEvent eventhorizon.EventType = "Update"
     UpdatedEvent eventhorizon.EventType = "Updated"
)


const (
     CreateEvent eventhorizon.EventType = "Create"
     CreatedEvent eventhorizon.EventType = "Created"
     DeleteEvent eventhorizon.EventType = "Delete"
     DeletedEvent eventhorizon.EventType = "Deleted"
     UpdateEvent eventhorizon.EventType = "Update"
     UpdatedEvent eventhorizon.EventType = "Updated"
)


const (
     CreateEvent eventhorizon.EventType = "Create"
     CreatedEvent eventhorizon.EventType = "Created"
     DeleteEvent eventhorizon.EventType = "Delete"
     DeletedEvent eventhorizon.EventType = "Deleted"
     UpdateEvent eventhorizon.EventType = "Update"
     UpdatedEvent eventhorizon.EventType = "Updated"
)


const (
     CreateEvent eventhorizon.EventType = "Create"
     CreatedEvent eventhorizon.EventType = "Created"
     DeleteEvent eventhorizon.EventType = "Delete"
     DeletedEvent eventhorizon.EventType = "Deleted"
     UpdateEvent eventhorizon.EventType = "Update"
     UpdatedEvent eventhorizon.EventType = "Updated"
)




type Create struct {
    Purpose *ExpensePurpose `json:"purpose" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Profile *person.Profile `json:"profile" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Created struct {
    Purpose *ExpensePurpose `json:"purpose" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Profile *person.Profile `json:"profile" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Deleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Update struct {
    Purpose *ExpensePurpose `json:"purpose" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Profile *person.Profile `json:"profile" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Updated struct {
    Purpose *ExpensePurpose `json:"purpose" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Profile *person.Profile `json:"profile" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Create struct {
    Name string `json:"name" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Created struct {
    Name string `json:"name" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Deleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Update struct {
    Name string `json:"name" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Updated struct {
    Name string `json:"name" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Create struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Kind *FeeKind `json:"kind" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Created struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Kind *FeeKind `json:"kind" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Deleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Update struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Kind *FeeKind `json:"kind" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Updated struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Kind *FeeKind `json:"kind" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Create struct {
    Name string `json:"name" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Created struct {
    Name string `json:"name" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Deleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Update struct {
    Name string `json:"name" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Updated struct {
    Name string `json:"name" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
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

func (o ExpenseEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *ExpenseEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := ExpenseEventTypes().ParseExpenseEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ExpenseEventType %q", lit.Name)
        }
	}
	return
}

func (o ExpenseEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *ExpenseEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := ExpenseEventTypes().ParseExpenseEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ExpenseEventType %q", lit)
        }
    }
    return
}

func (o *ExpenseEventType) IsCreate() bool {
    return o == _expenseEventTypes.Create()
}

func (o *ExpenseEventType) IsCreated() bool {
    return o == _expenseEventTypes.Created()
}

func (o *ExpenseEventType) IsDelete() bool {
    return o == _expenseEventTypes.Delete()
}

func (o *ExpenseEventType) IsDeleted() bool {
    return o == _expenseEventTypes.Deleted()
}

func (o *ExpenseEventType) IsUpdate() bool {
    return o == _expenseEventTypes.Update()
}

func (o *ExpenseEventType) IsUpdated() bool {
    return o == _expenseEventTypes.Updated()
}

type expenseEventTypes struct {
	values []*ExpenseEventType
    literals []enum.Literal
}

var _expenseEventTypes = &expenseEventTypes{values: []*ExpenseEventType{
    {name: "Create", ordinal: 0},
    {name: "Created", ordinal: 1},
    {name: "Delete", ordinal: 2},
    {name: "Deleted", ordinal: 3},
    {name: "Update", ordinal: 4},
    {name: "Updated", ordinal: 5}},
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

func (o *expenseEventTypes) Create() *ExpenseEventType {
    return _expenseEventTypes.values[0]
}

func (o *expenseEventTypes) Created() *ExpenseEventType {
    return _expenseEventTypes.values[1]
}

func (o *expenseEventTypes) Delete() *ExpenseEventType {
    return _expenseEventTypes.values[2]
}

func (o *expenseEventTypes) Deleted() *ExpenseEventType {
    return _expenseEventTypes.values[3]
}

func (o *expenseEventTypes) Update() *ExpenseEventType {
    return _expenseEventTypes.values[4]
}

func (o *expenseEventTypes) Updated() *ExpenseEventType {
    return _expenseEventTypes.values[5]
}

func (o *expenseEventTypes) ParseExpenseEventType(name string) (ret *ExpenseEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
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

func (o ExpensePurposeEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *ExpensePurposeEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := ExpensePurposeEventTypes().ParseExpensePurposeEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ExpensePurposeEventType %q", lit.Name)
        }
	}
	return
}

func (o ExpensePurposeEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *ExpensePurposeEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := ExpensePurposeEventTypes().ParseExpensePurposeEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ExpensePurposeEventType %q", lit)
        }
    }
    return
}

func (o *ExpensePurposeEventType) IsCreate() bool {
    return o == _expensePurposeEventTypes.Create()
}

func (o *ExpensePurposeEventType) IsCreated() bool {
    return o == _expensePurposeEventTypes.Created()
}

func (o *ExpensePurposeEventType) IsDelete() bool {
    return o == _expensePurposeEventTypes.Delete()
}

func (o *ExpensePurposeEventType) IsDeleted() bool {
    return o == _expensePurposeEventTypes.Deleted()
}

func (o *ExpensePurposeEventType) IsUpdate() bool {
    return o == _expensePurposeEventTypes.Update()
}

func (o *ExpensePurposeEventType) IsUpdated() bool {
    return o == _expensePurposeEventTypes.Updated()
}

type expensePurposeEventTypes struct {
	values []*ExpensePurposeEventType
    literals []enum.Literal
}

var _expensePurposeEventTypes = &expensePurposeEventTypes{values: []*ExpensePurposeEventType{
    {name: "Create", ordinal: 0},
    {name: "Created", ordinal: 1},
    {name: "Delete", ordinal: 2},
    {name: "Deleted", ordinal: 3},
    {name: "Update", ordinal: 4},
    {name: "Updated", ordinal: 5}},
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

func (o *expensePurposeEventTypes) Create() *ExpensePurposeEventType {
    return _expensePurposeEventTypes.values[0]
}

func (o *expensePurposeEventTypes) Created() *ExpensePurposeEventType {
    return _expensePurposeEventTypes.values[1]
}

func (o *expensePurposeEventTypes) Delete() *ExpensePurposeEventType {
    return _expensePurposeEventTypes.values[2]
}

func (o *expensePurposeEventTypes) Deleted() *ExpensePurposeEventType {
    return _expensePurposeEventTypes.values[3]
}

func (o *expensePurposeEventTypes) Update() *ExpensePurposeEventType {
    return _expensePurposeEventTypes.values[4]
}

func (o *expensePurposeEventTypes) Updated() *ExpensePurposeEventType {
    return _expensePurposeEventTypes.values[5]
}

func (o *expensePurposeEventTypes) ParseExpensePurposeEventType(name string) (ret *ExpensePurposeEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
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

func (o FeeEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *FeeEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := FeeEventTypes().ParseFeeEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid FeeEventType %q", lit.Name)
        }
	}
	return
}

func (o FeeEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *FeeEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := FeeEventTypes().ParseFeeEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid FeeEventType %q", lit)
        }
    }
    return
}

func (o *FeeEventType) IsCreate() bool {
    return o == _feeEventTypes.Create()
}

func (o *FeeEventType) IsCreated() bool {
    return o == _feeEventTypes.Created()
}

func (o *FeeEventType) IsDelete() bool {
    return o == _feeEventTypes.Delete()
}

func (o *FeeEventType) IsDeleted() bool {
    return o == _feeEventTypes.Deleted()
}

func (o *FeeEventType) IsUpdate() bool {
    return o == _feeEventTypes.Update()
}

func (o *FeeEventType) IsUpdated() bool {
    return o == _feeEventTypes.Updated()
}

type feeEventTypes struct {
	values []*FeeEventType
    literals []enum.Literal
}

var _feeEventTypes = &feeEventTypes{values: []*FeeEventType{
    {name: "Create", ordinal: 0},
    {name: "Created", ordinal: 1},
    {name: "Delete", ordinal: 2},
    {name: "Deleted", ordinal: 3},
    {name: "Update", ordinal: 4},
    {name: "Updated", ordinal: 5}},
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

func (o *feeEventTypes) Create() *FeeEventType {
    return _feeEventTypes.values[0]
}

func (o *feeEventTypes) Created() *FeeEventType {
    return _feeEventTypes.values[1]
}

func (o *feeEventTypes) Delete() *FeeEventType {
    return _feeEventTypes.values[2]
}

func (o *feeEventTypes) Deleted() *FeeEventType {
    return _feeEventTypes.values[3]
}

func (o *feeEventTypes) Update() *FeeEventType {
    return _feeEventTypes.values[4]
}

func (o *feeEventTypes) Updated() *FeeEventType {
    return _feeEventTypes.values[5]
}

func (o *feeEventTypes) ParseFeeEventType(name string) (ret *FeeEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
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

func (o FeeKindEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *FeeKindEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := FeeKindEventTypes().ParseFeeKindEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid FeeKindEventType %q", lit.Name)
        }
	}
	return
}

func (o FeeKindEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *FeeKindEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := FeeKindEventTypes().ParseFeeKindEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid FeeKindEventType %q", lit)
        }
    }
    return
}

func (o *FeeKindEventType) IsCreate() bool {
    return o == _feeKindEventTypes.Create()
}

func (o *FeeKindEventType) IsCreated() bool {
    return o == _feeKindEventTypes.Created()
}

func (o *FeeKindEventType) IsDelete() bool {
    return o == _feeKindEventTypes.Delete()
}

func (o *FeeKindEventType) IsDeleted() bool {
    return o == _feeKindEventTypes.Deleted()
}

func (o *FeeKindEventType) IsUpdate() bool {
    return o == _feeKindEventTypes.Update()
}

func (o *FeeKindEventType) IsUpdated() bool {
    return o == _feeKindEventTypes.Updated()
}

type feeKindEventTypes struct {
	values []*FeeKindEventType
    literals []enum.Literal
}

var _feeKindEventTypes = &feeKindEventTypes{values: []*FeeKindEventType{
    {name: "Create", ordinal: 0},
    {name: "Created", ordinal: 1},
    {name: "Delete", ordinal: 2},
    {name: "Deleted", ordinal: 3},
    {name: "Update", ordinal: 4},
    {name: "Updated", ordinal: 5}},
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

func (o *feeKindEventTypes) Create() *FeeKindEventType {
    return _feeKindEventTypes.values[0]
}

func (o *feeKindEventTypes) Created() *FeeKindEventType {
    return _feeKindEventTypes.values[1]
}

func (o *feeKindEventTypes) Delete() *FeeKindEventType {
    return _feeKindEventTypes.values[2]
}

func (o *feeKindEventTypes) Deleted() *FeeKindEventType {
    return _feeKindEventTypes.values[3]
}

func (o *feeKindEventTypes) Update() *FeeKindEventType {
    return _feeKindEventTypes.values[4]
}

func (o *feeKindEventTypes) Updated() *FeeKindEventType {
    return _feeKindEventTypes.values[5]
}

func (o *feeKindEventTypes) ParseFeeKindEventType(name string) (ret *FeeKindEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*FeeKindEventType), ok
	}
	return
}



