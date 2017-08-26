package library

import (
    "ee/schkola/person"
    "github.com/looplab/eventhorizon"
)
type BookQueryRepository struct {
}

func NewBookQueryRepository() (ret *BookQueryRepository) {
    ret = &BookQueryRepository{}
    return
}

func (o *BookQueryRepository) FindByTitle(title string)  {
    
}

func (o *BookQueryRepository) FindByAuthor(author *person.PersonName)  {
    
}

func (o *BookQueryRepository) FindByPattern(pattern string)  {
    
}

func (o *BookQueryRepository) FindAll() (ret []*Book) {
    
    return
}

func (o *BookQueryRepository) FindById(id eventhorizon.UUID) (ret *Book) {
    
    return
}

func (o *BookQueryRepository) CountAll() (ret int) {
    
    return
}

func (o *BookQueryRepository) CountById(id eventhorizon.UUID) (ret int) {
    
    return
}

func (o *BookQueryRepository) ExistAll() (ret bool) {
    
    return
}

func (o *BookQueryRepository) ExistById(id eventhorizon.UUID) (ret bool) {
    
    return
}









