package student

import (
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

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



func NewAttendanceAggregateInitializer(
	eventStore *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, eventPublisher *eventhorizon.EventPublisher,
	commandBus *eventhorizon.CommandBus) (ret *AttendanceAggregateInitializer) {
	ret = &AttendanceAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(AttendanceAggregateType,
        AttendanceCommandTypes().Literals(), AttendanceEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
    }
	return
}


func (o *AttendanceAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AttendanceEventTypes().AttendanceCreated())
}

func (o *AttendanceAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AttendanceEventTypes().AttendanceDeleted())
}

func (o *AttendanceAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AttendanceEventTypes().AttendanceUpdated())
}

type AttendanceAggregateInitializer struct {
    *eh.AggregateInitializer
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



func NewCourseAggregateInitializer(
	eventStore *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, eventPublisher *eventhorizon.EventPublisher,
	commandBus *eventhorizon.CommandBus) (ret *CourseAggregateInitializer) {
	ret = &CourseAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(CourseAggregateType,
        CourseCommandTypes().Literals(), CourseEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
    }
	return
}


func (o *CourseAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, CourseEventTypes().CourseCreated())
}

func (o *CourseAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, CourseEventTypes().CourseDeleted())
}

func (o *CourseAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, CourseEventTypes().CourseUpdated())
}

type CourseAggregateInitializer struct {
    *eh.AggregateInitializer
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



func NewGradeAggregateInitializer(
	eventStore *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, eventPublisher *eventhorizon.EventPublisher,
	commandBus *eventhorizon.CommandBus) (ret *GradeAggregateInitializer) {
	ret = &GradeAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(GradeAggregateType,
        GradeCommandTypes().Literals(), GradeEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
    }
	return
}


func (o *GradeAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GradeEventTypes().GradeCreated())
}

func (o *GradeAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GradeEventTypes().GradeDeleted())
}

func (o *GradeAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GradeEventTypes().GradeUpdated())
}

type GradeAggregateInitializer struct {
    *eh.AggregateInitializer
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



func NewGroupAggregateInitializer(
	eventStore *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, eventPublisher *eventhorizon.EventPublisher,
	commandBus *eventhorizon.CommandBus) (ret *GroupAggregateInitializer) {
	ret = &GroupAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(GroupAggregateType,
        GroupCommandTypes().Literals(), GroupEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
    }
	return
}


func (o *GroupAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GroupEventTypes().GroupCreated())
}

func (o *GroupAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GroupEventTypes().GroupDeleted())
}

func (o *GroupAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GroupEventTypes().GroupUpdated())
}

type GroupAggregateInitializer struct {
    *eh.AggregateInitializer
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



func NewSchoolApplicationAggregateInitializer(
	eventStore *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, eventPublisher *eventhorizon.EventPublisher,
	commandBus *eventhorizon.CommandBus) (ret *SchoolApplicationAggregateInitializer) {
	ret = &SchoolApplicationAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(SchoolApplicationAggregateType,
        SchoolApplicationCommandTypes().Literals(), SchoolApplicationEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
    }
	return
}


func (o *SchoolApplicationAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, SchoolApplicationEventTypes().SchoolApplicationCreated())
}

func (o *SchoolApplicationAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, SchoolApplicationEventTypes().SchoolApplicationDeleted())
}

func (o *SchoolApplicationAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, SchoolApplicationEventTypes().SchoolApplicationUpdated())
}

type SchoolApplicationAggregateInitializer struct {
    *eh.AggregateInitializer
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



func NewSchoolYearAggregateInitializer(
	eventStore *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, eventPublisher *eventhorizon.EventPublisher,
	commandBus *eventhorizon.CommandBus) (ret *SchoolYearAggregateInitializer) {
	ret = &SchoolYearAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(SchoolYearAggregateType,
        SchoolYearCommandTypes().Literals(), SchoolYearEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
    }
	return
}


func (o *SchoolYearAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, SchoolYearEventTypes().SchoolYearCreated())
}

func (o *SchoolYearAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, SchoolYearEventTypes().SchoolYearDeleted())
}

func (o *SchoolYearAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, SchoolYearEventTypes().SchoolYearUpdated())
}

type SchoolYearAggregateInitializer struct {
    *eh.AggregateInitializer
}



func NewStudentEventhorizonInitializer(
	eventStore *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, eventPublisher *eventhorizon.EventPublisher,
	commandBus *eventhorizon.CommandBus) (ret *StudentEventhorizonInitializer) {
	ret = &StudentEventhorizonInitializer{eventStore: eventStore, eventBus: eventBus, eventPublisher: eventPublisher,
            commandBus: commandBus, 
    attendanceAggregateInitializer: NewAttendanceAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    courseAggregateInitializer: NewCourseAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    gradeAggregateInitializer: NewGradeAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    groupAggregateInitializer: NewGroupAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    schoolApplicationAggregateInitializer: NewSchoolApplicationAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    schoolYearAggregateInitializer: NewSchoolYearAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus)}
	return
}

func (o *StudentEventhorizonInitializer) Setup() (err error) {
    
    if err = o.attendanceAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.courseAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.gradeAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.groupAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.schoolApplicationAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.schoolYearAggregateInitializer.Setup(); err != nil {
        return
    }
    return
}

type StudentEventhorizonInitializer struct {
    eventStore *eventhorizon.EventStore
    eventBus *eventhorizon.EventBus
    eventPublisher *eventhorizon.EventPublisher
    commandBus *eventhorizon.CommandBus
    attendanceAggregateInitializer *AttendanceAggregateInitializer
    courseAggregateInitializer *CourseAggregateInitializer
    gradeAggregateInitializer *GradeAggregateInitializer
    groupAggregateInitializer *GroupAggregateInitializer
    schoolApplicationAggregateInitializer *SchoolApplicationAggregateInitializer
    schoolYearAggregateInitializer *SchoolYearAggregateInitializer
}









