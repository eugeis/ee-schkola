package person

import (
    "github.com/eugeis/gee/enum"
    "time"
)

type ChurchCreate struct {
    Name  string
    Address  *Address
    Pastor  *PersonName
    Contact  *Contact
}

func NewChurchCreate(name string, address *Address, pastor *PersonName, contact *Contact) (ret *ChurchCreate, err error) {
    ret = &ChurchCreate{
        Name : name,
        Address : address,
        Pastor : pastor,
        Contact : contact,
    }
    return
}



type ChurchDelete struct {
    Id  string
}

func NewChurchDelete(id string) (ret *ChurchDelete, err error) {
    ret = &ChurchDelete{
        Id : id,
    }
    return
}



type ChurchUpdate struct {
    Name  string
    Address  *Address
    Pastor  *PersonName
    Contact  *Contact
}

func NewChurchUpdate(name string, address *Address, pastor *PersonName, contact *Contact) (ret *ChurchUpdate, err error) {
    ret = &ChurchUpdate{
        Name : name,
        Address : address,
        Pastor : pastor,
        Contact : contact,
    }
    return
}



type GraduationCreate struct {
    Name  string
    Level  *GraduationLevel
}

func NewGraduationCreate(name string, level *GraduationLevel) (ret *GraduationCreate, err error) {
    ret = &GraduationCreate{
        Name : name,
        Level : level,
    }
    return
}



type GraduationDelete struct {
    Id  string
}

func NewGraduationDelete(id string) (ret *GraduationDelete, err error) {
    ret = &GraduationDelete{
        Id : id,
    }
    return
}



type GraduationUpdate struct {
    Name  string
    Level  *GraduationLevel
}

func NewGraduationUpdate(name string, level *GraduationLevel) (ret *GraduationUpdate, err error) {
    ret = &GraduationUpdate{
        Name : name,
        Level : level,
    }
    return
}



