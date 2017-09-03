package library

import (
    "ee/schkola"
    "github.com/eugeis/gee/enum"
    "github.com/looplab/eventhorizon"
    "time"
)
const (
     LocationChangedBookEvent eventhorizon.EventType = "LocationChangedBook"
     BookCreatedEvent eventhorizon.EventType = "BookCreated"
     BookDeletedEvent eventhorizon.EventType = "BookDeleted"
     BookUpdatedEvent eventhorizon.EventType = "BookUpdated"
)




type LocationChangedBook struct {
    Location *Location
    Id eventhorizon.UUID
}

func NewLocationChanged() (ret *LocationChangedBook) {
    ret = &LocationChangedBook{}
    return
}


type BookCreated struct {
    Title string`eh:"optional"`
    Description string`eh:"optional"`
    Language string`eh:"optional"`
    ReleaseDate *time.Time`eh:"optional"`
    Edition string`eh:"optional"`
    Category string`eh:"optional"`
    Author *schkola.PersonName`eh:"optional"`
    Location *Location`eh:"optional"`
    Id eventhorizon.UUID`eh:"optional"`
}


type BookDeleted struct {
    Id eventhorizon.UUID
}


type BookUpdated struct {
    Title string`eh:"optional"`
    Description string`eh:"optional"`
    Language string`eh:"optional"`
    ReleaseDate *time.Time`eh:"optional"`
    Edition string`eh:"optional"`
    Category string`eh:"optional"`
    Author *schkola.PersonName`eh:"optional"`
    Location *Location`eh:"optional"`
    Id eventhorizon.UUID`eh:"optional"`
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

func (o *BookEventType) IsLocationChangedBook() bool {
    return o == _bookEventTypes.LocationChangedBook()
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
    {name: "LocationChangedBook", ordinal: 0},
    {name: "BookCreated", ordinal: 1},
    {name: "BookDeleted", ordinal: 2},
    {name: "BookUpdated", ordinal: 3}},
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

func (o *bookEventTypes) LocationChangedBook() *BookEventType {
    return _bookEventTypes.values[0]
}

func (o *bookEventTypes) BookCreated() *BookEventType {
    return _bookEventTypes.values[1]
}

func (o *bookEventTypes) BookDeleted() *BookEventType {
    return _bookEventTypes.values[2]
}

func (o *bookEventTypes) BookUpdated() *BookEventType {
    return _bookEventTypes.values[3]
}

func (o *bookEventTypes) ParseBookEventType(name string) (ret *BookEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*BookEventType), ok
	}
	return
}



