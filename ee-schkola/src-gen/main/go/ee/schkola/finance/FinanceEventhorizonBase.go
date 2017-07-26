package finance

import (
    "errors"
    "fmt"
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

type ExpenseCommandHandler struct {
    CreateHandler  func (*CreateExpense, *Expense, eh.AggregateStoreEvent) error
    DeleteHandler  func (*DeleteExpense, *Expense, eh.AggregateStoreEvent) error
    UpdateHandler  func (*UpdateExpense, *Expense, eh.AggregateStoreEvent) error
}

func (o *ExpenseCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (ret error) {
    
    switch cmd.CommandType() {
    case CreateExpenseCommand:
        ret = o.CreateHandler(cmd.(*CreateExpense), entity.(*Expense), store)
    case DeleteExpenseCommand:
        ret = o.DeleteHandler(cmd.(*DeleteExpense), entity.(*Expense), store)
    case UpdateExpenseCommand:
        ret = o.UpdateHandler(cmd.(*UpdateExpense), entity.(*Expense), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
    
}

func (o *ExpenseCommandHandler) SetupCommandHandler() (ret error) {
    
    if o.CreateHandler == nil {
        o.CreateHandler = func(command *CreateExpense, entity *Expense,
            store eh.AggregateStoreEvent) (ret error) {
            return
        }
    }
    
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteExpense, entity *Expense,
            store eh.AggregateStoreEvent) (ret error) {
            return
        }
    }
    
    if o.UpdateHandler == nil {
        o.UpdateHandler = func(command *UpdateExpense, entity *Expense,
            store eh.AggregateStoreEvent) (ret error) {
            return
        }
    }
    
    return
    
}



type ExpenseEventHandler struct {
    CreatedHandler  func (*ExpenseCreated, *Expense) error
    DeletedHandler  func (*ExpenseDeleted, *Expense) error
    UpdatedHandler  func (*ExpenseUpdated, *Expense) error
}

func (o *ExpenseEventHandler) Apply(event eventhorizon.Event, entity interface{}) (ret error) {
    
    switch event.EventType() {
    case ExpenseCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*ExpenseCreated), entity.(*Expense))
    case ExpenseDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*ExpenseDeleted), entity.(*Expense))
    case ExpenseUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*ExpenseUpdated), entity.(*Expense))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
    
}

func (o *ExpenseEventHandler) SetupEventHandler() (ret error) {
    
    if o.CreatedHandler == nil {
        o.CreatedHandler = func(event *ExpenseCreated, entity *Expense) (ret error) {
            return
        }
    }
    
    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *ExpenseDeleted, entity *Expense) (ret error) {
            return
        }
    }
    
    if o.UpdatedHandler == nil {
        o.UpdatedHandler = func(event *ExpenseUpdated, entity *Expense) (ret error) {
            return
        }
    }
    
    return
    
}



const ExpenseAggregateType eventhorizon.AggregateType = "ExpenseAggregateInitializer"

func NewExpenseAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *ExpenseAggregateInitializer) {
    commandHandler := &ExpenseCommandHandler{}
    eventHandler := &ExpenseEventHandler{}
	ret = &ExpenseAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ExpenseAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(ExpenseAggregateType, id, commandHandler, eventHandler, &Expense{})
        },
        ExpenseCommandTypes().Literals(), ExpenseEventTypes().Literals(),
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus),
        ExpenseCommandHandler: commandHandler,
        ExpenseEventHandler: eventHandler,
    }
	return
}


func (o *ExpenseAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ExpenseEventTypes().ExpenseCreated())
}

func (o *ExpenseAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ExpenseEventTypes().ExpenseDeleted())
}

func (o *ExpenseAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ExpenseEventTypes().ExpenseUpdated())
}

type ExpenseAggregateInitializer struct {
    *eh.AggregateInitializer
    *ExpenseCommandHandler
    *ExpenseEventHandler
}



type ExpensePurposeCommandHandler struct {
    CreateHandler  func (*CreateExpensePurpose, *ExpensePurpose, eh.AggregateStoreEvent) error
    DeleteHandler  func (*DeleteExpensePurpose, *ExpensePurpose, eh.AggregateStoreEvent) error
    UpdateHandler  func (*UpdateExpensePurpose, *ExpensePurpose, eh.AggregateStoreEvent) error
}

