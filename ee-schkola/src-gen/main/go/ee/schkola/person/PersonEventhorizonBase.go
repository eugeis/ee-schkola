package person

import (
    "github.com/looplab/eventhorizon"
)
type ChurchAggregate struct {
    *eventhorizon.AggregateBase
    *Church
}


type GraduationAggregate struct {
    *eventhorizon.AggregateBase
    *Graduation
}


type ProfileAggregate struct {
    *eventhorizon.AggregateBase
    *Profile
}




type ChurchChurchAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
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
      case "Church":
        ret = o.Church()
      case "Graduation":
        ret = o.Graduation()
      case "Profile":
        ret = o.Profile()
    }
    return
}



