package auth

import (
    "errors"
    "fmt"
    "github.com/eugeis/gee/eh"
    "github.com/looplab/eventhorizon"
    "github.com/looplab/eventhorizon/commandhandler/bus"
    "time"
)
type AccountCommandHandler struct {
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

func (o *AccountCommandHandler) AddSendEnabledConfirmationPreparer(preparer func (*SendEnabledConfirmation, *Account) (err error) ) {
    prevHandler := o.SendEnabledConfirmationHandler
	o.SendEnabledConfirmationHandler = func(command *SendEnabledConfirmation, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) AddSendDisabledConfirmationPreparer(preparer func (*SendDisabledConfirmation, *Account) (err error) ) {
    prevHandler := o.SendDisabledConfirmationHandler
	o.SendDisabledConfirmationHandler = func(command *SendDisabledConfirmation, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) AddLoginPreparer(preparer func (*Login, *Account) (err error) ) {
    prevHandler := o.LoginHandler
	o.LoginHandler = func(command *Login, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) AddSendCreatedConfirmationPreparer(preparer func (*SendCreatedConfirmation, *Account) (err error) ) {
    prevHandler := o.SendCreatedConfirmationHandler
	o.SendCreatedConfirmationHandler = func(command *SendCreatedConfirmation, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) AddCreatePreparer(preparer func (*Create, *Account) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *Create, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) AddDeletePreparer(preparer func (*Delete, *Account) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *Delete, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) AddDisablePreparer(preparer func (*Disable, *Account) (err error) ) {
    prevHandler := o.DisableHandler
	o.DisableHandler = func(command *Disable, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) AddEnablePreparer(preparer func (*Enable, *Account) (err error) ) {
    prevHandler := o.EnableHandler
	o.EnableHandler = func(command *Enable, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) AddUpdatePreparer(preparer func (*Update, *Account) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *Update, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case SendAccountEnabledConfirmationCommand:
        err = o.SendEnabledConfirmationHandler(cmd.(*SendEnabledConfirmation), entity.(*Account), store)
    case SendAccountDisabledConfirmationCommand:
        err = o.SendDisabledConfirmationHandler(cmd.(*SendDisabledConfirmation), entity.(*Account), store)
    case LoginAccountCommand:
        err = o.LoginHandler(cmd.(*Login), entity.(*Account), store)
    case SendAccountCreatedConfirmationCommand:
        err = o.SendCreatedConfirmationHandler(cmd.(*SendCreatedConfirmation), entity.(*Account), store)
    case CreateAccountCommand:
        err = o.CreateHandler(cmd.(*Create), entity.(*Account), store)
    case DeleteAccountCommand:
        err = o.DeleteHandler(cmd.(*Delete), entity.(*Account), store)
    case DisableAccountCommand:
        err = o.DisableHandler(cmd.(*Disable), entity.(*Account), store)
    case EnableAccountCommand:
        err = o.EnableHandler(cmd.(*Enable), entity.(*Account), store)
    case UpdateAccountCommand:
        err = o.UpdateHandler(cmd.(*Update), entity.(*Account), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *AccountCommandHandler) SetupCommandHandler() (err error) {
    o.SendEnabledConfirmationHandler = func(command *SendEnabledConfirmation, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(SendEnabledAccountConfirmationedEvent, &SendEnabledConfirmationed{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.SendDisabledConfirmationHandler = func(command *SendDisabledConfirmation, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(SendDisabledAccountConfirmationedEvent, &SendDisabledConfirmationed{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.LoginHandler = func(command *Login, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(AccountLoggedEvent, &Logged{
                Username: command.Username,
                Email: command.Email,
                Password: command.Password,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.SendCreatedConfirmationHandler = func(command *SendCreatedConfirmation, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(SendCreatedAccountConfirmationedEvent, &SendCreatedConfirmationed{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.CreateHandler = func(command *Create, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(AccountCreatedEvent, &Created{
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
            store.StoreEvent(AccountDeletedEvent, &Deleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DisableHandler = func(command *Disable, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(AccountDisabledEvent, &Disabled{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.EnableHandler = func(command *Enable, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(AccountEnabledEvent, &Enabled{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *Update, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(AccountUpdatedEvent, &Updated{
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


type AccountEventHandler struct {
    CreatedHandler func (*Created, *Account) (err error)  `json:"createdHandler" eh:"optional"`
    DeletedHandler func (*Deleted, *Account) (err error)  `json:"deletedHandler" eh:"optional"`
    LoggedHandler func (*Logged, *Account) (err error)  `json:"loggedHandler" eh:"optional"`
    SendCreatedConfirmationedHandler func (*SendCreatedConfirmationed, *Account) (err error)  `json:"sendCreatedConfirmationedHandler" eh:"optional"`
    SendEnabledConfirmationedHandler func (*SendEnabledConfirmationed, *Account) (err error)  `json:"sendEnabledConfirmationedHandler" eh:"optional"`
    SendDisabledConfirmationedHandler func (*SendDisabledConfirmationed, *Account) (err error)  `json:"sendDisabledConfirmationedHandler" eh:"optional"`
    UpdatedHandler func (*Updated, *Account) (err error)  `json:"updatedHandler" eh:"optional"`
    EnabledHandler func (*Enabled, *Account) (err error)  `json:"enabledHandler" eh:"optional"`
    DisabledHandler func (*Disabled, *Account) (err error)  `json:"disabledHandler" eh:"optional"`
}

func (o *AccountEventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case AccountCreatedEvent:
        err = o.CreatedHandler(event.Data().(*Created), entity.(*Account))
    case AccountDeletedEvent:
        err = o.DeletedHandler(event.Data().(*Deleted), entity.(*Account))
    case AccountLoggedEvent:
        err = o.LoggedHandler(event.Data().(*Logged), entity.(*Account))
    case SendCreatedAccountConfirmationedEvent:
        err = o.SendCreatedConfirmationedHandler(event.Data().(*SendCreatedConfirmationed), entity.(*Account))
    case SendEnabledAccountConfirmationedEvent:
        err = o.SendEnabledConfirmationedHandler(event.Data().(*SendEnabledConfirmationed), entity.(*Account))
    case SendDisabledAccountConfirmationedEvent:
        err = o.SendDisabledConfirmationedHandler(event.Data().(*SendDisabledConfirmationed), entity.(*Account))
    case AccountUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*Updated), entity.(*Account))
    case AccountEnabledEvent:
        err = o.EnabledHandler(event.Data().(*Enabled), entity.(*Account))
    case AccountDisabledEvent:
        err = o.DisabledHandler(event.Data().(*Disabled), entity.(*Account))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *AccountEventHandler) SetupEventHandler() (err error) {

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
    eventhorizon.RegisterEventData(DeletedEvent, func() eventhorizon.EventData {
		return &Deleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *Deleted, entity *Account) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AccountAggregateType); err == nil {
            *entity = *New@@EMPTY@@()
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

type AccountAggregateInitializer struct {
    *eh.AggregateInitializer
    *AccountCommandHandler
    *AccountEventHandler
    ProjectorHandler *AccountEventHandler `json:"projectorHandler" eh:"optional"`
}


func (o *AccountAggregateInitializer) RegisterForLogged(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountEventTypes().AccountLogged())
}

func (o *AccountAggregateInitializer) RegisterForSendCreatedConfirmationed(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountEventTypes().SendCreatedAccountConfirmationed())
}

func (o *AccountAggregateInitializer) RegisterForSendEnabledConfirmationed(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountEventTypes().SendEnabledAccountConfirmationed())
}

func (o *AccountAggregateInitializer) RegisterForSendDisabledConfirmationed(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountEventTypes().SendDisabledAccountConfirmationed())
}


func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *AccountAggregateInitializer) {
    
    commandHandler := &AccountCommandHandler{}
    eventHandler := &AccountEventHandler{}
    entityFactory := func() eventhorizon.Entity { return New@@EMPTY@@() }
    ret = &AccountAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(AccountAggregateType,
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
    AccountAggregateInitializer *AccountAggregateInitializer `json:"accountAggregateInitializer" eh:"optional"`
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









