package student

import (
    "ee/schkola/person"
    "ee/schkola/shared"
    "encoding/json"
    "fmt"
    "github.com/eugeis/gee/enum"
    "github.com/looplab/eventhorizon"
    "gopkg.in/mgo.v2/bson"
    "time"
)
const (
     RegisterCommand eventhorizon.CommandType = "Register"
     CreateCommand eventhorizon.CommandType = "Create"
     DeleteCommand eventhorizon.CommandType = "Delete"
     CancelCommand eventhorizon.CommandType = "Cancel"
     ConfirmCommand eventhorizon.CommandType = "Confirm"
     UpdateCommand eventhorizon.CommandType = "Update"
)


const (
     CreateCommand eventhorizon.CommandType = "Create"
     DeleteCommand eventhorizon.CommandType = "Delete"
     UpdateCommand eventhorizon.CommandType = "Update"
)


const (
     CreateCommand eventhorizon.CommandType = "Create"
     DeleteCommand eventhorizon.CommandType = "Delete"
     UpdateCommand eventhorizon.CommandType = "Update"
)


const (
     CreateCommand eventhorizon.CommandType = "Create"
     DeleteCommand eventhorizon.CommandType = "Delete"
     UpdateCommand eventhorizon.CommandType = "Update"
)


const (
     CreateCommand eventhorizon.CommandType = "Create"
     DeleteCommand eventhorizon.CommandType = "Delete"
     UpdateCommand eventhorizon.CommandType = "Update"
)


const (
     CreateCommand eventhorizon.CommandType = "Create"
     DeleteCommand eventhorizon.CommandType = "Delete"
     UpdateCommand eventhorizon.CommandType = "Update"
)




        
type Register struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Register) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Register) AggregateType() eventhorizon.AggregateType  { return AttendanceAggregateType }
func (o *Register) CommandType() eventhorizon.CommandType      { return RegisterCommand }



        
type Create struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Hours int `json:"hours" eh:"optional"`
    State *AttendanceState `json:"state" eh:"optional"`
    Token string `json:"token" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Create) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Create) AggregateType() eventhorizon.AggregateType  { return AttendanceAggregateType }
func (o *Create) CommandType() eventhorizon.CommandType      { return CreateCommand }



        
type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Delete) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Delete) AggregateType() eventhorizon.AggregateType  { return AttendanceAggregateType }
func (o *Delete) CommandType() eventhorizon.CommandType      { return DeleteCommand }



        
type Cancel struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Cancel) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Cancel) AggregateType() eventhorizon.AggregateType  { return AttendanceAggregateType }
func (o *Cancel) CommandType() eventhorizon.CommandType      { return CancelCommand }



        
type Confirm struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Confirm) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Confirm) AggregateType() eventhorizon.AggregateType  { return AttendanceAggregateType }
func (o *Confirm) CommandType() eventhorizon.CommandType      { return ConfirmCommand }



        
type Update struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Hours int `json:"hours" eh:"optional"`
    State *AttendanceState `json:"state" eh:"optional"`
    Token string `json:"token" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Update) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Update) AggregateType() eventhorizon.AggregateType  { return AttendanceAggregateType }
func (o *Update) CommandType() eventhorizon.CommandType      { return UpdateCommand }



        
type Create struct {
    Name string `json:"name" eh:"optional"`
    Begin *time.Time `json:"begin" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Teacher *shared.PersonName `json:"teacher" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Fee float64 `json:"fee" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Create) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Create) AggregateType() eventhorizon.AggregateType  { return CourseAggregateType }
func (o *Create) CommandType() eventhorizon.CommandType      { return CreateCommand }



        
type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Delete) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Delete) AggregateType() eventhorizon.AggregateType  { return CourseAggregateType }
func (o *Delete) CommandType() eventhorizon.CommandType      { return DeleteCommand }



        
type Update struct {
    Name string `json:"name" eh:"optional"`
    Begin *time.Time `json:"begin" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Teacher *shared.PersonName `json:"teacher" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Fee float64 `json:"fee" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Update) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Update) AggregateType() eventhorizon.AggregateType  { return CourseAggregateType }
