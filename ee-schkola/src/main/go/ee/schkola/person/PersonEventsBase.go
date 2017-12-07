package person

import (
    "ee/schkola/shared"
    "encoding/json"
    "fmt"
    "github.com/eugeis/gee/enum"
    "github.com/looplab/eventhorizon"
    "gopkg.in/mgo.v2/bson"
    "time"
)
const (
     CreateEvent eventhorizon.EventType = "Create"
     CreatedEvent eventhorizon.EventType = "Created"
     DeleteEvent eventhorizon.EventType = "Delete"
     DeletedEvent eventhorizon.EventType = "Deleted"
     UpdateEvent eventhorizon.EventType = "Update"
     UpdatedEvent eventhorizon.EventType = "Updated"
)


const (
     CreateEvent eventhorizon.EventType = "Create"
     CreatedEvent eventhorizon.EventType = "Created"
     DeleteEvent eventhorizon.EventType = "Delete"
     DeletedEvent eventhorizon.EventType = "Deleted"
     UpdateEvent eventhorizon.EventType = "Update"
     UpdatedEvent eventhorizon.EventType = "Updated"
)


const (
     CreateEvent eventhorizon.EventType = "Create"
     CreatedEvent eventhorizon.EventType = "Created"
     DeleteEvent eventhorizon.EventType = "Delete"
     DeletedEvent eventhorizon.EventType = "Deleted"
     UpdateEvent eventhorizon.EventType = "Update"
     UpdatedEvent eventhorizon.EventType = "Updated"
)




