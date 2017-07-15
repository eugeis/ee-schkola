package person

import (
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

const ChurchAggregateType eventhorizon.AggregateType = "ChurchAggregate"

type ChurchAggregate struct {
    *eventhorizon.AggregateBase
    *Church
}

func NewChurchAggregate(AggregateBase *eventhorizon.AggregateBase, Entity *Church) (ret *ChurchAggregate, err error) {
    ret = &ChurchAggregate{
        AggregateBase: AggregateBase,
        Church: Entity,
    }
    
    return
}



func NewChurchAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *ChurchAggregateInitializer) {
	ret = &ChurchAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ChurchAggregateType,
        ChurchCommandTypes().Literals(), ChurchEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
    }
	return
}


func (o *ChurchAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ChurchEventTypes().CreatedChurch())
}

func (o *ChurchAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ChurchEventTypes().DeletedChurch())
}

func (o *ChurchAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ChurchEventTypes().UpdatedChurch())
}

type ChurchAggregateInitializer struct {
    *eh.AggregateInitializer
}



const GraduationAggregateType eventhorizon.AggregateType = "GraduationAggregate"

type GraduationAggregate struct {
    *eventhorizon.AggregateBase
    *Graduation
}

func NewGraduationAggregate(AggregateBase *eventhorizon.AggregateBase, Entity *Graduation) (ret *GraduationAggregate, err error) {
    ret = &GraduationAggregate{
        AggregateBase: AggregateBase,
        Graduation: Entity,
    }
    
    return
}



func NewGraduationAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *GraduationAggregateInitializer) {
	ret = &GraduationAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(GraduationAggregateType,
        GraduationCommandTypes().Literals(), GraduationEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
    }
	return
}


func (o *GraduationAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GraduationEventTypes().CreatedGraduation())
}

func (o *GraduationAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GraduationEventTypes().DeletedGraduation())
}

func (o *GraduationAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GraduationEventTypes().UpdatedGraduation())
}

type GraduationAggregateInitializer struct {
    *eh.AggregateInitializer
}



const ProfileAggregateType eventhorizon.AggregateType = "ProfileAggregate"

type ProfileAggregate struct {
    *eventhorizon.AggregateBase
    *Profile
}

func NewProfileAggregate(AggregateBase *eventhorizon.AggregateBase, Entity *Profile) (ret *ProfileAggregate, err error) {
    ret = &ProfileAggregate{
        AggregateBase: AggregateBase,
        Profile: Entity,
    }
    
    return
}



func NewProfileAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *ProfileAggregateInitializer) {
	ret = &ProfileAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ProfileAggregateType,
        ProfileCommandTypes().Literals(), ProfileEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
    }
	return
}


func (o *ProfileAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ProfileEventTypes().CreatedProfile())
}

func (o *ProfileAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ProfileEventTypes().DeletedProfile())
}

func (o *ProfileAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ProfileEventTypes().UpdatedProfile())
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









