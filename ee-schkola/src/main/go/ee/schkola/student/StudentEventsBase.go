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
     RegisterEvent eventhorizon.EventType = "Register"
     CreateEvent eventhorizon.EventType = "Create"
     CreatedEvent eventhorizon.EventType = "Created"
     RegisteredEvent eventhorizon.EventType = "Registered"
     DeleteEvent eventhorizon.EventType = "Delete"
     DeletedEvent eventhorizon.EventType = "Deleted"
     CancelEvent eventhorizon.EventType = "Cancel"
     ConfirmEvent eventhorizon.EventType = "Confirm"
     UpdateEvent eventhorizon.EventType = "Update"
     UpdatedEvent eventhorizon.EventType = "Updated"
     ConfirmedEvent eventhorizon.EventType = "Confirmed"
     CanceledEvent eventhorizon.EventType = "Canceled"
)


const (
     CreateEvent eventhorizon.EventType = "Create"
     CreatedEvent eventhorizon.EventType = "Created"
     DeleteEvent eventhorizon.EventType = "Delete"
     DeletedEvent eventhorizon.EventType = "Deleted"
     UpdateEvent eventhorizon.EventType = "Update"
     UpdatedEvent eventhorizon.EventType = "Updated"
)


const (
     CreateEvent eventhorizon.EventType = "Create"
     CreatedEvent eventhorizon.EventType = "Created"
     DeleteEvent eventhorizon.EventType = "Delete"
     DeletedEvent eventhorizon.EventType = "Deleted"
     UpdateEvent eventhorizon.EventType = "Update"
     UpdatedEvent eventhorizon.EventType = "Updated"
)


const (
     CreateEvent eventhorizon.EventType = "Create"
     CreatedEvent eventhorizon.EventType = "Created"
     DeleteEvent eventhorizon.EventType = "Delete"
     DeletedEvent eventhorizon.EventType = "Deleted"
     UpdateEvent eventhorizon.EventType = "Update"
     UpdatedEvent eventhorizon.EventType = "Updated"
)


const (
     CreateEvent eventhorizon.EventType = "Create"
     CreatedEvent eventhorizon.EventType = "Created"
     DeleteEvent eventhorizon.EventType = "Delete"
     DeletedEvent eventhorizon.EventType = "Deleted"
     UpdateEvent eventhorizon.EventType = "Update"
     UpdatedEvent eventhorizon.EventType = "Updated"
)


const (
     CreateEvent eventhorizon.EventType = "Create"
     CreatedEvent eventhorizon.EventType = "Created"
     DeleteEvent eventhorizon.EventType = "Delete"
     DeletedEvent eventhorizon.EventType = "Deleted"
     UpdateEvent eventhorizon.EventType = "Update"
     UpdatedEvent eventhorizon.EventType = "Updated"
)




type Register struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Create struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Hours int `json:"hours" eh:"optional"`
    State *AttendanceState `json:"state" eh:"optional"`
    Token string `json:"token" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Created struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Hours int `json:"hours" eh:"optional"`
    State *AttendanceState `json:"state" eh:"optional"`
    Token string `json:"token" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Registered struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Deleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Cancel struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Confirm struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Update struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Hours int `json:"hours" eh:"optional"`
    State *AttendanceState `json:"state" eh:"optional"`
    Token string `json:"token" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Updated struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Hours int `json:"hours" eh:"optional"`
    State *AttendanceState `json:"state" eh:"optional"`
    Token string `json:"token" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Confirmed struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Canceled struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


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


