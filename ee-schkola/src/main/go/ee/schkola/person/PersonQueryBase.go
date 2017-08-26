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

func (o *ChurchQueryRepository) FindAll() (ret []*Church, err error) {
    
    return
}

func (o *ChurchQueryRepository) FindById(id eventhorizon.UUID) (ret *Church, err error) {
    
    return
}

func (o *ChurchQueryRepository) CountAll() (ret int, err error) {
    
    return
}

func (o *ChurchQueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    
    return
}

func (o *ChurchQueryRepository) ExistAll() (ret bool, err error) {
    
    return
}

func (o *ChurchQueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    
    return
}


type GraduationQueryRepository struct {
}

func NewGraduationQueryRepository() (ret *GraduationQueryRepository) {
    ret = &GraduationQueryRepository{}
    return
}

func (o *GraduationQueryRepository) FindAll() (ret []*Graduation, err error) {
    
    return
}

func (o *GraduationQueryRepository) FindById(id eventhorizon.UUID) (ret *Graduation, err error) {
    
    return
}

func (o *GraduationQueryRepository) CountAll() (ret int, err error) {
    
    return
}

func (o *GraduationQueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    
    return
}

func (o *GraduationQueryRepository) ExistAll() (ret bool, err error) {
    
    return
}

func (o *GraduationQueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    
    return
}


type ProfileQueryRepository struct {
}

func NewProfileQueryRepository() (ret *ProfileQueryRepository) {
    ret = &ProfileQueryRepository{}
    return
}

func (o *ProfileQueryRepository) FindByEmail(email string) (ret *Profile, err error) {
    
    return
}

func (o *ProfileQueryRepository) FindByPhone(phone string) (ret *Profile, err error) {
    
    return
}

func (o *ProfileQueryRepository) FindAll() (ret []*Profile, err error) {
    
    return
}

func (o *ProfileQueryRepository) FindById(id eventhorizon.UUID) (ret *Profile, err error) {
    
    return
}

func (o *ProfileQueryRepository) CountAll() (ret int, err error) {
    
    return
}

func (o *ProfileQueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    
    return
}

func (o *ProfileQueryRepository) ExistAll() (ret bool, err error) {
    
    return
}

func (o *ProfileQueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    
    return
}