func (o *Update) CommandType() eventhorizon.CommandType      { return UpdateCommand }



        
type Create struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Grade float64 `json:"grade" eh:"optional"`
    Comment string `json:"comment" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Create) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Create) AggregateType() eventhorizon.AggregateType  { return GradeAggregateType }
func (o *Create) CommandType() eventhorizon.CommandType      { return CreateCommand }



        
type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Delete) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Delete) AggregateType() eventhorizon.AggregateType  { return GradeAggregateType }
func (o *Delete) CommandType() eventhorizon.CommandType      { return DeleteCommand }



        
type Update struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Grade float64 `json:"grade" eh:"optional"`
    Comment string `json:"comment" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Update) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Update) AggregateType() eventhorizon.AggregateType  { return GradeAggregateType }
func (o *Update) CommandType() eventhorizon.CommandType      { return UpdateCommand }



        
type Create struct {
    Name string `json:"name" eh:"optional"`
    Category *GroupCategory `json:"category" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Representative *person.Profile `json:"representative" eh:"optional"`
    Students []*person.Profile `json:"students" eh:"optional"`
    Courses []*Course `json:"courses" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *Create) AddToStudents(item *person.Profile) *person.Profile {
    o.Students = append(o.Students, item)
    return item
}

func (o *Create) AddToCourses(item *Course) *Course {
    o.Courses = append(o.Courses, item)
    return item
}
func (o *Create) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Create) AggregateType() eventhorizon.AggregateType  { return GroupAggregateType }
func (o *Create) CommandType() eventhorizon.CommandType      { return CreateCommand }



        
type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Delete) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Delete) AggregateType() eventhorizon.AggregateType  { return GroupAggregateType }
func (o *Delete) CommandType() eventhorizon.CommandType      { return DeleteCommand }



        
type Update struct {
    Name string `json:"name" eh:"optional"`
    Category *GroupCategory `json:"category" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Representative *person.Profile `json:"representative" eh:"optional"`
    Students []*person.Profile `json:"students" eh:"optional"`
    Courses []*Course `json:"courses" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *Update) AddToStudents(item *person.Profile) *person.Profile {
    o.Students = append(o.Students, item)
    return item
}

func (o *Update) AddToCourses(item *Course) *Course {
    o.Courses = append(o.Courses, item)
    return item
}
func (o *Update) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Update) AggregateType() eventhorizon.AggregateType  { return GroupAggregateType }
func (o *Update) CommandType() eventhorizon.CommandType      { return UpdateCommand }



        
type Create struct {
    Profile *person.Profile `json:"profile" eh:"optional"`
    ChurchContactPerson *shared.PersonName `json:"churchContactPerson" eh:"optional"`
    ChurchContact *person.Contact `json:"churchContact" eh:"optional"`
    ChurchCommitment bool `json:"churchCommitment" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Group string `json:"group" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Create) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Create) AggregateType() eventhorizon.AggregateType  { return SchoolApplicationAggregateType }
func (o *Create) CommandType() eventhorizon.CommandType      { return CreateCommand }



        
type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Delete) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Delete) AggregateType() eventhorizon.AggregateType  { return SchoolApplicationAggregateType }
func (o *Delete) CommandType() eventhorizon.CommandType      { return DeleteCommand }



        
type Update struct {
    Profile *person.Profile `json:"profile" eh:"optional"`
    ChurchContactPerson *shared.PersonName `json:"churchContactPerson" eh:"optional"`
    ChurchContact *person.Contact `json:"churchContact" eh:"optional"`
    ChurchCommitment bool `json:"churchCommitment" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Group string `json:"group" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Update) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Update) AggregateType() eventhorizon.AggregateType  { return SchoolApplicationAggregateType }
func (o *Update) CommandType() eventhorizon.CommandType      { return UpdateCommand }



        
type Create struct {
    Name string `json:"name" eh:"optional"`
    Start *time.Time `json:"start" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Dates []*time.Time `json:"dates" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *Create) AddToDates(item *time.Time) *time.Time {
    o.Dates = append(o.Dates, item)
    return item
}
func (o *Create) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Create) AggregateType() eventhorizon.AggregateType  { return SchoolYearAggregateType }
func (o *Create) CommandType() eventhorizon.CommandType      { return CreateCommand }



        
type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *Delete) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Delete) AggregateType() eventhorizon.AggregateType  { return SchoolYearAggregateType }
func (o *Delete) CommandType() eventhorizon.CommandType      { return DeleteCommand }



        
type Update struct {
    Name string `json:"name" eh:"optional"`
    Start *time.Time `json:"start" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Dates []*time.Time `json:"dates" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *Update) AddToDates(item *time.Time) *time.Time {
    o.Dates = append(o.Dates, item)
    return item
}
func (o *Update) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *Update) AggregateType() eventhorizon.AggregateType  { return SchoolYearAggregateType }
func (o *Update) CommandType() eventhorizon.CommandType      { return UpdateCommand }





