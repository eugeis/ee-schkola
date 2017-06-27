package student

import (
    "ee/schkola"
    "ee/schkola/person"
    "time"
)
type AttendanceStatus struct {
	name  string
	ordinal int
}

func (o *AttendanceStatus) Name() string {
    return o.name
}

func (o *AttendanceStatus) Ordinal() int {
    return o.ordinal
}

func (o *AttendanceStatus) IsRegistered() bool {
    return o == _attendanceStatuss.Registered()
}

func (o *AttendanceStatus) IsConfirmed() bool {
    return o == _attendanceStatuss.Confirmed()
}

func (o *AttendanceStatus) IsCanceled() bool {
    return o == _attendanceStatuss.Canceled()
}

func (o *AttendanceStatus) IsPresent() bool {
    return o == _attendanceStatuss.Present()
}

type attendanceStatuss struct {
	values []*AttendanceStatus
}

var _attendanceStatuss = &attendanceStatuss{values: []*AttendanceStatus{
    {name: "Registered", ordinal: 0},
    {name: "Confirmed", ordinal: 1},
    {name: "Canceled", ordinal: 2},
    {name: "Present", ordinal: 3}},
}

func AttendanceStatuss() *attendanceStatuss {
	return _attendanceStatuss
}

func (o *attendanceStatuss) Values() []*AttendanceStatus {
	return o.values
}

func (o *attendanceStatuss) Registered() *AttendanceStatus {
    return _attendanceStatuss.values[0]
}

func (o *attendanceStatuss) Confirmed() *AttendanceStatus {
    return _attendanceStatuss.values[1]
}

func (o *attendanceStatuss) Canceled() *AttendanceStatus {
    return _attendanceStatuss.values[2]
}

func (o *attendanceStatuss) Present() *AttendanceStatus {
    return _attendanceStatuss.values[3]
}

func (o *attendanceStatuss) ParseAttendanceStatus(name string) (ret *AttendanceStatus, ok bool) {
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




type Attendance struct {
    Student  *person.User
    Date  *time.Time
    Course  *Course
    Hours  int
    Status  *AttendanceStatus
    StatusTrace  *schkola.Trace
    Token  string
    TokenTrace  *schkola.Trace
}

func NewAttendance(student *person.User, date *time.Time, course *Course, hours int, status *AttendanceStatus, statusTrace *schkola.Trace, 
                token string, tokenTrace *schkola.Trace) (ret Attendance, err error) {
    ret = Attendance{
        Student : student,
        Date : date,
        Course : course,
        Hours : hours,
        Status : status,
        StatusTrace : statusTrace,
        Token : token,
        TokenTrace : tokenTrace,
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
}

func NewCourse(name string, begin *time.Time, end *time.Time, teacher *person.PersonName, schoolYear *SchoolYear, fee float64, 
                description string) (ret Course, err error) {
    ret = Course{
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


type Grade struct {
    Student  *person.User
    Course  *Course
    Grade  float64
    GradeTrace  *schkola.Trace
    Comment  string
}

func NewGrade(student *person.User, course *Course, grade float64, gradeTrace *schkola.Trace, comment string) (ret Grade, err error) {
    ret = Grade{
        Student : student,
        Course : course,
        Grade : grade,
        GradeTrace : gradeTrace,
        Comment : comment,
    }
    return
}


type Group struct {
    Name  string
    Category  *GroupCategory
    SchoolYear  *SchoolYear
    Representative  *person.User
    Students  []*person.User
    Courses  []*Course
}

func NewGroup(name string, category *GroupCategory, schoolYear *SchoolYear, representative *person.User, students []*person.User, 
                courses []*Course) (ret Group, err error) {
    ret = Group{
        Name : name,
        Category : category,
        SchoolYear : schoolYear,
        Representative : representative,
        Students : students,
        Courses : courses,
    }
    return
}

func (o *Group) AddStudents(item *person.User) *person.User {
    o.Students  = append(o.Students , item)
    return item
}

func (o *Group) AddCourses(item *Course) *Course {
    o.Courses  = append(o.Courses , item)
    return item
}


type SchoolApplication struct {
    Person  *person.User
    HasRecommendation  bool
    RecommendationOf  *person.PersonName
    Mentor  *person.PersonName
    MentorContact  *person.Contact
    SchoolYear  *SchoolYear
    Group  string
}

func NewSchoolApplication(person *person.User, hasRecommendation bool, recommendationOf *person.PersonName, mentor *person.PersonName, 
                mentorContact *person.Contact, schoolYear *SchoolYear, group string) (ret SchoolApplication, err error) {
    ret = SchoolApplication{
        Person : person,
        HasRecommendation : hasRecommendation,
        RecommendationOf : recommendationOf,
        Mentor : mentor,
        MentorContact : mentorContact,
        SchoolYear : schoolYear,
        Group : group,
    }
    return
}


type SchoolYear struct {
    Name  string
    Start  *time.Time
    End  *time.Time
    Dates  []*time.Time
}

func NewSchoolYear(name string, start *time.Time, end *time.Time, dates []*time.Time) (ret SchoolYear, err error) {
    ret = SchoolYear{
        Name : name,
        Start : start,
        End : end,
        Dates : dates,
    }
    return
}

func (o *SchoolYear) AddDates(item *time.Time) *time.Time {
    o.Dates  = append(o.Dates , item)
    return item
}



