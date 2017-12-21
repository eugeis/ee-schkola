package library

import (
    "ee/schkola/shared"
    "github.com/eugeis/gee/eh"
    "github.com/looplab/eventhorizon"
    "time"
)
        
type Book struct {
    Title string `json:"title" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Language string `json:"language" eh:"optional"`
    ReleaseDate *time.Time `json:"releaseDate" eh:"optional"`
    Edition string `json:"edition" eh:"optional"`
    Category string `json:"category" eh:"optional"`
    Author *shared.PersonName `json:"author" eh:"optional"`
    Location *Location `json:"location" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
    *ChangeLocation
}

func NewBook() (ret *Book) {
    changeLocation := &ChangeLocation{}
    ret = &Book{
        ChangeLocation: changeLocation,
    }
    return
}

func (o *Book) FindByPattern(pattern string) (ret *Book, err error) {
    err = eh.QueryNotImplemented("findByPattern")
    return
}

func (o *Book) FindByTitle(title string) (ret *Book, err error) {
    err = eh.QueryNotImplemented("findByTitle")
    return
}
func (o *Book) EntityID() eventhorizon.UUID { return o.Id }







type Location struct {
    Shelf string `json:"shelf" eh:"optional"`
    Fold string `json:"fold" eh:"optional"`
}

func NewLocation() (ret *Location) {
    ret = &Location{}
    return
}





