package student

import (
    "ee/schkola"
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "time"
)

type AttendanceCreate struct {
    student *person.Profile
    date *time.Time
    course *Course
    hours int
    state *AttendanceState
    stateTrace *schkola.Trace
    token string
    tokenTrace *schkola.Trace
}

func NewAttendanceCreate(student *person.Profile, date *time.Time, course *Course, hours int, state *AttendanceState, stateTrace *schkola.Trace, 
                token string, tokenTrace *schkola.Trace) (ret *AttendanceCreate, err error) {
    ret = &AttendanceCreate{
        student: student,
        date: date,
        course: course,
        hours: hours,
        state: state,
        stateTrace: stateTrace,
        token: token,
        tokenTrace: tokenTrace,
    }
    
    return
}



type AttendanceDelete struct {
    id string
}

func NewAttendanceDelete(id string) (ret *AttendanceDelete, err error) {
    ret = &AttendanceDelete{
        id: id,
    }
    
    return
}



type AttendanceUpdate struct {
    student *person.Profile
    date *time.Time
    course *Course
    hours int
    state *AttendanceState
    stateTrace *schkola.Trace
    token string
    tokenTrace *schkola.Trace
}

func NewAttendanceUpdate(student *person.Profile, date *time.Time, course *Course, hours int, state *AttendanceState, stateTrace *schkola.Trace, 
                token string, tokenTrace *schkola.Trace) (ret *AttendanceUpdate, err error) {
    ret = &AttendanceUpdate{
        student: student,
        date: date,
        course: course,
        hours: hours,
        state: state,
        stateTrace: stateTrace,
        token: token,
        tokenTrace: tokenTrace,
    }
    
    return
}



type AttendanceConfirm struct {
}



type AttendanceCancel struct {
}



type AttendanceRegister struct {
    student *person.Profile
    course *Course
}

func NewAttendanceRegister(student *person.Profile, course *Course) (ret *AttendanceRegister, err error) {
    ret = &AttendanceRegister{
        student: student,
        course: course,
    }
    
    return
}



type CourseCreate struct {
    name string
    begin *time.Time
    end *time.Time
    teacher *person.PersonName
    schoolYear *SchoolYear
    fee float64
    description string
}

func NewCourseCreate(name string, begin *time.Time, end *time.Time, teacher *person.PersonName, schoolYear *SchoolYear, fee float64, 
                description string) (ret *CourseCreate, err error) {
    ret = &CourseCreate{
        name: name,
        begin: begin,
        end: end,
        teacher: teacher,
        schoolYear: schoolYear,
        fee: fee,
        description: description,
    }
    
    return
}



type CourseDelete struct {
    id string
}

func NewCourseDelete(id string) (ret *CourseDelete, err error) {
    ret = &CourseDelete{
        id: id,
    }
    
    return
}



type CourseUpdate struct {
    name string
    begin *time.Time
    end *time.Time
    teacher *person.PersonName
    schoolYear *SchoolYear
    fee float64
    description string
}

func NewCourseUpdate(name string, begin *time.Time, end *time.Time, teacher *person.PersonName, schoolYear *SchoolYear, fee float64, 
                description string) (ret *CourseUpdate, err error) {
    ret = &CourseUpdate{
        name: name,
        begin: begin,
        end: end,
        teacher: teacher,
        schoolYear: schoolYear,
        fee: fee,
        description: description,
    }
    
    return
}



type GradeCreate struct {
    student *person.Profile
    course *Course
    grade float64
    gradeTrace *schkola.Trace
    comment string
}

func NewGradeCreate(student *person.Profile, course *Course, grade float64, gradeTrace *schkola.Trace, comment string) (ret *GradeCreate, err error) {
    ret = &GradeCreate{
        student: student,
        course: course,
        grade: grade,
        gradeTrace: gradeTrace,
        comment: comment,
    }
    
    return
}



type GradeDelete struct {
    id string
}

func NewGradeDelete(id string) (ret *GradeDelete, err error) {
    ret = &GradeDelete{
        id: id,
    }
    
    return
}



type GradeUpdate struct {
    student *person.Profile
    course *Course
    grade float64
    gradeTrace *schkola.Trace
    comment string
}

func NewGradeUpdate(student *person.Profile, course *Course, grade float64, gradeTrace *schkola.Trace, comment string) (ret *GradeUpdate, err error) {
    ret = &GradeUpdate{
        student: student,
        course: course,
        grade: grade,
        gradeTrace: gradeTrace,
        comment: comment,
    }
    
    return
}



