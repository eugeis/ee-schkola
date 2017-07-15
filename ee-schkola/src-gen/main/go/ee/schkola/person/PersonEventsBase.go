package person

import (
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/enum"
    "time"
)

type ChurchCreated struct {
    Id  eventhorizon.UUID
    Name  string
    Address  *Address
    Pastor  *PersonName
    Contact  *Contact
}



type ChurchDeleted struct {
    Id  eventhorizon.UUID
}



type ChurchUpdated struct {
    Id  eventhorizon.UUID
    Name  string
    Address  *Address
    Pastor  *PersonName
    Contact  *Contact
}



type GraduationCreated struct {
    Id  eventhorizon.UUID
    Name  string
    Level  *GraduationLevel
}



type GraduationDeleted struct {
    Id  eventhorizon.UUID
}



type GraduationUpdated struct {
    Id  eventhorizon.UUID
    Name  string
    Level  *GraduationLevel
}



type ProfileCreated struct {
    Id  eventhorizon.UUID
    Gender  *Gender
    Name  *PersonName
    BirthName  string
    Birthday  *time.Time
    Address  *Address
    Contact  *Contact
    PhotoData  []byte
    Photo  string
    Family  *Family
    Church  *ChurchInfo
    Education  *Education
}



type ProfileDeleted struct {
    Id  eventhorizon.UUID
}



type ProfileUpdated struct {
    Id  eventhorizon.UUID
    Gender  *Gender
    Name  *PersonName
    BirthName  string
    Birthday  *time.Time
    Address  *Address
    Contact  *Contact
    PhotoData  []byte
    Photo  string
    Family  *Family
    Church  *ChurchInfo
    Education  *Education
}




type ChurchEventType struct {
	name  string
	ordinal int
}

func (o *ChurchEventType) Name() string {
    return o.name
}

func (o *ChurchEventType) Ordinal() int {
    return o.ordinal
}

func (o *ChurchEventType) IsChurchCreated() bool {
    return o == _churchEventTypes.ChurchCreated()
}

func (o *ChurchEventType) IsChurchDeleted() bool {
    return o == _churchEventTypes.ChurchDeleted()
}

func (o *ChurchEventType) IsChurchUpdated() bool {
    return o == _churchEventTypes.ChurchUpdated()
}

type churchEventTypes struct {
	values []*ChurchEventType
    literals []enum.Literal
}

var _churchEventTypes = &churchEventTypes{values: []*ChurchEventType{
    {name: "ChurchCreated", ordinal: 0},
    {name: "ChurchDeleted", ordinal: 1},
    {name: "ChurchUpdated", ordinal: 2}},
}

func ChurchEventTypes() *churchEventTypes {
	return _churchEventTypes
}

func (o *churchEventTypes) Values() []*ChurchEventType {
	return o.values
}

func (o *churchEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *churchEventTypes) ChurchCreated() *ChurchEventType {
    return _churchEventTypes.values[0]
}

func (o *churchEventTypes) ChurchDeleted() *ChurchEventType {
    return _churchEventTypes.values[1]
}

func (o *churchEventTypes) ChurchUpdated() *ChurchEventType {
    return _churchEventTypes.values[2]
}

func (o *churchEventTypes) ParseChurchEventType(name string) (ret *ChurchEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*ChurchEventType), ok
	}
	return
}


type GraduationEventType struct {
	name  string
	ordinal int
}

func (o *GraduationEventType) Name() string {
    return o.name
}

func (o *GraduationEventType) Ordinal() int {
    return o.ordinal
}

func (o *GraduationEventType) IsGraduationCreated() bool {
    return o == _graduationEventTypes.GraduationCreated()
}

func (o *GraduationEventType) IsGraduationDeleted() bool {
    return o == _graduationEventTypes.GraduationDeleted()
}

func (o *GraduationEventType) IsGraduationUpdated() bool {
    return o == _graduationEventTypes.GraduationUpdated()
}

type graduationEventTypes struct {
	values []*GraduationEventType
    literals []enum.Literal
}

var _graduationEventTypes = &graduationEventTypes{values: []*GraduationEventType{
    {name: "GraduationCreated", ordinal: 0},
    {name: "GraduationDeleted", ordinal: 1},
    {name: "GraduationUpdated", ordinal: 2}},
}

func GraduationEventTypes() *graduationEventTypes {
	return _graduationEventTypes
}

func (o *graduationEventTypes) Values() []*GraduationEventType {
	return o.values
}

func (o *graduationEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *graduationEventTypes) GraduationCreated() *GraduationEventType {
    return _graduationEventTypes.values[0]
}

func (o *graduationEventTypes) GraduationDeleted() *GraduationEventType {
    return _graduationEventTypes.values[1]
}

func (o *graduationEventTypes) GraduationUpdated() *GraduationEventType {
    return _graduationEventTypes.values[2]
}

func (o *graduationEventTypes) ParseGraduationEventType(name string) (ret *GraduationEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*GraduationEventType), ok
	}
	return
}


type ProfileEventType struct {
	name  string
	ordinal int
}

func (o *ProfileEventType) Name() string {
    return o.name
}

func (o *ProfileEventType) Ordinal() int {
    return o.ordinal
}

func (o *ProfileEventType) IsProfileCreated() bool {
    return o == _profileEventTypes.ProfileCreated()
}

func (o *ProfileEventType) IsProfileDeleted() bool {
    return o == _profileEventTypes.ProfileDeleted()
}

func (o *ProfileEventType) IsProfileUpdated() bool {
    return o == _profileEventTypes.ProfileUpdated()
}

type profileEventTypes struct {
	values []*ProfileEventType
    literals []enum.Literal
}

var _profileEventTypes = &profileEventTypes{values: []*ProfileEventType{
    {name: "ProfileCreated", ordinal: 0},
    {name: "ProfileDeleted", ordinal: 1},
    {name: "ProfileUpdated", ordinal: 2}},
}

func ProfileEventTypes() *profileEventTypes {
	return _profileEventTypes
}

func (o *profileEventTypes) Values() []*ProfileEventType {
	return o.values
}

func (o *profileEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *profileEventTypes) ProfileCreated() *ProfileEventType {
    return _profileEventTypes.values[0]
}

func (o *profileEventTypes) ProfileDeleted() *ProfileEventType {
    return _profileEventTypes.values[1]
}

func (o *profileEventTypes) ProfileUpdated() *ProfileEventType {
    return _profileEventTypes.values[2]
}

func (o *profileEventTypes) ParseProfileEventType(name string) (ret *ProfileEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*ProfileEventType), ok
	}
	return
}



