package library

import (
    "context"
    "errors"
    "fmt"
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

const BookAggregateType eventhorizon.AggregateType = "BookAggregate"

func NewBookAggregate(id eventhorizon.UUID) (ret *BookAggregate) {
    ret = &BookAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(BookAggregateType, id),
    }
	ret.CommandHandler = NewBookAggregateCommandHandler(ret)
    return
}

func (o *BookAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}


func NewBookAggregateCommandHandler(aggregate *BookAggregate) *BookAggregateCommandHandler {
	return &BookAggregateCommandHandler{
		aggregate: aggregate,
        handlers: make(map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *BookAggregate) error),
    }
}

type BookAggregateCommandHandler struct {
	aggregate *BookAggregate
	handlers  map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *BookAggregate) error
}

func (o *BookAggregateCommandHandler) AddHandler(commandType eventhorizon.CommandType,
	handler func(cmd eventhorizon.Command, aggregate *BookAggregate) error) {
	o.handlers[commandType] = handler
}

func (o *BookAggregateCommandHandler) HandleCommand(ctx context.Context, cmd eventhorizon.Command) (err error) {
	if handler, ok := o.handlers[cmd.CommandType()]; ok {
		err = handler(cmd, o.aggregate)
	} else {
		err = errors.New(fmt.Sprintf("There is no handlers for command %v registered in the aggregate %v",
			cmd.CommandType(), cmd.AggregateType()))
	}
	return
}

type BookAggregate struct {
    *eventhorizon.AggregateBase
    *Book
    eventhorizon.CommandHandler
}



func NewBookAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *BookAggregateInitializer) {
	ret = &BookAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(BookAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate { return NewBookAggregate(id) },
        BookCommandTypes().Literals(), BookEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
    }
	return
}


func (o *BookAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, BookEventTypes().BookCreated())
}

func (o *BookAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, BookEventTypes().BookDeleted())
}

func (o *BookAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, BookEventTypes().BookUpdated())
}

type BookAggregateInitializer struct {
    *eh.AggregateInitializer
}



func NewLibraryEventhorizonInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *LibraryEventhorizonInitializer) {
	ret = &LibraryEventhorizonInitializer{eventStore: eventStore, eventBus: eventBus, eventPublisher: eventPublisher,
            commandBus: commandBus, 
    BookAggregateInitializer: NewBookAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus)}
	return
}

func (o *LibraryEventhorizonInitializer) Setup() (err error) {
    
    if err = o.BookAggregateInitializer.Setup(); err != nil {
        return
    }
    return
}

type LibraryEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore
    eventBus eventhorizon.EventBus
    eventPublisher eventhorizon.EventPublisher
    commandBus eventhorizon.CommandBus
    BookAggregateInitializer  *BookAggregateInitializer
}









