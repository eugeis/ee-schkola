package student

import (
    "ee/schkola"
    "ee/schkola/person"
    "time"
)
type CourseCreated struct {
    Name  string
    Begin  *time.Time
    End  *time.Time
    Teacher  *person.PersonName
    SchoolYear  *SchoolYear
    Fee  float64
    Description  string
    *schkola.SchkolaBase
}


type CourseDeleted struct {
    Id  string
}


type CourseUpdated struct {
    Name  string
    Begin  *time.Time
    End  *time.Time
    Teacher  *person.PersonName
    SchoolYear  *SchoolYear
    Fee  float64
    Description  string
    *schkola.SchkolaBase
}


type GradeCreated struct {
    Student  *person.Profile
    Course  *Course
    Grade  float64
    GradeTrace  *schkola.Trace
    Comment  string
    *schkola.SchkolaBase
}


type GradeDeleted struct {
    Id  string
}


type GradeUpdated struct {
    Student  *person.Profile
    Course  *Course
    Grade  float64
    GradeTrace  *schkola.Trace
    Comment  string
    *schkola.SchkolaBase
}


type GroupCreated struct {
    Name  string
    Category  *GroupCategory
    SchoolYear  *SchoolYear
    Representative  *person.Profile
    Students  []*Course
    Courses  []*Course
    *schkola.SchkolaBase
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


type GroupUpdated struct {
    Name  string
    Category  *GroupCategory
    SchoolYear  *SchoolYear
    Representative  *person.Profile
    Students  []*Course
    Courses  []*Course
    *schkola.SchkolaBase
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
    *schkola.SchkolaBase
}


type SchoolApplicationDeleted struct {
    Id  string
}


type SchoolApplicationUpdated struct {
    Profile  *person.Profile
    RecommendationOf  *person.PersonName
    ChurchContactPerson  *person.PersonName
    ChurchContact  *person.Contact
    SchoolYear  *SchoolYear
    Group  string
    *schkola.SchkolaBase
}


type SchoolYearCreated struct {
    Name  string
    Start  *time.Time
    End  *time.Time
    Dates  []*Course
    *schkola.SchkolaBase
}

func (o *SchoolYearCreated) AddToDates(item *Course) *Course {
    o.Dates  = append(o.Dates , item)
    return item
}


type SchoolYearDeleted struct {
    Id  string
}


type SchoolYearUpdated struct {
    Name  string
    Start  *time.Time
    End  *time.Time
    Dates  []*Course
    *schkola.SchkolaBase
}

func (o *SchoolYearUpdated) AddToDates(item *Course) *Course {
    o.Dates  = append(o.Dates , item)
    return item
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
      case "CreatedCourse":
        ret = o.CreatedCourse()
      case "DeletedCourse":
        ret = o.DeletedCourse()
      case "UpdatedCourse":
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
      case "CreatedGrade":
        ret = o.CreatedGrade()
      case "DeletedGrade":
        ret = o.DeletedGrade()
      case "UpdatedGrade":
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
      case "CreatedGroup":
        ret = o.CreatedGroup()
      case "DeletedGroup":
        ret = o.DeletedGroup()
      case "UpdatedGroup":
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
      case "CreatedSchoolApplication":
        ret = o.CreatedSchoolApplication()
      case "DeletedSchoolApplication":
        ret = o.DeletedSchoolApplication()
      case "UpdatedSchoolApplication":
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
      case "CreatedSchoolYear":
        ret = o.CreatedSchoolYear()
      case "DeletedSchoolYear":
        ret = o.DeletedSchoolYear()
      case "UpdatedSchoolYear":
        ret = o.UpdatedSchoolYear()
    }
    return
}



