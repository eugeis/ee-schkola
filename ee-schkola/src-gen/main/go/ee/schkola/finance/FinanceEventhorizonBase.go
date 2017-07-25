package finance

import (
    "context"
    "errors"
    "fmt"
    "github.com/looplab/eventhorizon"
    "github.com/eugeis/gee/eh"
)

const ExpenseAggregateType eventhorizon.AggregateType = "ExpenseAggregate"

func NewExpenseAggregate(id eventhorizon.UUID) (ret *ExpenseAggregate) {
    ret = &ExpenseAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(ExpenseAggregateType, id),
    }
	//ret.CommandHandler = NewExpenseAggregateCommandHandler(ret)
    return
}

func (o *ExpenseAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type ExpenseAggregate struct {
    *eventhorizon.AggregateBase
    *Expense
    eventhorizon.CommandHandler
}



type ExpenseCommandHandler struct {
    CreateHandler  func (*CreateExpense, *ExpenseAggregate) error
    DeleteHandler  func (*DeleteExpense, *ExpenseAggregate) error
    UpdateHandler  func (*UpdateExpense, *ExpenseAggregate) error
}

func (o *ExpenseCommandHandler) HandleCommand(ctx *context.Context, cmd eventhorizon.Command, aggregate *ExpenseAggregate) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateExpenseCommand:
        ret = o.CreateHandler(cmd.(*CreateExpense), aggregate)
    case DeleteExpenseCommand:
        ret = o.DeleteHandler(cmd.(*DeleteExpense), aggregate)
    case UpdateExpenseCommand:
        ret = o.UpdateHandler(cmd.(*UpdateExpense), aggregate)
    default:
		ret = errors.New(fmt.Sprintf("Wrong comand type '%v' for aggregate '%v", cmd.CommandType(), aggregate))
	}
    return ret
    
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
    *ExpenseCommandHandler
}



const ExpensePurposeAggregateType eventhorizon.AggregateType = "ExpensePurposeAggregate"

func NewExpensePurposeAggregate(id eventhorizon.UUID) (ret *ExpensePurposeAggregate) {
    ret = &ExpensePurposeAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(ExpensePurposeAggregateType, id),
    }
	//ret.CommandHandler = NewExpensePurposeAggregateCommandHandler(ret)
    return
}

func (o *ExpensePurposeAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type ExpensePurposeAggregate struct {
    *eventhorizon.AggregateBase
    *ExpensePurpose
    eventhorizon.CommandHandler
}



type ExpensePurposeCommandHandler struct {
    CreateHandler  func (*CreateExpensePurpose, *ExpensePurposeAggregate) error
    DeleteHandler  func (*DeleteExpensePurpose, *ExpensePurposeAggregate) error
    UpdateHandler  func (*UpdateExpensePurpose, *ExpensePurposeAggregate) error
}

func (o *ExpensePurposeCommandHandler) HandleCommand(ctx *context.Context, cmd eventhorizon.Command, aggregate *ExpensePurposeAggregate) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateExpensePurposeCommand:
        ret = o.CreateHandler(cmd.(*CreateExpensePurpose), aggregate)
    case DeleteExpensePurposeCommand:
        ret = o.DeleteHandler(cmd.(*DeleteExpensePurpose), aggregate)
    case UpdateExpensePurposeCommand:
        ret = o.UpdateHandler(cmd.(*UpdateExpensePurpose), aggregate)
    default:
		ret = errors.New(fmt.Sprintf("Wrong comand type '%v' for aggregate '%v", cmd.CommandType(), aggregate))
	}
    return ret
    
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
    *ExpensePurposeCommandHandler
}



const FeeAggregateType eventhorizon.AggregateType = "FeeAggregate"

func NewFeeAggregate(id eventhorizon.UUID) (ret *FeeAggregate) {
    ret = &FeeAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(FeeAggregateType, id),
    }
	//ret.CommandHandler = NewFeeAggregateCommandHandler(ret)
    return
}

