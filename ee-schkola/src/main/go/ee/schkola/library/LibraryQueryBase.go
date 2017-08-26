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

func (o *BookQueryRepository) FindByTitle(title string) (ret *Book, err error) {
    
    return
}

func (o *BookQueryRepository) FindByAuthor(author *person.PersonName) (ret *Book, err error) {
    
    return
}

func (o *BookQueryRepository) FindByPattern(pattern string) (ret *Book, err error) {
    
    return
}

func (o *BookQueryRepository) FindAll() (ret []*Book, err error) {
    
    return
}

func (o *BookQueryRepository) FindById(id eventhorizon.UUID) (ret *Book, err error) {
    
    return
}

func (o *BookQueryRepository) CountAll() (ret int, err error) {
    
    return
}

func (o *BookQueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    
    return
}

func (o *BookQueryRepository) ExistAll() (ret bool, err error) {
    
    return
}

func (o *BookQueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    
    return
}









