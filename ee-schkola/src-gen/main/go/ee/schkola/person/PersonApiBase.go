package person

import (
    "ee/schkola"
    "github.com/eugeis/gee/enum"
    "time"
)

type Church struct {
    name string
    address *Address
    pastor *PersonName
    contact *Contact
    *schkola.SchkolaBase
}

func NewChurch(name string, address *Address, pastor *PersonName, contact *Contact, SchkolaBase *schkola.SchkolaBase) (ret *Church, err error) {
    ret = &Church{
        name: name,
        address: address,
        pastor: pastor,
        contact: contact,
        SchkolaBase: SchkolaBase,
    }
    
    return
}



type Graduation struct {
    name string
    level *GraduationLevel
    *schkola.SchkolaBase
}

func NewGraduation(name string, level *GraduationLevel, SchkolaBase *schkola.SchkolaBase) (ret *Graduation, err error) {
    ret = &Graduation{
        name: name,
        level: level,
        SchkolaBase: SchkolaBase,
    }
    
    return
}



type Profile struct {
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
    *schkola.SchkolaBase
}

func NewProfile(gender *Gender, name *PersonName, birthName string, birthday *time.Time, address *Address, contact *Contact, 
                photoData []byte, photo string, family *Family, church *ChurchInfo, education *Education, 
                SchkolaBase *schkola.SchkolaBase) (ret *Profile, err error) {
    ret = &Profile{
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
        SchkolaBase: SchkolaBase,
    }
    
    return
}







type Address struct {
    street string
    suite string
    city string
    code string
    country string
}

func NewAddress(street string, suite string, city string, code string, country string) (ret *Address, err error) {
    ret = &Address{
        street: street,
        suite: suite,
        city: city,
        code: code,
        country: country,
    }
    
    return
}



type ChurchInfo struct {
    church string
    association string
    member bool
    services string
}

func NewChurchInfo(church string, association string, member bool, services string) (ret *ChurchInfo, err error) {
    ret = &ChurchInfo{
        church: church,
        association: association,
        member: member,
        services: services,
    }
    
    return
}



type Contact struct {
    phone string
    email string
    cellphone string
}

func NewContact(phone string, email string, cellphone string) (ret *Contact, err error) {
    ret = &Contact{
        phone: phone,
        email: email,
        cellphone: cellphone,
    }
    
    return
}



type Education struct {
    graduation *Graduation
    profession string
}

func NewEducation(graduation *Graduation, profession string) (ret *Education, err error) {
    ret = &Education{
        graduation: graduation,
        profession: profession,
    }
    
    return
}



type Family struct {
    maritalState *MaritalState
    childrenCount int
    partner *PersonName
}

func NewFamily(maritalState *MaritalState, childrenCount int, partner *PersonName) (ret *Family, err error) {
    ret = &Family{
        maritalState: maritalState,
        childrenCount: childrenCount,
        partner: partner,
    }
    
    return
}



type PersonName struct {
    first string
    last string
}

func NewPersonName(first string, last string) (ret *PersonName, err error) {
    ret = &PersonName{
        first: first,
        last: last,
    }
    
    return
}




type Gender struct {
	name  string
	ordinal int
}

func (o *Gender) Name() string {
    return o.name
}

func (o *Gender) Ordinal() int {
    return o.ordinal
}

func (o *Gender) IsUnknown() bool {
    return o == _genders.Unknown()
}

func (o *Gender) IsMale() bool {
    return o == _genders.Male()
}

func (o *Gender) IsFemale() bool {
    return o == _genders.Female()
}

type genders struct {
	values []*Gender
    literals []enum.Literal
}

var _genders = &genders{values: []*Gender{
    {name: "Unknown", ordinal: 0},
    {name: "Male", ordinal: 1},
    {name: "Female", ordinal: 2}},
}

func Genders() *genders {
	return _genders
}

func (o *genders) Values() []*Gender {
	return o.values
}

func (o *genders) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *genders) Unknown() *Gender {
    return _genders.values[0]
}

func (o *genders) Male() *Gender {
    return _genders.values[1]
}

func (o *genders) Female() *Gender {
    return _genders.values[2]
}

func (o *genders) ParseGender(name string) (ret *Gender, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*Gender), ok
	}
	return
}


type GraduationLevel struct {
	name  string
	ordinal int
}

func (o *GraduationLevel) Name() string {
    return o.name
}

