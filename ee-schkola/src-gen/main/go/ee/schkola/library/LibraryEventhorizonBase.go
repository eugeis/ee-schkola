package library

import (
    "errors"
    "fmt"
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

type BookCommandHandler struct {
    CreateHandler  func (*CreateBook, *Book, eh.AggregateStoreEvent) error
    DeleteHandler  func (*DeleteBook, *Book, eh.AggregateStoreEvent) error
    UpdateHandler  func (*UpdateBook, *Book, eh.AggregateStoreEvent) error
    RegisterHandler  func (*RegisterBook, *Book, eh.AggregateStoreEvent) error
    UnregisterHandler  func (*UnregisterBook, *Book, eh.AggregateStoreEvent) error
    ChangeHandler  func (*ChangeBook, *Book, eh.AggregateStoreEvent) error
    ChangeLocationHandler  func (*ChangeBookLocation, *Book, eh.AggregateStoreEvent) error
}

func (o *BookCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateBookCommand:
        ret = o.CreateHandler(cmd.(*CreateBook), entity.(*Book), store)
    case DeleteBookCommand:
        ret = o.DeleteHandler(cmd.(*DeleteBook), entity.(*Book), store)
    case UpdateBookCommand:
        ret = o.UpdateHandler(cmd.(*UpdateBook), entity.(*Book), store)
    case RegisterBookCommand:
        ret = o.RegisterHandler(cmd.(*RegisterBook), entity.(*Book), store)
    case UnregisterBookCommand:
        ret = o.UnregisterHandler(cmd.(*UnregisterBook), entity.(*Book), store)
    case ChangeBookCommand:
        ret = o.ChangeHandler(cmd.(*ChangeBook), entity.(*Book), store)
    case ChangeBookLocationCommand:
        ret = o.ChangeLocationHandler(cmd.(*ChangeBookLocation), entity.(*Book), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return ret
    
}



type BookEventHandler struct {
    CreatedHandler  func (*BookCreated, *Book) error
    DeletedHandler  func (*BookDeleted, *Book) error
    UpdatedHandler  func (*BookUpdated, *Book) error
}

func (o *BookEventHandler) Apply(event eventhorizon.Event, entity interface{}) error {
    
    var ret error
    switch event.EventType() {
    case BookCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*BookCreated), entity.(*Book))
    case BookDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*BookDeleted), entity.(*Book))
    case BookUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*BookUpdated), entity.(*Book))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return ret
    
}



const BookAggregateType eventhorizon.AggregateType = "BookAggregate"

func NewBookAggregate(id eventhorizon.UUID,
    commandHandler eh.DelegateCommandHandler,
    eventHandler eh.DelegateEventHandler) (ret *BookAggregate) {
    ret = &BookAggregate{
		AggregateBase: eh.NewAggregateBase(BookAggregateType, id, commandHandler, eventHandler, &Book{}),
    }
    return
}

type BookAggregate struct {
    *eh.AggregateBase
}



func NewBookAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *BookAggregateInitializer) {
    commandHandler := &BookCommandHandler{}
    eventHandler := &BookEventHandler{}
	ret = &BookAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(BookAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate { return NewBookAggregate(id, commandHandler, eventHandler) },
        BookCommandTypes().Literals(), BookEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
        BookCommandHandler: commandHandler,
        BookEventHandler: eventHandler,
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
    *BookCommandHandler
    *BookEventHandler
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









