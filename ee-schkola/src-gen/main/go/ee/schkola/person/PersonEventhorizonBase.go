package person

import (
    "context"
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

const ChurchAggregateType eventhorizon.AggregateType = "ChurchAggregate"

func NewChurchAggregate(id eventhorizon.UUID) *ChurchAggregate {
	return &ChurchAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(ChurchAggregateType, id),
	}
}

func (o *ChurchAggregate) HandleCommand(ctx context.Context, cmd eventhorizon.Command) error {
    println("HandleCommand", cmd.CommandType())
    return nil
}

func (o *ChurchAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type ChurchAggregate struct {
    *eventhorizon.AggregateBase
    *Church
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
}



const GraduationAggregateType eventhorizon.AggregateType = "GraduationAggregate"

func NewGraduationAggregate(id eventhorizon.UUID) *GraduationAggregate {
	return &GraduationAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(GraduationAggregateType, id),
	}
}

func (o *GraduationAggregate) HandleCommand(ctx context.Context, cmd eventhorizon.Command) error {
    println("HandleCommand", cmd.CommandType())
    return nil
}

func (o *GraduationAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type GraduationAggregate struct {
    *eventhorizon.AggregateBase
    *Graduation
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
}



const ProfileAggregateType eventhorizon.AggregateType = "ProfileAggregate"

func NewProfileAggregate(id eventhorizon.UUID) *ProfileAggregate {
	return &ProfileAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(ProfileAggregateType, id),
	}
}

func (o *ProfileAggregate) HandleCommand(ctx context.Context, cmd eventhorizon.Command) error {
    println("HandleCommand", cmd.CommandType())
    return nil
}

func (o *ProfileAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type ProfileAggregate struct {
    *eventhorizon.AggregateBase
    *Profile
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
}



func NewPersonEventhorizonInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *PersonEventhorizonInitializer) {
	ret = &PersonEventhorizonInitializer{eventStore: eventStore, eventBus: eventBus, eventPublisher: eventPublisher,
            commandBus: commandBus, 
    churchAggregateInitializer: NewChurchAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    graduationAggregateInitializer: NewGraduationAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    profileAggregateInitializer: NewProfileAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus)}
	return
}

func (o *PersonEventhorizonInitializer) Setup() (err error) {
    
    if err = o.churchAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.graduationAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.profileAggregateInitializer.Setup(); err != nil {
        return
    }
    return
}

type PersonEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore
    eventBus eventhorizon.EventBus
    eventPublisher eventhorizon.EventPublisher
    commandBus eventhorizon.CommandBus
    churchAggregateInitializer *ChurchAggregateInitializer
    graduationAggregateInitializer *GraduationAggregateInitializer
    profileAggregateInitializer *ProfileAggregateInitializer
}









