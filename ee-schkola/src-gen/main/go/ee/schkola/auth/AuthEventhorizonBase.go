package auth

import (
    "context"
    "errors"
    "fmt"
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

const AccountAggregateType eventhorizon.AggregateType = "AccountAggregate"

func NewAccountAggregate(id eventhorizon.UUID) (ret *AccountAggregate) {
    ret = &AccountAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(AccountAggregateType, id),
    }
	//ret.CommandHandler = NewAccountAggregateCommandHandler(ret)
    return
}

func (o *AccountAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type AccountAggregate struct {
    *eventhorizon.AggregateBase
    *Account
    eventhorizon.CommandHandler
}



type AccountCommandHandler struct {
    CreateHandler  func (*CreateAccount, *AccountAggregate) error
    DeleteHandler  func (*DeleteAccount, *AccountAggregate) error
    UpdateHandler  func (*UpdateAccount, *AccountAggregate) error
    EnableHandler  func (*EnableAccount, *AccountAggregate) error
    DisableHandler  func (*DisableAccount, *AccountAggregate) error
    RegisterHandler  func (*RegisterAccount, *AccountAggregate) error
}

func (o *AccountCommandHandler) HandleCommand(ctx *context.Context, cmd eventhorizon.Command, aggregate *AccountAggregate) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateAccountCommand:
        ret = o.CreateHandler(cmd.(*CreateAccount), aggregate)
    case DeleteAccountCommand:
        ret = o.DeleteHandler(cmd.(*DeleteAccount), aggregate)
    case UpdateAccountCommand:
        ret = o.UpdateHandler(cmd.(*UpdateAccount), aggregate)
    case EnableAccountCommand:
        ret = o.EnableHandler(cmd.(*EnableAccount), aggregate)
    case DisableAccountCommand:
        ret = o.DisableHandler(cmd.(*DisableAccount), aggregate)
    case RegisterAccountCommand:
        ret = o.RegisterHandler(cmd.(*RegisterAccount), aggregate)
    default:
		ret = errors.New(fmt.Sprintf("Wrong comand type '%v' for aggregate '%v", cmd.CommandType(), aggregate))
	}
    return ret
    
}



func NewAccountAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *AccountAggregateInitializer) {
	ret = &AccountAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(AccountAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate { return NewAccountAggregate(id) },
        AccountCommandTypes().Literals(), AccountEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
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









