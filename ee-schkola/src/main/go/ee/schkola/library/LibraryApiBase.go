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
    schkolaBase := schkola.NewSchkolaBase()
    ret = &Book{
        SchkolaBase: schkolaBase,
    }
    return
}

func (o *Book) FindByAuthor(author *person.PersonName)  {
    
}

func (o *Book) FindByPattern(pattern string)  {
    
}

func (o *Book) FindByTitle(title string)  {
    
}






type Location struct {
    Shelf string
    Fold string
}

func NewLocation() (ret *Location) {
    ret = &Location{}
    return
}





