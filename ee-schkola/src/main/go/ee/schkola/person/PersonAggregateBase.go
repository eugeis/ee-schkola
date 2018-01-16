package person

import (
    "errors"
    "fmt"
    "github.com/eugeis/gee/eh"
    "github.com/looplab/eventhorizon"
    "github.com/looplab/eventhorizon/commandhandler/bus"
    "time"
)
type ChurchCommandHandler struct {
    CreateHandler func (*Create, *Church, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*Delete, *Church, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*Update, *Church, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *ChurchCommandHandler) AddCreatePreparer(preparer func (*Create, *Church) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *Create, entity *Church, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *ChurchCommandHandler) AddDeletePreparer(preparer func (*Delete, *Church) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *Delete, entity *Church, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *ChurchCommandHandler) AddUpdatePreparer(preparer func (*Update, *Church) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *Update, entity *Church, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *ChurchCommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateChurchCommand:
        err = o.CreateHandler(cmd.(*Create), entity.(*Church), store)
    case DeleteChurchCommand:
        err = o.DeleteHandler(cmd.(*Delete), entity.(*Church), store)
    case UpdateChurchCommand:
        err = o.UpdateHandler(cmd.(*Update), entity.(*Church), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *ChurchCommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *Create, entity *Church, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, ChurchAggregateType); err == nil {
            store.StoreEvent(ChurchCreatedEvent, &Created{
                Name: command.Name,
                Address: command.Address,
                Pastor: command.Pastor,
                Contact: command.Contact,
                Association: command.Association,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *Delete, entity *Church, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, ChurchAggregateType); err == nil {
            store.StoreEvent(ChurchDeletedEvent, &Deleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *Update, entity *Church, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, ChurchAggregateType); err == nil {
            store.StoreEvent(ChurchUpdatedEvent, &Updated{
                Name: command.Name,
                Address: command.Address,
                Pastor: command.Pastor,
                Contact: command.Contact,
                Association: command.Association,
                Id: command.Id,}, time.Now())
        }
        return
    }
    return
}


type ChurchEventHandler struct {
    CreatedHandler func (*Created, *Church) (err error)  `json:"createdHandler" eh:"optional"`
    DeletedHandler func (*Deleted, *Church) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdatedHandler func (*Updated, *Church) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *ChurchEventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case ChurchCreatedEvent:
        err = o.CreatedHandler(event.Data().(*Created), entity.(*Church))
    case ChurchDeletedEvent:
        err = o.DeletedHandler(event.Data().(*Deleted), entity.(*Church))
    case ChurchUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*Updated), entity.(*Church))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *ChurchEventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(CreatedEvent, func() eventhorizon.EventData {
		return &Created{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *Created, entity *Church) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, ChurchAggregateType); err == nil {
            entity.Name = event.Name
            entity.Address = event.Address
            entity.Pastor = event.Pastor
            entity.Contact = event.Contact
            entity.Association = event.Association
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DeletedEvent, func() eventhorizon.EventData {
		return &Deleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *Deleted, entity *Church) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, ChurchAggregateType); err == nil {
            *entity = *New@@EMPTY@@()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdatedEvent, func() eventhorizon.EventData {
		return &Updated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *Updated, entity *Church) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, ChurchAggregateType); err == nil {
            entity.Name = event.Name
            entity.Address = event.Address
            entity.Pastor = event.Pastor
            entity.Contact = event.Contact
            entity.Association = event.Association
        }
        return
    }
    return
}


const ChurchAggregateType eventhorizon.AggregateType = "Church"

type ChurchAggregateInitializer struct {
    *eh.AggregateInitializer
    *ChurchCommandHandler
    *ChurchEventHandler
    ProjectorHandler *ChurchEventHandler `json:"projectorHandler" eh:"optional"`
}



func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *ChurchAggregateInitializer) {
    
    commandHandler := &ChurchCommandHandler{}
    eventHandler := &ChurchEventHandler{}
    entityFactory := func() eventhorizon.Entity { return New@@EMPTY@@() }
    ret = &ChurchAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ChurchAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(ChurchAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        ChurchCommandTypes().Literals(), ChurchEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), ChurchCommandHandler: commandHandler, ChurchEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type GraduationCommandHandler struct {
    CreateHandler func (*Create, *Graduation, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*Delete, *Graduation, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*Update, *Graduation, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *GraduationCommandHandler) AddCreatePreparer(preparer func (*Create, *Graduation) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *Create, entity *Graduation, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *GraduationCommandHandler) AddDeletePreparer(preparer func (*Delete, *Graduation) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *Delete, entity *Graduation, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *GraduationCommandHandler) AddUpdatePreparer(preparer func (*Update, *Graduation) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *Update, entity *Graduation, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *GraduationCommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateGraduationCommand:
        err = o.CreateHandler(cmd.(*Create), entity.(*Graduation), store)
    case DeleteGraduationCommand:
        err = o.DeleteHandler(cmd.(*Delete), entity.(*Graduation), store)
    case UpdateGraduationCommand:
        err = o.UpdateHandler(cmd.(*Update), entity.(*Graduation), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *GraduationCommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *Create, entity *Graduation, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, GraduationAggregateType); err == nil {
            store.StoreEvent(GraduationCreatedEvent, &Created{
                Name: command.Name,
                Level: command.Level,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *Delete, entity *Graduation, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, GraduationAggregateType); err == nil {
            store.StoreEvent(GraduationDeletedEvent, &Deleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *Update, entity *Graduation, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, GraduationAggregateType); err == nil {
            store.StoreEvent(GraduationUpdatedEvent, &Updated{
                Name: command.Name,
                Level: command.Level,
                Id: command.Id,}, time.Now())
        }
        return
    }
    return
}


type GraduationEventHandler struct {
    CreatedHandler func (*Created, *Graduation) (err error)  `json:"createdHandler" eh:"optional"`
    DeletedHandler func (*Deleted, *Graduation) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdatedHandler func (*Updated, *Graduation) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *GraduationEventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case GraduationCreatedEvent:
        err = o.CreatedHandler(event.Data().(*Created), entity.(*Graduation))
    case GraduationDeletedEvent:
        err = o.DeletedHandler(event.Data().(*Deleted), entity.(*Graduation))
    case GraduationUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*Updated), entity.(*Graduation))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *GraduationEventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(CreatedEvent, func() eventhorizon.EventData {
		return &Created{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *Created, entity *Graduation) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, GraduationAggregateType); err == nil {
            entity.Name = event.Name
            entity.Level = event.Level
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DeletedEvent, func() eventhorizon.EventData {
		return &Deleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *Deleted, entity *Graduation) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, GraduationAggregateType); err == nil {
            *entity = *New@@EMPTY@@()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdatedEvent, func() eventhorizon.EventData {
		return &Updated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *Updated, entity *Graduation) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, GraduationAggregateType); err == nil {
            entity.Name = event.Name
            entity.Level = event.Level
        }
        return
    }
    return
}


const GraduationAggregateType eventhorizon.AggregateType = "Graduation"

type GraduationAggregateInitializer struct {
    *eh.AggregateInitializer
    *GraduationCommandHandler
    *GraduationEventHandler
    ProjectorHandler *GraduationEventHandler `json:"projectorHandler" eh:"optional"`
}



func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *GraduationAggregateInitializer) {
    
    commandHandler := &GraduationCommandHandler{}
    eventHandler := &GraduationEventHandler{}
    entityFactory := func() eventhorizon.Entity { return New@@EMPTY@@() }
    ret = &GraduationAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(GraduationAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(GraduationAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        GraduationCommandTypes().Literals(), GraduationEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), GraduationCommandHandler: commandHandler, GraduationEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type ProfileCommandHandler struct {
    CreateHandler func (*Create, *Profile, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*Delete, *Profile, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*Update, *Profile, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *ProfileCommandHandler) AddCreatePreparer(preparer func (*Create, *Profile) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *Create, entity *Profile, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *ProfileCommandHandler) AddDeletePreparer(preparer func (*Delete, *Profile) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *Delete, entity *Profile, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *ProfileCommandHandler) AddUpdatePreparer(preparer func (*Update, *Profile) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *Update, entity *Profile, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *ProfileCommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateProfileCommand:
        err = o.CreateHandler(cmd.(*Create), entity.(*Profile), store)
    case DeleteProfileCommand:
        err = o.DeleteHandler(cmd.(*Delete), entity.(*Profile), store)
    case UpdateProfileCommand:
        err = o.UpdateHandler(cmd.(*Update), entity.(*Profile), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *ProfileCommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *Create, entity *Profile, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, ProfileAggregateType); err == nil {
            store.StoreEvent(ProfileCreatedEvent, &Created{
                Gender: command.Gender,
                Name: command.Name,
                BirthName: command.BirthName,
                Birthday: command.Birthday,
                Address: command.Address,
                Contact: command.Contact,
                PhotoData: command.PhotoData,
                Photo: command.Photo,
                Family: command.Family,
                Church: command.Church,
                Education: command.Education,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *Delete, entity *Profile, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, ProfileAggregateType); err == nil {
            store.StoreEvent(ProfileDeletedEvent, &Deleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *Update, entity *Profile, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, ProfileAggregateType); err == nil {
            store.StoreEvent(ProfileUpdatedEvent, &Updated{
                Gender: command.Gender,
                Name: command.Name,
                BirthName: command.BirthName,
                Birthday: command.Birthday,
                Address: command.Address,
                Contact: command.Contact,
                PhotoData: command.PhotoData,
                Photo: command.Photo,
                Family: command.Family,
                Church: command.Church,
                Education: command.Education,
                Id: command.Id,}, time.Now())
        }
        return
    }
    return
}


type ProfileEventHandler struct {
    CreatedHandler func (*Created, *Profile) (err error)  `json:"createdHandler" eh:"optional"`
    DeletedHandler func (*Deleted, *Profile) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdatedHandler func (*Updated, *Profile) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *ProfileEventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case ProfileCreatedEvent:
        err = o.CreatedHandler(event.Data().(*Created), entity.(*Profile))
    case ProfileDeletedEvent:
        err = o.DeletedHandler(event.Data().(*Deleted), entity.(*Profile))
    case ProfileUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*Updated), entity.(*Profile))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *ProfileEventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(CreatedEvent, func() eventhorizon.EventData {
		return &Created{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *Created, entity *Profile) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, ProfileAggregateType); err == nil {
            entity.Gender = event.Gender
            entity.Name = event.Name
            entity.BirthName = event.BirthName
            entity.Birthday = event.Birthday
            entity.Address = event.Address
            entity.Contact = event.Contact
            entity.PhotoData = event.PhotoData
            entity.Photo = event.Photo
            entity.Family = event.Family
            entity.Church = event.Church
            entity.Education = event.Education
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DeletedEvent, func() eventhorizon.EventData {
		return &Deleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *Deleted, entity *Profile) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, ProfileAggregateType); err == nil {
            *entity = *New@@EMPTY@@()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdatedEvent, func() eventhorizon.EventData {
		return &Updated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *Updated, entity *Profile) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, ProfileAggregateType); err == nil {
            entity.Gender = event.Gender
            entity.Name = event.Name
            entity.BirthName = event.BirthName
            entity.Birthday = event.Birthday
            entity.Address = event.Address
            entity.Contact = event.Contact
            entity.PhotoData = event.PhotoData
            entity.Photo = event.Photo
            entity.Family = event.Family
            entity.Church = event.Church
            entity.Education = event.Education
        }
        return
    }
    return
}


const ProfileAggregateType eventhorizon.AggregateType = "Profile"

type ProfileAggregateInitializer struct {
    *eh.AggregateInitializer
    *ProfileCommandHandler
    *ProfileEventHandler
    ProjectorHandler *ProfileEventHandler `json:"projectorHandler" eh:"optional"`
}



func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *ProfileAggregateInitializer) {
    
    commandHandler := &ProfileCommandHandler{}
    eventHandler := &ProfileEventHandler{}
    entityFactory := func() eventhorizon.Entity { return New@@EMPTY@@() }
    ret = &ProfileAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ProfileAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(ProfileAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        ProfileCommandTypes().Literals(), ProfileEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), ProfileCommandHandler: commandHandler, ProfileEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type PersonEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore `json:"eventStore" eh:"optional"`
    eventBus eventhorizon.EventBus `json:"eventBus" eh:"optional"`
    eventPublisher eventhorizon.EventPublisher `json:"eventPublisher" eh:"optional"`
    commandBus *bus.CommandHandler `json:"commandBus" eh:"optional"`
    ChurchAggregateInitializer *ChurchAggregateInitializer `json:"churchAggregateInitializer" eh:"optional"`
    GraduationAggregateInitializer *GraduationAggregateInitializer `json:"graduationAggregateInitializer" eh:"optional"`
    ProfileAggregateInitializer *ProfileAggregateInitializer `json:"profileAggregateInitializer" eh:"optional"`
}

func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *PersonEventhorizonInitializer) {
    churchAggregateInitializer := New@@EMPTY@@(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    graduationAggregateInitializer := New@@EMPTY@@(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    profileAggregateInitializer := New@@EMPTY@@(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    ret = &PersonEventhorizonInitializer{
        eventStore: eventStore,
        eventBus: eventBus,
        eventPublisher: eventPublisher,
        commandBus: commandBus,
        ChurchAggregateInitializer: churchAggregateInitializer,
        GraduationAggregateInitializer: graduationAggregateInitializer,
        ProfileAggregateInitializer: profileAggregateInitializer,
    }
    return
}

func (o *PersonEventhorizonInitializer) Setup() (err error) {
    
    if err = o.ChurchAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.GraduationAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.ProfileAggregateInitializer.Setup(); err != nil {
        return
    }

    return
}









