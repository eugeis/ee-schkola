package student

import (
    "context"
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

const AttendanceAggregateType eventhorizon.AggregateType = "AttendanceAggregate"

func NewAttendanceAggregate(id eventhorizon.UUID) *AttendanceAggregate {
	return &AttendanceAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(AttendanceAggregateType, id),
	}
}

func (o *AttendanceAggregate) HandleCommand(ctx context.Context, cmd eventhorizon.Command) error {
    println("HandleCommand", cmd.CommandType())
    return nil
}

func (o *AttendanceAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type AttendanceAggregate struct {
    *eventhorizon.AggregateBase
    *Attendance
}



func NewAttendanceAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *AttendanceAggregateInitializer) {
	ret = &AttendanceAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(AttendanceAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate { return NewAttendanceAggregate(id) },
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

func NewCourseAggregate(id eventhorizon.UUID) *CourseAggregate {
	return &CourseAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(CourseAggregateType, id),
	}
}

func (o *CourseAggregate) HandleCommand(ctx context.Context, cmd eventhorizon.Command) error {
    println("HandleCommand", cmd.CommandType())
    return nil
}

func (o *CourseAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type CourseAggregate struct {
    *eventhorizon.AggregateBase
    *Course
}



func NewCourseAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *CourseAggregateInitializer) {
	ret = &CourseAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(CourseAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate { return NewCourseAggregate(id) },
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

func NewGradeAggregate(id eventhorizon.UUID) *GradeAggregate {
	return &GradeAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(GradeAggregateType, id),
	}
}

func (o *GradeAggregate) HandleCommand(ctx context.Context, cmd eventhorizon.Command) error {
    println("HandleCommand", cmd.CommandType())
    return nil
}

func (o *GradeAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type GradeAggregate struct {
    *eventhorizon.AggregateBase
    *Grade
}



func NewGradeAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *GradeAggregateInitializer) {
	ret = &GradeAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(GradeAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate { return NewGradeAggregate(id) },
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

func NewGroupAggregate(id eventhorizon.UUID) *GroupAggregate {
	return &GroupAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(GroupAggregateType, id),
	}
}

func (o *GroupAggregate) HandleCommand(ctx context.Context, cmd eventhorizon.Command) error {
    println("HandleCommand", cmd.CommandType())
    return nil
}

func (o *GroupAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type GroupAggregate struct {
    *eventhorizon.AggregateBase
    *Group
}



func NewGroupAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *GroupAggregateInitializer) {
	ret = &GroupAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(GroupAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate { return NewGroupAggregate(id) },
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

func NewSchoolApplicationAggregate(id eventhorizon.UUID) *SchoolApplicationAggregate {
	return &SchoolApplicationAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(SchoolApplicationAggregateType, id),
	}
}

func (o *SchoolApplicationAggregate) HandleCommand(ctx context.Context, cmd eventhorizon.Command) error {
    println("HandleCommand", cmd.CommandType())
    return nil
}

func (o *SchoolApplicationAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type SchoolApplicationAggregate struct {
    *eventhorizon.AggregateBase
    *SchoolApplication
}



func NewSchoolApplicationAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *SchoolApplicationAggregateInitializer) {
	ret = &SchoolApplicationAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(SchoolApplicationAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate { return NewSchoolApplicationAggregate(id) },
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

func NewSchoolYearAggregate(id eventhorizon.UUID) *SchoolYearAggregate {
	return &SchoolYearAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(SchoolYearAggregateType, id),
	}
}

func (o *SchoolYearAggregate) HandleCommand(ctx context.Context, cmd eventhorizon.Command) error {
    println("HandleCommand", cmd.CommandType())
    return nil
}

func (o *SchoolYearAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type SchoolYearAggregate struct {
    *eventhorizon.AggregateBase
    *SchoolYear
}



func NewSchoolYearAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *SchoolYearAggregateInitializer) {
	ret = &SchoolYearAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(SchoolYearAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate { return NewSchoolYearAggregate(id) },
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
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *StudentEventhorizonInitializer) {
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
    eventStore eventhorizon.EventStore
    eventBus eventhorizon.EventBus
    eventPublisher eventhorizon.EventPublisher
    commandBus eventhorizon.CommandBus
    attendanceAggregateInitializer *AttendanceAggregateInitializer
    courseAggregateInitializer *CourseAggregateInitializer
    gradeAggregateInitializer *GradeAggregateInitializer
    groupAggregateInitializer *GroupAggregateInitializer
    schoolApplicationAggregateInitializer *SchoolApplicationAggregateInitializer
    schoolYearAggregateInitializer *SchoolYearAggregateInitializer
}