func (o *ExpensePurposeCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (ret error) {
    
    switch cmd.CommandType() {
    case CreateExpensePurposeCommand:
        ret = o.CreateHandler(cmd.(*CreateExpensePurpose), entity.(*ExpensePurpose), store)
    case DeleteExpensePurposeCommand:
        ret = o.DeleteHandler(cmd.(*DeleteExpensePurpose), entity.(*ExpensePurpose), store)
    case UpdateExpensePurposeCommand:
        ret = o.UpdateHandler(cmd.(*UpdateExpensePurpose), entity.(*ExpensePurpose), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
    
}

func (o *ExpensePurposeCommandHandler) SetupCommandHandler() (ret error) {
    
    if o.CreateHandler == nil {
        o.CreateHandler = func(command *CreateExpensePurpose, entity *ExpensePurpose,
            store eh.AggregateStoreEvent) (ret error) {
            return
        }
    }
    
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteExpensePurpose, entity *ExpensePurpose,
            store eh.AggregateStoreEvent) (ret error) {
            return
        }
    }
    
    if o.UpdateHandler == nil {
        o.UpdateHandler = func(command *UpdateExpensePurpose, entity *ExpensePurpose,
            store eh.AggregateStoreEvent) (ret error) {
            return
        }
    }
    
    return
    
}



type ExpensePurposeEventHandler struct {
    CreatedHandler  func (*ExpensePurposeCreated, *ExpensePurpose) error
    DeletedHandler  func (*ExpensePurposeDeleted, *ExpensePurpose) error
    UpdatedHandler  func (*ExpensePurposeUpdated, *ExpensePurpose) error
}

func (o *ExpensePurposeEventHandler) Apply(event eventhorizon.Event, entity interface{}) (ret error) {
    
    switch event.EventType() {
    case ExpensePurposeCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*ExpensePurposeCreated), entity.(*ExpensePurpose))
    case ExpensePurposeDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*ExpensePurposeDeleted), entity.(*ExpensePurpose))
    case ExpensePurposeUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*ExpensePurposeUpdated), entity.(*ExpensePurpose))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
    
}

func (o *ExpensePurposeEventHandler) SetupEventHandler() (ret error) {
    
    if o.CreatedHandler == nil {
        o.CreatedHandler = func(event *ExpensePurposeCreated, entity *ExpensePurpose) (ret error) {
            return
        }
    }
    
    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *ExpensePurposeDeleted, entity *ExpensePurpose) (ret error) {
            return
        }
    }
    
    if o.UpdatedHandler == nil {
        o.UpdatedHandler = func(event *ExpensePurposeUpdated, entity *ExpensePurpose) (ret error) {
            return
        }
    }
    
    return
    
}



const ExpensePurposeAggregateType eventhorizon.AggregateType = "ExpensePurposeAggregateInitializer"

func NewExpensePurposeAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *ExpensePurposeAggregateInitializer) {
    commandHandler := &ExpensePurposeCommandHandler{}
    eventHandler := &ExpensePurposeEventHandler{}
	ret = &ExpensePurposeAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ExpensePurposeAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(ExpensePurposeAggregateType, id, commandHandler, eventHandler, &ExpensePurpose{})
        },
        ExpensePurposeCommandTypes().Literals(), ExpensePurposeEventTypes().Literals(),
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus),
        ExpensePurposeCommandHandler: commandHandler,
        ExpensePurposeEventHandler: eventHandler,
    }
	return
}


func (o *ExpensePurposeAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ExpensePurposeEventTypes().ExpensePurposeCreated())
}

func (o *ExpensePurposeAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ExpensePurposeEventTypes().ExpensePurposeDeleted())
}

func (o *ExpensePurposeAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ExpensePurposeEventTypes().ExpensePurposeUpdated())
}

type ExpensePurposeAggregateInitializer struct {
    *eh.AggregateInitializer
    *ExpensePurposeCommandHandler
    *ExpensePurposeEventHandler
}



