package auth

import (
    "errors"
    "fmt"
    "github.com/eugeis/gee/eh"
    "github.com/looplab/eventhorizon"
    "time"
)
type AccountCommandHandler struct {
    LoginHandler func (*LoginAccount, *Account, eh.AggregateStoreEvent) (err error) 
    CreateHandler func (*CreateAccount, *Account, eh.AggregateStoreEvent) (err error) 
    DeleteHandler func (*DeleteAccount, *Account, eh.AggregateStoreEvent) (err error) 
    EnableHandler func (*EnableAccount, *Account, eh.AggregateStoreEvent) (err error) 
    DisableHandler func (*DisableAccount, *Account, eh.AggregateStoreEvent) (err error) 
    UpdateHandler func (*UpdateAccount, *Account, eh.AggregateStoreEvent) (err error) 
}

func (o *AccountCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case LoginAccountCommand:
        err = o.LoginHandler(cmd.(*LoginAccount), entity.(*Account), store)
    case CreateAccountCommand:
        err = o.CreateHandler(cmd.(*CreateAccount), entity.(*Account), store)
    case DeleteAccountCommand:
        err = o.DeleteHandler(cmd.(*DeleteAccount), entity.(*Account), store)
    case EnableAccountCommand:
        err = o.EnableHandler(cmd.(*EnableAccount), entity.(*Account), store)
    case DisableAccountCommand:
        err = o.DisableHandler(cmd.(*DisableAccount), entity.(*Account), store)
    case UpdateAccountCommand:
        err = o.UpdateHandler(cmd.(*UpdateAccount), entity.(*Account), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *AccountCommandHandler) SetupCommandHandler() (err error) {
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
    o.CreateHandler = func(command *CreateAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(AccountCreatedEvent, &AccountCreated{
                Name: command.Name,
                Username: command.Username,
                Password: command.Password,
                Email: command.Email,
                Disabled: command.Disabled,
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
    o.EnableHandler = func(command *EnableAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(AccountEnabledEvent, &AccountEnabled{
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
    o.UpdateHandler = func(command *UpdateAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.StoreEvent(AccountUpdatedEvent, &AccountUpdated{
                Name: command.Name,
                Username: command.Username,
                Password: command.Password,
                Email: command.Email,
                Disabled: command.Disabled,
                Profile: command.Profile,
                Id: command.Id,}, time.Now())
        }
        return
    }
    return
}


type AccountEventHandler struct {
    CreatedHandler func (*AccountCreated, *Account) (err error) 
    DeletedHandler func (*AccountDeleted, *Account) (err error) 
    LoggedHandler func (*AccountLogged, *Account) (err error) 
    UpdatedHandler func (*AccountUpdated, *Account) (err error) 
    EnabledHandler func (*AccountEnabled, *Account) (err error) 
    DisabledHandler func (*AccountDisabled, *Account) (err error) 
}

func (o *AccountEventHandler) Apply(event eventhorizon.Event, entity interface{}) (err error) {
    switch event.EventType() {
    case AccountCreatedEvent:
        err = o.CreatedHandler(event.Data().(*AccountCreated), entity.(*Account))
    case AccountDeletedEvent:
        err = o.DeletedHandler(event.Data().(*AccountDeleted), entity.(*Account))
    case AccountLoggedEvent:
        err = o.LoggedHandler(event.Data().(*AccountLogged), entity.(*Account))
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

func (o *AccountEventHandler) SetupEventHandler() (err error) {

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
            entity.Profile = event.Profile
            entity.Id = event.Id
        }
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

type AccountAggregateInitializer struct {
    *eh.AggregateInitializer
    *AccountCommandHandler
    *AccountEventHandler
    ProjectorHandler *AccountEventHandler
}


func (o *AccountAggregateInitializer) RegisterForLogged(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountEventTypes().AccountLogged())
}


func NewAccountAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus, 
                readRepos func (string, func () (ret interface{}) ) (ret eventhorizon.ReadWriteRepo) ) (ret *AccountAggregateInitializer) {
    
    commandHandler := &AccountCommandHandler{}
    eventHandler := &AccountEventHandler{}
    modelFactory := func() interface{} { return NewAccount() }
    ret = &AccountAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(AccountAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(AccountAggregateType, id, commandHandler, eventHandler, modelFactory())
        }, modelFactory,
        AccountCommandTypes().Literals(), AccountEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), AccountCommandHandler: commandHandler, AccountEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type AuthEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore
    eventBus eventhorizon.EventBus
    eventPublisher eventhorizon.EventPublisher
    commandBus eventhorizon.CommandBus
    AccountAggregateInitializer *AccountAggregateInitializer
}

func NewAuthEventhorizonInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus, 
                readRepos func (string, func () (ret interface{}) ) (ret eventhorizon.ReadWriteRepo) ) (ret *AuthEventhorizonInitializer) {
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









