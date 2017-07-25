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
	//ret.CommandHandler = NewAttendanceAggregateCommandHandler(ret)
    return
}

func (o *AttendanceAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type AttendanceAggregate struct {
    *eventhorizon.AggregateBase
    *Attendance
    eventhorizon.CommandHandler
}



type AttendanceCommandHandler struct {
    CreateHandler  func (*CreateAttendance, *AttendanceAggregate) error
    DeleteHandler  func (*DeleteAttendance, *AttendanceAggregate) error
    UpdateHandler  func (*UpdateAttendance, *AttendanceAggregate) error
    ConfirmHandler  func (*ConfirmAttendance, *AttendanceAggregate) error
    CancelHandler  func (*CancelAttendance, *AttendanceAggregate) error
    RegisterHandler  func (*RegisterAttendance, *AttendanceAggregate) error
}

func (o *AttendanceCommandHandler) HandleCommand(ctx *context.Context, cmd eventhorizon.Command, aggregate *AttendanceAggregate) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateAttendanceCommand:
        ret = o.CreateHandler(cmd.(*CreateAttendance), aggregate)
    case DeleteAttendanceCommand:
        ret = o.DeleteHandler(cmd.(*DeleteAttendance), aggregate)
    case UpdateAttendanceCommand:
        ret = o.UpdateHandler(cmd.(*UpdateAttendance), aggregate)
    case ConfirmAttendanceCommand:
        ret = o.ConfirmHandler(cmd.(*ConfirmAttendance), aggregate)
    case CancelAttendanceCommand:
        ret = o.CancelHandler(cmd.(*CancelAttendance), aggregate)
    case RegisterAttendanceCommand:
        ret = o.RegisterHandler(cmd.(*RegisterAttendance), aggregate)
    default:
		ret = errors.New(fmt.Sprintf("Wrong comand type '%v' for aggregate '%v", cmd.CommandType(), aggregate))
	}
    return ret
    
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
    *AttendanceCommandHandler
}



const CourseAggregateType eventhorizon.AggregateType = "CourseAggregate"

func NewCourseAggregate(id eventhorizon.UUID) (ret *CourseAggregate) {
    ret = &CourseAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(CourseAggregateType, id),
    }
	//ret.CommandHandler = NewCourseAggregateCommandHandler(ret)
    return
}

func (o *CourseAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type CourseAggregate struct {
    *eventhorizon.AggregateBase
    *Course
    eventhorizon.CommandHandler
}



type CourseCommandHandler struct {
    CreateHandler  func (*CreateCourse, *CourseAggregate) error
    DeleteHandler  func (*DeleteCourse, *CourseAggregate) error
    UpdateHandler  func (*UpdateCourse, *CourseAggregate) error
}

func (o *CourseCommandHandler) HandleCommand(ctx *context.Context, cmd eventhorizon.Command, aggregate *CourseAggregate) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateCourseCommand:
        ret = o.CreateHandler(cmd.(*CreateCourse), aggregate)
    case DeleteCourseCommand:
        ret = o.DeleteHandler(cmd.(*DeleteCourse), aggregate)
    case UpdateCourseCommand:
        ret = o.UpdateHandler(cmd.(*UpdateCourse), aggregate)
    default:
		ret = errors.New(fmt.Sprintf("Wrong comand type '%v' for aggregate '%v", cmd.CommandType(), aggregate))
	}
    return ret
    
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
    *CourseCommandHandler
}



const GradeAggregateType eventhorizon.AggregateType = "GradeAggregate"

func NewGradeAggregate(id eventhorizon.UUID) (ret *GradeAggregate) {
    ret = &GradeAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(GradeAggregateType, id),
    }
	//ret.CommandHandler = NewGradeAggregateCommandHandler(ret)
    return
}

func (o *GradeAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type GradeAggregate struct {
    *eventhorizon.AggregateBase
    *Grade
    eventhorizon.CommandHandler
}



type GradeCommandHandler struct {
    CreateHandler  func (*CreateGrade, *GradeAggregate) error
    DeleteHandler  func (*DeleteGrade, *GradeAggregate) error
    UpdateHandler  func (*UpdateGrade, *GradeAggregate) error
}

func (o *GradeCommandHandler) HandleCommand(ctx *context.Context, cmd eventhorizon.Command, aggregate *GradeAggregate) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateGradeCommand:
        ret = o.CreateHandler(cmd.(*CreateGrade), aggregate)
    case DeleteGradeCommand:
        ret = o.DeleteHandler(cmd.(*DeleteGrade), aggregate)
    case UpdateGradeCommand:
        ret = o.UpdateHandler(cmd.(*UpdateGrade), aggregate)
    default:
		ret = errors.New(fmt.Sprintf("Wrong comand type '%v' for aggregate '%v", cmd.CommandType(), aggregate))
	}
    return ret
    
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
    *GradeCommandHandler
}