type GroupCreate struct {
    name string
    category *GroupCategory
    schoolYear *SchoolYear
    representative *person.Profile
    students []*Course
    courses []*Course
}

func NewGroupCreate(name string, category *GroupCategory, schoolYear *SchoolYear, representative *person.Profile, students []*Course, 
                courses []*Course) (ret *GroupCreate, err error) {
    ret = &GroupCreate{
        name: name,
        category: category,
        schoolYear: schoolYear,
        representative: representative,
        students: students,
        courses: courses,
    }
    
    return
}

func (o *GroupCreate) AddToStudents(item *Course) *Course {
    o.students = append(o.students, item)
    return item
}

func (o *GroupCreate) AddToCourses(item *Course) *Course {
    o.courses = append(o.courses, item)
    return item
}



type GroupDelete struct {
    id string
}

func NewGroupDelete(id string) (ret *GroupDelete, err error) {
    ret = &GroupDelete{
        id: id,
    }
    
    return
}



type GroupUpdate struct {
    name string
    category *GroupCategory
    schoolYear *SchoolYear
    representative *person.Profile
    students []*Course
    courses []*Course
}

func NewGroupUpdate(name string, category *GroupCategory, schoolYear *SchoolYear, representative *person.Profile, students []*Course, 
                courses []*Course) (ret *GroupUpdate, err error) {
    ret = &GroupUpdate{
        name: name,
        category: category,
        schoolYear: schoolYear,
        representative: representative,
        students: students,
        courses: courses,
    }
    
    return
}

func (o *GroupUpdate) AddToStudents(item *Course) *Course {
    o.students = append(o.students, item)
    return item
}

func (o *GroupUpdate) AddToCourses(item *Course) *Course {
    o.courses = append(o.courses, item)
    return item
}



type SchoolApplicationCreate struct {
    profile *person.Profile
    recommendationOf *person.PersonName
    churchContactPerson *person.PersonName
    churchContact *person.Contact
    schoolYear *SchoolYear
    group string
}

func NewSchoolApplicationCreate(profile *person.Profile, recommendationOf *person.PersonName, churchContactPerson *person.PersonName, 
                churchContact *person.Contact, schoolYear *SchoolYear, group string) (ret *SchoolApplicationCreate, err error) {
    ret = &SchoolApplicationCreate{
        profile: profile,
        recommendationOf: recommendationOf,
        churchContactPerson: churchContactPerson,
        churchContact: churchContact,
        schoolYear: schoolYear,
        group: group,
    }
    
    return
}



type SchoolApplicationDelete struct {
    id string
}

func NewSchoolApplicationDelete(id string) (ret *SchoolApplicationDelete, err error) {
    ret = &SchoolApplicationDelete{
        id: id,
    }
    
    return
}



type SchoolApplicationUpdate struct {
    profile *person.Profile
    recommendationOf *person.PersonName
    churchContactPerson *person.PersonName
    churchContact *person.Contact
    schoolYear *SchoolYear
    group string
}

func NewSchoolApplicationUpdate(profile *person.Profile, recommendationOf *person.PersonName, churchContactPerson *person.PersonName, 
                churchContact *person.Contact, schoolYear *SchoolYear, group string) (ret *SchoolApplicationUpdate, err error) {
    ret = &SchoolApplicationUpdate{
        profile: profile,
        recommendationOf: recommendationOf,
        churchContactPerson: churchContactPerson,
        churchContact: churchContact,
        schoolYear: schoolYear,
        group: group,
    }
    
    return
}



type SchoolYearCreate struct {
    name string
    start *time.Time
    end *time.Time
    dates []*Course
}

func NewSchoolYearCreate(name string, start *time.Time, end *time.Time, dates []*Course) (ret *SchoolYearCreate, err error) {
    ret = &SchoolYearCreate{
        name: name,
        start: start,
        end: end,
        dates: dates,
    }
    
    return
}

func (o *SchoolYearCreate) AddToDates(item *Course) *Course {
    o.dates = append(o.dates, item)
    return item
}



type SchoolYearDelete struct {
    id string
}

func NewSchoolYearDelete(id string) (ret *SchoolYearDelete, err error) {
    ret = &SchoolYearDelete{
        id: id,
    }
    
    return
}



type SchoolYearUpdate struct {
    name string
    start *time.Time
    end *time.Time
    dates []*Course
}

func NewSchoolYearUpdate(name string, start *time.Time, end *time.Time, dates []*Course) (ret *SchoolYearUpdate, err error) {
    ret = &SchoolYearUpdate{
        name: name,
        start: start,
        end: end,
        dates: dates,
    }
    
    return
}

