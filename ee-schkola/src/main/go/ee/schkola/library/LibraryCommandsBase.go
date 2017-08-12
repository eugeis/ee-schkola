package library

import (
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "github.com/looplab/eventhorizon"
    "time"
)
const (
     CreateBookCommand eventhorizon.CommandType = "CreateBook"
     DeleteBookCommand eventhorizon.CommandType = "DeleteBook"
     ChangeBookLocationCommand eventhorizon.CommandType = "ChangeBookLocation"
     UpdateBookCommand eventhorizon.CommandType = "UpdateBook"
)




        
type CreateBook struct {
    Id eventhorizon.UUID`eh:"optional"`
    Title string`eh:"optional"`
    Description string`eh:"optional"`
    Language string`eh:"optional"`
    ReleaseDate *time.Time`eh:"optional"`
    Edition string`eh:"optional"`
    Category string`eh:"optional"`
    Author *person.PersonName`eh:"optional"`
    Location *Location`eh:"optional"`
}
func (o *CreateBook) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *CreateBook) AggregateType() eventhorizon.AggregateType  { return BookAggregateType }
func (o *CreateBook) CommandType() eventhorizon.CommandType      { return CreateBookCommand }



        
type DeleteBook struct {
    Id eventhorizon.UUID
}
func (o *DeleteBook) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *DeleteBook) AggregateType() eventhorizon.AggregateType  { return BookAggregateType }
func (o *DeleteBook) CommandType() eventhorizon.CommandType      { return DeleteBookCommand }



        
type ChangeBookLocation struct {
    Id eventhorizon.UUID
    Location *Location
}

func NewChangeLocation() (ret *ChangeBookLocation) {
    ret = &ChangeBookLocation{}
    return
}
func (o *ChangeBookLocation) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *ChangeBookLocation) AggregateType() eventhorizon.AggregateType  { return BookAggregateType }
func (o *ChangeBookLocation) CommandType() eventhorizon.CommandType      { return ChangeBookLocationCommand }



        
type UpdateBook struct {
    Id eventhorizon.UUID`eh:"optional"`
    Title string`eh:"optional"`
    Description string`eh:"optional"`
    Language string`eh:"optional"`
    ReleaseDate *time.Time`eh:"optional"`
    Edition string`eh:"optional"`
    Category string`eh:"optional"`
    Author *person.PersonName`eh:"optional"`
    Location *Location`eh:"optional"`
}
func (o *UpdateBook) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *UpdateBook) AggregateType() eventhorizon.AggregateType  { return BookAggregateType }
func (o *UpdateBook) CommandType() eventhorizon.CommandType      { return UpdateBookCommand }





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

func (o *BookCommandType) IsCreateBook() bool {
    return o == _bookCommandTypes.CreateBook()
}

func (o *BookCommandType) IsDeleteBook() bool {
    return o == _bookCommandTypes.DeleteBook()
}

func (o *BookCommandType) IsChangeBookLocation() bool {
    return o == _bookCommandTypes.ChangeBookLocation()
}

func (o *BookCommandType) IsUpdateBook() bool {
    return o == _bookCommandTypes.UpdateBook()
}

type bookCommandTypes struct {
	values []*BookCommandType
    literals []enum.Literal
}

var _bookCommandTypes = &bookCommandTypes{values: []*BookCommandType{
    {name: "CreateBook", ordinal: 0},
    {name: "DeleteBook", ordinal: 1},
    {name: "ChangeBookLocation", ordinal: 2},
    {name: "UpdateBook", ordinal: 3}},
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

func (o *bookCommandTypes) CreateBook() *BookCommandType {
    return _bookCommandTypes.values[0]
}

func (o *bookCommandTypes) DeleteBook() *BookCommandType {
    return _bookCommandTypes.values[1]
}

func (o *bookCommandTypes) ChangeBookLocation() *BookCommandType {
    return _bookCommandTypes.values[2]
}

func (o *bookCommandTypes) UpdateBook() *BookCommandType {
    return _bookCommandTypes.values[3]
}

func (o *bookCommandTypes) ParseBookCommandType(name string) (ret *BookCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*BookCommandType), ok
	}
	return
}



