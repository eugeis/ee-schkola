package person

import (
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

type ChurchAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewChurchAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *ChurchAggregateInitializer, err error) {
    ret = &ChurchAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *ChurchAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, ChurchAggregateType, ChurchCommandTypes().Literals())
}




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



type GraduationAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewGraduationAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *GraduationAggregateInitializer, err error) {
    ret = &GraduationAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *GraduationAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, GraduationAggregateType, GraduationCommandTypes().Literals())
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



type ProfileAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewProfileAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *ProfileAggregateInitializer, err error) {
    ret = &ProfileAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *ProfileAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, ProfileAggregateType, ProfileCommandTypes().Literals())
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











