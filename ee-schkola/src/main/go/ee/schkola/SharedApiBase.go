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
    trace := NewTrace()
    ret = &SchkolaBase{
        Trace: trace,
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





