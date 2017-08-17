package finance

import (
    "errors"
    "fmt"
    "github.com/eugeis/gee/eh"
    "github.com/looplab/eventhorizon"
    "time"
)
type ExpenseCommandHandler struct {
    CreateHandler func (*CreateExpense, *Expense, eh.AggregateStoreEvent) error
    DeleteHandler func (*DeleteExpense, *Expense, eh.AggregateStoreEvent) error
    UpdateHandler func (*UpdateExpense, *Expense, eh.AggregateStoreEvent) error
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
        o.CreateHandler = func(command *CreateExpense, entity *Expense, store eh.AggregateStoreEvent) (ret error) {
            if ret = eh.ValidateNewId(entity.Id, command.Id, ExpenseAggregateType); ret == nil {store.StoreEvent(ExpenseCreatedEvent, &ExpenseCreated{
                    Id: command.Id,
                    Purpose: command.Purpose,
                    Amount: command.Amount,
                    Profile: command.Profile,
                    Date: command.Date,}, time.Now())
            }
            return
        }
    }
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteExpense, entity *Expense, store eh.AggregateStoreEvent) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, command.Id, ExpenseAggregateType); ret == nil {
                store.StoreEvent(ExpenseDeletedEvent, &ExpenseDeleted{
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    if o.UpdateHandler == nil {
        o.UpdateHandler = func(command *UpdateExpense, entity *Expense, store eh.AggregateStoreEvent) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, command.Id, ExpenseAggregateType); ret == nil {
                store.StoreEvent(ExpenseUpdatedEvent, &ExpenseUpdated{
                    Id: command.Id,
                    Purpose: command.Purpose,
                    Amount: command.Amount,
                    Profile: command.Profile,
                    Date: command.Date,}, time.Now())
            }
            return
        }
    }
    return
}


type ExpenseEventHandler struct {
    CreatedHandler func (*ExpenseCreated, *Expense) error
    DeletedHandler func (*ExpenseDeleted, *Expense) error
    UpdatedHandler func (*ExpenseUpdated, *Expense) error
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
            if ret = eh.ValidateNewId(entity.Id, event.Id, ExpenseAggregateType); ret == nil {
                entity.Id = event.Id
                entity.Purpose = event.Purpose
                entity.Amount = event.Amount
                entity.Profile = event.Profile
                entity.Date = event.Date
            }
            return
        }
    }
    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *ExpenseDeleted, entity *Expense) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, event.Id, ExpenseAggregateType); ret == nil {
                *entity = *NewExpense()
            }
            return
        }
    }
    if o.UpdatedHandler == nil {
        o.UpdatedHandler = func(event *ExpenseUpdated, entity *Expense) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, event.Id, ExpenseAggregateType); ret == nil {
                entity.Purpose = event.Purpose
                entity.Amount = event.Amount
                entity.Profile = event.Profile
                entity.Date = event.Date
            }
            return
        }
    }
    return
}


const ExpenseAggregateType eventhorizon.AggregateType = "Expense"

type ExpenseAggregateInitializer struct {
    *eh.AggregateInitializer
    *ExpenseCommandHandler
    *ExpenseEventHandler
}



func NewExpenseAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus) (ret *ExpenseAggregateInitializer) {
    
    commandHandler := &ExpenseCommandHandler{}
    eventHandler := &ExpenseEventHandler{}
	ret = &ExpenseAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ExpenseAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(ExpenseAggregateType, id, commandHandler, eventHandler, NewExpense())
        },
        ExpenseCommandTypes().Literals(), ExpenseEventTypes().Literals(),
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus),
        ExpenseCommandHandler: commandHandler,
        ExpenseEventHandler: eventHandler,
    }

    return
}