func (o *SchoolYearUpdate) AddToDates(item *Course) *Course {
    o.dates = append(o.dates, item)
    return item
}




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

func (o *AttendanceCommandType) IsCreateAttendance() bool {
    return o == _attendanceCommandTypes.CreateAttendance()
}

func (o *AttendanceCommandType) IsDeleteAttendance() bool {
    return o == _attendanceCommandTypes.DeleteAttendance()
}

func (o *AttendanceCommandType) IsUpdateAttendance() bool {
    return o == _attendanceCommandTypes.UpdateAttendance()
}

func (o *AttendanceCommandType) IsConfirmAttendance() bool {
    return o == _attendanceCommandTypes.ConfirmAttendance()
}

func (o *AttendanceCommandType) IsCancelAttendance() bool {
    return o == _attendanceCommandTypes.CancelAttendance()
}

func (o *AttendanceCommandType) IsRegisterAttendance() bool {
    return o == _attendanceCommandTypes.RegisterAttendance()
}

type attendanceCommandTypes struct {
	values []*AttendanceCommandType
    literals []enum.Literal
}

var _attendanceCommandTypes = &attendanceCommandTypes{values: []*AttendanceCommandType{
    {name: "createAttendance", ordinal: 0},
    {name: "deleteAttendance", ordinal: 1},
    {name: "updateAttendance", ordinal: 2},
    {name: "confirmAttendance", ordinal: 3},
    {name: "cancelAttendance", ordinal: 4},
    {name: "registerAttendance", ordinal: 5}},
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

func (o *attendanceCommandTypes) CreateAttendance() *AttendanceCommandType {
    return _attendanceCommandTypes.values[0]
}

func (o *attendanceCommandTypes) DeleteAttendance() *AttendanceCommandType {
    return _attendanceCommandTypes.values[1]
}

func (o *attendanceCommandTypes) UpdateAttendance() *AttendanceCommandType {
    return _attendanceCommandTypes.values[2]
}

func (o *attendanceCommandTypes) ConfirmAttendance() *AttendanceCommandType {
    return _attendanceCommandTypes.values[3]
}

func (o *attendanceCommandTypes) CancelAttendance() *AttendanceCommandType {
    return _attendanceCommandTypes.values[4]
}

func (o *attendanceCommandTypes) RegisterAttendance() *AttendanceCommandType {
    return _attendanceCommandTypes.values[5]
}

func (o *attendanceCommandTypes) ParseAttendanceCommandType(name string) (ret *AttendanceCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
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

func (o *CourseCommandType) IsCreateCourse() bool {
    return o == _courseCommandTypes.CreateCourse()
}

func (o *CourseCommandType) IsDeleteCourse() bool {
    return o == _courseCommandTypes.DeleteCourse()
}

func (o *CourseCommandType) IsUpdateCourse() bool {
    return o == _courseCommandTypes.UpdateCourse()
}

type courseCommandTypes struct {
	values []*CourseCommandType
    literals []enum.Literal
}

var _courseCommandTypes = &courseCommandTypes{values: []*CourseCommandType{
    {name: "createCourse", ordinal: 0},
    {name: "deleteCourse", ordinal: 1},
    {name: "updateCourse", ordinal: 2}},
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

func (o *courseCommandTypes) CreateCourse() *CourseCommandType {
    return _courseCommandTypes.values[0]
}

func (o *courseCommandTypes) DeleteCourse() *CourseCommandType {
    return _courseCommandTypes.values[1]
}

func (o *courseCommandTypes) UpdateCourse() *CourseCommandType {
    return _courseCommandTypes.values[2]
}

func (o *courseCommandTypes) ParseCourseCommandType(name string) (ret *CourseCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
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

func (o *GradeCommandType) IsCreateGrade() bool {
    return o == _gradeCommandTypes.CreateGrade()
}

func (o *GradeCommandType) IsDeleteGrade() bool {
    return o == _gradeCommandTypes.DeleteGrade()
}

func (o *GradeCommandType) IsUpdateGrade() bool {
    return o == _gradeCommandTypes.UpdateGrade()
}

type gradeCommandTypes struct {
	values []*GradeCommandType
    literals []enum.Literal
}

var _gradeCommandTypes = &gradeCommandTypes{values: []*GradeCommandType{
    {name: "createGrade", ordinal: 0},
    {name: "deleteGrade", ordinal: 1},
    {name: "updateGrade", ordinal: 2}},
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

func (o *gradeCommandTypes) CreateGrade() *GradeCommandType {
    return _gradeCommandTypes.values[0]
}

func (o *gradeCommandTypes) DeleteGrade() *GradeCommandType {
    return _gradeCommandTypes.values[1]
}

func (o *gradeCommandTypes) UpdateGrade() *GradeCommandType {
    return _gradeCommandTypes.values[2]
}

func (o *gradeCommandTypes) ParseGradeCommandType(name string) (ret *GradeCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
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

func (o *GroupCommandType) IsCreateGroup() bool {
    return o == _groupCommandTypes.CreateGroup()
}

func (o *GroupCommandType) IsDeleteGroup() bool {
    return o == _groupCommandTypes.DeleteGroup()
}

func (o *GroupCommandType) IsUpdateGroup() bool {
    return o == _groupCommandTypes.UpdateGroup()
}

type groupCommandTypes struct {
	values []*GroupCommandType
    literals []enum.Literal
}

var _groupCommandTypes = &groupCommandTypes{values: []*GroupCommandType{
    {name: "createGroup", ordinal: 0},
    {name: "deleteGroup", ordinal: 1},
    {name: "updateGroup", ordinal: 2}},
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

func (o *groupCommandTypes) CreateGroup() *GroupCommandType {
    return _groupCommandTypes.values[0]
}

func (o *groupCommandTypes) DeleteGroup() *GroupCommandType {
    return _groupCommandTypes.values[1]
}

func (o *groupCommandTypes) UpdateGroup() *GroupCommandType {
    return _groupCommandTypes.values[2]
}

func (o *groupCommandTypes) ParseGroupCommandType(name string) (ret *GroupCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
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

func (o *SchoolApplicationCommandType) IsCreateSchoolApplication() bool {
    return o == _schoolApplicationCommandTypes.CreateSchoolApplication()
}

func (o *SchoolApplicationCommandType) IsDeleteSchoolApplication() bool {
    return o == _schoolApplicationCommandTypes.DeleteSchoolApplication()
}

func (o *SchoolApplicationCommandType) IsUpdateSchoolApplication() bool {
    return o == _schoolApplicationCommandTypes.UpdateSchoolApplication()
}

type schoolApplicationCommandTypes struct {
	values []*SchoolApplicationCommandType
    literals []enum.Literal
}

var _schoolApplicationCommandTypes = &schoolApplicationCommandTypes{values: []*SchoolApplicationCommandType{
    {name: "createSchoolApplication", ordinal: 0},
    {name: "deleteSchoolApplication", ordinal: 1},
    {name: "updateSchoolApplication", ordinal: 2}},
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

func (o *schoolApplicationCommandTypes) CreateSchoolApplication() *SchoolApplicationCommandType {
    return _schoolApplicationCommandTypes.values[0]
}

func (o *schoolApplicationCommandTypes) DeleteSchoolApplication() *SchoolApplicationCommandType {
    return _schoolApplicationCommandTypes.values[1]
}

func (o *schoolApplicationCommandTypes) UpdateSchoolApplication() *SchoolApplicationCommandType {
    return _schoolApplicationCommandTypes.values[2]
}

func (o *schoolApplicationCommandTypes) ParseSchoolApplicationCommandType(name string) (ret *SchoolApplicationCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
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

func (o *SchoolYearCommandType) IsCreateSchoolYear() bool {
    return o == _schoolYearCommandTypes.CreateSchoolYear()
}

func (o *SchoolYearCommandType) IsDeleteSchoolYear() bool {
    return o == _schoolYearCommandTypes.DeleteSchoolYear()
}

func (o *SchoolYearCommandType) IsUpdateSchoolYear() bool {
    return o == _schoolYearCommandTypes.UpdateSchoolYear()
}

type schoolYearCommandTypes struct {
	values []*SchoolYearCommandType
    literals []enum.Literal
}

var _schoolYearCommandTypes = &schoolYearCommandTypes{values: []*SchoolYearCommandType{
    {name: "createSchoolYear", ordinal: 0},
    {name: "deleteSchoolYear", ordinal: 1},
    {name: "updateSchoolYear", ordinal: 2}},
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

func (o *schoolYearCommandTypes) CreateSchoolYear() *SchoolYearCommandType {
    return _schoolYearCommandTypes.values[0]
}

func (o *schoolYearCommandTypes) DeleteSchoolYear() *SchoolYearCommandType {
    return _schoolYearCommandTypes.values[1]
}

func (o *schoolYearCommandTypes) UpdateSchoolYear() *SchoolYearCommandType {
    return _schoolYearCommandTypes.values[2]
}

func (o *schoolYearCommandTypes) ParseSchoolYearCommandType(name string) (ret *SchoolYearCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*SchoolYearCommandType), ok
	}
	return
}