const GroupAggregateType eventhorizon.AggregateType = "GroupAggregate"

func NewGroupAggregate(id eventhorizon.UUID) (ret *GroupAggregate) {
    ret = &GroupAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(GroupAggregateType, id),
    }
	//ret.CommandHandler = NewGroupAggregateCommandHandler(ret)
    return
}

func (o *GroupAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type GroupAggregate struct {
    *eventhorizon.AggregateBase
    *Group
    eventhorizon.CommandHandler
}



type GroupCommandHandler struct {
    CreateHandler  func (*CreateGroup, *GroupAggregate) error
    DeleteHandler  func (*DeleteGroup, *GroupAggregate) error
    UpdateHandler  func (*UpdateGroup, *GroupAggregate) error
}

func (o *GroupCommandHandler) HandleCommand(ctx *context.Context, cmd eventhorizon.Command, aggregate *GroupAggregate) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateGroupCommand:
        ret = o.CreateHandler(cmd.(*CreateGroup), aggregate)
    case DeleteGroupCommand:
        ret = o.DeleteHandler(cmd.(*DeleteGroup), aggregate)
    case UpdateGroupCommand:
        ret = o.UpdateHandler(cmd.(*UpdateGroup), aggregate)
    default:
		ret = errors.New(fmt.Sprintf("Wrong comand type '%v' for aggregate '%v", cmd.CommandType(), aggregate))
	}
    return ret
    
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
    *GroupCommandHandler
}



const SchoolApplicationAggregateType eventhorizon.AggregateType = "SchoolApplicationAggregate"

func NewSchoolApplicationAggregate(id eventhorizon.UUID) (ret *SchoolApplicationAggregate) {
    ret = &SchoolApplicationAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(SchoolApplicationAggregateType, id),
    }
	//ret.CommandHandler = NewSchoolApplicationAggregateCommandHandler(ret)
    return
}

func (o *SchoolApplicationAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type SchoolApplicationAggregate struct {
    *eventhorizon.AggregateBase
    *SchoolApplication
    eventhorizon.CommandHandler
}



type SchoolApplicationCommandHandler struct {
    CreateHandler  func (*CreateSchoolApplication, *SchoolApplicationAggregate) error
    DeleteHandler  func (*DeleteSchoolApplication, *SchoolApplicationAggregate) error
    UpdateHandler  func (*UpdateSchoolApplication, *SchoolApplicationAggregate) error
}

func (o *SchoolApplicationCommandHandler) HandleCommand(ctx *context.Context, cmd eventhorizon.Command, aggregate *SchoolApplicationAggregate) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateSchoolApplicationCommand:
        ret = o.CreateHandler(cmd.(*CreateSchoolApplication), aggregate)
    case DeleteSchoolApplicationCommand:
        ret = o.DeleteHandler(cmd.(*DeleteSchoolApplication), aggregate)
    case UpdateSchoolApplicationCommand:
        ret = o.UpdateHandler(cmd.(*UpdateSchoolApplication), aggregate)
    default:
		ret = errors.New(fmt.Sprintf("Wrong comand type '%v' for aggregate '%v", cmd.CommandType(), aggregate))
	}
    return ret
    
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
    *SchoolApplicationCommandHandler
}



const SchoolYearAggregateType eventhorizon.AggregateType = "SchoolYearAggregate"

func NewSchoolYearAggregate(id eventhorizon.UUID) (ret *SchoolYearAggregate) {
    ret = &SchoolYearAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(SchoolYearAggregateType, id),
    }
	//ret.CommandHandler = NewSchoolYearAggregateCommandHandler(ret)
    return
}

func (o *SchoolYearAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type SchoolYearAggregate struct {
    *eventhorizon.AggregateBase
    *SchoolYear
    eventhorizon.CommandHandler
}



type SchoolYearCommandHandler struct {
    CreateHandler  func (*CreateSchoolYear, *SchoolYearAggregate) error
    DeleteHandler  func (*DeleteSchoolYear, *SchoolYearAggregate) error
    UpdateHandler  func (*UpdateSchoolYear, *SchoolYearAggregate) error
}

func (o *SchoolYearCommandHandler) HandleCommand(ctx *context.Context, cmd eventhorizon.Command, aggregate *SchoolYearAggregate) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateSchoolYearCommand:
        ret = o.CreateHandler(cmd.(*CreateSchoolYear), aggregate)
    case DeleteSchoolYearCommand:
        ret = o.DeleteHandler(cmd.(*DeleteSchoolYear), aggregate)
    case UpdateSchoolYearCommand:
        ret = o.UpdateHandler(cmd.(*UpdateSchoolYear), aggregate)
    default:
		ret = errors.New(fmt.Sprintf("Wrong comand type '%v' for aggregate '%v", cmd.CommandType(), aggregate))
	}
    return ret
    
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
    *SchoolYearCommandHandler
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









