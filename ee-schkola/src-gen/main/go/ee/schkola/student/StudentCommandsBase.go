package student

import (
    "ee/schkola"
    "ee/schkola/person"
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/enum"
    "time"
)
const (
     CreateAttendanceCommand eventhorizon.CommandType = "CreateAttendance"
     DeleteAttendanceCommand eventhorizon.CommandType = "DeleteAttendance"
     UpdateAttendanceCommand eventhorizon.CommandType = "UpdateAttendance"
     ConfirmAttendanceCommand eventhorizon.CommandType = "ConfirmAttendance"
     CancelAttendanceCommand eventhorizon.CommandType = "CancelAttendance"
     RegisterAttendanceCommand eventhorizon.CommandType = "RegisterAttendance"
)


const (
     CreateCourseCommand eventhorizon.CommandType = "CreateCourse"
     DeleteCourseCommand eventhorizon.CommandType = "DeleteCourse"
     UpdateCourseCommand eventhorizon.CommandType = "UpdateCourse"
)


const (
     CreateGradeCommand eventhorizon.CommandType = "CreateGrade"
     DeleteGradeCommand eventhorizon.CommandType = "DeleteGrade"
     UpdateGradeCommand eventhorizon.CommandType = "UpdateGrade"
)


const (
     CreateGroupCommand eventhorizon.CommandType = "CreateGroup"
     DeleteGroupCommand eventhorizon.CommandType = "DeleteGroup"
     UpdateGroupCommand eventhorizon.CommandType = "UpdateGroup"
)


const (
     CreateSchoolApplicationCommand eventhorizon.CommandType = "CreateSchoolApplication"
     DeleteSchoolApplicationCommand eventhorizon.CommandType = "DeleteSchoolApplication"
     UpdateSchoolApplicationCommand eventhorizon.CommandType = "UpdateSchoolApplication"
)


const (
     CreateSchoolYearCommand eventhorizon.CommandType = "CreateSchoolYear"
     DeleteSchoolYearCommand eventhorizon.CommandType = "DeleteSchoolYear"
     UpdateSchoolYearCommand eventhorizon.CommandType = "UpdateSchoolYear"
)




        

type CreateAttendance struct {
    Student  *person.Profile
    Date  *time.Time
    Course  *Course
    Hours  int
    State  *AttendanceState
    StateTrace  *schkola.Trace
    Token  string
    TokenTrace  *schkola.Trace
}

func NewCreateAttendance(student *person.Profile, date *time.Time, course *Course, hours int, state *AttendanceState, stateTrace *schkola.Trace, 
                token string, tokenTrace *schkola.Trace) (ret *CreateAttendance, err error) {
    ret = &CreateAttendance{
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
func (o *CreateAttendance) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *CreateAttendance) AggregateType() eventhorizon.AggregateType  { return AttendanceAggregateType }
func (o *CreateAttendance) CommandType() eventhorizon.CommandType      { return CreateAttendanceCommand }



        

type DeleteAttendance struct {
    Id  string
}

func NewDeleteAttendance(id string) (ret *DeleteAttendance, err error) {
    ret = &DeleteAttendance{
        Id : id,
    }
    
    return
}
func (o *DeleteAttendance) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *DeleteAttendance) AggregateType() eventhorizon.AggregateType  { return AttendanceAggregateType }
func (o *DeleteAttendance) CommandType() eventhorizon.CommandType      { return DeleteAttendanceCommand }



        

type UpdateAttendance struct {
    Student  *person.Profile
    Date  *time.Time
    Course  *Course
    Hours  int
    State  *AttendanceState
    StateTrace  *schkola.Trace
    Token  string
    TokenTrace  *schkola.Trace
}

func NewUpdateAttendance(student *person.Profile, date *time.Time, course *Course, hours int, state *AttendanceState, stateTrace *schkola.Trace, 
                token string, tokenTrace *schkola.Trace) (ret *UpdateAttendance, err error) {
    ret = &UpdateAttendance{
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
func (o *UpdateAttendance) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *UpdateAttendance) AggregateType() eventhorizon.AggregateType  { return AttendanceAggregateType }
func (o *UpdateAttendance) CommandType() eventhorizon.CommandType      { return UpdateAttendanceCommand }



        

type ConfirmAttendance struct {
}
func (o *ConfirmAttendance) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *ConfirmAttendance) AggregateType() eventhorizon.AggregateType  { return AttendanceAggregateType }
func (o *ConfirmAttendance) CommandType() eventhorizon.CommandType      { return ConfirmAttendanceCommand }



        