func (o *FeeAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type FeeAggregate struct {
    *eventhorizon.AggregateBase
    *Fee
    eventhorizon.CommandHandler
}



type FeeCommandHandler struct {
    CreateHandler  func (*CreateFee, *FeeAggregate) error
    DeleteHandler  func (*DeleteFee, *FeeAggregate) error
    UpdateHandler  func (*UpdateFee, *FeeAggregate) error
}

func (o *FeeCommandHandler) HandleCommand(ctx *context.Context, cmd eventhorizon.Command, aggregate *FeeAggregate) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateFeeCommand:
        ret = o.CreateHandler(cmd.(*CreateFee), aggregate)
    case DeleteFeeCommand:
        ret = o.DeleteHandler(cmd.(*DeleteFee), aggregate)
    case UpdateFeeCommand:
        ret = o.UpdateHandler(cmd.(*UpdateFee), aggregate)
    default:
		ret = errors.New(fmt.Sprintf("Wrong comand type '%v' for aggregate '%v", cmd.CommandType(), aggregate))
	}
    return ret
    
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
    *FeeCommandHandler
}



const FeeKindAggregateType eventhorizon.AggregateType = "FeeKindAggregate"

func NewFeeKindAggregate(id eventhorizon.UUID) (ret *FeeKindAggregate) {
    ret = &FeeKindAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(FeeKindAggregateType, id),
    }
	//ret.CommandHandler = NewFeeKindAggregateCommandHandler(ret)
    return
}

func (o *FeeKindAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}

type FeeKindAggregate struct {
    *eventhorizon.AggregateBase
    *FeeKind
    eventhorizon.CommandHandler
}



type FeeKindCommandHandler struct {
    CreateHandler  func (*CreateFeeKind, *FeeKindAggregate) error
    DeleteHandler  func (*DeleteFeeKind, *FeeKindAggregate) error
    UpdateHandler  func (*UpdateFeeKind, *FeeKindAggregate) error
}

func (o *FeeKindCommandHandler) HandleCommand(ctx *context.Context, cmd eventhorizon.Command, aggregate *FeeKindAggregate) error {
    
    var ret error
    switch cmd.CommandType() {
    case CreateFeeKindCommand:
        ret = o.CreateHandler(cmd.(*CreateFeeKind), aggregate)
    case DeleteFeeKindCommand:
        ret = o.DeleteHandler(cmd.(*DeleteFeeKind), aggregate)
    case UpdateFeeKindCommand:
        ret = o.UpdateHandler(cmd.(*UpdateFeeKind), aggregate)
    default:
		ret = errors.New(fmt.Sprintf("Wrong comand type '%v' for aggregate '%v", cmd.CommandType(), aggregate))
	}
    return ret
    
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
    *FeeKindCommandHandler
}



func NewFinanceEventhorizonInitializer(
	eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher,
	commandBus eventhorizon.CommandBus) (ret *FinanceEventhorizonInitializer) {
	ret = &FinanceEventhorizonInitializer{eventStore: eventStore, eventBus: eventBus, eventPublisher: eventPublisher,
            commandBus: commandBus, 
    ExpenseAggregateInitializer: NewExpenseAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    ExpensePurposeAggregateInitializer: NewExpensePurposeAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    FeeAggregateInitializer: NewFeeAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    FeeKindAggregateInitializer: NewFeeKindAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus)}
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

type FinanceEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore
    eventBus eventhorizon.EventBus
    eventPublisher eventhorizon.EventPublisher
    commandBus eventhorizon.CommandBus
    ExpenseAggregateInitializer  *ExpenseAggregateInitializer
    ExpensePurposeAggregateInitializer  *ExpensePurposeAggregateInitializer
    FeeAggregateInitializer  *FeeAggregateInitializer
    FeeKindAggregateInitializer  *FeeKindAggregateInitializer
}









