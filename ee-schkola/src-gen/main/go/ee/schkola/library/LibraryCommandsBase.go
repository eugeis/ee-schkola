package library

import (
    "ee/schkola/person"
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/enum"
    "time"
)
const (
     CreateBookCommand eventhorizon.CommandType = "CreateBook"
     DeleteBookCommand eventhorizon.CommandType = "DeleteBook"
     UpdateBookCommand eventhorizon.CommandType = "UpdateBook"
     RegisterBookCommand eventhorizon.CommandType = "RegisterBook"
     UnregisterBookCommand eventhorizon.CommandType = "UnregisterBook"
     ChangeBookCommand eventhorizon.CommandType = "ChangeBook"
     ChangeBookLocationCommand eventhorizon.CommandType = "ChangeBookLocation"
)




        

type CreateBook struct {
    Title  
    Description  string
    Language  string
    ReleaseDate  *time.Time
    Edition  string
    Category  string
    Author  *person.PersonName
    Location  *Location
}

func NewCreateBook(title , description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName, location *Location) (ret *CreateBook, err error) {
    ret = &CreateBook{
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
func (o *CreateBook) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *CreateBook) AggregateType() eventhorizon.AggregateType  { return BookAggregateType }
func (o *CreateBook) CommandType() eventhorizon.CommandType      { return CreateBookCommand }



        

type DeleteBook struct {
    Id  string
}

func NewDeleteBook(id string) (ret *DeleteBook, err error) {
    ret = &DeleteBook{
        Id : id,
    }
    
    return
}
func (o *DeleteBook) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *DeleteBook) AggregateType() eventhorizon.AggregateType  { return BookAggregateType }
func (o *DeleteBook) CommandType() eventhorizon.CommandType      { return DeleteBookCommand }



        

type UpdateBook struct {
    Title  
    Description  string
    Language  string
    ReleaseDate  *time.Time
    Edition  string
    Category  string
    Author  *person.PersonName
    Location  *Location
}

func NewUpdateBook(title , description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName, location *Location) (ret *UpdateBook, err error) {
    ret = &UpdateBook{
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
func (o *UpdateBook) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *UpdateBook) AggregateType() eventhorizon.AggregateType  { return BookAggregateType }
func (o *UpdateBook) CommandType() eventhorizon.CommandType      { return UpdateBookCommand }



        

type RegisterBook struct {
    Title  
    Description  string
    Language  string
    ReleaseDate  *time.Time
    Edition  string
    Category  string
    Author  *person.PersonName
}

func NewRegisterBook(title , description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName) (ret *RegisterBook, err error) {
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
func (o *RegisterBook) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *RegisterBook) AggregateType() eventhorizon.AggregateType  { return BookAggregateType }
func (o *RegisterBook) CommandType() eventhorizon.CommandType      { return RegisterBookCommand }



        

type UnregisterBook struct {
    Id  string
}

func NewUnregisterBook(id string) (ret *UnregisterBook, err error) {
    ret = &UnregisterBook{
        Id : id,
    }
    
    return
}
func (o *UnregisterBook) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *UnregisterBook) AggregateType() eventhorizon.AggregateType  { return BookAggregateType }
func (o *UnregisterBook) CommandType() eventhorizon.CommandType      { return UnregisterBookCommand }



        

type ChangeBook struct {
    Title  
    Description  string
    Language  string
    ReleaseDate  *time.Time
    Edition  string
    Category  string
    Author  *person.PersonName
}

func NewChangeBook(title , description string, language string, releaseDate *time.Time, edition string, category string, 
                author *person.PersonName) (ret *ChangeBook, err error) {
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
func (o *ChangeBook) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *ChangeBook) AggregateType() eventhorizon.AggregateType  { return BookAggregateType }
func (o *ChangeBook) CommandType() eventhorizon.CommandType      { return ChangeBookCommand }



        

type ChangeBookLocation struct {
    Shelf  string
    Fold  string
}

func NewChangeBookLocation(shelf string, fold string) (ret *ChangeBookLocation, err error) {
    ret = &ChangeBookLocation{
        Shelf : shelf,
        Fold : fold,
    }
    
    return
}
func (o *ChangeBookLocation) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *ChangeBookLocation) AggregateType() eventhorizon.AggregateType  { return BookAggregateType }
func (o *ChangeBookLocation) CommandType() eventhorizon.CommandType      { return ChangeBookLocationCommand }





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

func (o *BookCommandType) IsBookCreate() bool {
    return o == _bookCommandTypes.BookCreate()
}

func (o *BookCommandType) IsBookDelete() bool {
    return o == _bookCommandTypes.BookDelete()
}

func (o *BookCommandType) IsBookUpdate() bool {
    return o == _bookCommandTypes.BookUpdate()
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
    literals []enum.Literal
}

var _bookCommandTypes = &bookCommandTypes{values: []*BookCommandType{
    {name: "BookCreate", ordinal: 0},
    {name: "BookDelete", ordinal: 1},
    {name: "BookUpdate", ordinal: 2},
    {name: "BookRegister", ordinal: 3},
    {name: "BookUnregister", ordinal: 4},
    {name: "BookChange", ordinal: 5},
    {name: "changeLBookocation", ordinal: 6}},
}

func BookCommandTypes() *bookCommandTypes {
	return _bookCommandTypes
}

func (o *bookCommandTypes) Values() []*BookCommandType {
	return o.values
}

func (o *bookCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *bookCommandTypes) BookCreate() *BookCommandType {
    return _bookCommandTypes.values[0]
}

func (o *bookCommandTypes) BookDelete() *BookCommandType {
    return _bookCommandTypes.values[1]
}

func (o *bookCommandTypes) BookUpdate() *BookCommandType {
    return _bookCommandTypes.values[2]
}

func (o *bookCommandTypes) BookRegister() *BookCommandType {
    return _bookCommandTypes.values[3]
}

func (o *bookCommandTypes) BookUnregister() *BookCommandType {
    return _bookCommandTypes.values[4]
}

func (o *bookCommandTypes) BookChange() *BookCommandType {
    return _bookCommandTypes.values[5]
}

func (o *bookCommandTypes) ChangeLBookocation() *BookCommandType {
    return _bookCommandTypes.values[6]
}

func (o *bookCommandTypes) ParseBookCommandType(name string) (ret *BookCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*BookCommandType), ok
	}
	return
}



