package library

import (
    "ee/schkola/person"
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/enum"
    "time"
)
const (
     BookCreatedEvent eventhorizon.EventType = "BookCreated"
     BookDeletedEvent eventhorizon.EventType = "BookDeleted"
     BookUpdatedEvent eventhorizon.EventType = "BookUpdated"
)





type BookCreated struct {
    Id  eventhorizon.UUID`eh:"optional"`
    Title  string`eh:"optional"`
    Description  string`eh:"optional"`
    Language  string`eh:"optional"`
    ReleaseDate  *time.Time`eh:"optional"`
    Edition  string`eh:"optional"`
    Category  string`eh:"optional"`
    Author  *person.PersonName`eh:"optional"`
    Location  *Location`eh:"optional"`
}



type BookDeleted struct {
    Id  eventhorizon.UUID
}



type BookUpdated struct {
    Id  eventhorizon.UUID`eh:"optional"`
    Title  string`eh:"optional"`
    Description  string`eh:"optional"`
    Language  string`eh:"optional"`
    ReleaseDate  *time.Time`eh:"optional"`
    Edition  string`eh:"optional"`
    Category  string`eh:"optional"`
    Author  *person.PersonName`eh:"optional"`
    Location  *Location`eh:"optional"`
}




type BookEventType struct {
	name  string
	ordinal int
}

func (o *BookEventType) Name() string {
    return o.name
}

func (o *BookEventType) Ordinal() int {
    return o.ordinal
}

func (o *BookEventType) IsBookCreated() bool {
    return o == _bookEventTypes.BookCreated()
}

func (o *BookEventType) IsBookDeleted() bool {
    return o == _bookEventTypes.BookDeleted()
}

func (o *BookEventType) IsBookUpdated() bool {
    return o == _bookEventTypes.BookUpdated()
}

type bookEventTypes struct {
	values []*BookEventType
    literals []enum.Literal
}

var _bookEventTypes = &bookEventTypes{values: []*BookEventType{
    {name: "BookCreated", ordinal: 0},
    {name: "BookDeleted", ordinal: 1},
    {name: "BookUpdated", ordinal: 2}},
}

func BookEventTypes() *bookEventTypes {
	return _bookEventTypes
}

func (o *bookEventTypes) Values() []*BookEventType {
	return o.values
}

func (o *bookEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *bookEventTypes) BookCreated() *BookEventType {
    return _bookEventTypes.values[0]
}

func (o *bookEventTypes) BookDeleted() *BookEventType {
    return _bookEventTypes.values[1]
}

func (o *bookEventTypes) BookUpdated() *BookEventType {
    return _bookEventTypes.values[2]
}

func (o *bookEventTypes) ParseBookEventType(name string) (ret *BookEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*BookEventType), ok
	}
	return
}


