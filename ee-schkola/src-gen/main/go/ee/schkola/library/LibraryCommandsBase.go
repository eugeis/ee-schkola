package library

import (
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "time"
)

type CreateBook struct {
    Title  
    Description  string
    Language  string
    ReleaseDate  *time.Time
    Edition  string
    Category  string
    Author  *person.PersonName
    Location  *Location
}

func NewCreateBook(title , description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName, location *Location) (ret *CreateBook, err error) {
    ret = &CreateBook{
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



type DeleteBook struct {
    Id  string
}

func NewDeleteBook(id string) (ret *DeleteBook, err error) {
    ret = &DeleteBook{
        Id : id,
    }
    return
}



type UpdateBook struct {
    Title  
    Description  string
    Language  string
    ReleaseDate  *time.Time
    Edition  string
    Category  string
    Author  *person.PersonName
    Location  *Location
}

func NewUpdateBook(title , description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName, location *Location) (ret *UpdateBook, err error) {
    ret = &UpdateBook{
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



type RegisterBook struct {
    Title  
    Description  string
    Language  string
    ReleaseDate  *time.Time
    Edition  string
    Category  string
    Author  *person.PersonName
}

func NewRegisterBook(title , description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName) (ret *RegisterBook, err error) {
    ret = &RegisterBook{
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



type UnregisterBook struct {
    Id  string
}

func NewUnregisterBook(id string) (ret *UnregisterBook, err error) {
    ret = &UnregisterBook{
        Id : id,
    }
    return
}



type ChangeBook struct {
    Title  
    Description  string
    Language  string
    ReleaseDate  *time.Time
    Edition  string
    Category  string
    Author  *person.PersonName
}

func NewChangeBook(title , description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName) (ret *ChangeBook, err error) {
    ret = &ChangeBook{
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

func (o *BookCommandType) IsBookCreate() bool {
    return o == _bookCommandTypes.BookCreate()
}

func (o *BookCommandType) IsBookDelete() bool {
    return o == _bookCommandTypes.BookDelete()
}

func (o *BookCommandType) IsBookUpdate() bool {
    return o == _bookCommandTypes.BookUpdate()
}

func (o *BookCommandType) IsBookRegister() bool {
    return o == _bookCommandTypes.BookRegister()
}

func (o *BookCommandType) IsBookUnregister() bool {
    return o == _bookCommandTypes.BookUnregister()
}

func (o *BookCommandType) IsBookChange() bool {
    return o == _bookCommandTypes.BookChange()
}

func (o *BookCommandType) IsChangeLBookocation() bool {
    return o == _bookCommandTypes.ChangeLBookocation()
}

type bookCommandTypes struct {
	values []*BookCommandType
    literals []enum.Literal
}

var _bookCommandTypes = &bookCommandTypes{values: []*BookCommandType{
    {name: "BookCreate", ordinal: 0},
    {name: "BookDelete", ordinal: 1},
    {name: "BookUpdate", ordinal: 2},
    {name: "BookRegister", ordinal: 3},
    {name: "BookUnregister", ordinal: 4},
    {name: "BookChange", ordinal: 5},
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

func (o *bookCommandTypes) BookCreate() *BookCommandType {
    return _bookCommandTypes.values[0]
}

func (o *bookCommandTypes) BookDelete() *BookCommandType {
    return _bookCommandTypes.values[1]
}

func (o *bookCommandTypes) BookUpdate() *BookCommandType {
    return _bookCommandTypes.values[2]
}

func (o *bookCommandTypes) BookRegister() *BookCommandType {
    return _bookCommandTypes.values[3]
}

func (o *bookCommandTypes) BookUnregister() *BookCommandType {
    return _bookCommandTypes.values[4]
}

func (o *bookCommandTypes) BookChange() *BookCommandType {
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



