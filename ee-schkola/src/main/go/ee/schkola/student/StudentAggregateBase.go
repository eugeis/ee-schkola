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
    RegisterHandler func (*RegisterAttendance, *Attendance, eh.AggregateStoreEvent) (err error)  `json:"registerHandler" eh:"optional"`
    CreateHandler func (*CreateAttendance, *Attendance, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*DeleteAttendance, *Attendance, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    CancelHandler func (*CancelAttendance, *Attendance, eh.AggregateStoreEvent) (err error)  `json:"cancelHandler" eh:"optional"`
    ConfirmHandler func (*ConfirmAttendance, *Attendance, eh.AggregateStoreEvent) (err error)  `json:"confirmHandler" eh:"optional"`
    UpdateHandler func (*UpdateAttendance, *Attendance, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *CommandHandler) AddRegisterPreparer(preparer func (*RegisterAttendance, *Attendance) (err error) ) {
    prevHandler := o.RegisterHandler
	o.RegisterHandler = func(command *RegisterAttendance, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddCreatePreparer(preparer func (*CreateAttendance, *Attendance) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *CreateAttendance, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddDeletePreparer(preparer func (*DeleteAttendance, *Attendance) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *DeleteAttendance, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddCancelPreparer(preparer func (*CancelAttendance, *Attendance) (err error) ) {
    prevHandler := o.CancelHandler
	o.CancelHandler = func(command *CancelAttendance, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddConfirmPreparer(preparer func (*ConfirmAttendance, *Attendance) (err error) ) {
    prevHandler := o.ConfirmHandler
	o.ConfirmHandler = func(command *ConfirmAttendance, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddUpdatePreparer(preparer func (*UpdateAttendance, *Attendance) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *UpdateAttendance, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case RegisterAttendanceCommand:
        err = o.RegisterHandler(cmd.(*RegisterAttendance), entity.(*Attendance), store)
    case CreateAttendanceCommand:
        err = o.CreateHandler(cmd.(*CreateAttendance), entity.(*Attendance), store)
    case DeleteAttendanceCommand:
        err = o.DeleteHandler(cmd.(*DeleteAttendance), entity.(*Attendance), store)
    case CancelAttendanceCommand:
        err = o.CancelHandler(cmd.(*CancelAttendance), entity.(*Attendance), store)
    case ConfirmAttendanceCommand:
        err = o.ConfirmHandler(cmd.(*ConfirmAttendance), entity.(*Attendance), store)
    case UpdateAttendanceCommand:
        err = o.UpdateHandler(cmd.(*UpdateAttendance), entity.(*Attendance), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *CommandHandler) SetupCommandHandler() (err error) {
    o.RegisterHandler = func(command *RegisterAttendance, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, AttendanceAggregateType); err == nil {
            store.StoreEvent(AttendanceRegisteredEvent, &AttendanceRegistered{
                Student: command.Student,
                Course: command.Course,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.CreateHandler = func(command *CreateAttendance, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, AttendanceAggregateType); err == nil {
            store.StoreEvent(AttendanceCreatedEvent, &AttendanceCreated{
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
    o.DeleteHandler = func(command *DeleteAttendance, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AttendanceAggregateType); err == nil {
            store.StoreEvent(AttendanceDeletedEvent, &AttendanceDeleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.CancelHandler = func(command *CancelAttendance, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AttendanceAggregateType); err == nil {
            store.StoreEvent(AttendanceCanceledEvent, &AttendanceCanceled{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.ConfirmHandler = func(command *ConfirmAttendance, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AttendanceAggregateType); err == nil {
            store.StoreEvent(AttendanceConfirmedEvent, &AttendanceConfirmed{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *UpdateAttendance, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AttendanceAggregateType); err == nil {
            store.StoreEvent(AttendanceUpdatedEvent, &AttendanceUpdated{
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
    CreatedHandler func (*AttendanceCreated, *Attendance) (err error)  `json:"createdHandler" eh:"optional"`
    RegisteredHandler func (*AttendanceRegistered, *Attendance) (err error)  `json:"registeredHandler" eh:"optional"`
    DeletedHandler func (*AttendanceDeleted, *Attendance) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdatedHandler func (*AttendanceUpdated, *Attendance) (err error)  `json:"updatedHandler" eh:"optional"`
    ConfirmedHandler func (*AttendanceConfirmed, *Attendance) (err error)  `json:"confirmedHandler" eh:"optional"`
    CanceledHandler func (*AttendanceCanceled, *Attendance) (err error)  `json:"canceledHandler" eh:"optional"`
}

func (o *EventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case AttendanceCreatedEvent:
        err = o.CreatedHandler(event.Data().(*AttendanceCreated), entity.(*Attendance))
    case AttendanceRegisteredEvent:
        err = o.RegisteredHandler(event.Data().(*AttendanceRegistered), entity.(*Attendance))
    case AttendanceDeletedEvent:
        err = o.DeletedHandler(event.Data().(*AttendanceDeleted), entity.(*Attendance))
    case AttendanceUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*AttendanceUpdated), entity.(*Attendance))
    case AttendanceConfirmedEvent:
        err = o.ConfirmedHandler(event.Data().(*AttendanceConfirmed), entity.(*Attendance))
    case AttendanceCanceledEvent:
        err = o.CanceledHandler(event.Data().(*AttendanceCanceled), entity.(*Attendance))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *EventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(AttendanceCreatedEvent, func() eventhorizon.EventData {
		return &AttendanceCreated{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *AttendanceCreated, entity *Attendance) (err error) {
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
    eventhorizon.RegisterEventData(AttendanceRegisteredEvent, func() eventhorizon.EventData {
		return &AttendanceRegistered{}
	})

    //default handler implementation
    o.RegisteredHandler = func(event *AttendanceRegistered, entity *Attendance) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, AttendanceAggregateType); err == nil {
            entity.Student = event.Student
            entity.Course = event.Course
            entity.State = AttendanceStates().Registered()
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(AttendanceDeletedEvent, func() eventhorizon.EventData {
		return &AttendanceDeleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *AttendanceDeleted, entity *Attendance) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AttendanceAggregateType); err == nil {
            *entity = *NewAttendance()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(AttendanceUpdatedEvent, func() eventhorizon.EventData {
		return &AttendanceUpdated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *AttendanceUpdated, entity *Attendance) (err error) {
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
    eventhorizon.RegisterEventData(AttendanceConfirmedEvent, func() eventhorizon.EventData {
		return &AttendanceConfirmed{}
	})

    //default handler implementation
    o.ConfirmedHandler = func(event *AttendanceConfirmed, entity *Attendance) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AttendanceAggregateType); err == nil {
            entity.State = AttendanceStates().Confirmed()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(AttendanceCanceledEvent, func() eventhorizon.EventData {
		return &AttendanceCanceled{}
	})

    //default handler implementation
    o.CanceledHandler = func(event *AttendanceCanceled, entity *Attendance) (err error) {
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



func NewAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
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
    CreateHandler func (*CreateCourse, *Course, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*DeleteCourse, *Course, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*UpdateCourse, *Course, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *CommandHandler) AddCreatePreparer(preparer func (*CreateCourse, *Course) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *CreateCourse, entity *Course, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddDeletePreparer(preparer func (*DeleteCourse, *Course) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *DeleteCourse, entity *Course, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddUpdatePreparer(preparer func (*UpdateCourse, *Course) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *UpdateCourse, entity *Course, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateCourseCommand:
        err = o.CreateHandler(cmd.(*CreateCourse), entity.(*Course), store)
    case DeleteCourseCommand:
        err = o.DeleteHandler(cmd.(*DeleteCourse), entity.(*Course), store)
    case UpdateCourseCommand:
        err = o.UpdateHandler(cmd.(*UpdateCourse), entity.(*Course), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *CommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *CreateCourse, entity *Course, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, CourseAggregateType); err == nil {
            store.StoreEvent(CourseCreatedEvent, &CourseCreated{
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
    o.DeleteHandler = func(command *DeleteCourse, entity *Course, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, CourseAggregateType); err == nil {
            store.StoreEvent(CourseDeletedEvent, &CourseDeleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *UpdateCourse, entity *Course, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, CourseAggregateType); err == nil {
            store.StoreEvent(CourseUpdatedEvent, &CourseUpdated{
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
    CreatedHandler func (*CourseCreated, *Course) (err error)  `json:"createdHandler" eh:"optional"`
    DeletedHandler func (*CourseDeleted, *Course) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdatedHandler func (*CourseUpdated, *Course) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *EventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case CourseCreatedEvent:
        err = o.CreatedHandler(event.Data().(*CourseCreated), entity.(*Course))
    case CourseDeletedEvent:
        err = o.DeletedHandler(event.Data().(*CourseDeleted), entity.(*Course))
    case CourseUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*CourseUpdated), entity.(*Course))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *EventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(CourseCreatedEvent, func() eventhorizon.EventData {
		return &CourseCreated{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *CourseCreated, entity *Course) (err error) {
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
    eventhorizon.RegisterEventData(CourseDeletedEvent, func() eventhorizon.EventData {
		return &CourseDeleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *CourseDeleted, entity *Course) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, CourseAggregateType); err == nil {
            *entity = *NewCourse()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(CourseUpdatedEvent, func() eventhorizon.EventData {
		return &CourseUpdated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *CourseUpdated, entity *Course) (err error) {
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



func NewAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
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
    CreateHandler func (*CreateGrade, *Grade, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*DeleteGrade, *Grade, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*UpdateGrade, *Grade, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *CommandHandler) AddCreatePreparer(preparer func (*CreateGrade, *Grade) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *CreateGrade, entity *Grade, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddDeletePreparer(preparer func (*DeleteGrade, *Grade) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *DeleteGrade, entity *Grade, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddUpdatePreparer(preparer func (*UpdateGrade, *Grade) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *UpdateGrade, entity *Grade, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateGradeCommand:
        err = o.CreateHandler(cmd.(*CreateGrade), entity.(*Grade), store)
    case DeleteGradeCommand:
        err = o.DeleteHandler(cmd.(*DeleteGrade), entity.(*Grade), store)
    case UpdateGradeCommand:
        err = o.UpdateHandler(cmd.(*UpdateGrade), entity.(*Grade), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *CommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *CreateGrade, entity *Grade, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, GradeAggregateType); err == nil {
            store.StoreEvent(GradeCreatedEvent, &GradeCreated{
                Student: command.Student,
                Course: command.Course,
                Grade: command.Grade,
                Comment: command.Comment,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *DeleteGrade, entity *Grade, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, GradeAggregateType); err == nil {
            store.StoreEvent(GradeDeletedEvent, &GradeDeleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *UpdateGrade, entity *Grade, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, GradeAggregateType); err == nil {
            store.StoreEvent(GradeUpdatedEvent, &GradeUpdated{
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
    CreatedHandler func (*GradeCreated, *Grade) (err error)  `json:"createdHandler" eh:"optional"`
    DeletedHandler func (*GradeDeleted, *Grade) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdatedHandler func (*GradeUpdated, *Grade) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *EventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case GradeCreatedEvent:
        err = o.CreatedHandler(event.Data().(*GradeCreated), entity.(*Grade))
    case GradeDeletedEvent:
        err = o.DeletedHandler(event.Data().(*GradeDeleted), entity.(*Grade))
    case GradeUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*GradeUpdated), entity.(*Grade))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *EventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(GradeCreatedEvent, func() eventhorizon.EventData {
		return &GradeCreated{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *GradeCreated, entity *Grade) (err error) {
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
    eventhorizon.RegisterEventData(GradeDeletedEvent, func() eventhorizon.EventData {
		return &GradeDeleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *GradeDeleted, entity *Grade) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, GradeAggregateType); err == nil {
            *entity = *NewGrade()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(GradeUpdatedEvent, func() eventhorizon.EventData {
		return &GradeUpdated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *GradeUpdated, entity *Grade) (err error) {
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



func NewAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
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
    CreateHandler func (*CreateGroup, *Group, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*DeleteGroup, *Group, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*UpdateGroup, *Group, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *CommandHandler) AddCreatePreparer(preparer func (*CreateGroup, *Group) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *CreateGroup, entity *Group, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddDeletePreparer(preparer func (*DeleteGroup, *Group) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *DeleteGroup, entity *Group, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddUpdatePreparer(preparer func (*UpdateGroup, *Group) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *UpdateGroup, entity *Group, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateGroupCommand:
        err = o.CreateHandler(cmd.(*CreateGroup), entity.(*Group), store)
    case DeleteGroupCommand:
        err = o.DeleteHandler(cmd.(*DeleteGroup), entity.(*Group), store)
    case UpdateGroupCommand:
        err = o.UpdateHandler(cmd.(*UpdateGroup), entity.(*Group), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *CommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *CreateGroup, entity *Group, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, GroupAggregateType); err == nil {
            store.StoreEvent(GroupCreatedEvent, &GroupCreated{
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
    o.DeleteHandler = func(command *DeleteGroup, entity *Group, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, GroupAggregateType); err == nil {
            store.StoreEvent(GroupDeletedEvent, &GroupDeleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *UpdateGroup, entity *Group, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, GroupAggregateType); err == nil {
            store.StoreEvent(GroupUpdatedEvent, &GroupUpdated{
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
    CreatedHandler func (*GroupCreated, *Group) (err error)  `json:"createdHandler" eh:"optional"`
    DeletedHandler func (*GroupDeleted, *Group) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdatedHandler func (*GroupUpdated, *Group) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *EventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case GroupCreatedEvent:
        err = o.CreatedHandler(event.Data().(*GroupCreated), entity.(*Group))
    case GroupDeletedEvent:
        err = o.DeletedHandler(event.Data().(*GroupDeleted), entity.(*Group))
    case GroupUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*GroupUpdated), entity.(*Group))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *EventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(GroupCreatedEvent, func() eventhorizon.EventData {
		return &GroupCreated{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *GroupCreated, entity *Group) (err error) {
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
    eventhorizon.RegisterEventData(GroupDeletedEvent, func() eventhorizon.EventData {
		return &GroupDeleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *GroupDeleted, entity *Group) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, GroupAggregateType); err == nil {
            *entity = *NewGroup()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(GroupUpdatedEvent, func() eventhorizon.EventData {
		return &GroupUpdated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *GroupUpdated, entity *Group) (err error) {
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



func NewAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
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
    CreateHandler func (*CreateSchoolApplication, *SchoolApplication, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*DeleteSchoolApplication, *SchoolApplication, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*UpdateSchoolApplication, *SchoolApplication, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *CommandHandler) AddCreatePreparer(preparer func (*CreateSchoolApplication, *SchoolApplication) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *CreateSchoolApplication, entity *SchoolApplication, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddDeletePreparer(preparer func (*DeleteSchoolApplication, *SchoolApplication) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *DeleteSchoolApplication, entity *SchoolApplication, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddUpdatePreparer(preparer func (*UpdateSchoolApplication, *SchoolApplication) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *UpdateSchoolApplication, entity *SchoolApplication, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateSchoolApplicationCommand:
        err = o.CreateHandler(cmd.(*CreateSchoolApplication), entity.(*SchoolApplication), store)
    case DeleteSchoolApplicationCommand:
        err = o.DeleteHandler(cmd.(*DeleteSchoolApplication), entity.(*SchoolApplication), store)
    case UpdateSchoolApplicationCommand:
        err = o.UpdateHandler(cmd.(*UpdateSchoolApplication), entity.(*SchoolApplication), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *CommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *CreateSchoolApplication, entity *SchoolApplication, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, SchoolApplicationAggregateType); err == nil {
            store.StoreEvent(SchoolApplicationCreatedEvent, &SchoolApplicationCreated{
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
    o.DeleteHandler = func(command *DeleteSchoolApplication, entity *SchoolApplication, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, SchoolApplicationAggregateType); err == nil {
            store.StoreEvent(SchoolApplicationDeletedEvent, &SchoolApplicationDeleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *UpdateSchoolApplication, entity *SchoolApplication, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, SchoolApplicationAggregateType); err == nil {
            store.StoreEvent(SchoolApplicationUpdatedEvent, &SchoolApplicationUpdated{
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
    CreatedHandler func (*SchoolApplicationCreated, *SchoolApplication) (err error)  `json:"createdHandler" eh:"optional"`
    DeletedHandler func (*SchoolApplicationDeleted, *SchoolApplication) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdatedHandler func (*SchoolApplicationUpdated, *SchoolApplication) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *EventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case SchoolApplicationCreatedEvent:
        err = o.CreatedHandler(event.Data().(*SchoolApplicationCreated), entity.(*SchoolApplication))
    case SchoolApplicationDeletedEvent:
        err = o.DeletedHandler(event.Data().(*SchoolApplicationDeleted), entity.(*SchoolApplication))
    case SchoolApplicationUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*SchoolApplicationUpdated), entity.(*SchoolApplication))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *EventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(SchoolApplicationCreatedEvent, func() eventhorizon.EventData {
		return &SchoolApplicationCreated{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *SchoolApplicationCreated, entity *SchoolApplication) (err error) {
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
    eventhorizon.RegisterEventData(SchoolApplicationDeletedEvent, func() eventhorizon.EventData {
		return &SchoolApplicationDeleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *SchoolApplicationDeleted, entity *SchoolApplication) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, SchoolApplicationAggregateType); err == nil {
            *entity = *NewSchoolApplication()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(SchoolApplicationUpdatedEvent, func() eventhorizon.EventData {
		return &SchoolApplicationUpdated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *SchoolApplicationUpdated, entity *SchoolApplication) (err error) {
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



func NewAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
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
    CreateHandler func (*CreateSchoolYear, *SchoolYear, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*DeleteSchoolYear, *SchoolYear, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*UpdateSchoolYear, *SchoolYear, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *CommandHandler) AddCreatePreparer(preparer func (*CreateSchoolYear, *SchoolYear) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *CreateSchoolYear, entity *SchoolYear, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddDeletePreparer(preparer func (*DeleteSchoolYear, *SchoolYear) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *DeleteSchoolYear, entity *SchoolYear, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) AddUpdatePreparer(preparer func (*UpdateSchoolYear, *SchoolYear) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *UpdateSchoolYear, entity *SchoolYear, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateSchoolYearCommand:
        err = o.CreateHandler(cmd.(*CreateSchoolYear), entity.(*SchoolYear), store)
    case DeleteSchoolYearCommand:
        err = o.DeleteHandler(cmd.(*DeleteSchoolYear), entity.(*SchoolYear), store)
    case UpdateSchoolYearCommand:
        err = o.UpdateHandler(cmd.(*UpdateSchoolYear), entity.(*SchoolYear), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *CommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *CreateSchoolYear, entity *SchoolYear, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, SchoolYearAggregateType); err == nil {
            store.StoreEvent(SchoolYearCreatedEvent, &SchoolYearCreated{
                Name: command.Name,
                Start: command.Start,
                End: command.End,
                Dates: command.Dates,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *DeleteSchoolYear, entity *SchoolYear, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, SchoolYearAggregateType); err == nil {
            store.StoreEvent(SchoolYearDeletedEvent, &SchoolYearDeleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *UpdateSchoolYear, entity *SchoolYear, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, SchoolYearAggregateType); err == nil {
            store.StoreEvent(SchoolYearUpdatedEvent, &SchoolYearUpdated{
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
    CreatedHandler func (*SchoolYearCreated, *SchoolYear) (err error)  `json:"createdHandler" eh:"optional"`
    DeletedHandler func (*SchoolYearDeleted, *SchoolYear) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdatedHandler func (*SchoolYearUpdated, *SchoolYear) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *EventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case SchoolYearCreatedEvent:
        err = o.CreatedHandler(event.Data().(*SchoolYearCreated), entity.(*SchoolYear))
    case SchoolYearDeletedEvent:
        err = o.DeletedHandler(event.Data().(*SchoolYearDeleted), entity.(*SchoolYear))
    case SchoolYearUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*SchoolYearUpdated), entity.(*SchoolYear))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *EventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(SchoolYearCreatedEvent, func() eventhorizon.EventData {
		return &SchoolYearCreated{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *SchoolYearCreated, entity *SchoolYear) (err error) {
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
    eventhorizon.RegisterEventData(SchoolYearDeletedEvent, func() eventhorizon.EventData {
		return &SchoolYearDeleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *SchoolYearDeleted, entity *SchoolYear) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, SchoolYearAggregateType); err == nil {
            *entity = *NewSchoolYear()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(SchoolYearUpdatedEvent, func() eventhorizon.EventData {
		return &SchoolYearUpdated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *SchoolYearUpdated, entity *SchoolYear) (err error) {
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



func NewAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
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

func NewStudentEventhorizonInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *StudentEventhorizonInitializer) {
    attendanceAggregateInitializer := NewAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    courseAggregateInitializer := NewAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    gradeAggregateInitializer := NewAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    groupAggregateInitializer := NewAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    schoolApplicationAggregateInitializer := NewAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    schoolYearAggregateInitializer := NewAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
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









