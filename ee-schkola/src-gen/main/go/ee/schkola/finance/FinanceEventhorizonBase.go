package finance

import (
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

type ExpenseAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewExpenseAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *ExpenseAggregateInitializer, err error) {
    ret = &ExpenseAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *ExpenseAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, ExpenseAggregateType, ExpenseCommandTypes().Literals())
}




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



type ExpensePurposeAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewExpensePurposeAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *ExpensePurposeAggregateInitializer, err error) {
    ret = &ExpensePurposeAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *ExpensePurposeAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, ExpensePurposeAggregateType, ExpensePurposeCommandTypes().Literals())
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



type FeeAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewFeeAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *FeeAggregateInitializer, err error) {
    ret = &FeeAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *FeeAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, FeeAggregateType, FeeCommandTypes().Literals())
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



type FeeKindAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewFeeKindAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *FeeKindAggregateInitializer, err error) {
    ret = &FeeKindAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *FeeKindAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, FeeKindAggregateType, FeeKindCommandTypes().Literals())
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



type FinanceEventhorizonInitializer struct {
    Store  *eventhorizon.EventStore
    EventBus  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    CommandBus  *eventhorizon.CommandBus
}

func NewFinanceEventhorizonInitializer(store *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                commandBus *eventhorizon.CommandBus) (ret *FinanceEventhorizonInitializer, err error) {
    ret = &FinanceEventhorizonInitializer{
        Store : store,
        EventBus : eventBus,
        Publisher : publisher,
        CommandBus : commandBus,
    }
    return
}











