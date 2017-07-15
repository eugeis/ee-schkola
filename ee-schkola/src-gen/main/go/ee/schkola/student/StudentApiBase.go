package student

import (
    "ee/schkola"
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "time"
)

type Attendance struct {
    Student  *person.Profile
    date *time.Time
    Course  *Course
    hours int
    state *AttendanceState
    stateTrace *schkola.Trace
    token string
    tokenTrace *schkola.Trace
    *schkola.SchkolaBase
}

func NewAttendance(student *person.Profile, date *time.Time, course *Course, hours int, state *AttendanceState, stateTrace *schkola.Trace, 
                token string, tokenTrace *schkola.Trace, SchkolaBase *schkola.SchkolaBase) (ret *Attendance) {
    ret = &Attendance{
        Student : student,
        date: date,
        Course : course,
        hours: hours,
        state: state,
        stateTrace: stateTrace,
        token: token,
        tokenTrace: tokenTrace,
        SchkolaBase: SchkolaBase,
    }
    
    return
}



type Course struct {
    name string
    begin *time.Time
    end *time.Time
    teacher *person.PersonName
    schoolYear *SchoolYear
    fee float64
    description string
    *schkola.SchkolaBase
}

func NewCourse(name string, begin *time.Time, end *time.Time, teacher *person.PersonName, schoolYear *SchoolYear, fee float64, 
                description string, SchkolaBase *schkola.SchkolaBase) (ret *Course) {
    ret = &Course{
        name: name,
        begin: begin,
        end: end,
        teacher: teacher,
        schoolYear: schoolYear,
        fee: fee,
        description: description,
        SchkolaBase: SchkolaBase,
    }
    
    return
}



type Grade struct {
    student *person.Profile
    course *Course
    grade float64
    gradeTrace *schkola.Trace
    comment string
    *schkola.SchkolaBase
}

func NewGrade(student *person.Profile, course *Course, grade float64, gradeTrace *schkola.Trace, comment string, 
                SchkolaBase *schkola.SchkolaBase) (ret *Grade) {
    ret = &Grade{
        student: student,
        course: course,
        grade: grade,
        gradeTrace: gradeTrace,
        comment: comment,
        SchkolaBase: SchkolaBase,
    }
    
    return
}



type Group struct {
    name string
    category *GroupCategory
    schoolYear *SchoolYear
    representative *person.Profile
    students []*Course
    courses []*Course
    *schkola.SchkolaBase
}

func NewGroup(name string, category *GroupCategory, schoolYear *SchoolYear, representative *person.Profile, students []*Course, 
                courses []*Course, SchkolaBase *schkola.SchkolaBase) (ret *Group) {
    ret = &Group{
        name: name,
        category: category,
        schoolYear: schoolYear,
        representative: representative,
        students: students,
        courses: courses,
        SchkolaBase: SchkolaBase,
    }
    
    return
}

func (o *Group) AddToStudents(item *Course) *Course {
    o.students = append(o.students, item)
    return item
}

func (o *Group) AddToCourses(item *Course) *Course {
    o.courses = append(o.courses, item)
    return item
}



type SchoolApplication struct {
    profile *person.Profile
    recommendationOf *person.PersonName
    churchContactPerson *person.PersonName
    churchContact *person.Contact
    schoolYear *SchoolYear
    group string
    *schkola.SchkolaBase
}

func NewSchoolApplication(profile *person.Profile, recommendationOf *person.PersonName, churchContactPerson *person.PersonName, 
                churchContact *person.Contact, schoolYear *SchoolYear, group string, SchkolaBase *schkola.SchkolaBase) (ret *SchoolApplication) {
    ret = &SchoolApplication{
        profile: profile,
        recommendationOf: recommendationOf,
        churchContactPerson: churchContactPerson,
        churchContact: churchContact,
        schoolYear: schoolYear,
        group: group,
        SchkolaBase: SchkolaBase,
    }
    
    return
}



type SchoolYear struct {
    name string
    start *time.Time
    end *time.Time
    dates []*Course
    *schkola.SchkolaBase
}

func NewSchoolYear(name string, start *time.Time, end *time.Time, dates []*Course, SchkolaBase *schkola.SchkolaBase) (ret *SchoolYear) {
    ret = &SchoolYear{
        name: name,
        start: start,
        end: end,
        dates: dates,
        SchkolaBase: SchkolaBase,
    }
    
    return
}

func (o *SchoolYear) AddToDates(item *Course) *Course {
    o.dates = append(o.dates, item)
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



