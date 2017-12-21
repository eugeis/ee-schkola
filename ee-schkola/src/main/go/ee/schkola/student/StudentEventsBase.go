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
     AttendanceRegisterEvent eventhorizon.EventType = "AttendanceRegister"
     AttendanceCreateEvent eventhorizon.EventType = "AttendanceCreate"
     AttendanceCreatedEvent eventhorizon.EventType = "AttendanceCreated"
     AttendanceRegisteredEvent eventhorizon.EventType = "AttendanceRegistered"
     AttendanceDeleteEvent eventhorizon.EventType = "AttendanceDelete"
     AttendanceDeletedEvent eventhorizon.EventType = "AttendanceDeleted"
     AttendanceCancelEvent eventhorizon.EventType = "AttendanceCancel"
     AttendanceConfirmEvent eventhorizon.EventType = "AttendanceConfirm"
     AttendanceUpdateEvent eventhorizon.EventType = "AttendanceUpdate"
     AttendanceUpdatedEvent eventhorizon.EventType = "AttendanceUpdated"
     AttendanceConfirmedEvent eventhorizon.EventType = "AttendanceConfirmed"
     AttendanceCanceledEvent eventhorizon.EventType = "AttendanceCanceled"
)


const (
     CourseCreateEvent eventhorizon.EventType = "CourseCreate"
     CourseCreatedEvent eventhorizon.EventType = "CourseCreated"
     CourseDeleteEvent eventhorizon.EventType = "CourseDelete"
     CourseDeletedEvent eventhorizon.EventType = "CourseDeleted"
     CourseUpdateEvent eventhorizon.EventType = "CourseUpdate"
     CourseUpdatedEvent eventhorizon.EventType = "CourseUpdated"
)


const (
     GradeCreateEvent eventhorizon.EventType = "GradeCreate"
     GradeCreatedEvent eventhorizon.EventType = "GradeCreated"
     GradeDeleteEvent eventhorizon.EventType = "GradeDelete"
     GradeDeletedEvent eventhorizon.EventType = "GradeDeleted"
     GradeUpdateEvent eventhorizon.EventType = "GradeUpdate"
     GradeUpdatedEvent eventhorizon.EventType = "GradeUpdated"
)


const (
     GroupCreateEvent eventhorizon.EventType = "GroupCreate"
     GroupCreatedEvent eventhorizon.EventType = "GroupCreated"
     GroupDeleteEvent eventhorizon.EventType = "GroupDelete"
     GroupDeletedEvent eventhorizon.EventType = "GroupDeleted"
     GroupUpdateEvent eventhorizon.EventType = "GroupUpdate"
     GroupUpdatedEvent eventhorizon.EventType = "GroupUpdated"
)


const (
     SchoolApplicationCreateEvent eventhorizon.EventType = "SchoolApplicationCreate"
     SchoolApplicationCreatedEvent eventhorizon.EventType = "SchoolApplicationCreated"
     SchoolApplicationDeleteEvent eventhorizon.EventType = "SchoolApplicationDelete"
     SchoolApplicationDeletedEvent eventhorizon.EventType = "SchoolApplicationDeleted"
     SchoolApplicationUpdateEvent eventhorizon.EventType = "SchoolApplicationUpdate"
     SchoolApplicationUpdatedEvent eventhorizon.EventType = "SchoolApplicationUpdated"
)


const (
     SchoolYearCreateEvent eventhorizon.EventType = "SchoolYearCreate"
     SchoolYearCreatedEvent eventhorizon.EventType = "SchoolYearCreated"
     SchoolYearDeleteEvent eventhorizon.EventType = "SchoolYearDelete"
     SchoolYearDeletedEvent eventhorizon.EventType = "SchoolYearDeleted"
     SchoolYearUpdateEvent eventhorizon.EventType = "SchoolYearUpdate"
     SchoolYearUpdatedEvent eventhorizon.EventType = "SchoolYearUpdated"
)