type CancelAttendance struct {
}
func (o *CancelAttendance) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *CancelAttendance) AggregateType() eventhorizon.AggregateType  { return AttendanceAggregateType }
func (o *CancelAttendance) CommandType() eventhorizon.CommandType      { return CancelAttendanceCommand }



        

type RegisterAttendance struct {
    Student  *person.Profile
    Course  *Course
}

func NewRegisterAttendance(student *person.Profile, course *Course) (ret *RegisterAttendance, err error) {
    ret = &RegisterAttendance{
        Student : student,
        Course : course,
    }
    
    return
}
func (o *RegisterAttendance) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *RegisterAttendance) AggregateType() eventhorizon.AggregateType  { return AttendanceAggregateType }
func (o *RegisterAttendance) CommandType() eventhorizon.CommandType      { return RegisterAttendanceCommand }



        

type CreateCourse struct {
    Name  string
    Begin  *time.Time
    End  *time.Time
    Teacher  *person.PersonName
    SchoolYear  *SchoolYear
    Fee  float64
    Description  string
}

func NewCreateCourse(name string, begin *time.Time, end *time.Time, teacher *person.PersonName, schoolYear *SchoolYear, fee float64, 
                description string) (ret *CreateCourse, err error) {
    ret = &CreateCourse{
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
func (o *CreateCourse) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *CreateCourse) AggregateType() eventhorizon.AggregateType  { return CourseAggregateType }
func (o *CreateCourse) CommandType() eventhorizon.CommandType      { return CreateCourseCommand }



        

type DeleteCourse struct {
    Id  string
}

func NewDeleteCourse(id string) (ret *DeleteCourse, err error) {
    ret = &DeleteCourse{
        Id : id,
    }
    
    return
}
func (o *DeleteCourse) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *DeleteCourse) AggregateType() eventhorizon.AggregateType  { return CourseAggregateType }
func (o *DeleteCourse) CommandType() eventhorizon.CommandType      { return DeleteCourseCommand }



        

type UpdateCourse struct {
    Name  string
    Begin  *time.Time
    End  *time.Time
    Teacher  *person.PersonName
    SchoolYear  *SchoolYear
    Fee  float64
    Description  string
}

func NewUpdateCourse(name string, begin *time.Time, end *time.Time, teacher *person.PersonName, schoolYear *SchoolYear, fee float64, 
                description string) (ret *UpdateCourse, err error) {
    ret = &UpdateCourse{
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
func (o *UpdateCourse) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *UpdateCourse) AggregateType() eventhorizon.AggregateType  { return CourseAggregateType }
func (o *UpdateCourse) CommandType() eventhorizon.CommandType      { return UpdateCourseCommand }



        

type CreateGrade struct {
    Student  *person.Profile
    Course  *Course
    Grade  float64
    GradeTrace  *schkola.Trace
    Comment  string
}

func NewCreateGrade(student *person.Profile, course *Course, grade float64, gradeTrace *schkola.Trace, comment string) (ret *CreateGrade, err error) {
    ret = &CreateGrade{
        Student : student,
        Course : course,
        Grade : grade,
        GradeTrace : gradeTrace,
        Comment : comment,
    }
    
    return
}
func (o *CreateGrade) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *CreateGrade) AggregateType() eventhorizon.AggregateType  { return GradeAggregateType }
func (o *CreateGrade) CommandType() eventhorizon.CommandType      { return CreateGradeCommand }



        

type DeleteGrade struct {
    Id  string
}

func NewDeleteGrade(id string) (ret *DeleteGrade, err error) {
    ret = &DeleteGrade{
        Id : id,
    }
    
    return
}
func (o *DeleteGrade) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *DeleteGrade) AggregateType() eventhorizon.AggregateType  { return GradeAggregateType }
func (o *DeleteGrade) CommandType() eventhorizon.CommandType      { return DeleteGradeCommand }



        

