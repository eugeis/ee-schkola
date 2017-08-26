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

func (o *AccountQueryRepository) FindAll() (ret []*Account) {
    
    return
}

func (o *AccountQueryRepository) FindById(id eventhorizon.UUID) (ret *Account) {
    
    return
}

func (o *AccountQueryRepository) CountAll() (ret int) {
    
    return
}

func (o *AccountQueryRepository) CountById(id eventhorizon.UUID) (ret int) {
    
    return
}

func (o *AccountQueryRepository) ExistAll() (ret bool) {
    
    return
}

func (o *AccountQueryRepository) ExistById(id eventhorizon.UUID) (ret bool) {
    
    return
}