type FeeCommandHandler struct {
    CreateHandler  func (*CreateFee, *Fee, eh.AggregateStoreEvent) error
    DeleteHandler  func (*DeleteFee, *Fee, eh.AggregateStoreEvent) error
    UpdateHandler  func (*UpdateFee, *Fee, eh.AggregateStoreEvent) error
}

func (o *FeeCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (ret error) {
    
    switch cmd.CommandType() {
    case CreateFeeCommand:
        ret = o.CreateHandler(cmd.(*CreateFee), entity.(*Fee), store)
    case DeleteFeeCommand:
        ret = o.DeleteHandler(cmd.(*DeleteFee), entity.(*Fee), store)
    case UpdateFeeCommand:
        ret = o.UpdateHandler(cmd.(*UpdateFee), entity.(*Fee), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
    
}

func (o *FeeCommandHandler) SetupCommandHandler() (ret error) {
    
    if o.CreateHandler == nil {
        o.CreateHandler = func(command *CreateFee, entity *Fee,
            store eh.AggregateStoreEvent) (ret error) {
            return
        }
    }
    
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteFee, entity *Fee,
            store eh.AggregateStoreEvent) (ret error) {
            return
        }
    }
    
    if o.UpdateHandler == nil {
        o.UpdateHandler = func(command *UpdateFee, entity *Fee,
            store eh.AggregateStoreEvent) (ret error) {
            return
        }
    }
    
    return
    
}



type FeeEventHandler struct {
    CreatedHandler  func (*FeeCreated, *Fee) error
    DeletedHandler  func (*FeeDeleted, *Fee) error
    UpdatedHandler  func (*FeeUpdated, *Fee) error
}

func (o *FeeEventHandler) Apply(event eventhorizon.Event, entity interface{}) (ret error) {
    
    switch event.EventType() {
    case FeeCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*FeeCreated), entity.(*Fee))
    case FeeDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*FeeDeleted), entity.(*Fee))
    case FeeUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*FeeUpdated), entity.(*Fee))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
    
}

func (o *FeeEventHandler) SetupEventHandler() (ret error) {
    
    if o.CreatedHandler == nil {
        o.CreatedHandler = func(event *FeeCreated, entity *Fee) (ret error) {
            return
        }
    }
    
    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *FeeDeleted, entity *Fee) (ret error) {
            return
        }
    }
    
    if o.UpdatedHandler == nil {
        o.UpdatedHandler = func(event *FeeUpdated, entity *Fee) (ret error) {
            return
        }
    }
    
    return
    
}



const FeeAggregateType eventhorizon.AggregateType = "FeeAggregateInitializer"

func NewFeeAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *FeeAggregateInitializer) {
    commandHandler := &FeeCommandHandler{}
    eventHandler := &FeeEventHandler{}
	ret = &FeeAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(FeeAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(FeeAggregateType, id, commandHandler, eventHandler, &Fee{})
        },
        FeeCommandTypes().Literals(), FeeEventTypes().Literals(),
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus),
        FeeCommandHandler: commandHandler,
        FeeEventHandler: eventHandler,
    }
	return
}


func (o *FeeAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, FeeEventTypes().FeeCreated())
}

func (o *FeeAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, FeeEventTypes().FeeDeleted())
}

func (o *FeeAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, FeeEventTypes().FeeUpdated())
}

type FeeAggregateInitializer struct {
    *eh.AggregateInitializer
    *FeeCommandHandler
    *FeeEventHandler
}



type FeeKindCommandHandler struct {
    CreateHandler  func (*CreateFeeKind, *FeeKind, eh.AggregateStoreEvent) error
    DeleteHandler  func (*DeleteFeeKind, *FeeKind, eh.AggregateStoreEvent) error
    UpdateHandler  func (*UpdateFeeKind, *FeeKind, eh.AggregateStoreEvent) error
}

