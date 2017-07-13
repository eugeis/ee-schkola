package library

import (
    "ee/schkola/person"
    "time"
)
type RegisterBook struct {
    Title  
    Description  string
    Language  
    ReleaseDate  *time.Time
    Edition  
    Category  
    Author  *person.PersonName
}

func NewRegisterBook(title , description string, language , releaseDate *time.Time, edition , category , author *person.PersonName) (ret *RegisterBook, err error) {
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
    Id  
}

func NewUnregisterBook(id ) (ret *UnregisterBook, err error) {
    ret = &UnregisterBook{
        Id : id,
    }
    return
}


type ChangeBook struct {
    Title  
    Description  string
    Language  
    ReleaseDate  *time.Time
    Edition  
    Category  
    Author  *person.PersonName
}

func NewChangeBook(title , description string, language , releaseDate *time.Time, edition , category , author *person.PersonName) (ret *ChangeBook, err error) {
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
    Shelf  
    Fold  
}

func NewChangeBookLocation(shelf , fold ) (ret *ChangeBookLocation, err error) {
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
}

var _bookCommandTypes = &bookCommandTypes{values: []*BookCommandType{
    {name: "BookRegister", ordinal: 0},
    {name: "BookUnregister", ordinal: 1},
    {name: "BookChange", ordinal: 2},
    {name: "changeLBookocation", ordinal: 3}},
}

func BookCommandTypes() *bookCommandTypes {
	return _bookCommandTypes
}

func (o *bookCommandTypes) Values() []*BookCommandType {
	return o.values
}

func (o *bookCommandTypes) BookRegister() *BookCommandType {
    return _bookCommandTypes.values[0]
}

func (o *bookCommandTypes) BookUnregister() *BookCommandType {
    return _bookCommandTypes.values[1]
}

func (o *bookCommandTypes) BookChange() *BookCommandType {
    return _bookCommandTypes.values[2]
}

func (o *bookCommandTypes) ChangeLBookocation() *BookCommandType {
    return _bookCommandTypes.values[3]
}

func (o *bookCommandTypes) ParseBookCommandType(name string) (ret *BookCommandType, ok bool) {
    switch name {
      case "BookRegister":
        ret = o.BookRegister()
      case "BookUnregister":
        ret = o.BookUnregister()
      case "BookChange":
        ret = o.BookChange()
      case "ChangeLBookocation":
        ret = o.ChangeLBookocation()
    }
    return
}



