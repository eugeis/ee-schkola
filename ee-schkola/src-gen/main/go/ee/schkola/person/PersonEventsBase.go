package person

import (
    "github.com/eugeis/gee/enum"
    "time"
)

type CreatedChurch struct {
    Name  string
    Address  *Address
    Pastor  *PersonName
    Contact  *Contact
}

func NewCreatedChurch(name string, address *Address, pastor *PersonName, contact *Contact) (ret *CreatedChurch, err error) {
    ret = &CreatedChurch{
        Name : name,
        Address : address,
        Pastor : pastor,
        Contact : contact,
    }
    return
}



type DeletedChurch struct {
    Id  string
}

func NewDeletedChurch(id string) (ret *DeletedChurch, err error) {
    ret = &DeletedChurch{
        Id : id,
    }
    return
}



type UpdatedChurch struct {
    Name  string
    Address  *Address
    Pastor  *PersonName
    Contact  *Contact
}

func NewUpdatedChurch(name string, address *Address, pastor *PersonName, contact *Contact) (ret *UpdatedChurch, err error) {
    ret = &UpdatedChurch{
        Name : name,
        Address : address,
        Pastor : pastor,
        Contact : contact,
    }
    return
}



type CreatedGraduation struct {
    Name  string
    Level  *GraduationLevel
}

func NewCreatedGraduation(name string, level *GraduationLevel) (ret *CreatedGraduation, err error) {
    ret = &CreatedGraduation{
        Name : name,
        Level : level,
    }
    return
}



type DeletedGraduation struct {
    Id  string
}

func NewDeletedGraduation(id string) (ret *DeletedGraduation, err error) {
    ret = &DeletedGraduation{
        Id : id,
    }
    return
}



type UpdatedGraduation struct {
    Name  string
    Level  *GraduationLevel
}

func NewUpdatedGraduation(name string, level *GraduationLevel) (ret *UpdatedGraduation, err error) {
    ret = &UpdatedGraduation{
        Name : name,
        Level : level,
    }
    return
}



type CreatedProfile struct {
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

func NewCreatedProfile(gender *Gender, name *PersonName, birthName string, birthday *time.Time, address *Address, contact *Contact, 
                photoData []byte, photo string, family *Family, church *ChurchInfo, education *Education) (ret *CreatedProfile, err error) {
    ret = &CreatedProfile{
        Gender : gender,
        Name : name,
        BirthName : birthName,
        Birthday : birthday,
        Address : address,
        Contact : contact,
        PhotoData : photoData,
        Photo : photo,
        Family : family,
        Church : church,
        Education : education,
    }
    return
}



type DeletedProfile struct {
    Id  string
}

func NewDeletedProfile(id string) (ret *DeletedProfile, err error) {
    ret = &DeletedProfile{
        Id : id,
    }
    return
}



type UpdatedProfile struct {
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

func NewUpdatedProfile(gender *Gender, name *PersonName, birthName string, birthday *time.Time, address *Address, contact *Contact, 
                photoData []byte, photo string, family *Family, church *ChurchInfo, education *Education) (ret *UpdatedProfile, err error) {
    ret = &UpdatedProfile{
        Gender : gender,
        Name : name,
        BirthName : birthName,
        Birthday : birthday,
        Address : address,
        Contact : contact,
        PhotoData : photoData,
        Photo : photo,
        Family : family,
        Church : church,
        Education : education,
    }
    return
}




type ChurchAggregateEventType struct {
	name  string
	ordinal int
}

func (o *ChurchAggregateEventType) Name() string {
    return o.name
}

func (o *ChurchAggregateEventType) Ordinal() int {
    return o.ordinal
}

func (o *ChurchAggregateEventType) IsChurchCreated() bool {
    return o == _churchAggregateEventTypes.ChurchCreated()
}

func (o *ChurchAggregateEventType) IsChurchDeleted() bool {
    return o == _churchAggregateEventTypes.ChurchDeleted()
}

func (o *ChurchAggregateEventType) IsChurchUpdated() bool {
    return o == _churchAggregateEventTypes.ChurchUpdated()
}

type churchAggregateEventTypes struct {
	values []*ChurchAggregateEventType
    literals []enum.Literal
}

var _churchAggregateEventTypes = &churchAggregateEventTypes{values: []*ChurchAggregateEventType{
    {name: "ChurchCreated", ordinal: 0},
    {name: "ChurchDeleted", ordinal: 1},
    {name: "ChurchUpdated", ordinal: 2}},
}

func ChurchAggregateEventTypes() *churchAggregateEventTypes {
	return _churchAggregateEventTypes
}

func (o *churchAggregateEventTypes) Values() []*ChurchAggregateEventType {
	return o.values
}

func (o *churchAggregateEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *churchAggregateEventTypes) ChurchCreated() *ChurchAggregateEventType {
    return _churchAggregateEventTypes.values[0]
}

func (o *churchAggregateEventTypes) ChurchDeleted() *ChurchAggregateEventType {
    return _churchAggregateEventTypes.values[1]
}

func (o *churchAggregateEventTypes) ChurchUpdated() *ChurchAggregateEventType {
    return _churchAggregateEventTypes.values[2]
}

func (o *churchAggregateEventTypes) ParseChurchAggregateEventType(name string) (ret *ChurchAggregateEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*ChurchAggregateEventType), ok
	}
	return
}


