package person

import (
    "github.com/eugeis/gee/enum"
    "time"
)

type ChurchCreate struct {
    name string
    address *Address
    pastor *PersonName
    contact *Contact
}

func NewChurchCreate(name string, address *Address, pastor *PersonName, contact *Contact) (ret *ChurchCreate, err error) {
    ret = &ChurchCreate{
        name: name,
        address: address,
        pastor: pastor,
        contact: contact,
    }
    
    return
}



type ChurchDelete struct {
    id string
}

func NewChurchDelete(id string) (ret *ChurchDelete, err error) {
    ret = &ChurchDelete{
        id: id,
    }
    
    return
}



type ChurchUpdate struct {
    name string
    address *Address
    pastor *PersonName
    contact *Contact
}

func NewChurchUpdate(name string, address *Address, pastor *PersonName, contact *Contact) (ret *ChurchUpdate, err error) {
    ret = &ChurchUpdate{
        name: name,
        address: address,
        pastor: pastor,
        contact: contact,
    }
    
    return
}



type GraduationCreate struct {
    name string
    level *GraduationLevel
}

func NewGraduationCreate(name string, level *GraduationLevel) (ret *GraduationCreate, err error) {
    ret = &GraduationCreate{
        name: name,
        level: level,
    }
    
    return
}



type GraduationDelete struct {
    id string
}

func NewGraduationDelete(id string) (ret *GraduationDelete, err error) {
    ret = &GraduationDelete{
        id: id,
    }
    
    return
}



type GraduationUpdate struct {
    name string
    level *GraduationLevel
}

func NewGraduationUpdate(name string, level *GraduationLevel) (ret *GraduationUpdate, err error) {
    ret = &GraduationUpdate{
        name: name,
        level: level,
    }
    
    return
}



type ProfileCreate struct {
    gender *Gender
    name *PersonName
    birthName string
    birthday *time.Time
    address *Address
    contact *Contact
    photoData []byte
    photo string
    family *Family
    church *ChurchInfo
    education *Education
}

func NewProfileCreate(gender *Gender, name *PersonName, birthName string, birthday *time.Time, address *Address, contact *Contact, 
                photoData []byte, photo string, family *Family, church *ChurchInfo, education *Education) (ret *ProfileCreate, err error) {
    ret = &ProfileCreate{
        gender: gender,
        name: name,
        birthName: birthName,
        birthday: birthday,
        address: address,
        contact: contact,
        photoData: photoData,
        photo: photo,
        family: family,
        church: church,
        education: education,
    }
    
    return
}



type ProfileDelete struct {
    id string
}

func NewProfileDelete(id string) (ret *ProfileDelete, err error) {
    ret = &ProfileDelete{
        id: id,
    }
    
    return
}



type ProfileUpdate struct {
    gender *Gender
    name *PersonName
    birthName string
    birthday *time.Time
    address *Address
    contact *Contact
    photoData []byte
    photo string
    family *Family
    church *ChurchInfo
    education *Education
}

func NewProfileUpdate(gender *Gender, name *PersonName, birthName string, birthday *time.Time, address *Address, contact *Contact, 
                photoData []byte, photo string, family *Family, church *ChurchInfo, education *Education) (ret *ProfileUpdate, err error) {
    ret = &ProfileUpdate{
        gender: gender,
        name: name,
        birthName: birthName,
        birthday: birthday,
        address: address,
        contact: contact,
        photoData: photoData,
        photo: photo,
        family: family,
        church: church,
        education: education,
    }
    
    return
}




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

func (o *ChurchCommandType) IsCreateChurch() bool {
    return o == _churchCommandTypes.CreateChurch()
}

func (o *ChurchCommandType) IsDeleteChurch() bool {
    return o == _churchCommandTypes.DeleteChurch()
}

func (o *ChurchCommandType) IsUpdateChurch() bool {
    return o == _churchCommandTypes.UpdateChurch()
}

type churchCommandTypes struct {
	values []*ChurchCommandType
    literals []enum.Literal
}

var _churchCommandTypes = &churchCommandTypes{values: []*ChurchCommandType{
    {name: "createChurch", ordinal: 0},
    {name: "deleteChurch", ordinal: 1},
    {name: "updateChurch", ordinal: 2}},
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

func (o *churchCommandTypes) CreateChurch() *ChurchCommandType {
    return _churchCommandTypes.values[0]
}

func (o *churchCommandTypes) DeleteChurch() *ChurchCommandType {
    return _churchCommandTypes.values[1]
}

func (o *churchCommandTypes) UpdateChurch() *ChurchCommandType {
    return _churchCommandTypes.values[2]
}

func (o *churchCommandTypes) ParseChurchCommandType(name string) (ret *ChurchCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
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

func (o *GraduationCommandType) IsCreateGraduation() bool {
    return o == _graduationCommandTypes.CreateGraduation()
}

func (o *GraduationCommandType) IsDeleteGraduation() bool {
    return o == _graduationCommandTypes.DeleteGraduation()
}

func (o *GraduationCommandType) IsUpdateGraduation() bool {
    return o == _graduationCommandTypes.UpdateGraduation()
}

type graduationCommandTypes struct {
	values []*GraduationCommandType
    literals []enum.Literal
}

var _graduationCommandTypes = &graduationCommandTypes{values: []*GraduationCommandType{
    {name: "createGraduation", ordinal: 0},
    {name: "deleteGraduation", ordinal: 1},
    {name: "updateGraduation", ordinal: 2}},
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

func (o *graduationCommandTypes) CreateGraduation() *GraduationCommandType {
    return _graduationCommandTypes.values[0]
}

func (o *graduationCommandTypes) DeleteGraduation() *GraduationCommandType {
    return _graduationCommandTypes.values[1]
}

func (o *graduationCommandTypes) UpdateGraduation() *GraduationCommandType {
    return _graduationCommandTypes.values[2]
}

func (o *graduationCommandTypes) ParseGraduationCommandType(name string) (ret *GraduationCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
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

func (o *ProfileCommandType) IsCreateProfile() bool {
    return o == _profileCommandTypes.CreateProfile()
}

func (o *ProfileCommandType) IsDeleteProfile() bool {
    return o == _profileCommandTypes.DeleteProfile()
}

func (o *ProfileCommandType) IsUpdateProfile() bool {
    return o == _profileCommandTypes.UpdateProfile()
}

type profileCommandTypes struct {
	values []*ProfileCommandType
    literals []enum.Literal
}

var _profileCommandTypes = &profileCommandTypes{values: []*ProfileCommandType{
    {name: "createProfile", ordinal: 0},
    {name: "deleteProfile", ordinal: 1},
    {name: "updateProfile", ordinal: 2}},
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

func (o *profileCommandTypes) CreateProfile() *ProfileCommandType {
    return _profileCommandTypes.values[0]
}

func (o *profileCommandTypes) DeleteProfile() *ProfileCommandType {
    return _profileCommandTypes.values[1]
}

func (o *profileCommandTypes) UpdateProfile() *ProfileCommandType {
    return _profileCommandTypes.values[2]
}

func (o *profileCommandTypes) ParseProfileCommandType(name string) (ret *ProfileCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*ProfileCommandType), ok
	}
	return
}



