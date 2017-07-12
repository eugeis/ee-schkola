package person

import (
    "ee/schkola"
    "time"
)
type CreateChurch struct {
    Name  string
    Address  *Address
    Pastor  *PersonName
    Contact  *Contact
    *schkola.SchkolaBase
}


type DeleteChurch struct {
    Id  string
}


type UpdateChurch struct {
    Name  string
    Address  *Address
    Pastor  *PersonName
    Contact  *Contact
    *schkola.SchkolaBase
}


type CreateGraduation struct {
    Name  string
    Level  *GraduationLevel
    *schkola.SchkolaBase
}


type DeleteGraduation struct {
    Id  string
}


type UpdateGraduation struct {
    Name  string
    Level  *GraduationLevel
    *schkola.SchkolaBase
}


type CreateProfile struct {
    Gender  *Gender
    Name  *PersonName
    BirthName  string
    Birthday  *time.Time
    Address  *Address
    Contact  *Contact
    Photo  []byte
    Family  *Family
    Church  *ChurchInfo
    Education  *Education
    *schkola.SchkolaBase
}


type DeleteProfile struct {
    Id  string
}


type UpdateProfile struct {
    Gender  *Gender
    Name  *PersonName
    BirthName  string
    Birthday  *time.Time
    Address  *Address
    Contact  *Contact
    Photo  []byte
    Family  *Family
    Church  *ChurchInfo
    Education  *Education
    *schkola.SchkolaBase
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
    switch name {
      case "ChurchCreate":
        ret = o.ChurchCreate()
      case "ChurchDelete":
        ret = o.ChurchDelete()
      case "ChurchUpdate":
        ret = o.ChurchUpdate()
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
    switch name {
      case "GraduationCreate":
        ret = o.GraduationCreate()
      case "GraduationDelete":
        ret = o.GraduationDelete()
      case "GraduationUpdate":
        ret = o.GraduationUpdate()
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
    switch name {
      case "ProfileCreate":
        ret = o.ProfileCreate()
      case "ProfileDelete":
        ret = o.ProfileDelete()
      case "ProfileUpdate":
        ret = o.ProfileUpdate()
    }
    return
}