type AttendanceCommandType struct {
	name  string
	ordinal int
}

func (o *AttendanceCommandType) Name() string {
    return o.name
}

func (o *AttendanceCommandType) Ordinal() int {
    return o.ordinal
}

func (o AttendanceCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *AttendanceCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := AttendanceCommandTypes().ParseAttendanceCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid AttendanceCommandType %q", lit.Name)
        }
	}
	return
}

func (o AttendanceCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *AttendanceCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := AttendanceCommandTypes().ParseAttendanceCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid AttendanceCommandType %q", lit)
        }
    }
    return
}

func (o *AttendanceCommandType) IsRegister() bool {
    return o == _attendanceCommandTypes.Register()
}

func (o *AttendanceCommandType) IsCreate() bool {
    return o == _attendanceCommandTypes.Create()
}

func (o *AttendanceCommandType) IsDelete() bool {
    return o == _attendanceCommandTypes.Delete()
}

func (o *AttendanceCommandType) IsCancel() bool {
    return o == _attendanceCommandTypes.Cancel()
}

func (o *AttendanceCommandType) IsConfirm() bool {
    return o == _attendanceCommandTypes.Confirm()
}

func (o *AttendanceCommandType) IsUpdate() bool {
    return o == _attendanceCommandTypes.Update()
}

type attendanceCommandTypes struct {
	values []*AttendanceCommandType
    literals []enum.Literal
}

var _attendanceCommandTypes = &attendanceCommandTypes{values: []*AttendanceCommandType{
    {name: "Register", ordinal: 0},
    {name: "Create", ordinal: 1},
    {name: "Delete", ordinal: 2},
    {name: "Cancel", ordinal: 3},
    {name: "Confirm", ordinal: 4},
    {name: "Update", ordinal: 5}},
}

func AttendanceCommandTypes() *attendanceCommandTypes {
	return _attendanceCommandTypes
}

func (o *attendanceCommandTypes) Values() []*AttendanceCommandType {
	return o.values
}

func (o *attendanceCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *attendanceCommandTypes) Register() *AttendanceCommandType {
    return _attendanceCommandTypes.values[0]
}

func (o *attendanceCommandTypes) Create() *AttendanceCommandType {
    return _attendanceCommandTypes.values[1]
}

func (o *attendanceCommandTypes) Delete() *AttendanceCommandType {
    return _attendanceCommandTypes.values[2]
}

func (o *attendanceCommandTypes) Cancel() *AttendanceCommandType {
    return _attendanceCommandTypes.values[3]
}

func (o *attendanceCommandTypes) Confirm() *AttendanceCommandType {
    return _attendanceCommandTypes.values[4]
}

func (o *attendanceCommandTypes) Update() *AttendanceCommandType {
    return _attendanceCommandTypes.values[5]
}

func (o *attendanceCommandTypes) ParseAttendanceCommandType(name string) (ret *AttendanceCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*AttendanceCommandType), ok
	}
	return
}


type CourseCommandType struct {
	name  string
	ordinal int
}

func (o *CourseCommandType) Name() string {
    return o.name
}

func (o *CourseCommandType) Ordinal() int {
    return o.ordinal
}

func (o CourseCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *CourseCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := CourseCommandTypes().ParseCourseCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid CourseCommandType %q", lit.Name)
        }
	}
	return
}

func (o CourseCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *CourseCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := CourseCommandTypes().ParseCourseCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid CourseCommandType %q", lit)
        }
    }
    return
}

func (o *CourseCommandType) IsCreate() bool {
    return o == _courseCommandTypes.Create()
}

func (o *CourseCommandType) IsDelete() bool {
    return o == _courseCommandTypes.Delete()
}

func (o *CourseCommandType) IsUpdate() bool {
    return o == _courseCommandTypes.Update()
}

type courseCommandTypes struct {
	values []*CourseCommandType
    literals []enum.Literal
}

var _courseCommandTypes = &courseCommandTypes{values: []*CourseCommandType{
    {name: "Create", ordinal: 0},
    {name: "Delete", ordinal: 1},
    {name: "Update", ordinal: 2}},
}

func CourseCommandTypes() *courseCommandTypes {
	return _courseCommandTypes
}

