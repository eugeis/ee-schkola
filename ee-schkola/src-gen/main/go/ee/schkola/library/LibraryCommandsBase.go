package library

import (
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "time"
)

type BookCreate struct {
    Title  
    Description  string
    Language  string
    ReleaseDate  *time.Time
    Edition  string
    Category  string
    Author  *person.PersonName
    Location  *Location
}

func NewBookCreate(title , description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName, location *Location) (ret *BookCreate, err error) {
    ret = &BookCreate{
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



type BookDelete struct {
    Id  string
}

func NewBookDelete(id string) (ret *BookDelete, err error) {
    ret = &BookDelete{
        Id : id,
    }
    return
}



type BookUpdate struct {
    Title  
    Description  string
    Language  string
    ReleaseDate  *time.Time
    Edition  string
    Category  string
    Author  *person.PersonName
    Location  *Location
}

func NewBookUpdate(title , description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName, location *Location) (ret *BookUpdate, err error) {
    ret = &BookUpdate{
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



type BookRegister struct {
    Title  
    Description  string
    Language  string
    ReleaseDate  *time.Time
    Edition  string
    Category  string
    Author  *person.PersonName
}

func NewBookRegister(title , description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName) (ret *BookRegister, err error) {
    ret = &BookRegister{
        Title : title,
        Description : description,
        Language : language,
        ReleaseDate : releaseDate,
        Edition : edition,
        Category : category,
        Author : author,
    }
    return
}



type BookUnregister struct {
    Id  string
}

func NewBookUnregister(id string) (ret *BookUnregister, err error) {
    ret = &BookUnregister{
        Id : id,
    }
    return
}



type BookChange struct {
    Title  
    Description  string
    Language  string
    ReleaseDate  *time.Time
    Edition  string
    Category  string
    Author  *person.PersonName
}

func NewBookChange(title , description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName) (ret *BookChange, err error) {
    ret = &BookChange{
        Title : title,
        Description : description,
        Language : language,
        ReleaseDate : releaseDate,
        Edition : edition,
        Category : category,
        Author : author,
    }
    return
}



type ChangeBookLocation struct {
    Shelf  string
    Fold  string
}

func NewChangeBookLocation(shelf string, fold string) (ret *ChangeBookLocation, err error) {
    ret = &ChangeBookLocation{
        Shelf : shelf,
        Fold : fold,
    }
    return
}




type BookAggregateCommandType struct {
	name  string
	ordinal int
}

func (o *BookAggregateCommandType) Name() string {
    return o.name
}

func (o *BookAggregateCommandType) Ordinal() int {
    return o.ordinal
}

func (o *BookAggregateCommandType) IsCreateBook() bool {
    return o == _bookAggregateCommandTypes.CreateBook()
}

func (o *BookAggregateCommandType) IsDeleteBook() bool {
    return o == _bookAggregateCommandTypes.DeleteBook()
}

func (o *BookAggregateCommandType) IsUpdateBook() bool {
    return o == _bookAggregateCommandTypes.UpdateBook()
}

func (o *BookAggregateCommandType) IsRegisterBook() bool {
    return o == _bookAggregateCommandTypes.RegisterBook()
}

func (o *BookAggregateCommandType) IsUnregisterBook() bool {
    return o == _bookAggregateCommandTypes.UnregisterBook()
}

func (o *BookAggregateCommandType) IsChangeBook() bool {
    return o == _bookAggregateCommandTypes.ChangeBook()
}

func (o *BookAggregateCommandType) IsChangeLBookocation() bool {
    return o == _bookAggregateCommandTypes.ChangeLBookocation()
}

type bookAggregateCommandTypes struct {
	values []*BookAggregateCommandType
    literals []enum.Literal
}

var _bookAggregateCommandTypes = &bookAggregateCommandTypes{values: []*BookAggregateCommandType{
    {name: "createBook", ordinal: 0},
    {name: "deleteBook", ordinal: 1},
    {name: "updateBook", ordinal: 2},
    {name: "registerBook", ordinal: 3},
    {name: "unregisterBook", ordinal: 4},
    {name: "changeBook", ordinal: 5},
    {name: "changeLBookocation", ordinal: 6}},
}

func BookAggregateCommandTypes() *bookAggregateCommandTypes {
	return _bookAggregateCommandTypes
}

func (o *bookAggregateCommandTypes) Values() []*BookAggregateCommandType {
	return o.values
}

func (o *bookAggregateCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *bookAggregateCommandTypes) CreateBook() *BookAggregateCommandType {
    return _bookAggregateCommandTypes.values[0]
}

func (o *bookAggregateCommandTypes) DeleteBook() *BookAggregateCommandType {
    return _bookAggregateCommandTypes.values[1]
}

func (o *bookAggregateCommandTypes) UpdateBook() *BookAggregateCommandType {
    return _bookAggregateCommandTypes.values[2]
}

func (o *bookAggregateCommandTypes) RegisterBook() *BookAggregateCommandType {
    return _bookAggregateCommandTypes.values[3]
}

func (o *bookAggregateCommandTypes) UnregisterBook() *BookAggregateCommandType {
    return _bookAggregateCommandTypes.values[4]
}

func (o *bookAggregateCommandTypes) ChangeBook() *BookAggregateCommandType {
    return _bookAggregateCommandTypes.values[5]
}

func (o *bookAggregateCommandTypes) ChangeLBookocation() *BookAggregateCommandType {
    return _bookAggregateCommandTypes.values[6]
}

func (o *bookAggregateCommandTypes) ParseBookAggregateCommandType(name string) (ret *BookAggregateCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*BookAggregateCommandType), ok
	}
	return
}



