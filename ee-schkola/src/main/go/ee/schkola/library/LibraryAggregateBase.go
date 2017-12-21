package library

import (
    "errors"
    "fmt"
    "github.com/eugeis/gee/eh"
    "github.com/looplab/eventhorizon"
    "github.com/looplab/eventhorizon/commandhandler/bus"
    "time"
)
type CommandHandler struct {
    CreateHandler func (*CreateBook, *Book, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*DeleteBook, *Book, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    ChangeLocationHandler func (*ChangeBookLocation, *Book, eh.AggregateStoreEvent) (err error)  `json:"changeLocationHandler" eh:"optional"`
    UpdateHandler func (*UpdateBook, *Book, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *CommandHandler) AddCreatePreparer(preparer func (*CreateBook, *Book) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *CreateBook, entity *Book, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddDeletePreparer(preparer func (*DeleteBook, *Book) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *DeleteBook, entity *Book, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddChangeLocationPreparer(preparer func (*ChangeBookLocation, *Book) (err error) ) {
    prevHandler := o.ChangeLocationHandler
	o.ChangeLocationHandler = func(command *ChangeBookLocation, entity *Book, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddUpdatePreparer(preparer func (*UpdateBook, *Book) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *UpdateBook, entity *Book, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateBookCommand:
        err = o.CreateHandler(cmd.(*CreateBook), entity.(*Book), store)
    case DeleteBookCommand:
        err = o.DeleteHandler(cmd.(*DeleteBook), entity.(*Book), store)
    case ChangeBookLocationCommand:
        err = o.ChangeLocationHandler(cmd.(*ChangeBookLocation), entity.(*Book), store)
    case UpdateBookCommand:
        err = o.UpdateHandler(cmd.(*UpdateBook), entity.(*Book), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *CommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *CreateBook, entity *Book, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, BookAggregateType); err == nil {
            store.StoreEvent(BookCreatedEvent, &BookCreated{
                Location: command.Location,
                Title: command.Title,
                Description: command.Description,
                Language: command.Language,
                ReleaseDate: command.ReleaseDate,
                Edition: command.Edition,
                Category: command.Category,
                Author: command.Author,
                Location: command.Location,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *DeleteBook, entity *Book, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, BookAggregateType); err == nil {
            store.StoreEvent(BookDeletedEvent, &BookDeleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.ChangeLocationHandler = func(command *ChangeBookLocation, entity *Book, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, BookAggregateType); err == nil {
            store.StoreEvent(ChangeLocationedBookEvent, &ChangeLocationedBook{
                Location: command.Location,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *UpdateBook, entity *Book, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, BookAggregateType); err == nil {
            store.StoreEvent(BookUpdatedEvent, &BookUpdated{
                Location: command.Location,
                Title: command.Title,
                Description: command.Description,
                Language: command.Language,
                ReleaseDate: command.ReleaseDate,
                Edition: command.Edition,
                Category: command.Category,
                Author: command.Author,
                Location: command.Location,
                Id: command.Id,}, time.Now())
        }
        return
    }
    return
}


type EventHandler struct {
    CreateHandler func (*CreateBook, *Book) (err error)  `json:"createHandler" eh:"optional"`
    CreatedHandler func (*BookCreated, *Book) (err error)  `json:"createdHandler" eh:"optional"`
    DeleteHandler func (*DeleteBook, *Book) (err error)  `json:"deleteHandler" eh:"optional"`
    DeletedHandler func (*BookDeleted, *Book) (err error)  `json:"deletedHandler" eh:"optional"`
    ChangeLocationHandler func (*ChangeBookLocation, *Book) (err error)  `json:"changeLocationHandler" eh:"optional"`
    UpdateHandler func (*UpdateBook, *Book) (err error)  `json:"updateHandler" eh:"optional"`
    UpdatedHandler func (*BookUpdated, *Book) (err error)  `json:"updatedHandler" eh:"optional"`
    ChangeLocationedHandler func (*ChangeLocationedBook, *Book) (err error)  `json:"changeLocationedHandler" eh:"optional"`
}

func (o *EventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case BookCreateEvent:
        err = o.CreateHandler(event.Data().(*CreateBook), entity.(*Book))
    case BookCreatedEvent:
        err = o.CreatedHandler(event.Data().(*BookCreated), entity.(*Book))
    case BookDeleteEvent:
        err = o.DeleteHandler(event.Data().(*DeleteBook), entity.(*Book))
    case BookDeletedEvent:
        err = o.DeletedHandler(event.Data().(*BookDeleted), entity.(*Book))
    case ChangeLocationBookEvent:
        err = o.ChangeLocationHandler(event.Data().(*ChangeBookLocation), entity.(*Book))
    case BookUpdateEvent:
        err = o.UpdateHandler(event.Data().(*UpdateBook), entity.(*Book))
    case BookUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*BookUpdated), entity.(*Book))
    case ChangeLocationedBookEvent:
        err = o.ChangeLocationedHandler(event.Data().(*ChangeLocationedBook), entity.(*Book))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *EventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(CreateBookEvent, func() eventhorizon.EventData {
		return &CreateBook{}
	})

    //default handler implementation
    o.CreateHandler = func(event *CreateBook, entity *Book) (err error) {
        //err = eh.EventHandlerNotImplemented(CreateBookEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(BookCreatedEvent, func() eventhorizon.EventData {
		return &BookCreated{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *BookCreated, entity *Book) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, BookAggregateType); err == nil {
            entity.Location = event.Location
            entity.Title = event.Title
            entity.Description = event.Description
            entity.Language = event.Language
            entity.ReleaseDate = event.ReleaseDate
            entity.Edition = event.Edition
            entity.Category = event.Category
            entity.Author = event.Author
            entity.Location = event.Location
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DeleteBookEvent, func() eventhorizon.EventData {
		return &DeleteBook{}
	})

    //default handler implementation
    o.DeleteHandler = func(event *DeleteBook, entity *Book) (err error) {
        //err = eh.EventHandlerNotImplemented(DeleteBookEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(BookDeletedEvent, func() eventhorizon.EventData {
		return &BookDeleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *BookDeleted, entity *Book) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, BookAggregateType); err == nil {
            *entity = *NewBook()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(ChangeBookLocationEvent, func() eventhorizon.EventData {
		return &ChangeBookLocation{}
	})

    //default handler implementation
    o.ChangeLocationHandler = func(event *ChangeBookLocation, entity *Book) (err error) {
        //err = eh.EventHandlerNotImplemented(ChangeBookLocationEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdateBookEvent, func() eventhorizon.EventData {
		return &UpdateBook{}
	})

    //default handler implementation
    o.UpdateHandler = func(event *UpdateBook, entity *Book) (err error) {
        //err = eh.EventHandlerNotImplemented(UpdateBookEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(BookUpdatedEvent, func() eventhorizon.EventData {
		return &BookUpdated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *BookUpdated, entity *Book) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, BookAggregateType); err == nil {
            entity.Location = event.Location
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

    //register event object factory
    eventhorizon.RegisterEventData(ChangeLocationedBookEvent, func() eventhorizon.EventData {
		return &ChangeLocationedBook{}
	})

    //default handler implementation
    o.ChangeLocationedHandler = func(event *ChangeLocationedBook, entity *Book) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, BookAggregateType); err == nil {
            entity.Location = event.Location
        }
        return
    }
    return
}


const BookAggregateType eventhorizon.AggregateType = "Book"

type AggregateInitializer struct {
    *eh.AggregateInitializer
    *CommandHandler
    *EventHandler
    ProjectorHandler *EventHandler `json:"projectorHandler" eh:"optional"`
}



func NewBookAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *AggregateInitializer) {
    
    commandHandler := &BookCommandHandler{}
    eventHandler := &BookEventHandler{}
    entityFactory := func() eventhorizon.Entity { return NewBook() }
    ret = &AggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(BookAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(BookAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        BookCommandTypes().Literals(), BookEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), BookCommandHandler: commandHandler, BookEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type LibraryEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore `json:"eventStore" eh:"optional"`
    eventBus eventhorizon.EventBus `json:"eventBus" eh:"optional"`
    eventPublisher eventhorizon.EventPublisher `json:"eventPublisher" eh:"optional"`
    commandBus *bus.CommandHandler `json:"commandBus" eh:"optional"`
    BookAggregateInitializer *AggregateInitializer `json:"bookAggregateInitializer" eh:"optional"`
}

func NewLibraryEventhorizonInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *LibraryEventhorizonInitializer) {
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

func (o *LibraryEventhorizonInitializer) Setup() (err error) {
    
    if err = o.BookAggregateInitializer.Setup(); err != nil {
        return
    }

    return
}









