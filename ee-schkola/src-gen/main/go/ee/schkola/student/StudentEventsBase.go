package student

import (
    "ee/schkola"
    "ee/schkola/person"
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/enum"
    "time"
)
const (
     AttendanceCreatedEvent eventhorizon.EventType = "AttendanceCreated"
     AttendanceDeletedEvent eventhorizon.EventType = "AttendanceDeleted"
     AttendanceUpdatedEvent eventhorizon.EventType = "AttendanceUpdated"
)


const (
     CourseCreatedEvent eventhorizon.EventType = "CourseCreated"
     CourseDeletedEvent eventhorizon.EventType = "CourseDeleted"
     CourseUpdatedEvent eventhorizon.EventType = "CourseUpdated"
)


const (
     GradeCreatedEvent eventhorizon.EventType = "GradeCreated"
     GradeDeletedEvent eventhorizon.EventType = "GradeDeleted"
     GradeUpdatedEvent eventhorizon.EventType = "GradeUpdated"
)


const (
     GroupCreatedEvent eventhorizon.EventType = "GroupCreated"
     GroupDeletedEvent eventhorizon.EventType = "GroupDeleted"
     GroupUpdatedEvent eventhorizon.EventType = "GroupUpdated"
)


const (
     SchoolApplicationCreatedEvent eventhorizon.EventType = "SchoolApplicationCreated"
     SchoolApplicationDeletedEvent eventhorizon.EventType = "SchoolApplicationDeleted"
     SchoolApplicationUpdatedEvent eventhorizon.EventType = "SchoolApplicationUpdated"
)


const (
     SchoolYearCreatedEvent eventhorizon.EventType = "SchoolYearCreated"
     SchoolYearDeletedEvent eventhorizon.EventType = "SchoolYearDeleted"
     SchoolYearUpdatedEvent eventhorizon.EventType = "SchoolYearUpdated"
)





type AttendanceCreated struct {
    Id  eventhorizon.UUID`eh:"optional"`
    Student  *person.Profile`eh:"optional"`
    Date  *time.Time`eh:"optional"`
    Course  *Course`eh:"optional"`
    Hours  int`eh:"optional"`
    State  *AttendanceState`eh:"optional"`
    StateTrace  *schkola.Trace`eh:"optional"`
    Token  string`eh:"optional"`
    TokenTrace  *schkola.Trace`eh:"optional"`
}

func NewCreated() (ret *AttendanceCreated) {
    ret = &AttendanceCreated{}
    return
}



type AttendanceDeleted struct {
    Id  eventhorizon.UUID
}

func NewDeleted() (ret *AttendanceDeleted) {
    ret = &AttendanceDeleted{}
    return
}



type AttendanceUpdated struct {
    Id  eventhorizon.UUID`eh:"optional"`
    Student  *person.Profile`eh:"optional"`
    Date  *time.Time`eh:"optional"`
    Course  *Course`eh:"optional"`
    Hours  int`eh:"optional"`
    State  *AttendanceState`eh:"optional"`
    StateTrace  *schkola.Trace`eh:"optional"`
    Token  string`eh:"optional"`
    TokenTrace  *schkola.Trace`eh:"optional"`
}

func NewUpdated() (ret *AttendanceUpdated) {
    ret = &AttendanceUpdated{}
    return
}



type CourseCreated struct {
    Id  eventhorizon.UUID`eh:"optional"`
    Name  string`eh:"optional"`
    Begin  *time.Time`eh:"optional"`
    End  *time.Time`eh:"optional"`
    Teacher  *person.PersonName`eh:"optional"`
    SchoolYear  *SchoolYear`eh:"optional"`
    Fee  float64`eh:"optional"`
    Description  string`eh:"optional"`
}

func NewCreated() (ret *CourseCreated) {
    ret = &CourseCreated{}
    return
}



type CourseDeleted struct {
    Id  eventhorizon.UUID
}

func NewDeleted() (ret *CourseDeleted) {
    ret = &CourseDeleted{}
    return
}



type CourseUpdated struct {
    Id  eventhorizon.UUID`eh:"optional"`
    Name  string`eh:"optional"`
    Begin  *time.Time`eh:"optional"`
    End  *time.Time`eh:"optional"`
    Teacher  *person.PersonName`eh:"optional"`
    SchoolYear  *SchoolYear`eh:"optional"`
    Fee  float64`eh:"optional"`
    Description  string`eh:"optional"`
}

func NewUpdated() (ret *CourseUpdated) {
    ret = &CourseUpdated{}
    return
}



