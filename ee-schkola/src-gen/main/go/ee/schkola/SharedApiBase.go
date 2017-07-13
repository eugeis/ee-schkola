package schkola

import (
    "time"
)
type SchkolaBase struct {
    Trace  *Trace
    Id  string
}

func NewSchkolaBase(id string) (ret *SchkolaBase, err error) {
    ret = &SchkolaBase{
        Id : id,
    }
    return
}






type Trace struct {
    CreatedAt  *time.Time
    UpdatedAt  *time.Time
    ModifiedBy  string
}

func NewTrace(createdAt *time.Time, updatedAt *time.Time, modifiedBy string) (ret *Trace, err error) {
    ret = &Trace{
        CreatedAt : createdAt,
        UpdatedAt : updatedAt,
        ModifiedBy : modifiedBy,
    }
    return
}





