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
    SendEnabledConfirmationHandler func (*SendAccountEnabledConfirmation, *Account, eh.AggregateStoreEvent) (err error)  `json:"sendEnabledConfirmationHandler" eh:"optional"`
    SendDisabledConfirmationHandler func (*SendAccountDisabledConfirmation, *Account, eh.AggregateStoreEvent) (err error)  `json:"sendDisabledConfirmationHandler" eh:"optional"`
    LoginHandler func (*LoginAccount, *Account, eh.AggregateStoreEvent) (err error)  `json:"loginHandler" eh:"optional"`
    SendCreatedConfirmationHandler func (*SendAccountCreatedConfirmation, *Account, eh.AggregateStoreEvent) (err error)  `json:"sendCreatedConfirmationHandler" eh:"optional"`
    CreateHandler func (*CreateAccount, *Account, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*DeleteAccount, *Account, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    DisableHandler func (*DisableAccount, *Account, eh.AggregateStoreEvent) (err error)  `json:"disableHandler" eh:"optional"`
    EnableHandler func (*EnableAccount, *Account, eh.AggregateStoreEvent) (err error)  `json:"enableHandler" eh:"optional"`
    UpdateHandler func (*UpdateAccount, *Account, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *CommandHandler) AddSendEnabledConfirmationPreparer(preparer func (*SendAccountEnabledConfirmation, *Account) (err error) ) {
    prevHandler := o.SendEnabledConfirmationHandler
	o.SendEnabledConfirmationHandler = func(command *SendAccountEnabledConfirmation, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddSendDisabledConfirmationPreparer(preparer func (*SendAccountDisabledConfirmation, *Account) (err error) ) {
    prevHandler := o.SendDisabledConfirmationHandler
	o.SendDisabledConfirmationHandler = func(command *SendAccountDisabledConfirmation, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddLoginPreparer(preparer func (*LoginAccount, *Account) (err error) ) {
    prevHandler := o.LoginHandler
	o.LoginHandler = func(command *LoginAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddSendCreatedConfirmationPreparer(preparer func (*SendAccountCreatedConfirmation, *Account) (err error) ) {
    prevHandler := o.SendCreatedConfirmationHandler
	o.SendCreatedConfirmationHandler = func(command *SendAccountCreatedConfirmation, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddCreatePreparer(preparer func (*CreateAccount, *Account) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *CreateAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddDeletePreparer(preparer func (*DeleteAccount, *Account) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *DeleteAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddDisablePreparer(preparer func (*DisableAccount, *Account) (err error) ) {
    prevHandler := o.DisableHandler
	o.DisableHandler = func(command *DisableAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddEnablePreparer(preparer func (*EnableAccount, *Account) (err error) ) {
    prevHandler := o.EnableHandler
	o.EnableHandler = func(command *EnableAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddUpdatePreparer(preparer func (*UpdateAccount, *Account) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *UpdateAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case SendAccountEnabledConfirmationCommand:
        err = o.SendEnabledConfirmationHandler(cmd.(*SendAccountEnabledConfirmation), entity.(*Account), store)
    case SendAccountDisabledConfirmationCommand:
        err = o.SendDisabledConfirmationHandler(cmd.(*SendAccountDisabledConfirmation), entity.(*Account), store)
    case LoginAccountCommand:
        err = o.LoginHandler(cmd.(*LoginAccount), entity.(*Account), store)
    case SendAccountCreatedConfirmationCommand:
        err = o.SendCreatedConfirmationHandler(cmd.(*SendAccountCreatedConfirmation), entity.(*Account), store)
    case CreateAccountCommand:
        err = o.CreateHandler(cmd.(*CreateAccount), entity.(*Account), store)
    case DeleteAccountCommand:
        err = o.DeleteHandler(cmd.(*DeleteAccount), entity.(*Account), store)
    case DisableAccountCommand:
        err = o.DisableHandler(cmd.(*DisableAccount), entity.(*Account), store)
    case EnableAccountCommand:
        err = o.EnableHandler(cmd.(*EnableAccount), entity.(*Account), store)
    case UpdateAccountCommand:
        err = o.UpdateHandler(cmd.(*UpdateAccount), entity.(*Account), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *CommandHandler) SetupCommandHandler() (err error) {
    o.SendEnabledConfirmationHandler = func(command *SendAccountEnabledConfirmation, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(SendEnabledAccountConfirmationedEvent, &SendEnabledAccountConfirmationed{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.SendDisabledConfirmationHandler = func(command *SendAccountDisabledConfirmation, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(SendDisabledAccountConfirmationedEvent, &SendDisabledAccountConfirmationed{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.LoginHandler = func(command *LoginAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(AccountLoggedEvent, &AccountLogged{
                Username: command.Username,
                Email: command.Email,
                Password: command.Password,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.SendCreatedConfirmationHandler = func(command *SendAccountCreatedConfirmation, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(SendCreatedAccountConfirmationedEvent, &SendCreatedAccountConfirmationed{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.CreateHandler = func(command *CreateAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(AccountCreatedEvent, &AccountCreated{
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
    o.DeleteHandler = func(command *DeleteAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(AccountDeletedEvent, &AccountDeleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DisableHandler = func(command *DisableAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(AccountDisabledEvent, &AccountDisabled{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.EnableHandler = func(command *EnableAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(AccountEnabledEvent, &AccountEnabled{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *UpdateAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(AccountUpdatedEvent, &AccountUpdated{
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
    SendEnabledConfirmationHandler func (*SendAccountEnabledConfirmation, *Account) (err error)  `json:"sendEnabledConfirmationHandler" eh:"optional"`
    SendDisabledConfirmationHandler func (*SendAccountDisabledConfirmation, *Account) (err error)  `json:"sendDisabledConfirmationHandler" eh:"optional"`
    LoginHandler func (*LoginAccount, *Account) (err error)  `json:"loginHandler" eh:"optional"`
    SendCreatedConfirmationHandler func (*SendAccountCreatedConfirmation, *Account) (err error)  `json:"sendCreatedConfirmationHandler" eh:"optional"`
    CreateHandler func (*CreateAccount, *Account) (err error)  `json:"createHandler" eh:"optional"`
    CreatedHandler func (*AccountCreated, *Account) (err error)  `json:"createdHandler" eh:"optional"`
    DeleteHandler func (*DeleteAccount, *Account) (err error)  `json:"deleteHandler" eh:"optional"`
    DeletedHandler func (*AccountDeleted, *Account) (err error)  `json:"deletedHandler" eh:"optional"`
    LoggedHandler func (*AccountLogged, *Account) (err error)  `json:"loggedHandler" eh:"optional"`
    SendCreatedConfirmationedHandler func (*SendCreatedAccountConfirmationed, *Account) (err error)  `json:"sendCreatedConfirmationedHandler" eh:"optional"`
    SendEnabledConfirmationedHandler func (*SendEnabledAccountConfirmationed, *Account) (err error)  `json:"sendEnabledConfirmationedHandler" eh:"optional"`
    SendDisabledConfirmationedHandler func (*SendDisabledAccountConfirmationed, *Account) (err error)  `json:"sendDisabledConfirmationedHandler" eh:"optional"`
    DisableHandler func (*DisableAccount, *Account) (err error)  `json:"disableHandler" eh:"optional"`
    EnableHandler func (*EnableAccount, *Account) (err error)  `json:"enableHandler" eh:"optional"`
    UpdateHandler func (*UpdateAccount, *Account) (err error)  `json:"updateHandler" eh:"optional"`
    UpdatedHandler func (*AccountUpdated, *Account) (err error)  `json:"updatedHandler" eh:"optional"`
    EnabledHandler func (*AccountEnabled, *Account) (err error)  `json:"enabledHandler" eh:"optional"`
    DisabledHandler func (*AccountDisabled, *Account) (err error)  `json:"disabledHandler" eh:"optional"`
}

func (o *EventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case SendEnabledAccountConfirmationEvent:
        err = o.SendEnabledConfirmationHandler(event.Data().(*SendAccountEnabledConfirmation), entity.(*Account))
    case SendDisabledAccountConfirmationEvent:
        err = o.SendDisabledConfirmationHandler(event.Data().(*SendAccountDisabledConfirmation), entity.(*Account))
    case AccountLoginEvent:
        err = o.LoginHandler(event.Data().(*LoginAccount), entity.(*Account))
    case SendCreatedAccountConfirmationEvent:
        err = o.SendCreatedConfirmationHandler(event.Data().(*SendAccountCreatedConfirmation), entity.(*Account))
    case AccountCreateEvent:
        err = o.CreateHandler(event.Data().(*CreateAccount), entity.(*Account))
    case AccountCreatedEvent:
        err = o.CreatedHandler(event.Data().(*AccountCreated), entity.(*Account))
    case AccountDeleteEvent:
        err = o.DeleteHandler(event.Data().(*DeleteAccount), entity.(*Account))
    case AccountDeletedEvent:
        err = o.DeletedHandler(event.Data().(*AccountDeleted), entity.(*Account))
    case AccountLoggedEvent:
        err = o.LoggedHandler(event.Data().(*AccountLogged), entity.(*Account))
    case SendCreatedAccountConfirmationedEvent:
        err = o.SendCreatedConfirmationedHandler(event.Data().(*SendCreatedAccountConfirmationed), entity.(*Account))
    case SendEnabledAccountConfirmationedEvent:
        err = o.SendEnabledConfirmationedHandler(event.Data().(*SendEnabledAccountConfirmationed), entity.(*Account))
    case SendDisabledAccountConfirmationedEvent:
        err = o.SendDisabledConfirmationedHandler(event.Data().(*SendDisabledAccountConfirmationed), entity.(*Account))
    case AccountDisableEvent:
        err = o.DisableHandler(event.Data().(*DisableAccount), entity.(*Account))
    case AccountEnableEvent:
        err = o.EnableHandler(event.Data().(*EnableAccount), entity.(*Account))
    case AccountUpdateEvent:
        err = o.UpdateHandler(event.Data().(*UpdateAccount), entity.(*Account))
    case AccountUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*AccountUpdated), entity.(*Account))
    case AccountEnabledEvent:
        err = o.EnabledHandler(event.Data().(*AccountEnabled), entity.(*Account))
    case AccountDisabledEvent:
        err = o.DisabledHandler(event.Data().(*AccountDisabled), entity.(*Account))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *EventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(SendAccountEnabledConfirmationEvent, func() eventhorizon.EventData {
		return &SendAccountEnabledConfirmation{}
	})

    //default handler implementation
    o.SendEnabledConfirmationHandler = func(event *SendAccountEnabledConfirmation, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(SendAccountEnabledConfirmationEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(SendAccountDisabledConfirmationEvent, func() eventhorizon.EventData {
		return &SendAccountDisabledConfirmation{}
	})

    //default handler implementation
    o.SendDisabledConfirmationHandler = func(event *SendAccountDisabledConfirmation, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(SendAccountDisabledConfirmationEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(LoginAccountEvent, func() eventhorizon.EventData {
		return &LoginAccount{}
	})

    //default handler implementation
    o.LoginHandler = func(event *LoginAccount, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(LoginAccountEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(SendAccountCreatedConfirmationEvent, func() eventhorizon.EventData {
		return &SendAccountCreatedConfirmation{}
	})

    //default handler implementation
    o.SendCreatedConfirmationHandler = func(event *SendAccountCreatedConfirmation, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(SendAccountCreatedConfirmationEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(CreateAccountEvent, func() eventhorizon.EventData {
		return &CreateAccount{}
	})

    //default handler implementation
    o.CreateHandler = func(event *CreateAccount, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(CreateAccountEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(AccountCreatedEvent, func() eventhorizon.EventData {
		return &AccountCreated{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *AccountCreated, entity *Account) (err error) {
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
    eventhorizon.RegisterEventData(DeleteAccountEvent, func() eventhorizon.EventData {
		return &DeleteAccount{}
	})

    //default handler implementation
    o.DeleteHandler = func(event *DeleteAccount, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(DeleteAccountEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(AccountDeletedEvent, func() eventhorizon.EventData {
		return &AccountDeleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *AccountDeleted, entity *Account) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AccountAggregateType); err == nil {
            *entity = *NewAccount()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(AccountLoggedEvent, func() eventhorizon.EventData {
		return &AccountLogged{}
	})

    //default handler implementation
    o.LoggedHandler = func(event *AccountLogged, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(AccountLoggedEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(SendCreatedAccountConfirmationedEvent, func() eventhorizon.EventData {
		return &SendCreatedAccountConfirmationed{}
	})

    //default handler implementation
    o.SendCreatedConfirmationedHandler = func(event *SendCreatedAccountConfirmationed, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(SendCreatedAccountConfirmationedEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(SendEnabledAccountConfirmationedEvent, func() eventhorizon.EventData {
		return &SendEnabledAccountConfirmationed{}
	})

    //default handler implementation
    o.SendEnabledConfirmationedHandler = func(event *SendEnabledAccountConfirmationed, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(SendEnabledAccountConfirmationedEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(SendDisabledAccountConfirmationedEvent, func() eventhorizon.EventData {
		return &SendDisabledAccountConfirmationed{}
	})

    //default handler implementation
    o.SendDisabledConfirmationedHandler = func(event *SendDisabledAccountConfirmationed, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(SendDisabledAccountConfirmationedEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DisableAccountEvent, func() eventhorizon.EventData {
		return &DisableAccount{}
	})

    //default handler implementation
    o.DisableHandler = func(event *DisableAccount, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(DisableAccountEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(EnableAccountEvent, func() eventhorizon.EventData {
		return &EnableAccount{}
	})

    //default handler implementation
    o.EnableHandler = func(event *EnableAccount, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(EnableAccountEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdateAccountEvent, func() eventhorizon.EventData {
		return &UpdateAccount{}
	})

    //default handler implementation
    o.UpdateHandler = func(event *UpdateAccount, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(UpdateAccountEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(AccountUpdatedEvent, func() eventhorizon.EventData {
		return &AccountUpdated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *AccountUpdated, entity *Account) (err error) {
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
    eventhorizon.RegisterEventData(AccountEnabledEvent, func() eventhorizon.EventData {
		return &AccountEnabled{}
	})

    //default handler implementation
    o.EnabledHandler = func(event *AccountEnabled, entity *Account) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AccountAggregateType); err == nil {
            entity.Disabled = false
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(AccountDisabledEvent, func() eventhorizon.EventData {
		return &AccountDisabled{}
	})

    //default handler implementation
    o.DisabledHandler = func(event *AccountDisabled, entity *Account) (err error) {
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
    o.RegisterForEvent(handler, AccountEventTypes().AccountLogged())
}

func (o *AggregateInitializer) RegisterForSendCreatedConfirmationed(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountEventTypes().SendCreatedAccountConfirmationed())
}

func (o *AggregateInitializer) RegisterForSendEnabledConfirmationed(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountEventTypes().SendEnabledAccountConfirmationed())
}

func (o *AggregateInitializer) RegisterForSendDisabledConfirmationed(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountEventTypes().SendDisabledAccountConfirmationed())
}


func NewAccountAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
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

func NewAuthEventhorizonInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *AuthEventhorizonInitializer) {
    accountAggregateInitializer := NewAccountAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
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









