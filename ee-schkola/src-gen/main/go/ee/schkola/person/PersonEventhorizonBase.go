package person

import (
    "errors"
    "fmt"
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

type ChurchCommandHandler struct {
    CreateHandler  func (*CreateChurch, *Church, eh.AggregateStoreEvent) error
    DeleteHandler  func (*DeleteChurch, *Church, eh.AggregateStoreEvent) error
    UpdateHandler  func (*UpdateChurch, *Church, eh.AggregateStoreEvent) error
}

func (o *ChurchCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateChurchCommand:
        ret = o.CreateHandler(cmd.(*CreateChurch), entity.(*Church), store)
    case DeleteChurchCommand:
        ret = o.DeleteHandler(cmd.(*DeleteChurch), entity.(*Church), store)
    case UpdateChurchCommand:
        ret = o.UpdateHandler(cmd.(*UpdateChurch), entity.(*Church), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return ret
    
}

func (o *ChurchCommandHandler) SetupCommandHandler() error {
    
    var ret error
    return ret
    
}



type ChurchEventHandler struct {
    CreatedHandler  func (*ChurchCreated, *Church) error
    DeletedHandler  func (*ChurchDeleted, *Church) error
    UpdatedHandler  func (*ChurchUpdated, *Church) error
}

func (o *ChurchEventHandler) Apply(event eventhorizon.Event, entity interface{}) error {
    
    var ret error
    switch event.EventType() {
    case ChurchCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*ChurchCreated), entity.(*Church))
    case ChurchDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*ChurchDeleted), entity.(*Church))
    case ChurchUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*ChurchUpdated), entity.(*Church))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return ret
    
}

func (o *ChurchEventHandler) SetupEventHandler() error {
    
    var ret error
    return ret
    
}



const ChurchAggregateType eventhorizon.AggregateType = "ChurchAggregateInitializer"

func NewChurchAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *ChurchAggregateInitializer) {
    commandHandler := &ChurchCommandHandler{}
    eventHandler := &ChurchEventHandler{}
	ret = &ChurchAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ChurchAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(ChurchAggregateType, id, commandHandler, eventHandler, &Church{})
        },
        ChurchCommandTypes().Literals(), ChurchEventTypes().Literals(),
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus),
        ChurchCommandHandler: commandHandler,
        ChurchEventHandler: eventHandler,
    }
	return
}


func (o *ChurchAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ChurchEventTypes().ChurchCreated())
}

func (o *ChurchAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ChurchEventTypes().ChurchDeleted())
}

func (o *ChurchAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ChurchEventTypes().ChurchUpdated())
}

type ChurchAggregateInitializer struct {
    *eh.AggregateInitializer
    *ChurchCommandHandler
    *ChurchEventHandler
}



type GraduationCommandHandler struct {
    CreateHandler  func (*CreateGraduation, *Graduation, eh.AggregateStoreEvent) error
    DeleteHandler  func (*DeleteGraduation, *Graduation, eh.AggregateStoreEvent) error
    UpdateHandler  func (*UpdateGraduation, *Graduation, eh.AggregateStoreEvent) error
}

func (o *GraduationCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateGraduationCommand:
        ret = o.CreateHandler(cmd.(*CreateGraduation), entity.(*Graduation), store)
    case DeleteGraduationCommand:
        ret = o.DeleteHandler(cmd.(*DeleteGraduation), entity.(*Graduation), store)
    case UpdateGraduationCommand:
        ret = o.UpdateHandler(cmd.(*UpdateGraduation), entity.(*Graduation), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return ret
    
}

func (o *GraduationCommandHandler) SetupCommandHandler() error {
    
    var ret error
    return ret
    
}



type GraduationEventHandler struct {
    CreatedHandler  func (*GraduationCreated, *Graduation) error
    DeletedHandler  func (*GraduationDeleted, *Graduation) error
    UpdatedHandler  func (*GraduationUpdated, *Graduation) error
}

func (o *GraduationEventHandler) Apply(event eventhorizon.Event, entity interface{}) error {
    
    var ret error
    switch event.EventType() {
    case GraduationCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*GraduationCreated), entity.(*Graduation))
    case GraduationDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*GraduationDeleted), entity.(*Graduation))
    case GraduationUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*GraduationUpdated), entity.(*Graduation))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return ret
    
}

func (o *GraduationEventHandler) SetupEventHandler() error {
    
    var ret error
    return ret
    
}



const GraduationAggregateType eventhorizon.AggregateType = "GraduationAggregateInitializer"

func NewGraduationAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *GraduationAggregateInitializer) {
    commandHandler := &GraduationCommandHandler{}
    eventHandler := &GraduationEventHandler{}
	ret = &GraduationAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(GraduationAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(GraduationAggregateType, id, commandHandler, eventHandler, &Graduation{})
        },
        GraduationCommandTypes().Literals(), GraduationEventTypes().Literals(),
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus),
        GraduationCommandHandler: commandHandler,
        GraduationEventHandler: eventHandler,
    }
	return
}


func (o *GraduationAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GraduationEventTypes().GraduationCreated())
}

func (o *GraduationAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GraduationEventTypes().GraduationDeleted())
}

func (o *GraduationAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GraduationEventTypes().GraduationUpdated())
}

type GraduationAggregateInitializer struct {
    *eh.AggregateInitializer
    *GraduationCommandHandler
    *GraduationEventHandler
}



type ProfileCommandHandler struct {
    CreateHandler  func (*CreateProfile, *Profile, eh.AggregateStoreEvent) error
    DeleteHandler  func (*DeleteProfile, *Profile, eh.AggregateStoreEvent) error
    UpdateHandler  func (*UpdateProfile, *Profile, eh.AggregateStoreEvent) error
}

func (o *ProfileCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateProfileCommand:
        ret = o.CreateHandler(cmd.(*CreateProfile), entity.(*Profile), store)
    case DeleteProfileCommand:
        ret = o.DeleteHandler(cmd.(*DeleteProfile), entity.(*Profile), store)
    case UpdateProfileCommand:
        ret = o.UpdateHandler(cmd.(*UpdateProfile), entity.(*Profile), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return ret
    
}

func (o *ProfileCommandHandler) SetupCommandHandler() error {
    
    var ret error
    return ret
    
}



type ProfileEventHandler struct {
    CreatedHandler  func (*ProfileCreated, *Profile) error
    DeletedHandler  func (*ProfileDeleted, *Profile) error
    UpdatedHandler  func (*ProfileUpdated, *Profile) error
}

func (o *ProfileEventHandler) Apply(event eventhorizon.Event, entity interface{}) error {
    
    var ret error
    switch event.EventType() {
    case ProfileCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*ProfileCreated), entity.(*Profile))
    case ProfileDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*ProfileDeleted), entity.(*Profile))
    case ProfileUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*ProfileUpdated), entity.(*Profile))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return ret
    
}

func (o *ProfileEventHandler) SetupEventHandler() error {
    
    var ret error
    return ret
    
}



const ProfileAggregateType eventhorizon.AggregateType = "ProfileAggregateInitializer"

func NewProfileAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *ProfileAggregateInitializer) {
    commandHandler := &ProfileCommandHandler{}
    eventHandler := &ProfileEventHandler{}
	ret = &ProfileAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ProfileAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(ProfileAggregateType, id, commandHandler, eventHandler, &Profile{})
        },
        ProfileCommandTypes().Literals(), ProfileEventTypes().Literals(),
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus),
        ProfileCommandHandler: commandHandler,
        ProfileEventHandler: eventHandler,
    }
	return
}


func (o *ProfileAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ProfileEventTypes().ProfileCreated())
}

func (o *ProfileAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ProfileEventTypes().ProfileDeleted())
}

func (o *ProfileAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ProfileEventTypes().ProfileUpdated())
}

type ProfileAggregateInitializer struct {
    *eh.AggregateInitializer
    *ProfileCommandHandler
    *ProfileEventHandler
}



func NewPersonEventhorizonInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *PersonEventhorizonInitializer) {
	ret = &PersonEventhorizonInitializer{eventStore: eventStore, eventBus: eventBus, eventPublisher: eventPublisher,
            commandBus: commandBus, 
    ChurchAggregateInitializer: NewChurchAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    GraduationAggregateInitializer: NewGraduationAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    ProfileAggregateInitializer: NewProfileAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus)}
	return
}

func (o *PersonEventhorizonInitializer) Setup() (err error) {
    
    if err = o.ChurchAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.GraduationAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.ProfileAggregateInitializer.Setup(); err != nil {
        return
    }
    return
}

type PersonEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore
    eventBus eventhorizon.EventBus
    eventPublisher eventhorizon.EventPublisher
    commandBus eventhorizon.CommandBus
    ChurchAggregateInitializer  *ChurchAggregateInitializer
    GraduationAggregateInitializer  *GraduationAggregateInitializer
    ProfileAggregateInitializer  *ProfileAggregateInitializer
}









