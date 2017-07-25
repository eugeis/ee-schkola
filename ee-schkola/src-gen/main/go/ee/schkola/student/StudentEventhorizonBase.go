package student

import (
    "context"
    "errors"
    "fmt"
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

const AttendanceAggregateType eventhorizon.AggregateType = "AttendanceAggregate"

func NewAttendanceAggregate(id eventhorizon.UUID) (ret *AttendanceAggregate) {
    ret = &AttendanceAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(AttendanceAggregateType, id),
    }
	ret.CommandHandler = NewAttendanceAggregateCommandHandler(ret)
    return
}

func (o *AttendanceAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}


func NewAttendanceAggregateCommandHandler(aggregate *AttendanceAggregate) *AttendanceAggregateCommandHandler {
	return &AttendanceAggregateCommandHandler{
		aggregate: aggregate,
        handlers: make(map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *AttendanceAggregate) error),
    }
}

type AttendanceAggregateCommandHandler struct {
	aggregate *AttendanceAggregate
	handlers  map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *AttendanceAggregate) error
}

func (o *AttendanceAggregateCommandHandler) AddHandler(commandType eventhorizon.CommandType,
	handler func(cmd eventhorizon.Command, aggregate *AttendanceAggregate) error) {
	o.handlers[commandType] = handler
}

func (o *AttendanceAggregateCommandHandler) HandleCommand(ctx context.Context, cmd eventhorizon.Command) (err error) {
	if handler, ok := o.handlers[cmd.CommandType()]; ok {
		err = handler(cmd, o.aggregate)
	} else {
		err = errors.New(fmt.Sprintf("There is no handlers for command %v registered in the aggregate %v",
			cmd.CommandType(), cmd.AggregateType()))
	}
	return
}

type AttendanceAggregate struct {
    *eventhorizon.AggregateBase
    *Attendance
    eventhorizon.CommandHandler
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

func NewCourseAggregate(id eventhorizon.UUID) (ret *CourseAggregate) {
    ret = &CourseAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(CourseAggregateType, id),
    }
	ret.CommandHandler = NewCourseAggregateCommandHandler(ret)
    return
}

func (o *CourseAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}


func NewCourseAggregateCommandHandler(aggregate *CourseAggregate) *CourseAggregateCommandHandler {
	return &CourseAggregateCommandHandler{
		aggregate: aggregate,
        handlers: make(map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *CourseAggregate) error),
    }
}

type CourseAggregateCommandHandler struct {
	aggregate *CourseAggregate
	handlers  map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *CourseAggregate) error
}

func (o *CourseAggregateCommandHandler) AddHandler(commandType eventhorizon.CommandType,
	handler func(cmd eventhorizon.Command, aggregate *CourseAggregate) error) {
	o.handlers[commandType] = handler
}

func (o *CourseAggregateCommandHandler) HandleCommand(ctx context.Context, cmd eventhorizon.Command) (err error) {
	if handler, ok := o.handlers[cmd.CommandType()]; ok {
		err = handler(cmd, o.aggregate)
	} else {
		err = errors.New(fmt.Sprintf("There is no handlers for command %v registered in the aggregate %v",
			cmd.CommandType(), cmd.AggregateType()))
	}
	return
}

type CourseAggregate struct {
    *eventhorizon.AggregateBase
    *Course
    eventhorizon.CommandHandler
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

func NewGradeAggregate(id eventhorizon.UUID) (ret *GradeAggregate) {
    ret = &GradeAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(GradeAggregateType, id),
    }
	ret.CommandHandler = NewGradeAggregateCommandHandler(ret)
    return
}

func (o *GradeAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}


func NewGradeAggregateCommandHandler(aggregate *GradeAggregate) *GradeAggregateCommandHandler {
	return &GradeAggregateCommandHandler{
		aggregate: aggregate,
        handlers: make(map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *GradeAggregate) error),
    }
}

type GradeAggregateCommandHandler struct {
	aggregate *GradeAggregate
	handlers  map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *GradeAggregate) error
}

func (o *GradeAggregateCommandHandler) AddHandler(commandType eventhorizon.CommandType,
	handler func(cmd eventhorizon.Command, aggregate *GradeAggregate) error) {
	o.handlers[commandType] = handler
}

func (o *GradeAggregateCommandHandler) HandleCommand(ctx context.Context, cmd eventhorizon.Command) (err error) {
	if handler, ok := o.handlers[cmd.CommandType()]; ok {
		err = handler(cmd, o.aggregate)
	} else {
		err = errors.New(fmt.Sprintf("There is no handlers for command %v registered in the aggregate %v",
			cmd.CommandType(), cmd.AggregateType()))
	}
	return
}

type GradeAggregate struct {
    *eventhorizon.AggregateBase
    *Grade
    eventhorizon.CommandHandler
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

func NewGroupAggregate(id eventhorizon.UUID) (ret *GroupAggregate) {
    ret = &GroupAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(GroupAggregateType, id),
    }
	ret.CommandHandler = NewGroupAggregateCommandHandler(ret)
    return
}

func (o *GroupAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}


func NewGroupAggregateCommandHandler(aggregate *GroupAggregate) *GroupAggregateCommandHandler {
	return &GroupAggregateCommandHandler{
		aggregate: aggregate,
        handlers: make(map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *GroupAggregate) error),
    }
}

type GroupAggregateCommandHandler struct {
	aggregate *GroupAggregate
	handlers  map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *GroupAggregate) error
}

