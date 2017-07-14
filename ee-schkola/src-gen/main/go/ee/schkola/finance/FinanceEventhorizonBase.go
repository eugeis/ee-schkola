package finance

import (
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

type ExpenseExpenseAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewExpenseExpenseAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *ExpenseExpenseAggregateInitializer, err error) {
    ret = &ExpenseExpenseAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *ExpenseExpenseAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, ExpenseAggregateType, ExpenseCommandTypes().Literals())
}




const ExpenseExpenseType eventhorizon.AggregateType = "ExpenseExpense"
type ExpenseExpense struct {
    *eventhorizon.AggregateBase
    *Expense
}

func NewExpenseExpense(AggregateBase *eventhorizon.AggregateBase, Entity *Expense) (ret *ExpenseExpense, err error) {
    ret = &ExpenseExpense{
        AggregateBase: AggregateBase,
        Expense: Entity,
    }
    return
}



type ExpensePurposeExpensePurposeAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewExpensePurposeExpensePurposeAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *ExpensePurposeExpensePurposeAggregateInitializer, err error) {
    ret = &ExpensePurposeExpensePurposeAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *ExpensePurposeExpensePurposeAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, ExpensePurposeAggregateType, ExpensePurposeCommandTypes().Literals())
}




const ExpensePurposeExpensePurposeType eventhorizon.AggregateType = "ExpensePurposeExpensePurpose"
type ExpensePurposeExpensePurpose struct {
    *eventhorizon.AggregateBase
    *ExpensePurpose
}

func NewExpensePurposeExpensePurpose(AggregateBase *eventhorizon.AggregateBase, Entity *ExpensePurpose) (ret *ExpensePurposeExpensePurpose, err error) {
    ret = &ExpensePurposeExpensePurpose{
        AggregateBase: AggregateBase,
        ExpensePurpose: Entity,
    }
    return
}



type FeeFeeAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewFeeFeeAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *FeeFeeAggregateInitializer, err error) {
    ret = &FeeFeeAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *FeeFeeAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, FeeAggregateType, FeeCommandTypes().Literals())
}




const FeeFeeType eventhorizon.AggregateType = "FeeFee"
type FeeFee struct {
    *eventhorizon.AggregateBase
    *Fee
}

func NewFeeFee(AggregateBase *eventhorizon.AggregateBase, Entity *Fee) (ret *FeeFee, err error) {
    ret = &FeeFee{
        AggregateBase: AggregateBase,
        Fee: Entity,
    }
    return
}



type FeeKindFeeKindAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func NewFeeKindFeeKindAggregateInitializer(store *eventhorizon.EventStore, notifier *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher, 
                executor *eventhorizon.CommandBus) (ret *FeeKindFeeKindAggregateInitializer, err error) {
    ret = &FeeKindFeeKindAggregateInitializer{
        Store : store,
        Notifier : notifier,
        Publisher : publisher,
        Executor : executor,
    }
    return
}


func (o *FeeKindFeeKindAggregateInitializer) RegisterCommands(handler *eventhorizon.AggregateCommandHandler)  {
    eh.RegisterCommands(handler, FeeKindAggregateType, FeeKindCommandTypes().Literals())
}




const FeeKindFeeKindType eventhorizon.AggregateType = "FeeKindFeeKind"
type FeeKindFeeKind struct {
    *eventhorizon.AggregateBase
    *FeeKind
}

func NewFeeKindFeeKind(AggregateBase *eventhorizon.AggregateBase, Entity *FeeKind) (ret *FeeKindFeeKind, err error) {
    ret = &FeeKindFeeKind{
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