type RegisterAttendance struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type CreateAttendance struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Hours int `json:"hours" eh:"optional"`
    State *AttendanceState `json:"state" eh:"optional"`
    Token string `json:"token" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type AttendanceCreated struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Hours int `json:"hours" eh:"optional"`
    State *AttendanceState `json:"state" eh:"optional"`
    Token string `json:"token" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type AttendanceRegistered struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type DeleteAttendance struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type AttendanceDeleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type CancelAttendance struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type ConfirmAttendance struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type UpdateAttendance struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Hours int `json:"hours" eh:"optional"`
    State *AttendanceState `json:"state" eh:"optional"`
    Token string `json:"token" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type AttendanceUpdated struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Hours int `json:"hours" eh:"optional"`
    State *AttendanceState `json:"state" eh:"optional"`
    Token string `json:"token" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type AttendanceConfirmed struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type AttendanceCanceled struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type CreateCourse struct {
    Name string `json:"name" eh:"optional"`
    Begin *time.Time `json:"begin" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Teacher *shared.PersonName `json:"teacher" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Fee float64 `json:"fee" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type CourseCreated struct {
    Name string `json:"name" eh:"optional"`
    Begin *time.Time `json:"begin" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Teacher *shared.PersonName `json:"teacher" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Fee float64 `json:"fee" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type DeleteCourse struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type CourseDeleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type UpdateCourse struct {
    Name string `json:"name" eh:"optional"`
    Begin *time.Time `json:"begin" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Teacher *shared.PersonName `json:"teacher" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Fee float64 `json:"fee" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type CourseUpdated struct {
    Name string `json:"name" eh:"optional"`
    Begin *time.Time `json:"begin" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Teacher *shared.PersonName `json:"teacher" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Fee float64 `json:"fee" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type CreateGrade struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Grade float64 `json:"grade" eh:"optional"`
    Comment string `json:"comment" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type GradeCreated struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Grade float64 `json:"grade" eh:"optional"`
    Comment string `json:"comment" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type DeleteGrade struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type GradeDeleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type UpdateGrade struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Grade float64 `json:"grade" eh:"optional"`
    Comment string `json:"comment" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type GradeUpdated struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Grade float64 `json:"grade" eh:"optional"`
    Comment string `json:"comment" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type CreateGroup struct {
    Name string `json:"name" eh:"optional"`
    Category *GroupCategory `json:"category" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Representative *person.Profile `json:"representative" eh:"optional"`
    Students []*person.Profile `json:"students" eh:"optional"`
    Courses []*Course `json:"courses" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *CreateGroup) AddToStudents(item *person.Profile) *person.Profile {
    o.Students = append(o.Students, item)
    return item
}

func (o *CreateGroup) AddToCourses(item *Course) *Course {
    o.Courses = append(o.Courses, item)
    return item
}


type GroupCreated struct {
    Name string `json:"name" eh:"optional"`
    Category *GroupCategory `json:"category" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Representative *person.Profile `json:"representative" eh:"optional"`
    Students []*person.Profile `json:"students" eh:"optional"`
    Courses []*Course `json:"courses" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *GroupCreated) AddToStudents(item *person.Profile) *person.Profile {
    o.Students = append(o.Students, item)
    return item
}

func (o *GroupCreated) AddToCourses(item *Course) *Course {
    o.Courses = append(o.Courses, item)
    return item
}


type DeleteGroup struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type GroupDeleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type UpdateGroup struct {
    Name string `json:"name" eh:"optional"`
    Category *GroupCategory `json:"category" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Representative *person.Profile `json:"representative" eh:"optional"`
    Students []*person.Profile `json:"students" eh:"optional"`
    Courses []*Course `json:"courses" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *UpdateGroup) AddToStudents(item *person.Profile) *person.Profile {
    o.Students = append(o.Students, item)
    return item
}

func (o *UpdateGroup) AddToCourses(item *Course) *Course {
    o.Courses = append(o.Courses, item)
    return item
}


type GroupUpdated struct {
    Name string `json:"name" eh:"optional"`
    Category *GroupCategory `json:"category" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Representative *person.Profile `json:"representative" eh:"optional"`
    Students []*person.Profile `json:"students" eh:"optional"`
    Courses []*Course `json:"courses" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *GroupUpdated) AddToStudents(item *person.Profile) *person.Profile {
    o.Students = append(o.Students, item)
    return item
}

func (o *GroupUpdated) AddToCourses(item *Course) *Course {
    o.Courses = append(o.Courses, item)
    return item
}


type CreateSchoolApplication struct {
    Profile *person.Profile `json:"profile" eh:"optional"`
    ChurchContactPerson *shared.PersonName `json:"churchContactPerson" eh:"optional"`
    ChurchContact *person.Contact `json:"churchContact" eh:"optional"`
    ChurchCommitment bool `json:"churchCommitment" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Group string `json:"group" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type SchoolApplicationCreated struct {
    Profile *person.Profile `json:"profile" eh:"optional"`
    ChurchContactPerson *shared.PersonName `json:"churchContactPerson" eh:"optional"`
    ChurchContact *person.Contact `json:"churchContact" eh:"optional"`
    ChurchCommitment bool `json:"churchCommitment" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Group string `json:"group" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type DeleteSchoolApplication struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type SchoolApplicationDeleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type UpdateSchoolApplication struct {
    Profile *person.Profile `json:"profile" eh:"optional"`
    ChurchContactPerson *shared.PersonName `json:"churchContactPerson" eh:"optional"`
    ChurchContact *person.Contact `json:"churchContact" eh:"optional"`
    ChurchCommitment bool `json:"churchCommitment" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Group string `json:"group" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type SchoolApplicationUpdated struct {
    Profile *person.Profile `json:"profile" eh:"optional"`
    ChurchContactPerson *shared.PersonName `json:"churchContactPerson" eh:"optional"`
    ChurchContact *person.Contact `json:"churchContact" eh:"optional"`
    ChurchCommitment bool `json:"churchCommitment" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Group string `json:"group" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type CreateSchoolYear struct {
    Name string `json:"name" eh:"optional"`
    Start *time.Time `json:"start" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Dates []*time.Time `json:"dates" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *CreateSchoolYear) AddToDates(item *time.Time) *time.Time {
    o.Dates = append(o.Dates, item)
    return item
}


type SchoolYearCreated struct {
    Name string `json:"name" eh:"optional"`
    Start *time.Time `json:"start" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Dates []*time.Time `json:"dates" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *SchoolYearCreated) AddToDates(item *time.Time) *time.Time {
    o.Dates = append(o.Dates, item)
    return item
}


type DeleteSchoolYear struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type SchoolYearDeleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type UpdateSchoolYear struct {
    Name string `json:"name" eh:"optional"`
    Start *time.Time `json:"start" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Dates []*time.Time `json:"dates" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *UpdateSchoolYear) AddToDates(item *time.Time) *time.Time {
    o.Dates = append(o.Dates, item)
    return item
}


type SchoolYearUpdated struct {
    Name string `json:"name" eh:"optional"`
    Start *time.Time `json:"start" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Dates []*time.Time `json:"dates" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *SchoolYearUpdated) AddToDates(item *time.Time) *time.Time {
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

func (o *AttendanceEventType) IsAttendanceRegister() bool {
    return o == _attendanceEventTypes.AttendanceRegister()
}

func (o *AttendanceEventType) IsAttendanceCreate() bool {
    return o == _attendanceEventTypes.AttendanceCreate()
}

func (o *AttendanceEventType) IsAttendanceCreated() bool {
    return o == _attendanceEventTypes.AttendanceCreated()
}

func (o *AttendanceEventType) IsAttendanceRegistered() bool {
    return o == _attendanceEventTypes.AttendanceRegistered()
}

func (o *AttendanceEventType) IsAttendanceDelete() bool {
    return o == _attendanceEventTypes.AttendanceDelete()
}

func (o *AttendanceEventType) IsAttendanceDeleted() bool {
    return o == _attendanceEventTypes.AttendanceDeleted()
}

func (o *AttendanceEventType) IsAttendanceCancel() bool {
    return o == _attendanceEventTypes.AttendanceCancel()
}

func (o *AttendanceEventType) IsAttendanceConfirm() bool {
    return o == _attendanceEventTypes.AttendanceConfirm()
}

func (o *AttendanceEventType) IsAttendanceUpdate() bool {
    return o == _attendanceEventTypes.AttendanceUpdate()
}

func (o *AttendanceEventType) IsAttendanceUpdated() bool {
    return o == _attendanceEventTypes.AttendanceUpdated()
}

func (o *AttendanceEventType) IsAttendanceConfirmed() bool {
    return o == _attendanceEventTypes.AttendanceConfirmed()
}

func (o *AttendanceEventType) IsAttendanceCanceled() bool {
    return o == _attendanceEventTypes.AttendanceCanceled()
}

type attendanceEventTypes struct {
	values []*AttendanceEventType
    literals []enum.Literal
}

var _attendanceEventTypes = &attendanceEventTypes{values: []*AttendanceEventType{
    {name: "AttendanceRegister", ordinal: 0},
    {name: "AttendanceCreate", ordinal: 1},
    {name: "AttendanceCreated", ordinal: 2},
    {name: "AttendanceRegistered", ordinal: 3},
    {name: "AttendanceDelete", ordinal: 4},
    {name: "AttendanceDeleted", ordinal: 5},
    {name: "AttendanceCancel", ordinal: 6},
    {name: "AttendanceConfirm", ordinal: 7},
    {name: "AttendanceUpdate", ordinal: 8},
    {name: "AttendanceUpdated", ordinal: 9},
    {name: "AttendanceConfirmed", ordinal: 10},
    {name: "AttendanceCanceled", ordinal: 11}},
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

func (o *attendanceEventTypes) AttendanceRegister() *AttendanceEventType {
    return _attendanceEventTypes.values[0]
}

func (o *attendanceEventTypes) AttendanceCreate() *AttendanceEventType {
    return _attendanceEventTypes.values[1]
}

func (o *attendanceEventTypes) AttendanceCreated() *AttendanceEventType {
    return _attendanceEventTypes.values[2]
}

func (o *attendanceEventTypes) AttendanceRegistered() *AttendanceEventType {
    return _attendanceEventTypes.values[3]
}

func (o *attendanceEventTypes) AttendanceDelete() *AttendanceEventType {
    return _attendanceEventTypes.values[4]
}

func (o *attendanceEventTypes) AttendanceDeleted() *AttendanceEventType {
    return _attendanceEventTypes.values[5]
}

func (o *attendanceEventTypes) AttendanceCancel() *AttendanceEventType {
    return _attendanceEventTypes.values[6]
}

func (o *attendanceEventTypes) AttendanceConfirm() *AttendanceEventType {
    return _attendanceEventTypes.values[7]
}

func (o *attendanceEventTypes) AttendanceUpdate() *AttendanceEventType {
    return _attendanceEventTypes.values[8]
}

func (o *attendanceEventTypes) AttendanceUpdated() *AttendanceEventType {
    return _attendanceEventTypes.values[9]
}

func (o *attendanceEventTypes) AttendanceConfirmed() *AttendanceEventType {
    return _attendanceEventTypes.values[10]
}

func (o *attendanceEventTypes) AttendanceCanceled() *AttendanceEventType {
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

func (o *CourseEventType) IsCourseCreate() bool {
    return o == _courseEventTypes.CourseCreate()
}

func (o *CourseEventType) IsCourseCreated() bool {
    return o == _courseEventTypes.CourseCreated()
}

func (o *CourseEventType) IsCourseDelete() bool {
    return o == _courseEventTypes.CourseDelete()
}

func (o *CourseEventType) IsCourseDeleted() bool {
    return o == _courseEventTypes.CourseDeleted()
}

func (o *CourseEventType) IsCourseUpdate() bool {
    return o == _courseEventTypes.CourseUpdate()
}

func (o *CourseEventType) IsCourseUpdated() bool {
    return o == _courseEventTypes.CourseUpdated()
}

type courseEventTypes struct {
	values []*CourseEventType
    literals []enum.Literal
}

var _courseEventTypes = &courseEventTypes{values: []*CourseEventType{
    {name: "CourseCreate", ordinal: 0},
    {name: "CourseCreated", ordinal: 1},
    {name: "CourseDelete", ordinal: 2},
    {name: "CourseDeleted", ordinal: 3},
    {name: "CourseUpdate", ordinal: 4},
    {name: "CourseUpdated", ordinal: 5}},
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

func (o *courseEventTypes) CourseCreate() *CourseEventType {
    return _courseEventTypes.values[0]
}

func (o *courseEventTypes) CourseCreated() *CourseEventType {
    return _courseEventTypes.values[1]
}

func (o *courseEventTypes) CourseDelete() *CourseEventType {
    return _courseEventTypes.values[2]
}

func (o *courseEventTypes) CourseDeleted() *CourseEventType {
    return _courseEventTypes.values[3]
}

func (o *courseEventTypes) CourseUpdate() *CourseEventType {
    return _courseEventTypes.values[4]
}

func (o *courseEventTypes) CourseUpdated() *CourseEventType {
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

func (o *GradeEventType) IsGradeCreate() bool {
    return o == _gradeEventTypes.GradeCreate()
}

func (o *GradeEventType) IsGradeCreated() bool {
    return o == _gradeEventTypes.GradeCreated()
}

func (o *GradeEventType) IsGradeDelete() bool {
    return o == _gradeEventTypes.GradeDelete()
}

func (o *GradeEventType) IsGradeDeleted() bool {
    return o == _gradeEventTypes.GradeDeleted()
}

func (o *GradeEventType) IsGradeUpdate() bool {
    return o == _gradeEventTypes.GradeUpdate()
}

func (o *GradeEventType) IsGradeUpdated() bool {
    return o == _gradeEventTypes.GradeUpdated()
}

type gradeEventTypes struct {
	values []*GradeEventType
    literals []enum.Literal
}

var _gradeEventTypes = &gradeEventTypes{values: []*GradeEventType{
    {name: "GradeCreate", ordinal: 0},
    {name: "GradeCreated", ordinal: 1},
    {name: "GradeDelete", ordinal: 2},
    {name: "GradeDeleted", ordinal: 3},
    {name: "GradeUpdate", ordinal: 4},
    {name: "GradeUpdated", ordinal: 5}},
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

func (o *gradeEventTypes) GradeCreate() *GradeEventType {
    return _gradeEventTypes.values[0]
}

func (o *gradeEventTypes) GradeCreated() *GradeEventType {
    return _gradeEventTypes.values[1]
}

func (o *gradeEventTypes) GradeDelete() *GradeEventType {
    return _gradeEventTypes.values[2]
}

func (o *gradeEventTypes) GradeDeleted() *GradeEventType {
    return _gradeEventTypes.values[3]
}

func (o *gradeEventTypes) GradeUpdate() *GradeEventType {
    return _gradeEventTypes.values[4]
}

func (o *gradeEventTypes) GradeUpdated() *GradeEventType {
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

func (o *GroupEventType) IsGroupCreate() bool {
    return o == _groupEventTypes.GroupCreate()
}

func (o *GroupEventType) IsGroupCreated() bool {
    return o == _groupEventTypes.GroupCreated()
}

func (o *GroupEventType) IsGroupDelete() bool {
    return o == _groupEventTypes.GroupDelete()
}

func (o *GroupEventType) IsGroupDeleted() bool {
    return o == _groupEventTypes.GroupDeleted()
}

func (o *GroupEventType) IsGroupUpdate() bool {
    return o == _groupEventTypes.GroupUpdate()
}

func (o *GroupEventType) IsGroupUpdated() bool {
    return o == _groupEventTypes.GroupUpdated()
}

type groupEventTypes struct {
	values []*GroupEventType
    literals []enum.Literal
}

var _groupEventTypes = &groupEventTypes{values: []*GroupEventType{
    {name: "GroupCreate", ordinal: 0},
    {name: "GroupCreated", ordinal: 1},
    {name: "GroupDelete", ordinal: 2},
    {name: "GroupDeleted", ordinal: 3},
    {name: "GroupUpdate", ordinal: 4},
    {name: "GroupUpdated", ordinal: 5}},
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

func (o *groupEventTypes) GroupCreate() *GroupEventType {
    return _groupEventTypes.values[0]
}

func (o *groupEventTypes) GroupCreated() *GroupEventType {
    return _groupEventTypes.values[1]
}

func (o *groupEventTypes) GroupDelete() *GroupEventType {
    return _groupEventTypes.values[2]
}

func (o *groupEventTypes) GroupDeleted() *GroupEventType {
    return _groupEventTypes.values[3]
}

func (o *groupEventTypes) GroupUpdate() *GroupEventType {
    return _groupEventTypes.values[4]
}

func (o *groupEventTypes) GroupUpdated() *GroupEventType {
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

func (o *SchoolApplicationEventType) IsSchoolApplicationCreate() bool {
    return o == _schoolApplicationEventTypes.SchoolApplicationCreate()
}

func (o *SchoolApplicationEventType) IsSchoolApplicationCreated() bool {
    return o == _schoolApplicationEventTypes.SchoolApplicationCreated()
}

func (o *SchoolApplicationEventType) IsSchoolApplicationDelete() bool {
    return o == _schoolApplicationEventTypes.SchoolApplicationDelete()
}

func (o *SchoolApplicationEventType) IsSchoolApplicationDeleted() bool {
    return o == _schoolApplicationEventTypes.SchoolApplicationDeleted()
}

func (o *SchoolApplicationEventType) IsSchoolApplicationUpdate() bool {
    return o == _schoolApplicationEventTypes.SchoolApplicationUpdate()
}

func (o *SchoolApplicationEventType) IsSchoolApplicationUpdated() bool {
    return o == _schoolApplicationEventTypes.SchoolApplicationUpdated()
}

type schoolApplicationEventTypes struct {
	values []*SchoolApplicationEventType
    literals []enum.Literal
}

var _schoolApplicationEventTypes = &schoolApplicationEventTypes{values: []*SchoolApplicationEventType{
    {name: "SchoolApplicationCreate", ordinal: 0},
    {name: "SchoolApplicationCreated", ordinal: 1},
    {name: "SchoolApplicationDelete", ordinal: 2},
    {name: "SchoolApplicationDeleted", ordinal: 3},
    {name: "SchoolApplicationUpdate", ordinal: 4},
    {name: "SchoolApplicationUpdated", ordinal: 5}},
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

func (o *schoolApplicationEventTypes) SchoolApplicationCreate() *SchoolApplicationEventType {
    return _schoolApplicationEventTypes.values[0]
}

func (o *schoolApplicationEventTypes) SchoolApplicationCreated() *SchoolApplicationEventType {
    return _schoolApplicationEventTypes.values[1]
}

func (o *schoolApplicationEventTypes) SchoolApplicationDelete() *SchoolApplicationEventType {
    return _schoolApplicationEventTypes.values[2]
}

func (o *schoolApplicationEventTypes) SchoolApplicationDeleted() *SchoolApplicationEventType {
    return _schoolApplicationEventTypes.values[3]
}

func (o *schoolApplicationEventTypes) SchoolApplicationUpdate() *SchoolApplicationEventType {
    return _schoolApplicationEventTypes.values[4]
}

func (o *schoolApplicationEventTypes) SchoolApplicationUpdated() *SchoolApplicationEventType {
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

func (o *SchoolYearEventType) IsSchoolYearCreate() bool {
    return o == _schoolYearEventTypes.SchoolYearCreate()
}

func (o *SchoolYearEventType) IsSchoolYearCreated() bool {
    return o == _schoolYearEventTypes.SchoolYearCreated()
}

func (o *SchoolYearEventType) IsSchoolYearDelete() bool {
    return o == _schoolYearEventTypes.SchoolYearDelete()
}

func (o *SchoolYearEventType) IsSchoolYearDeleted() bool {
    return o == _schoolYearEventTypes.SchoolYearDeleted()
}

func (o *SchoolYearEventType) IsSchoolYearUpdate() bool {
    return o == _schoolYearEventTypes.SchoolYearUpdate()
}

func (o *SchoolYearEventType) IsSchoolYearUpdated() bool {
    return o == _schoolYearEventTypes.SchoolYearUpdated()
}

type schoolYearEventTypes struct {
	values []*SchoolYearEventType
    literals []enum.Literal
}

var _schoolYearEventTypes = &schoolYearEventTypes{values: []*SchoolYearEventType{
    {name: "SchoolYearCreate", ordinal: 0},
    {name: "SchoolYearCreated", ordinal: 1},
    {name: "SchoolYearDelete", ordinal: 2},
    {name: "SchoolYearDeleted", ordinal: 3},
    {name: "SchoolYearUpdate", ordinal: 4},
    {name: "SchoolYearUpdated", ordinal: 5}},
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

func (o *schoolYearEventTypes) SchoolYearCreate() *SchoolYearEventType {
    return _schoolYearEventTypes.values[0]
}

func (o *schoolYearEventTypes) SchoolYearCreated() *SchoolYearEventType {
    return _schoolYearEventTypes.values[1]
}

func (o *schoolYearEventTypes) SchoolYearDelete() *SchoolYearEventType {
    return _schoolYearEventTypes.values[2]
}

func (o *schoolYearEventTypes) SchoolYearDeleted() *SchoolYearEventType {
    return _schoolYearEventTypes.values[3]
}

func (o *schoolYearEventTypes) SchoolYearUpdate() *SchoolYearEventType {
    return _schoolYearEventTypes.values[4]
}

func (o *schoolYearEventTypes) SchoolYearUpdated() *SchoolYearEventType {
    return _schoolYearEventTypes.values[5]
}

func (o *schoolYearEventTypes) ParseSchoolYearEventType(name string) (ret *SchoolYearEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*SchoolYearEventType), ok
	}
	return
}



