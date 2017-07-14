package student

import (
    "github.com/looplab/eventhorizon"
)
type AttendanceAggregate struct {
    *eventhorizon.AggregateBase
    *Attendance
}

func NewAttendanceAggregate(AggregateBase *eventhorizon.AggregateBase, Entity *Attendance) (ret *AttendanceAggregate, err error) {
    ret = &AttendanceAggregate{
        AggregateBase: AggregateBase,
        Attendance: Entity,
    }
    return
}


type CourseAggregate struct {
    *eventhorizon.AggregateBase
    *Course
}

func NewCourseAggregate(AggregateBase *eventhorizon.AggregateBase, Entity *Course) (ret *CourseAggregate, err error) {
    ret = &CourseAggregate{
        AggregateBase: AggregateBase,
        Course: Entity,
    }
    return
}


type GradeAggregate struct {
    *eventhorizon.AggregateBase
    *Grade
}

func NewGradeAggregate(AggregateBase *eventhorizon.AggregateBase, Entity *Grade) (ret *GradeAggregate, err error) {
    ret = &GradeAggregate{
        AggregateBase: AggregateBase,
        Grade: Entity,
    }
    return
}


type GroupAggregate struct {
    *eventhorizon.AggregateBase
    *Group
}

func NewGroupAggregate(AggregateBase *eventhorizon.AggregateBase, Entity *Group) (ret *GroupAggregate, err error) {
    ret = &GroupAggregate{
        AggregateBase: AggregateBase,
        Group: Entity,
    }
    return
}


type SchoolApplicationAggregate struct {
    *eventhorizon.AggregateBase
    *SchoolApplication
}

func NewSchoolApplicationAggregate(AggregateBase *eventhorizon.AggregateBase, Entity *SchoolApplication) (ret *SchoolApplicationAggregate, err error) {
    ret = &SchoolApplicationAggregate{
        AggregateBase: AggregateBase,
        SchoolApplication: Entity,
    }
    return
}


type SchoolYearAggregate struct {
    *eventhorizon.AggregateBase
    *SchoolYear
}

func NewSchoolYearAggregate(AggregateBase *eventhorizon.AggregateBase, Entity *SchoolYear) (ret *SchoolYearAggregate, err error) {
    ret = &SchoolYearAggregate{
        AggregateBase: AggregateBase,
        SchoolYear: Entity,
    }
    return
}




type AttendanceAttendanceAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewAttendanceAttendanceAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *AttendanceAttendanceAggregateInitializer, err error) {
    ret = &AttendanceAttendanceAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
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

func NewCourseCourseAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *CourseCourseAggregateInitializer, err error) {
    ret = &CourseCourseAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
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

func NewGradeGradeAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *GradeGradeAggregateInitializer, err error) {
    ret = &GradeGradeAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
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

func NewGroupGroupAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *GroupGroupAggregateInitializer, err error) {
    ret = &GroupGroupAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
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

func NewSchoolApplicationSchoolApplicationAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *SchoolApplicationSchoolApplicationAggregateInitializer, err error) {
    ret = &SchoolApplicationSchoolApplicationAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
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

func NewSchoolYearSchoolYearAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *SchoolYearSchoolYearAggregateInitializer, err error) {
    ret = &SchoolYearSchoolYearAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
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

func NewStudentEventhorizonInitializer(store *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                commandBus *eventhorizon.CommandBus) (ret *StudentEventhorizonInitializer, err error) {
    ret = &StudentEventhorizonInitializer{
        Store : store,
        EventBus : eventBus,
        Publisher : publisher,
        CommandBus : commandBus,
    }
    return
}










type StudentAggregateType struct {
	name  string
	ordinal int
    commands string
    events string
}

func (o *StudentAggregateType) Name() string {
    return o.name
}

func (o *StudentAggregateType) Ordinal() int {
    return o.ordinal
}

func (o *StudentAggregateType) Commands() string {
    return o.commands
}
func (o *StudentAggregateType) Events() string {
    return o.events
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
      case o.Attendance().Name():
        ret = o.Attendance()
      case o.Course().Name():
        ret = o.Course()
      case o.Grade().Name():
        ret = o.Grade()
      case o.Group().Name():
        ret = o.Group()
      case o.SchoolApplication().Name():
        ret = o.SchoolApplication()
      case o.SchoolYear().Name():
        ret = o.SchoolYear()
    }
    return
}



