package person

import (
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

const ChurchAggregateType eventhorizon.AggregateType = "ChurchAggregate"

type ChurchAggregateInitializer struct {
    *eh.AggregateInitializer
}

func (o *ChurchAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ChurchAggregateEventTypes().ChurchCreated())
}

func (o *ChurchAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ChurchAggregateEventTypes().ChurchDeleted())
}

func (o *ChurchAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ChurchAggregateEventTypes().ChurchUpdated())
}

func NewChurchAggregateInitializer(
	store *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher,
	commandBus *eventhorizon.CommandBus) (ret *ChurchAggregateInitializer) {
	ret = &ChurchAggregateInitializer{
        AggregateInitializer: eh.NewAggregateInitializer(ChurchAggregateType, ChurchAggregateCommandTypes().Literals(),
		ChurchAggregateEventTypes().Literals(), store, eventBus, publisher, commandBus),
    }
	return
}

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



const GraduationAggregateType eventhorizon.AggregateType = "GraduationAggregate"

type GraduationAggregateInitializer struct {
    *eh.AggregateInitializer
}

func (o *GraduationAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GraduationAggregateEventTypes().GraduationCreated())
}

func (o *GraduationAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GraduationAggregateEventTypes().GraduationDeleted())
}

func (o *GraduationAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GraduationAggregateEventTypes().GraduationUpdated())
}

func NewGraduationAggregateInitializer(
	store *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher,
	commandBus *eventhorizon.CommandBus) (ret *GraduationAggregateInitializer) {
	ret = &GraduationAggregateInitializer{
        AggregateInitializer: eh.NewAggregateInitializer(GraduationAggregateType, GraduationAggregateCommandTypes().Literals(),
		ChurchAggregateEventTypes().Literals(), store, eventBus, publisher, commandBus),
    }
	return
}

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



const ProfileAggregateType eventhorizon.AggregateType = "ProfileAggregate"

type ProfileAggregateInitializer struct {
    *eh.AggregateInitializer
}

func (o *ProfileAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ProfileAggregateEventTypes().ProfileCreated())
}

func (o *ProfileAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ProfileAggregateEventTypes().ProfileDeleted())
}

func (o *ProfileAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ProfileAggregateEventTypes().ProfileUpdated())
}

func NewProfileAggregateInitializer(
	store *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher,
	commandBus *eventhorizon.CommandBus) (ret *ProfileAggregateInitializer) {
	ret = &ProfileAggregateInitializer{
        AggregateInitializer: eh.NewAggregateInitializer(ProfileAggregateType, ProfileAggregateCommandTypes().Literals(),
		ChurchAggregateEventTypes().Literals(), store, eventBus, publisher, commandBus),
    }
	return
}

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



type PersonEventhorizonInitializer struct {
    Store  *eventhorizon.EventStore
    EventBus  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    CommandBus  *eventhorizon.CommandBus
}

func NewPersonEventhorizonInitializer(store *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                commandBus *eventhorizon.CommandBus) (ret *PersonEventhorizonInitializer, err error) {
    ret = &PersonEventhorizonInitializer{
        Store : store,
        EventBus : eventBus,
        Publisher : publisher,
        CommandBus : commandBus,
    }
    return
}











