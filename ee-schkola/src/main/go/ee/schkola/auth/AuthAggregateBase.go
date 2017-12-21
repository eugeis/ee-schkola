package auth

import (
    "errors"
    "fmt"
    "github.com/eugeis/gee/eh"
    "github.com/looplab/eventhorizon"
    "github.com/looplab/eventhorizon/commandhandler/bus"
    "time"
)
type CommandHandler struct {
    SendEnabledConfirmationHandler func (*SendEnabledConfirmation, *Account, eh.AggregateStoreEvent) (err error)  `json:"sendEnabledConfirmationHandler" eh:"optional"`
    SendDisabledConfirmationHandler func (*SendDisabledConfirmation, *Account, eh.AggregateStoreEvent) (err error)  `json:"sendDisabledConfirmationHandler" eh:"optional"`
    LoginHandler func (*Login, *Account, eh.AggregateStoreEvent) (err error)  `json:"loginHandler" eh:"optional"`
    SendCreatedConfirmationHandler func (*SendCreatedConfirmation, *Account, eh.AggregateStoreEvent) (err error)  `json:"sendCreatedConfirmationHandler" eh:"optional"`
    CreateHandler func (*Create, *Account, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*Delete, *Account, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    DisableHandler func (*Disable, *Account, eh.AggregateStoreEvent) (err error)  `json:"disableHandler" eh:"optional"`
    EnableHandler func (*Enable, *Account, eh.AggregateStoreEvent) (err error)  `json:"enableHandler" eh:"optional"`
    UpdateHandler func (*Update, *Account, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *CommandHandler) AddSendEnabledConfirmationPreparer(preparer func (*SendEnabledConfirmation, *Account) (err error) ) {
    prevHandler := o.SendEnabledConfirmationHandler
	o.SendEnabledConfirmationHandler = func(command *SendEnabledConfirmation, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddSendDisabledConfirmationPreparer(preparer func (*SendDisabledConfirmation, *Account) (err error) ) {
    prevHandler := o.SendDisabledConfirmationHandler
	o.SendDisabledConfirmationHandler = func(command *SendDisabledConfirmation, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddLoginPreparer(preparer func (*Login, *Account) (err error) ) {
    prevHandler := o.LoginHandler
	o.LoginHandler = func(command *Login, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddSendCreatedConfirmationPreparer(preparer func (*SendCreatedConfirmation, *Account) (err error) ) {
    prevHandler := o.SendCreatedConfirmationHandler
	o.SendCreatedConfirmationHandler = func(command *SendCreatedConfirmation, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddCreatePreparer(preparer func (*Create, *Account) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *Create, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddDeletePreparer(preparer func (*Delete, *Account) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *Delete, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddDisablePreparer(preparer func (*Disable, *Account) (err error) ) {
    prevHandler := o.DisableHandler
	o.DisableHandler = func(command *Disable, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddEnablePreparer(preparer func (*Enable, *Account) (err error) ) {
    prevHandler := o.EnableHandler
	o.EnableHandler = func(command *Enable, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddUpdatePreparer(preparer func (*Update, *Account) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *Update, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case SendEnabledConfirmationCommand:
        err = o.SendEnabledConfirmationHandler(cmd.(*SendEnabledConfirmation), entity.(*Account), store)
    case SendDisabledConfirmationCommand:
        err = o.SendDisabledConfirmationHandler(cmd.(*SendDisabledConfirmation), entity.(*Account), store)
    case LoginCommand:
        err = o.LoginHandler(cmd.(*Login), entity.(*Account), store)
    case SendCreatedConfirmationCommand:
        err = o.SendCreatedConfirmationHandler(cmd.(*SendCreatedConfirmation), entity.(*Account), store)
    case CreateCommand:
        err = o.CreateHandler(cmd.(*Create), entity.(*Account), store)
    case DeleteCommand:
        err = o.DeleteHandler(cmd.(*Delete), entity.(*Account), store)
    case DisableCommand:
        err = o.DisableHandler(cmd.(*Disable), entity.(*Account), store)
    case EnableCommand:
        err = o.EnableHandler(cmd.(*Enable), entity.(*Account), store)
    case UpdateCommand:
        err = o.UpdateHandler(cmd.(*Update), entity.(*Account), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *CommandHandler) SetupCommandHandler() (err error) {
    o.SendEnabledConfirmationHandler = func(command *SendEnabledConfirmation, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(SendEnabledConfirmationedEvent, &SendEnabledConfirmationed{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.SendDisabledConfirmationHandler = func(command *SendDisabledConfirmation, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(SendDisabledConfirmationedEvent, &SendDisabledConfirmationed{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.LoginHandler = func(command *Login, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(LoggedEvent, &Logged{
                Username: command.Username,
                Email: command.Email,
                Password: command.Password,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.SendCreatedConfirmationHandler = func(command *SendCreatedConfirmation, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(SendCreatedConfirmationedEvent, &SendCreatedConfirmationed{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.CreateHandler = func(command *Create, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(createdEvent, &Created{
                Name: command.Name,
                Username: command.Username,
                Password: command.Password,
                Email: command.Email,
                Disabled: command.Disabled,
                Roles: command.Roles,
                Profile: command.Profile,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *Delete, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(deletedEvent, &Deleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DisableHandler = func(command *Disable, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(DisabledEvent, &Disabled{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.EnableHandler = func(command *Enable, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(EnabledEvent, &Enabled{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *Update, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(updatedEvent, &Updated{
                Name: command.Name,
                Username: command.Username,
                Password: command.Password,
                Email: command.Email,
                Disabled: command.Disabled,
                Roles: command.Roles,
                Profile: command.Profile,
                Id: command.Id,}, time.Now())
        }
        return
    }
    return
}


type EventHandler struct {
    SendEnabledConfirmationHandler func (*SendEnabledConfirmation, *Account) (err error)  `json:"sendEnabledConfirmationHandler" eh:"optional"`
    SendDisabledConfirmationHandler func (*SendDisabledConfirmation, *Account) (err error)  `json:"sendDisabledConfirmationHandler" eh:"optional"`
    LoginHandler func (*Login, *Account) (err error)  `json:"loginHandler" eh:"optional"`
    SendCreatedConfirmationHandler func (*SendCreatedConfirmation, *Account) (err error)  `json:"sendCreatedConfirmationHandler" eh:"optional"`
    CreateHandler func (*Create, *Account) (err error)  `json:"createHandler" eh:"optional"`
    CreatedHandler func (*Created, *Account) (err error)  `json:"createdHandler" eh:"optional"`
    DeleteHandler func (*Delete, *Account) (err error)  `json:"deleteHandler" eh:"optional"`
    DeletedHandler func (*Deleted, *Account) (err error)  `json:"deletedHandler" eh:"optional"`
    LoggedHandler func (*Logged, *Account) (err error)  `json:"loggedHandler" eh:"optional"`
    SendCreatedConfirmationedHandler func (*SendCreatedConfirmationed, *Account) (err error)  `json:"sendCreatedConfirmationedHandler" eh:"optional"`
    SendEnabledConfirmationedHandler func (*SendEnabledConfirmationed, *Account) (err error)  `json:"sendEnabledConfirmationedHandler" eh:"optional"`
    SendDisabledConfirmationedHandler func (*SendDisabledConfirmationed, *Account) (err error)  `json:"sendDisabledConfirmationedHandler" eh:"optional"`
    DisableHandler func (*Disable, *Account) (err error)  `json:"disableHandler" eh:"optional"`
    EnableHandler func (*Enable, *Account) (err error)  `json:"enableHandler" eh:"optional"`
    UpdateHandler func (*Update, *Account) (err error)  `json:"updateHandler" eh:"optional"`
    UpdatedHandler func (*Updated, *Account) (err error)  `json:"updatedHandler" eh:"optional"`
    EnabledHandler func (*Enabled, *Account) (err error)  `json:"enabledHandler" eh:"optional"`
    DisabledHandler func (*Disabled, *Account) (err error)  `json:"disabledHandler" eh:"optional"`
}

func (o *EventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case SendEnabledConfirmationEvent:
        err = o.SendEnabledConfirmationHandler(event.Data().(*SendEnabledConfirmation), entity.(*Account))
    case SendDisabledConfirmationEvent:
        err = o.SendDisabledConfirmationHandler(event.Data().(*SendDisabledConfirmation), entity.(*Account))
    case LoginEvent:
        err = o.LoginHandler(event.Data().(*Login), entity.(*Account))
    case SendCreatedConfirmationEvent:
        err = o.SendCreatedConfirmationHandler(event.Data().(*SendCreatedConfirmation), entity.(*Account))
    case CreateEvent:
        err = o.CreateHandler(event.Data().(*Create), entity.(*Account))
    case CreatedEvent:
        err = o.CreatedHandler(event.Data().(*Created), entity.(*Account))
    case DeleteEvent:
        err = o.DeleteHandler(event.Data().(*Delete), entity.(*Account))
    case DeletedEvent:
        err = o.DeletedHandler(event.Data().(*Deleted), entity.(*Account))
    case LoggedEvent:
        err = o.LoggedHandler(event.Data().(*Logged), entity.(*Account))
    case SendCreatedConfirmationedEvent:
        err = o.SendCreatedConfirmationedHandler(event.Data().(*SendCreatedConfirmationed), entity.(*Account))
    case SendEnabledConfirmationedEvent:
        err = o.SendEnabledConfirmationedHandler(event.Data().(*SendEnabledConfirmationed), entity.(*Account))
    case SendDisabledConfirmationedEvent:
        err = o.SendDisabledConfirmationedHandler(event.Data().(*SendDisabledConfirmationed), entity.(*Account))
    case DisableEvent:
        err = o.DisableHandler(event.Data().(*Disable), entity.(*Account))
    case EnableEvent:
        err = o.EnableHandler(event.Data().(*Enable), entity.(*Account))
    case UpdateEvent:
        err = o.UpdateHandler(event.Data().(*Update), entity.(*Account))
    case UpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*Updated), entity.(*Account))
    case EnabledEvent:
        err = o.EnabledHandler(event.Data().(*Enabled), entity.(*Account))
    case DisabledEvent:
        err = o.DisabledHandler(event.Data().(*Disabled), entity.(*Account))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *EventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(SendEnabledConfirmationEvent, func() eventhorizon.EventData {
		return &SendEnabledConfirmation{}
	})

    //default handler implementation
    o.SendEnabledConfirmationHandler = func(event *SendEnabledConfirmation, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(SendEnabledConfirmationEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(SendDisabledConfirmationEvent, func() eventhorizon.EventData {
		return &SendDisabledConfirmation{}
	})

    //default handler implementation
    o.SendDisabledConfirmationHandler = func(event *SendDisabledConfirmation, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(SendDisabledConfirmationEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(LoginEvent, func() eventhorizon.EventData {
		return &Login{}
	})

    //default handler implementation
    o.LoginHandler = func(event *Login, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(LoginEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(SendCreatedConfirmationEvent, func() eventhorizon.EventData {
		return &SendCreatedConfirmation{}
	})

    //default handler implementation
    o.SendCreatedConfirmationHandler = func(event *SendCreatedConfirmation, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(SendCreatedConfirmationEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(CreateEvent, func() eventhorizon.EventData {
		return &Create{}
	})

    //default handler implementation
    o.CreateHandler = func(event *Create, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(CreateEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(CreatedEvent, func() eventhorizon.EventData {
		return &Created{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *Created, entity *Account) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, AccountAggregateType); err == nil {
            entity.Name = event.Name
            entity.Username = event.Username
            entity.Password = event.Password
            entity.Email = event.Email
            entity.Disabled = event.Disabled
            entity.Roles = event.Roles
            entity.Profile = event.Profile
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DeleteEvent, func() eventhorizon.EventData {
		return &Delete{}
	})

    //default handler implementation
    o.DeleteHandler = func(event *Delete, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(DeleteEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DeletedEvent, func() eventhorizon.EventData {
		return &Deleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *Deleted, entity *Account) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AccountAggregateType); err == nil {
            *entity = *NewAccount()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(LoggedEvent, func() eventhorizon.EventData {
		return &Logged{}
	})

    //default handler implementation
    o.LoggedHandler = func(event *Logged, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(LoggedEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(SendCreatedConfirmationedEvent, func() eventhorizon.EventData {
		return &SendCreatedConfirmationed{}
	})

    //default handler implementation
    o.SendCreatedConfirmationedHandler = func(event *SendCreatedConfirmationed, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(SendCreatedConfirmationedEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(SendEnabledConfirmationedEvent, func() eventhorizon.EventData {
		return &SendEnabledConfirmationed{}
	})

    //default handler implementation
    o.SendEnabledConfirmationedHandler = func(event *SendEnabledConfirmationed, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(SendEnabledConfirmationedEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(SendDisabledConfirmationedEvent, func() eventhorizon.EventData {
		return &SendDisabledConfirmationed{}
	})

    //default handler implementation
    o.SendDisabledConfirmationedHandler = func(event *SendDisabledConfirmationed, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(SendDisabledConfirmationedEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DisableEvent, func() eventhorizon.EventData {
		return &Disable{}
	})

    //default handler implementation
    o.DisableHandler = func(event *Disable, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(DisableEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(EnableEvent, func() eventhorizon.EventData {
		return &Enable{}
	})

    //default handler implementation
    o.EnableHandler = func(event *Enable, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(EnableEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdateEvent, func() eventhorizon.EventData {
		return &Update{}
	})

    //default handler implementation
    o.UpdateHandler = func(event *Update, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(UpdateEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdatedEvent, func() eventhorizon.EventData {
		return &Updated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *Updated, entity *Account) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AccountAggregateType); err == nil {
            entity.Name = event.Name
            entity.Username = event.Username
            entity.Password = event.Password
            entity.Email = event.Email
            entity.Disabled = event.Disabled
            entity.Roles = event.Roles
            entity.Profile = event.Profile
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(EnabledEvent, func() eventhorizon.EventData {
		return &Enabled{}
	})

    //default handler implementation
    o.EnabledHandler = func(event *Enabled, entity *Account) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AccountAggregateType); err == nil {
            entity.Disabled = false
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DisabledEvent, func() eventhorizon.EventData {
		return &Disabled{}
	})

    //default handler implementation
    o.DisabledHandler = func(event *Disabled, entity *Account) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AccountAggregateType); err == nil {
            entity.Disabled = true
        }
        return
    }
    return
}


const AccountAggregateType eventhorizon.AggregateType = "Account"

type AggregateInitializer struct {
    *eh.AggregateInitializer
    *CommandHandler
    *EventHandler
    ProjectorHandler *EventHandler `json:"projectorHandler" eh:"optional"`
}


func (o *AggregateInitializer) RegisterForLogged(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountEventTypes().Logged())
}

func (o *AggregateInitializer) RegisterForSendCreatedConfirmationed(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountEventTypes().SendCreatedConfirmationed())
}

func (o *AggregateInitializer) RegisterForSendEnabledConfirmationed(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountEventTypes().SendEnabledConfirmationed())
}

func (o *AggregateInitializer) RegisterForSendDisabledConfirmationed(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountEventTypes().SendDisabledConfirmationed())
}


func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *AggregateInitializer) {
    
    commandHandler := &AccountCommandHandler{}
    eventHandler := &AccountEventHandler{}
    entityFactory := func() eventhorizon.Entity { return NewAccount() }
    ret = &AggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(AccountAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(AccountAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        AccountCommandTypes().Literals(), AccountEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), AccountCommandHandler: commandHandler, AccountEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type AuthEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore `json:"eventStore" eh:"optional"`
    eventBus eventhorizon.EventBus `json:"eventBus" eh:"optional"`
    eventPublisher eventhorizon.EventPublisher `json:"eventPublisher" eh:"optional"`
    commandBus *bus.CommandHandler `json:"commandBus" eh:"optional"`
    AccountAggregateInitializer *AggregateInitializer `json:"accountAggregateInitializer" eh:"optional"`
}

func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *AuthEventhorizonInitializer) {
    accountAggregateInitializer := New@@EMPTY@@(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    ret = &AuthEventhorizonInitializer{
        eventStore: eventStore,
        eventBus: eventBus,
        eventPublisher: eventPublisher,
        commandBus: commandBus,
        AccountAggregateInitializer: accountAggregateInitializer,
    }
    return
}

func (o *AuthEventhorizonInitializer) Setup() (err error) {
    
    if err = o.AccountAggregateInitializer.Setup(); err != nil {
        return
    }

    return
}









