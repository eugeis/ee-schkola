package person

import (
    "errors"
    "fmt"
    "github.com/eugeis/gee/eh"
    "github.com/looplab/eventhorizon"
)
type ChurchCommandHandler struct {
    CreateHandler func (*CreateChurch, *Church, eh.AggregateStoreEvent) error
    DeleteHandler func (*DeleteChurch, *Church, eh.AggregateStoreEvent) error
    UpdateHandler func (*UpdateChurch, *Church, eh.AggregateStoreEvent) error
}

func (o *ChurchCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (ret error) {
    switch cmd.CommandType() {
    case CreateChurchCommand:
        ret = o.CreateHandler(cmd.(*CreateChurch), entity.(*Church), store)
    case DeleteChurchCommand:
        ret = o.DeleteHandler(cmd.(*DeleteChurch), entity.(*Church), store)
    case UpdateChurchCommand:
        ret = o.UpdateHandler(cmd.(*UpdateChurch), entity.(*Church), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *ChurchCommandHandler) SetupCommandHandler() (ret error) {
    if o.CreateHandler == nil {
        o.CreateHandler = func(command *CreateChurch, entity *Church, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityAlreadyExists(entity.Id, ChurchAggregateType)
            } else {
                store.StoreEvent(ChurchCreatedEvent, &ChurchCreated{
                    Id: command.Id,
                    Name: command.Name,
                    Address: command.Address,
                    Pastor: command.Pastor,
                    Contact: command.Contact,})
            }
            return
        }
    }
    
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteChurch, entity *Church, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, ChurchAggregateType)
            } else if entity.Id != command.Id {
                ret = eh.IdsDismatch(entity.Id, command.Id, ChurchAggregateType)
            } else {
                store.StoreEvent(ChurchDeletedEvent, &ChurchDeleted{
                    Id: command.Id,})
            }
            return
        }
    }
    
    if o.UpdateHandler == nil {
        o.UpdateHandler = func(command *UpdateChurch, entity *Church, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, ChurchAggregateType)
            } else if entity.Id != command.Id {
                ret = eh.IdsDismatch(entity.Id, command.Id, ChurchAggregateType)
            } else {
                store.StoreEvent(ChurchUpdatedEvent, &ChurchUpdated{
                    Id: command.Id,
                    Name: command.Name,
                    Address: command.Address,
                    Pastor: command.Pastor,
                    Contact: command.Contact,})
            }
            return
        }
    }
    
    return
}


type ChurchEventHandler struct {
    CreatedHandler func (*ChurchCreated, *Church) error
    DeletedHandler func (*ChurchDeleted, *Church) error
    UpdatedHandler func (*ChurchUpdated, *Church) error
}

func (o *ChurchEventHandler) Apply(event eventhorizon.Event, entity interface{}) (ret error) {
    switch event.EventType() {
    case ChurchCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*ChurchCreated), entity.(*Church))
    case ChurchDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*ChurchDeleted), entity.(*Church))
    case ChurchUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*ChurchUpdated), entity.(*Church))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *ChurchEventHandler) SetupEventHandler() (ret error) {
    if o.CreatedHandler == nil {
        o.CreatedHandler = func(event *ChurchCreated, entity *Church) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityAlreadyExists(entity.Id, ChurchAggregateType)
            } else {
                entity.Id = event.Id
                entity.Name = event.Name
                entity.Address = event.Address
                entity.Pastor = event.Pastor
                entity.Contact = event.Contact
            }
            return
        }
    }
    
    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *ChurchDeleted, entity *Church) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityNotExists(entity.Id, ChurchAggregateType)
            } else if entity.Id != event.Id {
                ret = eh.IdsDismatch(entity.Id, event.Id, ChurchAggregateType)
            } else {
                entity.Id = ""
            }
            return
        }
    }
    
    if o.UpdatedHandler == nil {
        o.UpdatedHandler = func(event *ChurchUpdated, entity *Church) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityNotExists(entity.Id, ChurchAggregateType)
            } else if entity.Id != event.Id {
                ret = eh.IdsDismatch(entity.Id, event.Id, ChurchAggregateType)
            } else {
                entity.Id = event.Id
                entity.Name = event.Name
                entity.Address = event.Address
                entity.Pastor = event.Pastor
                entity.Contact = event.Contact
            }
            return
        }
    }
    
    return
}


