package auth

import (
    "github.com/looplab/eventhorizon"
)
type AccountQueryRepository struct {
}

func NewAccountQueryRepository() (ret *AccountQueryRepository) {
    ret = &AccountQueryRepository{}
    return
}

func (o *AccountQueryRepository) FindAll() (ret []*Account, err error) {
    
    return
}

func (o *AccountQueryRepository) FindById(id eventhorizon.UUID) (ret *Account, err error) {
    
    return
}

func (o *AccountQueryRepository) CountAll() (ret int, err error) {
    
    return
}

func (o *AccountQueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    
    return
}

func (o *AccountQueryRepository) ExistAll() (ret bool, err error) {
    
    return
}

func (o *AccountQueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    
    return
}