type Create struct {
    Name string `json:"name" eh:"optional"`
    Address *Address `json:"address" eh:"optional"`
    Pastor *shared.PersonName `json:"pastor" eh:"optional"`
    Contact *Contact `json:"contact" eh:"optional"`
    Association string `json:"association" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Created struct {
    Name string `json:"name" eh:"optional"`
    Address *Address `json:"address" eh:"optional"`
    Pastor *shared.PersonName `json:"pastor" eh:"optional"`
    Contact *Contact `json:"contact" eh:"optional"`
    Association string `json:"association" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Deleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Update struct {
    Name string `json:"name" eh:"optional"`
    Address *Address `json:"address" eh:"optional"`
    Pastor *shared.PersonName `json:"pastor" eh:"optional"`
    Contact *Contact `json:"contact" eh:"optional"`
    Association string `json:"association" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Updated struct {
    Name string `json:"name" eh:"optional"`
    Address *Address `json:"address" eh:"optional"`
    Pastor *shared.PersonName `json:"pastor" eh:"optional"`
    Contact *Contact `json:"contact" eh:"optional"`
    Association string `json:"association" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Create struct {
    Name string `json:"name" eh:"optional"`
    Level *GraduationLevel `json:"level" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Created struct {
    Name string `json:"name" eh:"optional"`
    Level *GraduationLevel `json:"level" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Deleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Update struct {
    Name string `json:"name" eh:"optional"`
    Level *GraduationLevel `json:"level" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Updated struct {
    Name string `json:"name" eh:"optional"`
    Level *GraduationLevel `json:"level" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Create struct {
    Gender *Gender `json:"gender" eh:"optional"`
    Name *shared.PersonName `json:"name" eh:"optional"`
    BirthName string `json:"birthName" eh:"optional"`
    Birthday *time.Time `json:"birthday" eh:"optional"`
    Address *Address `json:"address" eh:"optional"`
    Contact *Contact `json:"contact" eh:"optional"`
    PhotoData []byte `json:"photoData" eh:"optional"`
    Photo string `json:"photo" eh:"optional"`
    Family *Family `json:"family" eh:"optional"`
    Church *ChurchInfo `json:"church" eh:"optional"`
    Education *Education `json:"education" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Created struct {
    Gender *Gender `json:"gender" eh:"optional"`
    Name *shared.PersonName `json:"name" eh:"optional"`
    BirthName string `json:"birthName" eh:"optional"`
    Birthday *time.Time `json:"birthday" eh:"optional"`
    Address *Address `json:"address" eh:"optional"`
    Contact *Contact `json:"contact" eh:"optional"`
    PhotoData []byte `json:"photoData" eh:"optional"`
    Photo string `json:"photo" eh:"optional"`
    Family *Family `json:"family" eh:"optional"`
    Church *ChurchInfo `json:"church" eh:"optional"`
    Education *Education `json:"education" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Deleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Update struct {
    Gender *Gender `json:"gender" eh:"optional"`
    Name *shared.PersonName `json:"name" eh:"optional"`
    BirthName string `json:"birthName" eh:"optional"`
    Birthday *time.Time `json:"birthday" eh:"optional"`
    Address *Address `json:"address" eh:"optional"`
    Contact *Contact `json:"contact" eh:"optional"`
    PhotoData []byte `json:"photoData" eh:"optional"`
    Photo string `json:"photo" eh:"optional"`
    Family *Family `json:"family" eh:"optional"`
    Church *ChurchInfo `json:"church" eh:"optional"`
    Education *Education `json:"education" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Updated struct {
    Gender *Gender `json:"gender" eh:"optional"`
    Name *shared.PersonName `json:"name" eh:"optional"`
    BirthName string `json:"birthName" eh:"optional"`
    Birthday *time.Time `json:"birthday" eh:"optional"`
    Address *Address `json:"address" eh:"optional"`
    Contact *Contact `json:"contact" eh:"optional"`
    PhotoData []byte `json:"photoData" eh:"optional"`
    Photo string `json:"photo" eh:"optional"`
    Family *Family `json:"family" eh:"optional"`
    Church *ChurchInfo `json:"church" eh:"optional"`
    Education *Education `json:"education" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
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

func (o ChurchEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *ChurchEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := ChurchEventTypes().ParseChurchEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ChurchEventType %q", lit.Name)
        }
	}
	return
}

func (o ChurchEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *ChurchEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := ChurchEventTypes().ParseChurchEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ChurchEventType %q", lit)
        }
    }
    return
}

func (o *ChurchEventType) IsCreate() bool {
    return o == _churchEventTypes.Create()
}

func (o *ChurchEventType) IsCreated() bool {
    return o == _churchEventTypes.Created()
}

func (o *ChurchEventType) IsDelete() bool {
    return o == _churchEventTypes.Delete()
}

func (o *ChurchEventType) IsDeleted() bool {
    return o == _churchEventTypes.Deleted()
}

func (o *ChurchEventType) IsUpdate() bool {
    return o == _churchEventTypes.Update()
}

func (o *ChurchEventType) IsUpdated() bool {
    return o == _churchEventTypes.Updated()
}

type churchEventTypes struct {
	values []*ChurchEventType
    literals []enum.Literal
}

var _churchEventTypes = &churchEventTypes{values: []*ChurchEventType{
    {name: "Create", ordinal: 0},
    {name: "Created", ordinal: 1},
    {name: "Delete", ordinal: 2},
    {name: "Deleted", ordinal: 3},
    {name: "Update", ordinal: 4},
    {name: "Updated", ordinal: 5}},
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

func (o *churchEventTypes) Create() *ChurchEventType {
    return _churchEventTypes.values[0]
}

func (o *churchEventTypes) Created() *ChurchEventType {
    return _churchEventTypes.values[1]
}

func (o *churchEventTypes) Delete() *ChurchEventType {
    return _churchEventTypes.values[2]
}

func (o *churchEventTypes) Deleted() *ChurchEventType {
    return _churchEventTypes.values[3]
}

func (o *churchEventTypes) Update() *ChurchEventType {
    return _churchEventTypes.values[4]
}

func (o *churchEventTypes) Updated() *ChurchEventType {
    return _churchEventTypes.values[5]
}

func (o *churchEventTypes) ParseChurchEventType(name string) (ret *ChurchEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
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

func (o GraduationEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *GraduationEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := GraduationEventTypes().ParseGraduationEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GraduationEventType %q", lit.Name)
        }
	}
	return
}

func (o GraduationEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *GraduationEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := GraduationEventTypes().ParseGraduationEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GraduationEventType %q", lit)
        }
    }
    return
}