const ChurchAggregateType eventhorizon.AggregateType = "Church"

type ChurchAggregateInitializer struct {
    *eh.AggregateInitializer
    *ChurchCommandHandler
    *ChurchEventHandler
}

func NewChurchAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus) (ret *ChurchAggregateInitializer) {
    
    commandHandler := &ChurchCommandHandler{}
    eventHandler := &ChurchEventHandler{}
	ret = &ChurchAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ChurchAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(ChurchAggregateType, id, commandHandler, eventHandler, NewChurch())
        },
        ChurchCommandTypes().Literals(), ChurchEventTypes().Literals(),
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus),
        ChurchCommandHandler: commandHandler,
        ChurchEventHandler: eventHandler,
    }

    return
}


func (o *ChurchAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ChurchEventTypes().ChurchCreated())
}

func (o *ChurchAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ChurchEventTypes().ChurchDeleted())
}

func (o *ChurchAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ChurchEventTypes().ChurchUpdated())
}



type GraduationCommandHandler struct {
    CreateHandler func (*CreateGraduation, *Graduation, eh.AggregateStoreEvent) error
    DeleteHandler func (*DeleteGraduation, *Graduation, eh.AggregateStoreEvent) error
    UpdateHandler func (*UpdateGraduation, *Graduation, eh.AggregateStoreEvent) error
}

func (o *GraduationCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (ret error) {
    switch cmd.CommandType() {
    case CreateGraduationCommand:
        ret = o.CreateHandler(cmd.(*CreateGraduation), entity.(*Graduation), store)
    case DeleteGraduationCommand:
        ret = o.DeleteHandler(cmd.(*DeleteGraduation), entity.(*Graduation), store)
    case UpdateGraduationCommand:
        ret = o.UpdateHandler(cmd.(*UpdateGraduation), entity.(*Graduation), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *GraduationCommandHandler) SetupCommandHandler() (ret error) {
    if o.CreateHandler == nil {
        o.CreateHandler = func(command *CreateGraduation, entity *Graduation, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityAlreadyExists(entity.Id, GraduationAggregateType)
            } else {
                store.StoreEvent(GraduationCreatedEvent, &GraduationCreated{
                    Id: command.Id,
                    Name: command.Name,
                    Level: command.Level,})
            }
            return
        }
    }
    
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteGraduation, entity *Graduation, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, GraduationAggregateType)
            } else if entity.Id != command.Id {
                ret = eh.IdsDismatch(entity.Id, command.Id, GraduationAggregateType)
            } else {
                store.StoreEvent(GraduationDeletedEvent, &GraduationDeleted{
                    Id: command.Id,})
            }
            return
        }
    }
    
    if o.UpdateHandler == nil {
        o.UpdateHandler = func(command *UpdateGraduation, entity *Graduation, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, GraduationAggregateType)
            } else if entity.Id != command.Id {
                ret = eh.IdsDismatch(entity.Id, command.Id, GraduationAggregateType)
            } else {
                store.StoreEvent(GraduationUpdatedEvent, &GraduationUpdated{
                    Id: command.Id,
                    Name: command.Name,
                    Level: command.Level,})
            }
            return
        }
    }
    
    return
}


type GraduationEventHandler struct {
    CreatedHandler func (*GraduationCreated, *Graduation) error
    DeletedHandler func (*GraduationDeleted, *Graduation) error
    UpdatedHandler func (*GraduationUpdated, *Graduation) error
}

