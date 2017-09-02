package library

import (
    "ee/schkola/person"
    "github.com/eugeis/gee/eh"
    "github.com/looplab/eventhorizon"
    "time"
)
type Book struct {
    Title string
    Description string
    Language string
    ReleaseDate *time.Time
    Edition string
    Category string
    Author *person.PersonName
    Location *Location
    Id eventhorizon.UUID
}

func NewBook() (ret *Book) {
    ret = &Book{}
    return
}

func (o *Book) FindByPattern(pattern string) (ret *Book, err error) {
    err = eh.QueryNotImplemented("findBookByPattern")
    return
}

func (o *Book) FindByTitle(title string) (ret *Book, err error) {
    err = eh.QueryNotImplemented("findBookByTitle")
    return
}






type Location struct {
    Shelf string
    Fold string
}

func NewLocation() (ret *Location) {
    ret = &Location{}
    return
}





