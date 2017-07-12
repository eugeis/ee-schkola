package student

import (
    "github.com/looplab/eventhorizon"
)
type AttendanceAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func (o *AttendanceAggregateInitializer) setup() string {
    panic("Not implemented yet.")
}

func (o *AttendanceAggregateInitializer) registerCommands() string {
    panic("Not implemented yet.")
}


type CourseAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func (o *CourseAggregateInitializer) setup() string {
    panic("Not implemented yet.")
}

func (o *CourseAggregateInitializer) registerCommands() string {
    panic("Not implemented yet.")
}


type GradeAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func (o *GradeAggregateInitializer) setup() string {
    panic("Not implemented yet.")
}

func (o *GradeAggregateInitializer) registerCommands() string {
    panic("Not implemented yet.")
}


type GroupAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func (o *GroupAggregateInitializer) setup() string {
    panic("Not implemented yet.")
}

func (o *GroupAggregateInitializer) registerCommands() string {
    panic("Not implemented yet.")
}


type SchoolApplicationAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func (o *SchoolApplicationAggregateInitializer) setup() string {
    panic("Not implemented yet.")
}

func (o *SchoolApplicationAggregateInitializer) registerCommands() string {
    panic("Not implemented yet.")
}


type SchoolYearAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func (o *SchoolYearAggregateInitializer) setup() string {
    panic("Not implemented yet.")
}

func (o *SchoolYearAggregateInitializer) registerCommands() string {
    panic("Not implemented yet.")
}


type StudentEventhorizonInitializer struct {
    Store  *eventhorizon.EventStore
    EventBus  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    CommandBus  *eventhorizon.CommandBus
}

func (o *StudentEventhorizonInitializer) setup() string {
    panic("Not implemented yet.")
}

func (o *StudentEventhorizonInitializer) registerCommands() string {
    panic("Not implemented yet.")
}


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



