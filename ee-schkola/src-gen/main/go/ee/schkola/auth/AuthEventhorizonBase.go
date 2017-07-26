package auth

import (
    "errors"
    "fmt"
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

type AccountCommandHandler struct {
    CreateHandler  func (*CreateAccount, *Account, eh.AggregateStoreEvent) error
    DeleteHandler  func (*DeleteAccount, *Account, eh.AggregateStoreEvent) error
    UpdateHandler  func (*UpdateAccount, *Account, eh.AggregateStoreEvent) error
    EnableHandler  func (*EnableAccount, *Account, eh.AggregateStoreEvent) error
    DisableHandler  func (*DisableAccount, *Account, eh.AggregateStoreEvent) error
    RegisterHandler  func (*RegisterAccount, *Account, eh.AggregateStoreEvent) error
}

func (o *AccountCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateAccountCommand:
        ret = o.CreateHandler(cmd.(*CreateAccount), entity.(*Account), store)
    case DeleteAccountCommand:
        ret = o.DeleteHandler(cmd.(*DeleteAccount), entity.(*Account), store)
    case UpdateAccountCommand:
        ret = o.UpdateHandler(cmd.(*UpdateAccount), entity.(*Account), store)
    case EnableAccountCommand:
        ret = o.EnableHandler(cmd.(*EnableAccount), entity.(*Account), store)
    case DisableAccountCommand:
        ret = o.DisableHandler(cmd.(*DisableAccount), entity.(*Account), store)
    case RegisterAccountCommand:
        ret = o.RegisterHandler(cmd.(*RegisterAccount), entity.(*Account), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return ret
    
}

func (o *AccountCommandHandler) SetupCommandHandler() error {
    
    var ret error
    return ret
    
}



type AccountEventHandler struct {
    CreatedHandler  func (*AccountCreated, *Account) error
    DeletedHandler  func (*AccountDeleted, *Account) error
    UpdatedHandler  func (*AccountUpdated, *Account) error
}

func (o *AccountEventHandler) Apply(event eventhorizon.Event, entity interface{}) error {
    
    var ret error
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
    return ret
    
}

func (o *AccountEventHandler) SetupEventHandler() error {
    
    var ret error
    return ret
    
}



const AccountAggregateType eventhorizon.AggregateType = "AccountAggregateInitializer"

func NewAccountAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *AccountAggregateInitializer) {
    commandHandler := &AccountCommandHandler{}
    eventHandler := &AccountEventHandler{}
	ret = &AccountAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(AccountAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(AccountAggregateType, id, commandHandler, eventHandler, &Account{})
        },
        AccountCommandTypes().Literals(), AccountEventTypes().Literals(),
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus),
        AccountCommandHandler: commandHandler,
        AccountEventHandler: eventHandler,
    }
	return
}


func (o *AccountAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountEventTypes().AccountCreated())
}

func (o *AccountAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountEventTypes().AccountDeleted())
}

func (o *AccountAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountEventTypes().AccountUpdated())
}

type AccountAggregateInitializer struct {
    *eh.AggregateInitializer
    *AccountCommandHandler
    *AccountEventHandler
}



func NewAuthEventhorizonInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *AuthEventhorizonInitializer) {
	ret = &AuthEventhorizonInitializer{eventStore: eventStore, eventBus: eventBus, eventPublisher: eventPublisher,
            commandBus: commandBus, 
    AccountAggregateInitializer: NewAccountAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus)}
	return
}

func (o *AuthEventhorizonInitializer) Setup() (err error) {
    
    if err = o.AccountAggregateInitializer.Setup(); err != nil {
        return
    }
    return
}

type AuthEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore
    eventBus eventhorizon.EventBus
    eventPublisher eventhorizon.EventPublisher
    commandBus eventhorizon.CommandBus
    AccountAggregateInitializer  *AccountAggregateInitializer
}









