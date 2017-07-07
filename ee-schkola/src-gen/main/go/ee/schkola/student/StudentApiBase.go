package student

import (
    "ee/schkola"
    "ee/schkola/person"
    "time"
)
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
    switch name {
      case "Registered":
        ret = o.Registered()
      case "Confirmed":
        ret = o.Confirmed()
      case "Canceled":
        ret = o.Canceled()
      case "Present":
        ret = o.Present()
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

func (o *groupCategorys) CourseGroup() *GroupCategory {
    return _groupCategorys.values[0]
}

func (o *groupCategorys) YearGroup() *GroupCategory {
    return _groupCategorys.values[1]
}

func (o *groupCategorys) ParseGroupCategory(name string) (ret *GroupCategory, ok bool) {
    switch name {
      case "CourseGroup":
        ret = o.CourseGroup()
      case "YearGroup":
        ret = o.YearGroup()
    }
    return
}




type Attendance struct {
    Student  *person.Profile
    Date  *time.Time
    Course  *Course
    Hours  int
    State  *AttendanceState
    StateTrace  *schkola.Trace
    Token  string
    TokenTrace  *schkola.Trace
    *schkola.SchkolaBase
}

func NewAttendance(student *person.Profile, date *time.Time, course *Course, hours int, state *AttendanceState, stateTrace *schkola.Trace, 
                token string, tokenTrace *schkola.Trace, SchkolaBase *schkola.SchkolaBase) (ret *Attendance, err error) {
    ret = &Attendance{
        Student : student,
        Date : date,
        Course : course,
        Hours : hours,
        State : state,
        StateTrace : stateTrace,
        Token : token,
        TokenTrace : tokenTrace,
        SchkolaBase : SchkolaBase,
    }
    return
}


type Course struct {
    Name  string
    Begin  *time.Time
    End  *time.Time
    Teacher  *person.PersonName
    SchoolYear  *SchoolYear
    Fee  float64
    Description  string
    *schkola.SchkolaBase
}

func NewCourse(name string, begin *time.Time, end *time.Time, teacher *person.PersonName, schoolYear *SchoolYear, fee float64, 
                description string, SchkolaBase *schkola.SchkolaBase) (ret *Course, err error) {
    ret = &Course{
        Name : name,
        Begin : begin,
        End : end,
        Teacher : teacher,
        SchoolYear : schoolYear,
        Fee : fee,
        Description : description,
        SchkolaBase : SchkolaBase,
    }
    return
}


type Grade struct {
    Student  *person.Profile
    Course  *Course
    Grade  float64
    GradeTrace  *schkola.Trace
    Comment  string
    *schkola.SchkolaBase
}

func NewGrade(student *person.Profile, course *Course, grade float64, gradeTrace *schkola.Trace, comment string, 
                SchkolaBase *schkola.SchkolaBase) (ret *Grade, err error) {
    ret = &Grade{
        Student : student,
        Course : course,
        Grade : grade,
        GradeTrace : gradeTrace,
        Comment : comment,
        SchkolaBase : SchkolaBase,
    }
    return
}


type Group struct {
    Name  string
    Category  *GroupCategory
    SchoolYear  *SchoolYear
    Representative  *person.Profile
    Students  []*person.Profile
    Courses  []*Course
    *schkola.SchkolaBase
}

func NewGroup(name string, category *GroupCategory, schoolYear *SchoolYear, representative *person.Profile, students []*person.Profile, 
                courses []*Course, SchkolaBase *schkola.SchkolaBase) (ret *Group, err error) {
    ret = &Group{
        Name : name,
        Category : category,
        SchoolYear : schoolYear,
        Representative : representative,
        Students : students,
        Courses : courses,
        SchkolaBase : SchkolaBase,
    }
    return
}

func (o *Group) AddToStudents(item *person.Profile) *person.Profile {
    o.Students  = append(o.Students , item)
    return item
}

func (o *Group) AddToCourses(item *Course) *Course {
    o.Courses  = append(o.Courses , item)
    return item
}


type SchoolApplication struct {
    Profile  *person.Profile
    RecommendationOf  *person.PersonName
    ChurchContactPerson  *person.PersonName
    ChurchContact  *person.Contact
    SchoolYear  *SchoolYear
    Group  string
    *schkola.SchkolaBase
}

func NewSchoolApplication(profile *person.Profile, recommendationOf *person.PersonName, churchContactPerson *person.PersonName, 
                churchContact *person.Contact, schoolYear *SchoolYear, group string, SchkolaBase *schkola.SchkolaBase) (ret *SchoolApplication, err error) {
    ret = &SchoolApplication{
        Profile : profile,
        RecommendationOf : recommendationOf,
        ChurchContactPerson : churchContactPerson,
        ChurchContact : churchContact,
        SchoolYear : schoolYear,
        Group : group,
        SchkolaBase : SchkolaBase,
    }
    return
}


type SchoolYear struct {
    Name  string
    Start  *time.Time
    End  *time.Time
    Dates  []*time.Time
    *schkola.SchkolaBase
}

func NewSchoolYear(name string, start *time.Time, end *time.Time, dates []*time.Time, SchkolaBase *schkola.SchkolaBase) (ret *SchoolYear, err error) {
    ret = &SchoolYear{
        Name : name,
        Start : start,
        End : end,
        Dates : dates,
        SchkolaBase : SchkolaBase,
    }
    return
}

func (o *SchoolYear) AddToDates(item *time.Time) *time.Time {
    o.Dates  = append(o.Dates , item)
    return item
}