func (o *courseCommandTypes) Values() []*CourseCommandType {
	return o.values
}

func (o *courseCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *courseCommandTypes) Create() *CourseCommandType {
    return _courseCommandTypes.values[0]
}

func (o *courseCommandTypes) Delete() *CourseCommandType {
    return _courseCommandTypes.values[1]
}

func (o *courseCommandTypes) Update() *CourseCommandType {
    return _courseCommandTypes.values[2]
}

func (o *courseCommandTypes) ParseCourseCommandType(name string) (ret *CourseCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*CourseCommandType), ok
	}
	return
}


type GradeCommandType struct {
	name  string
	ordinal int
}

func (o *GradeCommandType) Name() string {
    return o.name
}

func (o *GradeCommandType) Ordinal() int {
    return o.ordinal
}

func (o GradeCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *GradeCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := GradeCommandTypes().ParseGradeCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GradeCommandType %q", lit.Name)
        }
	}
	return
}

func (o GradeCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *GradeCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := GradeCommandTypes().ParseGradeCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GradeCommandType %q", lit)
        }
    }
    return
}

func (o *GradeCommandType) IsCreate() bool {
    return o == _gradeCommandTypes.Create()
}

func (o *GradeCommandType) IsDelete() bool {
    return o == _gradeCommandTypes.Delete()
}

func (o *GradeCommandType) IsUpdate() bool {
    return o == _gradeCommandTypes.Update()
}

type gradeCommandTypes struct {
	values []*GradeCommandType
    literals []enum.Literal
}

var _gradeCommandTypes = &gradeCommandTypes{values: []*GradeCommandType{
    {name: "Create", ordinal: 0},
    {name: "Delete", ordinal: 1},
    {name: "Update", ordinal: 2}},
}

func GradeCommandTypes() *gradeCommandTypes {
	return _gradeCommandTypes
}

func (o *gradeCommandTypes) Values() []*GradeCommandType {
	return o.values
}

func (o *gradeCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *gradeCommandTypes) Create() *GradeCommandType {
    return _gradeCommandTypes.values[0]
}

func (o *gradeCommandTypes) Delete() *GradeCommandType {
    return _gradeCommandTypes.values[1]
}

func (o *gradeCommandTypes) Update() *GradeCommandType {
    return _gradeCommandTypes.values[2]
}

func (o *gradeCommandTypes) ParseGradeCommandType(name string) (ret *GradeCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*GradeCommandType), ok
	}
	return
}


type GroupCommandType struct {
	name  string
	ordinal int
}

func (o *GroupCommandType) Name() string {
    return o.name
}

func (o *GroupCommandType) Ordinal() int {
    return o.ordinal
}

func (o GroupCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *GroupCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := GroupCommandTypes().ParseGroupCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GroupCommandType %q", lit.Name)
        }
	}
	return
}

func (o GroupCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *GroupCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := GroupCommandTypes().ParseGroupCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GroupCommandType %q", lit)
        }
    }
    return
}

func (o *GroupCommandType) IsCreate() bool {
    return o == _groupCommandTypes.Create()
}

func (o *GroupCommandType) IsDelete() bool {
    return o == _groupCommandTypes.Delete()
}

func (o *GroupCommandType) IsUpdate() bool {
    return o == _groupCommandTypes.Update()
}

type groupCommandTypes struct {
	values []*GroupCommandType
    literals []enum.Literal
}

var _groupCommandTypes = &groupCommandTypes{values: []*GroupCommandType{
    {name: "Create", ordinal: 0},
    {name: "Delete", ordinal: 1},
    {name: "Update", ordinal: 2}},
}

func GroupCommandTypes() *groupCommandTypes {
	return _groupCommandTypes
}

func (o *groupCommandTypes) Values() []*GroupCommandType {
	return o.values
}

func (o *groupCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *groupCommandTypes) Create() *GroupCommandType {
    return _groupCommandTypes.values[0]
}

func (o *groupCommandTypes) Delete() *GroupCommandType {
    return _groupCommandTypes.values[1]
}

func (o *groupCommandTypes) Update() *GroupCommandType {
    return _groupCommandTypes.values[2]
}

func (o *groupCommandTypes) ParseGroupCommandType(name string) (ret *GroupCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*GroupCommandType), ok
	}
	return
}


type SchoolApplicationCommandType struct {
	name  string
	ordinal int
}

