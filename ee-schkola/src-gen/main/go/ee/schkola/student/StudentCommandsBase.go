package student

import (
    "ee/schkola"
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "time"
)

type AttendanceCreate struct {
    Student  *person.Profile
    Date  *time.Time
    Course  *Course
    Hours  int
    State  *AttendanceState
    StateTrace  *schkola.Trace
    Token  string
    TokenTrace  *schkola.Trace
}

func NewAttendanceCreate(student *person.Profile, date *time.Time, course *Course, hours int, state *AttendanceState, stateTrace *schkola.Trace, 
                token string, tokenTrace *schkola.Trace) (ret *AttendanceCreate, err error) {
    ret = &AttendanceCreate{
        Student : student,
        Date : date,
        Course : course,
        Hours : hours,
        State : state,
        StateTrace : stateTrace,
        Token : token,
        TokenTrace : tokenTrace,
    }
    return
}



type AttendanceDelete struct {
    Id  string
}

func NewAttendanceDelete(id string) (ret *AttendanceDelete, err error) {
    ret = &AttendanceDelete{
        Id : id,
    }
    return
}



type AttendanceUpdate struct {
    Student  *person.Profile
    Date  *time.Time
    Course  *Course
    Hours  int
    State  *AttendanceState
    StateTrace  *schkola.Trace
    Token  string
    TokenTrace  *schkola.Trace
}

func NewAttendanceUpdate(student *person.Profile, date *time.Time, course *Course, hours int, state *AttendanceState, stateTrace *schkola.Trace, 
                token string, tokenTrace *schkola.Trace) (ret *AttendanceUpdate, err error) {
    ret = &AttendanceUpdate{
        Student : student,
        Date : date,
        Course : course,
        Hours : hours,
        State : state,
        StateTrace : stateTrace,
        Token : token,
        TokenTrace : tokenTrace,
    }
    return
}



type AttendanceConfirm struct {
}



type AttendanceCancel struct {
}



type AttendanceRegister struct {
    Student  *person.Profile
    Course  *Course
}

func NewAttendanceRegister(student *person.Profile, course *Course) (ret *AttendanceRegister, err error) {
    ret = &AttendanceRegister{
        Student : student,
        Course : course,
    }
    return
}



type CourseCreate struct {
    Name  string
    Begin  *time.Time
    End  *time.Time
    Teacher  *person.PersonName
    SchoolYear  *SchoolYear
    Fee  float64
    Description  string
}

func NewCourseCreate(name string, begin *time.Time, end *time.Time, teacher *person.PersonName, schoolYear *SchoolYear, fee float64, 
                description string) (ret *CourseCreate, err error) {
    ret = &CourseCreate{
        Name : name,
        Begin : begin,
        End : end,
        Teacher : teacher,
        SchoolYear : schoolYear,
        Fee : fee,
        Description : description,
    }
    return
}



type CourseDelete struct {
    Id  string
}

func NewCourseDelete(id string) (ret *CourseDelete, err error) {
    ret = &CourseDelete{
        Id : id,
    }
    return
}



type CourseUpdate struct {
    Name  string
    Begin  *time.Time
    End  *time.Time
    Teacher  *person.PersonName
    SchoolYear  *SchoolYear
    Fee  float64
    Description  string
}

func NewCourseUpdate(name string, begin *time.Time, end *time.Time, teacher *person.PersonName, schoolYear *SchoolYear, fee float64, 
                description string) (ret *CourseUpdate, err error) {
    ret = &CourseUpdate{
        Name : name,
        Begin : begin,
        End : end,
        Teacher : teacher,
        SchoolYear : schoolYear,
        Fee : fee,
        Description : description,
    }
    return
}



type GradeCreate struct {
    Student  *person.Profile
    Course  *Course
    Grade  float64
    GradeTrace  *schkola.Trace
    Comment  string
}

