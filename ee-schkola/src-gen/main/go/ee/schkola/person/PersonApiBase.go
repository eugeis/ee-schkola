package person

import (
    "ee/schkola"
    "ee/schkola/student"
    "time"
)
type Church struct {
    Name  string
    Address  *Address
    Pastor  *PersonName
    Contact  *Contact
    *schkola.SchkolaBase
}

func NewChurch(name string, address *Address, pastor *PersonName, contact *Contact, SchkolaBase *schkola.SchkolaBase) (ret *Church, err error) {
    ret = &Church{
        Name : name,
        Address : address,
        Pastor : pastor,
        Contact : contact,
        SchkolaBase: SchkolaBase,
    }
    return
}


type Graduation struct {
    Name  string
    Level  *GraduationLevel
    *schkola.SchkolaBase
}

func NewGraduation(name string, level *GraduationLevel, SchkolaBase *schkola.SchkolaBase) (ret *Graduation, err error) {
    ret = &Graduation{
        Name : name,
        Level : level,
        SchkolaBase: SchkolaBase,
    }
    return
}


type Profile struct {
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

func NewProfile(gender *Gender, name *PersonName, birthName string, birthday *time.Time, address *Address, contact *Contact, 
                photoData []byte, photo string, family *Family, church *ChurchInfo, education *Education, 
                SchkolaBase *schkola.SchkolaBase) (ret *Profile, err error) {
    ret = &Profile{
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
        SchkolaBase: SchkolaBase,
    }
    return
}






type Address struct {
    Street  string
    Suite  string
    City  string
    Code  string
    Country  string
}

func NewAddress(street string, suite string, city string, code string, country string) (ret *Address, err error) {
    ret = &Address{
        Street : street,
        Suite : suite,
        City : city,
        Code : code,
        Country : country,
    }
    return
}


type ChurchInfo struct {
    Church  string
    Association  string
    Member  bool
    Services  []*student.Course
}

func NewChurchInfo(church string, association string, member bool, services []*student.Course) (ret *ChurchInfo, err error) {
    ret = &ChurchInfo{
        Church : church,
        Association : association,
        Member : member,
        Services : services,
    }
    return
}

func (o *ChurchInfo) AddToServices(item *student.Course) *student.Course {
    o.Services  = append(o.Services , item)
    return item
}


type Contact struct {
    Phone  string
    Email  string
    Cellphone  string
}

func NewContact(phone string, email string, cellphone string) (ret *Contact, err error) {
    ret = &Contact{
        Phone : phone,
        Email : email,
        Cellphone : cellphone,
    }
    return
}


type Education struct {
    Graduation  *Graduation
    Profession  string
}

func NewEducation(graduation *Graduation, profession string) (ret *Education, err error) {
    ret = &Education{
        Graduation : graduation,
        Profession : profession,
    }
    return
}


type Family struct {
    MaritalState  *MaritalState
    ChildrenCount  int
    Partner  *PersonName
}

func NewFamily(maritalState *MaritalState, childrenCount int, partner *PersonName) (ret *Family, err error) {
    ret = &Family{
        MaritalState : maritalState,
        ChildrenCount : childrenCount,
        Partner : partner,
    }
    return
}


type PersonName struct {
    First  string
    Last  string
}

func NewPersonName(first string, last string) (ret *PersonName, err error) {
    ret = &PersonName{
        First : first,
        Last : last,
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
    switch name {
      case o.Unknown().Name():
        ret = o.Unknown()
      case o.Male().Name():
        ret = o.Male()
      case o.Female().Name():
        ret = o.Female()
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
    switch name {
      case o.Unknown().Name():
        ret = o.Unknown()
      case o.MiddleSchool().Name():
        ret = o.MiddleSchool()
      case o.SecondarySchool().Name():
        ret = o.SecondarySchool()
      case o.HighSchool().Name():
        ret = o.HighSchool()
      case o.TechnicalCollege().Name():
        ret = o.TechnicalCollege()
      case o.College().Name():
        ret = o.College()
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
    switch name {
      case o.Unknown().Name():
        ret = o.Unknown()
      case o.Single().Name():
        ret = o.Single()
      case o.Married().Name():
        ret = o.Married()
      case o.Separated().Name():
        ret = o.Separated()
      case o.Divorced().Name():
        ret = o.Divorced()
      case o.Widowed().Name():
        ret = o.Widowed()
    }
    return
}



