package library

import (
    "ee/schkola/person"
    "time"
)


type Book struct {
    Title  string
    Description  string
    Language  string
    ReleaseDate  *time.Time
    Edition  string
    Category  string
    Author  *person.PersonName
    Location  *Location
}

func NewBook(@@EMPTY@@ string, title string, description string, language string, releaseDate *time.Time, edition string, 
                category string, author *person.PersonName, location *Location) (ret Book, err error) {
    ret = Book{
        @@EMPTY@@ : @@EMPTY@@,
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


type Location struct {
    Shelf  string
    Fold  string
}

func NewLocation(shelf string, fold string) (ret Location, err error) {
    ret = Location{
        Shelf : shelf,
        Fold : fold,
    }
    return
}



