package student

import (
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

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
    eh.RegisterCommands(handler, AttendanceAggregateType, AttendanceCommandTypes().Literals())
}




const AttendanceAttendanceType eventhorizon.AggregateType = "AttendanceAttendance"
type AttendanceAttendance struct {
    *eventhorizon.AggregateBase
    *Attendance
}

func NewAttendanceAttendance(AggregateBase *eventhorizon.AggregateBase, Entity *Attendance) (ret *AttendanceAttendance, err error) {
    ret = &AttendanceAttendance{
        AggregateBase: AggregateBase,
        Attendance: Entity,
    }
    return
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
    eh.RegisterCommands(handler, CourseAggregateType, CourseCommandTypes().Literals())
}




const CourseCourseType eventhorizon.AggregateType = "CourseCourse"
type CourseCourse struct {
    *eventhorizon.AggregateBase
    *Course
}

func NewCourseCourse(AggregateBase *eventhorizon.AggregateBase, Entity *Course) (ret *CourseCourse, err error) {
    ret = &CourseCourse{
        AggregateBase: AggregateBase,
        Course: Entity,
    }
    return
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
    eh.RegisterCommands(handler, GradeAggregateType, GradeCommandTypes().Literals())
}




const GradeGradeType eventhorizon.AggregateType = "GradeGrade"
type GradeGrade struct {
    *eventhorizon.AggregateBase
    *Grade
}

func NewGradeGrade(AggregateBase *eventhorizon.AggregateBase, Entity *Grade) (ret *GradeGrade, err error) {
    ret = &GradeGrade{
        AggregateBase: AggregateBase,
        Grade: Entity,
    }
    return
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
    eh.RegisterCommands(handler, GroupAggregateType, GroupCommandTypes().Literals())
}




const GroupGroupType eventhorizon.AggregateType = "GroupGroup"
type GroupGroup struct {
    *eventhorizon.AggregateBase
    *Group
}

func NewGroupGroup(AggregateBase *eventhorizon.AggregateBase, Entity *Group) (ret *GroupGroup, err error) {
    ret = &GroupGroup{
        AggregateBase: AggregateBase,
        Group: Entity,
    }
    return
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
    eh.RegisterCommands(handler, SchoolApplicationAggregateType, SchoolApplicationCommandTypes().Literals())
}




const SchoolApplicationSchoolApplicationType eventhorizon.AggregateType = "SchoolApplicationSchoolApplication"
type SchoolApplicationSchoolApplication struct {
    *eventhorizon.AggregateBase
    *SchoolApplication
}

func NewSchoolApplicationSchoolApplication(AggregateBase *eventhorizon.AggregateBase, Entity *SchoolApplication) (ret *SchoolApplicationSchoolApplication, err error) {
    ret = &SchoolApplicationSchoolApplication{
        AggregateBase: AggregateBase,
        SchoolApplication: Entity,
    }
    return
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
    eh.RegisterCommands(handler, SchoolYearAggregateType, SchoolYearCommandTypes().Literals())
}




const SchoolYearSchoolYearType eventhorizon.AggregateType = "SchoolYearSchoolYear"
type SchoolYearSchoolYear struct {
    *eventhorizon.AggregateBase
    *SchoolYear
}

func NewSchoolYearSchoolYear(AggregateBase *eventhorizon.AggregateBase, Entity *SchoolYear) (ret *SchoolYearSchoolYear, err error) {
    ret = &SchoolYearSchoolYear{
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