type GraduationAggregateEventType struct {
	name  string
	ordinal int
}

func (o *GraduationAggregateEventType) Name() string {
    return o.name
}

func (o *GraduationAggregateEventType) Ordinal() int {
    return o.ordinal
}

func (o *GraduationAggregateEventType) IsGraduationCreated() bool {
    return o == _graduationAggregateEventTypes.GraduationCreated()
}

func (o *GraduationAggregateEventType) IsGraduationDeleted() bool {
    return o == _graduationAggregateEventTypes.GraduationDeleted()
}

func (o *GraduationAggregateEventType) IsGraduationUpdated() bool {
    return o == _graduationAggregateEventTypes.GraduationUpdated()
}

type graduationAggregateEventTypes struct {
	values []*GraduationAggregateEventType
    literals []enum.Literal
}

var _graduationAggregateEventTypes = &graduationAggregateEventTypes{values: []*GraduationAggregateEventType{
    {name: "GraduationCreated", ordinal: 0},
    {name: "GraduationDeleted", ordinal: 1},
    {name: "GraduationUpdated", ordinal: 2}},
}

func GraduationAggregateEventTypes() *graduationAggregateEventTypes {
	return _graduationAggregateEventTypes
}

func (o *graduationAggregateEventTypes) Values() []*GraduationAggregateEventType {
	return o.values
}

func (o *graduationAggregateEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *graduationAggregateEventTypes) GraduationCreated() *GraduationAggregateEventType {
    return _graduationAggregateEventTypes.values[0]
}

func (o *graduationAggregateEventTypes) GraduationDeleted() *GraduationAggregateEventType {
    return _graduationAggregateEventTypes.values[1]
}

func (o *graduationAggregateEventTypes) GraduationUpdated() *GraduationAggregateEventType {
    return _graduationAggregateEventTypes.values[2]
}

func (o *graduationAggregateEventTypes) ParseGraduationAggregateEventType(name string) (ret *GraduationAggregateEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*GraduationAggregateEventType), ok
	}
	return
}


type ProfileAggregateEventType struct {
	name  string
	ordinal int
}

func (o *ProfileAggregateEventType) Name() string {
    return o.name
}

func (o *ProfileAggregateEventType) Ordinal() int {
    return o.ordinal
}

func (o *ProfileAggregateEventType) IsProfileCreated() bool {
    return o == _profileAggregateEventTypes.ProfileCreated()
}

func (o *ProfileAggregateEventType) IsProfileDeleted() bool {
    return o == _profileAggregateEventTypes.ProfileDeleted()
}

func (o *ProfileAggregateEventType) IsProfileUpdated() bool {
    return o == _profileAggregateEventTypes.ProfileUpdated()
}

type profileAggregateEventTypes struct {
	values []*ProfileAggregateEventType
    literals []enum.Literal
}

var _profileAggregateEventTypes = &profileAggregateEventTypes{values: []*ProfileAggregateEventType{
    {name: "ProfileCreated", ordinal: 0},
    {name: "ProfileDeleted", ordinal: 1},
    {name: "ProfileUpdated", ordinal: 2}},
}

func ProfileAggregateEventTypes() *profileAggregateEventTypes {
	return _profileAggregateEventTypes
}

func (o *profileAggregateEventTypes) Values() []*ProfileAggregateEventType {
	return o.values
}

func (o *profileAggregateEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *profileAggregateEventTypes) ProfileCreated() *ProfileAggregateEventType {
    return _profileAggregateEventTypes.values[0]
}

func (o *profileAggregateEventTypes) ProfileDeleted() *ProfileAggregateEventType {
    return _profileAggregateEventTypes.values[1]
}

func (o *profileAggregateEventTypes) ProfileUpdated() *ProfileAggregateEventType {
    return _profileAggregateEventTypes.values[2]
}

func (o *profileAggregateEventTypes) ParseProfileAggregateEventType(name string) (ret *ProfileAggregateEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*ProfileAggregateEventType), ok
	}
	return
}