type GradeCreated struct {
    Id  eventhorizon.UUID`eh:"optional"`
    Student  *person.Profile`eh:"optional"`
    Course  *Course`eh:"optional"`
    Grade  float64`eh:"optional"`
    GradeTrace  *schkola.Trace`eh:"optional"`
    Comment  string`eh:"optional"`
}

func NewCreated() (ret *GradeCreated) {
    ret = &GradeCreated{}
    return
}



type GradeDeleted struct {
    Id  eventhorizon.UUID
}

func NewDeleted() (ret *GradeDeleted) {
    ret = &GradeDeleted{}
    return
}



type GradeUpdated struct {
    Id  eventhorizon.UUID`eh:"optional"`
    Student  *person.Profile`eh:"optional"`
    Course  *Course`eh:"optional"`
    Grade  float64`eh:"optional"`
    GradeTrace  *schkola.Trace`eh:"optional"`
    Comment  string`eh:"optional"`
}

func NewUpdated() (ret *GradeUpdated) {
    ret = &GradeUpdated{}
    return
}



type GroupCreated struct {
    Id  eventhorizon.UUID`eh:"optional"`
    Name  string`eh:"optional"`
    Category  *GroupCategory`eh:"optional"`
    SchoolYear  *SchoolYear`eh:"optional"`
    Representative  *person.Profile`eh:"optional"`
    Students  []*Course`eh:"optional"`
    Courses  []*Course`eh:"optional"`
}

func NewCreated() (ret *GroupCreated) {
    ret = &GroupCreated{}
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
    Id  eventhorizon.UUID
}

func NewDeleted() (ret *GroupDeleted) {
    ret = &GroupDeleted{}
    return
}



type GroupUpdated struct {
    Id  eventhorizon.UUID`eh:"optional"`
    Name  string`eh:"optional"`
    Category  *GroupCategory`eh:"optional"`
    SchoolYear  *SchoolYear`eh:"optional"`
    Representative  *person.Profile`eh:"optional"`
    Students  []*Course`eh:"optional"`
    Courses  []*Course`eh:"optional"`
}

func NewUpdated() (ret *GroupUpdated) {
    ret = &GroupUpdated{}
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
    Id  eventhorizon.UUID`eh:"optional"`
    Profile  *person.Profile`eh:"optional"`
    RecommendationOf  *person.PersonName`eh:"optional"`
    ChurchContactPerson  *person.PersonName`eh:"optional"`
    ChurchContact  *person.Contact`eh:"optional"`
    SchoolYear  *SchoolYear`eh:"optional"`
    Group  string`eh:"optional"`
}

func NewCreated() (ret *SchoolApplicationCreated) {
    ret = &SchoolApplicationCreated{}
    return
}



type SchoolApplicationDeleted struct {
    Id  eventhorizon.UUID
}

func NewDeleted() (ret *SchoolApplicationDeleted) {
    ret = &SchoolApplicationDeleted{}
    return
}



type SchoolApplicationUpdated struct {
    Id  eventhorizon.UUID`eh:"optional"`
    Profile  *person.Profile`eh:"optional"`
    RecommendationOf  *person.PersonName`eh:"optional"`
    ChurchContactPerson  *person.PersonName`eh:"optional"`
    ChurchContact  *person.Contact`eh:"optional"`
    SchoolYear  *SchoolYear`eh:"optional"`
    Group  string`eh:"optional"`
}

func NewUpdated() (ret *SchoolApplicationUpdated) {
    ret = &SchoolApplicationUpdated{}
    return
}



type SchoolYearCreated struct {
    Id  eventhorizon.UUID`eh:"optional"`
    Name  string`eh:"optional"`
    Start  *time.Time`eh:"optional"`
    End  *time.Time`eh:"optional"`
    Dates  []*Course`eh:"optional"`
}

func NewCreated() (ret *SchoolYearCreated) {
    ret = &SchoolYearCreated{}
    return
}

func (o *SchoolYearCreated) AddToDates(item *Course) *Course {
    o.Dates  = append(o.Dates , item)
    return item
}



type SchoolYearDeleted struct {
    Id  eventhorizon.UUID
}

func NewDeleted() (ret *SchoolYearDeleted) {
    ret = &SchoolYearDeleted{}
    return
}



type SchoolYearUpdated struct {
    Id  eventhorizon.UUID`eh:"optional"`
    Name  string`eh:"optional"`
    Start  *time.Time`eh:"optional"`
    End  *time.Time`eh:"optional"`
    Dates  []*Course`eh:"optional"`
}

func NewUpdated() (ret *SchoolYearUpdated) {
    ret = &SchoolYearUpdated{}
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