func NewGradeCreate(student *person.Profile, course *Course, grade float64, gradeTrace *schkola.Trace, comment string) (ret *GradeCreate, err error) {
    ret = &GradeCreate{
        Student : student,
        Course : course,
        Grade : grade,
        GradeTrace : gradeTrace,
        Comment : comment,
    }
    return
}



type GradeDelete struct {
    Id  string
}

func NewGradeDelete(id string) (ret *GradeDelete, err error) {
    ret = &GradeDelete{
        Id : id,
    }
    return
}



type GradeUpdate struct {
    Student  *person.Profile
    Course  *Course
    Grade  float64
    GradeTrace  *schkola.Trace
    Comment  string
}

func NewGradeUpdate(student *person.Profile, course *Course, grade float64, gradeTrace *schkola.Trace, comment string) (ret *GradeUpdate, err error) {
    ret = &GradeUpdate{
        Student : student,
        Course : course,
        Grade : grade,
        GradeTrace : gradeTrace,
        Comment : comment,
    }
    return
}



type GroupCreate struct {
    Name  string
    Category  *GroupCategory
    SchoolYear  *SchoolYear
    Representative  *person.Profile
    Students  []*Course
    Courses  []*Course
}

func NewGroupCreate(name string, category *GroupCategory, schoolYear *SchoolYear, representative *person.Profile, students []*Course, 
                courses []*Course) (ret *GroupCreate, err error) {
    ret = &GroupCreate{
        Name : name,
        Category : category,
        SchoolYear : schoolYear,
        Representative : representative,
        Students : students,
        Courses : courses,
    }
    return
}

func (o *GroupCreate) AddToStudents(item *Course) *Course {
    o.Students  = append(o.Students , item)
    return item
}

func (o *GroupCreate) AddToCourses(item *Course) *Course {
    o.Courses  = append(o.Courses , item)
    return item
}



type GroupDelete struct {
    Id  string
}

func NewGroupDelete(id string) (ret *GroupDelete, err error) {
    ret = &GroupDelete{
        Id : id,
    }
    return
}



type GroupUpdate struct {
    Name  string
    Category  *GroupCategory
    SchoolYear  *SchoolYear
    Representative  *person.Profile
    Students  []*Course
    Courses  []*Course
}

func NewGroupUpdate(name string, category *GroupCategory, schoolYear *SchoolYear, representative *person.Profile, students []*Course, 
                courses []*Course) (ret *GroupUpdate, err error) {
    ret = &GroupUpdate{
        Name : name,
        Category : category,
        SchoolYear : schoolYear,
        Representative : representative,
        Students : students,
        Courses : courses,
    }
    return
}

func (o *GroupUpdate) AddToStudents(item *Course) *Course {
    o.Students  = append(o.Students , item)
    return item
}

func (o *GroupUpdate) AddToCourses(item *Course) *Course {
    o.Courses  = append(o.Courses , item)
    return item
}



type SchoolApplicationCreate struct {
    Profile  *person.Profile
    RecommendationOf  *person.PersonName
    ChurchContactPerson  *person.PersonName
    ChurchContact  *person.Contact
    SchoolYear  *SchoolYear
    Group  string
}

func NewSchoolApplicationCreate(profile *person.Profile, recommendationOf *person.PersonName, churchContactPerson *person.PersonName, 
                churchContact *person.Contact, schoolYear *SchoolYear, group string) (ret *SchoolApplicationCreate, err error) {
    ret = &SchoolApplicationCreate{
        Profile : profile,
        RecommendationOf : recommendationOf,
        ChurchContactPerson : churchContactPerson,
        ChurchContact : churchContact,
        SchoolYear : schoolYear,
        Group : group,
    }
    return
}



type SchoolApplicationDelete struct {
    Id  string
}

func NewSchoolApplicationDelete(id string) (ret *SchoolApplicationDelete, err error) {
    ret = &SchoolApplicationDelete{
        Id : id,
    }
    return
}



type SchoolApplicationUpdate struct {
    Profile  *person.Profile
    RecommendationOf  *person.PersonName
    ChurchContactPerson  *person.PersonName
    ChurchContact  *person.Contact
    SchoolYear  *SchoolYear
    Group  string
}