func (o *GraduationEventHandler) Apply(event eventhorizon.Event, entity interface{}) (ret error) {
    switch event.EventType() {
    case GraduationCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*GraduationCreated), entity.(*Graduation))
    case GraduationDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*GraduationDeleted), entity.(*Graduation))
    case GraduationUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*GraduationUpdated), entity.(*Graduation))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *GraduationEventHandler) SetupEventHandler() (ret error) {
    if o.CreatedHandler == nil {
        o.CreatedHandler = func(event *GraduationCreated, entity *Graduation) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityAlreadyExists(entity.Id, GraduationAggregateType)
            } else {
                entity.Id = event.Id
                entity.Name = event.Name
                entity.Level = event.Level
            }
            return
        }
    }
    
    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *GraduationDeleted, entity *Graduation) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityNotExists(entity.Id, GraduationAggregateType)
            } else if entity.Id != event.Id {
                ret = eh.IdsDismatch(entity.Id, event.Id, GraduationAggregateType)
            } else {
                entity.Id = ""
            }
            return
        }
    }
    
    if o.UpdatedHandler == nil {
        o.UpdatedHandler = func(event *GraduationUpdated, entity *Graduation) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityNotExists(entity.Id, GraduationAggregateType)
            } else if entity.Id != event.Id {
                ret = eh.IdsDismatch(entity.Id, event.Id, GraduationAggregateType)
            } else {
                entity.Id = event.Id
                entity.Name = event.Name
                entity.Level = event.Level
            }
            return
        }
    }
    
    return
}


const GraduationAggregateType eventhorizon.AggregateType = "Graduation"

type GraduationAggregateInitializer struct {
    *eh.AggregateInitializer
    *GraduationCommandHandler
    *GraduationEventHandler
}

func NewGraduationAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus) (ret *GraduationAggregateInitializer) {
    
    commandHandler := &GraduationCommandHandler{}
    eventHandler := &GraduationEventHandler{}
	ret = &GraduationAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(GraduationAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(GraduationAggregateType, id, commandHandler, eventHandler, NewGraduation())
        },
        GraduationCommandTypes().Literals(), GraduationEventTypes().Literals(),
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus),
        GraduationCommandHandler: commandHandler,
        GraduationEventHandler: eventHandler,
    }

    return
}


func (o *GraduationAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GraduationEventTypes().GraduationCreated())
}

func (o *GraduationAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GraduationEventTypes().GraduationDeleted())
}

func (o *GraduationAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, GraduationEventTypes().GraduationUpdated())
}



type ProfileCommandHandler struct {
    CreateHandler func (*CreateProfile, *Profile, eh.AggregateStoreEvent) error
    DeleteHandler func (*DeleteProfile, *Profile, eh.AggregateStoreEvent) error
    UpdateHandler func (*UpdateProfile, *Profile, eh.AggregateStoreEvent) error
}

func (o *ProfileCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (ret error) {
    switch cmd.CommandType() {
    case CreateProfileCommand:
        ret = o.CreateHandler(cmd.(*CreateProfile), entity.(*Profile), store)
    case DeleteProfileCommand:
        ret = o.DeleteHandler(cmd.(*DeleteProfile), entity.(*Profile), store)
    case UpdateProfileCommand:
        ret = o.UpdateHandler(cmd.(*UpdateProfile), entity.(*Profile), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *ProfileCommandHandler) SetupCommandHandler() (ret error) {
    if o.CreateHandler == nil {
        o.CreateHandler = func(command *CreateProfile, entity *Profile, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityAlreadyExists(entity.Id, ProfileAggregateType)
            } else {
                store.StoreEvent(ProfileCreatedEvent, &ProfileCreated{
                    Id: command.Id,
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
                    Education: command.Education,})
            }
            return
        }
    }
    
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteProfile, entity *Profile, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, ProfileAggregateType)
            } else if entity.Id != command.Id {
                ret = eh.IdsDismatch(entity.Id, command.Id, ProfileAggregateType)
            } else {
                store.StoreEvent(ProfileDeletedEvent, &ProfileDeleted{
                    Id: command.Id,})
            }
            return
        }
    }
    
    if o.UpdateHandler == nil {
        o.UpdateHandler = func(command *UpdateProfile, entity *Profile, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, ProfileAggregateType)
            } else if entity.Id != command.Id {
                ret = eh.IdsDismatch(entity.Id, command.Id, ProfileAggregateType)
            } else {
                store.StoreEvent(ProfileUpdatedEvent, &ProfileUpdated{
                    Id: command.Id,
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
                    Education: command.Education,})
            }
            return
        }
    }
    
    return
}


