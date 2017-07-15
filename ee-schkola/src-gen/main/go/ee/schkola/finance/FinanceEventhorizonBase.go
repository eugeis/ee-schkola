package finance

import (
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

const ExpenseAggregateType eventhorizon.AggregateType = "ExpenseAggregate"

type ExpenseAggregateInitializer struct {
    *eh.AggregateInitializer
}

func (o *ExpenseAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ExpenseAggregateEventTypes().ExpenseCreated())
}

func (o *ExpenseAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ExpenseAggregateEventTypes().ExpenseDeleted())
}

func (o *ExpenseAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ExpenseAggregateEventTypes().ExpenseUpdated())
}

func NewExpenseAggregateInitializer(
	store *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher,
	commandBus *eventhorizon.CommandBus) (ret *ExpenseAggregateInitializer) {
	ret = &ExpenseAggregateInitializer{
        AggregateInitializer: eh.NewAggregateInitializer(ExpenseAggregateType, ExpenseAggregateCommandTypes().Literals(),
		ChurchAggregateEventTypes().Literals(), store, eventBus, publisher, commandBus),
    }
	return
}

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



const ExpensePurposeAggregateType eventhorizon.AggregateType = "ExpensePurposeAggregate"

type ExpensePurposeAggregateInitializer struct {
    *eh.AggregateInitializer
}

func (o *ExpensePurposeAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ExpensePurposeAggregateEventTypes().ExpensePurposeCreated())
}

func (o *ExpensePurposeAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ExpensePurposeAggregateEventTypes().ExpensePurposeDeleted())
}

func (o *ExpensePurposeAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, ExpensePurposeAggregateEventTypes().ExpensePurposeUpdated())
}

func NewExpensePurposeAggregateInitializer(
	store *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher,
	commandBus *eventhorizon.CommandBus) (ret *ExpensePurposeAggregateInitializer) {
	ret = &ExpensePurposeAggregateInitializer{
        AggregateInitializer: eh.NewAggregateInitializer(ExpensePurposeAggregateType, ExpensePurposeAggregateCommandTypes().Literals(),
		ChurchAggregateEventTypes().Literals(), store, eventBus, publisher, commandBus),
    }
	return
}

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



const FeeAggregateType eventhorizon.AggregateType = "FeeAggregate"

type FeeAggregateInitializer struct {
    *eh.AggregateInitializer
}

func (o *FeeAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, FeeAggregateEventTypes().FeeCreated())
}

func (o *FeeAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, FeeAggregateEventTypes().FeeDeleted())
}

func (o *FeeAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, FeeAggregateEventTypes().FeeUpdated())
}

func NewFeeAggregateInitializer(
	store *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher,
	commandBus *eventhorizon.CommandBus) (ret *FeeAggregateInitializer) {
	ret = &FeeAggregateInitializer{
        AggregateInitializer: eh.NewAggregateInitializer(FeeAggregateType, FeeAggregateCommandTypes().Literals(),
		ChurchAggregateEventTypes().Literals(), store, eventBus, publisher, commandBus),
    }
	return
}

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



const FeeKindAggregateType eventhorizon.AggregateType = "FeeKindAggregate"

type FeeKindAggregateInitializer struct {
    *eh.AggregateInitializer
}

func (o *FeeKindAggregateInitializer) RegisterForCreated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, FeeKindAggregateEventTypes().FeeKindCreated())
}

func (o *FeeKindAggregateInitializer) RegisterForDeleted(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, FeeKindAggregateEventTypes().FeeKindDeleted())
}

func (o *FeeKindAggregateInitializer) RegisterForUpdated(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, FeeKindAggregateEventTypes().FeeKindUpdated())
}

func NewFeeKindAggregateInitializer(
	store *eventhorizon.EventStore, eventBus *eventhorizon.EventBus, publisher *eventhorizon.EventPublisher,
	commandBus *eventhorizon.CommandBus) (ret *FeeKindAggregateInitializer) {
	ret = &FeeKindAggregateInitializer{
        AggregateInitializer: eh.NewAggregateInitializer(FeeKindAggregateType, FeeKindAggregateCommandTypes().Literals(),
		ChurchAggregateEventTypes().Literals(), store, eventBus, publisher, commandBus),
    }
	return
}

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