func (o *GraduationEventType) IsCreate() bool {
    return o == _graduationEventTypes.Create()
}

func (o *GraduationEventType) IsCreated() bool {
    return o == _graduationEventTypes.Created()
}

func (o *GraduationEventType) IsDelete() bool {
    return o == _graduationEventTypes.Delete()
}

func (o *GraduationEventType) IsDeleted() bool {
    return o == _graduationEventTypes.Deleted()
}

func (o *GraduationEventType) IsUpdate() bool {
    return o == _graduationEventTypes.Update()
}

func (o *GraduationEventType) IsUpdated() bool {
    return o == _graduationEventTypes.Updated()
}

type graduationEventTypes struct {
	values []*GraduationEventType
    literals []enum.Literal
}

var _graduationEventTypes = &graduationEventTypes{values: []*GraduationEventType{
    {name: "Create", ordinal: 0},
    {name: "Created", ordinal: 1},
    {name: "Delete", ordinal: 2},
    {name: "Deleted", ordinal: 3},
    {name: "Update", ordinal: 4},
    {name: "Updated", ordinal: 5}},
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

func (o *graduationEventTypes) Create() *GraduationEventType {
    return _graduationEventTypes.values[0]
}

func (o *graduationEventTypes) Created() *GraduationEventType {
    return _graduationEventTypes.values[1]
}

func (o *graduationEventTypes) Delete() *GraduationEventType {
    return _graduationEventTypes.values[2]
}

func (o *graduationEventTypes) Deleted() *GraduationEventType {
    return _graduationEventTypes.values[3]
}

func (o *graduationEventTypes) Update() *GraduationEventType {
    return _graduationEventTypes.values[4]
}

func (o *graduationEventTypes) Updated() *GraduationEventType {
    return _graduationEventTypes.values[5]
}

func (o *graduationEventTypes) ParseGraduationEventType(name string) (ret *GraduationEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
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

func (o ProfileEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *ProfileEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := ProfileEventTypes().ParseProfileEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ProfileEventType %q", lit.Name)
        }
	}
	return
}

func (o ProfileEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *ProfileEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := ProfileEventTypes().ParseProfileEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ProfileEventType %q", lit)
        }
    }
    return
}

func (o *ProfileEventType) IsCreate() bool {
    return o == _profileEventTypes.Create()
}

func (o *ProfileEventType) IsCreated() bool {
    return o == _profileEventTypes.Created()
}

func (o *ProfileEventType) IsDelete() bool {
    return o == _profileEventTypes.Delete()
}

func (o *ProfileEventType) IsDeleted() bool {
    return o == _profileEventTypes.Deleted()
}

func (o *ProfileEventType) IsUpdate() bool {
    return o == _profileEventTypes.Update()
}

func (o *ProfileEventType) IsUpdated() bool {
    return o == _profileEventTypes.Updated()
}

type profileEventTypes struct {
	values []*ProfileEventType
    literals []enum.Literal
}

var _profileEventTypes = &profileEventTypes{values: []*ProfileEventType{
    {name: "Create", ordinal: 0},
    {name: "Created", ordinal: 1},
    {name: "Delete", ordinal: 2},
    {name: "Deleted", ordinal: 3},
    {name: "Update", ordinal: 4},
    {name: "Updated", ordinal: 5}},
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

func (o *profileEventTypes) Create() *ProfileEventType {
    return _profileEventTypes.values[0]
}

func (o *profileEventTypes) Created() *ProfileEventType {
    return _profileEventTypes.values[1]
}

func (o *profileEventTypes) Delete() *ProfileEventType {
    return _profileEventTypes.values[2]
}

func (o *profileEventTypes) Deleted() *ProfileEventType {
    return _profileEventTypes.values[3]
}

func (o *profileEventTypes) Update() *ProfileEventType {
    return _profileEventTypes.values[4]
}

func (o *profileEventTypes) Updated() *ProfileEventType {
    return _profileEventTypes.values[5]
}

func (o *profileEventTypes) ParseProfileEventType(name string) (ret *ProfileEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*ProfileEventType), ok
	}
	return
}



