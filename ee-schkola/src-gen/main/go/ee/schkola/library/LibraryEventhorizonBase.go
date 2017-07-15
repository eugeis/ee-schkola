package library

import (
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

type BookAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewBookAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *BookAggregateInitializer, err error) {
    ret = &BookAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *BookAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, BookAggregateType, BookCommandTypes().Literals())
}




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











