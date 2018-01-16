package finance

import (
    "errors"
    "fmt"
    "github.com/eugeis/gee/eh"
    "github.com/looplab/eventhorizon"
    "github.com/looplab/eventhorizon/commandhandler/bus"
    "time"
)
type ExpenseCommandHandler struct {
    CreateHandler func (*Create, *Expense, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*Delete, *Expense, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*Update, *Expense, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *ExpenseCommandHandler) AddCreatePreparer(preparer func (*Create, *Expense) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *Create, entity *Expense, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *ExpenseCommandHandler) AddDeletePreparer(preparer func (*Delete, *Expense) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *Delete, entity *Expense, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *ExpenseCommandHandler) AddUpdatePreparer(preparer func (*Update, *Expense) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *Update, entity *Expense, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *ExpenseCommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateExpenseCommand:
        err = o.CreateHandler(cmd.(*Create), entity.(*Expense), store)
    case DeleteExpenseCommand:
        err = o.DeleteHandler(cmd.(*Delete), entity.(*Expense), store)
    case UpdateExpenseCommand:
        err = o.UpdateHandler(cmd.(*Update), entity.(*Expense), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *ExpenseCommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *Create, entity *Expense, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, ExpenseAggregateType); err == nil {
            store.StoreEvent(ExpenseCreatedEvent, &Created{
                Purpose: command.Purpose,
                Amount: command.Amount,
                Profile: command.Profile,
                Date: command.Date,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *Delete, entity *Expense, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, ExpenseAggregateType); err == nil {
            store.StoreEvent(ExpenseDeletedEvent, &Deleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *Update, entity *Expense, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, ExpenseAggregateType); err == nil {
            store.StoreEvent(ExpenseUpdatedEvent, &Updated{
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
    CreatedHandler func (*Created, *Expense) (err error)  `json:"createdHandler" eh:"optional"`
    DeletedHandler func (*Deleted, *Expense) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdatedHandler func (*Updated, *Expense) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *ExpenseEventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case ExpenseCreatedEvent:
        err = o.CreatedHandler(event.Data().(*Created), entity.(*Expense))
    case ExpenseDeletedEvent:
        err = o.DeletedHandler(event.Data().(*Deleted), entity.(*Expense))
    case ExpenseUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*Updated), entity.(*Expense))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *ExpenseEventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(CreatedEvent, func() eventhorizon.EventData {
		return &Created{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *Created, entity *Expense) (err error) {
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
    eventhorizon.RegisterEventData(DeletedEvent, func() eventhorizon.EventData {
		return &Deleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *Deleted, entity *Expense) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, ExpenseAggregateType); err == nil {
            *entity = *New@@EMPTY@@()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdatedEvent, func() eventhorizon.EventData {
		return &Updated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *Updated, entity *Expense) (err error) {
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
    ProjectorHandler *ExpenseEventHandler `json:"projectorHandler" eh:"optional"`
}



func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *ExpenseAggregateInitializer) {
    
    commandHandler := &ExpenseCommandHandler{}
    eventHandler := &ExpenseEventHandler{}
    entityFactory := func() eventhorizon.Entity { return New@@EMPTY@@() }
    ret = &ExpenseAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ExpenseAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(ExpenseAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        ExpenseCommandTypes().Literals(), ExpenseEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), ExpenseCommandHandler: commandHandler, ExpenseEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type ExpensePurposeCommandHandler struct {
    CreateHandler func (*Create, *ExpensePurpose, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*Delete, *ExpensePurpose, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*Update, *ExpensePurpose, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *ExpensePurposeCommandHandler) AddCreatePreparer(preparer func (*Create, *ExpensePurpose) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *Create, entity *ExpensePurpose, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *ExpensePurposeCommandHandler) AddDeletePreparer(preparer func (*Delete, *ExpensePurpose) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *Delete, entity *ExpensePurpose, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *ExpensePurposeCommandHandler) AddUpdatePreparer(preparer func (*Update, *ExpensePurpose) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *Update, entity *ExpensePurpose, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *ExpensePurposeCommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateExpensePurposeCommand:
        err = o.CreateHandler(cmd.(*Create), entity.(*ExpensePurpose), store)
    case DeleteExpensePurposeCommand:
        err = o.DeleteHandler(cmd.(*Delete), entity.(*ExpensePurpose), store)
    case UpdateExpensePurposeCommand:
        err = o.UpdateHandler(cmd.(*Update), entity.(*ExpensePurpose), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *ExpensePurposeCommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *Create, entity *ExpensePurpose, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, ExpensePurposeAggregateType); err == nil {
            store.StoreEvent(ExpensePurposeCreatedEvent, &Created{
                Name: command.Name,
                Description: command.Description,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *Delete, entity *ExpensePurpose, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, ExpensePurposeAggregateType); err == nil {
            store.StoreEvent(ExpensePurposeDeletedEvent, &Deleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *Update, entity *ExpensePurpose, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, ExpensePurposeAggregateType); err == nil {
            store.StoreEvent(ExpensePurposeUpdatedEvent, &Updated{
                Name: command.Name,
                Description: command.Description,
                Id: command.Id,}, time.Now())
        }
        return
    }
    return
}


type ExpensePurposeEventHandler struct {
    CreatedHandler func (*Created, *ExpensePurpose) (err error)  `json:"createdHandler" eh:"optional"`
    DeletedHandler func (*Deleted, *ExpensePurpose) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdatedHandler func (*Updated, *ExpensePurpose) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *ExpensePurposeEventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case ExpensePurposeCreatedEvent:
        err = o.CreatedHandler(event.Data().(*Created), entity.(*ExpensePurpose))
    case ExpensePurposeDeletedEvent:
        err = o.DeletedHandler(event.Data().(*Deleted), entity.(*ExpensePurpose))
    case ExpensePurposeUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*Updated), entity.(*ExpensePurpose))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *ExpensePurposeEventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(CreatedEvent, func() eventhorizon.EventData {
		return &Created{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *Created, entity *ExpensePurpose) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, ExpensePurposeAggregateType); err == nil {
            entity.Name = event.Name
            entity.Description = event.Description
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DeletedEvent, func() eventhorizon.EventData {
		return &Deleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *Deleted, entity *ExpensePurpose) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, ExpensePurposeAggregateType); err == nil {
            *entity = *New@@EMPTY@@()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdatedEvent, func() eventhorizon.EventData {
		return &Updated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *Updated, entity *ExpensePurpose) (err error) {
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
    ProjectorHandler *ExpensePurposeEventHandler `json:"projectorHandler" eh:"optional"`
}



func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *ExpensePurposeAggregateInitializer) {
    
    commandHandler := &ExpensePurposeCommandHandler{}
    eventHandler := &ExpensePurposeEventHandler{}
    entityFactory := func() eventhorizon.Entity { return New@@EMPTY@@() }
    ret = &ExpensePurposeAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ExpensePurposeAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(ExpensePurposeAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        ExpensePurposeCommandTypes().Literals(), ExpensePurposeEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), ExpensePurposeCommandHandler: commandHandler, ExpensePurposeEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type FeeCommandHandler struct {
    CreateHandler func (*Create, *Fee, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*Delete, *Fee, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*Update, *Fee, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *FeeCommandHandler) AddCreatePreparer(preparer func (*Create, *Fee) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *Create, entity *Fee, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *FeeCommandHandler) AddDeletePreparer(preparer func (*Delete, *Fee) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *Delete, entity *Fee, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *FeeCommandHandler) AddUpdatePreparer(preparer func (*Update, *Fee) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *Update, entity *Fee, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *FeeCommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateFeeCommand:
        err = o.CreateHandler(cmd.(*Create), entity.(*Fee), store)
    case DeleteFeeCommand:
        err = o.DeleteHandler(cmd.(*Delete), entity.(*Fee), store)
    case UpdateFeeCommand:
        err = o.UpdateHandler(cmd.(*Update), entity.(*Fee), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *FeeCommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *Create, entity *Fee, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, FeeAggregateType); err == nil {
            store.StoreEvent(FeeCreatedEvent, &Created{
                Student: command.Student,
                Amount: command.Amount,
                Kind: command.Kind,
                Date: command.Date,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *Delete, entity *Fee, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, FeeAggregateType); err == nil {
            store.StoreEvent(FeeDeletedEvent, &Deleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *Update, entity *Fee, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, FeeAggregateType); err == nil {
            store.StoreEvent(FeeUpdatedEvent, &Updated{
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
    CreatedHandler func (*Created, *Fee) (err error)  `json:"createdHandler" eh:"optional"`
    DeletedHandler func (*Deleted, *Fee) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdatedHandler func (*Updated, *Fee) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *FeeEventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case FeeCreatedEvent:
        err = o.CreatedHandler(event.Data().(*Created), entity.(*Fee))
    case FeeDeletedEvent:
        err = o.DeletedHandler(event.Data().(*Deleted), entity.(*Fee))
    case FeeUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*Updated), entity.(*Fee))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *FeeEventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(CreatedEvent, func() eventhorizon.EventData {
		return &Created{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *Created, entity *Fee) (err error) {
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
    eventhorizon.RegisterEventData(DeletedEvent, func() eventhorizon.EventData {
		return &Deleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *Deleted, entity *Fee) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, FeeAggregateType); err == nil {
            *entity = *New@@EMPTY@@()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdatedEvent, func() eventhorizon.EventData {
		return &Updated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *Updated, entity *Fee) (err error) {
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
    ProjectorHandler *FeeEventHandler `json:"projectorHandler" eh:"optional"`
}



func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *FeeAggregateInitializer) {
    
    commandHandler := &FeeCommandHandler{}
    eventHandler := &FeeEventHandler{}
    entityFactory := func() eventhorizon.Entity { return New@@EMPTY@@() }
    ret = &FeeAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(FeeAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(FeeAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        FeeCommandTypes().Literals(), FeeEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), FeeCommandHandler: commandHandler, FeeEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type FeeKindCommandHandler struct {
    CreateHandler func (*Create, *FeeKind, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*Delete, *FeeKind, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*Update, *FeeKind, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *FeeKindCommandHandler) AddCreatePreparer(preparer func (*Create, *FeeKind) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *Create, entity *FeeKind, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *FeeKindCommandHandler) AddDeletePreparer(preparer func (*Delete, *FeeKind) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *Delete, entity *FeeKind, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *FeeKindCommandHandler) AddUpdatePreparer(preparer func (*Update, *FeeKind) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *Update, entity *FeeKind, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *FeeKindCommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateFeeKindCommand:
        err = o.CreateHandler(cmd.(*Create), entity.(*FeeKind), store)
    case DeleteFeeKindCommand:
        err = o.DeleteHandler(cmd.(*Delete), entity.(*FeeKind), store)
    case UpdateFeeKindCommand:
        err = o.UpdateHandler(cmd.(*Update), entity.(*FeeKind), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *FeeKindCommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *Create, entity *FeeKind, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, FeeKindAggregateType); err == nil {
            store.StoreEvent(FeeKindCreatedEvent, &Created{
                Name: command.Name,
                Amount: command.Amount,
                Description: command.Description,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *Delete, entity *FeeKind, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, FeeKindAggregateType); err == nil {
            store.StoreEvent(FeeKindDeletedEvent, &Deleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *Update, entity *FeeKind, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, FeeKindAggregateType); err == nil {
            store.StoreEvent(FeeKindUpdatedEvent, &Updated{
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
    CreatedHandler func (*Created, *FeeKind) (err error)  `json:"createdHandler" eh:"optional"`
    DeletedHandler func (*Deleted, *FeeKind) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdatedHandler func (*Updated, *FeeKind) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *FeeKindEventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case FeeKindCreatedEvent:
        err = o.CreatedHandler(event.Data().(*Created), entity.(*FeeKind))
    case FeeKindDeletedEvent:
        err = o.DeletedHandler(event.Data().(*Deleted), entity.(*FeeKind))
    case FeeKindUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*Updated), entity.(*FeeKind))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *FeeKindEventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(CreatedEvent, func() eventhorizon.EventData {
		return &Created{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *Created, entity *FeeKind) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, FeeKindAggregateType); err == nil {
            entity.Name = event.Name
            entity.Amount = event.Amount
            entity.Description = event.Description
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DeletedEvent, func() eventhorizon.EventData {
		return &Deleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *Deleted, entity *FeeKind) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, FeeKindAggregateType); err == nil {
            *entity = *New@@EMPTY@@()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdatedEvent, func() eventhorizon.EventData {
		return &Updated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *Updated, entity *FeeKind) (err error) {
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
    ProjectorHandler *FeeKindEventHandler `json:"projectorHandler" eh:"optional"`
}



func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *FeeKindAggregateInitializer) {
    
    commandHandler := &FeeKindCommandHandler{}
    eventHandler := &FeeKindEventHandler{}
    entityFactory := func() eventhorizon.Entity { return New@@EMPTY@@() }
    ret = &FeeKindAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(FeeKindAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(FeeKindAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        FeeKindCommandTypes().Literals(), FeeKindEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), FeeKindCommandHandler: commandHandler, FeeKindEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type FinanceEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore `json:"eventStore" eh:"optional"`
    eventBus eventhorizon.EventBus `json:"eventBus" eh:"optional"`
    eventPublisher eventhorizon.EventPublisher `json:"eventPublisher" eh:"optional"`
    commandBus *bus.CommandHandler `json:"commandBus" eh:"optional"`
    ExpenseAggregateInitializer *ExpenseAggregateInitializer `json:"expenseAggregateInitializer" eh:"optional"`
    ExpensePurposeAggregateInitializer *ExpensePurposeAggregateInitializer `json:"expensePurposeAggregateInitializer" eh:"optional"`
    FeeAggregateInitializer *FeeAggregateInitializer `json:"feeAggregateInitializer" eh:"optional"`
    FeeKindAggregateInitializer *FeeKindAggregateInitializer `json:"feeKindAggregateInitializer" eh:"optional"`
}

func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *FinanceEventhorizonInitializer) {
    expenseAggregateInitializer := New@@EMPTY@@(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    expensePurposeAggregateInitializer := New@@EMPTY@@(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    feeAggregateInitializer := New@@EMPTY@@(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    feeKindAggregateInitializer := New@@EMPTY@@(eventStore, eventBus, eventPublisher, commandBus, readRepos)
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









