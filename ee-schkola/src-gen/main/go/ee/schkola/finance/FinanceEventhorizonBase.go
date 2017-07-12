package finance

import (
    "github.com/looplab/eventhorizon"
)
type ExpenseAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func (o *ExpenseAggregateInitializer) setup() string {
    panic("Not implemented yet.")
}

func (o *ExpenseAggregateInitializer) registerCommands() string {
    panic("Not implemented yet.")
}


type ExpensePurposeAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func (o *ExpensePurposeAggregateInitializer) setup() string {
    panic("Not implemented yet.")
}

func (o *ExpensePurposeAggregateInitializer) registerCommands() string {
    panic("Not implemented yet.")
}


type FeeAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func (o *FeeAggregateInitializer) setup() string {
    panic("Not implemented yet.")
}

func (o *FeeAggregateInitializer) registerCommands() string {
    panic("Not implemented yet.")
}


type FeeKindAggregateInitializer struct {
    Store  *eventhorizon.EventStore
    Notifier  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    Executor  *eventhorizon.CommandBus
}

func (o *FeeKindAggregateInitializer) setup() string {
    panic("Not implemented yet.")
}

func (o *FeeKindAggregateInitializer) registerCommands() string {
    panic("Not implemented yet.")
}


type FinanceEventhorizonInitializer struct {
    Store  *eventhorizon.EventStore
    EventBus  *eventhorizon.EventBus
    Publisher  *eventhorizon.EventPublisher
    CommandBus  *eventhorizon.CommandBus
}

func (o *FinanceEventhorizonInitializer) setup() string {
    panic("Not implemented yet.")
}

func (o *FinanceEventhorizonInitializer) registerCommands() string {
    panic("Not implemented yet.")
}


type ExpenseAggregate struct {
    *eventhorizon.AggregateBase
    *Expense
}


type ExpensePurposeAggregate struct {
    *eventhorizon.AggregateBase
    *ExpensePurpose
}


type FeeAggregate struct {
    *eventhorizon.AggregateBase
    *Fee
}


type FeeKindAggregate struct {
    *eventhorizon.AggregateBase
    *FeeKind
}




type FinanceAggregateType struct {
	name  string
	ordinal int
}

func (o *FinanceAggregateType) Name() string {
    return o.name
}

func (o *FinanceAggregateType) Ordinal() int {
    return o.ordinal
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
    {name: "Expense", ordinal: 0},
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
      case "Expense":
        ret = o.Expense()
      case "ExpensePurpose":
        ret = o.ExpensePurpose()
      case "Fee":
        ret = o.Fee()
      case "FeeKind":
        ret = o.FeeKind()
    }
    return
}



