package library

import (
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

const BookAggregateType eventhorizon.AggregateType = "BookAggregate"

type BookAggregate struct {
    *eventhorizon.AggregateBase
    *Book
}

func NewBookAggregate(AggregateBase *eventhorizon.AggregateBase, Entity *Book) (ret *BookAggregate, err error) {
    ret = &BookAggregate{
        AggregateBase: AggregateBase,
        Book: Entity,
    }
    
    return
}



func NewBookAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *BookAggregateInitializer) {
	ret = &BookAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(BookAggregateType,
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









