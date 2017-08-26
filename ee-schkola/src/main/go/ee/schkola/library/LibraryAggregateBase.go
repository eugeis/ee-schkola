package library

import (
    "errors"
    "fmt"
    "github.com/eugeis/gee/eh"
    "github.com/looplab/eventhorizon"
    "time"
)
type BookCommandHandler struct {
    CreateHandler func (*CreateBook, *Book, eh.AggregateStoreEvent) error
    DeleteHandler func (*DeleteBook, *Book, eh.AggregateStoreEvent) error
    ChangeLocationHandler func (*ChangeBookLocation, *Book, eh.AggregateStoreEvent) error
    UpdateHandler func (*UpdateBook, *Book, eh.AggregateStoreEvent) error
}

func (o *BookCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (ret error) {
    switch cmd.CommandType() {
    case CreateBookCommand:
        ret = o.CreateHandler(cmd.(*CreateBook), entity.(*Book), store)
    case DeleteBookCommand:
        ret = o.DeleteHandler(cmd.(*DeleteBook), entity.(*Book), store)
    case ChangeBookLocationCommand:
        ret = o.ChangeLocationHandler(cmd.(*ChangeBookLocation), entity.(*Book), store)
    case UpdateBookCommand:
        ret = o.UpdateHandler(cmd.(*UpdateBook), entity.(*Book), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *BookCommandHandler) SetupCommandHandler() (ret error) {
    if o.CreateHandler == nil {
        o.CreateHandler = func(command *CreateBook, entity *Book, store eh.AggregateStoreEvent) (ret error) {
            if ret = eh.ValidateNewId(entity.Id, command.Id, BookAggregateType); ret == nil {store.StoreEvent(BookCreatedEvent, &BookCreated{
                    Id: command.Id,
                    Title: command.Title,
                    Description: command.Description,
                    Language: command.Language,
                    ReleaseDate: command.ReleaseDate,
                    Edition: command.Edition,
                    Category: command.Category,
                    Author: command.Author,
                    Location: command.Location,}, time.Now())
            }
            return
        }
    }
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteBook, entity *Book, store eh.AggregateStoreEvent) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, command.Id, BookAggregateType); ret == nil {
                store.StoreEvent(BookDeletedEvent, &BookDeleted{
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    if o.ChangeLocationHandler == nil {
        o.ChangeLocationHandler = func(command *ChangeBookLocation, entity *Book, store eh.AggregateStoreEvent) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, command.Id, BookAggregateType); ret == nil {
                store.StoreEvent(LocationChangedBookEvent, &LocationChangedBook{
                    Location: command.Location,
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    if o.UpdateHandler == nil {
        o.UpdateHandler = func(command *UpdateBook, entity *Book, store eh.AggregateStoreEvent) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, command.Id, BookAggregateType); ret == nil {
                store.StoreEvent(BookUpdatedEvent, &BookUpdated{
                    Id: command.Id,
                    Title: command.Title,
                    Description: command.Description,
                    Language: command.Language,
                    ReleaseDate: command.ReleaseDate,
                    Edition: command.Edition,
                    Category: command.Category,
                    Author: command.Author,
                    Location: command.Location,}, time.Now())
            }
            return
        }
    }
    return
}


type BookEventHandler struct {
    LocationChangedHandler func (*LocationChangedBook, *Book) error
    CreatedHandler func (*BookCreated, *Book) error
    DeletedHandler func (*BookDeleted, *Book) error
    UpdatedHandler func (*BookUpdated, *Book) error
}

func (o *BookEventHandler) Apply(event eventhorizon.Event, entity interface{}) (ret error) {
    switch event.EventType() {
    case LocationChangedBookEvent:
        ret = o.LocationChangedHandler(event.Data().(*LocationChangedBook), entity.(*Book))
    case BookCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*BookCreated), entity.(*Book))
    case BookDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*BookDeleted), entity.(*Book))
    case BookUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*BookUpdated), entity.(*Book))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *BookEventHandler) SetupEventHandler() (ret error) {
    if o.LocationChangedHandler == nil {
        o.LocationChangedHandler = func(event *LocationChangedBook, entity *Book) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, event.Id, BookAggregateType); ret == nil {
                entity.Location = event.Location
            }
            return
        }
    }
    if o.CreatedHandler == nil {
        o.CreatedHandler = func(event *BookCreated, entity *Book) (ret error) {
            if ret = eh.ValidateNewId(entity.Id, event.Id, BookAggregateType); ret == nil {
                entity.Id = event.Id
                entity.Title = event.Title
                entity.Description = event.Description
                entity.Language = event.Language
                entity.ReleaseDate = event.ReleaseDate
                entity.Edition = event.Edition
                entity.Category = event.Category
                entity.Author = event.Author
                entity.Location = event.Location
            }
            return
        }
    }
    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *BookDeleted, entity *Book) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, event.Id, BookAggregateType); ret == nil {
                *entity = *NewBook()
            }
            return
        }
    }
    if o.UpdatedHandler == nil {
        o.UpdatedHandler = func(event *BookUpdated, entity *Book) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, event.Id, BookAggregateType); ret == nil {
                entity.Title = event.Title
                entity.Description = event.Description
                entity.Language = event.Language
                entity.ReleaseDate = event.ReleaseDate
                entity.Edition = event.Edition
                entity.Category = event.Category
                entity.Author = event.Author
                entity.Location = event.Location
            }
            return
        }
    }
    return
}


const BookAggregateType eventhorizon.AggregateType = "Book"

type BookAggregateInitializer struct {
    *eh.AggregateInitializer
    *BookCommandHandler
    *BookEventHandler
    ProjectorHandler *BookEventHandler
}



func NewBookAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus, readRepos func (string) eventhorizon.ReadWriteRepo) (ret *BookAggregateInitializer) {
    
    commandHandler := &BookCommandHandler{}
    eventHandler := &BookEventHandler{}
    modelFactory := func() interface{} { return NewBook() }
    ret = &BookAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(BookAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(BookAggregateType, id, commandHandler, eventHandler, modelFactory())
        }, modelFactory,
        BookCommandTypes().Literals(), BookEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), BookCommandHandler: commandHandler, BookEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type LibraryEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore
    eventBus eventhorizon.EventBus
    eventPublisher eventhorizon.EventPublisher
    commandBus eventhorizon.CommandBus
    BookAggregateInitializer *BookAggregateInitializer
}

func NewLibraryEventhorizonInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus, readRepos func (string) eventhorizon.ReadWriteRepo) (ret *LibraryEventhorizonInitializer) {
    bookAggregateInitializer := NewBookAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    ret = &LibraryEventhorizonInitializer{
        eventStore: eventStore,
        eventBus: eventBus,
        eventPublisher: eventPublisher,
        commandBus: commandBus,
        BookAggregateInitializer: bookAggregateInitializer,
    }
    return
}

func (o *LibraryEventhorizonInitializer) Setup() (ret error) {
    
    if ret = o.BookAggregateInitializer.Setup(); ret != nil {
        return
    }

    return
}









