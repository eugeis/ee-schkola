package library

import (
    "ee/schkola"
    "ee/schkola/person"
    "time"
)
type Book struct {
    Title  
    Description  string
    Language  
    ReleaseDate  *time.Time
    Edition  
    Category  
    Author  *person.PersonName
    Location  *Location
    *schkola.SchkolaBase
}

func NewBook(title , description string, language , releaseDate *time.Time, edition , category , author *person.PersonName, 
                location *Location, SchkolaBase *schkola.SchkolaBase) (ret *Book, err error) {
    ret = (id)
        
    ret.Title = title
    ret.Description = description
    ret.Language = language
    ret.ReleaseDate = releaseDate
    ret.Edition = edition
    ret.Category = category
    ret.Author = author
    ret.Location = location
    ret.SchkolaBase = SchkolaBase
    return
}






type Location struct {
    Shelf  
    Fold  
}

func NewLocation(shelf , fold ) (ret *Location, err error) {
    ret = &Location{
        Shelf : shelf,
        Fold : fold,
    }
    return
}