type Created struct {
    Name string `json:"name" eh:"optional"`
    Begin *time.Time `json:"begin" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Teacher *shared.PersonName `json:"teacher" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Fee float64 `json:"fee" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Deleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


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


type Updated struct {
    Name string `json:"name" eh:"optional"`
    Begin *time.Time `json:"begin" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Teacher *shared.PersonName `json:"teacher" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Fee float64 `json:"fee" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Create struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Grade float64 `json:"grade" eh:"optional"`
    Comment string `json:"comment" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Created struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Grade float64 `json:"grade" eh:"optional"`
    Comment string `json:"comment" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Deleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Update struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Grade float64 `json:"grade" eh:"optional"`
    Comment string `json:"comment" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Updated struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Grade float64 `json:"grade" eh:"optional"`
    Comment string `json:"comment" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


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


type Created struct {
    Name string `json:"name" eh:"optional"`
    Category *GroupCategory `json:"category" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Representative *person.Profile `json:"representative" eh:"optional"`
    Students []*person.Profile `json:"students" eh:"optional"`
    Courses []*Course `json:"courses" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *Created) AddToStudents(item *person.Profile) *person.Profile {
    o.Students = append(o.Students, item)
    return item
}

func (o *Created) AddToCourses(item *Course) *Course {
    o.Courses = append(o.Courses, item)
    return item
}


type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Deleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


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


type Updated struct {
    Name string `json:"name" eh:"optional"`
    Category *GroupCategory `json:"category" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Representative *person.Profile `json:"representative" eh:"optional"`
    Students []*person.Profile `json:"students" eh:"optional"`
    Courses []*Course `json:"courses" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *Updated) AddToStudents(item *person.Profile) *person.Profile {
    o.Students = append(o.Students, item)
    return item
}

func (o *Updated) AddToCourses(item *Course) *Course {
    o.Courses = append(o.Courses, item)
    return item
}


type Create struct {
    Profile *person.Profile `json:"profile" eh:"optional"`
    ChurchContactPerson *shared.PersonName `json:"churchContactPerson" eh:"optional"`
    ChurchContact *person.Contact `json:"churchContact" eh:"optional"`
    ChurchCommitment bool `json:"churchCommitment" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Group string `json:"group" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Created struct {
    Profile *person.Profile `json:"profile" eh:"optional"`
    ChurchContactPerson *shared.PersonName `json:"churchContactPerson" eh:"optional"`
    ChurchContact *person.Contact `json:"churchContact" eh:"optional"`
    ChurchCommitment bool `json:"churchCommitment" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Group string `json:"group" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Deleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Update struct {
    Profile *person.Profile `json:"profile" eh:"optional"`
    ChurchContactPerson *shared.PersonName `json:"churchContactPerson" eh:"optional"`
    ChurchContact *person.Contact `json:"churchContact" eh:"optional"`
    ChurchCommitment bool `json:"churchCommitment" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Group string `json:"group" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Updated struct {
    Profile *person.Profile `json:"profile" eh:"optional"`
    ChurchContactPerson *shared.PersonName `json:"churchContactPerson" eh:"optional"`
    ChurchContact *person.Contact `json:"churchContact" eh:"optional"`
    ChurchCommitment bool `json:"churchCommitment" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Group string `json:"group" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


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


type Created struct {
    Name string `json:"name" eh:"optional"`
    Start *time.Time `json:"start" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Dates []*time.Time `json:"dates" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *Created) AddToDates(item *time.Time) *time.Time {
    o.Dates = append(o.Dates, item)
    return item
}


type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Deleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


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


type Updated struct {
    Name string `json:"name" eh:"optional"`
    Start *time.Time `json:"start" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Dates []*time.Time `json:"dates" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *Updated) AddToDates(item *time.Time) *time.Time {
    o.Dates = append(o.Dates, item)
    return item
}




type AttendanceEventType struct {
	name  string
	ordinal int
}

func (o *AttendanceEventType) Name() string {
    return o.name
}

func (o *AttendanceEventType) Ordinal() int {
    return o.ordinal
}

func (o AttendanceEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *AttendanceEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := AttendanceEventTypes().ParseAttendanceEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid AttendanceEventType %q", lit.Name)
        }
	}
	return
}

func (o AttendanceEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *AttendanceEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := AttendanceEventTypes().ParseAttendanceEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid AttendanceEventType %q", lit)
        }
    }
    return
}

func (o *AttendanceEventType) IsRegister() bool {
    return o == _attendanceEventTypes.Register()
}

func (o *AttendanceEventType) IsCreate() bool {
    return o == _attendanceEventTypes.Create()
}

func (o *AttendanceEventType) IsCreated() bool {
    return o == _attendanceEventTypes.Created()
}

func (o *AttendanceEventType) IsRegistered() bool {
    return o == _attendanceEventTypes.Registered()
}

func (o *AttendanceEventType) IsDelete() bool {
    return o == _attendanceEventTypes.Delete()
}

func (o *AttendanceEventType) IsDeleted() bool {
    return o == _attendanceEventTypes.Deleted()
}

func (o *AttendanceEventType) IsCancel() bool {
    return o == _attendanceEventTypes.Cancel()
}

func (o *AttendanceEventType) IsConfirm() bool {
    return o == _attendanceEventTypes.Confirm()
}

func (o *AttendanceEventType) IsUpdate() bool {
    return o == _attendanceEventTypes.Update()
}

func (o *AttendanceEventType) IsUpdated() bool {
    return o == _attendanceEventTypes.Updated()
}

func (o *AttendanceEventType) IsConfirmed() bool {
    return o == _attendanceEventTypes.Confirmed()
}

func (o *AttendanceEventType) IsCanceled() bool {
    return o == _attendanceEventTypes.Canceled()
}

type attendanceEventTypes struct {
	values []*AttendanceEventType
    literals []enum.Literal
}

var _attendanceEventTypes = &attendanceEventTypes{values: []*AttendanceEventType{
    {name: "Register", ordinal: 0},
    {name: "Create", ordinal: 1},
    {name: "Created", ordinal: 2},
    {name: "Registered", ordinal: 3},
    {name: "Delete", ordinal: 4},
    {name: "Deleted", ordinal: 5},
    {name: "Cancel", ordinal: 6},
    {name: "Confirm", ordinal: 7},
    {name: "Update", ordinal: 8},
    {name: "Updated", ordinal: 9},
    {name: "Confirmed", ordinal: 10},
    {name: "Canceled", ordinal: 11}},
}

func AttendanceEventTypes() *attendanceEventTypes {
	return _attendanceEventTypes
}

func (o *attendanceEventTypes) Values() []*AttendanceEventType {
	return o.values
}

func (o *attendanceEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *attendanceEventTypes) Register() *AttendanceEventType {
    return _attendanceEventTypes.values[0]
}

func (o *attendanceEventTypes) Create() *AttendanceEventType {
    return _attendanceEventTypes.values[1]
}

func (o *attendanceEventTypes) Created() *AttendanceEventType {
    return _attendanceEventTypes.values[2]
}

func (o *attendanceEventTypes) Registered() *AttendanceEventType {
    return _attendanceEventTypes.values[3]
}

func (o *attendanceEventTypes) Delete() *AttendanceEventType {
    return _attendanceEventTypes.values[4]
}

func (o *attendanceEventTypes) Deleted() *AttendanceEventType {
    return _attendanceEventTypes.values[5]
}

func (o *attendanceEventTypes) Cancel() *AttendanceEventType {
    return _attendanceEventTypes.values[6]
}

func (o *attendanceEventTypes) Confirm() *AttendanceEventType {
    return _attendanceEventTypes.values[7]
}

func (o *attendanceEventTypes) Update() *AttendanceEventType {
    return _attendanceEventTypes.values[8]
}

func (o *attendanceEventTypes) Updated() *AttendanceEventType {
    return _attendanceEventTypes.values[9]
}

func (o *attendanceEventTypes) Confirmed() *AttendanceEventType {
    return _attendanceEventTypes.values[10]
}

func (o *attendanceEventTypes) Canceled() *AttendanceEventType {
    return _attendanceEventTypes.values[11]
}

func (o *attendanceEventTypes) ParseAttendanceEventType(name string) (ret *AttendanceEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*AttendanceEventType), ok
	}
	return
}