type ProfileCreate struct {
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

func NewProfileCreate(gender *Gender, name *PersonName, birthName string, birthday *time.Time, address *Address, contact *Contact, 
                photoData []byte, photo string, family *Family, church *ChurchInfo, education *Education) (ret *ProfileCreate, err error) {
    ret = &ProfileCreate{
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



type ProfileDelete struct {
    Id  string
}

func NewProfileDelete(id string) (ret *ProfileDelete, err error) {
    ret = &ProfileDelete{
        Id : id,
    }
    return
}



type ProfileUpdate struct {
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

func NewProfileUpdate(gender *Gender, name *PersonName, birthName string, birthday *time.Time, address *Address, contact *Contact, 
                photoData []byte, photo string, family *Family, church *ChurchInfo, education *Education) (ret *ProfileUpdate, err error) {
    ret = &ProfileUpdate{
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




type ChurchAggregateCommandType struct {
	name  string
	ordinal int
}

func (o *ChurchAggregateCommandType) Name() string {
    return o.name
}

func (o *ChurchAggregateCommandType) Ordinal() int {
    return o.ordinal
}

func (o *ChurchAggregateCommandType) IsCreateChurch() bool {
    return o == _churchAggregateCommandTypes.CreateChurch()
}

func (o *ChurchAggregateCommandType) IsDeleteChurch() bool {
    return o == _churchAggregateCommandTypes.DeleteChurch()
}

func (o *ChurchAggregateCommandType) IsUpdateChurch() bool {
    return o == _churchAggregateCommandTypes.UpdateChurch()
}

type churchAggregateCommandTypes struct {
	values []*ChurchAggregateCommandType
    literals []enum.Literal
}

var _churchAggregateCommandTypes = &churchAggregateCommandTypes{values: []*ChurchAggregateCommandType{
    {name: "createChurch", ordinal: 0},
    {name: "deleteChurch", ordinal: 1},
    {name: "updateChurch", ordinal: 2}},
}

func ChurchAggregateCommandTypes() *churchAggregateCommandTypes {
	return _churchAggregateCommandTypes
}

func (o *churchAggregateCommandTypes) Values() []*ChurchAggregateCommandType {
	return o.values
}

func (o *churchAggregateCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *churchAggregateCommandTypes) CreateChurch() *ChurchAggregateCommandType {
    return _churchAggregateCommandTypes.values[0]
}

func (o *churchAggregateCommandTypes) DeleteChurch() *ChurchAggregateCommandType {
    return _churchAggregateCommandTypes.values[1]
}

func (o *churchAggregateCommandTypes) UpdateChurch() *ChurchAggregateCommandType {
    return _churchAggregateCommandTypes.values[2]
}

func (o *churchAggregateCommandTypes) ParseChurchAggregateCommandType(name string) (ret *ChurchAggregateCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*ChurchAggregateCommandType), ok
	}
	return
}


type GraduationAggregateCommandType struct {
	name  string
	ordinal int
}

func (o *GraduationAggregateCommandType) Name() string {
    return o.name
}

func (o *GraduationAggregateCommandType) Ordinal() int {
    return o.ordinal
}

func (o *GraduationAggregateCommandType) IsCreateGraduation() bool {
    return o == _graduationAggregateCommandTypes.CreateGraduation()
}

func (o *GraduationAggregateCommandType) IsDeleteGraduation() bool {
    return o == _graduationAggregateCommandTypes.DeleteGraduation()
}

func (o *GraduationAggregateCommandType) IsUpdateGraduation() bool {
    return o == _graduationAggregateCommandTypes.UpdateGraduation()
}

type graduationAggregateCommandTypes struct {
	values []*GraduationAggregateCommandType
    literals []enum.Literal
}

var _graduationAggregateCommandTypes = &graduationAggregateCommandTypes{values: []*GraduationAggregateCommandType{
    {name: "createGraduation", ordinal: 0},
    {name: "deleteGraduation", ordinal: 1},
    {name: "updateGraduation", ordinal: 2}},
}

func GraduationAggregateCommandTypes() *graduationAggregateCommandTypes {
	return _graduationAggregateCommandTypes
}

func (o *graduationAggregateCommandTypes) Values() []*GraduationAggregateCommandType {
	return o.values
}

func (o *graduationAggregateCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *graduationAggregateCommandTypes) CreateGraduation() *GraduationAggregateCommandType {
    return _graduationAggregateCommandTypes.values[0]
}

func (o *graduationAggregateCommandTypes) DeleteGraduation() *GraduationAggregateCommandType {
    return _graduationAggregateCommandTypes.values[1]
}

func (o *graduationAggregateCommandTypes) UpdateGraduation() *GraduationAggregateCommandType {
    return _graduationAggregateCommandTypes.values[2]
}

func (o *graduationAggregateCommandTypes) ParseGraduationAggregateCommandType(name string) (ret *GraduationAggregateCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*GraduationAggregateCommandType), ok
	}
	return
}


type ProfileAggregateCommandType struct {
	name  string
	ordinal int
}

func (o *ProfileAggregateCommandType) Name() string {
    return o.name
}

func (o *ProfileAggregateCommandType) Ordinal() int {
    return o.ordinal
}

func (o *ProfileAggregateCommandType) IsCreateProfile() bool {
    return o == _profileAggregateCommandTypes.CreateProfile()
}

func (o *ProfileAggregateCommandType) IsDeleteProfile() bool {
    return o == _profileAggregateCommandTypes.DeleteProfile()
}

func (o *ProfileAggregateCommandType) IsUpdateProfile() bool {
    return o == _profileAggregateCommandTypes.UpdateProfile()
}

type profileAggregateCommandTypes struct {
	values []*ProfileAggregateCommandType
    literals []enum.Literal
}

var _profileAggregateCommandTypes = &profileAggregateCommandTypes{values: []*ProfileAggregateCommandType{
    {name: "createProfile", ordinal: 0},
    {name: "deleteProfile", ordinal: 1},
    {name: "updateProfile", ordinal: 2}},
}

func ProfileAggregateCommandTypes() *profileAggregateCommandTypes {
	return _profileAggregateCommandTypes
}

func (o *profileAggregateCommandTypes) Values() []*ProfileAggregateCommandType {
	return o.values
}

func (o *profileAggregateCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *profileAggregateCommandTypes) CreateProfile() *ProfileAggregateCommandType {
    return _profileAggregateCommandTypes.values[0]
}

func (o *profileAggregateCommandTypes) DeleteProfile() *ProfileAggregateCommandType {
    return _profileAggregateCommandTypes.values[1]
}

func (o *profileAggregateCommandTypes) UpdateProfile() *ProfileAggregateCommandType {
    return _profileAggregateCommandTypes.values[2]
}

func (o *profileAggregateCommandTypes) ParseProfileAggregateCommandType(name string) (ret *ProfileAggregateCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*ProfileAggregateCommandType), ok
	}
	return
}



