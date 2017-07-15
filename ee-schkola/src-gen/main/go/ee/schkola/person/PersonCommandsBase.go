package person

import (
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/enum"
    "time"
)
const (
     CreateChurchCommand eventhorizon.CommandType = "CreateChurch"
     DeleteChurchCommand eventhorizon.CommandType = "DeleteChurch"
     UpdateChurchCommand eventhorizon.CommandType = "UpdateChurch"
)


const (
     CreateGraduationCommand eventhorizon.CommandType = "CreateGraduation"
     DeleteGraduationCommand eventhorizon.CommandType = "DeleteGraduation"
     UpdateGraduationCommand eventhorizon.CommandType = "UpdateGraduation"
)


const (
     CreateProfileCommand eventhorizon.CommandType = "CreateProfile"
     DeleteProfileCommand eventhorizon.CommandType = "DeleteProfile"
     UpdateProfileCommand eventhorizon.CommandType = "UpdateProfile"
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
func (o *CreateChurch) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *CreateChurch) AggregateType() eventhorizon.AggregateType  { return ChurchAggregateType }
func (o *CreateChurch) CommandType() eventhorizon.CommandType      { return CreateChurchCommand }



        

type DeleteChurch struct {
    Id  string
}

func NewDeleteChurch(id string) (ret *DeleteChurch, err error) {
    ret = &DeleteChurch{
        Id : id,
    }
    
    return
}
func (o *DeleteChurch) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *DeleteChurch) AggregateType() eventhorizon.AggregateType  { return ChurchAggregateType }
func (o *DeleteChurch) CommandType() eventhorizon.CommandType      { return DeleteChurchCommand }



        

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
func (o *UpdateChurch) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *UpdateChurch) AggregateType() eventhorizon.AggregateType  { return ChurchAggregateType }
func (o *UpdateChurch) CommandType() eventhorizon.CommandType      { return UpdateChurchCommand }



        

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
func (o *CreateGraduation) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *CreateGraduation) AggregateType() eventhorizon.AggregateType  { return GraduationAggregateType }
func (o *CreateGraduation) CommandType() eventhorizon.CommandType      { return CreateGraduationCommand }



        

type DeleteGraduation struct {
    Id  string
}

func NewDeleteGraduation(id string) (ret *DeleteGraduation, err error) {
    ret = &DeleteGraduation{
        Id : id,
    }
    
    return
}
func (o *DeleteGraduation) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *DeleteGraduation) AggregateType() eventhorizon.AggregateType  { return GraduationAggregateType }
func (o *DeleteGraduation) CommandType() eventhorizon.CommandType      { return DeleteGraduationCommand }



        

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
func (o *UpdateGraduation) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *UpdateGraduation) AggregateType() eventhorizon.AggregateType  { return GraduationAggregateType }
func (o *UpdateGraduation) CommandType() eventhorizon.CommandType      { return UpdateGraduationCommand }



        

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
func (o *CreateProfile) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *CreateProfile) AggregateType() eventhorizon.AggregateType  { return ProfileAggregateType }
func (o *CreateProfile) CommandType() eventhorizon.CommandType      { return CreateProfileCommand }



        

type DeleteProfile struct {
    Id  string
}

func NewDeleteProfile(id string) (ret *DeleteProfile, err error) {
    ret = &DeleteProfile{
        Id : id,
    }
    
    return
}
func (o *DeleteProfile) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *DeleteProfile) AggregateType() eventhorizon.AggregateType  { return ProfileAggregateType }
func (o *DeleteProfile) CommandType() eventhorizon.CommandType      { return DeleteProfileCommand }



        

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
func (o *UpdateProfile) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *UpdateProfile) AggregateType() eventhorizon.AggregateType  { return ProfileAggregateType }
func (o *UpdateProfile) CommandType() eventhorizon.CommandType      { return UpdateProfileCommand }





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