func (o *GroupAggregateCommandHandler) AddHandler(commandType eventhorizon.CommandType,
	handler func(cmd eventhorizon.Command, aggregate *GroupAggregate) error) {
	o.handlers[commandType] = handler
}

func (o *GroupAggregateCommandHandler) HandleCommand(ctx context.Context, cmd eventhorizon.Command) (err error) {
	if handler, ok := o.handlers[cmd.CommandType()]; ok {
		err = handler(cmd, o.aggregate)
	} else {
		err = errors.New(fmt.Sprintf("There is no handlers for command %v registered in the aggregate %v",
			cmd.CommandType(), cmd.AggregateType()))
	}
	return
}

type GroupAggregate struct {
    *eventhorizon.AggregateBase
    *Group
    eventhorizon.CommandHandler
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

func NewSchoolApplicationAggregate(id eventhorizon.UUID) (ret *SchoolApplicationAggregate) {
    ret = &SchoolApplicationAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(SchoolApplicationAggregateType, id),
    }
	ret.CommandHandler = NewSchoolApplicationAggregateCommandHandler(ret)
    return
}

func (o *SchoolApplicationAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}


func NewSchoolApplicationAggregateCommandHandler(aggregate *SchoolApplicationAggregate) *SchoolApplicationAggregateCommandHandler {
	return &SchoolApplicationAggregateCommandHandler{
		aggregate: aggregate,
        handlers: make(map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *SchoolApplicationAggregate) error),
    }
}

type SchoolApplicationAggregateCommandHandler struct {
	aggregate *SchoolApplicationAggregate
	handlers  map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *SchoolApplicationAggregate) error
}

func (o *SchoolApplicationAggregateCommandHandler) AddHandler(commandType eventhorizon.CommandType,
	handler func(cmd eventhorizon.Command, aggregate *SchoolApplicationAggregate) error) {
	o.handlers[commandType] = handler
}

func (o *SchoolApplicationAggregateCommandHandler) HandleCommand(ctx context.Context, cmd eventhorizon.Command) (err error) {
	if handler, ok := o.handlers[cmd.CommandType()]; ok {
		err = handler(cmd, o.aggregate)
	} else {
		err = errors.New(fmt.Sprintf("There is no handlers for command %v registered in the aggregate %v",
			cmd.CommandType(), cmd.AggregateType()))
	}
	return
}

type SchoolApplicationAggregate struct {
    *eventhorizon.AggregateBase
    *SchoolApplication
    eventhorizon.CommandHandler
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

func NewSchoolYearAggregate(id eventhorizon.UUID) (ret *SchoolYearAggregate) {
    ret = &SchoolYearAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(SchoolYearAggregateType, id),
    }
	ret.CommandHandler = NewSchoolYearAggregateCommandHandler(ret)
    return
}

func (o *SchoolYearAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}


func NewSchoolYearAggregateCommandHandler(aggregate *SchoolYearAggregate) *SchoolYearAggregateCommandHandler {
	return &SchoolYearAggregateCommandHandler{
		aggregate: aggregate,
        handlers: make(map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *SchoolYearAggregate) error),
    }
}

type SchoolYearAggregateCommandHandler struct {
	aggregate *SchoolYearAggregate
	handlers  map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *SchoolYearAggregate) error
}

func (o *SchoolYearAggregateCommandHandler) AddHandler(commandType eventhorizon.CommandType,
	handler func(cmd eventhorizon.Command, aggregate *SchoolYearAggregate) error) {
	o.handlers[commandType] = handler
}

func (o *SchoolYearAggregateCommandHandler) HandleCommand(ctx context.Context, cmd eventhorizon.Command) (err error) {
	if handler, ok := o.handlers[cmd.CommandType()]; ok {
		err = handler(cmd, o.aggregate)
	} else {
		err = errors.New(fmt.Sprintf("There is no handlers for command %v registered in the aggregate %v",
			cmd.CommandType(), cmd.AggregateType()))
	}
	return
}

type SchoolYearAggregate struct {
    *eventhorizon.AggregateBase
    *SchoolYear
    eventhorizon.CommandHandler
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
    AttendanceAggregateInitializer: NewAttendanceAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    CourseAggregateInitializer: NewCourseAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    GradeAggregateInitializer: NewGradeAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    GroupAggregateInitializer: NewGroupAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    SchoolApplicationAggregateInitializer: NewSchoolApplicationAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    SchoolYearAggregateInitializer: NewSchoolYearAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus)}
	return
}

func (o *StudentEventhorizonInitializer) Setup() (err error) {
    
    if err = o.AttendanceAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.CourseAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.GradeAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.GroupAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.SchoolApplicationAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.SchoolYearAggregateInitializer.Setup(); err != nil {
        return
    }
    return
}

type StudentEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore
    eventBus eventhorizon.EventBus
    eventPublisher eventhorizon.EventPublisher
    commandBus eventhorizon.CommandBus
    AttendanceAggregateInitializer  *AttendanceAggregateInitializer
    CourseAggregateInitializer  *CourseAggregateInitializer
    GradeAggregateInitializer  *GradeAggregateInitializer
    GroupAggregateInitializer  *GroupAggregateInitializer
    SchoolApplicationAggregateInitializer  *SchoolApplicationAggregateInitializer
    SchoolYearAggregateInitializer  *SchoolYearAggregateInitializer
}









