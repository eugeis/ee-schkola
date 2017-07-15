package library

import (
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "time"
)

type CreatedBook struct {
    Title  
    Description  string
    Language  string
    ReleaseDate  *time.Time
    Edition  string
    Category  string
    Author  *person.PersonName
    Location  *Location
}

func NewCreatedBook(title , description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName, location *Location) (ret *CreatedBook, err error) {
    ret = &CreatedBook{
        Title : title,
        Description : description,
        Language : language,
        ReleaseDate : releaseDate,
        Edition : edition,
        Category : category,
        Author : author,
        Location : location,
    }
    return
}



type DeletedBook struct {
    Id  string
}

func NewDeletedBook(id string) (ret *DeletedBook, err error) {
    ret = &DeletedBook{
        Id : id,
    }
    return
}



type UpdatedBook struct {
    Title  
    Description  string
    Language  string
    ReleaseDate  *time.Time
    Edition  string
    Category  string
    Author  *person.PersonName
    Location  *Location
}

func NewUpdatedBook(title , description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName, location *Location) (ret *UpdatedBook, err error) {
    ret = &UpdatedBook{
        Title : title,
        Description : description,
        Language : language,
        ReleaseDate : releaseDate,
        Edition : edition,
        Category : category,
        Author : author,
        Location : location,
    }
    return
}




type BookAggregateEventType struct {
	name  string
	ordinal int
}

func (o *BookAggregateEventType) Name() string {
    return o.name
}

func (o *BookAggregateEventType) Ordinal() int {
    return o.ordinal
}

func (o *BookAggregateEventType) IsBookCreated() bool {
    return o == _bookAggregateEventTypes.BookCreated()
}

func (o *BookAggregateEventType) IsBookDeleted() bool {
    return o == _bookAggregateEventTypes.BookDeleted()
}

func (o *BookAggregateEventType) IsBookUpdated() bool {
    return o == _bookAggregateEventTypes.BookUpdated()
}

type bookAggregateEventTypes struct {
	values []*BookAggregateEventType
    literals []enum.Literal
}

var _bookAggregateEventTypes = &bookAggregateEventTypes{values: []*BookAggregateEventType{
    {name: "BookCreated", ordinal: 0},
    {name: "BookDeleted", ordinal: 1},
    {name: "BookUpdated", ordinal: 2}},
}

func BookAggregateEventTypes() *bookAggregateEventTypes {
	return _bookAggregateEventTypes
}

func (o *bookAggregateEventTypes) Values() []*BookAggregateEventType {
	return o.values
}

func (o *bookAggregateEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *bookAggregateEventTypes) BookCreated() *BookAggregateEventType {
    return _bookAggregateEventTypes.values[0]
}

func (o *bookAggregateEventTypes) BookDeleted() *BookAggregateEventType {
    return _bookAggregateEventTypes.values[1]
}

func (o *bookAggregateEventTypes) BookUpdated() *BookAggregateEventType {
    return _bookAggregateEventTypes.values[2]
}

func (o *bookAggregateEventTypes) ParseBookAggregateEventType(name string) (ret *BookAggregateEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*BookAggregateEventType), ok
	}
	return
}