func (o *GraduationLevel) Ordinal() int {
    return o.ordinal
}

func (o *GraduationLevel) IsUnknown() bool {
    return o == _graduationLevels.Unknown()
}

func (o *GraduationLevel) IsMiddleSchool() bool {
    return o == _graduationLevels.MiddleSchool()
}

func (o *GraduationLevel) IsSecondarySchool() bool {
    return o == _graduationLevels.SecondarySchool()
}

func (o *GraduationLevel) IsHighSchool() bool {
    return o == _graduationLevels.HighSchool()
}

func (o *GraduationLevel) IsTechnicalCollege() bool {
    return o == _graduationLevels.TechnicalCollege()
}

func (o *GraduationLevel) IsCollege() bool {
    return o == _graduationLevels.College()
}

type graduationLevels struct {
	values []*GraduationLevel
    literals []enum.Literal
}

var _graduationLevels = &graduationLevels{values: []*GraduationLevel{
    {name: "Unknown", ordinal: 0},
    {name: "MiddleSchool", ordinal: 1},
    {name: "SecondarySchool", ordinal: 2},
    {name: "HighSchool", ordinal: 3},
    {name: "TechnicalCollege", ordinal: 4},
    {name: "College", ordinal: 5}},
}

func GraduationLevels() *graduationLevels {
	return _graduationLevels
}

func (o *graduationLevels) Values() []*GraduationLevel {
	return o.values
}

func (o *graduationLevels) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *graduationLevels) Unknown() *GraduationLevel {
    return _graduationLevels.values[0]
}

func (o *graduationLevels) MiddleSchool() *GraduationLevel {
    return _graduationLevels.values[1]
}

func (o *graduationLevels) SecondarySchool() *GraduationLevel {
    return _graduationLevels.values[2]
}

func (o *graduationLevels) HighSchool() *GraduationLevel {
    return _graduationLevels.values[3]
}

func (o *graduationLevels) TechnicalCollege() *GraduationLevel {
    return _graduationLevels.values[4]
}

func (o *graduationLevels) College() *GraduationLevel {
    return _graduationLevels.values[5]
}

func (o *graduationLevels) ParseGraduationLevel(name string) (ret *GraduationLevel, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*GraduationLevel), ok
	}
	return
}


type MaritalState struct {
	name  string
	ordinal int
}

func (o *MaritalState) Name() string {
    return o.name
}

func (o *MaritalState) Ordinal() int {
    return o.ordinal
}

func (o *MaritalState) IsUnknown() bool {
    return o == _maritalStates.Unknown()
}

func (o *MaritalState) IsSingle() bool {
    return o == _maritalStates.Single()
}

func (o *MaritalState) IsMarried() bool {
    return o == _maritalStates.Married()
}

func (o *MaritalState) IsSeparated() bool {
    return o == _maritalStates.Separated()
}

func (o *MaritalState) IsDivorced() bool {
    return o == _maritalStates.Divorced()
}

func (o *MaritalState) IsWidowed() bool {
    return o == _maritalStates.Widowed()
}

type maritalStates struct {
	values []*MaritalState
    literals []enum.Literal
}

var _maritalStates = &maritalStates{values: []*MaritalState{
    {name: "Unknown", ordinal: 0},
    {name: "Single", ordinal: 1},
    {name: "Married", ordinal: 2},
    {name: "Separated", ordinal: 3},
    {name: "Divorced", ordinal: 4},
    {name: "Widowed", ordinal: 5}},
}

func MaritalStates() *maritalStates {
	return _maritalStates
}

func (o *maritalStates) Values() []*MaritalState {
	return o.values
}

func (o *maritalStates) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *maritalStates) Unknown() *MaritalState {
    return _maritalStates.values[0]
}

func (o *maritalStates) Single() *MaritalState {
    return _maritalStates.values[1]
}

func (o *maritalStates) Married() *MaritalState {
    return _maritalStates.values[2]
}

func (o *maritalStates) Separated() *MaritalState {
    return _maritalStates.values[3]
}

func (o *maritalStates) Divorced() *MaritalState {
    return _maritalStates.values[4]
}

func (o *maritalStates) Widowed() *MaritalState {
    return _maritalStates.values[5]
}

func (o *maritalStates) ParseMaritalState(name string) (ret *MaritalState, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*MaritalState), ok
	}
	return
}



