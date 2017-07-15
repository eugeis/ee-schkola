package student

import (
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

const AttendanceAggregateType eventhorizon.AggregateType = "AttendanceAggregate"

type AttendanceAggregateInitializer struct {
    *eh.AggregateInitializer
}

func (o *AttendanceAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AttendanceAggregateEventTypes().AttendanceCreated())
}

func (o *AttendanceAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AttendanceAggregateEventTypes().AttendanceDeleted())
}

func (o *AttendanceAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AttendanceAggregateEventTypes().AttendanceUpdated())
}

func NewAttendanceAggregateInitializer(
	store *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher,
	commandBus *eventhorizon.CommandBus) (ret *AttendanceAggregateInitializer) {
	ret = &AttendanceAggregateInitializer{
        AggregateInitializer: eh.NewAggregateInitializer(AttendanceAggregateType, AttendanceAggregateCommandTypes().Literals(),
		AttendanceAggregateEventTypes().Literals(), store, eventBus, publisher, commandBus),
    }
	return
}

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



const CourseAggregateType eventhorizon.AggregateType = "CourseAggregate"

type CourseAggregateInitializer struct {
    *eh.AggregateInitializer
}

func (o *CourseAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, CourseAggregateEventTypes().CourseCreated())
}

func (o *CourseAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, CourseAggregateEventTypes().CourseDeleted())
}

func (o *CourseAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, CourseAggregateEventTypes().CourseUpdated())
}

func NewCourseAggregateInitializer(
	store *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher,
	commandBus *eventhorizon.CommandBus) (ret *CourseAggregateInitializer) {
	ret = &CourseAggregateInitializer{
        AggregateInitializer: eh.NewAggregateInitializer(CourseAggregateType, CourseAggregateCommandTypes().Literals(),
		CourseAggregateEventTypes().Literals(), store, eventBus, publisher, commandBus),
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



const GradeAggregateType eventhorizon.AggregateType = "GradeAggregate"

type GradeAggregateInitializer struct {
    *eh.AggregateInitializer
}

func (o *GradeAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GradeAggregateEventTypes().GradeCreated())
}

func (o *GradeAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GradeAggregateEventTypes().GradeDeleted())
}

func (o *GradeAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GradeAggregateEventTypes().GradeUpdated())
}

func NewGradeAggregateInitializer(
	store *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher,
	commandBus *eventhorizon.CommandBus) (ret *GradeAggregateInitializer) {
	ret = &GradeAggregateInitializer{
        AggregateInitializer: eh.NewAggregateInitializer(GradeAggregateType, GradeAggregateCommandTypes().Literals(),
		GradeAggregateEventTypes().Literals(), store, eventBus, publisher, commandBus),
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



const GroupAggregateType eventhorizon.AggregateType = "GroupAggregate"

type GroupAggregateInitializer struct {
    *eh.AggregateInitializer
}

func (o *GroupAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GroupAggregateEventTypes().GroupCreated())
}

func (o *GroupAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GroupAggregateEventTypes().GroupDeleted())
}

func (o *GroupAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GroupAggregateEventTypes().GroupUpdated())
}

func NewGroupAggregateInitializer(
	store *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher,
	commandBus *eventhorizon.CommandBus) (ret *GroupAggregateInitializer) {
	ret = &GroupAggregateInitializer{
        AggregateInitializer: eh.NewAggregateInitializer(GroupAggregateType, GroupAggregateCommandTypes().Literals(),
		GroupAggregateEventTypes().Literals(), store, eventBus, publisher, commandBus),
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



const SchoolApplicationAggregateType eventhorizon.AggregateType = "SchoolApplicationAggregate"

type SchoolApplicationAggregateInitializer struct {
    *eh.AggregateInitializer
}

func (o *SchoolApplicationAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, SchoolApplicationAggregateEventTypes().SchoolApplicationCreated())
}

func (o *SchoolApplicationAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, SchoolApplicationAggregateEventTypes().SchoolApplicationDeleted())
}

func (o *SchoolApplicationAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, SchoolApplicationAggregateEventTypes().SchoolApplicationUpdated())
}

func NewSchoolApplicationAggregateInitializer(
	store *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher,
	commandBus *eventhorizon.CommandBus) (ret *SchoolApplicationAggregateInitializer) {
	ret = &SchoolApplicationAggregateInitializer{
        AggregateInitializer: eh.NewAggregateInitializer(SchoolApplicationAggregateType, SchoolApplicationAggregateCommandTypes().Literals(),
		SchoolApplicationAggregateEventTypes().Literals(), store, eventBus, publisher, commandBus),
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



const SchoolYearAggregateType eventhorizon.AggregateType = "SchoolYearAggregate"

type SchoolYearAggregateInitializer struct {
    *eh.AggregateInitializer
}

func (o *SchoolYearAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, SchoolYearAggregateEventTypes().SchoolYearCreated())
}

func (o *SchoolYearAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, SchoolYearAggregateEventTypes().SchoolYearDeleted())
}

func (o *SchoolYearAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, SchoolYearAggregateEventTypes().SchoolYearUpdated())
}

func NewSchoolYearAggregateInitializer(
	store *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher,
	commandBus *eventhorizon.CommandBus) (ret *SchoolYearAggregateInitializer) {
	ret = &SchoolYearAggregateInitializer{
        AggregateInitializer: eh.NewAggregateInitializer(SchoolYearAggregateType, SchoolYearAggregateCommandTypes().Literals(),
		SchoolYearAggregateEventTypes().Literals(), store, eventBus, publisher, commandBus),
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











