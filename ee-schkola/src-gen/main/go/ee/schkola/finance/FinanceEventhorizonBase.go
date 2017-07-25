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
	ret.CommandHandler = NewExpenseAggregateCommandHandler(ret)
    return
}

func (o *ExpenseAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}


func NewExpenseAggregateCommandHandler(aggregate *ExpenseAggregate) *ExpenseAggregateCommandHandler {
	return &ExpenseAggregateCommandHandler{
		aggregate: aggregate,
        handlers: make(map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *ExpenseAggregate) error),
    }
}

type ExpenseAggregateCommandHandler struct {
	aggregate *ExpenseAggregate
	handlers  map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *ExpenseAggregate) error
}

func (o *ExpenseAggregateCommandHandler) AddHandler(commandType eventhorizon.CommandType,
	handler func(cmd eventhorizon.Command, aggregate *ExpenseAggregate) error) {
	o.handlers[commandType] = handler
}

func (o *ExpenseAggregateCommandHandler) HandleCommand(ctx context.Context, cmd eventhorizon.Command) (err error) {
	if handler, ok := o.handlers[cmd.CommandType()]; ok {
		err = handler(cmd, o.aggregate)
	} else {
		err = errors.New(fmt.Sprintf("There is no handlers for command %v registered in the aggregate %v",
			cmd.CommandType(), cmd.AggregateType()))
	}
	return
}

type ExpenseAggregate struct {
    *eventhorizon.AggregateBase
    *Expense
    eventhorizon.CommandHandler
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

func NewExpensePurposeAggregate(id eventhorizon.UUID) (ret *ExpensePurposeAggregate) {
    ret = &ExpensePurposeAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(ExpensePurposeAggregateType, id),
    }
	ret.CommandHandler = NewExpensePurposeAggregateCommandHandler(ret)
    return
}

func (o *ExpensePurposeAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}


func NewExpensePurposeAggregateCommandHandler(aggregate *ExpensePurposeAggregate) *ExpensePurposeAggregateCommandHandler {
	return &ExpensePurposeAggregateCommandHandler{
		aggregate: aggregate,
        handlers: make(map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *ExpensePurposeAggregate) error),
    }
}

type ExpensePurposeAggregateCommandHandler struct {
	aggregate *ExpensePurposeAggregate
	handlers  map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *ExpensePurposeAggregate) error
}

func (o *ExpensePurposeAggregateCommandHandler) AddHandler(commandType eventhorizon.CommandType,
	handler func(cmd eventhorizon.Command, aggregate *ExpensePurposeAggregate) error) {
	o.handlers[commandType] = handler
}

func (o *ExpensePurposeAggregateCommandHandler) HandleCommand(ctx context.Context, cmd eventhorizon.Command) (err error) {
	if handler, ok := o.handlers[cmd.CommandType()]; ok {
		err = handler(cmd, o.aggregate)
	} else {
		err = errors.New(fmt.Sprintf("There is no handlers for command %v registered in the aggregate %v",
			cmd.CommandType(), cmd.AggregateType()))
	}
	return
}

type ExpensePurposeAggregate struct {
    *eventhorizon.AggregateBase
    *ExpensePurpose
    eventhorizon.CommandHandler
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

func NewFeeAggregate(id eventhorizon.UUID) (ret *FeeAggregate) {
    ret = &FeeAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(FeeAggregateType, id),
    }
	ret.CommandHandler = NewFeeAggregateCommandHandler(ret)
    return
}

func (o *FeeAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}


func NewFeeAggregateCommandHandler(aggregate *FeeAggregate) *FeeAggregateCommandHandler {
	return &FeeAggregateCommandHandler{
		aggregate: aggregate,
        handlers: make(map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *FeeAggregate) error),
    }
}

type FeeAggregateCommandHandler struct {
	aggregate *FeeAggregate
	handlers  map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *FeeAggregate) error
}

func (o *FeeAggregateCommandHandler) AddHandler(commandType eventhorizon.CommandType,
	handler func(cmd eventhorizon.Command, aggregate *FeeAggregate) error) {
	o.handlers[commandType] = handler
}

func (o *FeeAggregateCommandHandler) HandleCommand(ctx context.Context, cmd eventhorizon.Command) (err error) {
	if handler, ok := o.handlers[cmd.CommandType()]; ok {
		err = handler(cmd, o.aggregate)
	} else {
		err = errors.New(fmt.Sprintf("There is no handlers for command %v registered in the aggregate %v",
			cmd.CommandType(), cmd.AggregateType()))
	}
	return
}

type FeeAggregate struct {
    *eventhorizon.AggregateBase
    *Fee
    eventhorizon.CommandHandler
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

func NewFeeKindAggregate(id eventhorizon.UUID) (ret *FeeKindAggregate) {
    ret = &FeeKindAggregate{
		AggregateBase: eventhorizon.NewAggregateBase(FeeKindAggregateType, id),
    }
	ret.CommandHandler = NewFeeKindAggregateCommandHandler(ret)
    return
}

func (o *FeeKindAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
    println("ApplyEvent", event.EventType())
    return nil
}


func NewFeeKindAggregateCommandHandler(aggregate *FeeKindAggregate) *FeeKindAggregateCommandHandler {
	return &FeeKindAggregateCommandHandler{
		aggregate: aggregate,
        handlers: make(map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *FeeKindAggregate) error),
    }
}

type FeeKindAggregateCommandHandler struct {
	aggregate *FeeKindAggregate
	handlers  map[eventhorizon.CommandType]func(cmd eventhorizon.Command, aggregate *FeeKindAggregate) error
}

func (o *FeeKindAggregateCommandHandler) AddHandler(commandType eventhorizon.CommandType,
	handler func(cmd eventhorizon.Command, aggregate *FeeKindAggregate) error) {
	o.handlers[commandType] = handler
}

func (o *FeeKindAggregateCommandHandler) HandleCommand(ctx context.Context, cmd eventhorizon.Command) (err error) {
	if handler, ok := o.handlers[cmd.CommandType()]; ok {
		err = handler(cmd, o.aggregate)
	} else {
		err = errors.New(fmt.Sprintf("There is no handlers for command %v registered in the aggregate %v",
			cmd.CommandType(), cmd.AggregateType()))
	}
	return
}

type FeeKindAggregate struct {
    *eventhorizon.AggregateBase
    *FeeKind
    eventhorizon.CommandHandler
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









