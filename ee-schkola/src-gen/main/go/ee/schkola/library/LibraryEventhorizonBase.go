package library

import (
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

const BookAggregateType eventhorizon.AggregateType = "BookAggregate"

type BookAggregateInitializer struct {
    *eh.AggregateInitializer
}

func (o *BookAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, BookAggregateEventTypes().BookCreated())
}

func (o *BookAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, BookAggregateEventTypes().BookDeleted())
}

func (o *BookAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, BookAggregateEventTypes().BookUpdated())
}

func NewBookAggregateInitializer(
	store *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher,
	commandBus *eventhorizon.CommandBus) (ret *BookAggregateInitializer) {
	ret = &BookAggregateInitializer{
        AggregateInitializer: eh.NewAggregateInitializer(BookAggregateType, BookAggregateCommandTypes().Literals(),
		ChurchAggregateEventTypes().Literals(), store, eventBus, publisher, commandBus),
    }
	return
}

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



type LibraryEventhorizonInitializer struct {
    Store  *eventhorizon.EventStore
    EventBus  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    CommandBus  *eventhorizon.CommandBus
}

func NewLibraryEventhorizonInitializer(store *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                commandBus *eventhorizon.CommandBus) (ret *LibraryEventhorizonInitializer, err error) {
    ret = &LibraryEventhorizonInitializer{
        Store : store,
        EventBus : eventBus,
        Publisher : publisher,
        CommandBus : commandBus,
    }
    return
}