type CourseEventType struct {
	name  string
	ordinal int
}

func (o *CourseEventType) Name() string {
    return o.name
}

func (o *CourseEventType) Ordinal() int {
    return o.ordinal
}

func (o CourseEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *CourseEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := CourseEventTypes().ParseCourseEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid CourseEventType %q", lit.Name)
        }
	}
	return
}

func (o CourseEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *CourseEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := CourseEventTypes().ParseCourseEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid CourseEventType %q", lit)
        }
    }
    return
}

func (o *CourseEventType) IsCreate() bool {
    return o == _courseEventTypes.Create()
}

func (o *CourseEventType) IsCreated() bool {
    return o == _courseEventTypes.Created()
}

func (o *CourseEventType) IsDelete() bool {
    return o == _courseEventTypes.Delete()
}

func (o *CourseEventType) IsDeleted() bool {
    return o == _courseEventTypes.Deleted()
}

func (o *CourseEventType) IsUpdate() bool {
    return o == _courseEventTypes.Update()
}

func (o *CourseEventType) IsUpdated() bool {
    return o == _courseEventTypes.Updated()
}

type courseEventTypes struct {
	values []*CourseEventType
    literals []enum.Literal
}

var _courseEventTypes = &courseEventTypes{values: []*CourseEventType{
    {name: "Create", ordinal: 0},
    {name: "Created", ordinal: 1},
    {name: "Delete", ordinal: 2},
    {name: "Deleted", ordinal: 3},
    {name: "Update", ordinal: 4},
    {name: "Updated", ordinal: 5}},
}

