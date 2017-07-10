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


type AttendanceCommands struct {
	name  string
	ordinal int
}

func (o *AttendanceCommands) Name() string {
    return o.name
}

func (o *AttendanceCommands) Ordinal() int {
    return o.ordinal
}

func (o *AttendanceCommands) IsConfirmAttendance() bool {
    return o == _attendanceCommandss.ConfirmAttendance()
}

func (o *AttendanceCommands) IsCancelAttendance() bool {
    return o == _attendanceCommandss.CancelAttendance()
}

func (o *AttendanceCommands) IsRegisterAttendance() bool {
    return o == _attendanceCommandss.RegisterAttendance()
}

type attendanceCommandss struct {
	values []*AttendanceCommands
}

var _attendanceCommandss = &attendanceCommandss{values: []*AttendanceCommands{
    {name: "confirmAttendance", ordinal: 0},
    {name: "cancelAttendance", ordinal: 1},
    {name: "registerAttendance", ordinal: 2}},
}

func AttendanceCommandss() *attendanceCommandss {
	return _attendanceCommandss
}

func (o *attendanceCommandss) Values() []*AttendanceCommands {
	return o.values
}

func (o *attendanceCommandss) ConfirmAttendance() *AttendanceCommands {
    return _attendanceCommandss.values[0]
}

func (o *attendanceCommandss) CancelAttendance() *AttendanceCommands {
    return _attendanceCommandss.values[1]
}

func (o *attendanceCommandss) RegisterAttendance() *AttendanceCommands {
    return _attendanceCommandss.values[2]
}

func (o *attendanceCommandss) ParseAttendanceCommands(name string) (ret *AttendanceCommands, ok bool) {
    switch name {
      case "ConfirmAttendance":
        ret = o.ConfirmAttendance()
      case "CancelAttendance":
        ret = o.CancelAttendance()
      case "RegisterAttendance":
        ret = o.RegisterAttendance()
    }
    return
}


type CourseCommands struct {
	name  string
	ordinal int
}

func (o *CourseCommands) Name() string {
    return o.name
}

func (o *CourseCommands) Ordinal() int {
    return o.ordinal
}

func (o *CourseCommands) IsRegisterCourse() bool {
    return o == _courseCommandss.RegisterCourse()
}

func (o *CourseCommands) IsDeleteCourse() bool {
    return o == _courseCommandss.DeleteCourse()
}

func (o *CourseCommands) IsChangeCourse() bool {
    return o == _courseCommandss.ChangeCourse()
}

type courseCommandss struct {
	values []*CourseCommands
}

var _courseCommandss = &courseCommandss{values: []*CourseCommands{
    {name: "registerCourse", ordinal: 0},
    {name: "deleteCourse", ordinal: 1},
    {name: "changeCourse", ordinal: 2}},
}

func CourseCommandss() *courseCommandss {
	return _courseCommandss
}

func (o *courseCommandss) Values() []*CourseCommands {
	return o.values
}

func (o *courseCommandss) RegisterCourse() *CourseCommands {
    return _courseCommandss.values[0]
}

func (o *courseCommandss) DeleteCourse() *CourseCommands {
    return _courseCommandss.values[1]
}

func (o *courseCommandss) ChangeCourse() *CourseCommands {
    return _courseCommandss.values[2]
}

func (o *courseCommandss) ParseCourseCommands(name string) (ret *CourseCommands, ok bool) {
    switch name {
      case "RegisterCourse":
        ret = o.RegisterCourse()
      case "DeleteCourse":
        ret = o.DeleteCourse()
      case "ChangeCourse":
        ret = o.ChangeCourse()
    }
    return
}


type GradeCommands struct {
	name  string
	ordinal int
}

func (o *GradeCommands) Name() string {
    return o.name
}

func (o *GradeCommands) Ordinal() int {
    return o.ordinal
}

func (o *GradeCommands) IsRegisterGrade() bool {
    return o == _gradeCommandss.RegisterGrade()
}

func (o *GradeCommands) IsDeleteGrade() bool {
    return o == _gradeCommandss.DeleteGrade()
}

func (o *GradeCommands) IsChangeGrade() bool {
    return o == _gradeCommandss.ChangeGrade()
}

type gradeCommandss struct {
	values []*GradeCommands
}

var _gradeCommandss = &gradeCommandss{values: []*GradeCommands{
    {name: "registerGrade", ordinal: 0},
    {name: "deleteGrade", ordinal: 1},
    {name: "changeGrade", ordinal: 2}},
}

func GradeCommandss() *gradeCommandss {
	return _gradeCommandss
}

func (o *gradeCommandss) Values() []*GradeCommands {
	return o.values
}

func (o *gradeCommandss) RegisterGrade() *GradeCommands {
    return _gradeCommandss.values[0]
}

func (o *gradeCommandss) DeleteGrade() *GradeCommands {
    return _gradeCommandss.values[1]
}

func (o *gradeCommandss) ChangeGrade() *GradeCommands {
    return _gradeCommandss.values[2]
}

func (o *gradeCommandss) ParseGradeCommands(name string) (ret *GradeCommands, ok bool) {
    switch name {
      case "RegisterGrade":
        ret = o.RegisterGrade()
      case "DeleteGrade":
        ret = o.DeleteGrade()
      case "ChangeGrade":
        ret = o.ChangeGrade()
    }
    return
}


type GroupCommands struct {
	name  string
	ordinal int
}

func (o *GroupCommands) Name() string {
    return o.name
}

func (o *GroupCommands) Ordinal() int {
    return o.ordinal
}

func (o *GroupCommands) IsRegisterGroup() bool {
    return o == _groupCommandss.RegisterGroup()
}