func NewSchoolApplicationUpdate(profile *person.Profile, recommendationOf *person.PersonName, churchContactPerson *person.PersonName, 
                churchContact *person.Contact, schoolYear *SchoolYear, group string) (ret *SchoolApplicationUpdate, err error) {
    ret = &SchoolApplicationUpdate{
        Profile : profile,
        RecommendationOf : recommendationOf,
        ChurchContactPerson : churchContactPerson,
        ChurchContact : churchContact,
        SchoolYear : schoolYear,
        Group : group,
    }
    return
}



type SchoolYearCreate struct {
    Name  string
    Start  *time.Time
    End  *time.Time
    Dates  []*Course
}

func NewSchoolYearCreate(name string, start *time.Time, end *time.Time, dates []*Course) (ret *SchoolYearCreate, err error) {
    ret = &SchoolYearCreate{
        Name : name,
        Start : start,
        End : end,
        Dates : dates,
    }
    return
}

func (o *SchoolYearCreate) AddToDates(item *Course) *Course {
    o.Dates  = append(o.Dates , item)
    return item
}



type SchoolYearDelete struct {
    Id  string
}

func NewSchoolYearDelete(id string) (ret *SchoolYearDelete, err error) {
    ret = &SchoolYearDelete{
        Id : id,
    }
    return
}



type SchoolYearUpdate struct {
    Name  string
    Start  *time.Time
    End  *time.Time
    Dates  []*Course
}

func NewSchoolYearUpdate(name string, start *time.Time, end *time.Time, dates []*Course) (ret *SchoolYearUpdate, err error) {
    ret = &SchoolYearUpdate{
        Name : name,
        Start : start,
        End : end,
        Dates : dates,
    }
    return
}

func (o *SchoolYearUpdate) AddToDates(item *Course) *Course {
    o.Dates  = append(o.Dates , item)
    return item
}




type AttendanceAggregateCommandType struct {
	name  string
	ordinal int
}

func (o *AttendanceAggregateCommandType) Name() string {
    return o.name
}

func (o *AttendanceAggregateCommandType) Ordinal() int {
    return o.ordinal
}

func (o *AttendanceAggregateCommandType) IsCreateAttendance() bool {
    return o == _attendanceAggregateCommandTypes.CreateAttendance()
}

func (o *AttendanceAggregateCommandType) IsDeleteAttendance() bool {
    return o == _attendanceAggregateCommandTypes.DeleteAttendance()
}

func (o *AttendanceAggregateCommandType) IsUpdateAttendance() bool {
    return o == _attendanceAggregateCommandTypes.UpdateAttendance()
}

func (o *AttendanceAggregateCommandType) IsConfirmAttendance() bool {
    return o == _attendanceAggregateCommandTypes.ConfirmAttendance()
}

func (o *AttendanceAggregateCommandType) IsCancelAttendance() bool {
    return o == _attendanceAggregateCommandTypes.CancelAttendance()
}

func (o *AttendanceAggregateCommandType) IsRegisterAttendance() bool {
    return o == _attendanceAggregateCommandTypes.RegisterAttendance()
}

type attendanceAggregateCommandTypes struct {
	values []*AttendanceAggregateCommandType
    literals []enum.Literal
}

var _attendanceAggregateCommandTypes = &attendanceAggregateCommandTypes{values: []*AttendanceAggregateCommandType{
    {name: "createAttendance", ordinal: 0},
    {name: "deleteAttendance", ordinal: 1},
    {name: "updateAttendance", ordinal: 2},
    {name: "confirmAttendance", ordinal: 3},
    {name: "cancelAttendance", ordinal: 4},
    {name: "registerAttendance", ordinal: 5}},
}

func AttendanceAggregateCommandTypes() *attendanceAggregateCommandTypes {
	return _attendanceAggregateCommandTypes
}

func (o *attendanceAggregateCommandTypes) Values() []*AttendanceAggregateCommandType {
	return o.values
}

