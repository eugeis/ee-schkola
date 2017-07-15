package schkola

import (
    "github.com/looplab/eventhorizon"
    "time"
)

type SchkolaBase struct {
    *Trace
    Id  eventhorizon.UUID
}

func NewSchkolaBase(id eventhorizon.UUID) (ret *SchkolaBase) {
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

func NewTrace(createdAt *time.Time, updatedAt *time.Time, modifiedBy string) (ret *Trace) {
    ret = &Trace{
        CreatedAt : createdAt,
        UpdatedAt : updatedAt,
        ModifiedBy : modifiedBy,
    }
    
    return
}





