package schkola

import (
    "time"
)
type SchkolaBaseCommands struct {
	name  string
	ordinal int
}

func (o *SchkolaBaseCommands) Name() string {
    return o.name
}

func (o *SchkolaBaseCommands) Ordinal() int {
    return o.ordinal
}

func (o *SchkolaBaseCommands) IsRegisterSchkolaBase() bool {
    return o == _schkolaBaseCommandss.RegisterSchkolaBase()
}

func (o *SchkolaBaseCommands) IsDeleteSchkolaBase() bool {
    return o == _schkolaBaseCommandss.DeleteSchkolaBase()
}

func (o *SchkolaBaseCommands) IsChangeSchkolaBase() bool {
    return o == _schkolaBaseCommandss.ChangeSchkolaBase()
}

type schkolaBaseCommandss struct {
	values []*SchkolaBaseCommands
}

var _schkolaBaseCommandss = &schkolaBaseCommandss{values: []*SchkolaBaseCommands{
    {name: "registerSchkolaBase", ordinal: 0},
    {name: "deleteSchkolaBase", ordinal: 1},
    {name: "changeSchkolaBase", ordinal: 2}},
}

func SchkolaBaseCommandss() *schkolaBaseCommandss {
	return _schkolaBaseCommandss
}

func (o *schkolaBaseCommandss) Values() []*SchkolaBaseCommands {
	return o.values
}

func (o *schkolaBaseCommandss) RegisterSchkolaBase() *SchkolaBaseCommands {
    return _schkolaBaseCommandss.values[0]
}

func (o *schkolaBaseCommandss) DeleteSchkolaBase() *SchkolaBaseCommands {
    return _schkolaBaseCommandss.values[1]
}

func (o *schkolaBaseCommandss) ChangeSchkolaBase() *SchkolaBaseCommands {
    return _schkolaBaseCommandss.values[2]
}

func (o *schkolaBaseCommandss) ParseSchkolaBaseCommands(name string) (ret *SchkolaBaseCommands, ok bool) {
    switch name {
      case "RegisterSchkolaBase":
        ret = o.RegisterSchkolaBase()
      case "DeleteSchkolaBase":
        ret = o.DeleteSchkolaBase()
      case "ChangeSchkolaBase":
        ret = o.ChangeSchkolaBase()
    }
    return
}




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



