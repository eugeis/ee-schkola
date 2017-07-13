package student

import (
    "github.com/looplab/eventhorizon"
)
type AttendanceAggregate struct {
    *eventhorizon.AggregateBase
    *Attendance
}


type CourseAggregate struct {
    *eventhorizon.AggregateBase
    *Course
}


type GradeAggregate struct {
    *eventhorizon.AggregateBase
    *Grade
}


type GroupAggregate struct {
    *eventhorizon.AggregateBase
    *Group
}


type SchoolApplicationAggregate struct {
    *eventhorizon.AggregateBase
    *SchoolApplication
}


type SchoolYearAggregate struct {
    *eventhorizon.AggregateBase
    *SchoolYear
}




type AttendanceAttendanceAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}


func (o *AttendanceAttendanceAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    
    aggregateType := eventhorizon.AggregateType(StudentAggregateTypes().Attendance)
    for _, command := range AttendanceCommandTypes().Values() {
        handler.SetAggregate(aggregateType, eventhorizon.CommandType(command.Name()))
    }
}


type CourseCourseAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}


func (o *CourseCourseAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    
    aggregateType := eventhorizon.AggregateType(StudentAggregateTypes().Course)
    for _, command := range CourseCommandTypes().Values() {
        handler.SetAggregate(aggregateType, eventhorizon.CommandType(command.Name()))
    }
}


type GradeGradeAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}


func (o *GradeGradeAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    
    aggregateType := eventhorizon.AggregateType(StudentAggregateTypes().Grade)
    for _, command := range GradeCommandTypes().Values() {
        handler.SetAggregate(aggregateType, eventhorizon.CommandType(command.Name()))
    }
}


type GroupGroupAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}


func (o *GroupGroupAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    
    aggregateType := eventhorizon.AggregateType(StudentAggregateTypes().Group)
    for _, command := range GroupCommandTypes().Values() {
        handler.SetAggregate(aggregateType, eventhorizon.CommandType(command.Name()))
    }
}


type SchoolApplicationSchoolApplicationAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}


func (o *SchoolApplicationSchoolApplicationAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    
    aggregateType := eventhorizon.AggregateType(StudentAggregateTypes().SchoolApplication)
    for _, command := range SchoolApplicationCommandTypes().Values() {
        handler.SetAggregate(aggregateType, eventhorizon.CommandType(command.Name()))
    }
}


type SchoolYearSchoolYearAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}


func (o *SchoolYearSchoolYearAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    
    aggregateType := eventhorizon.AggregateType(StudentAggregateTypes().SchoolYear)
    for _, command := range SchoolYearCommandTypes().Values() {
        handler.SetAggregate(aggregateType, eventhorizon.CommandType(command.Name()))
    }
}


type StudentEventhorizonInitializer struct {
    Store  *eventhorizon.EventStore
    EventBus  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    CommandBus  *eventhorizon.CommandBus
}










type StudentAggregateType struct {
	name  string
	ordinal int
}

func (o *StudentAggregateType) Name() string {
    return o.name
}

func (o *StudentAggregateType) Ordinal() int {
    return o.ordinal
}

func (o *StudentAggregateType) IsAttendance() bool {
    return o == _studentAggregateTypes.Attendance()
}

func (o *StudentAggregateType) IsCourse() bool {
    return o == _studentAggregateTypes.Course()
}

func (o *StudentAggregateType) IsGrade() bool {
    return o == _studentAggregateTypes.Grade()
}

func (o *StudentAggregateType) IsGroup() bool {
    return o == _studentAggregateTypes.Group()
}

func (o *StudentAggregateType) IsSchoolApplication() bool {
    return o == _studentAggregateTypes.SchoolApplication()
}

func (o *StudentAggregateType) IsSchoolYear() bool {
    return o == _studentAggregateTypes.SchoolYear()
}

type studentAggregateTypes struct {
	values []*StudentAggregateType
}

var _studentAggregateTypes = &studentAggregateTypes{values: []*StudentAggregateType{
    {name: "Attendance", ordinal: 0},
    {name: "Course", ordinal: 1},
    {name: "Grade", ordinal: 2},
    {name: "Group", ordinal: 3},
    {name: "SchoolApplication", ordinal: 4},
    {name: "SchoolYear", ordinal: 5}},
}

func StudentAggregateTypes() *studentAggregateTypes {
	return _studentAggregateTypes
}

func (o *studentAggregateTypes) Values() []*StudentAggregateType {
	return o.values
}

func (o *studentAggregateTypes) Attendance() *StudentAggregateType {
    return _studentAggregateTypes.values[0]
}

func (o *studentAggregateTypes) Course() *StudentAggregateType {
    return _studentAggregateTypes.values[1]
}

func (o *studentAggregateTypes) Grade() *StudentAggregateType {
    return _studentAggregateTypes.values[2]
}

func (o *studentAggregateTypes) Group() *StudentAggregateType {
    return _studentAggregateTypes.values[3]
}

func (o *studentAggregateTypes) SchoolApplication() *StudentAggregateType {
    return _studentAggregateTypes.values[4]
}

func (o *studentAggregateTypes) SchoolYear() *StudentAggregateType {
    return _studentAggregateTypes.values[5]
}

func (o *studentAggregateTypes) ParseStudentAggregateType(name string) (ret *StudentAggregateType, ok bool) {
    switch name {
      case "Attendance":
        ret = o.Attendance()
      case "Course":
        ret = o.Course()
      case "Grade":
        ret = o.Grade()
      case "Group":
        ret = o.Group()
      case "SchoolApplication":
        ret = o.SchoolApplication()
      case "SchoolYear":
        ret = o.SchoolYear()
    }
    return
}