type UpdateGrade struct {
    Student  *person.Profile
    Course  *Course
    Grade  float64
    GradeTrace  *schkola.Trace
    Comment  string
}

func NewUpdateGrade(student *person.Profile, course *Course, grade float64, gradeTrace *schkola.Trace, comment string) (ret *UpdateGrade, err error) {
    ret = &UpdateGrade{
        Student : student,
        Course : course,
        Grade : grade,
        GradeTrace : gradeTrace,
        Comment : comment,
    }
    
    return
}
func (o *UpdateGrade) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *UpdateGrade) AggregateType() eventhorizon.AggregateType  { return GradeAggregateType }
func (o *UpdateGrade) CommandType() eventhorizon.CommandType      { return UpdateGradeCommand }



        

type CreateGroup struct {
    Name  string
    Category  *GroupCategory
    SchoolYear  *SchoolYear
    Representative  *person.Profile
    Students  []*Course
    Courses  []*Course
}

func NewCreateGroup(name string, category *GroupCategory, schoolYear *SchoolYear, representative *person.Profile, students []*Course, 
                courses []*Course) (ret *CreateGroup, err error) {
    ret = &CreateGroup{
        Name : name,
        Category : category,
        SchoolYear : schoolYear,
        Representative : representative,
        Students : students,
        Courses : courses,
    }
    
    return
}

func (o *CreateGroup) AddToStudents(item *Course) *Course {
    o.Students  = append(o.Students , item)
    return item
}

func (o *CreateGroup) AddToCourses(item *Course) *Course {
    o.Courses  = append(o.Courses , item)
    return item
}
func (o *CreateGroup) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *CreateGroup) AggregateType() eventhorizon.AggregateType  { return GroupAggregateType }
func (o *CreateGroup) CommandType() eventhorizon.CommandType      { return CreateGroupCommand }



        

type DeleteGroup struct {
    Id  string
}

func NewDeleteGroup(id string) (ret *DeleteGroup, err error) {
    ret = &DeleteGroup{
        Id : id,
    }
    
    return
}
func (o *DeleteGroup) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *DeleteGroup) AggregateType() eventhorizon.AggregateType  { return GroupAggregateType }
func (o *DeleteGroup) CommandType() eventhorizon.CommandType      { return DeleteGroupCommand }



        

type UpdateGroup struct {
    Name  string
    Category  *GroupCategory
    SchoolYear  *SchoolYear
    Representative  *person.Profile
    Students  []*Course
    Courses  []*Course
}

func NewUpdateGroup(name string, category *GroupCategory, schoolYear *SchoolYear, representative *person.Profile, students []*Course, 
                courses []*Course) (ret *UpdateGroup, err error) {
    ret = &UpdateGroup{
        Name : name,
        Category : category,
        SchoolYear : schoolYear,
        Representative : representative,
        Students : students,
        Courses : courses,
    }
    
    return
}

func (o *UpdateGroup) AddToStudents(item *Course) *Course {
    o.Students  = append(o.Students , item)
    return item
}

func (o *UpdateGroup) AddToCourses(item *Course) *Course {
    o.Courses  = append(o.Courses , item)
    return item
}
func (o *UpdateGroup) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *UpdateGroup) AggregateType() eventhorizon.AggregateType  { return GroupAggregateType }
func (o *UpdateGroup) CommandType() eventhorizon.CommandType      { return UpdateGroupCommand }



        

type CreateSchoolApplication struct {
    Profile  *person.Profile
    RecommendationOf  *person.PersonName
    ChurchContactPerson  *person.PersonName
    ChurchContact  *person.Contact
    SchoolYear  *SchoolYear
    Group  string
}

func NewCreateSchoolApplication(profile *person.Profile, recommendationOf *person.PersonName, churchContactPerson *person.PersonName, 
                churchContact *person.Contact, schoolYear *SchoolYear, group string) (ret *CreateSchoolApplication, err error) {
    ret = &CreateSchoolApplication{
        Profile : profile,
        RecommendationOf : recommendationOf,
        ChurchContactPerson : churchContactPerson,
        ChurchContact : churchContact,
        SchoolYear : schoolYear,
        Group : group,
    }
    
    return
}
func (o *CreateSchoolApplication) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *CreateSchoolApplication) AggregateType() eventhorizon.AggregateType  { return SchoolApplicationAggregateType }
func (o *CreateSchoolApplication) CommandType() eventhorizon.CommandType      { return CreateSchoolApplicationCommand }



        

