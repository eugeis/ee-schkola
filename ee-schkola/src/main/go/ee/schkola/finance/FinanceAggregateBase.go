package finance

import (
    "errors"
    "fmt"
    "github.com/eugeis/gee/eh"
    "github.com/looplab/eventhorizon"
    "time"
)
type ExpenseCommandHandler struct {
    CreateHandler func (*CreateExpense, *Expense, eh.AggregateStoreEvent) (err error) 
    DeleteHandler func (*DeleteExpense, *Expense, eh.AggregateStoreEvent) (err error) 
    UpdateHandler func (*UpdateExpense, *Expense, eh.AggregateStoreEvent) (err error) 
}

func (o *ExpenseCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateExpenseCommand:
        err = o.CreateHandler(cmd.(*CreateExpense), entity.(*Expense), store)
    case DeleteExpenseCommand:
        err = o.DeleteHandler(cmd.(*DeleteExpense), entity.(*Expense), store)
    case UpdateExpenseCommand:
        err = o.UpdateHandler(cmd.(*UpdateExpense), entity.(*Expense), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *ExpenseCommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *CreateExpense, entity *Expense, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, ExpenseAggregateType); err == nil {
            store.StoreEvent(ExpenseCreatedEvent, &ExpenseCreated{
                Purpose: command.Purpose,
                Amount: command.Amount,
                Profile: command.Profile,
                Date: command.Date,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *DeleteExpense, entity *Expense, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, ExpenseAggregateType); err == nil {
            store.StoreEvent(ExpenseDeletedEvent, &ExpenseDeleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *UpdateExpense, entity *Expense, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, ExpenseAggregateType); err == nil {
            store.StoreEvent(ExpenseUpdatedEvent, &ExpenseUpdated{
                Purpose: command.Purpose,
                Amount: command.Amount,
                Profile: command.Profile,
                Date: command.Date,
                Id: command.Id,}, time.Now())
        }
        return
    }
    return
}


type ExpenseEventHandler struct {
    CreatedHandler func (*ExpenseCreated, *Expense) (err error) 
    DeletedHandler func (*ExpenseDeleted, *Expense) (err error) 
    UpdatedHandler func (*ExpenseUpdated, *Expense) (err error) 
}

func (o *ExpenseEventHandler) Apply(event eventhorizon.Event, entity interface{}) (err error) {
    switch event.EventType() {
    case ExpenseCreatedEvent:
        err = o.CreatedHandler(event.Data().(*ExpenseCreated), entity.(*Expense))
    case ExpenseDeletedEvent:
        err = o.DeletedHandler(event.Data().(*ExpenseDeleted), entity.(*Expense))
    case ExpenseUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*ExpenseUpdated), entity.(*Expense))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *ExpenseEventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(ExpenseCreatedEvent, func() eventhorizon.EventData {
		return &ExpenseCreated{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *ExpenseCreated, entity *Expense) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, ExpenseAggregateType); err == nil {
            entity.Purpose = event.Purpose
            entity.Amount = event.Amount
            entity.Profile = event.Profile
            entity.Date = event.Date
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(ExpenseDeletedEvent, func() eventhorizon.EventData {
		return &ExpenseDeleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *ExpenseDeleted, entity *Expense) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, ExpenseAggregateType); err == nil {
            *entity = *NewExpense()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(ExpenseUpdatedEvent, func() eventhorizon.EventData {
		return &ExpenseUpdated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *ExpenseUpdated, entity *Expense) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, ExpenseAggregateType); err == nil {
            entity.Purpose = event.Purpose
            entity.Amount = event.Amount
            entity.Profile = event.Profile
            entity.Date = event.Date
        }
        return
    }
    return
}


const ExpenseAggregateType eventhorizon.AggregateType = "Expense"

type ExpenseAggregateInitializer struct {
    *eh.AggregateInitializer
    *ExpenseCommandHandler
    *ExpenseEventHandler
    ProjectorHandler *ExpenseEventHandler
}



func NewExpenseAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus, 
                readRepos func (string, func () (ret interface{}) ) (ret eventhorizon.ReadWriteRepo) ) (ret *ExpenseAggregateInitializer) {
    
    commandHandler := &ExpenseCommandHandler{}
    eventHandler := &ExpenseEventHandler{}
    modelFactory := func() interface{} { return NewExpense() }
    ret = &ExpenseAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ExpenseAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(ExpenseAggregateType, id, commandHandler, eventHandler, modelFactory())
        }, modelFactory,
        ExpenseCommandTypes().Literals(), ExpenseEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), ExpenseCommandHandler: commandHandler, ExpenseEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type ExpensePurposeCommandHandler struct {
    CreateHandler func (*CreateExpensePurpose, *ExpensePurpose, eh.AggregateStoreEvent) (err error) 
    DeleteHandler func (*DeleteExpensePurpose, *ExpensePurpose, eh.AggregateStoreEvent) (err error) 
    UpdateHandler func (*UpdateExpensePurpose, *ExpensePurpose, eh.AggregateStoreEvent) (err error) 
}

func (o *ExpensePurposeCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateExpensePurposeCommand:
        err = o.CreateHandler(cmd.(*CreateExpensePurpose), entity.(*ExpensePurpose), store)
    case DeleteExpensePurposeCommand:
        err = o.DeleteHandler(cmd.(*DeleteExpensePurpose), entity.(*ExpensePurpose), store)
    case UpdateExpensePurposeCommand:
        err = o.UpdateHandler(cmd.(*UpdateExpensePurpose), entity.(*ExpensePurpose), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *ExpensePurposeCommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *CreateExpensePurpose, entity *ExpensePurpose, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, ExpensePurposeAggregateType); err == nil {
            store.StoreEvent(ExpensePurposeCreatedEvent, &ExpensePurposeCreated{
                Name: command.Name,
                Description: command.Description,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *DeleteExpensePurpose, entity *ExpensePurpose, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, ExpensePurposeAggregateType); err == nil {
            store.StoreEvent(ExpensePurposeDeletedEvent, &ExpensePurposeDeleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *UpdateExpensePurpose, entity *ExpensePurpose, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, ExpensePurposeAggregateType); err == nil {
            store.StoreEvent(ExpensePurposeUpdatedEvent, &ExpensePurposeUpdated{
                Name: command.Name,
                Description: command.Description,
                Id: command.Id,}, time.Now())
        }
        return
    }
    return
}


type ExpensePurposeEventHandler struct {
    CreatedHandler func (*ExpensePurposeCreated, *ExpensePurpose) (err error) 
    DeletedHandler func (*ExpensePurposeDeleted, *ExpensePurpose) (err error) 
    UpdatedHandler func (*ExpensePurposeUpdated, *ExpensePurpose) (err error) 
}

func (o *ExpensePurposeEventHandler) Apply(event eventhorizon.Event, entity interface{}) (err error) {
    switch event.EventType() {
    case ExpensePurposeCreatedEvent:
        err = o.CreatedHandler(event.Data().(*ExpensePurposeCreated), entity.(*ExpensePurpose))
    case ExpensePurposeDeletedEvent:
        err = o.DeletedHandler(event.Data().(*ExpensePurposeDeleted), entity.(*ExpensePurpose))
    case ExpensePurposeUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*ExpensePurposeUpdated), entity.(*ExpensePurpose))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *ExpensePurposeEventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(ExpensePurposeCreatedEvent, func() eventhorizon.EventData {
		return &ExpensePurposeCreated{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *ExpensePurposeCreated, entity *ExpensePurpose) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, ExpensePurposeAggregateType); err == nil {
            entity.Name = event.Name
            entity.Description = event.Description
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(ExpensePurposeDeletedEvent, func() eventhorizon.EventData {
		return &ExpensePurposeDeleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *ExpensePurposeDeleted, entity *ExpensePurpose) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, ExpensePurposeAggregateType); err == nil {
            *entity = *NewExpensePurpose()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(ExpensePurposeUpdatedEvent, func() eventhorizon.EventData {
		return &ExpensePurposeUpdated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *ExpensePurposeUpdated, entity *ExpensePurpose) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, ExpensePurposeAggregateType); err == nil {
            entity.Name = event.Name
            entity.Description = event.Description
        }
        return
    }
    return
}


const ExpensePurposeAggregateType eventhorizon.AggregateType = "ExpensePurpose"

type ExpensePurposeAggregateInitializer struct {
    *eh.AggregateInitializer
    *ExpensePurposeCommandHandler
    *ExpensePurposeEventHandler
    ProjectorHandler *ExpensePurposeEventHandler
}



func NewExpensePurposeAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus, 
                readRepos func (string, func () (ret interface{}) ) (ret eventhorizon.ReadWriteRepo) ) (ret *ExpensePurposeAggregateInitializer) {
    
    commandHandler := &ExpensePurposeCommandHandler{}
    eventHandler := &ExpensePurposeEventHandler{}
    modelFactory := func() interface{} { return NewExpensePurpose() }
    ret = &ExpensePurposeAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ExpensePurposeAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(ExpensePurposeAggregateType, id, commandHandler, eventHandler, modelFactory())
        }, modelFactory,
        ExpensePurposeCommandTypes().Literals(), ExpensePurposeEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), ExpensePurposeCommandHandler: commandHandler, ExpensePurposeEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type FeeCommandHandler struct {
    CreateHandler func (*CreateFee, *Fee, eh.AggregateStoreEvent) (err error) 
    DeleteHandler func (*DeleteFee, *Fee, eh.AggregateStoreEvent) (err error) 
    UpdateHandler func (*UpdateFee, *Fee, eh.AggregateStoreEvent) (err error) 
}

func (o *FeeCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateFeeCommand:
        err = o.CreateHandler(cmd.(*CreateFee), entity.(*Fee), store)
    case DeleteFeeCommand:
        err = o.DeleteHandler(cmd.(*DeleteFee), entity.(*Fee), store)
    case UpdateFeeCommand:
        err = o.UpdateHandler(cmd.(*UpdateFee), entity.(*Fee), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *FeeCommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *CreateFee, entity *Fee, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, FeeAggregateType); err == nil {
            store.StoreEvent(FeeCreatedEvent, &FeeCreated{
                Student: command.Student,
                Amount: command.Amount,
                Kind: command.Kind,
                Date: command.Date,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *DeleteFee, entity *Fee, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, FeeAggregateType); err == nil {
            store.StoreEvent(FeeDeletedEvent, &FeeDeleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *UpdateFee, entity *Fee, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, FeeAggregateType); err == nil {
            store.StoreEvent(FeeUpdatedEvent, &FeeUpdated{
                Student: command.Student,
                Amount: command.Amount,
                Kind: command.Kind,
                Date: command.Date,
                Id: command.Id,}, time.Now())
        }
        return
    }
    return
}


type FeeEventHandler struct {
    CreatedHandler func (*FeeCreated, *Fee) (err error) 
    DeletedHandler func (*FeeDeleted, *Fee) (err error) 
    UpdatedHandler func (*FeeUpdated, *Fee) (err error) 
}

func (o *FeeEventHandler) Apply(event eventhorizon.Event, entity interface{}) (err error) {
    switch event.EventType() {
    case FeeCreatedEvent:
        err = o.CreatedHandler(event.Data().(*FeeCreated), entity.(*Fee))
    case FeeDeletedEvent:
        err = o.DeletedHandler(event.Data().(*FeeDeleted), entity.(*Fee))
    case FeeUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*FeeUpdated), entity.(*Fee))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *FeeEventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(FeeCreatedEvent, func() eventhorizon.EventData {
		return &FeeCreated{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *FeeCreated, entity *Fee) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, FeeAggregateType); err == nil {
            entity.Student = event.Student
            entity.Amount = event.Amount
            entity.Kind = event.Kind
            entity.Date = event.Date
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(FeeDeletedEvent, func() eventhorizon.EventData {
		return &FeeDeleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *FeeDeleted, entity *Fee) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, FeeAggregateType); err == nil {
            *entity = *NewFee()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(FeeUpdatedEvent, func() eventhorizon.EventData {
		return &FeeUpdated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *FeeUpdated, entity *Fee) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, FeeAggregateType); err == nil {
            entity.Student = event.Student
            entity.Amount = event.Amount
            entity.Kind = event.Kind
            entity.Date = event.Date
        }
        return
    }
    return
}


const FeeAggregateType eventhorizon.AggregateType = "Fee"

type FeeAggregateInitializer struct {
    *eh.AggregateInitializer
    *FeeCommandHandler
    *FeeEventHandler
    ProjectorHandler *FeeEventHandler
}



func NewFeeAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus, 
                readRepos func (string, func () (ret interface{}) ) (ret eventhorizon.ReadWriteRepo) ) (ret *FeeAggregateInitializer) {
    
    commandHandler := &FeeCommandHandler{}
    eventHandler := &FeeEventHandler{}
    modelFactory := func() interface{} { return NewFee() }
    ret = &FeeAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(FeeAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(FeeAggregateType, id, commandHandler, eventHandler, modelFactory())
        }, modelFactory,
        FeeCommandTypes().Literals(), FeeEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), FeeCommandHandler: commandHandler, FeeEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type FeeKindCommandHandler struct {
    CreateHandler func (*CreateFeeKind, *FeeKind, eh.AggregateStoreEvent) (err error) 
    DeleteHandler func (*DeleteFeeKind, *FeeKind, eh.AggregateStoreEvent) (err error) 
    UpdateHandler func (*UpdateFeeKind, *FeeKind, eh.AggregateStoreEvent) (err error) 
}

func (o *FeeKindCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateFeeKindCommand:
        err = o.CreateHandler(cmd.(*CreateFeeKind), entity.(*FeeKind), store)
    case DeleteFeeKindCommand:
        err = o.DeleteHandler(cmd.(*DeleteFeeKind), entity.(*FeeKind), store)
    case UpdateFeeKindCommand:
        err = o.UpdateHandler(cmd.(*UpdateFeeKind), entity.(*FeeKind), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *FeeKindCommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *CreateFeeKind, entity *FeeKind, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, FeeKindAggregateType); err == nil {
            store.StoreEvent(FeeKindCreatedEvent, &FeeKindCreated{
                Name: command.Name,
                Amount: command.Amount,
                Description: command.Description,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *DeleteFeeKind, entity *FeeKind, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, FeeKindAggregateType); err == nil {
            store.StoreEvent(FeeKindDeletedEvent, &FeeKindDeleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *UpdateFeeKind, entity *FeeKind, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, FeeKindAggregateType); err == nil {
            store.StoreEvent(FeeKindUpdatedEvent, &FeeKindUpdated{
                Name: command.Name,
                Amount: command.Amount,
                Description: command.Description,
                Id: command.Id,}, time.Now())
        }
        return
    }
    return
}


type FeeKindEventHandler struct {
    CreatedHandler func (*FeeKindCreated, *FeeKind) (err error) 
    DeletedHandler func (*FeeKindDeleted, *FeeKind) (err error) 
    UpdatedHandler func (*FeeKindUpdated, *FeeKind) (err error) 
}

func (o *FeeKindEventHandler) Apply(event eventhorizon.Event, entity interface{}) (err error) {
    switch event.EventType() {
    case FeeKindCreatedEvent:
        err = o.CreatedHandler(event.Data().(*FeeKindCreated), entity.(*FeeKind))
    case FeeKindDeletedEvent:
        err = o.DeletedHandler(event.Data().(*FeeKindDeleted), entity.(*FeeKind))
    case FeeKindUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*FeeKindUpdated), entity.(*FeeKind))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *FeeKindEventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(FeeKindCreatedEvent, func() eventhorizon.EventData {
		return &FeeKindCreated{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *FeeKindCreated, entity *FeeKind) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, FeeKindAggregateType); err == nil {
            entity.Name = event.Name
            entity.Amount = event.Amount
            entity.Description = event.Description
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(FeeKindDeletedEvent, func() eventhorizon.EventData {
		return &FeeKindDeleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *FeeKindDeleted, entity *FeeKind) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, FeeKindAggregateType); err == nil {
            *entity = *NewFeeKind()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(FeeKindUpdatedEvent, func() eventhorizon.EventData {
		return &FeeKindUpdated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *FeeKindUpdated, entity *FeeKind) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, FeeKindAggregateType); err == nil {
            entity.Name = event.Name
            entity.Amount = event.Amount
            entity.Description = event.Description
        }
        return
    }
    return
}


const FeeKindAggregateType eventhorizon.AggregateType = "FeeKind"

type FeeKindAggregateInitializer struct {
    *eh.AggregateInitializer
    *FeeKindCommandHandler
    *FeeKindEventHandler
    ProjectorHandler *FeeKindEventHandler
}



func NewFeeKindAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus, 
                readRepos func (string, func () (ret interface{}) ) (ret eventhorizon.ReadWriteRepo) ) (ret *FeeKindAggregateInitializer) {
    
    commandHandler := &FeeKindCommandHandler{}
    eventHandler := &FeeKindEventHandler{}
    modelFactory := func() interface{} { return NewFeeKind() }
    ret = &FeeKindAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(FeeKindAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(FeeKindAggregateType, id, commandHandler, eventHandler, modelFactory())
        }, modelFactory,
        FeeKindCommandTypes().Literals(), FeeKindEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), FeeKindCommandHandler: commandHandler, FeeKindEventHandler: eventHandler, ProjectorHandler: eventHandler,
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
                commandBus eventhorizon.CommandBus, 
                readRepos func (string, func () (ret interface{}) ) (ret eventhorizon.ReadWriteRepo) ) (ret *FinanceEventhorizonInitializer) {
    expenseAggregateInitializer := NewExpenseAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    expensePurposeAggregateInitializer := NewExpensePurposeAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    feeAggregateInitializer := NewFeeAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    feeKindAggregateInitializer := NewFeeKindAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    ret = &FinanceEventhorizonInitializer{
        eventStore: eventStore,
        eventBus: eventBus,
        eventPublisher: eventPublisher,
        commandBus: commandBus,
        ExpenseAggregateInitializer: expenseAggregateInitializer,
        ExpensePurposeAggregateInitializer: expensePurposeAggregateInitializer,
        FeeAggregateInitializer: feeAggregateInitializer,
        FeeKindAggregateInitializer: feeKindAggregateInitializer,
    }
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









