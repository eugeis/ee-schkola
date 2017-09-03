package student

import (
    "ee/schkola"
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "github.com/looplab/eventhorizon"
    "time"
)
type Attendance struct {
    Student *person.Profile
    Date *time.Time
    Course *Course
    Hours int
    State *AttendanceState
    Token string
    Id eventhorizon.UUID
}

func NewAttendance() (ret *Attendance) {
    ret = &Attendance{}
    return
}


type Course struct {
    Name string
    Begin *time.Time
    End *time.Time
    Teacher *schkola.PersonName
    SchoolYear *SchoolYear
    Fee float64
    Description string
    Id eventhorizon.UUID
}

func NewCourse() (ret *Course) {
    ret = &Course{}
    return
}


type Grade struct {
    Student *person.Profile
    Course *Course
    Grade float64
    Comment string
    Id eventhorizon.UUID
}

func NewGrade() (ret *Grade) {
    ret = &Grade{}
    return
}


type Group struct {
    Name string
    Category *GroupCategory
    SchoolYear *SchoolYear
    Representative *person.Profile
    Students []*person.Profile
    Courses []*Course
    Id eventhorizon.UUID
}

func NewGroup() (ret *Group) {
    ret = &Group{}
    return
}

func (o *Group) AddToStudents(item *person.Profile) *person.Profile {
    o.Students = append(o.Students, item)
    return item
}

func (o *Group) AddToCourses(item *Course) *Course {
    o.Courses = append(o.Courses, item)
    return item
}


type SchoolApplication struct {
    Profile *person.Profile
    RecommendationOf *schkola.PersonName
    ChurchContactPerson *schkola.PersonName
    ChurchContact *person.Contact
    SchoolYear *SchoolYear
    Group string
    Id eventhorizon.UUID
}

func NewSchoolApplication() (ret *SchoolApplication) {
    ret = &SchoolApplication{}
    return
}


type SchoolYear struct {
    Name string
    Start *time.Time
    End *time.Time
    Dates []*time.Time
    Id eventhorizon.UUID
}

func NewSchoolYear() (ret *SchoolYear) {
    ret = &SchoolYear{}
    return
}

func (o *SchoolYear) AddToDates(item *time.Time) *time.Time {
    o.Dates = append(o.Dates, item)
    return item
}








type AttendanceState struct {
	name  string
	ordinal int
}

func (o *AttendanceState) Name() string {
    return o.name
}

func (o *AttendanceState) Ordinal() int {
    return o.ordinal
}

func (o *AttendanceState) IsRegistered() bool {
    return o == _attendanceStates.Registered()
}

func (o *AttendanceState) IsConfirmed() bool {
    return o == _attendanceStates.Confirmed()
}

func (o *AttendanceState) IsCanceled() bool {
    return o == _attendanceStates.Canceled()
}

func (o *AttendanceState) IsPresent() bool {
    return o == _attendanceStates.Present()
}

type attendanceStates struct {
	values []*AttendanceState
    literals []enum.Literal
}

var _attendanceStates = &attendanceStates{values: []*AttendanceState{
    {name: "Registered", ordinal: 0},
    {name: "Confirmed", ordinal: 1},
    {name: "Canceled", ordinal: 2},
    {name: "Present", ordinal: 3}},
}

func AttendanceStates() *attendanceStates {
	return _attendanceStates
}

func (o *attendanceStates) Values() []*AttendanceState {
	return o.values
}

func (o *attendanceStates) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *attendanceStates) Registered() *AttendanceState {
    return _attendanceStates.values[0]
}

func (o *attendanceStates) Confirmed() *AttendanceState {
    return _attendanceStates.values[1]
}

func (o *attendanceStates) Canceled() *AttendanceState {
    return _attendanceStates.values[2]
}

func (o *attendanceStates) Present() *AttendanceState {
    return _attendanceStates.values[3]
}

func (o *attendanceStates) ParseAttendanceState(name string) (ret *AttendanceState, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*AttendanceState), ok
	}
	return
}


type GroupCategory struct {
	name  string
	ordinal int
}

func (o *GroupCategory) Name() string {
    return o.name
}

func (o *GroupCategory) Ordinal() int {
    return o.ordinal
}

func (o *GroupCategory) IsCourseGroup() bool {
    return o == _groupCategorys.CourseGroup()
}

func (o *GroupCategory) IsYearGroup() bool {
    return o == _groupCategorys.YearGroup()
}

type groupCategorys struct {
	values []*GroupCategory
    literals []enum.Literal
}

var _groupCategorys = &groupCategorys{values: []*GroupCategory{
    {name: "CourseGroup", ordinal: 0},
    {name: "YearGroup", ordinal: 1}},
}

func GroupCategorys() *groupCategorys {
	return _groupCategorys
}

func (o *groupCategorys) Values() []*GroupCategory {
	return o.values
}

func (o *groupCategorys) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *groupCategorys) CourseGroup() *GroupCategory {
    return _groupCategorys.values[0]
}

func (o *groupCategorys) YearGroup() *GroupCategory {
    return _groupCategorys.values[1]
}

func (o *groupCategorys) ParseGroupCategory(name string) (ret *GroupCategory, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*GroupCategory), ok
	}
	return
}