func (o *attendanceAggregateCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *attendanceAggregateCommandTypes) CreateAttendance() *AttendanceAggregateCommandType {
    return _attendanceAggregateCommandTypes.values[0]
}

func (o *attendanceAggregateCommandTypes) DeleteAttendance() *AttendanceAggregateCommandType {
    return _attendanceAggregateCommandTypes.values[1]
}

func (o *attendanceAggregateCommandTypes) UpdateAttendance() *AttendanceAggregateCommandType {
    return _attendanceAggregateCommandTypes.values[2]
}

func (o *attendanceAggregateCommandTypes) ConfirmAttendance() *AttendanceAggregateCommandType {
    return _attendanceAggregateCommandTypes.values[3]
}

func (o *attendanceAggregateCommandTypes) CancelAttendance() *AttendanceAggregateCommandType {
    return _attendanceAggregateCommandTypes.values[4]
}

func (o *attendanceAggregateCommandTypes) RegisterAttendance() *AttendanceAggregateCommandType {
    return _attendanceAggregateCommandTypes.values[5]
}

func (o *attendanceAggregateCommandTypes) ParseAttendanceAggregateCommandType(name string) (ret *AttendanceAggregateCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*AttendanceAggregateCommandType), ok
	}
	return
}


type CourseAggregateCommandType struct {
	name  string
	ordinal int
}

func (o *CourseAggregateCommandType) Name() string {
    return o.name
}

func (o *CourseAggregateCommandType) Ordinal() int {
    return o.ordinal
}

func (o *CourseAggregateCommandType) IsCreateCourse() bool {
    return o == _courseAggregateCommandTypes.CreateCourse()
}

func (o *CourseAggregateCommandType) IsDeleteCourse() bool {
    return o == _courseAggregateCommandTypes.DeleteCourse()
}

func (o *CourseAggregateCommandType) IsUpdateCourse() bool {
    return o == _courseAggregateCommandTypes.UpdateCourse()
}

type courseAggregateCommandTypes struct {
	values []*CourseAggregateCommandType
    literals []enum.Literal
}

var _courseAggregateCommandTypes = &courseAggregateCommandTypes{values: []*CourseAggregateCommandType{
    {name: "createCourse", ordinal: 0},
    {name: "deleteCourse", ordinal: 1},
    {name: "updateCourse", ordinal: 2}},
}

func CourseAggregateCommandTypes() *courseAggregateCommandTypes {
	return _courseAggregateCommandTypes
}

func (o *courseAggregateCommandTypes) Values() []*CourseAggregateCommandType {
	return o.values
}

func (o *courseAggregateCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *courseAggregateCommandTypes) CreateCourse() *CourseAggregateCommandType {
    return _courseAggregateCommandTypes.values[0]
}

func (o *courseAggregateCommandTypes) DeleteCourse() *CourseAggregateCommandType {
    return _courseAggregateCommandTypes.values[1]
}

func (o *courseAggregateCommandTypes) UpdateCourse() *CourseAggregateCommandType {
    return _courseAggregateCommandTypes.values[2]
}

func (o *courseAggregateCommandTypes) ParseCourseAggregateCommandType(name string) (ret *CourseAggregateCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*CourseAggregateCommandType), ok
	}
	return
}


type GradeAggregateCommandType struct {
	name  string
	ordinal int
}

func (o *GradeAggregateCommandType) Name() string {
    return o.name
}

func (o *GradeAggregateCommandType) Ordinal() int {
    return o.ordinal
}

func (o *GradeAggregateCommandType) IsCreateGrade() bool {
    return o == _gradeAggregateCommandTypes.CreateGrade()
}

func (o *GradeAggregateCommandType) IsDeleteGrade() bool {
    return o == _gradeAggregateCommandTypes.DeleteGrade()
}

func (o *GradeAggregateCommandType) IsUpdateGrade() bool {
    return o == _gradeAggregateCommandTypes.UpdateGrade()
}

type gradeAggregateCommandTypes struct {
	values []*GradeAggregateCommandType
    literals []enum.Literal
}

