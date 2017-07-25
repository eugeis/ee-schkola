package student

import (
    "errors"
    "fmt"
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

type AttendanceCommandHandler struct {
    CreateHandler  func (*CreateAttendance, *Attendance, eh.AggregateStoreEvent) error
    DeleteHandler  func (*DeleteAttendance, *Attendance, eh.AggregateStoreEvent) error
    UpdateHandler  func (*UpdateAttendance, *Attendance, eh.AggregateStoreEvent) error
    ConfirmHandler  func (*ConfirmAttendance, *Attendance, eh.AggregateStoreEvent) error
    CancelHandler  func (*CancelAttendance, *Attendance, eh.AggregateStoreEvent) error
    RegisterHandler  func (*RegisterAttendance, *Attendance, eh.AggregateStoreEvent) error
}

func (o *AttendanceCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateAttendanceCommand:
        ret = o.CreateHandler(cmd.(*CreateAttendance), entity.(*Attendance), store)
    case DeleteAttendanceCommand:
        ret = o.DeleteHandler(cmd.(*DeleteAttendance), entity.(*Attendance), store)
    case UpdateAttendanceCommand:
        ret = o.UpdateHandler(cmd.(*UpdateAttendance), entity.(*Attendance), store)
    case ConfirmAttendanceCommand:
        ret = o.ConfirmHandler(cmd.(*ConfirmAttendance), entity.(*Attendance), store)
    case CancelAttendanceCommand:
        ret = o.CancelHandler(cmd.(*CancelAttendance), entity.(*Attendance), store)
    case RegisterAttendanceCommand:
        ret = o.RegisterHandler(cmd.(*RegisterAttendance), entity.(*Attendance), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return ret
    
}

func (o *AttendanceCommandHandler) SetupCommandHandler() error {
    
    var ret error
    return ret
    
}



type AttendanceEventHandler struct {
    CreatedHandler  func (*AttendanceCreated, *Attendance) error
    DeletedHandler  func (*AttendanceDeleted, *Attendance) error
    UpdatedHandler  func (*AttendanceUpdated, *Attendance) error
}

func (o *AttendanceEventHandler) Apply(event eventhorizon.Event, entity interface{}) error {
    
    var ret error
    switch event.EventType() {
    case AttendanceCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*AttendanceCreated), entity.(*Attendance))
    case AttendanceDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*AttendanceDeleted), entity.(*Attendance))
    case AttendanceUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*AttendanceUpdated), entity.(*Attendance))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return ret
    
}

func (o *AttendanceEventHandler) SetupEventHandler() error {
    
    var ret error
    return ret
    
}



const AttendanceAggregateType eventhorizon.AggregateType = "AttendanceAggregateInitializer"

func NewAttendanceAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *AttendanceAggregateInitializer) {
    commandHandler := &AttendanceCommandHandler{}
    eventHandler := &AttendanceEventHandler{}
	ret = &AttendanceAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(AttendanceAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(AttendanceAggregateType, id, commandHandler, eventHandler, &Attendance{})
        },
        AttendanceCommandTypes().Literals(), AttendanceEventTypes().Literals(), nil,
        eventStore, eventBus, eventPublisher, commandBus),
        AttendanceCommandHandler: commandHandler,
        AttendanceEventHandler: eventHandler,
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
    *AttendanceEventHandler
}



type CourseCommandHandler struct {
    CreateHandler  func (*CreateCourse, *Course, eh.AggregateStoreEvent) error
    DeleteHandler  func (*DeleteCourse, *Course, eh.AggregateStoreEvent) error
    UpdateHandler  func (*UpdateCourse, *Course, eh.AggregateStoreEvent) error
}

func (o *CourseCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateCourseCommand:
        ret = o.CreateHandler(cmd.(*CreateCourse), entity.(*Course), store)
    case DeleteCourseCommand:
        ret = o.DeleteHandler(cmd.(*DeleteCourse), entity.(*Course), store)
    case UpdateCourseCommand:
        ret = o.UpdateHandler(cmd.(*UpdateCourse), entity.(*Course), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return ret
    
}

func (o *CourseCommandHandler) SetupCommandHandler() error {
    
    var ret error
    return ret
    
}



type CourseEventHandler struct {
    CreatedHandler  func (*CourseCreated, *Course) error
    DeletedHandler  func (*CourseDeleted, *Course) error
    UpdatedHandler  func (*CourseUpdated, *Course) error
}

func (o *CourseEventHandler) Apply(event eventhorizon.Event, entity interface{}) error {
    
    var ret error
    switch event.EventType() {
    case CourseCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*CourseCreated), entity.(*Course))
    case CourseDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*CourseDeleted), entity.(*Course))
    case CourseUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*CourseUpdated), entity.(*Course))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return ret
    
}

func (o *CourseEventHandler) SetupEventHandler() error {
    
    var ret error
    return ret
    
}



const CourseAggregateType eventhorizon.AggregateType = "CourseAggregateInitializer"

func NewCourseAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *CourseAggregateInitializer) {
    commandHandler := &CourseCommandHandler{}
    eventHandler := &CourseEventHandler{}
	ret = &CourseAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(CourseAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(CourseAggregateType, id, commandHandler, eventHandler, &Course{})
        },
        CourseCommandTypes().Literals(), CourseEventTypes().Literals(), nil,
        eventStore, eventBus, eventPublisher, commandBus),
        CourseCommandHandler: commandHandler,
        CourseEventHandler: eventHandler,
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
    *CourseEventHandler
}



type GradeCommandHandler struct {
    CreateHandler  func (*CreateGrade, *Grade, eh.AggregateStoreEvent) error
    DeleteHandler  func (*DeleteGrade, *Grade, eh.AggregateStoreEvent) error
    UpdateHandler  func (*UpdateGrade, *Grade, eh.AggregateStoreEvent) error
}

func (o *GradeCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateGradeCommand:
        ret = o.CreateHandler(cmd.(*CreateGrade), entity.(*Grade), store)
    case DeleteGradeCommand:
        ret = o.DeleteHandler(cmd.(*DeleteGrade), entity.(*Grade), store)
    case UpdateGradeCommand:
        ret = o.UpdateHandler(cmd.(*UpdateGrade), entity.(*Grade), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return ret
    
}

func (o *GradeCommandHandler) SetupCommandHandler() error {
    
    var ret error
    return ret
    
}



type GradeEventHandler struct {
    CreatedHandler  func (*GradeCreated, *Grade) error
    DeletedHandler  func (*GradeDeleted, *Grade) error
    UpdatedHandler  func (*GradeUpdated, *Grade) error
}

func (o *GradeEventHandler) Apply(event eventhorizon.Event, entity interface{}) error {
    
    var ret error
    switch event.EventType() {
    case GradeCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*GradeCreated), entity.(*Grade))
    case GradeDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*GradeDeleted), entity.(*Grade))
    case GradeUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*GradeUpdated), entity.(*Grade))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return ret
    
}

func (o *GradeEventHandler) SetupEventHandler() error {
    
    var ret error
    return ret
    
}



const GradeAggregateType eventhorizon.AggregateType = "GradeAggregateInitializer"

func NewGradeAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *GradeAggregateInitializer) {
    commandHandler := &GradeCommandHandler{}
    eventHandler := &GradeEventHandler{}
	ret = &GradeAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(GradeAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(GradeAggregateType, id, commandHandler, eventHandler, &Grade{})
        },
        GradeCommandTypes().Literals(), GradeEventTypes().Literals(), nil,
        eventStore, eventBus, eventPublisher, commandBus),
        GradeCommandHandler: commandHandler,
        GradeEventHandler: eventHandler,
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
    *GradeEventHandler
}



type GroupCommandHandler struct {
    CreateHandler  func (*CreateGroup, *Group, eh.AggregateStoreEvent) error
    DeleteHandler  func (*DeleteGroup, *Group, eh.AggregateStoreEvent) error
    UpdateHandler  func (*UpdateGroup, *Group, eh.AggregateStoreEvent) error
}

func (o *GroupCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateGroupCommand:
        ret = o.CreateHandler(cmd.(*CreateGroup), entity.(*Group), store)
    case DeleteGroupCommand:
        ret = o.DeleteHandler(cmd.(*DeleteGroup), entity.(*Group), store)
    case UpdateGroupCommand:
        ret = o.UpdateHandler(cmd.(*UpdateGroup), entity.(*Group), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return ret
    
}

func (o *GroupCommandHandler) SetupCommandHandler() error {
    
    var ret error
    return ret
    
}



type GroupEventHandler struct {
    CreatedHandler  func (*GroupCreated, *Group) error
    DeletedHandler  func (*GroupDeleted, *Group) error
    UpdatedHandler  func (*GroupUpdated, *Group) error
}

func (o *GroupEventHandler) Apply(event eventhorizon.Event, entity interface{}) error {
    
    var ret error
    switch event.EventType() {
    case GroupCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*GroupCreated), entity.(*Group))
    case GroupDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*GroupDeleted), entity.(*Group))
    case GroupUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*GroupUpdated), entity.(*Group))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return ret
    
}

func (o *GroupEventHandler) SetupEventHandler() error {
    
    var ret error
    return ret
    
}



const GroupAggregateType eventhorizon.AggregateType = "GroupAggregateInitializer"

func NewGroupAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *GroupAggregateInitializer) {
    commandHandler := &GroupCommandHandler{}
    eventHandler := &GroupEventHandler{}
	ret = &GroupAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(GroupAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(GroupAggregateType, id, commandHandler, eventHandler, &Group{})
        },
        GroupCommandTypes().Literals(), GroupEventTypes().Literals(), nil,
        eventStore, eventBus, eventPublisher, commandBus),
        GroupCommandHandler: commandHandler,
        GroupEventHandler: eventHandler,
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
    *GroupEventHandler
}



type SchoolApplicationCommandHandler struct {
    CreateHandler  func (*CreateSchoolApplication, *SchoolApplication, eh.AggregateStoreEvent) error
    DeleteHandler  func (*DeleteSchoolApplication, *SchoolApplication, eh.AggregateStoreEvent) error
    UpdateHandler  func (*UpdateSchoolApplication, *SchoolApplication, eh.AggregateStoreEvent) error
}