func (o *FeeKindCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (ret error) {
    
    switch cmd.CommandType() {
    case CreateFeeKindCommand:
        ret = o.CreateHandler(cmd.(*CreateFeeKind), entity.(*FeeKind), store)
    case DeleteFeeKindCommand:
        ret = o.DeleteHandler(cmd.(*DeleteFeeKind), entity.(*FeeKind), store)
    case UpdateFeeKindCommand:
        ret = o.UpdateHandler(cmd.(*UpdateFeeKind), entity.(*FeeKind), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
    
}

func (o *FeeKindCommandHandler) SetupCommandHandler() (ret error) {
    
    if o.CreateHandler == nil {
        o.CreateHandler = func(command *CreateFeeKind, entity *FeeKind,
            store eh.AggregateStoreEvent) (ret error) {
            return
        }
    }
    
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteFeeKind, entity *FeeKind,
            store eh.AggregateStoreEvent) (ret error) {
            return
        }
    }
    
    if o.UpdateHandler == nil {
        o.UpdateHandler = func(command *UpdateFeeKind, entity *FeeKind,
            store eh.AggregateStoreEvent) (ret error) {
            return
        }
    }
    
    return
    
}



type FeeKindEventHandler struct {
    CreatedHandler  func (*FeeKindCreated, *FeeKind) error
    DeletedHandler  func (*FeeKindDeleted, *FeeKind) error
    UpdatedHandler  func (*FeeKindUpdated, *FeeKind) error
}

func (o *FeeKindEventHandler) Apply(event eventhorizon.Event, entity interface{}) (ret error) {
    
    switch event.EventType() {
    case FeeKindCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*FeeKindCreated), entity.(*FeeKind))
    case FeeKindDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*FeeKindDeleted), entity.(*FeeKind))
    case FeeKindUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*FeeKindUpdated), entity.(*FeeKind))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
    
}

func (o *FeeKindEventHandler) SetupEventHandler() (ret error) {
    
    if o.CreatedHandler == nil {
        o.CreatedHandler = func(event *FeeKindCreated, entity *FeeKind) (ret error) {
            return
        }
    }
    
    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *FeeKindDeleted, entity *FeeKind) (ret error) {
            return
        }
    }
    
    if o.UpdatedHandler == nil {
        o.UpdatedHandler = func(event *FeeKindUpdated, entity *FeeKind) (ret error) {
            return
        }
    }
    
    return
    
}



const FeeKindAggregateType eventhorizon.AggregateType = "FeeKindAggregateInitializer"

func NewFeeKindAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *FeeKindAggregateInitializer) {
    commandHandler := &FeeKindCommandHandler{}
    eventHandler := &FeeKindEventHandler{}
	ret = &FeeKindAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(FeeKindAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(FeeKindAggregateType, id, commandHandler, eventHandler, &FeeKind{})
        },
        FeeKindCommandTypes().Literals(), FeeKindEventTypes().Literals(),
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus),
        FeeKindCommandHandler: commandHandler,
        FeeKindEventHandler: eventHandler,
    }
	return
}


func (o *FeeKindAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, FeeKindEventTypes().FeeKindCreated())
}

func (o *FeeKindAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, FeeKindEventTypes().FeeKindDeleted())
}

func (o *FeeKindAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, FeeKindEventTypes().FeeKindUpdated())
}

type FeeKindAggregateInitializer struct {
    *eh.AggregateInitializer
    *FeeKindCommandHandler
    *FeeKindEventHandler
}



func NewFinanceEventhorizonInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *FinanceEventhorizonInitializer) {
	ret = &FinanceEventhorizonInitializer{eventStore: eventStore, eventBus: eventBus, eventPublisher: eventPublisher,
            commandBus: commandBus, 
    ExpenseAggregateInitializer: NewExpenseAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    ExpensePurposeAggregateInitializer: NewExpensePurposeAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    FeeAggregateInitializer: NewFeeAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    FeeKindAggregateInitializer: NewFeeKindAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus)}
	return
}

func (o *FinanceEventhorizonInitializer) Setup() (err error) {
    
    if err = o.ExpenseAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.ExpensePurposeAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.FeeAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.FeeKindAggregateInitializer.Setup(); err != nil {
        return
    }
    return
}

type FinanceEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore
    eventBus eventhorizon.EventBus
    eventPublisher eventhorizon.EventPublisher
    commandBus eventhorizon.CommandBus
    ExpenseAggregateInitializer  *ExpenseAggregateInitializer
    ExpensePurposeAggregateInitializer  *ExpensePurposeAggregateInitializer
    FeeAggregateInitializer  *FeeAggregateInitializer
    FeeKindAggregateInitializer  *FeeKindAggregateInitializer
}