var _gradeAggregateCommandTypes = &gradeAggregateCommandTypes{values: []*GradeAggregateCommandType{
    {name: "createGrade", ordinal: 0},
    {name: "deleteGrade", ordinal: 1},
    {name: "updateGrade", ordinal: 2}},
}

func GradeAggregateCommandTypes() *gradeAggregateCommandTypes {
	return _gradeAggregateCommandTypes
}

func (o *gradeAggregateCommandTypes) Values() []*GradeAggregateCommandType {
	return o.values
}

func (o *gradeAggregateCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *gradeAggregateCommandTypes) CreateGrade() *GradeAggregateCommandType {
    return _gradeAggregateCommandTypes.values[0]
}

func (o *gradeAggregateCommandTypes) DeleteGrade() *GradeAggregateCommandType {
    return _gradeAggregateCommandTypes.values[1]
}

func (o *gradeAggregateCommandTypes) UpdateGrade() *GradeAggregateCommandType {
    return _gradeAggregateCommandTypes.values[2]
}

func (o *gradeAggregateCommandTypes) ParseGradeAggregateCommandType(name string) (ret *GradeAggregateCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*GradeAggregateCommandType), ok
	}
	return
}


type GroupAggregateCommandType struct {
	name  string
	ordinal int
}

func (o *GroupAggregateCommandType) Name() string {
    return o.name
}

func (o *GroupAggregateCommandType) Ordinal() int {
    return o.ordinal
}

func (o *GroupAggregateCommandType) IsCreateGroup() bool {
    return o == _groupAggregateCommandTypes.CreateGroup()
}

func (o *GroupAggregateCommandType) IsDeleteGroup() bool {
    return o == _groupAggregateCommandTypes.DeleteGroup()
}

func (o *GroupAggregateCommandType) IsUpdateGroup() bool {
    return o == _groupAggregateCommandTypes.UpdateGroup()
}

type groupAggregateCommandTypes struct {
	values []*GroupAggregateCommandType
    literals []enum.Literal
}

var _groupAggregateCommandTypes = &groupAggregateCommandTypes{values: []*GroupAggregateCommandType{
    {name: "createGroup", ordinal: 0},
    {name: "deleteGroup", ordinal: 1},
    {name: "updateGroup", ordinal: 2}},
}

func GroupAggregateCommandTypes() *groupAggregateCommandTypes {
	return _groupAggregateCommandTypes
}

func (o *groupAggregateCommandTypes) Values() []*GroupAggregateCommandType {
	return o.values
}

func (o *groupAggregateCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *groupAggregateCommandTypes) CreateGroup() *GroupAggregateCommandType {
    return _groupAggregateCommandTypes.values[0]
}

func (o *groupAggregateCommandTypes) DeleteGroup() *GroupAggregateCommandType {
    return _groupAggregateCommandTypes.values[1]
}

func (o *groupAggregateCommandTypes) UpdateGroup() *GroupAggregateCommandType {
    return _groupAggregateCommandTypes.values[2]
}

func (o *groupAggregateCommandTypes) ParseGroupAggregateCommandType(name string) (ret *GroupAggregateCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*GroupAggregateCommandType), ok
	}
	return
}


type SchoolApplicationAggregateCommandType struct {
	name  string
	ordinal int
}

func (o *SchoolApplicationAggregateCommandType) Name() string {
    return o.name
}

func (o *SchoolApplicationAggregateCommandType) Ordinal() int {
    return o.ordinal
}

func (o *SchoolApplicationAggregateCommandType) IsCreateSchoolApplication() bool {
    return o == _schoolApplicationAggregateCommandTypes.CreateSchoolApplication()
}

func (o *SchoolApplicationAggregateCommandType) IsDeleteSchoolApplication() bool {
    return o == _schoolApplicationAggregateCommandTypes.DeleteSchoolApplication()
}

func (o *SchoolApplicationAggregateCommandType) IsUpdateSchoolApplication() bool {
    return o == _schoolApplicationAggregateCommandTypes.UpdateSchoolApplication()
}

