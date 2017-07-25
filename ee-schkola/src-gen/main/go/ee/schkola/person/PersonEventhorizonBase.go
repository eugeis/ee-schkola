package person

import (
    "context"
    "errors"
    "fmt"
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

const ChurchAggregateType eventhorizon.AggregateType = "ChurchAggregate"

func NewChurchAggregate(id eventhorizon.UUID) (ret *ChurchAggregate) {
    ret = &ChurchAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(ChurchAggregateType, id),
    }
	//ret.CommandHandler = NewChurchAggregateCommandHandler(ret)
    return
}

func (o *ChurchAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type ChurchAggregate struct {
    *eventhorizon.AggregateBase
    *Church
    eventhorizon.CommandHandler
}



type ChurchCommandHandler struct {
    CreateHandler  func (*CreateChurch, *ChurchAggregate) error
    DeleteHandler  func (*DeleteChurch, *ChurchAggregate) error
    UpdateHandler  func (*UpdateChurch, *ChurchAggregate) error
}

func (o *ChurchCommandHandler) HandleCommand(ctx *context.Context, cmd eventhorizon.Command, aggregate *ChurchAggregate) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateChurchCommand:
        ret = o.CreateHandler(cmd.(*CreateChurch), aggregate)
    case DeleteChurchCommand:
        ret = o.DeleteHandler(cmd.(*DeleteChurch), aggregate)
    case UpdateChurchCommand:
        ret = o.UpdateHandler(cmd.(*UpdateChurch), aggregate)
    default:
		ret = errors.New(fmt.Sprintf("Wrong comand type '%v' for aggregate '%v", cmd.CommandType(), aggregate))
	}
    return ret
    
}



func NewChurchAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *ChurchAggregateInitializer) {
	ret = &ChurchAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ChurchAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate { return NewChurchAggregate(id) },
        ChurchCommandTypes().Literals(), ChurchEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
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
}



const GraduationAggregateType eventhorizon.AggregateType = "GraduationAggregate"

func NewGraduationAggregate(id eventhorizon.UUID) (ret *GraduationAggregate) {
    ret = &GraduationAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(GraduationAggregateType, id),
    }
	//ret.CommandHandler = NewGraduationAggregateCommandHandler(ret)
    return
}

func (o *GraduationAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type GraduationAggregate struct {
    *eventhorizon.AggregateBase
    *Graduation
    eventhorizon.CommandHandler
}



type GraduationCommandHandler struct {
    CreateHandler  func (*CreateGraduation, *GraduationAggregate) error
    DeleteHandler  func (*DeleteGraduation, *GraduationAggregate) error
    UpdateHandler  func (*UpdateGraduation, *GraduationAggregate) error
}

func (o *GraduationCommandHandler) HandleCommand(ctx *context.Context, cmd eventhorizon.Command, aggregate *GraduationAggregate) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateGraduationCommand:
        ret = o.CreateHandler(cmd.(*CreateGraduation), aggregate)
    case DeleteGraduationCommand:
        ret = o.DeleteHandler(cmd.(*DeleteGraduation), aggregate)
    case UpdateGraduationCommand:
        ret = o.UpdateHandler(cmd.(*UpdateGraduation), aggregate)
    default:
		ret = errors.New(fmt.Sprintf("Wrong comand type '%v' for aggregate '%v", cmd.CommandType(), aggregate))
	}
    return ret
    
}



func NewGraduationAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *GraduationAggregateInitializer) {
	ret = &GraduationAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(GraduationAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate { return NewGraduationAggregate(id) },
        GraduationCommandTypes().Literals(), GraduationEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
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
}



const ProfileAggregateType eventhorizon.AggregateType = "ProfileAggregate"

func NewProfileAggregate(id eventhorizon.UUID) (ret *ProfileAggregate) {
    ret = &ProfileAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(ProfileAggregateType, id),
    }
	//ret.CommandHandler = NewProfileAggregateCommandHandler(ret)
    return
}

func (o *ProfileAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type ProfileAggregate struct {
    *eventhorizon.AggregateBase
    *Profile
    eventhorizon.CommandHandler
}



type ProfileCommandHandler struct {
    CreateHandler  func (*CreateProfile, *ProfileAggregate) error
    DeleteHandler  func (*DeleteProfile, *ProfileAggregate) error
    UpdateHandler  func (*UpdateProfile, *ProfileAggregate) error
}

func (o *ProfileCommandHandler) HandleCommand(ctx *context.Context, cmd eventhorizon.Command, aggregate *ProfileAggregate) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateProfileCommand:
        ret = o.CreateHandler(cmd.(*CreateProfile), aggregate)
    case DeleteProfileCommand:
        ret = o.DeleteHandler(cmd.(*DeleteProfile), aggregate)
    case UpdateProfileCommand:
        ret = o.UpdateHandler(cmd.(*UpdateProfile), aggregate)
    default:
		ret = errors.New(fmt.Sprintf("Wrong comand type '%v' for aggregate '%v", cmd.CommandType(), aggregate))
	}
    return ret
    
}



func NewProfileAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *ProfileAggregateInitializer) {
	ret = &ProfileAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ProfileAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate { return NewProfileAggregate(id) },
        ProfileCommandTypes().Literals(), ProfileEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
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