type ProfileEventHandler struct {
    CreatedHandler func (*ProfileCreated, *Profile) error
    DeletedHandler func (*ProfileDeleted, *Profile) error
    UpdatedHandler func (*ProfileUpdated, *Profile) error
}

func (o *ProfileEventHandler) Apply(event eventhorizon.Event, entity interface{}) (ret error) {
    switch event.EventType() {
    case ProfileCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*ProfileCreated), entity.(*Profile))
    case ProfileDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*ProfileDeleted), entity.(*Profile))
    case ProfileUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*ProfileUpdated), entity.(*Profile))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *ProfileEventHandler) SetupEventHandler() (ret error) {
    if o.CreatedHandler == nil {
        o.CreatedHandler = func(event *ProfileCreated, entity *Profile) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityAlreadyExists(entity.Id, ProfileAggregateType)
            } else {
                entity.Id = event.Id
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
    }
    
    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *ProfileDeleted, entity *Profile) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityNotExists(entity.Id, ProfileAggregateType)
            } else if entity.Id != event.Id {
                ret = eh.IdsDismatch(entity.Id, event.Id, ProfileAggregateType)
            } else {
                entity.Id = ""
            }
            return
        }
    }
    
    if o.UpdatedHandler == nil {
        o.UpdatedHandler = func(event *ProfileUpdated, entity *Profile) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityNotExists(entity.Id, ProfileAggregateType)
            } else if entity.Id != event.Id {
                ret = eh.IdsDismatch(entity.Id, event.Id, ProfileAggregateType)
            } else {
                entity.Id = event.Id
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
    }
    
    return
}


const ProfileAggregateType eventhorizon.AggregateType = "Profile"

type ProfileAggregateInitializer struct {
    *eh.AggregateInitializer
    *ProfileCommandHandler
    *ProfileEventHandler
}

func NewProfileAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus) (ret *ProfileAggregateInitializer) {
    
    commandHandler := &ProfileCommandHandler{}
    eventHandler := &ProfileEventHandler{}
	ret = &ProfileAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ProfileAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(ProfileAggregateType, id, commandHandler, eventHandler, NewProfile())
        },
        ProfileCommandTypes().Literals(), ProfileEventTypes().Literals(),
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus),
        ProfileCommandHandler: commandHandler,
        ProfileEventHandler: eventHandler,
    }

    return
}


func (o *ProfileAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ProfileEventTypes().ProfileCreated())
}

func (o *ProfileAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ProfileEventTypes().ProfileDeleted())
}

func (o *ProfileAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ProfileEventTypes().ProfileUpdated())
}



type PersonEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore
    eventBus eventhorizon.EventBus
    eventPublisher eventhorizon.EventPublisher
    commandBus eventhorizon.CommandBus
    ChurchAggregateInitializer *ChurchAggregateInitializer
    GraduationAggregateInitializer *GraduationAggregateInitializer
    ProfileAggregateInitializer *ProfileAggregateInitializer
}

func NewPersonEventhorizonInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus) (ret *PersonEventhorizonInitializer) {
    ret = &PersonEventhorizonInitializer{
        eventStore: eventStore,
        eventBus: eventBus,
        eventPublisher: eventPublisher,
        commandBus: commandBus,
        ChurchAggregateInitializer: NewChurchAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
        GraduationAggregateInitializer: NewGraduationAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
        ProfileAggregateInitializer: NewProfileAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    }
    return
}

func (o *PersonEventhorizonInitializer) Setup() (ret error) {
    
    if ret = o.ChurchAggregateInitializer.Setup(); ret != nil {
        return
    }
    
    if ret = o.GraduationAggregateInitializer.Setup(); ret != nil {
        return
    }
    
    if ret = o.ProfileAggregateInitializer.Setup(); ret != nil {
        return
    }

    return
}









