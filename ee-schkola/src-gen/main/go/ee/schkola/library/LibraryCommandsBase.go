package library

import (
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "time"
)

type BookCreate struct {
    title 
    description string
    language string
    releaseDate *time.Time
    edition string
    category string
    author *person.PersonName
    location *Location
}

func NewBookCreate(title , description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName, location *Location) (ret *BookCreate, err error) {
    ret = &BookCreate{
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



type BookDelete struct {
    id string
}

func NewBookDelete(id string) (ret *BookDelete, err error) {
    ret = &BookDelete{
        id: id,
    }
    
    return
}



type BookUpdate struct {
    title 
    description string
    language string
    releaseDate *time.Time
    edition string
    category string
    author *person.PersonName
    location *Location
}

func NewBookUpdate(title , description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName, location *Location) (ret *BookUpdate, err error) {
    ret = &BookUpdate{
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



type BookRegister struct {
    title 
    description string
    language string
    releaseDate *time.Time
    edition string
    category string
    author *person.PersonName
}

func NewBookRegister(title , description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName) (ret *BookRegister, err error) {
    ret = &BookRegister{
        title: title,
        description: description,
        language: language,
        releaseDate: releaseDate,
        edition: edition,
        category: category,
        author: author,
    }
    
    return
}



type BookUnregister struct {
    id string
}

func NewBookUnregister(id string) (ret *BookUnregister, err error) {
    ret = &BookUnregister{
        id: id,
    }
    
    return
}



type BookChange struct {
    title 
    description string
    language string
    releaseDate *time.Time
    edition string
    category string
    author *person.PersonName
}

func NewBookChange(title , description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName) (ret *BookChange, err error) {
    ret = &BookChange{
        title: title,
        description: description,
        language: language,
        releaseDate: releaseDate,
        edition: edition,
        category: category,
        author: author,
    }
    
    return
}



type ChangeBookLocation struct {
    shelf string
    fold string
}

func NewChangeBookLocation(shelf string, fold string) (ret *ChangeBookLocation, err error) {
    ret = &ChangeBookLocation{
        shelf: shelf,
        fold: fold,
    }
    
    return
}




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

func (o *BookCommandType) IsUpdateBook() bool {
    return o == _bookCommandTypes.UpdateBook()
}

func (o *BookCommandType) IsRegisterBook() bool {
    return o == _bookCommandTypes.RegisterBook()
}

func (o *BookCommandType) IsUnregisterBook() bool {
    return o == _bookCommandTypes.UnregisterBook()
}

func (o *BookCommandType) IsChangeBook() bool {
    return o == _bookCommandTypes.ChangeBook()
}

func (o *BookCommandType) IsChangeLBookocation() bool {
    return o == _bookCommandTypes.ChangeLBookocation()
}

type bookCommandTypes struct {
	values []*BookCommandType
    literals []enum.Literal
}

var _bookCommandTypes = &bookCommandTypes{values: []*BookCommandType{
    {name: "createBook", ordinal: 0},
    {name: "deleteBook", ordinal: 1},
    {name: "updateBook", ordinal: 2},
    {name: "registerBook", ordinal: 3},
    {name: "unregisterBook", ordinal: 4},
    {name: "changeBook", ordinal: 5},
    {name: "changeLBookocation", ordinal: 6}},
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

func (o *bookCommandTypes) UpdateBook() *BookCommandType {
    return _bookCommandTypes.values[2]
}

func (o *bookCommandTypes) RegisterBook() *BookCommandType {
    return _bookCommandTypes.values[3]
}

func (o *bookCommandTypes) UnregisterBook() *BookCommandType {
    return _bookCommandTypes.values[4]
}

func (o *bookCommandTypes) ChangeBook() *BookCommandType {
    return _bookCommandTypes.values[5]
}

func (o *bookCommandTypes) ChangeLBookocation() *BookCommandType {
    return _bookCommandTypes.values[6]
}

func (o *bookCommandTypes) ParseBookCommandType(name string) (ret *BookCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*BookCommandType), ok
	}
	return
}



