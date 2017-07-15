package student

import (
    "ee/schkola"
    "ee/schkola/person"
    "github.com/eugeis/gee/enum"
    "time"
)

type CreatedAttendance struct {
    Student  *person.Profile
    Date  *time.Time
    Course  *Course
    Hours  int
    State  *AttendanceState
    StateTrace  *schkola.Trace
    Token  string
    TokenTrace  *schkola.Trace
}

func NewCreatedAttendance(student *person.Profile, date *time.Time, course *Course, hours int, state *AttendanceState, stateTrace *schkola.Trace, 
                token string, tokenTrace *schkola.Trace) (ret *CreatedAttendance, err error) {
    ret = &CreatedAttendance{
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



type DeletedAttendance struct {
    Id  string
}

func NewDeletedAttendance(id string) (ret *DeletedAttendance, err error) {
    ret = &DeletedAttendance{
        Id : id,
    }
    return
}



type UpdatedAttendance struct {
    Student  *person.Profile
    Date  *time.Time
    Course  *Course
    Hours  int
    State  *AttendanceState
    StateTrace  *schkola.Trace
    Token  string
    TokenTrace  *schkola.Trace
}

func NewUpdatedAttendance(student *person.Profile, date *time.Time, course *Course, hours int, state *AttendanceState, stateTrace *schkola.Trace, 
                token string, tokenTrace *schkola.Trace) (ret *UpdatedAttendance, err error) {
    ret = &UpdatedAttendance{
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



type CreatedCourse struct {
    Name  string
    Begin  *time.Time
    End  *time.Time
    Teacher  *person.PersonName
    SchoolYear  *SchoolYear
    Fee  float64
    Description  string
}

func NewCreatedCourse(name string, begin *time.Time, end *time.Time, teacher *person.PersonName, schoolYear *SchoolYear, fee float64, 
                description string) (ret *CreatedCourse, err error) {
    ret = &CreatedCourse{
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



type DeletedCourse struct {
    Id  string
}

func NewDeletedCourse(id string) (ret *DeletedCourse, err error) {
    ret = &DeletedCourse{
        Id : id,
    }
    return
}



type UpdatedCourse struct {
    Name  string
    Begin  *time.Time
    End  *time.Time
    Teacher  *person.PersonName
    SchoolYear  *SchoolYear
    Fee  float64
    Description  string
}

func NewUpdatedCourse(name string, begin *time.Time, end *time.Time, teacher *person.PersonName, schoolYear *SchoolYear, fee float64, 
                description string) (ret *UpdatedCourse, err error) {
    ret = &UpdatedCourse{
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



type CreatedGrade struct {
    Student  *person.Profile
    Course  *Course
    Grade  float64
    GradeTrace  *schkola.Trace
    Comment  string
}

func NewCreatedGrade(student *person.Profile, course *Course, grade float64, gradeTrace *schkola.Trace, comment string) (ret *CreatedGrade, err error) {
    ret = &CreatedGrade{
        Student : student,
        Course : course,
        Grade : grade,
        GradeTrace : gradeTrace,
        Comment : comment,
    }
    return
}



type DeletedGrade struct {
    Id  string
}

func NewDeletedGrade(id string) (ret *DeletedGrade, err error) {
    ret = &DeletedGrade{
        Id : id,
    }
    return
}



type UpdatedGrade struct {
    Student  *person.Profile
    Course  *Course
    Grade  float64
    GradeTrace  *schkola.Trace
    Comment  string
}

func NewUpdatedGrade(student *person.Profile, course *Course, grade float64, gradeTrace *schkola.Trace, comment string) (ret *UpdatedGrade, err error) {
    ret = &UpdatedGrade{
        Student : student,
        Course : course,
        Grade : grade,
        GradeTrace : gradeTrace,
        Comment : comment,
    }
    return
}



type CreatedGroup struct {
    Name  string
    Category  *GroupCategory
    SchoolYear  *SchoolYear
    Representative  *person.Profile
    Students  []*Course
    Courses  []*Course
}

func NewCreatedGroup(name string, category *GroupCategory, schoolYear *SchoolYear, representative *person.Profile, students []*Course, 
                courses []*Course) (ret *CreatedGroup, err error) {
    ret = &CreatedGroup{
        Name : name,
        Category : category,
        SchoolYear : schoolYear,
        Representative : representative,
        Students : students,
        Courses : courses,
    }
    return
}

func (o *CreatedGroup) AddToStudents(item *Course) *Course {
    o.Students  = append(o.Students , item)
    return item
}

func (o *CreatedGroup) AddToCourses(item *Course) *Course {
    o.Courses  = append(o.Courses , item)
    return item
}



type DeletedGroup struct {
    Id  string
}

func NewDeletedGroup(id string) (ret *DeletedGroup, err error) {
    ret = &DeletedGroup{
        Id : id,
    }
    return
}



type UpdatedGroup struct {
    Name  string
    Category  *GroupCategory
    SchoolYear  *SchoolYear
    Representative  *person.Profile
    Students  []*Course
    Courses  []*Course
}

func NewUpdatedGroup(name string, category *GroupCategory, schoolYear *SchoolYear, representative *person.Profile, students []*Course, 
                courses []*Course) (ret *UpdatedGroup, err error) {
    ret = &UpdatedGroup{
        Name : name,
        Category : category,
        SchoolYear : schoolYear,
        Representative : representative,
        Students : students,
        Courses : courses,
    }
    return
}

func (o *UpdatedGroup) AddToStudents(item *Course) *Course {
    o.Students  = append(o.Students , item)
    return item
}

func (o *UpdatedGroup) AddToCourses(item *Course) *Course {
    o.Courses  = append(o.Courses , item)
    return item
}



type CreatedSchoolApplication struct {
    Profile  *person.Profile
    RecommendationOf  *person.PersonName
    ChurchContactPerson  *person.PersonName
    ChurchContact  *person.Contact
    SchoolYear  *SchoolYear
    Group  string
}

func NewCreatedSchoolApplication(profile *person.Profile, recommendationOf *person.PersonName, churchContactPerson *person.PersonName, 
                churchContact *person.Contact, schoolYear *SchoolYear, group string) (ret *CreatedSchoolApplication, err error) {
    ret = &CreatedSchoolApplication{
        Profile : profile,
        RecommendationOf : recommendationOf,
        ChurchContactPerson : churchContactPerson,
        ChurchContact : churchContact,
        SchoolYear : schoolYear,
        Group : group,
    }
    return
}



type DeletedSchoolApplication struct {
    Id  string
}

func NewDeletedSchoolApplication(id string) (ret *DeletedSchoolApplication, err error) {
    ret = &DeletedSchoolApplication{
        Id : id,
    }
    return
}



type UpdatedSchoolApplication struct {
    Profile  *person.Profile
    RecommendationOf  *person.PersonName
    ChurchContactPerson  *person.PersonName
    ChurchContact  *person.Contact
    SchoolYear  *SchoolYear
    Group  string
}

func NewUpdatedSchoolApplication(profile *person.Profile, recommendationOf *person.PersonName, churchContactPerson *person.PersonName, 
                churchContact *person.Contact, schoolYear *SchoolYear, group string) (ret *UpdatedSchoolApplication, err error) {
    ret = &UpdatedSchoolApplication{
        Profile : profile,
        RecommendationOf : recommendationOf,
        ChurchContactPerson : churchContactPerson,
        ChurchContact : churchContact,
        SchoolYear : schoolYear,
        Group : group,
    }
    return
}



type CreatedSchoolYear struct {
    Name  string
    Start  *time.Time
    End  *time.Time
    Dates  []*Course
}

func NewCreatedSchoolYear(name string, start *time.Time, end *time.Time, dates []*Course) (ret *CreatedSchoolYear, err error) {
    ret = &CreatedSchoolYear{
        Name : name,
        Start : start,
        End : end,
        Dates : dates,
    }
    return
}

func (o *CreatedSchoolYear) AddToDates(item *Course) *Course {
    o.Dates  = append(o.Dates , item)
    return item
}



type DeletedSchoolYear struct {
    Id  string
}

func NewDeletedSchoolYear(id string) (ret *DeletedSchoolYear, err error) {
    ret = &DeletedSchoolYear{
        Id : id,
    }
    return
}



type UpdatedSchoolYear struct {
    Name  string
    Start  *time.Time
    End  *time.Time
    Dates  []*Course
}

func NewUpdatedSchoolYear(name string, start *time.Time, end *time.Time, dates []*Course) (ret *UpdatedSchoolYear, err error) {
    ret = &UpdatedSchoolYear{
        Name : name,
        Start : start,
        End : end,
        Dates : dates,
    }
    return
}

func (o *UpdatedSchoolYear) AddToDates(item *Course) *Course {
    o.Dates  = append(o.Dates , item)
    return item
}




type AttendanceAggregateEventType struct {
	name  string
	ordinal int
}

func (o *AttendanceAggregateEventType) Name() string {
    return o.name
}

func (o *AttendanceAggregateEventType) Ordinal() int {
    return o.ordinal
}

func (o *AttendanceAggregateEventType) IsAttendanceCreated() bool {
    return o == _attendanceAggregateEventTypes.AttendanceCreated()
}

func (o *AttendanceAggregateEventType) IsAttendanceDeleted() bool {
    return o == _attendanceAggregateEventTypes.AttendanceDeleted()
}

func (o *AttendanceAggregateEventType) IsAttendanceUpdated() bool {
    return o == _attendanceAggregateEventTypes.AttendanceUpdated()
}

type attendanceAggregateEventTypes struct {
	values []*AttendanceAggregateEventType
    literals []enum.Literal
}

var _attendanceAggregateEventTypes = &attendanceAggregateEventTypes{values: []*AttendanceAggregateEventType{
    {name: "AttendanceCreated", ordinal: 0},
    {name: "AttendanceDeleted", ordinal: 1},
    {name: "AttendanceUpdated", ordinal: 2}},
}

func AttendanceAggregateEventTypes() *attendanceAggregateEventTypes {
	return _attendanceAggregateEventTypes
}

func (o *attendanceAggregateEventTypes) Values() []*AttendanceAggregateEventType {
	return o.values
}

func (o *attendanceAggregateEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *attendanceAggregateEventTypes) AttendanceCreated() *AttendanceAggregateEventType {
    return _attendanceAggregateEventTypes.values[0]
}

func (o *attendanceAggregateEventTypes) AttendanceDeleted() *AttendanceAggregateEventType {
    return _attendanceAggregateEventTypes.values[1]
}

func (o *attendanceAggregateEventTypes) AttendanceUpdated() *AttendanceAggregateEventType {
    return _attendanceAggregateEventTypes.values[2]
}

func (o *attendanceAggregateEventTypes) ParseAttendanceAggregateEventType(name string) (ret *AttendanceAggregateEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*AttendanceAggregateEventType), ok
	}
	return
}


type CourseAggregateEventType struct {
	name  string
	ordinal int
}

func (o *CourseAggregateEventType) Name() string {
    return o.name
}

func (o *CourseAggregateEventType) Ordinal() int {
    return o.ordinal
}

func (o *CourseAggregateEventType) IsCourseCreated() bool {
    return o == _courseAggregateEventTypes.CourseCreated()
}

func (o *CourseAggregateEventType) IsCourseDeleted() bool {
    return o == _courseAggregateEventTypes.CourseDeleted()
}

func (o *CourseAggregateEventType) IsCourseUpdated() bool {
    return o == _courseAggregateEventTypes.CourseUpdated()
}

type courseAggregateEventTypes struct {
	values []*CourseAggregateEventType
    literals []enum.Literal
}

var _courseAggregateEventTypes = &courseAggregateEventTypes{values: []*CourseAggregateEventType{
    {name: "CourseCreated", ordinal: 0},
    {name: "CourseDeleted", ordinal: 1},
    {name: "CourseUpdated", ordinal: 2}},
}

func CourseAggregateEventTypes() *courseAggregateEventTypes {
	return _courseAggregateEventTypes
}

func (o *courseAggregateEventTypes) Values() []*CourseAggregateEventType {
	return o.values
}

func (o *courseAggregateEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *courseAggregateEventTypes) CourseCreated() *CourseAggregateEventType {
    return _courseAggregateEventTypes.values[0]
}

func (o *courseAggregateEventTypes) CourseDeleted() *CourseAggregateEventType {
    return _courseAggregateEventTypes.values[1]
}

func (o *courseAggregateEventTypes) CourseUpdated() *CourseAggregateEventType {
    return _courseAggregateEventTypes.values[2]
}

func (o *courseAggregateEventTypes) ParseCourseAggregateEventType(name string) (ret *CourseAggregateEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*CourseAggregateEventType), ok
	}
	return
}


type GradeAggregateEventType struct {
	name  string
	ordinal int
}

func (o *GradeAggregateEventType) Name() string {
    return o.name
}

func (o *GradeAggregateEventType) Ordinal() int {
    return o.ordinal
}

func (o *GradeAggregateEventType) IsGradeCreated() bool {
    return o == _gradeAggregateEventTypes.GradeCreated()
}

func (o *GradeAggregateEventType) IsGradeDeleted() bool {
    return o == _gradeAggregateEventTypes.GradeDeleted()
}

func (o *GradeAggregateEventType) IsGradeUpdated() bool {
    return o == _gradeAggregateEventTypes.GradeUpdated()
}

type gradeAggregateEventTypes struct {
	values []*GradeAggregateEventType
    literals []enum.Literal
}

var _gradeAggregateEventTypes = &gradeAggregateEventTypes{values: []*GradeAggregateEventType{
    {name: "GradeCreated", ordinal: 0},
    {name: "GradeDeleted", ordinal: 1},
    {name: "GradeUpdated", ordinal: 2}},
}

func GradeAggregateEventTypes() *gradeAggregateEventTypes {
	return _gradeAggregateEventTypes
}

func (o *gradeAggregateEventTypes) Values() []*GradeAggregateEventType {
	return o.values
}

func (o *gradeAggregateEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *gradeAggregateEventTypes) GradeCreated() *GradeAggregateEventType {
    return _gradeAggregateEventTypes.values[0]
}

func (o *gradeAggregateEventTypes) GradeDeleted() *GradeAggregateEventType {
    return _gradeAggregateEventTypes.values[1]
}

func (o *gradeAggregateEventTypes) GradeUpdated() *GradeAggregateEventType {
    return _gradeAggregateEventTypes.values[2]
}

func (o *gradeAggregateEventTypes) ParseGradeAggregateEventType(name string) (ret *GradeAggregateEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*GradeAggregateEventType), ok
	}
	return
}


type GroupAggregateEventType struct {
	name  string
	ordinal int
}

func (o *GroupAggregateEventType) Name() string {
    return o.name
}

func (o *GroupAggregateEventType) Ordinal() int {
    return o.ordinal
}

func (o *GroupAggregateEventType) IsGroupCreated() bool {
    return o == _groupAggregateEventTypes.GroupCreated()
}

func (o *GroupAggregateEventType) IsGroupDeleted() bool {
    return o == _groupAggregateEventTypes.GroupDeleted()
}

func (o *GroupAggregateEventType) IsGroupUpdated() bool {
    return o == _groupAggregateEventTypes.GroupUpdated()
}

type groupAggregateEventTypes struct {
	values []*GroupAggregateEventType
    literals []enum.Literal
}

var _groupAggregateEventTypes = &groupAggregateEventTypes{values: []*GroupAggregateEventType{
    {name: "GroupCreated", ordinal: 0},
    {name: "GroupDeleted", ordinal: 1},
    {name: "GroupUpdated", ordinal: 2}},
}

func GroupAggregateEventTypes() *groupAggregateEventTypes {
	return _groupAggregateEventTypes
}

func (o *groupAggregateEventTypes) Values() []*GroupAggregateEventType {
	return o.values
}

func (o *groupAggregateEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *groupAggregateEventTypes) GroupCreated() *GroupAggregateEventType {
    return _groupAggregateEventTypes.values[0]
}

func (o *groupAggregateEventTypes) GroupDeleted() *GroupAggregateEventType {
    return _groupAggregateEventTypes.values[1]
}

func (o *groupAggregateEventTypes) GroupUpdated() *GroupAggregateEventType {
    return _groupAggregateEventTypes.values[2]
}

func (o *groupAggregateEventTypes) ParseGroupAggregateEventType(name string) (ret *GroupAggregateEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*GroupAggregateEventType), ok
	}
	return
}


