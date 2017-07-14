package library

import (
    "github.com/looplab/eventhorizon"
)
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
    
    aggregateType := eventhorizon.AggregateType(LibraryAggregateTypes().Book)
    for _, command := range BookCommandTypes().Values() {
        handler.SetAggregate(aggregateType, eventhorizon.CommandType(command.Name()))
    }
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










type LibraryAggregateType struct {
	name  string
	ordinal int
}

func (o *LibraryAggregateType) Name() string {
    return o.name
}

func (o *LibraryAggregateType) Ordinal() int {
    return o.ordinal
}

func (o *LibraryAggregateType) IsBook() bool {
    return o == _libraryAggregateTypes.Book()
}

type libraryAggregateTypes struct {
	values []*LibraryAggregateType
}

var _libraryAggregateTypes = &libraryAggregateTypes{values: []*LibraryAggregateType{
    {name: "Book", ordinal: 0}},
}

func LibraryAggregateTypes() *libraryAggregateTypes {
	return _libraryAggregateTypes
}

func (o *libraryAggregateTypes) Values() []*LibraryAggregateType {
	return o.values
}

func (o *libraryAggregateTypes) Book() *LibraryAggregateType {
    return _libraryAggregateTypes.values[0]
}

func (o *libraryAggregateTypes) ParseLibraryAggregateType(name string) (ret *LibraryAggregateType, ok bool) {
    switch name {
      case o.Book().Name():
        ret = o.Book()
    }
    return
}