type ExpensePurposeCommandHandler struct {
    CreateHandler func (*CreateExpensePurpose, *ExpensePurpose, eh.AggregateStoreEvent) error
    DeleteHandler func (*DeleteExpensePurpose, *ExpensePurpose, eh.AggregateStoreEvent) error
    UpdateHandler func (*UpdateExpensePurpose, *ExpensePurpose, eh.AggregateStoreEvent) error
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
        o.CreateHandler = func(command *CreateExpensePurpose, entity *ExpensePurpose, store eh.AggregateStoreEvent) (ret error) {
            if ret = eh.ValidateNewId(entity.Id, command.Id, ExpensePurposeAggregateType); ret == nil {store.StoreEvent(ExpensePurposeCreatedEvent, &ExpensePurposeCreated{
                    Id: command.Id,
                    Name: command.Name,
                    Description: command.Description,}, time.Now())
            }
            return
        }
    }
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteExpensePurpose, entity *ExpensePurpose, store eh.AggregateStoreEvent) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, command.Id, ExpensePurposeAggregateType); ret == nil {
                store.StoreEvent(ExpensePurposeDeletedEvent, &ExpensePurposeDeleted{
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    if o.UpdateHandler == nil {
        o.UpdateHandler = func(command *UpdateExpensePurpose, entity *ExpensePurpose, store eh.AggregateStoreEvent) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, command.Id, ExpensePurposeAggregateType); ret == nil {
                store.StoreEvent(ExpensePurposeUpdatedEvent, &ExpensePurposeUpdated{
                    Id: command.Id,
                    Name: command.Name,
                    Description: command.Description,}, time.Now())
            }
            return
        }
    }
    return
}


type ExpensePurposeEventHandler struct {
    CreatedHandler func (*ExpensePurposeCreated, *ExpensePurpose) error
    DeletedHandler func (*ExpensePurposeDeleted, *ExpensePurpose) error
    UpdatedHandler func (*ExpensePurposeUpdated, *ExpensePurpose) error
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
            if ret = eh.ValidateNewId(entity.Id, event.Id, ExpensePurposeAggregateType); ret == nil {
                entity.Id = event.Id
                entity.Name = event.Name
                entity.Description = event.Description
            }
            return
        }
    }
    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *ExpensePurposeDeleted, entity *ExpensePurpose) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, event.Id, ExpensePurposeAggregateType); ret == nil {
                *entity = *NewExpensePurpose()
            }
            return
        }
    }
    if o.UpdatedHandler == nil {
        o.UpdatedHandler = func(event *ExpensePurposeUpdated, entity *ExpensePurpose) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, event.Id, ExpensePurposeAggregateType); ret == nil {
                entity.Name = event.Name
                entity.Description = event.Description
            }
            return
        }
    }
    return
}


const ExpensePurposeAggregateType eventhorizon.AggregateType = "ExpensePurpose"

type ExpensePurposeAggregateInitializer struct {
    *eh.AggregateInitializer
    *ExpensePurposeCommandHandler
    *ExpensePurposeEventHandler
}



func NewExpensePurposeAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus) (ret *ExpensePurposeAggregateInitializer) {
    
    commandHandler := &ExpensePurposeCommandHandler{}
    eventHandler := &ExpensePurposeEventHandler{}
	ret = &ExpensePurposeAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ExpensePurposeAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(ExpensePurposeAggregateType, id, commandHandler, eventHandler, NewExpensePurpose())
        },
        ExpensePurposeCommandTypes().Literals(), ExpensePurposeEventTypes().Literals(),
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus),
        ExpensePurposeCommandHandler: commandHandler,
        ExpensePurposeEventHandler: eventHandler,
    }

    return
}