type DeleteSchoolApplication struct {
    Id  string
}

func NewDeleteSchoolApplication(id string) (ret *DeleteSchoolApplication, err error) {
    ret = &DeleteSchoolApplication{
        Id : id,
    }
    
    return
}
func (o *DeleteSchoolApplication) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *DeleteSchoolApplication) AggregateType() eventhorizon.AggregateType  { return SchoolApplicationAggregateType }
func (o *DeleteSchoolApplication) CommandType() eventhorizon.CommandType      { return DeleteSchoolApplicationCommand }



        

type UpdateSchoolApplication struct {
    Profile  *person.Profile
    RecommendationOf  *person.PersonName
    ChurchContactPerson  *person.PersonName
    ChurchContact  *person.Contact
    SchoolYear  *SchoolYear
    Group  string
}

func NewUpdateSchoolApplication(profile *person.Profile, recommendationOf *person.PersonName, churchContactPerson *person.PersonName, 
                churchContact *person.Contact, schoolYear *SchoolYear, group string) (ret *UpdateSchoolApplication, err error) {
    ret = &UpdateSchoolApplication{
        Profile : profile,
        RecommendationOf : recommendationOf,
        ChurchContactPerson : churchContactPerson,
        ChurchContact : churchContact,
        SchoolYear : schoolYear,
        Group : group,
    }
    
    return
}
func (o *UpdateSchoolApplication) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *UpdateSchoolApplication) AggregateType() eventhorizon.AggregateType  { return SchoolApplicationAggregateType }
func (o *UpdateSchoolApplication) CommandType() eventhorizon.CommandType      { return UpdateSchoolApplicationCommand }



        

type CreateSchoolYear struct {
    Name  string
    Start  *time.Time
    End  *time.Time
    Dates  []*Course
}

func NewCreateSchoolYear(name string, start *time.Time, end *time.Time, dates []*Course) (ret *CreateSchoolYear, err error) {
    ret = &CreateSchoolYear{
        Name : name,
        Start : start,
        End : end,
        Dates : dates,
    }
    
    return
}

func (o *CreateSchoolYear) AddToDates(item *Course) *Course {
    o.Dates  = append(o.Dates , item)
    return item
}
func (o *CreateSchoolYear) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *CreateSchoolYear) AggregateType() eventhorizon.AggregateType  { return SchoolYearAggregateType }
func (o *CreateSchoolYear) CommandType() eventhorizon.CommandType      { return CreateSchoolYearCommand }



        

type DeleteSchoolYear struct {
    Id  string
}

func NewDeleteSchoolYear(id string) (ret *DeleteSchoolYear, err error) {
    ret = &DeleteSchoolYear{
        Id : id,
    }
    
    return
}
func (o *DeleteSchoolYear) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *DeleteSchoolYear) AggregateType() eventhorizon.AggregateType  { return SchoolYearAggregateType }
func (o *DeleteSchoolYear) CommandType() eventhorizon.CommandType      { return DeleteSchoolYearCommand }



        

type UpdateSchoolYear struct {
    Name  string
    Start  *time.Time
    End  *time.Time
    Dates  []*Course
}

func NewUpdateSchoolYear(name string, start *time.Time, end *time.Time, dates []*Course) (ret *UpdateSchoolYear, err error) {
    ret = &UpdateSchoolYear{
        Name : name,
        Start : start,
        End : end,
        Dates : dates,
    }
    
    return
}

func (o *UpdateSchoolYear) AddToDates(item *Course) *Course {
    o.Dates  = append(o.Dates , item)
    return item
}
func (o *UpdateSchoolYear) AggregateID() eventhorizon.UUID            { return o.Id  }
func (o *UpdateSchoolYear) AggregateType() eventhorizon.AggregateType  { return SchoolYearAggregateType }
func (o *UpdateSchoolYear) CommandType() eventhorizon.CommandType      { return UpdateSchoolYearCommand }





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

