package finance

import (
    "context"
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

const ExpenseAggregateType eventhorizon.AggregateType = "ExpenseAggregate"

func NewExpenseAggregate(id eventhorizon.UUID) *ExpenseAggregate {
	return &ExpenseAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(ExpenseAggregateType, id),
	}
}

func (o *ExpenseAggregate) HandleCommand(ctx context.Context, cmd eventhorizon.Command) error {
    println("HandleCommand", cmd.CommandType())
    return nil
}

func (o *ExpenseAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type ExpenseAggregate struct {
    *eventhorizon.AggregateBase
    *Expense
}



func NewExpenseAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *ExpenseAggregateInitializer) {
	ret = &ExpenseAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ExpenseAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate { return NewExpenseAggregate(id) },
        ExpenseCommandTypes().Literals(), ExpenseEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
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
}



const ExpensePurposeAggregateType eventhorizon.AggregateType = "ExpensePurposeAggregate"

func NewExpensePurposeAggregate(id eventhorizon.UUID) *ExpensePurposeAggregate {
	return &ExpensePurposeAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(ExpensePurposeAggregateType, id),
	}
}

func (o *ExpensePurposeAggregate) HandleCommand(ctx context.Context, cmd eventhorizon.Command) error {
    println("HandleCommand", cmd.CommandType())
    return nil
}

func (o *ExpensePurposeAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type ExpensePurposeAggregate struct {
    *eventhorizon.AggregateBase
    *ExpensePurpose
}



func NewExpensePurposeAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *ExpensePurposeAggregateInitializer) {
	ret = &ExpensePurposeAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ExpensePurposeAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate { return NewExpensePurposeAggregate(id) },
        ExpensePurposeCommandTypes().Literals(), ExpensePurposeEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
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
}



const FeeAggregateType eventhorizon.AggregateType = "FeeAggregate"

func NewFeeAggregate(id eventhorizon.UUID) *FeeAggregate {
	return &FeeAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(FeeAggregateType, id),
	}
}

func (o *FeeAggregate) HandleCommand(ctx context.Context, cmd eventhorizon.Command) error {
    println("HandleCommand", cmd.CommandType())
    return nil
}

func (o *FeeAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type FeeAggregate struct {
    *eventhorizon.AggregateBase
    *Fee
}



func NewFeeAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *FeeAggregateInitializer) {
	ret = &FeeAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(FeeAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate { return NewFeeAggregate(id) },
        FeeCommandTypes().Literals(), FeeEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
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
}



const FeeKindAggregateType eventhorizon.AggregateType = "FeeKindAggregate"

func NewFeeKindAggregate(id eventhorizon.UUID) *FeeKindAggregate {
	return &FeeKindAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(FeeKindAggregateType, id),
	}
}

func (o *FeeKindAggregate) HandleCommand(ctx context.Context, cmd eventhorizon.Command) error {
    println("HandleCommand", cmd.CommandType())
    return nil
}

func (o *FeeKindAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type FeeKindAggregate struct {
    *eventhorizon.AggregateBase
    *FeeKind
}



func NewFeeKindAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *FeeKindAggregateInitializer) {
	ret = &FeeKindAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(FeeKindAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate { return NewFeeKindAggregate(id) },
        FeeKindCommandTypes().Literals(), FeeKindEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
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
}



func NewFinanceEventhorizonInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *FinanceEventhorizonInitializer) {
	ret = &FinanceEventhorizonInitializer{eventStore: eventStore, eventBus: eventBus, eventPublisher: eventPublisher,
            commandBus: commandBus, 
    expenseAggregateInitializer: NewExpenseAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    expensePurposeAggregateInitializer: NewExpensePurposeAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    feeAggregateInitializer: NewFeeAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    feeKindAggregateInitializer: NewFeeKindAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus)}
	return
}

func (o *FinanceEventhorizonInitializer) Setup() (err error) {
    
    if err = o.expenseAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.expensePurposeAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.feeAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.feeKindAggregateInitializer.Setup(); err != nil {
        return
    }
    return
}

type FinanceEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore
    eventBus eventhorizon.EventBus
    eventPublisher eventhorizon.EventPublisher
    commandBus eventhorizon.CommandBus
    expenseAggregateInitializer *ExpenseAggregateInitializer
    expensePurposeAggregateInitializer *ExpensePurposeAggregateInitializer
    feeAggregateInitializer *FeeAggregateInitializer
    feeKindAggregateInitializer *FeeKindAggregateInitializer
}