func CourseEventTypes() *courseEventTypes {
	return _courseEventTypes
}

func (o *courseEventTypes) Values() []*CourseEventType {
	return o.values
}

func (o *courseEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *courseEventTypes) Create() *CourseEventType {
    return _courseEventTypes.values[0]
}

func (o *courseEventTypes) Created() *CourseEventType {
    return _courseEventTypes.values[1]
}

func (o *courseEventTypes) Delete() *CourseEventType {
    return _courseEventTypes.values[2]
}

func (o *courseEventTypes) Deleted() *CourseEventType {
    return _courseEventTypes.values[3]
}

func (o *courseEventTypes) Update() *CourseEventType {
    return _courseEventTypes.values[4]
}

func (o *courseEventTypes) Updated() *CourseEventType {
    return _courseEventTypes.values[5]
}

func (o *courseEventTypes) ParseCourseEventType(name string) (ret *CourseEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*CourseEventType), ok
	}
	return
}


type GradeEventType struct {
	name  string
	ordinal int
}

func (o *GradeEventType) Name() string {
    return o.name
}

func (o *GradeEventType) Ordinal() int {
    return o.ordinal
}

func (o GradeEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *GradeEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := GradeEventTypes().ParseGradeEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GradeEventType %q", lit.Name)
        }
	}
	return
}

func (o GradeEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *GradeEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := GradeEventTypes().ParseGradeEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GradeEventType %q", lit)
        }
    }
    return
}

func (o *GradeEventType) IsCreate() bool {
    return o == _gradeEventTypes.Create()
}

func (o *GradeEventType) IsCreated() bool {
    return o == _gradeEventTypes.Created()
}

func (o *GradeEventType) IsDelete() bool {
    return o == _gradeEventTypes.Delete()
}

func (o *GradeEventType) IsDeleted() bool {
    return o == _gradeEventTypes.Deleted()
}

func (o *GradeEventType) IsUpdate() bool {
    return o == _gradeEventTypes.Update()
}

func (o *GradeEventType) IsUpdated() bool {
    return o == _gradeEventTypes.Updated()
}

type gradeEventTypes struct {
	values []*GradeEventType
    literals []enum.Literal
}

var _gradeEventTypes = &gradeEventTypes{values: []*GradeEventType{
    {name: "Create", ordinal: 0},
    {name: "Created", ordinal: 1},
    {name: "Delete", ordinal: 2},
    {name: "Deleted", ordinal: 3},
    {name: "Update", ordinal: 4},
    {name: "Updated", ordinal: 5}},
}

func GradeEventTypes() *gradeEventTypes {
	return _gradeEventTypes
}

func (o *gradeEventTypes) Values() []*GradeEventType {
	return o.values
}

func (o *gradeEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *gradeEventTypes) Create() *GradeEventType {
    return _gradeEventTypes.values[0]
}

func (o *gradeEventTypes) Created() *GradeEventType {
    return _gradeEventTypes.values[1]
}

func (o *gradeEventTypes) Delete() *GradeEventType {
    return _gradeEventTypes.values[2]
}

func (o *gradeEventTypes) Deleted() *GradeEventType {
    return _gradeEventTypes.values[3]
}

func (o *gradeEventTypes) Update() *GradeEventType {
    return _gradeEventTypes.values[4]
}

func (o *gradeEventTypes) Updated() *GradeEventType {
    return _gradeEventTypes.values[5]
}

func (o *gradeEventTypes) ParseGradeEventType(name string) (ret *GradeEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*GradeEventType), ok
	}
	return
}


type GroupEventType struct {
	name  string
	ordinal int
}

func (o *GroupEventType) Name() string {
    return o.name
}