func (o *GroupCommands) IsDeleteGroup() bool {
    return o == _groupCommandss.DeleteGroup()
}

func (o *GroupCommands) IsChangeGroup() bool {
    return o == _groupCommandss.ChangeGroup()
}

type groupCommandss struct {
	values []*GroupCommands
}

var _groupCommandss = &groupCommandss{values: []*GroupCommands{
    {name: "registerGroup", ordinal: 0},
    {name: "deleteGroup", ordinal: 1},
    {name: "changeGroup", ordinal: 2}},
}

func GroupCommandss() *groupCommandss {
	return _groupCommandss
}

func (o *groupCommandss) Values() []*GroupCommands {
	return o.values
}

func (o *groupCommandss) RegisterGroup() *GroupCommands {
    return _groupCommandss.values[0]
}

func (o *groupCommandss) DeleteGroup() *GroupCommands {
    return _groupCommandss.values[1]
}

func (o *groupCommandss) ChangeGroup() *GroupCommands {
    return _groupCommandss.values[2]
}

func (o *groupCommandss) ParseGroupCommands(name string) (ret *GroupCommands, ok bool) {
    switch name {
      case "RegisterGroup":
        ret = o.RegisterGroup()
      case "DeleteGroup":
        ret = o.DeleteGroup()
      case "ChangeGroup":
        ret = o.ChangeGroup()
    }
    return
}


type SchoolApplicationCommands struct {
	name  string
	ordinal int
}

func (o *SchoolApplicationCommands) Name() string {
    return o.name
}

func (o *SchoolApplicationCommands) Ordinal() int {
    return o.ordinal
}

func (o *SchoolApplicationCommands) IsRegisterSchoolApplication() bool {
    return o == _schoolApplicationCommandss.RegisterSchoolApplication()
}

func (o *SchoolApplicationCommands) IsDeleteSchoolApplication() bool {
    return o == _schoolApplicationCommandss.DeleteSchoolApplication()
}

func (o *SchoolApplicationCommands) IsChangeSchoolApplication() bool {
    return o == _schoolApplicationCommandss.ChangeSchoolApplication()
}

type schoolApplicationCommandss struct {
	values []*SchoolApplicationCommands
}

var _schoolApplicationCommandss = &schoolApplicationCommandss{values: []*SchoolApplicationCommands{
    {name: "registerSchoolApplication", ordinal: 0},
    {name: "deleteSchoolApplication", ordinal: 1},
    {name: "changeSchoolApplication", ordinal: 2}},
}

func SchoolApplicationCommandss() *schoolApplicationCommandss {
	return _schoolApplicationCommandss
}

func (o *schoolApplicationCommandss) Values() []*SchoolApplicationCommands {
	return o.values
}

func (o *schoolApplicationCommandss) RegisterSchoolApplication() *SchoolApplicationCommands {
    return _schoolApplicationCommandss.values[0]
}

func (o *schoolApplicationCommandss) DeleteSchoolApplication() *SchoolApplicationCommands {
    return _schoolApplicationCommandss.values[1]
}

func (o *schoolApplicationCommandss) ChangeSchoolApplication() *SchoolApplicationCommands {
    return _schoolApplicationCommandss.values[2]
}

func (o *schoolApplicationCommandss) ParseSchoolApplicationCommands(name string) (ret *SchoolApplicationCommands, ok bool) {
    switch name {
      case "RegisterSchoolApplication":
        ret = o.RegisterSchoolApplication()
      case "DeleteSchoolApplication":
        ret = o.DeleteSchoolApplication()
      case "ChangeSchoolApplication":
        ret = o.ChangeSchoolApplication()
    }
    return
}


type SchoolYearCommands struct {
	name  string
	ordinal int
}

func (o *SchoolYearCommands) Name() string {
    return o.name
}

func (o *SchoolYearCommands) Ordinal() int {
    return o.ordinal
}

func (o *SchoolYearCommands) IsRegisterSchoolYear() bool {
    return o == _schoolYearCommandss.RegisterSchoolYear()
}

func (o *SchoolYearCommands) IsDeleteSchoolYear() bool {
    return o == _schoolYearCommandss.DeleteSchoolYear()
}

func (o *SchoolYearCommands) IsChangeSchoolYear() bool {
    return o == _schoolYearCommandss.ChangeSchoolYear()
}

type schoolYearCommandss struct {
	values []*SchoolYearCommands
}

var _schoolYearCommandss = &schoolYearCommandss{values: []*SchoolYearCommands{
    {name: "registerSchoolYear", ordinal: 0},
    {name: "deleteSchoolYear", ordinal: 1},
    {name: "changeSchoolYear", ordinal: 2}},
}

func SchoolYearCommandss() *schoolYearCommandss {
	return _schoolYearCommandss
}

func (o *schoolYearCommandss) Values() []*SchoolYearCommands {
	return o.values
}

func (o *schoolYearCommandss) RegisterSchoolYear() *SchoolYearCommands {
    return _schoolYearCommandss.values[0]
}

func (o *schoolYearCommandss) DeleteSchoolYear() *SchoolYearCommands {
    return _schoolYearCommandss.values[1]
}

func (o *schoolYearCommandss) ChangeSchoolYear() *SchoolYearCommands {
    return _schoolYearCommandss.values[2]
}

func (o *schoolYearCommandss) ParseSchoolYearCommands(name string) (ret *SchoolYearCommands, ok bool) {
    switch name {
      case "RegisterSchoolYear":
        ret = o.RegisterSchoolYear()
      case "DeleteSchoolYear":
        ret = o.DeleteSchoolYear()
      case "ChangeSchoolYear":
        ret = o.ChangeSchoolYear()
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



