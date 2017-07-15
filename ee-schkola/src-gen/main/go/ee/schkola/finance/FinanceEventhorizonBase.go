package finance

import (
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

const ExpenseAggregateType eventhorizon.AggregateType = "ExpenseAggregate"

type ExpenseAggregate struct {
    *eventhorizon.AggregateBase
    *Expense
}

func NewExpenseAggregate(AggregateBase *eventhorizon.AggregateBase, Entity *Expense) (ret *ExpenseAggregate, err error) {
    ret = &ExpenseAggregate{
        AggregateBase: AggregateBase,
        Expense: Entity,
    }
    
    return
}



func NewExpenseAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *ExpenseAggregateInitializer) {
	ret = &ExpenseAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ExpenseAggregateType,
        ExpenseCommandTypes().Literals(), ExpenseEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
    }
	return
}


func (o *ExpenseAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ExpenseEventTypes().CreatedExpense())
}

func (o *ExpenseAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ExpenseEventTypes().DeletedExpense())
}

func (o *ExpenseAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ExpenseEventTypes().UpdatedExpense())
}

type ExpenseAggregateInitializer struct {
    *eh.AggregateInitializer
}



const ExpensePurposeAggregateType eventhorizon.AggregateType = "ExpensePurposeAggregate"

type ExpensePurposeAggregate struct {
    *eventhorizon.AggregateBase
    *ExpensePurpose
}

func NewExpensePurposeAggregate(AggregateBase *eventhorizon.AggregateBase, Entity *ExpensePurpose) (ret *ExpensePurposeAggregate, err error) {
    ret = &ExpensePurposeAggregate{
        AggregateBase: AggregateBase,
        ExpensePurpose: Entity,
    }
    
    return
}



func NewExpensePurposeAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *ExpensePurposeAggregateInitializer) {
	ret = &ExpensePurposeAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ExpensePurposeAggregateType,
        ExpensePurposeCommandTypes().Literals(), ExpensePurposeEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
    }
	return
}


func (o *ExpensePurposeAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ExpensePurposeEventTypes().CreatedExpensePurpose())
}

func (o *ExpensePurposeAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ExpensePurposeEventTypes().DeletedExpensePurpose())
}

func (o *ExpensePurposeAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ExpensePurposeEventTypes().UpdatedExpensePurpose())
}

type ExpensePurposeAggregateInitializer struct {
    *eh.AggregateInitializer
}



const FeeAggregateType eventhorizon.AggregateType = "FeeAggregate"

type FeeAggregate struct {
    *eventhorizon.AggregateBase
    *Fee
}

func NewFeeAggregate(AggregateBase *eventhorizon.AggregateBase, Entity *Fee) (ret *FeeAggregate, err error) {
    ret = &FeeAggregate{
        AggregateBase: AggregateBase,
        Fee: Entity,
    }
    
    return
}



func NewFeeAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *FeeAggregateInitializer) {
	ret = &FeeAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(FeeAggregateType,
        FeeCommandTypes().Literals(), FeeEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
    }
	return
}


func (o *FeeAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, FeeEventTypes().CreatedFee())
}

func (o *FeeAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, FeeEventTypes().DeletedFee())
}

func (o *FeeAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, FeeEventTypes().UpdatedFee())
}

type FeeAggregateInitializer struct {
    *eh.AggregateInitializer
}



const FeeKindAggregateType eventhorizon.AggregateType = "FeeKindAggregate"

type FeeKindAggregate struct {
    *eventhorizon.AggregateBase
    *FeeKind
}

func NewFeeKindAggregate(AggregateBase *eventhorizon.AggregateBase, Entity *FeeKind) (ret *FeeKindAggregate, err error) {
    ret = &FeeKindAggregate{
        AggregateBase: AggregateBase,
        FeeKind: Entity,
    }
    
    return
}



func NewFeeKindAggregateInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *FeeKindAggregateInitializer) {
	ret = &FeeKindAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(FeeKindAggregateType,
        FeeKindCommandTypes().Literals(), FeeKindEventTypes().Literals(), eventStore, eventBus, eventPublisher, commandBus),
    }
	return
}


func (o *FeeKindAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, FeeKindEventTypes().CreatedFeeKind())
}

func (o *FeeKindAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, FeeKindEventTypes().DeletedFeeKind())
}

func (o *FeeKindAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, FeeKindEventTypes().UpdatedFeeKind())
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









