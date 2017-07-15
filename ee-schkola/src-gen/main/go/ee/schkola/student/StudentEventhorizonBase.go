package student

import (
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

type AttendanceAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewAttendanceAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *AttendanceAggregateInitializer, err error) {
    ret = &AttendanceAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *AttendanceAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, AttendanceAggregateType, AttendanceCommandTypes().Literals())
}




const AttendanceAggregateType eventhorizon.AggregateType = "AttendanceAggregate"
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



type CourseAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewCourseAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *CourseAggregateInitializer, err error) {
    ret = &CourseAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *CourseAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, CourseAggregateType, CourseCommandTypes().Literals())
}




const CourseAggregateType eventhorizon.AggregateType = "CourseAggregate"
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



type GradeAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewGradeAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *GradeAggregateInitializer, err error) {
    ret = &GradeAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *GradeAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, GradeAggregateType, GradeCommandTypes().Literals())
}




const GradeAggregateType eventhorizon.AggregateType = "GradeAggregate"
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



type GroupAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewGroupAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *GroupAggregateInitializer, err error) {
    ret = &GroupAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *GroupAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, GroupAggregateType, GroupCommandTypes().Literals())
}




const GroupAggregateType eventhorizon.AggregateType = "GroupAggregate"
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



type SchoolApplicationAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewSchoolApplicationAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *SchoolApplicationAggregateInitializer, err error) {
    ret = &SchoolApplicationAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *SchoolApplicationAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, SchoolApplicationAggregateType, SchoolApplicationCommandTypes().Literals())
}




const SchoolApplicationAggregateType eventhorizon.AggregateType = "SchoolApplicationAggregate"
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



type SchoolYearAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewSchoolYearAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *SchoolYearAggregateInitializer, err error) {
    ret = &SchoolYearAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *SchoolYearAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, SchoolYearAggregateType, SchoolYearCommandTypes().Literals())
}




const SchoolYearAggregateType eventhorizon.AggregateType = "SchoolYearAggregate"
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