func (o *GroupEventType) Ordinal() int {
    return o.ordinal
}

func (o GroupEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *GroupEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := GroupEventTypes().ParseGroupEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GroupEventType %q", lit.Name)
        }
	}
	return
}

func (o GroupEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *GroupEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := GroupEventTypes().ParseGroupEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GroupEventType %q", lit)
        }
    }
    return
}

func (o *GroupEventType) IsCreate() bool {
    return o == _groupEventTypes.Create()
}

func (o *GroupEventType) IsCreated() bool {
    return o == _groupEventTypes.Created()
}

func (o *GroupEventType) IsDelete() bool {
    return o == _groupEventTypes.Delete()
}

func (o *GroupEventType) IsDeleted() bool {
    return o == _groupEventTypes.Deleted()
}

func (o *GroupEventType) IsUpdate() bool {
    return o == _groupEventTypes.Update()
}

func (o *GroupEventType) IsUpdated() bool {
    return o == _groupEventTypes.Updated()
}

type groupEventTypes struct {
	values []*GroupEventType
    literals []enum.Literal
}

var _groupEventTypes = &groupEventTypes{values: []*GroupEventType{
    {name: "Create", ordinal: 0},
    {name: "Created", ordinal: 1},
    {name: "Delete", ordinal: 2},
    {name: "Deleted", ordinal: 3},
    {name: "Update", ordinal: 4},
    {name: "Updated", ordinal: 5}},
}

func GroupEventTypes() *groupEventTypes {
	return _groupEventTypes
}

func (o *groupEventTypes) Values() []*GroupEventType {
	return o.values
}

func (o *groupEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *groupEventTypes) Create() *GroupEventType {
    return _groupEventTypes.values[0]
}

func (o *groupEventTypes) Created() *GroupEventType {
    return _groupEventTypes.values[1]
}

func (o *groupEventTypes) Delete() *GroupEventType {
    return _groupEventTypes.values[2]
}

func (o *groupEventTypes) Deleted() *GroupEventType {
    return _groupEventTypes.values[3]
}

func (o *groupEventTypes) Update() *GroupEventType {
    return _groupEventTypes.values[4]
}

func (o *groupEventTypes) Updated() *GroupEventType {
    return _groupEventTypes.values[5]
}

func (o *groupEventTypes) ParseGroupEventType(name string) (ret *GroupEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*GroupEventType), ok
	}
	return
}


type SchoolApplicationEventType struct {
	name  string
	ordinal int
}

func (o *SchoolApplicationEventType) Name() string {
    return o.name
}

func (o *SchoolApplicationEventType) Ordinal() int {
    return o.ordinal
}

func (o SchoolApplicationEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *SchoolApplicationEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := SchoolApplicationEventTypes().ParseSchoolApplicationEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid SchoolApplicationEventType %q", lit.Name)
        }
	}
	return
}

func (o SchoolApplicationEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *SchoolApplicationEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := SchoolApplicationEventTypes().ParseSchoolApplicationEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid SchoolApplicationEventType %q", lit)
        }
    }
    return
}

func (o *SchoolApplicationEventType) IsCreate() bool {
    return o == _schoolApplicationEventTypes.Create()
}

func (o *SchoolApplicationEventType) IsCreated() bool {
    return o == _schoolApplicationEventTypes.Created()
}

func (o *SchoolApplicationEventType) IsDelete() bool {
    return o == _schoolApplicationEventTypes.Delete()
}

func (o *SchoolApplicationEventType) IsDeleted() bool {
    return o == _schoolApplicationEventTypes.Deleted()
}

func (o *SchoolApplicationEventType) IsUpdate() bool {
    return o == _schoolApplicationEventTypes.Update()
}

func (o *SchoolApplicationEventType) IsUpdated() bool {
    return o == _schoolApplicationEventTypes.Updated()
}

type schoolApplicationEventTypes struct {
	values []*SchoolApplicationEventType
    literals []enum.Literal
}

var _schoolApplicationEventTypes = &schoolApplicationEventTypes{values: []*SchoolApplicationEventType{
    {name: "Create", ordinal: 0},
    {name: "Created", ordinal: 1},
    {name: "Delete", ordinal: 2},
    {name: "Deleted", ordinal: 3},
    {name: "Update", ordinal: 4},
    {name: "Updated", ordinal: 5}},
}

