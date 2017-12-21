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
     CreateCommand eventhorizon.CommandType = "Create"
     DeleteCommand eventhorizon.CommandType = "Delete"
     UpdateCommand eventhorizon.CommandType = "Update"
)


const (
     CreateCommand eventhorizon.CommandType = "Create"
     DeleteCommand eventhorizon.CommandType = "Delete"
     UpdateCommand eventhorizon.CommandType = "Update"
)


const (
     CreateCommand eventhorizon.CommandType = "Create"
     DeleteCommand eventhorizon.CommandType = "Delete"
     UpdateCommand eventhorizon.CommandType = "Update"
)




        
type Create struct {
    Name string `json:"name" eh:"optional"`
    Address *Address `json:"address" eh:"optional"`
    Pastor *shared.PersonName `json:"pastor" eh:"optional"`
    Contact *Contact `json:"contact" eh:"optional"`
    Association string `json:"association" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Create) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Create) AggregateType() eventhorizon.AggregateType  { return ChurchAggregateType }
func (o *Create) CommandType() eventhorizon.CommandType      { return CreateCommand }



        
type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Delete) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Delete) AggregateType() eventhorizon.AggregateType  { return ChurchAggregateType }
func (o *Delete) CommandType() eventhorizon.CommandType      { return DeleteCommand }



        
type Update struct {
    Name string `json:"name" eh:"optional"`
    Address *Address `json:"address" eh:"optional"`
    Pastor *shared.PersonName `json:"pastor" eh:"optional"`
    Contact *Contact `json:"contact" eh:"optional"`
    Association string `json:"association" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Update) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Update) AggregateType() eventhorizon.AggregateType  { return ChurchAggregateType }
func (o *Update) CommandType() eventhorizon.CommandType      { return UpdateCommand }



        
type Create struct {
    Name string `json:"name" eh:"optional"`
    Level *GraduationLevel `json:"level" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Create) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Create) AggregateType() eventhorizon.AggregateType  { return GraduationAggregateType }
func (o *Create) CommandType() eventhorizon.CommandType      { return CreateCommand }



        
type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Delete) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Delete) AggregateType() eventhorizon.AggregateType  { return GraduationAggregateType }
func (o *Delete) CommandType() eventhorizon.CommandType      { return DeleteCommand }



        
type Update struct {
    Name string `json:"name" eh:"optional"`
    Level *GraduationLevel `json:"level" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Update) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Update) AggregateType() eventhorizon.AggregateType  { return GraduationAggregateType }
func (o *Update) CommandType() eventhorizon.CommandType      { return UpdateCommand }



        
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
func (o *Create) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Create) AggregateType() eventhorizon.AggregateType  { return ProfileAggregateType }
func (o *Create) CommandType() eventhorizon.CommandType      { return CreateCommand }



        
type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Delete) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Delete) AggregateType() eventhorizon.AggregateType  { return ProfileAggregateType }
func (o *Delete) CommandType() eventhorizon.CommandType      { return DeleteCommand }



        
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
func (o *Update) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Update) AggregateType() eventhorizon.AggregateType  { return ProfileAggregateType }
func (o *Update) CommandType() eventhorizon.CommandType      { return UpdateCommand }





type ChurchCommandType struct {
	name  string
	ordinal int
}

func (o *ChurchCommandType) Name() string {
    return o.name
}

func (o *ChurchCommandType) Ordinal() int {
    return o.ordinal
}

func (o ChurchCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *ChurchCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := ChurchCommandTypes().ParseChurchCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ChurchCommandType %q", lit.Name)
        }
	}
	return
}

func (o ChurchCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *ChurchCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := ChurchCommandTypes().ParseChurchCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ChurchCommandType %q", lit)
        }
    }
    return
}

func (o *ChurchCommandType) IsCreate() bool {
    return o == _churchCommandTypes.Create()
}

func (o *ChurchCommandType) IsDelete() bool {
    return o == _churchCommandTypes.Delete()
}

func (o *ChurchCommandType) IsUpdate() bool {
    return o == _churchCommandTypes.Update()
}

type churchCommandTypes struct {
	values []*ChurchCommandType
    literals []enum.Literal
}

var _churchCommandTypes = &churchCommandTypes{values: []*ChurchCommandType{
    {name: "Create", ordinal: 0},
    {name: "Delete", ordinal: 1},
    {name: "Update", ordinal: 2}},
}

func ChurchCommandTypes() *churchCommandTypes {
	return _churchCommandTypes
}

