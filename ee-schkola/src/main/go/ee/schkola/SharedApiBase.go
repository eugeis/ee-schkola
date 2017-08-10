package schkola

import (
    "github.com/looplab/eventhorizon"
    "time"
)
type SchkolaBase struct {
    *Trace
    Id eventhorizon.UUID
}

func NewSchkolaBase() (ret *SchkolaBase) {
    ret = &SchkolaBase{
        Trace: NewTrace(),
    }
    return
}






type Trace struct {
    CreatedAt *time.Time
    UpdatedAt *time.Time
    ModifiedBy string
}

func NewTrace() (ret *Trace) {
    ret = &Trace{}
    return
}