func SchoolApplicationEventTypes() *schoolApplicationEventTypes {
	return _schoolApplicationEventTypes
}

func (o *schoolApplicationEventTypes) Values() []*SchoolApplicationEventType {
	return o.values
}

func (o *schoolApplicationEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *schoolApplicationEventTypes) Create() *SchoolApplicationEventType {
    return _schoolApplicationEventTypes.values[0]
}

func (o *schoolApplicationEventTypes) Created() *SchoolApplicationEventType {
    return _schoolApplicationEventTypes.values[1]
}

func (o *schoolApplicationEventTypes) Delete() *SchoolApplicationEventType {
    return _schoolApplicationEventTypes.values[2]
}

func (o *schoolApplicationEventTypes) Deleted() *SchoolApplicationEventType {
    return _schoolApplicationEventTypes.values[3]
}

func (o *schoolApplicationEventTypes) Update() *SchoolApplicationEventType {
    return _schoolApplicationEventTypes.values[4]
}

func (o *schoolApplicationEventTypes) Updated() *SchoolApplicationEventType {
    return _schoolApplicationEventTypes.values[5]
}

func (o *schoolApplicationEventTypes) ParseSchoolApplicationEventType(name string) (ret *SchoolApplicationEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*SchoolApplicationEventType), ok
	}
	return
}


type SchoolYearEventType struct {
	name  string
	ordinal int
}

func (o *SchoolYearEventType) Name() string {
    return o.name
}

func (o *SchoolYearEventType) Ordinal() int {
    return o.ordinal
}

func (o SchoolYearEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *SchoolYearEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := SchoolYearEventTypes().ParseSchoolYearEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid SchoolYearEventType %q", lit.Name)
        }
	}
	return
}

func (o SchoolYearEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *SchoolYearEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := SchoolYearEventTypes().ParseSchoolYearEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid SchoolYearEventType %q", lit)
        }
    }
    return
}

func (o *SchoolYearEventType) IsCreate() bool {
    return o == _schoolYearEventTypes.Create()
}

func (o *SchoolYearEventType) IsCreated() bool {
    return o == _schoolYearEventTypes.Created()
}

func (o *SchoolYearEventType) IsDelete() bool {
    return o == _schoolYearEventTypes.Delete()
}

func (o *SchoolYearEventType) IsDeleted() bool {
    return o == _schoolYearEventTypes.Deleted()
}

func (o *SchoolYearEventType) IsUpdate() bool {
    return o == _schoolYearEventTypes.Update()
}

func (o *SchoolYearEventType) IsUpdated() bool {
    return o == _schoolYearEventTypes.Updated()
}

type schoolYearEventTypes struct {
	values []*SchoolYearEventType
    literals []enum.Literal
}

var _schoolYearEventTypes = &schoolYearEventTypes{values: []*SchoolYearEventType{
    {name: "Create", ordinal: 0},
    {name: "Created", ordinal: 1},
    {name: "Delete", ordinal: 2},
    {name: "Deleted", ordinal: 3},
    {name: "Update", ordinal: 4},
    {name: "Updated", ordinal: 5}},
}

func SchoolYearEventTypes() *schoolYearEventTypes {
	return _schoolYearEventTypes
}

func (o *schoolYearEventTypes) Values() []*SchoolYearEventType {
	return o.values
}

func (o *schoolYearEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *schoolYearEventTypes) Create() *SchoolYearEventType {
    return _schoolYearEventTypes.values[0]
}

func (o *schoolYearEventTypes) Created() *SchoolYearEventType {
    return _schoolYearEventTypes.values[1]
}

func (o *schoolYearEventTypes) Delete() *SchoolYearEventType {
    return _schoolYearEventTypes.values[2]
}

func (o *schoolYearEventTypes) Deleted() *SchoolYearEventType {
    return _schoolYearEventTypes.values[3]
}

func (o *schoolYearEventTypes) Update() *SchoolYearEventType {
    return _schoolYearEventTypes.values[4]
}

func (o *schoolYearEventTypes) Updated() *SchoolYearEventType {
    return _schoolYearEventTypes.values[5]
}

func (o *schoolYearEventTypes) ParseSchoolYearEventType(name string) (ret *SchoolYearEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*SchoolYearEventType), ok
	}
	return
}



