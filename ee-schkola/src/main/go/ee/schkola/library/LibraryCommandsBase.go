package library

import (
    "ee/schkola/shared"
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
     ChangeLocationCommand eventhorizon.CommandType = "ChangeLocation"
     UpdateCommand eventhorizon.CommandType = "Update"
)




        
type Create struct {
    Location *Location `json:"location" eh:"optional"`
    Title string `json:"title" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Language string `json:"language" eh:"optional"`
    ReleaseDate *time.Time `json:"releaseDate" eh:"optional"`
    Edition string `json:"edition" eh:"optional"`
    Category string `json:"category" eh:"optional"`
    Author *shared.PersonName `json:"author" eh:"optional"`
    Location *Location `json:"location" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Create) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Create) AggregateType() eventhorizon.AggregateType  { return BookAggregateType }
func (o *Create) CommandType() eventhorizon.CommandType      { return CreateCommand }



        
type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Delete) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Delete) AggregateType() eventhorizon.AggregateType  { return BookAggregateType }
func (o *Delete) CommandType() eventhorizon.CommandType      { return DeleteCommand }



        
type ChangeLocation struct {
    Location *Location `json:"location" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *ChangeLocation) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *ChangeLocation) AggregateType() eventhorizon.AggregateType  { return BookAggregateType }
func (o *ChangeLocation) CommandType() eventhorizon.CommandType      { return ChangeLocationCommand }



        
type Update struct {
    Location *Location `json:"location" eh:"optional"`
    Title string `json:"title" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Language string `json:"language" eh:"optional"`
    ReleaseDate *time.Time `json:"releaseDate" eh:"optional"`
    Edition string `json:"edition" eh:"optional"`
    Category string `json:"category" eh:"optional"`
    Author *shared.PersonName `json:"author" eh:"optional"`
    Location *Location `json:"location" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Update) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Update) AggregateType() eventhorizon.AggregateType  { return BookAggregateType }
func (o *Update) CommandType() eventhorizon.CommandType      { return UpdateCommand }





type BookCommandType struct {
	name  string
	ordinal int
}

func (o *BookCommandType) Name() string {
    return o.name
}

func (o *BookCommandType) Ordinal() int {
    return o.ordinal
}

func (o BookCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *BookCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := BookCommandTypes().ParseBookCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid BookCommandType %q", lit.Name)
        }
	}
	return
}

func (o BookCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *BookCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := BookCommandTypes().ParseBookCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid BookCommandType %q", lit)
        }
    }
    return
}

func (o *BookCommandType) IsCreate() bool {
    return o == _bookCommandTypes.Create()
}

func (o *BookCommandType) IsDelete() bool {
    return o == _bookCommandTypes.Delete()
}

func (o *BookCommandType) IsChangeLocation() bool {
    return o == _bookCommandTypes.ChangeLocation()
}

func (o *BookCommandType) IsUpdate() bool {
    return o == _bookCommandTypes.Update()
}

type bookCommandTypes struct {
	values []*BookCommandType
    literals []enum.Literal
}

var _bookCommandTypes = &bookCommandTypes{values: []*BookCommandType{
    {name: "Create", ordinal: 0},
    {name: "Delete", ordinal: 1},
    {name: "ChangeLocation", ordinal: 2},
    {name: "Update", ordinal: 3}},
}

func BookCommandTypes() *bookCommandTypes {
	return _bookCommandTypes
}

func (o *bookCommandTypes) Values() []*BookCommandType {
	return o.values
}

func (o *bookCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *bookCommandTypes) Create() *BookCommandType {
    return _bookCommandTypes.values[0]
}

func (o *bookCommandTypes) Delete() *BookCommandType {
    return _bookCommandTypes.values[1]
}

func (o *bookCommandTypes) ChangeLocation() *BookCommandType {
    return _bookCommandTypes.values[2]
}

func (o *bookCommandTypes) Update() *BookCommandType {
    return _bookCommandTypes.values[3]
}

func (o *bookCommandTypes) ParseBookCommandType(name string) (ret *BookCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*BookCommandType), ok
	}
	return
}



