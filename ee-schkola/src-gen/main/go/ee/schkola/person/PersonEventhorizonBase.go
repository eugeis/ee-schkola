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
	ret.CommandHandler = NewChurchAggregateCommandHandler(ret)
    return
}

func (o *ChurchAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}


func NewChurchAggregateCommandHandler(aggregate *ChurchAggregate) *ChurchAggregateCommandHandler {
	return &ChurchAggregateCommandHandler{
		aggregate: aggregate,
        handlers: make(map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *ChurchAggregate) error),
    }
}

type ChurchAggregateCommandHandler struct {
	aggregate *ChurchAggregate
	handlers  map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *ChurchAggregate) error
}

func (o *ChurchAggregateCommandHandler) AddHandler(commandType eventhorizon.CommandType,
	handler func(cmd eventhorizon.Command, aggregate *ChurchAggregate) error) {
	o.handlers[commandType] = handler
}

func (o *ChurchAggregateCommandHandler) HandleCommand(ctx context.Context, cmd eventhorizon.Command) (err error) {
	if handler, ok := o.handlers[cmd.CommandType()]; ok {
		err = handler(cmd, o.aggregate)
	} else {
		err = errors.New(fmt.Sprintf("There is no handlers for command %v registered in the aggregate %v",
			cmd.CommandType(), cmd.AggregateType()))
	}
	return
}

type ChurchAggregate struct {
    *eventhorizon.AggregateBase
    *Church
    eventhorizon.CommandHandler
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

func NewGraduationAggregate(id eventhorizon.UUID) (ret *GraduationAggregate) {
    ret = &GraduationAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(GraduationAggregateType, id),
    }
	ret.CommandHandler = NewGraduationAggregateCommandHandler(ret)
    return
}

func (o *GraduationAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}


func NewGraduationAggregateCommandHandler(aggregate *GraduationAggregate) *GraduationAggregateCommandHandler {
	return &GraduationAggregateCommandHandler{
		aggregate: aggregate,
        handlers: make(map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *GraduationAggregate) error),
    }
}

type GraduationAggregateCommandHandler struct {
	aggregate *GraduationAggregate
	handlers  map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *GraduationAggregate) error
}

func (o *GraduationAggregateCommandHandler) AddHandler(commandType eventhorizon.CommandType,
	handler func(cmd eventhorizon.Command, aggregate *GraduationAggregate) error) {
	o.handlers[commandType] = handler
}

func (o *GraduationAggregateCommandHandler) HandleCommand(ctx context.Context, cmd eventhorizon.Command) (err error) {
	if handler, ok := o.handlers[cmd.CommandType()]; ok {
		err = handler(cmd, o.aggregate)
	} else {
		err = errors.New(fmt.Sprintf("There is no handlers for command %v registered in the aggregate %v",
			cmd.CommandType(), cmd.AggregateType()))
	}
	return
}

type GraduationAggregate struct {
    *eventhorizon.AggregateBase
    *Graduation
    eventhorizon.CommandHandler
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

func NewProfileAggregate(id eventhorizon.UUID) (ret *ProfileAggregate) {
    ret = &ProfileAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(ProfileAggregateType, id),
    }
	ret.CommandHandler = NewProfileAggregateCommandHandler(ret)
    return
}

func (o *ProfileAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}


func NewProfileAggregateCommandHandler(aggregate *ProfileAggregate) *ProfileAggregateCommandHandler {
	return &ProfileAggregateCommandHandler{
		aggregate: aggregate,
        handlers: make(map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *ProfileAggregate) error),
    }
}

type ProfileAggregateCommandHandler struct {
	aggregate *ProfileAggregate
	handlers  map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *ProfileAggregate) error
}

func (o *ProfileAggregateCommandHandler) AddHandler(commandType eventhorizon.CommandType,
	handler func(cmd eventhorizon.Command, aggregate *ProfileAggregate) error) {
	o.handlers[commandType] = handler
}

func (o *ProfileAggregateCommandHandler) HandleCommand(ctx context.Context, cmd eventhorizon.Command) (err error) {
	if handler, ok := o.handlers[cmd.CommandType()]; ok {
		err = handler(cmd, o.aggregate)
	} else {
		err = errors.New(fmt.Sprintf("There is no handlers for command %v registered in the aggregate %v",
			cmd.CommandType(), cmd.AggregateType()))
	}
	return
}

type ProfileAggregate struct {
    *eventhorizon.AggregateBase
    *Profile
    eventhorizon.CommandHandler
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









