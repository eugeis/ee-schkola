package library

import (
    "ee/schkola"
    "ee/schkola/person"
    "time"
)
type Book struct {
    Title  
    Description  string
    Language  string
    ReleaseDate  *time.Time
    Edition  string
    Category  string
    Author  *person.PersonName
    Location  *Location
    *schkola.SchkolaBase
}

func NewBook(title , description string, language string, releaseDate *time.Time, edition string, category string, 
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





