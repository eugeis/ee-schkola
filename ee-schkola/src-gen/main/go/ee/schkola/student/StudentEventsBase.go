package student

import (
    "ee/schkola"
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "time"
)

type CreatedAttendance struct {
    student *person.Profile
    date *time.Time
    course *Course
    hours int
    state *AttendanceState
    stateTrace *schkola.Trace
    token string
    tokenTrace *schkola.Trace
}

func NewCreatedAttendance(student *person.Profile, date *time.Time, course *Course, hours int, state *AttendanceState, stateTrace *schkola.Trace, 
                token string, tokenTrace *schkola.Trace) (ret *CreatedAttendance, err error) {
    ret = &CreatedAttendance{
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



type DeletedAttendance struct {
    id string
}

func NewDeletedAttendance(id string) (ret *DeletedAttendance, err error) {
    ret = &DeletedAttendance{
        id: id,
    }
    
    return
}



type UpdatedAttendance struct {
    student *person.Profile
    date *time.Time
    course *Course
    hours int
    state *AttendanceState
    stateTrace *schkola.Trace
    token string
    tokenTrace *schkola.Trace
}

func NewUpdatedAttendance(student *person.Profile, date *time.Time, course *Course, hours int, state *AttendanceState, stateTrace *schkola.Trace, 
                token string, tokenTrace *schkola.Trace) (ret *UpdatedAttendance, err error) {
    ret = &UpdatedAttendance{
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



type CreatedCourse struct {
    name string
    begin *time.Time
    end *time.Time
    teacher *person.PersonName
    schoolYear *SchoolYear
    fee float64
    description string
}

func NewCreatedCourse(name string, begin *time.Time, end *time.Time, teacher *person.PersonName, schoolYear *SchoolYear, fee float64, 
                description string) (ret *CreatedCourse, err error) {
    ret = &CreatedCourse{
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



type DeletedCourse struct {
    id string
}

func NewDeletedCourse(id string) (ret *DeletedCourse, err error) {
    ret = &DeletedCourse{
        id: id,
    }
    
    return
}



type UpdatedCourse struct {
    name string
    begin *time.Time
    end *time.Time
    teacher *person.PersonName
    schoolYear *SchoolYear
    fee float64
    description string
}

func NewUpdatedCourse(name string, begin *time.Time, end *time.Time, teacher *person.PersonName, schoolYear *SchoolYear, fee float64, 
                description string) (ret *UpdatedCourse, err error) {
    ret = &UpdatedCourse{
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



type CreatedGrade struct {
    student *person.Profile
    course *Course
    grade float64
    gradeTrace *schkola.Trace
    comment string
}

func NewCreatedGrade(student *person.Profile, course *Course, grade float64, gradeTrace *schkola.Trace, comment string) (ret *CreatedGrade, err error) {
    ret = &CreatedGrade{
        student: student,
        course: course,
        grade: grade,
        gradeTrace: gradeTrace,
        comment: comment,
    }
    
    return
}



type DeletedGrade struct {
    id string
}

func NewDeletedGrade(id string) (ret *DeletedGrade, err error) {
    ret = &DeletedGrade{
        id: id,
    }
    
    return
}



type UpdatedGrade struct {
    student *person.Profile
    course *Course
    grade float64
    gradeTrace *schkola.Trace
    comment string
}

func NewUpdatedGrade(student *person.Profile, course *Course, grade float64, gradeTrace *schkola.Trace, comment string) (ret *UpdatedGrade, err error) {
    ret = &UpdatedGrade{
        student: student,
        course: course,
        grade: grade,
        gradeTrace: gradeTrace,
        comment: comment,
    }
    
    return
}



type CreatedGroup struct {
    name string
    category *GroupCategory
    schoolYear *SchoolYear
    representative *person.Profile
    students []*Course
    courses []*Course
}

func NewCreatedGroup(name string, category *GroupCategory, schoolYear *SchoolYear, representative *person.Profile, students []*Course, 
                courses []*Course) (ret *CreatedGroup, err error) {
    ret = &CreatedGroup{
        name: name,
        category: category,
        schoolYear: schoolYear,
        representative: representative,
        students: students,
        courses: courses,
    }
    
    return
}

func (o *CreatedGroup) AddToStudents(item *Course) *Course {
    o.students = append(o.students, item)
    return item
}

func (o *CreatedGroup) AddToCourses(item *Course) *Course {
    o.courses = append(o.courses, item)
    return item
}



type DeletedGroup struct {
    id string
}

func NewDeletedGroup(id string) (ret *DeletedGroup, err error) {
    ret = &DeletedGroup{
        id: id,
    }
    
    return
}



type UpdatedGroup struct {
    name string
    category *GroupCategory
    schoolYear *SchoolYear
    representative *person.Profile
    students []*Course
    courses []*Course
}

func NewUpdatedGroup(name string, category *GroupCategory, schoolYear *SchoolYear, representative *person.Profile, students []*Course, 
                courses []*Course) (ret *UpdatedGroup, err error) {
    ret = &UpdatedGroup{
        name: name,
        category: category,
        schoolYear: schoolYear,
        representative: representative,
        students: students,
        courses: courses,
    }
    
    return
}

func (o *UpdatedGroup) AddToStudents(item *Course) *Course {
    o.students = append(o.students, item)
    return item
}

func (o *UpdatedGroup) AddToCourses(item *Course) *Course {
    o.courses = append(o.courses, item)
    return item
}



type CreatedSchoolApplication struct {
    profile *person.Profile
    recommendationOf *person.PersonName
    churchContactPerson *person.PersonName
    churchContact *person.Contact
    schoolYear *SchoolYear
    group string
}

func NewCreatedSchoolApplication(profile *person.Profile, recommendationOf *person.PersonName, churchContactPerson *person.PersonName, 
                churchContact *person.Contact, schoolYear *SchoolYear, group string) (ret *CreatedSchoolApplication, err error) {
    ret = &CreatedSchoolApplication{
        profile: profile,
        recommendationOf: recommendationOf,
        churchContactPerson: churchContactPerson,
        churchContact: churchContact,
        schoolYear: schoolYear,
        group: group,
    }
    
    return
}



type DeletedSchoolApplication struct {
    id string
}

func NewDeletedSchoolApplication(id string) (ret *DeletedSchoolApplication, err error) {
    ret = &DeletedSchoolApplication{
        id: id,
    }
    
    return
}



type UpdatedSchoolApplication struct {
    profile *person.Profile
    recommendationOf *person.PersonName
    churchContactPerson *person.PersonName
    churchContact *person.Contact
    schoolYear *SchoolYear
    group string
}

func NewUpdatedSchoolApplication(profile *person.Profile, recommendationOf *person.PersonName, churchContactPerson *person.PersonName, 
                churchContact *person.Contact, schoolYear *SchoolYear, group string) (ret *UpdatedSchoolApplication, err error) {
    ret = &UpdatedSchoolApplication{
        profile: profile,
        recommendationOf: recommendationOf,
        churchContactPerson: churchContactPerson,
        churchContact: churchContact,
        schoolYear: schoolYear,
        group: group,
    }
    
    return
}



type CreatedSchoolYear struct {
    name string
    start *time.Time
    end *time.Time
    dates []*Course
}

func NewCreatedSchoolYear(name string, start *time.Time, end *time.Time, dates []*Course) (ret *CreatedSchoolYear, err error) {
    ret = &CreatedSchoolYear{
        name: name,
        start: start,
        end: end,
        dates: dates,
    }
    
    return
}

func (o *CreatedSchoolYear) AddToDates(item *Course) *Course {
    o.dates = append(o.dates, item)
    return item
}



type DeletedSchoolYear struct {
    id string
}

func NewDeletedSchoolYear(id string) (ret *DeletedSchoolYear, err error) {
    ret = &DeletedSchoolYear{
        id: id,
    }
    
    return
}



type UpdatedSchoolYear struct {
    name string
    start *time.Time
    end *time.Time
    dates []*Course
}

func NewUpdatedSchoolYear(name string, start *time.Time, end *time.Time, dates []*Course) (ret *UpdatedSchoolYear, err error) {
    ret = &UpdatedSchoolYear{
        name: name,
        start: start,
        end: end,
        dates: dates,
    }
    
    return
}

func (o *UpdatedSchoolYear) AddToDates(item *Course) *Course {
    o.dates = append(o.dates, item)
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

func (o *AttendanceEventType) IsAttendanceCreated() bool {
    return o == _attendanceEventTypes.AttendanceCreated()
}

func (o *AttendanceEventType) IsAttendanceDeleted() bool {
    return o == _attendanceEventTypes.AttendanceDeleted()
}

func (o *AttendanceEventType) IsAttendanceUpdated() bool {
    return o == _attendanceEventTypes.AttendanceUpdated()
}

type attendanceEventTypes struct {
	values []*AttendanceEventType
    literals []enum.Literal
}

var _attendanceEventTypes = &attendanceEventTypes{values: []*AttendanceEventType{
    {name: "AttendanceCreated", ordinal: 0},
    {name: "AttendanceDeleted", ordinal: 1},
    {name: "AttendanceUpdated", ordinal: 2}},
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

func (o *attendanceEventTypes) AttendanceCreated() *AttendanceEventType {
    return _attendanceEventTypes.values[0]
}

func (o *attendanceEventTypes) AttendanceDeleted() *AttendanceEventType {
    return _attendanceEventTypes.values[1]
}

func (o *attendanceEventTypes) AttendanceUpdated() *AttendanceEventType {
    return _attendanceEventTypes.values[2]
}

func (o *attendanceEventTypes) ParseAttendanceEventType(name string) (ret *AttendanceEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
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

func (o *CourseEventType) IsCourseCreated() bool {
    return o == _courseEventTypes.CourseCreated()
}

func (o *CourseEventType) IsCourseDeleted() bool {
    return o == _courseEventTypes.CourseDeleted()
}

func (o *CourseEventType) IsCourseUpdated() bool {
    return o == _courseEventTypes.CourseUpdated()
}

type courseEventTypes struct {
	values []*CourseEventType
    literals []enum.Literal
}

var _courseEventTypes = &courseEventTypes{values: []*CourseEventType{
    {name: "CourseCreated", ordinal: 0},
    {name: "CourseDeleted", ordinal: 1},
    {name: "CourseUpdated", ordinal: 2}},
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

func (o *courseEventTypes) CourseCreated() *CourseEventType {
    return _courseEventTypes.values[0]
}

func (o *courseEventTypes) CourseDeleted() *CourseEventType {
    return _courseEventTypes.values[1]
}

func (o *courseEventTypes) CourseUpdated() *CourseEventType {
    return _courseEventTypes.values[2]
}

func (o *courseEventTypes) ParseCourseEventType(name string) (ret *CourseEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
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

func (o *GradeEventType) IsGradeCreated() bool {
    return o == _gradeEventTypes.GradeCreated()
}

func (o *GradeEventType) IsGradeDeleted() bool {
    return o == _gradeEventTypes.GradeDeleted()
}

func (o *GradeEventType) IsGradeUpdated() bool {
    return o == _gradeEventTypes.GradeUpdated()
}

type gradeEventTypes struct {
	values []*GradeEventType
    literals []enum.Literal
}

var _gradeEventTypes = &gradeEventTypes{values: []*GradeEventType{
    {name: "GradeCreated", ordinal: 0},
    {name: "GradeDeleted", ordinal: 1},
    {name: "GradeUpdated", ordinal: 2}},
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

func (o *gradeEventTypes) GradeCreated() *GradeEventType {
    return _gradeEventTypes.values[0]
}

func (o *gradeEventTypes) GradeDeleted() *GradeEventType {
    return _gradeEventTypes.values[1]
}

func (o *gradeEventTypes) GradeUpdated() *GradeEventType {
    return _gradeEventTypes.values[2]
}

func (o *gradeEventTypes) ParseGradeEventType(name string) (ret *GradeEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
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

func (o *GroupEventType) IsGroupCreated() bool {
    return o == _groupEventTypes.GroupCreated()
}

func (o *GroupEventType) IsGroupDeleted() bool {
    return o == _groupEventTypes.GroupDeleted()
}

func (o *GroupEventType) IsGroupUpdated() bool {
    return o == _groupEventTypes.GroupUpdated()
}

type groupEventTypes struct {
	values []*GroupEventType
    literals []enum.Literal
}

var _groupEventTypes = &groupEventTypes{values: []*GroupEventType{
    {name: "GroupCreated", ordinal: 0},
    {name: "GroupDeleted", ordinal: 1},
    {name: "GroupUpdated", ordinal: 2}},
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

func (o *groupEventTypes) GroupCreated() *GroupEventType {
    return _groupEventTypes.values[0]
}

func (o *groupEventTypes) GroupDeleted() *GroupEventType {
    return _groupEventTypes.values[1]
}

func (o *groupEventTypes) GroupUpdated() *GroupEventType {
    return _groupEventTypes.values[2]
}

func (o *groupEventTypes) ParseGroupEventType(name string) (ret *GroupEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
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

func (o *SchoolApplicationEventType) IsSchoolApplicationCreated() bool {
    return o == _schoolApplicationEventTypes.SchoolApplicationCreated()
}

func (o *SchoolApplicationEventType) IsSchoolApplicationDeleted() bool {
    return o == _schoolApplicationEventTypes.SchoolApplicationDeleted()
}

func (o *SchoolApplicationEventType) IsSchoolApplicationUpdated() bool {
    return o == _schoolApplicationEventTypes.SchoolApplicationUpdated()
}

type schoolApplicationEventTypes struct {
	values []*SchoolApplicationEventType
    literals []enum.Literal
}

var _schoolApplicationEventTypes = &schoolApplicationEventTypes{values: []*SchoolApplicationEventType{
    {name: "SchoolApplicationCreated", ordinal: 0},
    {name: "SchoolApplicationDeleted", ordinal: 1},
    {name: "SchoolApplicationUpdated", ordinal: 2}},
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

func (o *schoolApplicationEventTypes) SchoolApplicationCreated() *SchoolApplicationEventType {
    return _schoolApplicationEventTypes.values[0]
}

func (o *schoolApplicationEventTypes) SchoolApplicationDeleted() *SchoolApplicationEventType {
    return _schoolApplicationEventTypes.values[1]
}

func (o *schoolApplicationEventTypes) SchoolApplicationUpdated() *SchoolApplicationEventType {
    return _schoolApplicationEventTypes.values[2]
}

func (o *schoolApplicationEventTypes) ParseSchoolApplicationEventType(name string) (ret *SchoolApplicationEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
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

func (o *SchoolYearEventType) IsSchoolYearCreated() bool {
    return o == _schoolYearEventTypes.SchoolYearCreated()
}

func (o *SchoolYearEventType) IsSchoolYearDeleted() bool {
    return o == _schoolYearEventTypes.SchoolYearDeleted()
}

func (o *SchoolYearEventType) IsSchoolYearUpdated() bool {
    return o == _schoolYearEventTypes.SchoolYearUpdated()
}

type schoolYearEventTypes struct {
	values []*SchoolYearEventType
    literals []enum.Literal
}

var _schoolYearEventTypes = &schoolYearEventTypes{values: []*SchoolYearEventType{
    {name: "SchoolYearCreated", ordinal: 0},
    {name: "SchoolYearDeleted", ordinal: 1},
    {name: "SchoolYearUpdated", ordinal: 2}},
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

func (o *schoolYearEventTypes) SchoolYearCreated() *SchoolYearEventType {
    return _schoolYearEventTypes.values[0]
}

func (o *schoolYearEventTypes) SchoolYearDeleted() *SchoolYearEventType {
    return _schoolYearEventTypes.values[1]
}

func (o *schoolYearEventTypes) SchoolYearUpdated() *SchoolYearEventType {
    return _schoolYearEventTypes.values[2]
}

func (o *schoolYearEventTypes) ParseSchoolYearEventType(name string) (ret *SchoolYearEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*SchoolYearEventType), ok
	}
	return
}