func (o *SchoolApplicationCommandType) Name() string {
    return o.name
}

func (o *SchoolApplicationCommandType) Ordinal() int {
    return o.ordinal
}

func (o SchoolApplicationCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *SchoolApplicationCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := SchoolApplicationCommandTypes().ParseSchoolApplicationCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid SchoolApplicationCommandType %q", lit.Name)
        }
	}
	return
}

func (o SchoolApplicationCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *SchoolApplicationCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := SchoolApplicationCommandTypes().ParseSchoolApplicationCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid SchoolApplicationCommandType %q", lit)
        }
    }
    return
}

func (o *SchoolApplicationCommandType) IsCreate() bool {
    return o == _schoolApplicationCommandTypes.Create()
}

func (o *SchoolApplicationCommandType) IsDelete() bool {
    return o == _schoolApplicationCommandTypes.Delete()
}

func (o *SchoolApplicationCommandType) IsUpdate() bool {
    return o == _schoolApplicationCommandTypes.Update()
}

type schoolApplicationCommandTypes struct {
	values []*SchoolApplicationCommandType
    literals []enum.Literal
}

var _schoolApplicationCommandTypes = &schoolApplicationCommandTypes{values: []*SchoolApplicationCommandType{
    {name: "Create", ordinal: 0},
    {name: "Delete", ordinal: 1},
    {name: "Update", ordinal: 2}},
}

func SchoolApplicationCommandTypes() *schoolApplicationCommandTypes {
	return _schoolApplicationCommandTypes
}

func (o *schoolApplicationCommandTypes) Values() []*SchoolApplicationCommandType {
	return o.values
}

func (o *schoolApplicationCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *schoolApplicationCommandTypes) Create() *SchoolApplicationCommandType {
    return _schoolApplicationCommandTypes.values[0]
}

func (o *schoolApplicationCommandTypes) Delete() *SchoolApplicationCommandType {
    return _schoolApplicationCommandTypes.values[1]
}

func (o *schoolApplicationCommandTypes) Update() *SchoolApplicationCommandType {
    return _schoolApplicationCommandTypes.values[2]
}

func (o *schoolApplicationCommandTypes) ParseSchoolApplicationCommandType(name string) (ret *SchoolApplicationCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*SchoolApplicationCommandType), ok
	}
	return
}


type SchoolYearCommandType struct {
	name  string
	ordinal int
}

func (o *SchoolYearCommandType) Name() string {
    return o.name
}

func (o *SchoolYearCommandType) Ordinal() int {
    return o.ordinal
}

func (o SchoolYearCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *SchoolYearCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := SchoolYearCommandTypes().ParseSchoolYearCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid SchoolYearCommandType %q", lit.Name)
        }
	}
	return
}

func (o SchoolYearCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *SchoolYearCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := SchoolYearCommandTypes().ParseSchoolYearCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid SchoolYearCommandType %q", lit)
        }
    }
    return
}

func (o *SchoolYearCommandType) IsCreate() bool {
    return o == _schoolYearCommandTypes.Create()
}

func (o *SchoolYearCommandType) IsDelete() bool {
    return o == _schoolYearCommandTypes.Delete()
}

func (o *SchoolYearCommandType) IsUpdate() bool {
    return o == _schoolYearCommandTypes.Update()
}

type schoolYearCommandTypes struct {
	values []*SchoolYearCommandType
    literals []enum.Literal
}

var _schoolYearCommandTypes = &schoolYearCommandTypes{values: []*SchoolYearCommandType{
    {name: "Create", ordinal: 0},
    {name: "Delete", ordinal: 1},
    {name: "Update", ordinal: 2}},
}

func SchoolYearCommandTypes() *schoolYearCommandTypes {
	return _schoolYearCommandTypes
}

func (o *schoolYearCommandTypes) Values() []*SchoolYearCommandType {
	return o.values
}

func (o *schoolYearCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *schoolYearCommandTypes) Create() *SchoolYearCommandType {
    return _schoolYearCommandTypes.values[0]
}

func (o *schoolYearCommandTypes) Delete() *SchoolYearCommandType {
    return _schoolYearCommandTypes.values[1]
}

func (o *schoolYearCommandTypes) Update() *SchoolYearCommandType {
    return _schoolYearCommandTypes.values[2]
}

func (o *schoolYearCommandTypes) ParseSchoolYearCommandType(name string) (ret *SchoolYearCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*SchoolYearCommandType), ok
	}
	return
}



