package person

import (
    "ee/schkola"
    "time"
)
type ChurchCreated struct {
    Name  string
    Address  *Address
    Pastor  *PersonName
    Contact  *Contact
    *schkola.SchkolaBase
}


type ChurchDeleted struct {
    Id  string
}


type ChurchUpdated struct {
    Name  string
    Address  *Address
    Pastor  *PersonName
    Contact  *Contact
    *schkola.SchkolaBase
}


type GraduationCreated struct {
    Name  string
    Level  *GraduationLevel
    *schkola.SchkolaBase
}


type GraduationDeleted struct {
    Id  string
}


type GraduationUpdated struct {
    Name  string
    Level  *GraduationLevel
    *schkola.SchkolaBase
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
    *schkola.SchkolaBase
}


type ProfileDeleted struct {
    Id  string
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
    *schkola.SchkolaBase
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
    switch name {
      case "CreatedChurch":
        ret = o.CreatedChurch()
      case "DeletedChurch":
        ret = o.DeletedChurch()
      case "UpdatedChurch":
        ret = o.UpdatedChurch()
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
    switch name {
      case "CreatedGraduation":
        ret = o.CreatedGraduation()
      case "DeletedGraduation":
        ret = o.DeletedGraduation()
      case "UpdatedGraduation":
        ret = o.UpdatedGraduation()
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
    switch name {
      case "CreatedProfile":
        ret = o.CreatedProfile()
      case "DeletedProfile":
        ret = o.DeletedProfile()
      case "UpdatedProfile":
        ret = o.UpdatedProfile()
    }
    return
}



