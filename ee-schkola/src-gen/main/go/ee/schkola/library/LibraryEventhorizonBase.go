package library

import (
    "context"
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

const BookAggregateType eventhorizon.AggregateType = "BookAggregate"

func NewBookAggregate(id eventhorizon.UUID) *BookAggregate {
	return &BookAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(BookAggregateType, id),
	}
}

func (o *BookAggregate) HandleCommand(ctx context.Context, cmd eventhorizon.Command) error {
    println("HandleCommand %v - %v", ctx, cmd)
    return nil
}

func (o *BookAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent %v - %v", ctx, event)
    return nil
}

type BookAggregate struct {
    *eventhorizon.AggregateBase
    *Book
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
    bookAggregateInitializer: NewBookAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus)}
	return
}

func (o *LibraryEventhorizonInitializer) Setup() (err error) {
    
    if err = o.bookAggregateInitializer.Setup(); err != nil {
        return
    }
    return
}

type LibraryEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore
    eventBus eventhorizon.EventBus
    eventPublisher eventhorizon.EventPublisher
    commandBus eventhorizon.CommandBus
    bookAggregateInitializer *BookAggregateInitializer
}









