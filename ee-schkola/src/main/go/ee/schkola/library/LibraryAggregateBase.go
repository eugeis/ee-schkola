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
    CreateHandler func (*Create, *Book, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*Delete, *Book, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    ChangeLocationHandler func (*ChangeLocation, *Book, eh.AggregateStoreEvent) (err error)  `json:"changeLocationHandler" eh:"optional"`
    UpdateHandler func (*Update, *Book, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *CommandHandler) AddCreatePreparer(preparer func (*Create, *Book) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *Create, entity *Book, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddDeletePreparer(preparer func (*Delete, *Book) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *Delete, entity *Book, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddChangeLocationPreparer(preparer func (*ChangeLocation, *Book) (err error) ) {
    prevHandler := o.ChangeLocationHandler
	o.ChangeLocationHandler = func(command *ChangeLocation, entity *Book, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddUpdatePreparer(preparer func (*Update, *Book) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *Update, entity *Book, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateCommand:
        err = o.CreateHandler(cmd.(*Create), entity.(*Book), store)
    case DeleteCommand:
        err = o.DeleteHandler(cmd.(*Delete), entity.(*Book), store)
    case ChangeLocationCommand:
        err = o.ChangeLocationHandler(cmd.(*ChangeLocation), entity.(*Book), store)
    case UpdateCommand:
        err = o.UpdateHandler(cmd.(*Update), entity.(*Book), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *CommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *Create, entity *Book, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, BookAggregateType); err == nil {
            store.StoreEvent(createdEvent, &Created{
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
            store.StoreEvent(deletedEvent, &Deleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.ChangeLocationHandler = func(command *ChangeLocation, entity *Book, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, BookAggregateType); err == nil {
            store.StoreEvent(ChangeLocationedEvent, &ChangeLocationed{
                Location: command.Location,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *Update, entity *Book, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, BookAggregateType); err == nil {
            store.StoreEvent(updatedEvent, &Updated{
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
    CreateHandler func (*Create, *Book) (err error)  `json:"createHandler" eh:"optional"`
    CreatedHandler func (*Created, *Book) (err error)  `json:"createdHandler" eh:"optional"`
    DeleteHandler func (*Delete, *Book) (err error)  `json:"deleteHandler" eh:"optional"`
    DeletedHandler func (*Deleted, *Book) (err error)  `json:"deletedHandler" eh:"optional"`
    ChangeLocationHandler func (*ChangeLocation, *Book) (err error)  `json:"changeLocationHandler" eh:"optional"`
    UpdateHandler func (*Update, *Book) (err error)  `json:"updateHandler" eh:"optional"`
    UpdatedHandler func (*Updated, *Book) (err error)  `json:"updatedHandler" eh:"optional"`
    ChangeLocationedHandler func (*ChangeLocationed, *Book) (err error)  `json:"changeLocationedHandler" eh:"optional"`
}

func (o *EventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case CreateEvent:
        err = o.CreateHandler(event.Data().(*Create), entity.(*Book))
    case CreatedEvent:
        err = o.CreatedHandler(event.Data().(*Created), entity.(*Book))
    case DeleteEvent:
        err = o.DeleteHandler(event.Data().(*Delete), entity.(*Book))
    case DeletedEvent:
        err = o.DeletedHandler(event.Data().(*Deleted), entity.(*Book))
    case ChangeLocationEvent:
        err = o.ChangeLocationHandler(event.Data().(*ChangeLocation), entity.(*Book))
    case UpdateEvent:
        err = o.UpdateHandler(event.Data().(*Update), entity.(*Book))
    case UpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*Updated), entity.(*Book))
    case ChangeLocationedEvent:
        err = o.ChangeLocationedHandler(event.Data().(*ChangeLocationed), entity.(*Book))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *EventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(CreateEvent, func() eventhorizon.EventData {
		return &Create{}
	})

    //default handler implementation
    o.CreateHandler = func(event *Create, entity *Book) (err error) {
        //err = eh.EventHandlerNotImplemented(CreateEvent)
        return
    }

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
    eventhorizon.RegisterEventData(DeleteEvent, func() eventhorizon.EventData {
		return &Delete{}
	})

    //default handler implementation
    o.DeleteHandler = func(event *Delete, entity *Book) (err error) {
        //err = eh.EventHandlerNotImplemented(DeleteEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DeletedEvent, func() eventhorizon.EventData {
		return &Deleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *Deleted, entity *Book) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, BookAggregateType); err == nil {
            *entity = *NewBook()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(ChangeLocationEvent, func() eventhorizon.EventData {
		return &ChangeLocation{}
	})

    //default handler implementation
    o.ChangeLocationHandler = func(event *ChangeLocation, entity *Book) (err error) {
        //err = eh.EventHandlerNotImplemented(ChangeLocationEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdateEvent, func() eventhorizon.EventData {
		return &Update{}
	})

    //default handler implementation
    o.UpdateHandler = func(event *Update, entity *Book) (err error) {
        //err = eh.EventHandlerNotImplemented(UpdateEvent)
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

type AggregateInitializer struct {
    *eh.AggregateInitializer
    *CommandHandler
    *EventHandler
    ProjectorHandler *EventHandler `json:"projectorHandler" eh:"optional"`
}



func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
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