type schoolApplicationAggregateCommandTypes struct {
	values []*SchoolApplicationAggregateCommandType
    literals []enum.Literal
}

var _schoolApplicationAggregateCommandTypes = &schoolApplicationAggregateCommandTypes{values: []*SchoolApplicationAggregateCommandType{
    {name: "createSchoolApplication", ordinal: 0},
    {name: "deleteSchoolApplication", ordinal: 1},
    {name: "updateSchoolApplication", ordinal: 2}},
}

func SchoolApplicationAggregateCommandTypes() *schoolApplicationAggregateCommandTypes {
	return _schoolApplicationAggregateCommandTypes
}

func (o *schoolApplicationAggregateCommandTypes) Values() []*SchoolApplicationAggregateCommandType {
	return o.values
}

func (o *schoolApplicationAggregateCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *schoolApplicationAggregateCommandTypes) CreateSchoolApplication() *SchoolApplicationAggregateCommandType {
    return _schoolApplicationAggregateCommandTypes.values[0]
}

func (o *schoolApplicationAggregateCommandTypes) DeleteSchoolApplication() *SchoolApplicationAggregateCommandType {
    return _schoolApplicationAggregateCommandTypes.values[1]
}

func (o *schoolApplicationAggregateCommandTypes) UpdateSchoolApplication() *SchoolApplicationAggregateCommandType {
    return _schoolApplicationAggregateCommandTypes.values[2]
}

func (o *schoolApplicationAggregateCommandTypes) ParseSchoolApplicationAggregateCommandType(name string) (ret *SchoolApplicationAggregateCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*SchoolApplicationAggregateCommandType), ok
	}
	return
}


type SchoolYearAggregateCommandType struct {
	name  string
	ordinal int
}

func (o *SchoolYearAggregateCommandType) Name() string {
    return o.name
}

func (o *SchoolYearAggregateCommandType) Ordinal() int {
    return o.ordinal
}

func (o *SchoolYearAggregateCommandType) IsCreateSchoolYear() bool {
    return o == _schoolYearAggregateCommandTypes.CreateSchoolYear()
}

func (o *SchoolYearAggregateCommandType) IsDeleteSchoolYear() bool {
    return o == _schoolYearAggregateCommandTypes.DeleteSchoolYear()
}

func (o *SchoolYearAggregateCommandType) IsUpdateSchoolYear() bool {
    return o == _schoolYearAggregateCommandTypes.UpdateSchoolYear()
}

type schoolYearAggregateCommandTypes struct {
	values []*SchoolYearAggregateCommandType
    literals []enum.Literal
}

var _schoolYearAggregateCommandTypes = &schoolYearAggregateCommandTypes{values: []*SchoolYearAggregateCommandType{
    {name: "createSchoolYear", ordinal: 0},
    {name: "deleteSchoolYear", ordinal: 1},
    {name: "updateSchoolYear", ordinal: 2}},
}

func SchoolYearAggregateCommandTypes() *schoolYearAggregateCommandTypes {
	return _schoolYearAggregateCommandTypes
}

func (o *schoolYearAggregateCommandTypes) Values() []*SchoolYearAggregateCommandType {
	return o.values
}

func (o *schoolYearAggregateCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *schoolYearAggregateCommandTypes) CreateSchoolYear() *SchoolYearAggregateCommandType {
    return _schoolYearAggregateCommandTypes.values[0]
}

func (o *schoolYearAggregateCommandTypes) DeleteSchoolYear() *SchoolYearAggregateCommandType {
    return _schoolYearAggregateCommandTypes.values[1]
}

func (o *schoolYearAggregateCommandTypes) UpdateSchoolYear() *SchoolYearAggregateCommandType {
    return _schoolYearAggregateCommandTypes.values[2]
}

func (o *schoolYearAggregateCommandTypes) ParseSchoolYearAggregateCommandType(name string) (ret *SchoolYearAggregateCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*SchoolYearAggregateCommandType), ok
	}
	return
}



