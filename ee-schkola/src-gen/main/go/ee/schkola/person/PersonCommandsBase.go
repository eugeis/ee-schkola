package person

import (
    "github.com/eugeis/gee/enum"
    "time"
)

type CreateChurch struct {
    Name  string
    Address  *Address
    Pastor  *PersonName
    Contact  *Contact
}

func NewCreateChurch(name string, address *Address, pastor *PersonName, contact *Contact) (ret *CreateChurch, err error) {
    ret = &CreateChurch{
        Name : name,
        Address : address,
        Pastor : pastor,
        Contact : contact,
    }
    return
}



type DeleteChurch struct {
    Id  string
}

func NewDeleteChurch(id string) (ret *DeleteChurch, err error) {
    ret = &DeleteChurch{
        Id : id,
    }
    return
}



type UpdateChurch struct {
    Name  string
    Address  *Address
    Pastor  *PersonName
    Contact  *Contact
}

func NewUpdateChurch(name string, address *Address, pastor *PersonName, contact *Contact) (ret *UpdateChurch, err error) {
    ret = &UpdateChurch{
        Name : name,
        Address : address,
        Pastor : pastor,
        Contact : contact,
    }
    return
}



type CreateGraduation struct {
    Name  string
    Level  *GraduationLevel
}

func NewCreateGraduation(name string, level *GraduationLevel) (ret *CreateGraduation, err error) {
    ret = &CreateGraduation{
        Name : name,
        Level : level,
    }
    return
}



type DeleteGraduation struct {
    Id  string
}

func NewDeleteGraduation(id string) (ret *DeleteGraduation, err error) {
    ret = &DeleteGraduation{
        Id : id,
    }
    return
}



type UpdateGraduation struct {
    Name  string
    Level  *GraduationLevel
}

func NewUpdateGraduation(name string, level *GraduationLevel) (ret *UpdateGraduation, err error) {
    ret = &UpdateGraduation{
        Name : name,
        Level : level,
    }
    return
}



type CreateProfile struct {
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

func NewCreateProfile(gender *Gender, name *PersonName, birthName string, birthday *time.Time, address *Address, contact *Contact, 
                photoData []byte, photo string, family *Family, church *ChurchInfo, education *Education) (ret *CreateProfile, err error) {
    ret = &CreateProfile{
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



type DeleteProfile struct {
    Id  string
}

func NewDeleteProfile(id string) (ret *DeleteProfile, err error) {
    ret = &DeleteProfile{
        Id : id,
    }
    return
}



type UpdateProfile struct {
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

func NewUpdateProfile(gender *Gender, name *PersonName, birthName string, birthday *time.Time, address *Address, contact *Contact, 
                photoData []byte, photo string, family *Family, church *ChurchInfo, education *Education) (ret *UpdateProfile, err error) {
    ret = &UpdateProfile{
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

func (o *ChurchCommandType) IsChurchCreate() bool {
    return o == _churchCommandTypes.ChurchCreate()
}

func (o *ChurchCommandType) IsChurchDelete() bool {
    return o == _churchCommandTypes.ChurchDelete()
}

func (o *ChurchCommandType) IsChurchUpdate() bool {
    return o == _churchCommandTypes.ChurchUpdate()
}

type churchCommandTypes struct {
	values []*ChurchCommandType
    literals []enum.Literal
}

var _churchCommandTypes = &churchCommandTypes{values: []*ChurchCommandType{
    {name: "ChurchCreate", ordinal: 0},
    {name: "ChurchDelete", ordinal: 1},
    {name: "ChurchUpdate", ordinal: 2}},
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

func (o *churchCommandTypes) ChurchCreate() *ChurchCommandType {
    return _churchCommandTypes.values[0]
}

func (o *churchCommandTypes) ChurchDelete() *ChurchCommandType {
    return _churchCommandTypes.values[1]
}

func (o *churchCommandTypes) ChurchUpdate() *ChurchCommandType {
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

func (o *GraduationCommandType) IsGraduationCreate() bool {
    return o == _graduationCommandTypes.GraduationCreate()
}

func (o *GraduationCommandType) IsGraduationDelete() bool {
    return o == _graduationCommandTypes.GraduationDelete()
}

func (o *GraduationCommandType) IsGraduationUpdate() bool {
    return o == _graduationCommandTypes.GraduationUpdate()
}

type graduationCommandTypes struct {
	values []*GraduationCommandType
    literals []enum.Literal
}

var _graduationCommandTypes = &graduationCommandTypes{values: []*GraduationCommandType{
    {name: "GraduationCreate", ordinal: 0},
    {name: "GraduationDelete", ordinal: 1},
    {name: "GraduationUpdate", ordinal: 2}},
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

func (o *graduationCommandTypes) GraduationCreate() *GraduationCommandType {
    return _graduationCommandTypes.values[0]
}

func (o *graduationCommandTypes) GraduationDelete() *GraduationCommandType {
    return _graduationCommandTypes.values[1]
}

func (o *graduationCommandTypes) GraduationUpdate() *GraduationCommandType {
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

func (o *ProfileCommandType) IsProfileCreate() bool {
    return o == _profileCommandTypes.ProfileCreate()
}

func (o *ProfileCommandType) IsProfileDelete() bool {
    return o == _profileCommandTypes.ProfileDelete()
}

func (o *ProfileCommandType) IsProfileUpdate() bool {
    return o == _profileCommandTypes.ProfileUpdate()
}

type profileCommandTypes struct {
	values []*ProfileCommandType
    literals []enum.Literal
}

var _profileCommandTypes = &profileCommandTypes{values: []*ProfileCommandType{
    {name: "ProfileCreate", ordinal: 0},
    {name: "ProfileDelete", ordinal: 1},
    {name: "ProfileUpdate", ordinal: 2}},
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

func (o *profileCommandTypes) ProfileCreate() *ProfileCommandType {
    return _profileCommandTypes.values[0]
}

func (o *profileCommandTypes) ProfileDelete() *ProfileCommandType {
    return _profileCommandTypes.values[1]
}

func (o *profileCommandTypes) ProfileUpdate() *ProfileCommandType {
    return _profileCommandTypes.values[2]
}

func (o *profileCommandTypes) ParseProfileCommandType(name string) (ret *ProfileCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*ProfileCommandType), ok
	}
	return
}



