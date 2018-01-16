package student

import (
    "errors"
    "fmt"
    "github.com/eugeis/gee/eh"
    "github.com/looplab/eventhorizon"
    "github.com/looplab/eventhorizon/commandhandler/bus"
    "time"
)
type AttendanceCommandHandler struct {
    RegisterHandler func (*Register, *Attendance, eh.AggregateStoreEvent) (err error)  `json:"registerHandler" eh:"optional"`
    CreateHandler func (*Create, *Attendance, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*Delete, *Attendance, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    CancelHandler func (*Cancel, *Attendance, eh.AggregateStoreEvent) (err error)  `json:"cancelHandler" eh:"optional"`
    ConfirmHandler func (*Confirm, *Attendance, eh.AggregateStoreEvent) (err error)  `json:"confirmHandler" eh:"optional"`
    UpdateHandler func (*Update, *Attendance, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *AttendanceCommandHandler) AddRegisterPreparer(preparer func (*Register, *Attendance) (err error) ) {
    prevHandler := o.RegisterHandler
	o.RegisterHandler = func(command *Register, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AttendanceCommandHandler) AddCreatePreparer(preparer func (*Create, *Attendance) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *Create, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AttendanceCommandHandler) AddDeletePreparer(preparer func (*Delete, *Attendance) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *Delete, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AttendanceCommandHandler) AddCancelPreparer(preparer func (*Cancel, *Attendance) (err error) ) {
    prevHandler := o.CancelHandler
	o.CancelHandler = func(command *Cancel, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AttendanceCommandHandler) AddConfirmPreparer(preparer func (*Confirm, *Attendance) (err error) ) {
    prevHandler := o.ConfirmHandler
	o.ConfirmHandler = func(command *Confirm, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AttendanceCommandHandler) AddUpdatePreparer(preparer func (*Update, *Attendance) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *Update, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AttendanceCommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case RegisterAttendanceCommand:
        err = o.RegisterHandler(cmd.(*Register), entity.(*Attendance), store)
    case CreateAttendanceCommand:
        err = o.CreateHandler(cmd.(*Create), entity.(*Attendance), store)
    case DeleteAttendanceCommand:
        err = o.DeleteHandler(cmd.(*Delete), entity.(*Attendance), store)
    case CancelAttendanceCommand:
        err = o.CancelHandler(cmd.(*Cancel), entity.(*Attendance), store)
    case ConfirmAttendanceCommand:
        err = o.ConfirmHandler(cmd.(*Confirm), entity.(*Attendance), store)
    case UpdateAttendanceCommand:
        err = o.UpdateHandler(cmd.(*Update), entity.(*Attendance), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *AttendanceCommandHandler) SetupCommandHandler() (err error) {
    o.RegisterHandler = func(command *Register, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, AttendanceAggregateType); err == nil {
            store.StoreEvent(AttendanceRegisteredEvent, &Registered{
                Student: command.Student,
                Course: command.Course,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.CreateHandler = func(command *Create, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, AttendanceAggregateType); err == nil {
            store.StoreEvent(AttendanceCreatedEvent, &Created{
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
            store.StoreEvent(AttendanceDeletedEvent, &Deleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.CancelHandler = func(command *Cancel, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AttendanceAggregateType); err == nil {
            store.StoreEvent(AttendanceCanceledEvent, &Canceled{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.ConfirmHandler = func(command *Confirm, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AttendanceAggregateType); err == nil {
            store.StoreEvent(AttendanceConfirmedEvent, &Confirmed{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *Update, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AttendanceAggregateType); err == nil {
            store.StoreEvent(AttendanceUpdatedEvent, &Updated{
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


type AttendanceEventHandler struct {
    CreatedHandler func (*Created, *Attendance) (err error)  `json:"createdHandler" eh:"optional"`
    RegisteredHandler func (*Registered, *Attendance) (err error)  `json:"registeredHandler" eh:"optional"`
    DeletedHandler func (*Deleted, *Attendance) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdatedHandler func (*Updated, *Attendance) (err error)  `json:"updatedHandler" eh:"optional"`
    ConfirmedHandler func (*Confirmed, *Attendance) (err error)  `json:"confirmedHandler" eh:"optional"`
    CanceledHandler func (*Canceled, *Attendance) (err error)  `json:"canceledHandler" eh:"optional"`
}

func (o *AttendanceEventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case AttendanceCreatedEvent:
        err = o.CreatedHandler(event.Data().(*Created), entity.(*Attendance))
    case AttendanceRegisteredEvent:
        err = o.RegisteredHandler(event.Data().(*Registered), entity.(*Attendance))
    case AttendanceDeletedEvent:
        err = o.DeletedHandler(event.Data().(*Deleted), entity.(*Attendance))
    case AttendanceUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*Updated), entity.(*Attendance))
    case AttendanceConfirmedEvent:
        err = o.ConfirmedHandler(event.Data().(*Confirmed), entity.(*Attendance))
    case AttendanceCanceledEvent:
        err = o.CanceledHandler(event.Data().(*Canceled), entity.(*Attendance))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *AttendanceEventHandler) SetupEventHandler() (err error) {

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
    eventhorizon.RegisterEventData(DeletedEvent, func() eventhorizon.EventData {
		return &Deleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *Deleted, entity *Attendance) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AttendanceAggregateType); err == nil {
            *entity = *New@@EMPTY@@()
        }
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

type AttendanceAggregateInitializer struct {
    *eh.AggregateInitializer
    *AttendanceCommandHandler
    *AttendanceEventHandler
    ProjectorHandler *AttendanceEventHandler `json:"projectorHandler" eh:"optional"`
}



func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *AttendanceAggregateInitializer) {
    
    commandHandler := &AttendanceCommandHandler{}
    eventHandler := &AttendanceEventHandler{}
    entityFactory := func() eventhorizon.Entity { return New@@EMPTY@@() }
    ret = &AttendanceAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(AttendanceAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(AttendanceAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        AttendanceCommandTypes().Literals(), AttendanceEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), AttendanceCommandHandler: commandHandler, AttendanceEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type CourseCommandHandler struct {
    CreateHandler func (*Create, *Course, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*Delete, *Course, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*Update, *Course, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *CourseCommandHandler) AddCreatePreparer(preparer func (*Create, *Course) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *Create, entity *Course, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CourseCommandHandler) AddDeletePreparer(preparer func (*Delete, *Course) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *Delete, entity *Course, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CourseCommandHandler) AddUpdatePreparer(preparer func (*Update, *Course) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *Update, entity *Course, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CourseCommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateCourseCommand:
        err = o.CreateHandler(cmd.(*Create), entity.(*Course), store)
    case DeleteCourseCommand:
        err = o.DeleteHandler(cmd.(*Delete), entity.(*Course), store)
    case UpdateCourseCommand:
        err = o.UpdateHandler(cmd.(*Update), entity.(*Course), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *CourseCommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *Create, entity *Course, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, CourseAggregateType); err == nil {
            store.StoreEvent(CourseCreatedEvent, &Created{
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
            store.StoreEvent(CourseDeletedEvent, &Deleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *Update, entity *Course, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, CourseAggregateType); err == nil {
            store.StoreEvent(CourseUpdatedEvent, &Updated{
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


type CourseEventHandler struct {
    CreatedHandler func (*Created, *Course) (err error)  `json:"createdHandler" eh:"optional"`
    DeletedHandler func (*Deleted, *Course) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdatedHandler func (*Updated, *Course) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *CourseEventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case CourseCreatedEvent:
        err = o.CreatedHandler(event.Data().(*Created), entity.(*Course))
    case CourseDeletedEvent:
        err = o.DeletedHandler(event.Data().(*Deleted), entity.(*Course))
    case CourseUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*Updated), entity.(*Course))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *CourseEventHandler) SetupEventHandler() (err error) {

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
    eventhorizon.RegisterEventData(DeletedEvent, func() eventhorizon.EventData {
		return &Deleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *Deleted, entity *Course) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, CourseAggregateType); err == nil {
            *entity = *New@@EMPTY@@()
        }
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

type CourseAggregateInitializer struct {
    *eh.AggregateInitializer
    *CourseCommandHandler
    *CourseEventHandler
    ProjectorHandler *CourseEventHandler `json:"projectorHandler" eh:"optional"`
}



func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *CourseAggregateInitializer) {
    
    commandHandler := &CourseCommandHandler{}
    eventHandler := &CourseEventHandler{}
    entityFactory := func() eventhorizon.Entity { return New@@EMPTY@@() }
    ret = &CourseAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(CourseAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(CourseAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        CourseCommandTypes().Literals(), CourseEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), CourseCommandHandler: commandHandler, CourseEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type GradeCommandHandler struct {
    CreateHandler func (*Create, *Grade, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*Delete, *Grade, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*Update, *Grade, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *GradeCommandHandler) AddCreatePreparer(preparer func (*Create, *Grade) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *Create, entity *Grade, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *GradeCommandHandler) AddDeletePreparer(preparer func (*Delete, *Grade) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *Delete, entity *Grade, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *GradeCommandHandler) AddUpdatePreparer(preparer func (*Update, *Grade) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *Update, entity *Grade, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *GradeCommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateGradeCommand:
        err = o.CreateHandler(cmd.(*Create), entity.(*Grade), store)
    case DeleteGradeCommand:
        err = o.DeleteHandler(cmd.(*Delete), entity.(*Grade), store)
    case UpdateGradeCommand:
        err = o.UpdateHandler(cmd.(*Update), entity.(*Grade), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *GradeCommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *Create, entity *Grade, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, GradeAggregateType); err == nil {
            store.StoreEvent(GradeCreatedEvent, &Created{
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
            store.StoreEvent(GradeDeletedEvent, &Deleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *Update, entity *Grade, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, GradeAggregateType); err == nil {
            store.StoreEvent(GradeUpdatedEvent, &Updated{
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


type GradeEventHandler struct {
    CreatedHandler func (*Created, *Grade) (err error)  `json:"createdHandler" eh:"optional"`
    DeletedHandler func (*Deleted, *Grade) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdatedHandler func (*Updated, *Grade) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *GradeEventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case GradeCreatedEvent:
        err = o.CreatedHandler(event.Data().(*Created), entity.(*Grade))
    case GradeDeletedEvent:
        err = o.DeletedHandler(event.Data().(*Deleted), entity.(*Grade))
    case GradeUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*Updated), entity.(*Grade))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *GradeEventHandler) SetupEventHandler() (err error) {

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
    eventhorizon.RegisterEventData(DeletedEvent, func() eventhorizon.EventData {
		return &Deleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *Deleted, entity *Grade) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, GradeAggregateType); err == nil {
            *entity = *New@@EMPTY@@()
        }
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

type GradeAggregateInitializer struct {
    *eh.AggregateInitializer
    *GradeCommandHandler
    *GradeEventHandler
    ProjectorHandler *GradeEventHandler `json:"projectorHandler" eh:"optional"`
}



func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *GradeAggregateInitializer) {
    
    commandHandler := &GradeCommandHandler{}
    eventHandler := &GradeEventHandler{}
    entityFactory := func() eventhorizon.Entity { return New@@EMPTY@@() }
    ret = &GradeAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(GradeAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(GradeAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        GradeCommandTypes().Literals(), GradeEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), GradeCommandHandler: commandHandler, GradeEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type GroupCommandHandler struct {
    CreateHandler func (*Create, *Group, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*Delete, *Group, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*Update, *Group, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *GroupCommandHandler) AddCreatePreparer(preparer func (*Create, *Group) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *Create, entity *Group, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *GroupCommandHandler) AddDeletePreparer(preparer func (*Delete, *Group) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *Delete, entity *Group, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *GroupCommandHandler) AddUpdatePreparer(preparer func (*Update, *Group) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *Update, entity *Group, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *GroupCommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateGroupCommand:
        err = o.CreateHandler(cmd.(*Create), entity.(*Group), store)
    case DeleteGroupCommand:
        err = o.DeleteHandler(cmd.(*Delete), entity.(*Group), store)
    case UpdateGroupCommand:
        err = o.UpdateHandler(cmd.(*Update), entity.(*Group), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *GroupCommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *Create, entity *Group, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, GroupAggregateType); err == nil {
            store.StoreEvent(GroupCreatedEvent, &Created{
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
            store.StoreEvent(GroupDeletedEvent, &Deleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *Update, entity *Group, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, GroupAggregateType); err == nil {
            store.StoreEvent(GroupUpdatedEvent, &Updated{
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


type GroupEventHandler struct {
    CreatedHandler func (*Created, *Group) (err error)  `json:"createdHandler" eh:"optional"`
    DeletedHandler func (*Deleted, *Group) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdatedHandler func (*Updated, *Group) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *GroupEventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case GroupCreatedEvent:
        err = o.CreatedHandler(event.Data().(*Created), entity.(*Group))
    case GroupDeletedEvent:
        err = o.DeletedHandler(event.Data().(*Deleted), entity.(*Group))
    case GroupUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*Updated), entity.(*Group))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *GroupEventHandler) SetupEventHandler() (err error) {

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
    eventhorizon.RegisterEventData(DeletedEvent, func() eventhorizon.EventData {
		return &Deleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *Deleted, entity *Group) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, GroupAggregateType); err == nil {
            *entity = *New@@EMPTY@@()
        }
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

type GroupAggregateInitializer struct {
    *eh.AggregateInitializer
    *GroupCommandHandler
    *GroupEventHandler
    ProjectorHandler *GroupEventHandler `json:"projectorHandler" eh:"optional"`
}



func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *GroupAggregateInitializer) {
    
    commandHandler := &GroupCommandHandler{}
    eventHandler := &GroupEventHandler{}
    entityFactory := func() eventhorizon.Entity { return New@@EMPTY@@() }
    ret = &GroupAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(GroupAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(GroupAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        GroupCommandTypes().Literals(), GroupEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), GroupCommandHandler: commandHandler, GroupEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type SchoolApplicationCommandHandler struct {
    CreateHandler func (*Create, *SchoolApplication, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*Delete, *SchoolApplication, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*Update, *SchoolApplication, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *SchoolApplicationCommandHandler) AddCreatePreparer(preparer func (*Create, *SchoolApplication) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *Create, entity *SchoolApplication, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *SchoolApplicationCommandHandler) AddDeletePreparer(preparer func (*Delete, *SchoolApplication) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *Delete, entity *SchoolApplication, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *SchoolApplicationCommandHandler) AddUpdatePreparer(preparer func (*Update, *SchoolApplication) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *Update, entity *SchoolApplication, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *SchoolApplicationCommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateSchoolApplicationCommand:
        err = o.CreateHandler(cmd.(*Create), entity.(*SchoolApplication), store)
    case DeleteSchoolApplicationCommand:
        err = o.DeleteHandler(cmd.(*Delete), entity.(*SchoolApplication), store)
    case UpdateSchoolApplicationCommand:
        err = o.UpdateHandler(cmd.(*Update), entity.(*SchoolApplication), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *SchoolApplicationCommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *Create, entity *SchoolApplication, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, SchoolApplicationAggregateType); err == nil {
            store.StoreEvent(SchoolApplicationCreatedEvent, &Created{
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
            store.StoreEvent(SchoolApplicationDeletedEvent, &Deleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *Update, entity *SchoolApplication, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, SchoolApplicationAggregateType); err == nil {
            store.StoreEvent(SchoolApplicationUpdatedEvent, &Updated{
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


type SchoolApplicationEventHandler struct {
    CreatedHandler func (*Created, *SchoolApplication) (err error)  `json:"createdHandler" eh:"optional"`
    DeletedHandler func (*Deleted, *SchoolApplication) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdatedHandler func (*Updated, *SchoolApplication) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *SchoolApplicationEventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case SchoolApplicationCreatedEvent:
        err = o.CreatedHandler(event.Data().(*Created), entity.(*SchoolApplication))
    case SchoolApplicationDeletedEvent:
        err = o.DeletedHandler(event.Data().(*Deleted), entity.(*SchoolApplication))
    case SchoolApplicationUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*Updated), entity.(*SchoolApplication))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *SchoolApplicationEventHandler) SetupEventHandler() (err error) {

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
    eventhorizon.RegisterEventData(DeletedEvent, func() eventhorizon.EventData {
		return &Deleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *Deleted, entity *SchoolApplication) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, SchoolApplicationAggregateType); err == nil {
            *entity = *New@@EMPTY@@()
        }
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

type SchoolApplicationAggregateInitializer struct {
    *eh.AggregateInitializer
    *SchoolApplicationCommandHandler
    *SchoolApplicationEventHandler
    ProjectorHandler *SchoolApplicationEventHandler `json:"projectorHandler" eh:"optional"`
}



func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *SchoolApplicationAggregateInitializer) {
    
    commandHandler := &SchoolApplicationCommandHandler{}
    eventHandler := &SchoolApplicationEventHandler{}
    entityFactory := func() eventhorizon.Entity { return New@@EMPTY@@() }
    ret = &SchoolApplicationAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(SchoolApplicationAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(SchoolApplicationAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        SchoolApplicationCommandTypes().Literals(), SchoolApplicationEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), SchoolApplicationCommandHandler: commandHandler, SchoolApplicationEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type SchoolYearCommandHandler struct {
    CreateHandler func (*Create, *SchoolYear, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*Delete, *SchoolYear, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*Update, *SchoolYear, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *SchoolYearCommandHandler) AddCreatePreparer(preparer func (*Create, *SchoolYear) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *Create, entity *SchoolYear, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *SchoolYearCommandHandler) AddDeletePreparer(preparer func (*Delete, *SchoolYear) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *Delete, entity *SchoolYear, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *SchoolYearCommandHandler) AddUpdatePreparer(preparer func (*Update, *SchoolYear) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *Update, entity *SchoolYear, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *SchoolYearCommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateSchoolYearCommand:
        err = o.CreateHandler(cmd.(*Create), entity.(*SchoolYear), store)
    case DeleteSchoolYearCommand:
        err = o.DeleteHandler(cmd.(*Delete), entity.(*SchoolYear), store)
    case UpdateSchoolYearCommand:
        err = o.UpdateHandler(cmd.(*Update), entity.(*SchoolYear), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *SchoolYearCommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *Create, entity *SchoolYear, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, SchoolYearAggregateType); err == nil {
            store.StoreEvent(SchoolYearCreatedEvent, &Created{
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
            store.StoreEvent(SchoolYearDeletedEvent, &Deleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *Update, entity *SchoolYear, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, SchoolYearAggregateType); err == nil {
            store.StoreEvent(SchoolYearUpdatedEvent, &Updated{
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


type SchoolYearEventHandler struct {
    CreatedHandler func (*Created, *SchoolYear) (err error)  `json:"createdHandler" eh:"optional"`
    DeletedHandler func (*Deleted, *SchoolYear) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdatedHandler func (*Updated, *SchoolYear) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *SchoolYearEventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case SchoolYearCreatedEvent:
        err = o.CreatedHandler(event.Data().(*Created), entity.(*SchoolYear))
    case SchoolYearDeletedEvent:
        err = o.DeletedHandler(event.Data().(*Deleted), entity.(*SchoolYear))
    case SchoolYearUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*Updated), entity.(*SchoolYear))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *SchoolYearEventHandler) SetupEventHandler() (err error) {

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
    eventhorizon.RegisterEventData(DeletedEvent, func() eventhorizon.EventData {
		return &Deleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *Deleted, entity *SchoolYear) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, SchoolYearAggregateType); err == nil {
            *entity = *New@@EMPTY@@()
        }
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

type SchoolYearAggregateInitializer struct {
    *eh.AggregateInitializer
    *SchoolYearCommandHandler
    *SchoolYearEventHandler
    ProjectorHandler *SchoolYearEventHandler `json:"projectorHandler" eh:"optional"`
}



func New@@EMPTY@@(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *SchoolYearAggregateInitializer) {
    
    commandHandler := &SchoolYearCommandHandler{}
    eventHandler := &SchoolYearEventHandler{}
    entityFactory := func() eventhorizon.Entity { return New@@EMPTY@@() }
    ret = &SchoolYearAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(SchoolYearAggregateType,
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
    AttendanceAggregateInitializer *AttendanceAggregateInitializer `json:"attendanceAggregateInitializer" eh:"optional"`
    CourseAggregateInitializer *CourseAggregateInitializer `json:"courseAggregateInitializer" eh:"optional"`
    GradeAggregateInitializer *GradeAggregateInitializer `json:"gradeAggregateInitializer" eh:"optional"`
    GroupAggregateInitializer *GroupAggregateInitializer `json:"groupAggregateInitializer" eh:"optional"`
    SchoolApplicationAggregateInitializer *SchoolApplicationAggregateInitializer `json:"schoolApplicationAggregateInitializer" eh:"optional"`
    SchoolYearAggregateInitializer *SchoolYearAggregateInitializer `json:"schoolYearAggregateInitializer" eh:"optional"`
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