func (o *SchoolApplicationCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateSchoolApplicationCommand:
        ret = o.CreateHandler(cmd.(*CreateSchoolApplication), entity.(*SchoolApplication), store)
    case DeleteSchoolApplicationCommand:
        ret = o.DeleteHandler(cmd.(*DeleteSchoolApplication), entity.(*SchoolApplication), store)
    case UpdateSchoolApplicationCommand:
        ret = o.UpdateHandler(cmd.(*UpdateSchoolApplication), entity.(*SchoolApplication), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return ret
    
}

func (o *SchoolApplicationCommandHandler) SetupCommandHandler() error {
    
    var ret error
    return ret
    
}



type SchoolApplicationEventHandler struct {
    CreatedHandler  func (*SchoolApplicationCreated, *SchoolApplication) error
    DeletedHandler  func (*SchoolApplicationDeleted, *SchoolApplication) error
    UpdatedHandler  func (*SchoolApplicationUpdated, *SchoolApplication) error
}

func (o *SchoolApplicationEventHandler) Apply(event eventhorizon.Event, entity interface{}) error {
    
    var ret error
    switch event.EventType() {
    case SchoolApplicationCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*SchoolApplicationCreated), entity.(*SchoolApplication))
    case SchoolApplicationDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*SchoolApplicationDeleted), entity.(*SchoolApplication))
    case SchoolApplicationUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*SchoolApplicationUpdated), entity.(*SchoolApplication))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return ret
    
}

func (o *SchoolApplicationEventHandler) SetupEventHandler() error {
    
    var ret error
    return ret
    
}



const SchoolApplicationAggregateType eventhorizon.AggregateType = "SchoolApplicationAggregateInitializer"

func NewSchoolApplicationAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *SchoolApplicationAggregateInitializer) {
    commandHandler := &SchoolApplicationCommandHandler{}
    eventHandler := &SchoolApplicationEventHandler{}
	ret = &SchoolApplicationAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(SchoolApplicationAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(SchoolApplicationAggregateType, id, commandHandler, eventHandler, &SchoolApplication{})
        },
        SchoolApplicationCommandTypes().Literals(), SchoolApplicationEventTypes().Literals(), nil,
        eventStore, eventBus, eventPublisher, commandBus),
        SchoolApplicationCommandHandler: commandHandler,
        SchoolApplicationEventHandler: eventHandler,
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
    *SchoolApplicationEventHandler
}



type SchoolYearCommandHandler struct {
    CreateHandler  func (*CreateSchoolYear, *SchoolYear, eh.AggregateStoreEvent) error
    DeleteHandler  func (*DeleteSchoolYear, *SchoolYear, eh.AggregateStoreEvent) error
    UpdateHandler  func (*UpdateSchoolYear, *SchoolYear, eh.AggregateStoreEvent) error
}

func (o *SchoolYearCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateSchoolYearCommand:
        ret = o.CreateHandler(cmd.(*CreateSchoolYear), entity.(*SchoolYear), store)
    case DeleteSchoolYearCommand:
        ret = o.DeleteHandler(cmd.(*DeleteSchoolYear), entity.(*SchoolYear), store)
    case UpdateSchoolYearCommand:
        ret = o.UpdateHandler(cmd.(*UpdateSchoolYear), entity.(*SchoolYear), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return ret
    
}

func (o *SchoolYearCommandHandler) SetupCommandHandler() error {
    
    var ret error
    return ret
    
}



type SchoolYearEventHandler struct {
    CreatedHandler  func (*SchoolYearCreated, *SchoolYear) error
    DeletedHandler  func (*SchoolYearDeleted, *SchoolYear) error
    UpdatedHandler  func (*SchoolYearUpdated, *SchoolYear) error
}

func (o *SchoolYearEventHandler) Apply(event eventhorizon.Event, entity interface{}) error {
    
    var ret error
    switch event.EventType() {
    case SchoolYearCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*SchoolYearCreated), entity.(*SchoolYear))
    case SchoolYearDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*SchoolYearDeleted), entity.(*SchoolYear))
    case SchoolYearUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*SchoolYearUpdated), entity.(*SchoolYear))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return ret
    
}

func (o *SchoolYearEventHandler) SetupEventHandler() error {
    
    var ret error
    return ret
    
}



const SchoolYearAggregateType eventhorizon.AggregateType = "SchoolYearAggregateInitializer"

func NewSchoolYearAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *SchoolYearAggregateInitializer) {
    commandHandler := &SchoolYearCommandHandler{}
    eventHandler := &SchoolYearEventHandler{}
	ret = &SchoolYearAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(SchoolYearAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(SchoolYearAggregateType, id, commandHandler, eventHandler, &SchoolYear{})
        },
        SchoolYearCommandTypes().Literals(), SchoolYearEventTypes().Literals(), nil,
        eventStore, eventBus, eventPublisher, commandBus),
        SchoolYearCommandHandler: commandHandler,
        SchoolYearEventHandler: eventHandler,
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
    *SchoolYearEventHandler
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









