package library

import (
    "ee/schkola/person"
    "time"
)
type BookCreated struct {
    Title  
    Description  string
    Language  string
    ReleaseDate  *time.Time
    Edition  string
    Category  string
    Author  *person.PersonName
    Location  *Location
}

func NewBookCreated(title , description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName, location *Location) (ret *BookCreated, err error) {
    ret = &BookCreated{
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


type BookDeleted struct {
    Id  string
}

func NewBookDeleted(id string) (ret *BookDeleted, err error) {
    ret = &BookDeleted{
        Id : id,
    }
    return
}


type BookUpdated struct {
    Title  
    Description  string
    Language  string
    ReleaseDate  *time.Time
    Edition  string
    Category  string
    Author  *person.PersonName
    Location  *Location
}

func NewBookUpdated(title , description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName, location *Location) (ret *BookUpdated, err error) {
    ret = &BookUpdated{
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

func (o *BookEventType) IsCreatedBook() bool {
    return o == _bookEventTypes.CreatedBook()
}

func (o *BookEventType) IsDeletedBook() bool {
    return o == _bookEventTypes.DeletedBook()
}

func (o *BookEventType) IsUpdatedBook() bool {
    return o == _bookEventTypes.UpdatedBook()
}

type bookEventTypes struct {
	values []*BookEventType
}

var _bookEventTypes = &bookEventTypes{values: []*BookEventType{
    {name: "createdBook", ordinal: 0},
    {name: "deletedBook", ordinal: 1},
    {name: "updatedBook", ordinal: 2}},
}

func BookEventTypes() *bookEventTypes {
	return _bookEventTypes
}

func (o *bookEventTypes) Values() []*BookEventType {
	return o.values
}

func (o *bookEventTypes) CreatedBook() *BookEventType {
    return _bookEventTypes.values[0]
}

func (o *bookEventTypes) DeletedBook() *BookEventType {
    return _bookEventTypes.values[1]
}

func (o *bookEventTypes) UpdatedBook() *BookEventType {
    return _bookEventTypes.values[2]
}

func (o *bookEventTypes) ParseBookEventType(name string) (ret *BookEventType, ok bool) {
    switch name {
      case o.CreatedBook().Name():
        ret = o.CreatedBook()
      case o.DeletedBook().Name():
        ret = o.DeletedBook()
      case o.UpdatedBook().Name():
        ret = o.UpdatedBook()
    }
    return
}



