package finance

import (
    "github.com/looplab/eventhorizon"
	"ee/schkola/finance"
)
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

    aggregateType := eventhorizon.AggregateType(FinanceAggregateTypes().Expense)
    for _, command := range ExpenseCommandTypes().Values() {
        handler.SetAggregate(aggregateType, eventhorizon.CommandType(command.Name()))
    }
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

    aggregateType := eventhorizon.AggregateType(FinanceAggregateTypes().ExpensePurpose)
    for _, command := range ExpensePurposeCommandTypes().Values() {
        handler.SetAggregate(aggregateType, eventhorizon.CommandType(command.Name()))
    }
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

    aggregateType := eventhorizon.AggregateType(FinanceAggregateTypes().Fee)
    for _, command := range FeeCommandTypes().Values() {
        handler.SetAggregate(aggregateType, eventhorizon.CommandType(command.Name()))
    }
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

    aggregateType := eventhorizon.AggregateType(FinanceAggregateTypes().FeeKind)
    for _, command := range FeeKindCommandTypes().Values() {
        handler.SetAggregate(aggregateType, eventhorizon.CommandType(command.Name()))
    }
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

type Named interface {
	Name() string
}








type FinanceAggregateType struct {
	name  string
	ordinal int
    commands []Named
    events []Named
}

func (o *FinanceAggregateType) Name() string {
    return o.name
}

func (o *FinanceAggregateType) Ordinal() int {
    return o.ordinal
}

func (o *FinanceAggregateType) Commands() []Named {
    return o.commands
}
func (o *FinanceAggregateType) Events() []Named {
    return o.events
}

func (o *FinanceAggregateType) IsExpense() bool {
    return o == _financeAggregateTypes.Expense()
}

func (o *FinanceAggregateType) IsExpensePurpose() bool {
    return o == _financeAggregateTypes.ExpensePurpose()
}

func (o *FinanceAggregateType) IsFee() bool {
    return o == _financeAggregateTypes.Fee()
}

func (o *FinanceAggregateType) IsFeeKind() bool {
    return o == _financeAggregateTypes.FeeKind()
}

type financeAggregateTypes struct {
	values []*FinanceAggregateType
}

var _financeAggregateTypes = &financeAggregateTypes{values: []*FinanceAggregateType{
    {name: "Expense", ordinal: 0, commands: finance.ExpenseCommandTypes().Names()},
    {name: "ExpensePurpose", ordinal: 1},
    {name: "Fee", ordinal: 2},
    {name: "FeeKind", ordinal: 3}},
}

func FinanceAggregateTypes() *financeAggregateTypes {
	return _financeAggregateTypes
}

func (o *financeAggregateTypes) Values() []*FinanceAggregateType {
	return o.values
}

func (o *financeAggregateTypes) Expense() *FinanceAggregateType {
    return _financeAggregateTypes.values[0]
}

func (o *financeAggregateTypes) ExpensePurpose() *FinanceAggregateType {
    return _financeAggregateTypes.values[1]
}

func (o *financeAggregateTypes) Fee() *FinanceAggregateType {
    return _financeAggregateTypes.values[2]
}

func (o *financeAggregateTypes) FeeKind() *FinanceAggregateType {
    return _financeAggregateTypes.values[3]
}

func (o *financeAggregateTypes) ParseFinanceAggregateType(name string) (ret *FinanceAggregateType, ok bool) {
    switch name {
      case o.Expense().Name():
        ret = o.Expense()
      case o.ExpensePurpose().Name():
        ret = o.ExpensePurpose()
      case o.Fee().Name():
        ret = o.Fee()
      case o.FeeKind().Name():
        ret = o.FeeKind()
    }
    return
}