type FeeCommandHandler struct {
    CreateHandler func (*CreateFee, *Fee, eh.AggregateStoreEvent) error
    DeleteHandler func (*DeleteFee, *Fee, eh.AggregateStoreEvent) error
    UpdateHandler func (*UpdateFee, *Fee, eh.AggregateStoreEvent) error
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
        o.CreateHandler = func(command *CreateFee, entity *Fee, store eh.AggregateStoreEvent) (ret error) {
            if ret = eh.ValidateNewId(entity.Id, command.Id, FeeAggregateType); ret == nil {store.StoreEvent(FeeCreatedEvent, &FeeCreated{
                    Id: command.Id,
                    Student: command.Student,
                    Amount: command.Amount,
                    Kind: command.Kind,
                    Date: command.Date,}, time.Now())
            }
            return
        }
    }
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteFee, entity *Fee, store eh.AggregateStoreEvent) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, command.Id, FeeAggregateType); ret == nil {
                store.StoreEvent(FeeDeletedEvent, &FeeDeleted{
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    if o.UpdateHandler == nil {
        o.UpdateHandler = func(command *UpdateFee, entity *Fee, store eh.AggregateStoreEvent) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, command.Id, FeeAggregateType); ret == nil {
                store.StoreEvent(FeeUpdatedEvent, &FeeUpdated{
                    Id: command.Id,
                    Student: command.Student,
                    Amount: command.Amount,
                    Kind: command.Kind,
                    Date: command.Date,}, time.Now())
            }
            return
        }
    }
    return
}


type FeeEventHandler struct {
    CreatedHandler func (*FeeCreated, *Fee) error
    DeletedHandler func (*FeeDeleted, *Fee) error
    UpdatedHandler func (*FeeUpdated, *Fee) error
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
            if ret = eh.ValidateNewId(entity.Id, event.Id, FeeAggregateType); ret == nil {
                entity.Id = event.Id
                entity.Student = event.Student
                entity.Amount = event.Amount
                entity.Kind = event.Kind
                entity.Date = event.Date
            }
            return
        }
    }
    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *FeeDeleted, entity *Fee) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, event.Id, FeeAggregateType); ret == nil {
                *entity = *NewFee()
            }
            return
        }
    }
    if o.UpdatedHandler == nil {
        o.UpdatedHandler = func(event *FeeUpdated, entity *Fee) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, event.Id, FeeAggregateType); ret == nil {
                entity.Student = event.Student
                entity.Amount = event.Amount
                entity.Kind = event.Kind
                entity.Date = event.Date
            }
            return
        }
    }
    return
}


const FeeAggregateType eventhorizon.AggregateType = "Fee"

type FeeAggregateInitializer struct {
    *eh.AggregateInitializer
    *FeeCommandHandler
    *FeeEventHandler
}



func NewFeeAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus) (ret *FeeAggregateInitializer) {
    
    commandHandler := &FeeCommandHandler{}
    eventHandler := &FeeEventHandler{}
	ret = &FeeAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(FeeAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(FeeAggregateType, id, commandHandler, eventHandler, NewFee())
        },
        FeeCommandTypes().Literals(), FeeEventTypes().Literals(),
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus),
        FeeCommandHandler: commandHandler,
        FeeEventHandler: eventHandler,
    }

    return
}


type FeeKindCommandHandler struct {
    CreateHandler func (*CreateFeeKind, *FeeKind, eh.AggregateStoreEvent) error
    DeleteHandler func (*DeleteFeeKind, *FeeKind, eh.AggregateStoreEvent) error
    UpdateHandler func (*UpdateFeeKind, *FeeKind, eh.AggregateStoreEvent) error
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
        o.CreateHandler = func(command *CreateFeeKind, entity *FeeKind, store eh.AggregateStoreEvent) (ret error) {
            if ret = eh.ValidateNewId(entity.Id, command.Id, FeeKindAggregateType); ret == nil {store.StoreEvent(FeeKindCreatedEvent, &FeeKindCreated{
                    Id: command.Id,
                    Name: command.Name,
                    Amount: command.Amount,
                    Description: command.Description,}, time.Now())
            }
            return
        }
    }
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteFeeKind, entity *FeeKind, store eh.AggregateStoreEvent) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, command.Id, FeeKindAggregateType); ret == nil {
                store.StoreEvent(FeeKindDeletedEvent, &FeeKindDeleted{
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    if o.UpdateHandler == nil {
        o.UpdateHandler = func(command *UpdateFeeKind, entity *FeeKind, store eh.AggregateStoreEvent) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, command.Id, FeeKindAggregateType); ret == nil {
                store.StoreEvent(FeeKindUpdatedEvent, &FeeKindUpdated{
                    Id: command.Id,
                    Name: command.Name,
                    Amount: command.Amount,
                    Description: command.Description,}, time.Now())
            }
            return
        }
    }
    return
}


