package student

import (
    "ee/schkola"
    "ee/schkola/person"
    "time"
)
type AttendanceCreated struct {
    Student  *person.Profile
    Date  *time.Time
    Course  *Course
    Hours  int
    State  *AttendanceState
    StateTrace  *schkola.Trace
    Token  string
    TokenTrace  *schkola.Trace
}

func NewAttendanceCreated(student *person.Profile, date *time.Time, course *Course, hours int, state *AttendanceState, stateTrace *schkola.Trace, 
                token string, tokenTrace *schkola.Trace) (ret *AttendanceCreated, err error) {
    ret = &AttendanceCreated{
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


type AttendanceDeleted struct {
    Id  string
}

func NewAttendanceDeleted(id string) (ret *AttendanceDeleted, err error) {
    ret = &AttendanceDeleted{
        Id : id,
    }
    return
}


type AttendanceUpdated struct {
    Student  *person.Profile
    Date  *time.Time
    Course  *Course
    Hours  int
    State  *AttendanceState
    StateTrace  *schkola.Trace
    Token  string
    TokenTrace  *schkola.Trace
}

func NewAttendanceUpdated(student *person.Profile, date *time.Time, course *Course, hours int, state *AttendanceState, stateTrace *schkola.Trace, 
                token string, tokenTrace *schkola.Trace) (ret *AttendanceUpdated, err error) {
    ret = &AttendanceUpdated{
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


type CourseCreated struct {
    Name  string
    Begin  *time.Time
    End  *time.Time
    Teacher  *person.PersonName
    SchoolYear  *SchoolYear
    Fee  float64
    Description  string
}

func NewCourseCreated(name string, begin *time.Time, end *time.Time, teacher *person.PersonName, schoolYear *SchoolYear, fee float64, 
                description string) (ret *CourseCreated, err error) {
    ret = &CourseCreated{
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


type CourseDeleted struct {
    Id  string
}

func NewCourseDeleted(id string) (ret *CourseDeleted, err error) {
    ret = &CourseDeleted{
        Id : id,
    }
    return
}


type CourseUpdated struct {
    Name  string
    Begin  *time.Time
    End  *time.Time
    Teacher  *person.PersonName
    SchoolYear  *SchoolYear
    Fee  float64
    Description  string
}

func NewCourseUpdated(name string, begin *time.Time, end *time.Time, teacher *person.PersonName, schoolYear *SchoolYear, fee float64, 
                description string) (ret *CourseUpdated, err error) {
    ret = &CourseUpdated{
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


type GradeCreated struct {
    Student  *person.Profile
    Course  *Course
    Grade  float64
    GradeTrace  *schkola.Trace
    Comment  string
}

func NewGradeCreated(student *person.Profile, course *Course, grade float64, gradeTrace *schkola.Trace, comment string) (ret *GradeCreated, err error) {
    ret = &GradeCreated{
        Student : student,
        Course : course,
        Grade : grade,
        GradeTrace : gradeTrace,
        Comment : comment,
    }
    return
}


type GradeDeleted struct {
    Id  string
}

func NewGradeDeleted(id string) (ret *GradeDeleted, err error) {
    ret = &GradeDeleted{
        Id : id,
    }
    return
}


type GradeUpdated struct {
    Student  *person.Profile
    Course  *Course
    Grade  float64
    GradeTrace  *schkola.Trace
    Comment  string
}

func NewGradeUpdated(student *person.Profile, course *Course, grade float64, gradeTrace *schkola.Trace, comment string) (ret *GradeUpdated, err error) {
    ret = &GradeUpdated{
        Student : student,
        Course : course,
        Grade : grade,
        GradeTrace : gradeTrace,
        Comment : comment,
    }
    return
}


type GroupCreated struct {
    Name  string
    Category  *GroupCategory
    SchoolYear  *SchoolYear
    Representative  *person.Profile
    Students  []*Course
    Courses  []*Course
}

func NewGroupCreated(name string, category *GroupCategory, schoolYear *SchoolYear, representative *person.Profile, students []*Course, 
                courses []*Course) (ret *GroupCreated, err error) {
    ret = &GroupCreated{
        Name : name,
        Category : category,
        SchoolYear : schoolYear,
        Representative : representative,
        Students : students,
        Courses : courses,
    }
    return
}

func (o *GroupCreated) AddToStudents(item *Course) *Course {
    o.Students  = append(o.Students , item)
    return item
}

func (o *GroupCreated) AddToCourses(item *Course) *Course {
    o.Courses  = append(o.Courses , item)
    return item
}


type GroupDeleted struct {
    Id  string
}

func NewGroupDeleted(id string) (ret *GroupDeleted, err error) {
    ret = &GroupDeleted{
        Id : id,
    }
    return
}


type GroupUpdated struct {
    Name  string
    Category  *GroupCategory
    SchoolYear  *SchoolYear
    Representative  *person.Profile
    Students  []*Course
    Courses  []*Course
}

func NewGroupUpdated(name string, category *GroupCategory, schoolYear *SchoolYear, representative *person.Profile, students []*Course, 
                courses []*Course) (ret *GroupUpdated, err error) {
    ret = &GroupUpdated{
        Name : name,
        Category : category,
        SchoolYear : schoolYear,
        Representative : representative,
        Students : students,
        Courses : courses,
    }
    return
}

func (o *GroupUpdated) AddToStudents(item *Course) *Course {
    o.Students  = append(o.Students , item)
    return item
}

func (o *GroupUpdated) AddToCourses(item *Course) *Course {
    o.Courses  = append(o.Courses , item)
    return item
}


type SchoolApplicationCreated struct {
    Profile  *person.Profile
    RecommendationOf  *person.PersonName
    ChurchContactPerson  *person.PersonName
    ChurchContact  *person.Contact
    SchoolYear  *SchoolYear
    Group  string
}

func NewSchoolApplicationCreated(profile *person.Profile, recommendationOf *person.PersonName, churchContactPerson *person.PersonName, 
                churchContact *person.Contact, schoolYear *SchoolYear, group string) (ret *SchoolApplicationCreated, err error) {
    ret = &SchoolApplicationCreated{
        Profile : profile,
        RecommendationOf : recommendationOf,
        ChurchContactPerson : churchContactPerson,
        ChurchContact : churchContact,
        SchoolYear : schoolYear,
        Group : group,
    }
    return
}


type SchoolApplicationDeleted struct {
    Id  string
}

func NewSchoolApplicationDeleted(id string) (ret *SchoolApplicationDeleted, err error) {
    ret = &SchoolApplicationDeleted{
        Id : id,
    }
    return
}


type SchoolApplicationUpdated struct {
    Profile  *person.Profile
    RecommendationOf  *person.PersonName
    ChurchContactPerson  *person.PersonName
    ChurchContact  *person.Contact
    SchoolYear  *SchoolYear
    Group  string
}

func NewSchoolApplicationUpdated(profile *person.Profile, recommendationOf *person.PersonName, churchContactPerson *person.PersonName, 
                churchContact *person.Contact, schoolYear *SchoolYear, group string) (ret *SchoolApplicationUpdated, err error) {
    ret = &SchoolApplicationUpdated{
        Profile : profile,
        RecommendationOf : recommendationOf,
        ChurchContactPerson : churchContactPerson,
        ChurchContact : churchContact,
        SchoolYear : schoolYear,
        Group : group,
    }
    return
}


type SchoolYearCreated struct {
    Name  string
    Start  *time.Time
    End  *time.Time
    Dates  []*Course
}

func NewSchoolYearCreated(name string, start *time.Time, end *time.Time, dates []*Course) (ret *SchoolYearCreated, err error) {
    ret = &SchoolYearCreated{
        Name : name,
        Start : start,
        End : end,
        Dates : dates,
    }
    return
}

func (o *SchoolYearCreated) AddToDates(item *Course) *Course {
    o.Dates  = append(o.Dates , item)
    return item
}


type SchoolYearDeleted struct {
    Id  string
}

func NewSchoolYearDeleted(id string) (ret *SchoolYearDeleted, err error) {
    ret = &SchoolYearDeleted{
        Id : id,
    }
    return
}


type SchoolYearUpdated struct {
    Name  string
    Start  *time.Time
    End  *time.Time
    Dates  []*Course
}

func NewSchoolYearUpdated(name string, start *time.Time, end *time.Time, dates []*Course) (ret *SchoolYearUpdated, err error) {
    ret = &SchoolYearUpdated{
        Name : name,
        Start : start,
        End : end,
        Dates : dates,
    }
    return
}

func (o *SchoolYearUpdated) AddToDates(item *Course) *Course {
    o.Dates  = append(o.Dates , item)
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

func (o *AttendanceEventType) IsCreatedAttendance() bool {
    return o == _attendanceEventTypes.CreatedAttendance()
}

func (o *AttendanceEventType) IsDeletedAttendance() bool {
    return o == _attendanceEventTypes.DeletedAttendance()
}

func (o *AttendanceEventType) IsUpdatedAttendance() bool {
    return o == _attendanceEventTypes.UpdatedAttendance()
}

type attendanceEventTypes struct {
	values []*AttendanceEventType
}

var _attendanceEventTypes = &attendanceEventTypes{values: []*AttendanceEventType{
    {name: "createdAttendance", ordinal: 0},
    {name: "deletedAttendance", ordinal: 1},
    {name: "updatedAttendance", ordinal: 2}},
}

func AttendanceEventTypes() *attendanceEventTypes {
	return _attendanceEventTypes
}

func (o *attendanceEventTypes) Values() []*AttendanceEventType {
	return o.values
}

func (o *attendanceEventTypes) CreatedAttendance() *AttendanceEventType {
    return _attendanceEventTypes.values[0]
}

func (o *attendanceEventTypes) DeletedAttendance() *AttendanceEventType {
    return _attendanceEventTypes.values[1]
}

func (o *attendanceEventTypes) UpdatedAttendance() *AttendanceEventType {
    return _attendanceEventTypes.values[2]
}

func (o *attendanceEventTypes) ParseAttendanceEventType(name string) (ret *AttendanceEventType, ok bool) {
    switch name {
      case o.CreatedAttendance().Name():
        ret = o.CreatedAttendance()
      case o.DeletedAttendance().Name():
        ret = o.DeletedAttendance()
      case o.UpdatedAttendance().Name():
        ret = o.UpdatedAttendance()
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

func (o *CourseEventType) IsCreatedCourse() bool {
    return o == _courseEventTypes.CreatedCourse()
}

func (o *CourseEventType) IsDeletedCourse() bool {
    return o == _courseEventTypes.DeletedCourse()
}

func (o *CourseEventType) IsUpdatedCourse() bool {
    return o == _courseEventTypes.UpdatedCourse()
}

type courseEventTypes struct {
	values []*CourseEventType
}

var _courseEventTypes = &courseEventTypes{values: []*CourseEventType{
    {name: "createdCourse", ordinal: 0},
    {name: "deletedCourse", ordinal: 1},
    {name: "updatedCourse", ordinal: 2}},
}

func CourseEventTypes() *courseEventTypes {
	return _courseEventTypes
}

func (o *courseEventTypes) Values() []*CourseEventType {
	return o.values
}

func (o *courseEventTypes) CreatedCourse() *CourseEventType {
    return _courseEventTypes.values[0]
}

func (o *courseEventTypes) DeletedCourse() *CourseEventType {
    return _courseEventTypes.values[1]
}

func (o *courseEventTypes) UpdatedCourse() *CourseEventType {
    return _courseEventTypes.values[2]
}

func (o *courseEventTypes) ParseCourseEventType(name string) (ret *CourseEventType, ok bool) {
    switch name {
      case o.CreatedCourse().Name():
        ret = o.CreatedCourse()
      case o.DeletedCourse().Name():
        ret = o.DeletedCourse()
      case o.UpdatedCourse().Name():
        ret = o.UpdatedCourse()
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

func (o *GradeEventType) IsCreatedGrade() bool {
    return o == _gradeEventTypes.CreatedGrade()
}

func (o *GradeEventType) IsDeletedGrade() bool {
    return o == _gradeEventTypes.DeletedGrade()
}

func (o *GradeEventType) IsUpdatedGrade() bool {
    return o == _gradeEventTypes.UpdatedGrade()
}

type gradeEventTypes struct {
	values []*GradeEventType
}

var _gradeEventTypes = &gradeEventTypes{values: []*GradeEventType{
    {name: "createdGrade", ordinal: 0},
    {name: "deletedGrade", ordinal: 1},
    {name: "updatedGrade", ordinal: 2}},
}

func GradeEventTypes() *gradeEventTypes {
	return _gradeEventTypes
}

func (o *gradeEventTypes) Values() []*GradeEventType {
	return o.values
}

func (o *gradeEventTypes) CreatedGrade() *GradeEventType {
    return _gradeEventTypes.values[0]
}

func (o *gradeEventTypes) DeletedGrade() *GradeEventType {
    return _gradeEventTypes.values[1]
}

func (o *gradeEventTypes) UpdatedGrade() *GradeEventType {
    return _gradeEventTypes.values[2]
}

func (o *gradeEventTypes) ParseGradeEventType(name string) (ret *GradeEventType, ok bool) {
    switch name {
      case o.CreatedGrade().Name():
        ret = o.CreatedGrade()
      case o.DeletedGrade().Name():
        ret = o.DeletedGrade()
      case o.UpdatedGrade().Name():
        ret = o.UpdatedGrade()
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

func (o *GroupEventType) IsCreatedGroup() bool {
    return o == _groupEventTypes.CreatedGroup()
}

func (o *GroupEventType) IsDeletedGroup() bool {
    return o == _groupEventTypes.DeletedGroup()
}

func (o *GroupEventType) IsUpdatedGroup() bool {
    return o == _groupEventTypes.UpdatedGroup()
}

type groupEventTypes struct {
	values []*GroupEventType
}

var _groupEventTypes = &groupEventTypes{values: []*GroupEventType{
    {name: "createdGroup", ordinal: 0},
    {name: "deletedGroup", ordinal: 1},
    {name: "updatedGroup", ordinal: 2}},
}

func GroupEventTypes() *groupEventTypes {
	return _groupEventTypes
}

func (o *groupEventTypes) Values() []*GroupEventType {
	return o.values
}

func (o *groupEventTypes) CreatedGroup() *GroupEventType {
    return _groupEventTypes.values[0]
}

func (o *groupEventTypes) DeletedGroup() *GroupEventType {
    return _groupEventTypes.values[1]
}

func (o *groupEventTypes) UpdatedGroup() *GroupEventType {
    return _groupEventTypes.values[2]
}

func (o *groupEventTypes) ParseGroupEventType(name string) (ret *GroupEventType, ok bool) {
    switch name {
      case o.CreatedGroup().Name():
        ret = o.CreatedGroup()
      case o.DeletedGroup().Name():
        ret = o.DeletedGroup()
      case o.UpdatedGroup().Name():
        ret = o.UpdatedGroup()
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

func (o *SchoolApplicationEventType) IsCreatedSchoolApplication() bool {
    return o == _schoolApplicationEventTypes.CreatedSchoolApplication()
}

func (o *SchoolApplicationEventType) IsDeletedSchoolApplication() bool {
    return o == _schoolApplicationEventTypes.DeletedSchoolApplication()
}

func (o *SchoolApplicationEventType) IsUpdatedSchoolApplication() bool {
    return o == _schoolApplicationEventTypes.UpdatedSchoolApplication()
}

type schoolApplicationEventTypes struct {
	values []*SchoolApplicationEventType
}

var _schoolApplicationEventTypes = &schoolApplicationEventTypes{values: []*SchoolApplicationEventType{
    {name: "createdSchoolApplication", ordinal: 0},
    {name: "deletedSchoolApplication", ordinal: 1},
    {name: "updatedSchoolApplication", ordinal: 2}},
}

func SchoolApplicationEventTypes() *schoolApplicationEventTypes {
	return _schoolApplicationEventTypes
}

func (o *schoolApplicationEventTypes) Values() []*SchoolApplicationEventType {
	return o.values
}

func (o *schoolApplicationEventTypes) CreatedSchoolApplication() *SchoolApplicationEventType {
    return _schoolApplicationEventTypes.values[0]
}

func (o *schoolApplicationEventTypes) DeletedSchoolApplication() *SchoolApplicationEventType {
    return _schoolApplicationEventTypes.values[1]
}

func (o *schoolApplicationEventTypes) UpdatedSchoolApplication() *SchoolApplicationEventType {
    return _schoolApplicationEventTypes.values[2]
}

func (o *schoolApplicationEventTypes) ParseSchoolApplicationEventType(name string) (ret *SchoolApplicationEventType, ok bool) {
    switch name {
      case o.CreatedSchoolApplication().Name():
        ret = o.CreatedSchoolApplication()
      case o.DeletedSchoolApplication().Name():
        ret = o.DeletedSchoolApplication()
      case o.UpdatedSchoolApplication().Name():
        ret = o.UpdatedSchoolApplication()
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

func (o *SchoolYearEventType) IsCreatedSchoolYear() bool {
    return o == _schoolYearEventTypes.CreatedSchoolYear()
}

func (o *SchoolYearEventType) IsDeletedSchoolYear() bool {
    return o == _schoolYearEventTypes.DeletedSchoolYear()
}

func (o *SchoolYearEventType) IsUpdatedSchoolYear() bool {
    return o == _schoolYearEventTypes.UpdatedSchoolYear()
}

type schoolYearEventTypes struct {
	values []*SchoolYearEventType
}

var _schoolYearEventTypes = &schoolYearEventTypes{values: []*SchoolYearEventType{
    {name: "createdSchoolYear", ordinal: 0},
    {name: "deletedSchoolYear", ordinal: 1},
    {name: "updatedSchoolYear", ordinal: 2}},
}

func SchoolYearEventTypes() *schoolYearEventTypes {
	return _schoolYearEventTypes
}

func (o *schoolYearEventTypes) Values() []*SchoolYearEventType {
	return o.values
}

func (o *schoolYearEventTypes) CreatedSchoolYear() *SchoolYearEventType {
    return _schoolYearEventTypes.values[0]
}

func (o *schoolYearEventTypes) DeletedSchoolYear() *SchoolYearEventType {
    return _schoolYearEventTypes.values[1]
}

func (o *schoolYearEventTypes) UpdatedSchoolYear() *SchoolYearEventType {
    return _schoolYearEventTypes.values[2]
}

func (o *schoolYearEventTypes) ParseSchoolYearEventType(name string) (ret *SchoolYearEventType, ok bool) {
    switch name {
      case o.CreatedSchoolYear().Name():
        ret = o.CreatedSchoolYear()
      case o.DeletedSchoolYear().Name():
        ret = o.DeletedSchoolYear()
      case o.UpdatedSchoolYear().Name():
        ret = o.UpdatedSchoolYear()
    }
    return
}



