package library

import (
    "ee/schkola"
    "ee/schkola/person"
    "time"
)

type Book struct {
    title 
    description string
    language string
    releaseDate *time.Time
    edition string
    category string
    author *person.PersonName
    location *Location
    *schkola.SchkolaBase
}

func NewBook(title , description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName, location *Location, SchkolaBase *schkola.SchkolaBase) (ret *Book, err error) {
    ret = &Book{
        title: title,
        description: description,
        language: language,
        releaseDate: releaseDate,
        edition: edition,
        category: category,
        author: author,
        location: location,
        SchkolaBase: SchkolaBase,
    }
    
    return
}







type Location struct {
    shelf string
    fold string
}

func NewLocation(shelf string, fold string) (ret *Location, err error) {
    ret = &Location{
        shelf: shelf,
        fold: fold,
    }
    
    return
}





