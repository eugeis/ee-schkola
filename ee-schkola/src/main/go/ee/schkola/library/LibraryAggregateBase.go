package library

import (
    "errors"
    "fmt"
    "github.com/eugeis/gee/eh"
    "github.com/looplab/eventhorizon"
    "github.com/looplab/eventhorizon/commandhandler/bus"
    "time"
)
type BookCommandHandler struct {
    CreateHandler func (*Create, *Book, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*Delete, *Book, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    ChangeLocationHandler func (*ChangeLocation, *Book, eh.AggregateStoreEvent) (err error)  `json:"changeLocationHandler" eh:"optional"`
    UpdateHandler func (*Update, *Book, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *BookCommandHandler) AddCreatePreparer(preparer func (*Create, *Book) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *Create, entity *Book, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *BookCommandHandler) AddDeletePreparer(preparer func (*Delete, *Book) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *Delete, entity *Book, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *BookCommandHandler) AddChangeLocationPreparer(preparer func (*ChangeLocation, *Book) (err error) ) {
    prevHandler := o.ChangeLocationHandler
	o.ChangeLocationHandler = func(command *ChangeLocation, entity *Book, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *BookCommandHandler) AddUpdatePreparer(preparer func (*Update, *Book) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *Update, entity *Book, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *BookCommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateBookCommand:
        err = o.CreateHandler(cmd.(*Create), entity.(*Book), store)
    case DeleteBookCommand:
        err = o.DeleteHandler(cmd.(*Delete), entity.(*Book), store)
    case ChangeBookLocationCommand:
        err = o.ChangeLocationHandler(cmd.(*ChangeLocation), entity.(*Book), store)
    case UpdateBookCommand:
        err = o.UpdateHandler(cmd.(*Update), entity.(*Book), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *BookCommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *Create, entity *Book, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, BookAggregateType); err == nil {
            store.StoreEvent(BookCreatedEvent, &Created{
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
    o.DeleteHandler = func(command *Delete, entity *Book, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, BookAggregateType); err == nil {
            store.StoreEvent(BookDeletedEvent, &Deleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.ChangeLocationHandler = func(command *ChangeLocation, entity *Book, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, BookAggregateType); err == nil {
            store.StoreEvent(ChangeLocationedBookEvent, &ChangeLocationed{
                Location: command.Location,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *Update, entity *Book, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, BookAggregateType); err == nil {
            store.StoreEvent(BookUpdatedEvent, &Updated{
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


type BookEventHandler struct {
    CreatedHandler func (*Created, *Book) (err error)  `json:"createdHandler" eh:"optional"`
    DeletedHandler func (*Deleted, *Book) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdatedHandler func (*Updated, *Book) (err error)  `json:"updatedHandler" eh:"optional"`
    ChangeLocationedHandler func (*ChangeLocationed, *Book) (err error)  `json:"changeLocationedHandler" eh:"optional"`
}

func (o *BookEventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case BookCreatedEvent:
        err = o.CreatedHandler(event.Data().(*Created), entity.(*Book))
    case BookDeletedEvent:
        err = o.DeletedHandler(event.Data().(*Deleted), entity.(*Book))
    case BookUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*Updated), entity.(*Book))
    case ChangeLocationedBookEvent:
        err = o.ChangeLocationedHandler(event.Data().(*ChangeLocationed), entity.(*Book))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *BookEventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(CreatedEvent, func() eventhorizon.EventData {
		return &Created{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *Created, entity *Book) (err error) {
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
    eventhorizon.RegisterEventData(DeletedEvent, func() eventhorizon.EventData {
		return &Deleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *Deleted, entity *Book) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, BookAggregateType); err == nil {
            *entity = *New@@EMPTY@@()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdatedEvent, func() eventhorizon.EventData {
		return &Updated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *Updated, entity *Book) (err error) {
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
    eventhorizon.RegisterEventData(ChangeLocationedEvent, func() eventhorizon.EventData {
		return &ChangeLocationed{}
	})

    //default handler implementation
    o.ChangeLocationedHandler = func(event *ChangeLocationed, entity *Book) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, BookAggregateType); err == nil {
            entity.Location = event.Location
        }
        return
    }
    return
}


const BookAggregateType eventhorizon.AggregateType = "Book"

type BookAggregateInitializer struct {
    *eh.AggregateInitializer
    *BookCommandHandler
    *BookEventHandler
    ProjectorHandler *BookEventHandler `json:"projectorHandler" eh:"optional"`
}



func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *BookAggregateInitializer) {
    
    commandHandler := &BookCommandHandler{}
    eventHandler := &BookEventHandler{}
    entityFactory := func() eventhorizon.Entity { return New@@EMPTY@@() }
    ret = &BookAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(BookAggregateType,
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
    BookAggregateInitializer *BookAggregateInitializer `json:"bookAggregateInitializer" eh:"optional"`
}

func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *LibraryEventhorizonInitializer) {
    bookAggregateInitializer := New@@EMPTY@@(eventStore, eventBus, eventPublisher, commandBus, readRepos)
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









