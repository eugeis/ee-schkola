package library

import (
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

type BookBookAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewBookBookAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *BookBookAggregateInitializer, err error) {
    ret = &BookBookAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *BookBookAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, BookAggregateType, BookCommandTypes().Literals())
}




const BookBookType eventhorizon.AggregateType = "BookBook"
type BookBook struct {
    *eventhorizon.AggregateBase
    *Book
}

func NewBookBook(AggregateBase *eventhorizon.AggregateBase, Entity *Book) (ret *BookBook, err error) {
    ret = &BookBook{
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











