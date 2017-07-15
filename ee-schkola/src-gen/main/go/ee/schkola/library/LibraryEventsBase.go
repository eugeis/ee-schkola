package library

import (
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "time"
)

type CreatedBook struct {
    title 
    description string
    language string
    releaseDate *time.Time
    edition string
    category string
    author *person.PersonName
    location *Location
}

func NewCreatedBook(title , description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName, location *Location) (ret *CreatedBook, err error) {
    ret = &CreatedBook{
        title: title,
        description: description,
        language: language,
        releaseDate: releaseDate,
        edition: edition,
        category: category,
        author: author,
        location: location,
    }
    
    return
}



type DeletedBook struct {
    id string
}

func NewDeletedBook(id string) (ret *DeletedBook, err error) {
    ret = &DeletedBook{
        id: id,
    }
    
    return
}



type UpdatedBook struct {
    title 
    description string
    language string
    releaseDate *time.Time
    edition string
    category string
    author *person.PersonName
    location *Location
}

func NewUpdatedBook(title , description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName, location *Location) (ret *UpdatedBook, err error) {
    ret = &UpdatedBook{
        title: title,
        description: description,
        language: language,
        releaseDate: releaseDate,
        edition: edition,
        category: category,
        author: author,
        location: location,
    }
    
    return
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