type SchoolApplicationAggregateEventType struct {
	name  string
	ordinal int
}

func (o *SchoolApplicationAggregateEventType) Name() string {
    return o.name
}

func (o *SchoolApplicationAggregateEventType) Ordinal() int {
    return o.ordinal
}

func (o *SchoolApplicationAggregateEventType) IsSchoolApplicationCreated() bool {
    return o == _schoolApplicationAggregateEventTypes.SchoolApplicationCreated()
}

func (o *SchoolApplicationAggregateEventType) IsSchoolApplicationDeleted() bool {
    return o == _schoolApplicationAggregateEventTypes.SchoolApplicationDeleted()
}

func (o *SchoolApplicationAggregateEventType) IsSchoolApplicationUpdated() bool {
    return o == _schoolApplicationAggregateEventTypes.SchoolApplicationUpdated()
}

type schoolApplicationAggregateEventTypes struct {
	values []*SchoolApplicationAggregateEventType
    literals []enum.Literal
}

var _schoolApplicationAggregateEventTypes = &schoolApplicationAggregateEventTypes{values: []*SchoolApplicationAggregateEventType{
    {name: "SchoolApplicationCreated", ordinal: 0},
    {name: "SchoolApplicationDeleted", ordinal: 1},
    {name: "SchoolApplicationUpdated", ordinal: 2}},
}

