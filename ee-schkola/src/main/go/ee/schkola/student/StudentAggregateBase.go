package student

import (
    "errors"
    "fmt"
    "github.com/eugeis/gee/eh"
    "github.com/looplab/eventhorizon"
    "github.com/looplab/eventhorizon/commandhandler/bus"
    "time"
)
type CommandHandler struct {
    RegisterHandler func (*Register, *Attendance, eh.AggregateStoreEvent) (err error)  `json:"registerHandler" eh:"optional"`
    CreateHandler func (*Create, *Attendance, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*Delete, *Attendance, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    CancelHandler func (*Cancel, *Attendance, eh.AggregateStoreEvent) (err error)  `json:"cancelHandler" eh:"optional"`
    ConfirmHandler func (*Confirm, *Attendance, eh.AggregateStoreEvent) (err error)  `json:"confirmHandler" eh:"optional"`
    UpdateHandler func (*Update, *Attendance, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *CommandHandler) AddRegisterPreparer(preparer func (*Register, *Attendance) (err error) ) {
    prevHandler := o.RegisterHandler
	o.RegisterHandler = func(command *Register, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddCreatePreparer(preparer func (*Create, *Attendance) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *Create, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddDeletePreparer(preparer func (*Delete, *Attendance) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *Delete, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddCancelPreparer(preparer func (*Cancel, *Attendance) (err error) ) {
    prevHandler := o.CancelHandler
	o.CancelHandler = func(command *Cancel, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddConfirmPreparer(preparer func (*Confirm, *Attendance) (err error) ) {
    prevHandler := o.ConfirmHandler
	o.ConfirmHandler = func(command *Confirm, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddUpdatePreparer(preparer func (*Update, *Attendance) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *Update, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case RegisterCommand:
        err = o.RegisterHandler(cmd.(*Register), entity.(*Attendance), store)
    case CreateCommand:
        err = o.CreateHandler(cmd.(*Create), entity.(*Attendance), store)
    case DeleteCommand:
        err = o.DeleteHandler(cmd.(*Delete), entity.(*Attendance), store)
    case CancelCommand:
        err = o.CancelHandler(cmd.(*Cancel), entity.(*Attendance), store)
    case ConfirmCommand:
        err = o.ConfirmHandler(cmd.(*Confirm), entity.(*Attendance), store)
    case UpdateCommand:
        err = o.UpdateHandler(cmd.(*Update), entity.(*Attendance), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *CommandHandler) SetupCommandHandler() (err error) {
    o.RegisterHandler = func(command *Register, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, AttendanceAggregateType); err == nil {
            store.StoreEvent(RegisteredEvent, &Registered{
                Student: command.Student,
                Course: command.Course,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.CreateHandler = func(command *Create, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, AttendanceAggregateType); err == nil {
            store.StoreEvent(createdEvent, &Created{
                Student: command.Student,
                Date: command.Date,
                Course: command.Course,
                Hours: command.Hours,
                State: command.State,
                Token: command.Token,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *Delete, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AttendanceAggregateType); err == nil {
            store.StoreEvent(deletedEvent, &Deleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.CancelHandler = func(command *Cancel, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AttendanceAggregateType); err == nil {
            store.StoreEvent(CanceledEvent, &Canceled{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.ConfirmHandler = func(command *Confirm, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AttendanceAggregateType); err == nil {
            store.StoreEvent(ConfirmedEvent, &Confirmed{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *Update, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AttendanceAggregateType); err == nil {
            store.StoreEvent(updatedEvent, &Updated{
                Student: command.Student,
                Date: command.Date,
                Course: command.Course,
                Hours: command.Hours,
                State: command.State,
                Token: command.Token,
                Id: command.Id,}, time.Now())
        }
        return
    }
    return
}


type EventHandler struct {
    RegisterHandler func (*Register, *Attendance) (err error)  `json:"registerHandler" eh:"optional"`
    CreateHandler func (*Create, *Attendance) (err error)  `json:"createHandler" eh:"optional"`
    CreatedHandler func (*Created, *Attendance) (err error)  `json:"createdHandler" eh:"optional"`
    RegisteredHandler func (*Registered, *Attendance) (err error)  `json:"registeredHandler" eh:"optional"`
    DeleteHandler func (*Delete, *Attendance) (err error)  `json:"deleteHandler" eh:"optional"`
    DeletedHandler func (*Deleted, *Attendance) (err error)  `json:"deletedHandler" eh:"optional"`
    CancelHandler func (*Cancel, *Attendance) (err error)  `json:"cancelHandler" eh:"optional"`
    ConfirmHandler func (*Confirm, *Attendance) (err error)  `json:"confirmHandler" eh:"optional"`
    UpdateHandler func (*Update, *Attendance) (err error)  `json:"updateHandler" eh:"optional"`
    UpdatedHandler func (*Updated, *Attendance) (err error)  `json:"updatedHandler" eh:"optional"`
    ConfirmedHandler func (*Confirmed, *Attendance) (err error)  `json:"confirmedHandler" eh:"optional"`
    CanceledHandler func (*Canceled, *Attendance) (err error)  `json:"canceledHandler" eh:"optional"`
}

func (o *EventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case RegisterEvent:
        err = o.RegisterHandler(event.Data().(*Register), entity.(*Attendance))
    case CreateEvent:
        err = o.CreateHandler(event.Data().(*Create), entity.(*Attendance))
    case CreatedEvent:
        err = o.CreatedHandler(event.Data().(*Created), entity.(*Attendance))
    case RegisteredEvent:
        err = o.RegisteredHandler(event.Data().(*Registered), entity.(*Attendance))
    case DeleteEvent:
        err = o.DeleteHandler(event.Data().(*Delete), entity.(*Attendance))
    case DeletedEvent:
        err = o.DeletedHandler(event.Data().(*Deleted), entity.(*Attendance))
    case CancelEvent:
        err = o.CancelHandler(event.Data().(*Cancel), entity.(*Attendance))
    case ConfirmEvent:
        err = o.ConfirmHandler(event.Data().(*Confirm), entity.(*Attendance))
    case UpdateEvent:
        err = o.UpdateHandler(event.Data().(*Update), entity.(*Attendance))
    case UpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*Updated), entity.(*Attendance))
    case ConfirmedEvent:
        err = o.ConfirmedHandler(event.Data().(*Confirmed), entity.(*Attendance))
    case CanceledEvent:
        err = o.CanceledHandler(event.Data().(*Canceled), entity.(*Attendance))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *EventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(RegisterEvent, func() eventhorizon.EventData {
		return &Register{}
	})

    //default handler implementation
    o.RegisterHandler = func(event *Register, entity *Attendance) (err error) {
        //err = eh.EventHandlerNotImplemented(RegisterEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(CreateEvent, func() eventhorizon.EventData {
		return &Create{}
	})

    //default handler implementation
    o.CreateHandler = func(event *Create, entity *Attendance) (err error) {
        //err = eh.EventHandlerNotImplemented(CreateEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(CreatedEvent, func() eventhorizon.EventData {
		return &Created{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *Created, entity *Attendance) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, AttendanceAggregateType); err == nil {
            entity.State = AttendanceStates().Canceled()
            entity.Student = event.Student
            entity.Date = event.Date
            entity.Course = event.Course
            entity.Hours = event.Hours
            entity.State = event.State
            entity.Token = event.Token
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(RegisteredEvent, func() eventhorizon.EventData {
		return &Registered{}
	})

    //default handler implementation
    o.RegisteredHandler = func(event *Registered, entity *Attendance) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, AttendanceAggregateType); err == nil {
            entity.Student = event.Student
            entity.Course = event.Course
            entity.State = AttendanceStates().Registered()
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DeleteEvent, func() eventhorizon.EventData {
		return &Delete{}
	})

    //default handler implementation
    o.DeleteHandler = func(event *Delete, entity *Attendance) (err error) {
        //err = eh.EventHandlerNotImplemented(DeleteEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DeletedEvent, func() eventhorizon.EventData {
		return &Deleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *Deleted, entity *Attendance) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AttendanceAggregateType); err == nil {
            *entity = *NewAttendance()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(CancelEvent, func() eventhorizon.EventData {
		return &Cancel{}
	})

    //default handler implementation
    o.CancelHandler = func(event *Cancel, entity *Attendance) (err error) {
        //err = eh.EventHandlerNotImplemented(CancelEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(ConfirmEvent, func() eventhorizon.EventData {
		return &Confirm{}
	})

    //default handler implementation
    o.ConfirmHandler = func(event *Confirm, entity *Attendance) (err error) {
        //err = eh.EventHandlerNotImplemented(ConfirmEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdateEvent, func() eventhorizon.EventData {
		return &Update{}
	})

    //default handler implementation
    o.UpdateHandler = func(event *Update, entity *Attendance) (err error) {
        //err = eh.EventHandlerNotImplemented(UpdateEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdatedEvent, func() eventhorizon.EventData {
		return &Updated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *Updated, entity *Attendance) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AttendanceAggregateType); err == nil {
            entity.State = AttendanceStates().Canceled()
            entity.Student = event.Student
            entity.Date = event.Date
            entity.Course = event.Course
            entity.Hours = event.Hours
            entity.State = event.State
            entity.Token = event.Token
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(ConfirmedEvent, func() eventhorizon.EventData {
		return &Confirmed{}
	})

    //default handler implementation
    o.ConfirmedHandler = func(event *Confirmed, entity *Attendance) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AttendanceAggregateType); err == nil {
            entity.State = AttendanceStates().Confirmed()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(CanceledEvent, func() eventhorizon.EventData {
		return &Canceled{}
	})

    //default handler implementation
    o.CanceledHandler = func(event *Canceled, entity *Attendance) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AttendanceAggregateType); err == nil {
            entity.State = AttendanceStates().Canceled()
        }
        return
    }
    return
}


const AttendanceAggregateType eventhorizon.AggregateType = "Attendance"

type AggregateInitializer struct {
    *eh.AggregateInitializer
    *CommandHandler
    *EventHandler
    ProjectorHandler *EventHandler `json:"projectorHandler" eh:"optional"`
}



func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *AggregateInitializer) {
    
    commandHandler := &AttendanceCommandHandler{}
    eventHandler := &AttendanceEventHandler{}
    entityFactory := func() eventhorizon.Entity { return NewAttendance() }
    ret = &AggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(AttendanceAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(AttendanceAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        AttendanceCommandTypes().Literals(), AttendanceEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), AttendanceCommandHandler: commandHandler, AttendanceEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type CommandHandler struct {
    CreateHandler func (*Create, *Course, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*Delete, *Course, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*Update, *Course, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *CommandHandler) AddCreatePreparer(preparer func (*Create, *Course) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *Create, entity *Course, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddDeletePreparer(preparer func (*Delete, *Course) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *Delete, entity *Course, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddUpdatePreparer(preparer func (*Update, *Course) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *Update, entity *Course, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateCommand:
        err = o.CreateHandler(cmd.(*Create), entity.(*Course), store)
    case DeleteCommand:
        err = o.DeleteHandler(cmd.(*Delete), entity.(*Course), store)
    case UpdateCommand:
        err = o.UpdateHandler(cmd.(*Update), entity.(*Course), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *CommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *Create, entity *Course, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, CourseAggregateType); err == nil {
            store.StoreEvent(createdEvent, &Created{
                Name: command.Name,
                Begin: command.Begin,
                End: command.End,
                Teacher: command.Teacher,
                SchoolYear: command.SchoolYear,
                Fee: command.Fee,
                Description: command.Description,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *Delete, entity *Course, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, CourseAggregateType); err == nil {
            store.StoreEvent(deletedEvent, &Deleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *Update, entity *Course, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, CourseAggregateType); err == nil {
            store.StoreEvent(updatedEvent, &Updated{
                Name: command.Name,
                Begin: command.Begin,
                End: command.End,
                Teacher: command.Teacher,
                SchoolYear: command.SchoolYear,
                Fee: command.Fee,
                Description: command.Description,
                Id: command.Id,}, time.Now())
        }
        return
    }
    return
}


type EventHandler struct {
    CreateHandler func (*Create, *Course) (err error)  `json:"createHandler" eh:"optional"`
    CreatedHandler func (*Created, *Course) (err error)  `json:"createdHandler" eh:"optional"`
    DeleteHandler func (*Delete, *Course) (err error)  `json:"deleteHandler" eh:"optional"`
    DeletedHandler func (*Deleted, *Course) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdateHandler func (*Update, *Course) (err error)  `json:"updateHandler" eh:"optional"`
    UpdatedHandler func (*Updated, *Course) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *EventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case CreateEvent:
        err = o.CreateHandler(event.Data().(*Create), entity.(*Course))
    case CreatedEvent:
        err = o.CreatedHandler(event.Data().(*Created), entity.(*Course))
    case DeleteEvent:
        err = o.DeleteHandler(event.Data().(*Delete), entity.(*Course))
    case DeletedEvent:
        err = o.DeletedHandler(event.Data().(*Deleted), entity.(*Course))
    case UpdateEvent:
        err = o.UpdateHandler(event.Data().(*Update), entity.(*Course))
    case UpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*Updated), entity.(*Course))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *EventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(CreateEvent, func() eventhorizon.EventData {
		return &Create{}
	})

    //default handler implementation
    o.CreateHandler = func(event *Create, entity *Course) (err error) {
        //err = eh.EventHandlerNotImplemented(CreateEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(CreatedEvent, func() eventhorizon.EventData {
		return &Created{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *Created, entity *Course) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, CourseAggregateType); err == nil {
            entity.Name = event.Name
            entity.Begin = event.Begin
            entity.End = event.End
            entity.Teacher = event.Teacher
            entity.SchoolYear = event.SchoolYear
            entity.Fee = event.Fee
            entity.Description = event.Description
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DeleteEvent, func() eventhorizon.EventData {
		return &Delete{}
	})

    //default handler implementation
    o.DeleteHandler = func(event *Delete, entity *Course) (err error) {
        //err = eh.EventHandlerNotImplemented(DeleteEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DeletedEvent, func() eventhorizon.EventData {
		return &Deleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *Deleted, entity *Course) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, CourseAggregateType); err == nil {
            *entity = *NewCourse()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdateEvent, func() eventhorizon.EventData {
		return &Update{}
	})

    //default handler implementation
    o.UpdateHandler = func(event *Update, entity *Course) (err error) {
        //err = eh.EventHandlerNotImplemented(UpdateEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdatedEvent, func() eventhorizon.EventData {
		return &Updated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *Updated, entity *Course) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, CourseAggregateType); err == nil {
            entity.Name = event.Name
            entity.Begin = event.Begin
            entity.End = event.End
            entity.Teacher = event.Teacher
            entity.SchoolYear = event.SchoolYear
            entity.Fee = event.Fee
            entity.Description = event.Description
        }
        return
    }
    return
}


const CourseAggregateType eventhorizon.AggregateType = "Course"

type AggregateInitializer struct {
    *eh.AggregateInitializer
    *CommandHandler
    *EventHandler
    ProjectorHandler *EventHandler `json:"projectorHandler" eh:"optional"`
}



func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *AggregateInitializer) {
    
    commandHandler := &CourseCommandHandler{}
    eventHandler := &CourseEventHandler{}
    entityFactory := func() eventhorizon.Entity { return NewCourse() }
    ret = &AggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(CourseAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(CourseAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        CourseCommandTypes().Literals(), CourseEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), CourseCommandHandler: commandHandler, CourseEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type CommandHandler struct {
    CreateHandler func (*Create, *Grade, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*Delete, *Grade, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*Update, *Grade, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *CommandHandler) AddCreatePreparer(preparer func (*Create, *Grade) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *Create, entity *Grade, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddDeletePreparer(preparer func (*Delete, *Grade) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *Delete, entity *Grade, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddUpdatePreparer(preparer func (*Update, *Grade) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *Update, entity *Grade, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateCommand:
        err = o.CreateHandler(cmd.(*Create), entity.(*Grade), store)
    case DeleteCommand:
        err = o.DeleteHandler(cmd.(*Delete), entity.(*Grade), store)
    case UpdateCommand:
        err = o.UpdateHandler(cmd.(*Update), entity.(*Grade), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *CommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *Create, entity *Grade, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, GradeAggregateType); err == nil {
            store.StoreEvent(createdEvent, &Created{
                Student: command.Student,
                Course: command.Course,
                Grade: command.Grade,
                Comment: command.Comment,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *Delete, entity *Grade, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, GradeAggregateType); err == nil {
            store.StoreEvent(deletedEvent, &Deleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *Update, entity *Grade, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, GradeAggregateType); err == nil {
            store.StoreEvent(updatedEvent, &Updated{
                Student: command.Student,
                Course: command.Course,
                Grade: command.Grade,
                Comment: command.Comment,
                Id: command.Id,}, time.Now())
        }
        return
    }
    return
}


type EventHandler struct {
    CreateHandler func (*Create, *Grade) (err error)  `json:"createHandler" eh:"optional"`
    CreatedHandler func (*Created, *Grade) (err error)  `json:"createdHandler" eh:"optional"`
    DeleteHandler func (*Delete, *Grade) (err error)  `json:"deleteHandler" eh:"optional"`
    DeletedHandler func (*Deleted, *Grade) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdateHandler func (*Update, *Grade) (err error)  `json:"updateHandler" eh:"optional"`
    UpdatedHandler func (*Updated, *Grade) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *EventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case CreateEvent:
        err = o.CreateHandler(event.Data().(*Create), entity.(*Grade))
    case CreatedEvent:
        err = o.CreatedHandler(event.Data().(*Created), entity.(*Grade))
    case DeleteEvent:
        err = o.DeleteHandler(event.Data().(*Delete), entity.(*Grade))
    case DeletedEvent:
        err = o.DeletedHandler(event.Data().(*Deleted), entity.(*Grade))
    case UpdateEvent:
        err = o.UpdateHandler(event.Data().(*Update), entity.(*Grade))
    case UpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*Updated), entity.(*Grade))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *EventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(CreateEvent, func() eventhorizon.EventData {
		return &Create{}
	})

    //default handler implementation
    o.CreateHandler = func(event *Create, entity *Grade) (err error) {
        //err = eh.EventHandlerNotImplemented(CreateEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(CreatedEvent, func() eventhorizon.EventData {
		return &Created{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *Created, entity *Grade) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, GradeAggregateType); err == nil {
            entity.Student = event.Student
            entity.Course = event.Course
            entity.Grade = event.Grade
            entity.Comment = event.Comment
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DeleteEvent, func() eventhorizon.EventData {
		return &Delete{}
	})

    //default handler implementation
    o.DeleteHandler = func(event *Delete, entity *Grade) (err error) {
        //err = eh.EventHandlerNotImplemented(DeleteEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DeletedEvent, func() eventhorizon.EventData {
		return &Deleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *Deleted, entity *Grade) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, GradeAggregateType); err == nil {
            *entity = *NewGrade()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdateEvent, func() eventhorizon.EventData {
		return &Update{}
	})

    //default handler implementation
    o.UpdateHandler = func(event *Update, entity *Grade) (err error) {
        //err = eh.EventHandlerNotImplemented(UpdateEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdatedEvent, func() eventhorizon.EventData {
		return &Updated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *Updated, entity *Grade) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, GradeAggregateType); err == nil {
            entity.Student = event.Student
            entity.Course = event.Course
            entity.Grade = event.Grade
            entity.Comment = event.Comment
        }
        return
    }
    return
}


const GradeAggregateType eventhorizon.AggregateType = "Grade"

type AggregateInitializer struct {
    *eh.AggregateInitializer
    *CommandHandler
    *EventHandler
    ProjectorHandler *EventHandler `json:"projectorHandler" eh:"optional"`
}



func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *AggregateInitializer) {
    
    commandHandler := &GradeCommandHandler{}
    eventHandler := &GradeEventHandler{}
    entityFactory := func() eventhorizon.Entity { return NewGrade() }
    ret = &AggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(GradeAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(GradeAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        GradeCommandTypes().Literals(), GradeEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), GradeCommandHandler: commandHandler, GradeEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type CommandHandler struct {
    CreateHandler func (*Create, *Group, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*Delete, *Group, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*Update, *Group, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *CommandHandler) AddCreatePreparer(preparer func (*Create, *Group) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *Create, entity *Group, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddDeletePreparer(preparer func (*Delete, *Group) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *Delete, entity *Group, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddUpdatePreparer(preparer func (*Update, *Group) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *Update, entity *Group, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateCommand:
        err = o.CreateHandler(cmd.(*Create), entity.(*Group), store)
    case DeleteCommand:
        err = o.DeleteHandler(cmd.(*Delete), entity.(*Group), store)
    case UpdateCommand:
        err = o.UpdateHandler(cmd.(*Update), entity.(*Group), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *CommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *Create, entity *Group, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, GroupAggregateType); err == nil {
            store.StoreEvent(createdEvent, &Created{
                Name: command.Name,
                Category: command.Category,
                SchoolYear: command.SchoolYear,
                Representative: command.Representative,
                Students: command.Students,
                Courses: command.Courses,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *Delete, entity *Group, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, GroupAggregateType); err == nil {
            store.StoreEvent(deletedEvent, &Deleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *Update, entity *Group, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, GroupAggregateType); err == nil {
            store.StoreEvent(updatedEvent, &Updated{
                Name: command.Name,
                Category: command.Category,
                SchoolYear: command.SchoolYear,
                Representative: command.Representative,
                Students: command.Students,
                Courses: command.Courses,
                Id: command.Id,}, time.Now())
        }
        return
    }
    return
}


type EventHandler struct {
    CreateHandler func (*Create, *Group) (err error)  `json:"createHandler" eh:"optional"`
    CreatedHandler func (*Created, *Group) (err error)  `json:"createdHandler" eh:"optional"`
    DeleteHandler func (*Delete, *Group) (err error)  `json:"deleteHandler" eh:"optional"`
    DeletedHandler func (*Deleted, *Group) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdateHandler func (*Update, *Group) (err error)  `json:"updateHandler" eh:"optional"`
    UpdatedHandler func (*Updated, *Group) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *EventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case CreateEvent:
        err = o.CreateHandler(event.Data().(*Create), entity.(*Group))
    case CreatedEvent:
        err = o.CreatedHandler(event.Data().(*Created), entity.(*Group))
    case DeleteEvent:
        err = o.DeleteHandler(event.Data().(*Delete), entity.(*Group))
    case DeletedEvent:
        err = o.DeletedHandler(event.Data().(*Deleted), entity.(*Group))
    case UpdateEvent:
        err = o.UpdateHandler(event.Data().(*Update), entity.(*Group))
    case UpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*Updated), entity.(*Group))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *EventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(CreateEvent, func() eventhorizon.EventData {
		return &Create{}
	})

    //default handler implementation
    o.CreateHandler = func(event *Create, entity *Group) (err error) {
        //err = eh.EventHandlerNotImplemented(CreateEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(CreatedEvent, func() eventhorizon.EventData {
		return &Created{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *Created, entity *Group) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, GroupAggregateType); err == nil {
            entity.Name = event.Name
            entity.Category = event.Category
            entity.SchoolYear = event.SchoolYear
            entity.Representative = event.Representative
            entity.Students = event.Students
            entity.Courses = event.Courses
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DeleteEvent, func() eventhorizon.EventData {
		return &Delete{}
	})

    //default handler implementation
    o.DeleteHandler = func(event *Delete, entity *Group) (err error) {
        //err = eh.EventHandlerNotImplemented(DeleteEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DeletedEvent, func() eventhorizon.EventData {
		return &Deleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *Deleted, entity *Group) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, GroupAggregateType); err == nil {
            *entity = *NewGroup()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdateEvent, func() eventhorizon.EventData {
		return &Update{}
	})

    //default handler implementation
    o.UpdateHandler = func(event *Update, entity *Group) (err error) {
        //err = eh.EventHandlerNotImplemented(UpdateEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdatedEvent, func() eventhorizon.EventData {
		return &Updated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *Updated, entity *Group) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, GroupAggregateType); err == nil {
            entity.Name = event.Name
            entity.Category = event.Category
            entity.SchoolYear = event.SchoolYear
            entity.Representative = event.Representative
            entity.Students = event.Students
            entity.Courses = event.Courses
        }
        return
    }
    return
}


const GroupAggregateType eventhorizon.AggregateType = "Group"

type AggregateInitializer struct {
    *eh.AggregateInitializer
    *CommandHandler
    *EventHandler
    ProjectorHandler *EventHandler `json:"projectorHandler" eh:"optional"`
}



func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *AggregateInitializer) {
    
    commandHandler := &GroupCommandHandler{}
    eventHandler := &GroupEventHandler{}
    entityFactory := func() eventhorizon.Entity { return NewGroup() }
    ret = &AggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(GroupAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(GroupAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        GroupCommandTypes().Literals(), GroupEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), GroupCommandHandler: commandHandler, GroupEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type CommandHandler struct {
    CreateHandler func (*Create, *SchoolApplication, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*Delete, *SchoolApplication, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*Update, *SchoolApplication, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *CommandHandler) AddCreatePreparer(preparer func (*Create, *SchoolApplication) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *Create, entity *SchoolApplication, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddDeletePreparer(preparer func (*Delete, *SchoolApplication) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *Delete, entity *SchoolApplication, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddUpdatePreparer(preparer func (*Update, *SchoolApplication) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *Update, entity *SchoolApplication, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateCommand:
        err = o.CreateHandler(cmd.(*Create), entity.(*SchoolApplication), store)
    case DeleteCommand:
        err = o.DeleteHandler(cmd.(*Delete), entity.(*SchoolApplication), store)
    case UpdateCommand:
        err = o.UpdateHandler(cmd.(*Update), entity.(*SchoolApplication), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *CommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *Create, entity *SchoolApplication, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, SchoolApplicationAggregateType); err == nil {
            store.StoreEvent(createdEvent, &Created{
                Profile: command.Profile,
                ChurchContactPerson: command.ChurchContactPerson,
                ChurchContact: command.ChurchContact,
                ChurchCommitment: command.ChurchCommitment,
                SchoolYear: command.SchoolYear,
                Group: command.Group,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *Delete, entity *SchoolApplication, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, SchoolApplicationAggregateType); err == nil {
            store.StoreEvent(deletedEvent, &Deleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *Update, entity *SchoolApplication, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, SchoolApplicationAggregateType); err == nil {
            store.StoreEvent(updatedEvent, &Updated{
                Profile: command.Profile,
                ChurchContactPerson: command.ChurchContactPerson,
                ChurchContact: command.ChurchContact,
                ChurchCommitment: command.ChurchCommitment,
                SchoolYear: command.SchoolYear,
                Group: command.Group,
                Id: command.Id,}, time.Now())
        }
        return
    }
    return
}


type EventHandler struct {
    CreateHandler func (*Create, *SchoolApplication) (err error)  `json:"createHandler" eh:"optional"`
    CreatedHandler func (*Created, *SchoolApplication) (err error)  `json:"createdHandler" eh:"optional"`
    DeleteHandler func (*Delete, *SchoolApplication) (err error)  `json:"deleteHandler" eh:"optional"`
    DeletedHandler func (*Deleted, *SchoolApplication) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdateHandler func (*Update, *SchoolApplication) (err error)  `json:"updateHandler" eh:"optional"`
    UpdatedHandler func (*Updated, *SchoolApplication) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *EventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case CreateEvent:
        err = o.CreateHandler(event.Data().(*Create), entity.(*SchoolApplication))
    case CreatedEvent:
        err = o.CreatedHandler(event.Data().(*Created), entity.(*SchoolApplication))
    case DeleteEvent:
        err = o.DeleteHandler(event.Data().(*Delete), entity.(*SchoolApplication))
    case DeletedEvent:
        err = o.DeletedHandler(event.Data().(*Deleted), entity.(*SchoolApplication))
    case UpdateEvent:
        err = o.UpdateHandler(event.Data().(*Update), entity.(*SchoolApplication))
    case UpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*Updated), entity.(*SchoolApplication))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *EventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(CreateEvent, func() eventhorizon.EventData {
		return &Create{}
	})

    //default handler implementation
    o.CreateHandler = func(event *Create, entity *SchoolApplication) (err error) {
        //err = eh.EventHandlerNotImplemented(CreateEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(CreatedEvent, func() eventhorizon.EventData {
		return &Created{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *Created, entity *SchoolApplication) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, SchoolApplicationAggregateType); err == nil {
            entity.Profile = event.Profile
            entity.ChurchContactPerson = event.ChurchContactPerson
            entity.ChurchContact = event.ChurchContact
            entity.ChurchCommitment = event.ChurchCommitment
            entity.SchoolYear = event.SchoolYear
            entity.Group = event.Group
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DeleteEvent, func() eventhorizon.EventData {
		return &Delete{}
	})

    //default handler implementation
    o.DeleteHandler = func(event *Delete, entity *SchoolApplication) (err error) {
        //err = eh.EventHandlerNotImplemented(DeleteEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DeletedEvent, func() eventhorizon.EventData {
		return &Deleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *Deleted, entity *SchoolApplication) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, SchoolApplicationAggregateType); err == nil {
            *entity = *NewSchoolApplication()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdateEvent, func() eventhorizon.EventData {
		return &Update{}
	})

    //default handler implementation
    o.UpdateHandler = func(event *Update, entity *SchoolApplication) (err error) {
        //err = eh.EventHandlerNotImplemented(UpdateEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdatedEvent, func() eventhorizon.EventData {
		return &Updated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *Updated, entity *SchoolApplication) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, SchoolApplicationAggregateType); err == nil {
            entity.Profile = event.Profile
            entity.ChurchContactPerson = event.ChurchContactPerson
            entity.ChurchContact = event.ChurchContact
            entity.ChurchCommitment = event.ChurchCommitment
            entity.SchoolYear = event.SchoolYear
            entity.Group = event.Group
        }
        return
    }
    return
}


const SchoolApplicationAggregateType eventhorizon.AggregateType = "SchoolApplication"

type AggregateInitializer struct {
    *eh.AggregateInitializer
    *CommandHandler
    *EventHandler
    ProjectorHandler *EventHandler `json:"projectorHandler" eh:"optional"`
}



func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *AggregateInitializer) {
    
    commandHandler := &SchoolApplicationCommandHandler{}
    eventHandler := &SchoolApplicationEventHandler{}
    entityFactory := func() eventhorizon.Entity { return NewSchoolApplication() }
    ret = &AggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(SchoolApplicationAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(SchoolApplicationAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        SchoolApplicationCommandTypes().Literals(), SchoolApplicationEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), SchoolApplicationCommandHandler: commandHandler, SchoolApplicationEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type CommandHandler struct {
    CreateHandler func (*Create, *SchoolYear, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*Delete, *SchoolYear, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*Update, *SchoolYear, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *CommandHandler) AddCreatePreparer(preparer func (*Create, *SchoolYear) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *Create, entity *SchoolYear, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddDeletePreparer(preparer func (*Delete, *SchoolYear) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *Delete, entity *SchoolYear, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddUpdatePreparer(preparer func (*Update, *SchoolYear) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *Update, entity *SchoolYear, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateCommand:
        err = o.CreateHandler(cmd.(*Create), entity.(*SchoolYear), store)
    case DeleteCommand:
        err = o.DeleteHandler(cmd.(*Delete), entity.(*SchoolYear), store)
    case UpdateCommand:
        err = o.UpdateHandler(cmd.(*Update), entity.(*SchoolYear), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *CommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *Create, entity *SchoolYear, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, SchoolYearAggregateType); err == nil {
            store.StoreEvent(createdEvent, &Created{
                Name: command.Name,
                Start: command.Start,
                End: command.End,
                Dates: command.Dates,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *Delete, entity *SchoolYear, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, SchoolYearAggregateType); err == nil {
            store.StoreEvent(deletedEvent, &Deleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *Update, entity *SchoolYear, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, SchoolYearAggregateType); err == nil {
            store.StoreEvent(updatedEvent, &Updated{
                Name: command.Name,
                Start: command.Start,
                End: command.End,
                Dates: command.Dates,
                Id: command.Id,}, time.Now())
        }
        return
    }
    return
}


type EventHandler struct {
    CreateHandler func (*Create, *SchoolYear) (err error)  `json:"createHandler" eh:"optional"`
    CreatedHandler func (*Created, *SchoolYear) (err error)  `json:"createdHandler" eh:"optional"`
    DeleteHandler func (*Delete, *SchoolYear) (err error)  `json:"deleteHandler" eh:"optional"`
    DeletedHandler func (*Deleted, *SchoolYear) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdateHandler func (*Update, *SchoolYear) (err error)  `json:"updateHandler" eh:"optional"`
    UpdatedHandler func (*Updated, *SchoolYear) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *EventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case CreateEvent:
        err = o.CreateHandler(event.Data().(*Create), entity.(*SchoolYear))
    case CreatedEvent:
        err = o.CreatedHandler(event.Data().(*Created), entity.(*SchoolYear))
    case DeleteEvent:
        err = o.DeleteHandler(event.Data().(*Delete), entity.(*SchoolYear))
    case DeletedEvent:
        err = o.DeletedHandler(event.Data().(*Deleted), entity.(*SchoolYear))
    case UpdateEvent:
        err = o.UpdateHandler(event.Data().(*Update), entity.(*SchoolYear))
    case UpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*Updated), entity.(*SchoolYear))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *EventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(CreateEvent, func() eventhorizon.EventData {
		return &Create{}
	})

    //default handler implementation
    o.CreateHandler = func(event *Create, entity *SchoolYear) (err error) {
        //err = eh.EventHandlerNotImplemented(CreateEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(CreatedEvent, func() eventhorizon.EventData {
		return &Created{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *Created, entity *SchoolYear) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, SchoolYearAggregateType); err == nil {
            entity.Name = event.Name
            entity.Start = event.Start
            entity.End = event.End
            entity.Dates = event.Dates
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DeleteEvent, func() eventhorizon.EventData {
		return &Delete{}
	})

    //default handler implementation
    o.DeleteHandler = func(event *Delete, entity *SchoolYear) (err error) {
        //err = eh.EventHandlerNotImplemented(DeleteEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(DeletedEvent, func() eventhorizon.EventData {
		return &Deleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *Deleted, entity *SchoolYear) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, SchoolYearAggregateType); err == nil {
            *entity = *NewSchoolYear()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdateEvent, func() eventhorizon.EventData {
		return &Update{}
	})

    //default handler implementation
    o.UpdateHandler = func(event *Update, entity *SchoolYear) (err error) {
        //err = eh.EventHandlerNotImplemented(UpdateEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(UpdatedEvent, func() eventhorizon.EventData {
		return &Updated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *Updated, entity *SchoolYear) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, SchoolYearAggregateType); err == nil {
            entity.Name = event.Name
            entity.Start = event.Start
            entity.End = event.End
            entity.Dates = event.Dates
        }
        return
    }
    return
}


const SchoolYearAggregateType eventhorizon.AggregateType = "SchoolYear"

type AggregateInitializer struct {
    *eh.AggregateInitializer
    *CommandHandler
    *EventHandler
    ProjectorHandler *EventHandler `json:"projectorHandler" eh:"optional"`
}



func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *AggregateInitializer) {
    
    commandHandler := &SchoolYearCommandHandler{}
    eventHandler := &SchoolYearEventHandler{}
    entityFactory := func() eventhorizon.Entity { return NewSchoolYear() }
    ret = &AggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(SchoolYearAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(SchoolYearAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        SchoolYearCommandTypes().Literals(), SchoolYearEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), SchoolYearCommandHandler: commandHandler, SchoolYearEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type StudentEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore `json:"eventStore" eh:"optional"`
    eventBus eventhorizon.EventBus `json:"eventBus" eh:"optional"`
    eventPublisher eventhorizon.EventPublisher `json:"eventPublisher" eh:"optional"`
    commandBus *bus.CommandHandler `json:"commandBus" eh:"optional"`
    AttendanceAggregateInitializer *AggregateInitializer `json:"attendanceAggregateInitializer" eh:"optional"`
    CourseAggregateInitializer *AggregateInitializer `json:"courseAggregateInitializer" eh:"optional"`
    GradeAggregateInitializer *AggregateInitializer `json:"gradeAggregateInitializer" eh:"optional"`
    GroupAggregateInitializer *AggregateInitializer `json:"groupAggregateInitializer" eh:"optional"`
    SchoolApplicationAggregateInitializer *AggregateInitializer `json:"schoolApplicationAggregateInitializer" eh:"optional"`
    SchoolYearAggregateInitializer *AggregateInitializer `json:"schoolYearAggregateInitializer" eh:"optional"`
}

func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *StudentEventhorizonInitializer) {
    attendanceAggregateInitializer := New@@EMPTY@@(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    courseAggregateInitializer := New@@EMPTY@@(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    gradeAggregateInitializer := New@@EMPTY@@(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    groupAggregateInitializer := New@@EMPTY@@(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    schoolApplicationAggregateInitializer := New@@EMPTY@@(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    schoolYearAggregateInitializer := New@@EMPTY@@(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    ret = &StudentEventhorizonInitializer{
        eventStore: eventStore,
        eventBus: eventBus,
        eventPublisher: eventPublisher,
        commandBus: commandBus,
        AttendanceAggregateInitializer: attendanceAggregateInitializer,
        CourseAggregateInitializer: courseAggregateInitializer,
        GradeAggregateInitializer: gradeAggregateInitializer,
        GroupAggregateInitializer: groupAggregateInitializer,
        SchoolApplicationAggregateInitializer: schoolApplicationAggregateInitializer,
        SchoolYearAggregateInitializer: schoolYearAggregateInitializer,
    }
    return
}

func (o *StudentEventhorizonInitializer) Setup() (err error) {
    
    if err = o.AttendanceAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.CourseAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.GradeAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.GroupAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.SchoolApplicationAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.SchoolYearAggregateInitializer.Setup(); err != nil {
        return
    }

    return
}









