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
	//ret.CommandHandler = NewBookAggregateCommandHandler(ret)
    return
}

func (o *BookAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type BookAggregate struct {
    *eventhorizon.AggregateBase
    *Book
    eventhorizon.CommandHandler
}



type BookCommandHandler struct {
    CreateHandler  func (*CreateBook, *BookAggregate) error
    DeleteHandler  func (*DeleteBook, *BookAggregate) error
    UpdateHandler  func (*UpdateBook, *BookAggregate) error
    RegisterHandler  func (*RegisterBook, *BookAggregate) error
    UnregisterHandler  func (*UnregisterBook, *BookAggregate) error
    ChangeHandler  func (*ChangeBook, *BookAggregate) error
    ChangeLocationHandler  func (*ChangeBookLocation, *BookAggregate) error
}

func (o *BookCommandHandler) HandleCommand(ctx *context.Context, cmd eventhorizon.Command, aggregate *BookAggregate) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateBookCommand:
        ret = o.CreateHandler(cmd.(*CreateBook), aggregate)
    case DeleteBookCommand:
        ret = o.DeleteHandler(cmd.(*DeleteBook), aggregate)
    case UpdateBookCommand:
        ret = o.UpdateHandler(cmd.(*UpdateBook), aggregate)
    case RegisterBookCommand:
        ret = o.RegisterHandler(cmd.(*RegisterBook), aggregate)
    case UnregisterBookCommand:
        ret = o.UnregisterHandler(cmd.(*UnregisterBook), aggregate)
    case ChangeBookCommand:
        ret = o.ChangeHandler(cmd.(*ChangeBook), aggregate)
    case ChangeBookLocationCommand:
        ret = o.ChangeLocationHandler(cmd.(*ChangeBookLocation), aggregate)
    default:
		ret = errors.New(fmt.Sprintf("Wrong comand type '%v' for aggregate '%v", cmd.CommandType(), aggregate))
	}
    return ret
    
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
    *BookCommandHandler
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