func SchoolApplicationAggregateEventTypes() *schoolApplicationAggregateEventTypes {
	return _schoolApplicationAggregateEventTypes
}

func (o *schoolApplicationAggregateEventTypes) Values() []*SchoolApplicationAggregateEventType {
	return o.values
}

func (o *schoolApplicationAggregateEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *schoolApplicationAggregateEventTypes) SchoolApplicationCreated() *SchoolApplicationAggregateEventType {
    return _schoolApplicationAggregateEventTypes.values[0]
}

func (o *schoolApplicationAggregateEventTypes) SchoolApplicationDeleted() *SchoolApplicationAggregateEventType {
    return _schoolApplicationAggregateEventTypes.values[1]
}

func (o *schoolApplicationAggregateEventTypes) SchoolApplicationUpdated() *SchoolApplicationAggregateEventType {
    return _schoolApplicationAggregateEventTypes.values[2]
}

func (o *schoolApplicationAggregateEventTypes) ParseSchoolApplicationAggregateEventType(name string) (ret *SchoolApplicationAggregateEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*SchoolApplicationAggregateEventType), ok
	}
	return
}


type SchoolYearAggregateEventType struct {
	name  string
	ordinal int
}

func (o *SchoolYearAggregateEventType) Name() string {
    return o.name
}

