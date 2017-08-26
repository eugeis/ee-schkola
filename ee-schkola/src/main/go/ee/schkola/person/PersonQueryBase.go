package person

import (
    "github.com/looplab/eventhorizon"
)
type ChurchQueryRepository struct {
}

func NewChurchQueryRepository() (ret *ChurchQueryRepository) {
    ret = &ChurchQueryRepository{}
    return
}

func (o *ChurchQueryRepository) FindAll() (ret []*Church) {
    
    return
}

func (o *ChurchQueryRepository) FindById(id eventhorizon.UUID) (ret *Church) {
    
    return
}

func (o *ChurchQueryRepository) CountAll() (ret int) {
    
    return
}

func (o *ChurchQueryRepository) CountById(id eventhorizon.UUID) (ret int) {
    
    return
}

func (o *ChurchQueryRepository) ExistAll() (ret bool) {
    
    return
}

func (o *ChurchQueryRepository) ExistById(id eventhorizon.UUID) (ret bool) {
    
    return
}


type GraduationQueryRepository struct {
}

func NewGraduationQueryRepository() (ret *GraduationQueryRepository) {
    ret = &GraduationQueryRepository{}
    return
}

func (o *GraduationQueryRepository) FindAll() (ret []*Graduation) {
    
    return
}

func (o *GraduationQueryRepository) FindById(id eventhorizon.UUID) (ret *Graduation) {
    
    return
}

func (o *GraduationQueryRepository) CountAll() (ret int) {
    
    return
}

func (o *GraduationQueryRepository) CountById(id eventhorizon.UUID) (ret int) {
    
    return
}

func (o *GraduationQueryRepository) ExistAll() (ret bool) {
    
    return
}

func (o *GraduationQueryRepository) ExistById(id eventhorizon.UUID) (ret bool) {
    
    return
}


type ProfileQueryRepository struct {
}

func NewProfileQueryRepository() (ret *ProfileQueryRepository) {
    ret = &ProfileQueryRepository{}
    return
}

func (o *ProfileQueryRepository) FindByName(name *PersonName)  {
    
}

func (o *ProfileQueryRepository) FindByEmail(email string)  {
    
}

func (o *ProfileQueryRepository) FindByPhone(phone string)  {
    
}

func (o *ProfileQueryRepository) FindAll() (ret []*Profile) {
    
    return
}

func (o *ProfileQueryRepository) FindById(id eventhorizon.UUID) (ret *Profile) {
    
    return
}

func (o *ProfileQueryRepository) CountAll() (ret int) {
    
    return
}

func (o *ProfileQueryRepository) CountById(id eventhorizon.UUID) (ret int) {
    
    return
}

func (o *ProfileQueryRepository) ExistAll() (ret bool) {
    
    return
}

func (o *ProfileQueryRepository) ExistById(id eventhorizon.UUID) (ret bool) {
    
    return
}









