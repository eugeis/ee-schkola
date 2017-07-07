package schkola

import (
    "time"
)


type SchkolaBase struct {
    Trace  *Trace
    @@EMPTY@@  string
}

func NewSchkolaBase(@@EMPTY@@ string) (ret SchkolaBase, err error) {
    ret = SchkolaBase{
        @@EMPTY@@ : @@EMPTY@@,
    }
    return
}


type Trace struct {
    CreatedAt  *time.Time
    UpdatedAt  *time.Time
    ModifiedBy  string
}

func NewTrace(createdAt *time.Time, updatedAt *time.Time, modifiedBy string) (ret Trace, err error) {
    ret = Trace{
        CreatedAt : createdAt,
        UpdatedAt : updatedAt,
        ModifiedBy : modifiedBy,
    }
    return
}




type SchkolaBase struct {
    Trace  *Trace
    @@EMPTY@@  string
}

func NewSchkolaBase(@@EMPTY@@ string) (ret SchkolaBase, err error) {
    ret = SchkolaBase{
        @@EMPTY@@ : @@EMPTY@@,
    }
    return
}


type Trace struct {
    CreatedAt  *time.Time
    UpdatedAt  *time.Time
    ModifiedBy  string
}

func NewTrace(createdAt *time.Time, updatedAt *time.Time, modifiedBy string) (ret Trace, err error) {
    ret = Trace{
        CreatedAt : createdAt,
        UpdatedAt : updatedAt,
        ModifiedBy : modifiedBy,
    }
    return
}



