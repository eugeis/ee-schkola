package auth

import (
    "errors"
    "fmt"
    "github.com/eugeis/gee/eh"
    "github.com/looplab/eventhorizon"
    "time"
)
type AccountCommandHandler struct {
    EnableHandler func (*EnableAccount, *Account, eh.AggregateStoreEvent) error
    DisableHandler func (*DisableAccount, *Account, eh.AggregateStoreEvent) error
    RegisterHandler func (*RegisterAccount, *Account, eh.AggregateStoreEvent) error
    CreateHandler func (*CreateAccount, *Account, eh.AggregateStoreEvent) error
    DeleteHandler func (*DeleteAccount, *Account, eh.AggregateStoreEvent) error
    UpdateHandler func (*UpdateAccount, *Account, eh.AggregateStoreEvent) error
}

func (o *AccountCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (ret error) {
    switch cmd.CommandType() {
    case EnableAccountCommand:
        ret = o.EnableHandler(cmd.(*EnableAccount), entity.(*Account), store)
    case DisableAccountCommand:
        ret = o.DisableHandler(cmd.(*DisableAccount), entity.(*Account), store)
    case RegisterAccountCommand:
        ret = o.RegisterHandler(cmd.(*RegisterAccount), entity.(*Account), store)
    case CreateAccountCommand:
        ret = o.CreateHandler(cmd.(*CreateAccount), entity.(*Account), store)
    case DeleteAccountCommand:
        ret = o.DeleteHandler(cmd.(*DeleteAccount), entity.(*Account), store)
    case UpdateAccountCommand:
        ret = o.UpdateHandler(cmd.(*UpdateAccount), entity.(*Account), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *AccountCommandHandler) SetupCommandHandler() (ret error) {
    if o.EnableHandler == nil {
        o.EnableHandler = func(command *EnableAccount, entity *Account, store eh.AggregateStoreEvent) (ret error) {ret = eh.CommandHandlerNotImplemented(EnableAccountCommand)
            return
        }
    }
    if o.DisableHandler == nil {
        o.DisableHandler = func(command *DisableAccount, entity *Account, store eh.AggregateStoreEvent) (ret error) {ret = eh.CommandHandlerNotImplemented(DisableAccountCommand)
            return
        }
    }
    if o.RegisterHandler == nil {
        o.RegisterHandler = func(command *RegisterAccount, entity *Account, store eh.AggregateStoreEvent) (ret error) {ret = eh.CommandHandlerNotImplemented(RegisterAccountCommand)
            return
        }
    }
    if o.CreateHandler == nil {
        o.CreateHandler = func(command *CreateAccount, entity *Account, store eh.AggregateStoreEvent) (ret error) {
            if ret = eh.ValidateNewId(entity.Id, command.Id, AccountAggregateType); ret == nil {store.StoreEvent(AccountCreatedEvent, &AccountCreated{
                    Id: command.Id,
                    Username: command.Username,
                    Password: command.Password,
                    Email: command.Email,
                    Disabled: command.Disabled,
                    LastLoginAt: command.LastLoginAt,
                    Profile: command.Profile,}, time.Now())
            }
            return
        }
    }
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteAccount, entity *Account, store eh.AggregateStoreEvent) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); ret == nil {
                store.StoreEvent(AccountDeletedEvent, &AccountDeleted{
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    if o.UpdateHandler == nil {
        o.UpdateHandler = func(command *UpdateAccount, entity *Account, store eh.AggregateStoreEvent) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); ret == nil {
                store.StoreEvent(AccountUpdatedEvent, &AccountUpdated{
                    Id: command.Id,
                    Username: command.Username,
                    Password: command.Password,
                    Email: command.Email,
                    Disabled: command.Disabled,
                    LastLoginAt: command.LastLoginAt,
                    Profile: command.Profile,}, time.Now())
            }
            return
        }
    }
    return
}


type AccountEventHandler struct {
    CreatedHandler func (*AccountCreated, *Account) error
    DeletedHandler func (*AccountDeleted, *Account) error
    UpdatedHandler func (*AccountUpdated, *Account) error
}

func (o *AccountEventHandler) Apply(event eventhorizon.Event, entity interface{}) (ret error) {
    switch event.EventType() {
    case AccountCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*AccountCreated), entity.(*Account))
    case AccountDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*AccountDeleted), entity.(*Account))
    case AccountUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*AccountUpdated), entity.(*Account))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *AccountEventHandler) SetupEventHandler() (ret error) {
    if o.CreatedHandler == nil {
        o.CreatedHandler = func(event *AccountCreated, entity *Account) (ret error) {
            if ret = eh.ValidateNewId(entity.Id, event.Id, AccountAggregateType); ret == nil {
                entity.Id = event.Id
                entity.Username = event.Username
                entity.Password = event.Password
                entity.Email = event.Email
                entity.Disabled = event.Disabled
                entity.LastLoginAt = event.LastLoginAt
                entity.Profile = event.Profile
            }
            return
        }
    }
    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *AccountDeleted, entity *Account) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, event.Id, AccountAggregateType); ret == nil {
                *entity = *NewAccount()
            }
            return
        }
    }
    if o.UpdatedHandler == nil {
        o.UpdatedHandler = func(event *AccountUpdated, entity *Account) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, event.Id, AccountAggregateType); ret == nil {
                entity.Username = event.Username
                entity.Password = event.Password
                entity.Email = event.Email
                entity.Disabled = event.Disabled
                entity.LastLoginAt = event.LastLoginAt
                entity.Profile = event.Profile
            }
            return
        }
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



func NewAccountAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus, readRepos func (string) eventhorizon.ReadWriteRepo) (ret *AccountAggregateInitializer) {
    
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
                commandBus eventhorizon.CommandBus, readRepos func (string) eventhorizon.ReadWriteRepo) (ret *AuthEventhorizonInitializer) {
    AccountAggregateInitializer:= NewAccountAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    ret = &AuthEventhorizonInitializer{
        eventStore: eventStore,
        eventBus: eventBus,
        eventPublisher: eventPublisher,
        commandBus: commandBus,
        AccountAggregateInitializer: AccountAggregateInitializer,
    }
    return
}

func (o *AuthEventhorizonInitializer) Setup() (ret error) {
    
    if ret = o.AccountAggregateInitializer.Setup(); ret != nil {
        return
    }

    return
}









