package person

import (
    "github.com/eugeis/gee/enum"
    "time"
)

type ChurchCreated struct {
    Name  string
    Address  *Address
    Pastor  *PersonName
    Contact  *Contact
}

func NewChurchCreated(name string, address *Address, pastor *PersonName, contact *Contact) (ret *ChurchCreated, err error) {
    ret = &ChurchCreated{
        Name : name,
        Address : address,
        Pastor : pastor,
        Contact : contact,
    }
    return
}



type ChurchDeleted struct {
    Id  string
}

func NewChurchDeleted(id string) (ret *ChurchDeleted, err error) {
    ret = &ChurchDeleted{
        Id : id,
    }
    return
}



type ChurchUpdated struct {
    Name  string
    Address  *Address
    Pastor  *PersonName
    Contact  *Contact
}

func NewChurchUpdated(name string, address *Address, pastor *PersonName, contact *Contact) (ret *ChurchUpdated, err error) {
    ret = &ChurchUpdated{
        Name : name,
        Address : address,
        Pastor : pastor,
        Contact : contact,
    }
    return
}



type GraduationCreated struct {
    Name  string
    Level  *GraduationLevel
}

func NewGraduationCreated(name string, level *GraduationLevel) (ret *GraduationCreated, err error) {
    ret = &GraduationCreated{
        Name : name,
        Level : level,
    }
    return
}



type GraduationDeleted struct {
    Id  string
}

func NewGraduationDeleted(id string) (ret *GraduationDeleted, err error) {
    ret = &GraduationDeleted{
        Id : id,
    }
    return
}



type GraduationUpdated struct {
    Name  string
    Level  *GraduationLevel
}

func NewGraduationUpdated(name string, level *GraduationLevel) (ret *GraduationUpdated, err error) {
    ret = &GraduationUpdated{
        Name : name,
        Level : level,
    }
    return
}



type ProfileCreated struct {
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

func NewProfileCreated(gender *Gender, name *PersonName, birthName string, birthday *time.Time, address *Address, contact *Contact, 
                photoData []byte, photo string, family *Family, church *ChurchInfo, education *Education) (ret *ProfileCreated, err error) {
    ret = &ProfileCreated{
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



type ProfileDeleted struct {
    Id  string
}

func NewProfileDeleted(id string) (ret *ProfileDeleted, err error) {
    ret = &ProfileDeleted{
        Id : id,
    }
    return
}



type ProfileUpdated struct {
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

func NewProfileUpdated(gender *Gender, name *PersonName, birthName string, birthday *time.Time, address *Address, contact *Contact, 
                photoData []byte, photo string, family *Family, church *ChurchInfo, education *Education) (ret *ProfileUpdated, err error) {
    ret = &ProfileUpdated{
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

func (o *ChurchEventType) IsCreatedChurch() bool {
    return o == _churchEventTypes.CreatedChurch()
}

func (o *ChurchEventType) IsDeletedChurch() bool {
    return o == _churchEventTypes.DeletedChurch()
}

func (o *ChurchEventType) IsUpdatedChurch() bool {
    return o == _churchEventTypes.UpdatedChurch()
}

type churchEventTypes struct {
	values []*ChurchEventType
    literals []enum.Literal
}

var _churchEventTypes = &churchEventTypes{values: []*ChurchEventType{
    {name: "createdChurch", ordinal: 0},
    {name: "deletedChurch", ordinal: 1},
    {name: "updatedChurch", ordinal: 2}},
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

func (o *churchEventTypes) CreatedChurch() *ChurchEventType {
    return _churchEventTypes.values[0]
}

func (o *churchEventTypes) DeletedChurch() *ChurchEventType {
    return _churchEventTypes.values[1]
}

func (o *churchEventTypes) UpdatedChurch() *ChurchEventType {
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

func (o *GraduationEventType) IsCreatedGraduation() bool {
    return o == _graduationEventTypes.CreatedGraduation()
}

func (o *GraduationEventType) IsDeletedGraduation() bool {
    return o == _graduationEventTypes.DeletedGraduation()
}

func (o *GraduationEventType) IsUpdatedGraduation() bool {
    return o == _graduationEventTypes.UpdatedGraduation()
}

type graduationEventTypes struct {
	values []*GraduationEventType
    literals []enum.Literal
}

var _graduationEventTypes = &graduationEventTypes{values: []*GraduationEventType{
    {name: "createdGraduation", ordinal: 0},
    {name: "deletedGraduation", ordinal: 1},
    {name: "updatedGraduation", ordinal: 2}},
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

func (o *graduationEventTypes) CreatedGraduation() *GraduationEventType {
    return _graduationEventTypes.values[0]
}

func (o *graduationEventTypes) DeletedGraduation() *GraduationEventType {
    return _graduationEventTypes.values[1]
}

func (o *graduationEventTypes) UpdatedGraduation() *GraduationEventType {
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

func (o *ProfileEventType) IsCreatedProfile() bool {
    return o == _profileEventTypes.CreatedProfile()
}

func (o *ProfileEventType) IsDeletedProfile() bool {
    return o == _profileEventTypes.DeletedProfile()
}

func (o *ProfileEventType) IsUpdatedProfile() bool {
    return o == _profileEventTypes.UpdatedProfile()
}

type profileEventTypes struct {
	values []*ProfileEventType
    literals []enum.Literal
}

var _profileEventTypes = &profileEventTypes{values: []*ProfileEventType{
    {name: "createdProfile", ordinal: 0},
    {name: "deletedProfile", ordinal: 1},
    {name: "updatedProfile", ordinal: 2}},
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

func (o *profileEventTypes) CreatedProfile() *ProfileEventType {
    return _profileEventTypes.values[0]
}

func (o *profileEventTypes) DeletedProfile() *ProfileEventType {
    return _profileEventTypes.values[1]
}

func (o *profileEventTypes) UpdatedProfile() *ProfileEventType {
    return _profileEventTypes.values[2]
}

func (o *profileEventTypes) ParseProfileEventType(name string) (ret *ProfileEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*ProfileEventType), ok
	}
	return
}