func (o *churchCommandTypes) Values() []*ChurchCommandType {
	return o.values
}

func (o *churchCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *churchCommandTypes) Create() *ChurchCommandType {
    return _churchCommandTypes.values[0]
}

func (o *churchCommandTypes) Delete() *ChurchCommandType {
    return _churchCommandTypes.values[1]
}

func (o *churchCommandTypes) Update() *ChurchCommandType {
    return _churchCommandTypes.values[2]
}

func (o *churchCommandTypes) ParseChurchCommandType(name string) (ret *ChurchCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*ChurchCommandType), ok
	}
	return
}


type GraduationCommandType struct {
	name  string
	ordinal int
}

func (o *GraduationCommandType) Name() string {
    return o.name
}

func (o *GraduationCommandType) Ordinal() int {
    return o.ordinal
}

func (o GraduationCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *GraduationCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := GraduationCommandTypes().ParseGraduationCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GraduationCommandType %q", lit.Name)
        }
	}
	return
}

func (o GraduationCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *GraduationCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := GraduationCommandTypes().ParseGraduationCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GraduationCommandType %q", lit)
        }
    }
    return
}

func (o *GraduationCommandType) IsCreate() bool {
    return o == _graduationCommandTypes.Create()
}

func (o *GraduationCommandType) IsDelete() bool {
    return o == _graduationCommandTypes.Delete()
}

func (o *GraduationCommandType) IsUpdate() bool {
    return o == _graduationCommandTypes.Update()
}

type graduationCommandTypes struct {
	values []*GraduationCommandType
    literals []enum.Literal
}

var _graduationCommandTypes = &graduationCommandTypes{values: []*GraduationCommandType{
    {name: "Create", ordinal: 0},
    {name: "Delete", ordinal: 1},
    {name: "Update", ordinal: 2}},
}

func GraduationCommandTypes() *graduationCommandTypes {
	return _graduationCommandTypes
}

func (o *graduationCommandTypes) Values() []*GraduationCommandType {
	return o.values
}

func (o *graduationCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *graduationCommandTypes) Create() *GraduationCommandType {
    return _graduationCommandTypes.values[0]
}

func (o *graduationCommandTypes) Delete() *GraduationCommandType {
    return _graduationCommandTypes.values[1]
}

func (o *graduationCommandTypes) Update() *GraduationCommandType {
    return _graduationCommandTypes.values[2]
}

func (o *graduationCommandTypes) ParseGraduationCommandType(name string) (ret *GraduationCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*GraduationCommandType), ok
	}
	return
}


type ProfileCommandType struct {
	name  string
	ordinal int
}

func (o *ProfileCommandType) Name() string {
    return o.name
}

func (o *ProfileCommandType) Ordinal() int {
    return o.ordinal
}

func (o ProfileCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *ProfileCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := ProfileCommandTypes().ParseProfileCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ProfileCommandType %q", lit.Name)
        }
	}
	return
}

func (o ProfileCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *ProfileCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := ProfileCommandTypes().ParseProfileCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ProfileCommandType %q", lit)
        }
    }
    return
}

func (o *ProfileCommandType) IsCreate() bool {
    return o == _profileCommandTypes.Create()
}

func (o *ProfileCommandType) IsDelete() bool {
    return o == _profileCommandTypes.Delete()
}

func (o *ProfileCommandType) IsUpdate() bool {
    return o == _profileCommandTypes.Update()
}

type profileCommandTypes struct {
	values []*ProfileCommandType
    literals []enum.Literal
}

var _profileCommandTypes = &profileCommandTypes{values: []*ProfileCommandType{
    {name: "Create", ordinal: 0},
    {name: "Delete", ordinal: 1},
    {name: "Update", ordinal: 2}},
}

func ProfileCommandTypes() *profileCommandTypes {
	return _profileCommandTypes
}

func (o *profileCommandTypes) Values() []*ProfileCommandType {
	return o.values
}

func (o *profileCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *profileCommandTypes) Create() *ProfileCommandType {
    return _profileCommandTypes.values[0]
}

func (o *profileCommandTypes) Delete() *ProfileCommandType {
    return _profileCommandTypes.values[1]
}

func (o *profileCommandTypes) Update() *ProfileCommandType {
    return _profileCommandTypes.values[2]
}

func (o *profileCommandTypes) ParseProfileCommandType(name string) (ret *ProfileCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*ProfileCommandType), ok
	}
	return
}



