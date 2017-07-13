package student

import (
    "ee/schkola"
    "ee/schkola/person"
    "time"
)
type CourseCreated struct {
    Name  
    Begin  *time.Time
    End  *time.Time
    Teacher  *person.PersonName
    SchoolYear  *SchoolYear
    Fee  float64
    Description  string
    *schkola.SchkolaBase
}


type CourseDeleted struct {
    Id  
}


type CourseUpdated struct {
    Name  
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
    Comment  
    *schkola.SchkolaBase
}


type GradeDeleted struct {
    Id  
}


type GradeUpdated struct {
    Student  *person.Profile
    Course  *Course
    Grade  float64
    GradeTrace  *schkola.Trace
    Comment  
    *schkola.SchkolaBase
}


type GroupCreated struct {
    Name  
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
    Id  
}


type GroupUpdated struct {
    Name  
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
    Group  
    *schkola.SchkolaBase
}


type SchoolApplicationDeleted struct {
    Id  
}


type SchoolApplicationUpdated struct {
    Profile  *person.Profile
    RecommendationOf  *person.PersonName
    ChurchContactPerson  *person.PersonName
    ChurchContact  *person.Contact
    SchoolYear  *SchoolYear
    Group  
    *schkola.SchkolaBase
}


type SchoolYearCreated struct {
    Name  
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
    Id  
}


type SchoolYearUpdated struct {
    Name  
    Start  *time.Time
    End  *time.Time
    Dates  []*Course
    *schkola.SchkolaBase
}

func (o *SchoolYearUpdated) AddToDates(item *Course) *Course {
    o.Dates  = append(o.Dates , item)
    return item
}





