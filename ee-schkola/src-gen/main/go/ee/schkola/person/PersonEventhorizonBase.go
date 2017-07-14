package person

import (
    "github.com/looplab/eventhorizon"
)
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
    
    aggregateType := eventhorizon.AggregateType(PersonAggregateTypes().Church)
    for _, command := range ChurchCommandTypes().Values() {
        handler.SetAggregate(aggregateType, eventhorizon.CommandType(command.Name()))
    }
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
    
    aggregateType := eventhorizon.AggregateType(PersonAggregateTypes().Graduation)
    for _, command := range GraduationCommandTypes().Values() {
        handler.SetAggregate(aggregateType, eventhorizon.CommandType(command.Name()))
    }
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
    
    aggregateType := eventhorizon.AggregateType(PersonAggregateTypes().Profile)
    for _, command := range ProfileCommandTypes().Values() {
        handler.SetAggregate(aggregateType, eventhorizon.CommandType(command.Name()))
    }
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










type PersonAggregateType struct {
	name  string
	ordinal int
}

func (o *PersonAggregateType) Name() string {
    return o.name
}

func (o *PersonAggregateType) Ordinal() int {
    return o.ordinal
}

func (o *PersonAggregateType) IsChurch() bool {
    return o == _personAggregateTypes.Church()
}

func (o *PersonAggregateType) IsGraduation() bool {
    return o == _personAggregateTypes.Graduation()
}

func (o *PersonAggregateType) IsProfile() bool {
    return o == _personAggregateTypes.Profile()
}

type personAggregateTypes struct {
	values []*PersonAggregateType
}

var _personAggregateTypes = &personAggregateTypes{values: []*PersonAggregateType{
    {name: "Church", ordinal: 0},
    {name: "Graduation", ordinal: 1},
    {name: "Profile", ordinal: 2}},
}

func PersonAggregateTypes() *personAggregateTypes {
	return _personAggregateTypes
}

func (o *personAggregateTypes) Values() []*PersonAggregateType {
	return o.values
}

func (o *personAggregateTypes) Church() *PersonAggregateType {
    return _personAggregateTypes.values[0]
}

func (o *personAggregateTypes) Graduation() *PersonAggregateType {
    return _personAggregateTypes.values[1]
}

func (o *personAggregateTypes) Profile() *PersonAggregateType {
    return _personAggregateTypes.values[2]
}

func (o *personAggregateTypes) ParsePersonAggregateType(name string) (ret *PersonAggregateType, ok bool) {
    switch name {
      case o.Church().Name():
        ret = o.Church()
      case o.Graduation().Name():
        ret = o.Graduation()
      case o.Profile().Name():
        ret = o.Profile()
    }
    return
}



