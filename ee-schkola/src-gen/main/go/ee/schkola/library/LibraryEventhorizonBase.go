package library

import (
    "github.com/looplab/eventhorizon"
)
type BookAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func (o *BookAggregateInitializer) setup() string {
    panic("Not implemented yet.")
}

func (o *BookAggregateInitializer) registerCommands() string {
    panic("Not implemented yet.")
}


type LibraryEventhorizonInitializer struct {
    Store  *eventhorizon.EventStore
    EventBus  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    CommandBus  *eventhorizon.CommandBus
}

func (o *LibraryEventhorizonInitializer) setup() string {
    panic("Not implemented yet.")
}

func (o *LibraryEventhorizonInitializer) registerCommands() string {
    panic("Not implemented yet.")
}


type BookAggregate struct {
    *eventhorizon.AggregateBase
    *Book
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
      case "Book":
        ret = o.Book()
    }
    return
}