func (o *SchoolYearAggregateEventType) Ordinal() int {
    return o.ordinal
}

func (o *SchoolYearAggregateEventType) IsSchoolYearCreated() bool {
    return o == _schoolYearAggregateEventTypes.SchoolYearCreated()
}

func (o *SchoolYearAggregateEventType) IsSchoolYearDeleted() bool {
    return o == _schoolYearAggregateEventTypes.SchoolYearDeleted()
}

func (o *SchoolYearAggregateEventType) IsSchoolYearUpdated() bool {
    return o == _schoolYearAggregateEventTypes.SchoolYearUpdated()
}

type schoolYearAggregateEventTypes struct {
	values []*SchoolYearAggregateEventType
    literals []enum.Literal
}

var _schoolYearAggregateEventTypes = &schoolYearAggregateEventTypes{values: []*SchoolYearAggregateEventType{
    {name: "SchoolYearCreated", ordinal: 0},
    {name: "SchoolYearDeleted", ordinal: 1},
    {name: "SchoolYearUpdated", ordinal: 2}},
}

func SchoolYearAggregateEventTypes() *schoolYearAggregateEventTypes {
	return _schoolYearAggregateEventTypes
}

func (o *schoolYearAggregateEventTypes) Values() []*SchoolYearAggregateEventType {
	return o.values
}

func (o *schoolYearAggregateEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *schoolYearAggregateEventTypes) SchoolYearCreated() *SchoolYearAggregateEventType {
    return _schoolYearAggregateEventTypes.values[0]
}

func (o *schoolYearAggregateEventTypes) SchoolYearDeleted() *SchoolYearAggregateEventType {
    return _schoolYearAggregateEventTypes.values[1]
}

func (o *schoolYearAggregateEventTypes) SchoolYearUpdated() *SchoolYearAggregateEventType {
    return _schoolYearAggregateEventTypes.values[2]
}

func (o *schoolYearAggregateEventTypes) ParseSchoolYearAggregateEventType(name string) (ret *SchoolYearAggregateEventType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*SchoolYearAggregateEventType), ok
	}
	return
}



