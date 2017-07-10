package library

import (
    "ee/schkola"
    "ee/schkola/person"
    "time"
)
type BookCommands struct {
	name  string
	ordinal int
}

func (o *BookCommands) Name() string {
    return o.name
}

func (o *BookCommands) Ordinal() int {
    return o.ordinal
}

func (o *BookCommands) IsRegisterBook() bool {
    return o == _bookCommandss.RegisterBook()
}

func (o *BookCommands) IsUnregisterBook() bool {
    return o == _bookCommandss.UnregisterBook()
}

func (o *BookCommands) IsChangeBook() bool {
    return o == _bookCommandss.ChangeBook()
}

func (o *BookCommands) IsChangeBookLocation() bool {
    return o == _bookCommandss.ChangeBookLocation()
}

type bookCommandss struct {
	values []*BookCommands
}

var _bookCommandss = &bookCommandss{values: []*BookCommands{
    {name: "registerBook", ordinal: 0},
    {name: "unregisterBook", ordinal: 1},
    {name: "changeBook", ordinal: 2},
    {name: "changeBookLocation", ordinal: 3}},
}

func BookCommandss() *bookCommandss {
	return _bookCommandss
}

func (o *bookCommandss) Values() []*BookCommands {
	return o.values
}

func (o *bookCommandss) RegisterBook() *BookCommands {
    return _bookCommandss.values[0]
}

func (o *bookCommandss) UnregisterBook() *BookCommands {
    return _bookCommandss.values[1]
}

func (o *bookCommandss) ChangeBook() *BookCommands {
    return _bookCommandss.values[2]
}

func (o *bookCommandss) ChangeBookLocation() *BookCommands {
    return _bookCommandss.values[3]
}

func (o *bookCommandss) ParseBookCommands(name string) (ret *BookCommands, ok bool) {
    switch name {
      case "RegisterBook":
        ret = o.RegisterBook()
      case "UnregisterBook":
        ret = o.UnregisterBook()
      case "ChangeBook":
        ret = o.ChangeBook()
      case "ChangeBookLocation":
        ret = o.ChangeBookLocation()
    }
    return
}




type Book struct {
    Title  string
    Description  string
    Language  string
    ReleaseDate  *time.Time
    Edition  string
    Category  string
    Author  *person.PersonName
    Location  *Location
    *schkola.SchkolaBase
}

func NewBook(title string, description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName, location *Location, SchkolaBase *schkola.SchkolaBase) (ret *Book, err error) {
    ret = &Book{
        Title : title,
        Description : description,
        Language : language,
        ReleaseDate : releaseDate,
        Edition : edition,
        Category : category,
        Author : author,
        Location : location,
        SchkolaBase : SchkolaBase,
    }
    return
}


type Location struct {
    Shelf  string
    Fold  string
}

func NewLocation(shelf string, fold string) (ret *Location, err error) {
    ret = &Location{
        Shelf : shelf,
        Fold : fold,
    }
    return
}



