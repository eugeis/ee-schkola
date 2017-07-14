package person

import (
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

type ChurchChurchAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewChurchChurchAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *ChurchChurchAggregateInitializer, err error) {
    ret = &ChurchChurchAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *ChurchChurchAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, ChurchAggregateType, ChurchCommandTypes().Literals())
}




const ChurchChurchType eventhorizon.AggregateType = "ChurchChurch"
type ChurchChurch struct {
    *eventhorizon.AggregateBase
    *Church
}

func NewChurchChurch(AggregateBase *eventhorizon.AggregateBase, Entity *Church) (ret *ChurchChurch, err error) {
    ret = &ChurchChurch{
        AggregateBase: AggregateBase,
        Church: Entity,
    }
    return
}



type GraduationGraduationAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewGraduationGraduationAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *GraduationGraduationAggregateInitializer, err error) {
    ret = &GraduationGraduationAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *GraduationGraduationAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, GraduationAggregateType, GraduationCommandTypes().Literals())
}




const GraduationGraduationType eventhorizon.AggregateType = "GraduationGraduation"
type GraduationGraduation struct {
    *eventhorizon.AggregateBase
    *Graduation
}

func NewGraduationGraduation(AggregateBase *eventhorizon.AggregateBase, Entity *Graduation) (ret *GraduationGraduation, err error) {
    ret = &GraduationGraduation{
        AggregateBase: AggregateBase,
        Graduation: Entity,
    }
    return
}



type ProfileProfileAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewProfileProfileAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *ProfileProfileAggregateInitializer, err error) {
    ret = &ProfileProfileAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *ProfileProfileAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, ProfileAggregateType, ProfileCommandTypes().Literals())
}




const ProfileProfileType eventhorizon.AggregateType = "ProfileProfile"
type ProfileProfile struct {
    *eventhorizon.AggregateBase
    *Profile
}

func NewProfileProfile(AggregateBase *eventhorizon.AggregateBase, Entity *Profile) (ret *ProfileProfile, err error) {
    ret = &ProfileProfile{
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