type FeeKindEventHandler struct {
    CreatedHandler func (*FeeKindCreated, *FeeKind) error
    DeletedHandler func (*FeeKindDeleted, *FeeKind) error
    UpdatedHandler func (*FeeKindUpdated, *FeeKind) error
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
            if ret = eh.ValidateNewId(entity.Id, event.Id, FeeKindAggregateType); ret == nil {
                entity.Id = event.Id
                entity.Name = event.Name
                entity.Amount = event.Amount
                entity.Description = event.Description
            }
            return
        }
    }
    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *FeeKindDeleted, entity *FeeKind) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, event.Id, FeeKindAggregateType); ret == nil {
                *entity = *NewFeeKind()
            }
            return
        }
    }
    if o.UpdatedHandler == nil {
        o.UpdatedHandler = func(event *FeeKindUpdated, entity *FeeKind) (ret error) {
            if ret = eh.ValidateIdsMatch(entity.Id, event.Id, FeeKindAggregateType); ret == nil {
                entity.Name = event.Name
                entity.Amount = event.Amount
                entity.Description = event.Description
            }
            return
        }
    }
    return
}


const FeeKindAggregateType eventhorizon.AggregateType = "FeeKind"

type FeeKindAggregateInitializer struct {
    *eh.AggregateInitializer
    *FeeKindCommandHandler
    *FeeKindEventHandler
}



func NewFeeKindAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus) (ret *FeeKindAggregateInitializer) {
    
    commandHandler := &FeeKindCommandHandler{}
    eventHandler := &FeeKindEventHandler{}
	ret = &FeeKindAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(FeeKindAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(FeeKindAggregateType, id, commandHandler, eventHandler, NewFeeKind())
        },
        FeeKindCommandTypes().Literals(), FeeKindEventTypes().Literals(),
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus),
        FeeKindCommandHandler: commandHandler,
        FeeKindEventHandler: eventHandler,
    }

    return
}


type FinanceEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore
    eventBus eventhorizon.EventBus
    eventPublisher eventhorizon.EventPublisher
    commandBus eventhorizon.CommandBus
    ExpenseAggregateInitializer *ExpenseAggregateInitializer
    ExpensePurposeAggregateInitializer *ExpensePurposeAggregateInitializer
    FeeAggregateInitializer *FeeAggregateInitializer
    FeeKindAggregateInitializer *FeeKindAggregateInitializer
}

func NewFinanceEventhorizonInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus) (ret *FinanceEventhorizonInitializer) {
    ret = &FinanceEventhorizonInitializer{
        eventStore: eventStore,
        eventBus: eventBus,
        eventPublisher: eventPublisher,
        commandBus: commandBus,
        ExpenseAggregateInitializer: NewExpenseAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
        ExpensePurposeAggregateInitializer: NewExpensePurposeAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
        FeeAggregateInitializer: NewFeeAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
        FeeKindAggregateInitializer: NewFeeKindAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    }
    return
}

func (o *FinanceEventhorizonInitializer) Setup() (ret error) {
    
    if ret = o.ExpenseAggregateInitializer.Setup(); ret != nil {
        return
    }
    
    if ret = o.ExpensePurposeAggregateInitializer.Setup(); ret != nil {
        return
    }
    
    if ret = o.FeeAggregateInitializer.Setup(); ret != nil {
        return
    }
    
    if ret = o.FeeKindAggregateInitializer.Setup(); ret != nil {
        return
    }

    return
}









