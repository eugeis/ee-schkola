package library

import (
    "ee/schkola"
    "ee/schkola/person"
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
    *schkola.SchkolaBase
}

func NewBook() (ret *Book) {
    ret = &Book{
        SchkolaBase: schkola.NewSchkolaBase(),
    }
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