func (o *AttendanceCommandType) IsAttendanceCreate() bool {
    return o == _attendanceCommandTypes.AttendanceCreate()
}

func (o *AttendanceCommandType) IsAttendanceDelete() bool {
    return o == _attendanceCommandTypes.AttendanceDelete()
}

func (o *AttendanceCommandType) IsAttendanceUpdate() bool {
    return o == _attendanceCommandTypes.AttendanceUpdate()
}

func (o *AttendanceCommandType) IsAttendanceConfirm() bool {
    return o == _attendanceCommandTypes.AttendanceConfirm()
}

func (o *AttendanceCommandType) IsAttendanceCancel() bool {
    return o == _attendanceCommandTypes.AttendanceCancel()
}

func (o *AttendanceCommandType) IsAttendanceRegister() bool {
    return o == _attendanceCommandTypes.AttendanceRegister()
}

type attendanceCommandTypes struct {
	values []*AttendanceCommandType
    literals []enum.Literal
}

var _attendanceCommandTypes = &attendanceCommandTypes{values: []*AttendanceCommandType{
    {name: "AttendanceCreate", ordinal: 0},
    {name: "AttendanceDelete", ordinal: 1},
    {name: "AttendanceUpdate", ordinal: 2},
    {name: "AttendanceConfirm", ordinal: 3},
    {name: "AttendanceCancel", ordinal: 4},
    {name: "AttendanceRegister", ordinal: 5}},
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

func (o *attendanceCommandTypes) AttendanceCreate() *AttendanceCommandType {
    return _attendanceCommandTypes.values[0]
}

func (o *attendanceCommandTypes) AttendanceDelete() *AttendanceCommandType {
    return _attendanceCommandTypes.values[1]
}

func (o *attendanceCommandTypes) AttendanceUpdate() *AttendanceCommandType {
    return _attendanceCommandTypes.values[2]
}

func (o *attendanceCommandTypes) AttendanceConfirm() *AttendanceCommandType {
    return _attendanceCommandTypes.values[3]
}

func (o *attendanceCommandTypes) AttendanceCancel() *AttendanceCommandType {
    return _attendanceCommandTypes.values[4]
}

func (o *attendanceCommandTypes) AttendanceRegister() *AttendanceCommandType {
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

func (o *CourseCommandType) IsCourseCreate() bool {
    return o == _courseCommandTypes.CourseCreate()
}

func (o *CourseCommandType) IsCourseDelete() bool {
    return o == _courseCommandTypes.CourseDelete()
}

func (o *CourseCommandType) IsCourseUpdate() bool {
    return o == _courseCommandTypes.CourseUpdate()
}

type courseCommandTypes struct {
	values []*CourseCommandType
    literals []enum.Literal
}

var _courseCommandTypes = &courseCommandTypes{values: []*CourseCommandType{
    {name: "CourseCreate", ordinal: 0},
    {name: "CourseDelete", ordinal: 1},
    {name: "CourseUpdate", ordinal: 2}},
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

func (o *courseCommandTypes) CourseCreate() *CourseCommandType {
    return _courseCommandTypes.values[0]
}

func (o *courseCommandTypes) CourseDelete() *CourseCommandType {
    return _courseCommandTypes.values[1]
}

func (o *courseCommandTypes) CourseUpdate() *CourseCommandType {
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

func (o *GradeCommandType) IsGradeCreate() bool {
    return o == _gradeCommandTypes.GradeCreate()
}

func (o *GradeCommandType) IsGradeDelete() bool {
    return o == _gradeCommandTypes.GradeDelete()
}

func (o *GradeCommandType) IsGradeUpdate() bool {
    return o == _gradeCommandTypes.GradeUpdate()
}

type gradeCommandTypes struct {
	values []*GradeCommandType
    literals []enum.Literal
}

var _gradeCommandTypes = &gradeCommandTypes{values: []*GradeCommandType{
    {name: "GradeCreate", ordinal: 0},
    {name: "GradeDelete", ordinal: 1},
    {name: "GradeUpdate", ordinal: 2}},
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

func (o *gradeCommandTypes) GradeCreate() *GradeCommandType {
    return _gradeCommandTypes.values[0]
}

func (o *gradeCommandTypes) GradeDelete() *GradeCommandType {
    return _gradeCommandTypes.values[1]
}

func (o *gradeCommandTypes) GradeUpdate() *GradeCommandType {
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

func (o *GroupCommandType) IsGroupCreate() bool {
    return o == _groupCommandTypes.GroupCreate()
}

func (o *GroupCommandType) IsGroupDelete() bool {
    return o == _groupCommandTypes.GroupDelete()
}

func (o *GroupCommandType) IsGroupUpdate() bool {
    return o == _groupCommandTypes.GroupUpdate()
}

type groupCommandTypes struct {
	values []*GroupCommandType
    literals []enum.Literal
}

var _groupCommandTypes = &groupCommandTypes{values: []*GroupCommandType{
    {name: "GroupCreate", ordinal: 0},
    {name: "GroupDelete", ordinal: 1},
    {name: "GroupUpdate", ordinal: 2}},
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

func (o *groupCommandTypes) GroupCreate() *GroupCommandType {
    return _groupCommandTypes.values[0]
}

func (o *groupCommandTypes) GroupDelete() *GroupCommandType {
    return _groupCommandTypes.values[1]
}

func (o *groupCommandTypes) GroupUpdate() *GroupCommandType {
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

func (o *SchoolApplicationCommandType) IsSchoolApplicationCreate() bool {
    return o == _schoolApplicationCommandTypes.SchoolApplicationCreate()
}

func (o *SchoolApplicationCommandType) IsSchoolApplicationDelete() bool {
    return o == _schoolApplicationCommandTypes.SchoolApplicationDelete()
}

func (o *SchoolApplicationCommandType) IsSchoolApplicationUpdate() bool {
    return o == _schoolApplicationCommandTypes.SchoolApplicationUpdate()
}

type schoolApplicationCommandTypes struct {
	values []*SchoolApplicationCommandType
    literals []enum.Literal
}

var _schoolApplicationCommandTypes = &schoolApplicationCommandTypes{values: []*SchoolApplicationCommandType{
    {name: "SchoolApplicationCreate", ordinal: 0},
    {name: "SchoolApplicationDelete", ordinal: 1},
    {name: "SchoolApplicationUpdate", ordinal: 2}},
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

func (o *schoolApplicationCommandTypes) SchoolApplicationCreate() *SchoolApplicationCommandType {
    return _schoolApplicationCommandTypes.values[0]
}

func (o *schoolApplicationCommandTypes) SchoolApplicationDelete() *SchoolApplicationCommandType {
    return _schoolApplicationCommandTypes.values[1]
}

func (o *schoolApplicationCommandTypes) SchoolApplicationUpdate() *SchoolApplicationCommandType {
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

func (o *SchoolYearCommandType) IsSchoolYearCreate() bool {
    return o == _schoolYearCommandTypes.SchoolYearCreate()
}

func (o *SchoolYearCommandType) IsSchoolYearDelete() bool {
    return o == _schoolYearCommandTypes.SchoolYearDelete()
}

func (o *SchoolYearCommandType) IsSchoolYearUpdate() bool {
    return o == _schoolYearCommandTypes.SchoolYearUpdate()
}

type schoolYearCommandTypes struct {
	values []*SchoolYearCommandType
    literals []enum.Literal
}

var _schoolYearCommandTypes = &schoolYearCommandTypes{values: []*SchoolYearCommandType{
    {name: "SchoolYearCreate", ordinal: 0},
    {name: "SchoolYearDelete", ordinal: 1},
    {name: "SchoolYearUpdate", ordinal: 2}},
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

func (o *schoolYearCommandTypes) SchoolYearCreate() *SchoolYearCommandType {
    return _schoolYearCommandTypes.values[0]
}

func (o *schoolYearCommandTypes) SchoolYearDelete() *SchoolYearCommandType {
    return _schoolYearCommandTypes.values[1]
}

func (o *schoolYearCommandTypes) SchoolYearUpdate() *SchoolYearCommandType {
    return _schoolYearCommandTypes.values[2]
}

func (o *schoolYearCommandTypes) ParseSchoolYearCommandType(name string) (ret *SchoolYearCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.literals); ok {
		return item.(*SchoolYearCommandType), ok
	}
	return
}



