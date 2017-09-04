package student

import (
    "errors"
    "fmt"
    "github.com/eugeis/gee/eh"
    "github.com/looplab/eventhorizon"
    "time"
)
type AttendanceCommandHandler struct {
    ConfirmHandler func (*ConfirmAttendance, *Attendance, eh.AggregateStoreEvent) (err error) 
    CancelHandler func (*CancelAttendance, *Attendance, eh.AggregateStoreEvent) (err error) 
    RegisterHandler func (*RegisterAttendance, *Attendance, eh.AggregateStoreEvent) (err error) 
    CreateHandler func (*CreateAttendance, *Attendance, eh.AggregateStoreEvent) (err error) 
    DeleteHandler func (*DeleteAttendance, *Attendance, eh.AggregateStoreEvent) (err error) 
    UpdateHandler func (*UpdateAttendance, *Attendance, eh.AggregateStoreEvent) (err error) 
    AddItem *T
}

func (o *AttendanceCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case ConfirmAttendanceCommand:
        err = o.ConfirmHandler(cmd.(*ConfirmAttendance), entity.(*Attendance), store)
    case CancelAttendanceCommand:
        err = o.CancelHandler(cmd.(*CancelAttendance), entity.(*Attendance), store)
    case RegisterAttendanceCommand:
        err = o.RegisterHandler(cmd.(*RegisterAttendance), entity.(*Attendance), store)
    case CreateAttendanceCommand:
        err = o.CreateHandler(cmd.(*CreateAttendance), entity.(*Attendance), store)
    case DeleteAttendanceCommand:
        err = o.DeleteHandler(cmd.(*DeleteAttendance), entity.(*Attendance), store)
    case UpdateAttendanceCommand:
        err = o.UpdateHandler(cmd.(*UpdateAttendance), entity.(*Attendance), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *AttendanceCommandHandler) SetupCommandHandler() (err error) {
    if o.ConfirmHandler == nil {
        o.ConfirmHandler = func(command *ConfirmAttendance, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
            if err = eh.ValidateIdsMatch(entity.Id, command.Id, AttendanceAggregateType); err == nil {
                store.StoreEvent(AttendanceConfirmdEvent, &AttendanceConfirmd{
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    if o.CancelHandler == nil {
        o.CancelHandler = func(command *CancelAttendance, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
            if err = eh.ValidateIdsMatch(entity.Id, command.Id, AttendanceAggregateType); err == nil {
                store.StoreEvent(AttendanceCanceldEvent, &AttendanceCanceld{
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    if o.RegisterHandler == nil {
        o.RegisterHandler = func(command *RegisterAttendance, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
            if err = eh.ValidateNewId(entity.Id, command.Id, AttendanceAggregateType); err == nil {
                store.StoreEvent(AttendanceRegisterdEvent, &AttendanceRegisterd{
                    Student: command.Student,
                    Course: command.Course,
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    if o.CreateHandler == nil {
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
    }
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteAttendance, entity *Attendance, store eh.AggregateStoreEvent) (err error) {
            if err = eh.ValidateIdsMatch(entity.Id, command.Id, AttendanceAggregateType); err == nil {
                store.StoreEvent(AttendanceDeletedEvent, &AttendanceDeleted{
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    if o.UpdateHandler == nil {
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
    }
    return
}


type AttendanceEventHandler struct {
    CreatedHandler func (*AttendanceCreated, *Attendance) (err error) 
    RegisterdHandler func (*AttendanceRegisterd, *Attendance) (err error) 
    DeletedHandler func (*AttendanceDeleted, *Attendance) (err error) 
    ConfirmdHandler func (*AttendanceConfirmd, *Attendance) (err error) 
    CanceldHandler func (*AttendanceCanceld, *Attendance) (err error) 
    UpdatedHandler func (*AttendanceUpdated, *Attendance) (err error) 
    AddItem *T
}

func (o *AttendanceEventHandler) Apply(event eventhorizon.Event, entity interface{}) (err error) {
    switch event.EventType() {
    case AttendanceCreatedEvent:
        err = o.CreatedHandler(event.Data().(*AttendanceCreated), entity.(*Attendance))
    case AttendanceRegisterdEvent:
        err = o.RegisterdHandler(event.Data().(*AttendanceRegisterd), entity.(*Attendance))
    case AttendanceDeletedEvent:
        err = o.DeletedHandler(event.Data().(*AttendanceDeleted), entity.(*Attendance))
    case AttendanceConfirmdEvent:
        err = o.ConfirmdHandler(event.Data().(*AttendanceConfirmd), entity.(*Attendance))
    case AttendanceCanceldEvent:
        err = o.CanceldHandler(event.Data().(*AttendanceCanceld), entity.(*Attendance))
    case AttendanceUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*AttendanceUpdated), entity.(*Attendance))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *AttendanceEventHandler) SetupEventHandler() (err error) {
	eventhorizon.RegisterEventData(AttendanceCreatedEvent, func() eventhorizon.EventData {
		return &AttendanceCreated{}
	})

    if o.CreatedHandler == nil {
        o.CreatedHandler = func(event *AttendanceCreated, entity *Attendance) (err error) {
            if err = eh.ValidateNewId(entity.Id, event.Id, AttendanceAggregateType); err == nil {
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
    }
	eventhorizon.RegisterEventData(AttendanceRegisterdEvent, func() eventhorizon.EventData {
		return &AttendanceRegisterd{}
	})

    if o.RegisterdHandler == nil {
        o.RegisterdHandler = func(event *AttendanceRegisterd, entity *Attendance) (err error) {
            if err = eh.ValidateNewId(entity.Id, event.Id, AttendanceAggregateType); err == nil {
                entity.Student = event.Student
                entity.Course = event.Course
                entity.Id = event.Id
            }
            return
        }
    }
	eventhorizon.RegisterEventData(AttendanceDeletedEvent, func() eventhorizon.EventData {
		return &AttendanceDeleted{}
	})

    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *AttendanceDeleted, entity *Attendance) (err error) {
            if err = eh.ValidateIdsMatch(entity.Id, event.Id, AttendanceAggregateType); err == nil {
                *entity = *NewAttendance()
            }
            return
        }
    }
	eventhorizon.RegisterEventData(AttendanceConfirmdEvent, func() eventhorizon.EventData {
		return &AttendanceConfirmd{}
	})

    if o.ConfirmdHandler == nil {
        o.ConfirmdHandler = func(event *AttendanceConfirmd, entity *Attendance) (err error) {    err = eh.EventHandlerNotImplemented(AttendanceConfirmdEvent)
            return
        }
    }
	eventhorizon.RegisterEventData(AttendanceCanceldEvent, func() eventhorizon.EventData {
		return &AttendanceCanceld{}
	})

    if o.CanceldHandler == nil {
        o.CanceldHandler = func(event *AttendanceCanceld, entity *Attendance) (err error) {    err = eh.EventHandlerNotImplemented(AttendanceCanceldEvent)
            return
        }
    }
	eventhorizon.RegisterEventData(AttendanceUpdatedEvent, func() eventhorizon.EventData {
		return &AttendanceUpdated{}
	})

    if o.UpdatedHandler == nil {
        o.UpdatedHandler = func(event *AttendanceUpdated, entity *Attendance) (err error) {
            if err = eh.ValidateIdsMatch(entity.Id, event.Id, AttendanceAggregateType); err == nil {
                entity.Student = event.Student
                entity.Date = event.Date
                entity.Course = event.Course
                entity.Hours = event.Hours
                entity.State = event.State
                entity.Token = event.Token
            }
            return
        }
    }
    return
}


const AttendanceAggregateType eventhorizon.AggregateType = "Attendance"

type AttendanceAggregateInitializer struct {
    *eh.AggregateInitializer
    *AttendanceCommandHandler
    *AttendanceEventHandler
    ProjectorHandler *AttendanceEventHandler
    AddItem *T
}


func (o *AttendanceAggregateInitializer) RegisterForConfirmd(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AttendanceEventTypes().AttendanceConfirmd())
}

func (o *AttendanceAggregateInitializer) RegisterForCanceld(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AttendanceEventTypes().AttendanceCanceld())
}


func NewAttendanceAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus, 
                readRepos func (string, func () (ret interface{}) ) (ret eventhorizon.ReadWriteRepo) ) (ret *AttendanceAggregateInitializer) {
    
    commandHandler := &AttendanceCommandHandler{}
    eventHandler := &AttendanceEventHandler{}
    modelFactory := func() interface{} { return NewAttendance() }
    ret = &AttendanceAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(AttendanceAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(AttendanceAggregateType, id, commandHandler, eventHandler, modelFactory())
        }, modelFactory,
        AttendanceCommandTypes().Literals(), AttendanceEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), AttendanceCommandHandler: commandHandler, AttendanceEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type CourseCommandHandler struct {
    CreateHandler func (*CreateCourse, *Course, eh.AggregateStoreEvent) (err error) 
    DeleteHandler func (*DeleteCourse, *Course, eh.AggregateStoreEvent) (err error) 
    UpdateHandler func (*UpdateCourse, *Course, eh.AggregateStoreEvent) (err error) 
    AddItem *T
}

func (o *CourseCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (err error) {
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

func (o *CourseCommandHandler) SetupCommandHandler() (err error) {
    if o.CreateHandler == nil {
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
    }
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteCourse, entity *Course, store eh.AggregateStoreEvent) (err error) {
            if err = eh.ValidateIdsMatch(entity.Id, command.Id, CourseAggregateType); err == nil {
                store.StoreEvent(CourseDeletedEvent, &CourseDeleted{
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    if o.UpdateHandler == nil {
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
    }
    return
}


type CourseEventHandler struct {
    CreatedHandler func (*CourseCreated, *Course) (err error) 
    DeletedHandler func (*CourseDeleted, *Course) (err error) 
    UpdatedHandler func (*CourseUpdated, *Course) (err error) 
    AddItem *T
}

func (o *CourseEventHandler) Apply(event eventhorizon.Event, entity interface{}) (err error) {
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

func (o *CourseEventHandler) SetupEventHandler() (err error) {
	eventhorizon.RegisterEventData(CourseCreatedEvent, func() eventhorizon.EventData {
		return &CourseCreated{}
	})

    if o.CreatedHandler == nil {
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
    }
	eventhorizon.RegisterEventData(CourseDeletedEvent, func() eventhorizon.EventData {
		return &CourseDeleted{}
	})

    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *CourseDeleted, entity *Course) (err error) {
            if err = eh.ValidateIdsMatch(entity.Id, event.Id, CourseAggregateType); err == nil {
                *entity = *NewCourse()
            }
            return
        }
    }
	eventhorizon.RegisterEventData(CourseUpdatedEvent, func() eventhorizon.EventData {
		return &CourseUpdated{}
	})

    if o.UpdatedHandler == nil {
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
    }
    return
}


const CourseAggregateType eventhorizon.AggregateType = "Course"

type CourseAggregateInitializer struct {
    *eh.AggregateInitializer
    *CourseCommandHandler
    *CourseEventHandler
    ProjectorHandler *CourseEventHandler
    AddItem *T
}



func NewCourseAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus, 
                readRepos func (string, func () (ret interface{}) ) (ret eventhorizon.ReadWriteRepo) ) (ret *CourseAggregateInitializer) {
    
    commandHandler := &CourseCommandHandler{}
    eventHandler := &CourseEventHandler{}
    modelFactory := func() interface{} { return NewCourse() }
    ret = &CourseAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(CourseAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(CourseAggregateType, id, commandHandler, eventHandler, modelFactory())
        }, modelFactory,
        CourseCommandTypes().Literals(), CourseEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), CourseCommandHandler: commandHandler, CourseEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type GradeCommandHandler struct {
    CreateHandler func (*CreateGrade, *Grade, eh.AggregateStoreEvent) (err error) 
    DeleteHandler func (*DeleteGrade, *Grade, eh.AggregateStoreEvent) (err error) 
    UpdateHandler func (*UpdateGrade, *Grade, eh.AggregateStoreEvent) (err error) 
    AddItem *T
}

func (o *GradeCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (err error) {
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

func (o *GradeCommandHandler) SetupCommandHandler() (err error) {
    if o.CreateHandler == nil {
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
    }
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteGrade, entity *Grade, store eh.AggregateStoreEvent) (err error) {
            if err = eh.ValidateIdsMatch(entity.Id, command.Id, GradeAggregateType); err == nil {
                store.StoreEvent(GradeDeletedEvent, &GradeDeleted{
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    if o.UpdateHandler == nil {
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
    }
    return
}


type GradeEventHandler struct {
    CreatedHandler func (*GradeCreated, *Grade) (err error) 
    DeletedHandler func (*GradeDeleted, *Grade) (err error) 
    UpdatedHandler func (*GradeUpdated, *Grade) (err error) 
    AddItem *T
}

func (o *GradeEventHandler) Apply(event eventhorizon.Event, entity interface{}) (err error) {
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

func (o *GradeEventHandler) SetupEventHandler() (err error) {
	eventhorizon.RegisterEventData(GradeCreatedEvent, func() eventhorizon.EventData {
		return &GradeCreated{}
	})

    if o.CreatedHandler == nil {
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
    }
	eventhorizon.RegisterEventData(GradeDeletedEvent, func() eventhorizon.EventData {
		return &GradeDeleted{}
	})

    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *GradeDeleted, entity *Grade) (err error) {
            if err = eh.ValidateIdsMatch(entity.Id, event.Id, GradeAggregateType); err == nil {
                *entity = *NewGrade()
            }
            return
        }
    }
	eventhorizon.RegisterEventData(GradeUpdatedEvent, func() eventhorizon.EventData {
		return &GradeUpdated{}
	})

    if o.UpdatedHandler == nil {
        o.UpdatedHandler = func(event *GradeUpdated, entity *Grade) (err error) {
            if err = eh.ValidateIdsMatch(entity.Id, event.Id, GradeAggregateType); err == nil {
                entity.Student = event.Student
                entity.Course = event.Course
                entity.Grade = event.Grade
                entity.Comment = event.Comment
            }
            return
        }
    }
    return
}


const GradeAggregateType eventhorizon.AggregateType = "Grade"

type GradeAggregateInitializer struct {
    *eh.AggregateInitializer
    *GradeCommandHandler
    *GradeEventHandler
    ProjectorHandler *GradeEventHandler
    AddItem *T
}



func NewGradeAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus, 
                readRepos func (string, func () (ret interface{}) ) (ret eventhorizon.ReadWriteRepo) ) (ret *GradeAggregateInitializer) {
    
    commandHandler := &GradeCommandHandler{}
    eventHandler := &GradeEventHandler{}
    modelFactory := func() interface{} { return NewGrade() }
    ret = &GradeAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(GradeAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(GradeAggregateType, id, commandHandler, eventHandler, modelFactory())
        }, modelFactory,
        GradeCommandTypes().Literals(), GradeEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), GradeCommandHandler: commandHandler, GradeEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type GroupCommandHandler struct {
    CreateHandler func (*CreateGroup, *Group, eh.AggregateStoreEvent) (err error) 
    DeleteHandler func (*DeleteGroup, *Group, eh.AggregateStoreEvent) (err error) 
    UpdateHandler func (*UpdateGroup, *Group, eh.AggregateStoreEvent) (err error) 
    AddItem *T
}

func (o *GroupCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (err error) {
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

func (o *GroupCommandHandler) SetupCommandHandler() (err error) {
    if o.CreateHandler == nil {
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
    }
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteGroup, entity *Group, store eh.AggregateStoreEvent) (err error) {
            if err = eh.ValidateIdsMatch(entity.Id, command.Id, GroupAggregateType); err == nil {
                store.StoreEvent(GroupDeletedEvent, &GroupDeleted{
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    if o.UpdateHandler == nil {
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
    }
    return
}


type GroupEventHandler struct {
    CreatedHandler func (*GroupCreated, *Group) (err error) 
    DeletedHandler func (*GroupDeleted, *Group) (err error) 
    UpdatedHandler func (*GroupUpdated, *Group) (err error) 
    AddItem *T
}

func (o *GroupEventHandler) Apply(event eventhorizon.Event, entity interface{}) (err error) {
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

func (o *GroupEventHandler) SetupEventHandler() (err error) {
	eventhorizon.RegisterEventData(GroupCreatedEvent, func() eventhorizon.EventData {
		return &GroupCreated{}
	})

    if o.CreatedHandler == nil {
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
    }
	eventhorizon.RegisterEventData(GroupDeletedEvent, func() eventhorizon.EventData {
		return &GroupDeleted{}
	})

    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *GroupDeleted, entity *Group) (err error) {
            if err = eh.ValidateIdsMatch(entity.Id, event.Id, GroupAggregateType); err == nil {
                *entity = *NewGroup()
            }
            return
        }
    }
	eventhorizon.RegisterEventData(GroupUpdatedEvent, func() eventhorizon.EventData {
		return &GroupUpdated{}
	})

    if o.UpdatedHandler == nil {
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
    }
    return
}


const GroupAggregateType eventhorizon.AggregateType = "Group"

type GroupAggregateInitializer struct {
    *eh.AggregateInitializer
    *GroupCommandHandler
    *GroupEventHandler
    ProjectorHandler *GroupEventHandler
    AddItem *T
}



func NewGroupAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus, 
                readRepos func (string, func () (ret interface{}) ) (ret eventhorizon.ReadWriteRepo) ) (ret *GroupAggregateInitializer) {
    
    commandHandler := &GroupCommandHandler{}
    eventHandler := &GroupEventHandler{}
    modelFactory := func() interface{} { return NewGroup() }
    ret = &GroupAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(GroupAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(GroupAggregateType, id, commandHandler, eventHandler, modelFactory())
        }, modelFactory,
        GroupCommandTypes().Literals(), GroupEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), GroupCommandHandler: commandHandler, GroupEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type SchoolApplicationCommandHandler struct {
    CreateHandler func (*CreateSchoolApplication, *SchoolApplication, eh.AggregateStoreEvent) (err error) 
    DeleteHandler func (*DeleteSchoolApplication, *SchoolApplication, eh.AggregateStoreEvent) (err error) 
    UpdateHandler func (*UpdateSchoolApplication, *SchoolApplication, eh.AggregateStoreEvent) (err error) 
    AddItem *T
}

func (o *SchoolApplicationCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (err error) {
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

func (o *SchoolApplicationCommandHandler) SetupCommandHandler() (err error) {
    if o.CreateHandler == nil {
        o.CreateHandler = func(command *CreateSchoolApplication, entity *SchoolApplication, store eh.AggregateStoreEvent) (err error) {
            if err = eh.ValidateNewId(entity.Id, command.Id, SchoolApplicationAggregateType); err == nil {
                store.StoreEvent(SchoolApplicationCreatedEvent, &SchoolApplicationCreated{
                    Profile: command.Profile,
                    RecommendationOf: command.RecommendationOf,
                    ChurchContactPerson: command.ChurchContactPerson,
                    ChurchContact: command.ChurchContact,
                    SchoolYear: command.SchoolYear,
                    Group: command.Group,
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteSchoolApplication, entity *SchoolApplication, store eh.AggregateStoreEvent) (err error) {
            if err = eh.ValidateIdsMatch(entity.Id, command.Id, SchoolApplicationAggregateType); err == nil {
                store.StoreEvent(SchoolApplicationDeletedEvent, &SchoolApplicationDeleted{
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    if o.UpdateHandler == nil {
        o.UpdateHandler = func(command *UpdateSchoolApplication, entity *SchoolApplication, store eh.AggregateStoreEvent) (err error) {
            if err = eh.ValidateIdsMatch(entity.Id, command.Id, SchoolApplicationAggregateType); err == nil {
                store.StoreEvent(SchoolApplicationUpdatedEvent, &SchoolApplicationUpdated{
                    Profile: command.Profile,
                    RecommendationOf: command.RecommendationOf,
                    ChurchContactPerson: command.ChurchContactPerson,
                    ChurchContact: command.ChurchContact,
                    SchoolYear: command.SchoolYear,
                    Group: command.Group,
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    return
}


type SchoolApplicationEventHandler struct {
    CreatedHandler func (*SchoolApplicationCreated, *SchoolApplication) (err error) 
    DeletedHandler func (*SchoolApplicationDeleted, *SchoolApplication) (err error) 
    UpdatedHandler func (*SchoolApplicationUpdated, *SchoolApplication) (err error) 
    AddItem *T
}

func (o *SchoolApplicationEventHandler) Apply(event eventhorizon.Event, entity interface{}) (err error) {
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

func (o *SchoolApplicationEventHandler) SetupEventHandler() (err error) {
	eventhorizon.RegisterEventData(SchoolApplicationCreatedEvent, func() eventhorizon.EventData {
		return &SchoolApplicationCreated{}
	})

    if o.CreatedHandler == nil {
        o.CreatedHandler = func(event *SchoolApplicationCreated, entity *SchoolApplication) (err error) {
            if err = eh.ValidateNewId(entity.Id, event.Id, SchoolApplicationAggregateType); err == nil {
                entity.Profile = event.Profile
                entity.RecommendationOf = event.RecommendationOf
                entity.ChurchContactPerson = event.ChurchContactPerson
                entity.ChurchContact = event.ChurchContact
                entity.SchoolYear = event.SchoolYear
                entity.Group = event.Group
                entity.Id = event.Id
            }
            return
        }
    }
	eventhorizon.RegisterEventData(SchoolApplicationDeletedEvent, func() eventhorizon.EventData {
		return &SchoolApplicationDeleted{}
	})

    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *SchoolApplicationDeleted, entity *SchoolApplication) (err error) {
            if err = eh.ValidateIdsMatch(entity.Id, event.Id, SchoolApplicationAggregateType); err == nil {
                *entity = *NewSchoolApplication()
            }
            return
        }
    }
	eventhorizon.RegisterEventData(SchoolApplicationUpdatedEvent, func() eventhorizon.EventData {
		return &SchoolApplicationUpdated{}
	})

    if o.UpdatedHandler == nil {
        o.UpdatedHandler = func(event *SchoolApplicationUpdated, entity *SchoolApplication) (err error) {
            if err = eh.ValidateIdsMatch(entity.Id, event.Id, SchoolApplicationAggregateType); err == nil {
                entity.Profile = event.Profile
                entity.RecommendationOf = event.RecommendationOf
                entity.ChurchContactPerson = event.ChurchContactPerson
                entity.ChurchContact = event.ChurchContact
                entity.SchoolYear = event.SchoolYear
                entity.Group = event.Group
            }
            return
        }
    }
    return
}


const SchoolApplicationAggregateType eventhorizon.AggregateType = "SchoolApplication"

type SchoolApplicationAggregateInitializer struct {
    *eh.AggregateInitializer
    *SchoolApplicationCommandHandler
    *SchoolApplicationEventHandler
    ProjectorHandler *SchoolApplicationEventHandler
    AddItem *T
}



func NewSchoolApplicationAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus, 
                readRepos func (string, func () (ret interface{}) ) (ret eventhorizon.ReadWriteRepo) ) (ret *SchoolApplicationAggregateInitializer) {
    
    commandHandler := &SchoolApplicationCommandHandler{}
    eventHandler := &SchoolApplicationEventHandler{}
    modelFactory := func() interface{} { return NewSchoolApplication() }
    ret = &SchoolApplicationAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(SchoolApplicationAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(SchoolApplicationAggregateType, id, commandHandler, eventHandler, modelFactory())
        }, modelFactory,
        SchoolApplicationCommandTypes().Literals(), SchoolApplicationEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), SchoolApplicationCommandHandler: commandHandler, SchoolApplicationEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type SchoolYearCommandHandler struct {
    CreateHandler func (*CreateSchoolYear, *SchoolYear, eh.AggregateStoreEvent) (err error) 
    DeleteHandler func (*DeleteSchoolYear, *SchoolYear, eh.AggregateStoreEvent) (err error) 
    UpdateHandler func (*UpdateSchoolYear, *SchoolYear, eh.AggregateStoreEvent) (err error) 
    AddItem *T
}

func (o *SchoolYearCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (err error) {
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

func (o *SchoolYearCommandHandler) SetupCommandHandler() (err error) {
    if o.CreateHandler == nil {
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
    }
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteSchoolYear, entity *SchoolYear, store eh.AggregateStoreEvent) (err error) {
            if err = eh.ValidateIdsMatch(entity.Id, command.Id, SchoolYearAggregateType); err == nil {
                store.StoreEvent(SchoolYearDeletedEvent, &SchoolYearDeleted{
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    if o.UpdateHandler == nil {
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
    }
    return
}


type SchoolYearEventHandler struct {
    CreatedHandler func (*SchoolYearCreated, *SchoolYear) (err error) 
    DeletedHandler func (*SchoolYearDeleted, *SchoolYear) (err error) 
    UpdatedHandler func (*SchoolYearUpdated, *SchoolYear) (err error) 
    AddItem *T
}

func (o *SchoolYearEventHandler) Apply(event eventhorizon.Event, entity interface{}) (err error) {
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

func (o *SchoolYearEventHandler) SetupEventHandler() (err error) {
	eventhorizon.RegisterEventData(SchoolYearCreatedEvent, func() eventhorizon.EventData {
		return &SchoolYearCreated{}
	})

    if o.CreatedHandler == nil {
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
    }
	eventhorizon.RegisterEventData(SchoolYearDeletedEvent, func() eventhorizon.EventData {
		return &SchoolYearDeleted{}
	})

    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *SchoolYearDeleted, entity *SchoolYear) (err error) {
            if err = eh.ValidateIdsMatch(entity.Id, event.Id, SchoolYearAggregateType); err == nil {
                *entity = *NewSchoolYear()
            }
            return
        }
    }
	eventhorizon.RegisterEventData(SchoolYearUpdatedEvent, func() eventhorizon.EventData {
		return &SchoolYearUpdated{}
	})

    if o.UpdatedHandler == nil {
        o.UpdatedHandler = func(event *SchoolYearUpdated, entity *SchoolYear) (err error) {
            if err = eh.ValidateIdsMatch(entity.Id, event.Id, SchoolYearAggregateType); err == nil {
                entity.Name = event.Name
                entity.Start = event.Start
                entity.End = event.End
                entity.Dates = event.Dates
            }
            return
        }
    }
    return
}


const SchoolYearAggregateType eventhorizon.AggregateType = "SchoolYear"

type SchoolYearAggregateInitializer struct {
    *eh.AggregateInitializer
    *SchoolYearCommandHandler
    *SchoolYearEventHandler
    ProjectorHandler *SchoolYearEventHandler
    AddItem *T
}



func NewSchoolYearAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus, 
                readRepos func (string, func () (ret interface{}) ) (ret eventhorizon.ReadWriteRepo) ) (ret *SchoolYearAggregateInitializer) {
    
    commandHandler := &SchoolYearCommandHandler{}
    eventHandler := &SchoolYearEventHandler{}
    modelFactory := func() interface{} { return NewSchoolYear() }
    ret = &SchoolYearAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(SchoolYearAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(SchoolYearAggregateType, id, commandHandler, eventHandler, modelFactory())
        }, modelFactory,
        SchoolYearCommandTypes().Literals(), SchoolYearEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus, readRepos), SchoolYearCommandHandler: commandHandler, SchoolYearEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type StudentEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore
    eventBus eventhorizon.EventBus
    eventPublisher eventhorizon.EventPublisher
    commandBus eventhorizon.CommandBus
    AttendanceAggregateInitializer *AttendanceAggregateInitializer
    CourseAggregateInitializer *CourseAggregateInitializer
    GradeAggregateInitializer *GradeAggregateInitializer
    GroupAggregateInitializer *GroupAggregateInitializer
    SchoolApplicationAggregateInitializer *SchoolApplicationAggregateInitializer
    SchoolYearAggregateInitializer *SchoolYearAggregateInitializer
}

func NewStudentEventhorizonInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus, 
                readRepos func (string, func () (ret interface{}) ) (ret eventhorizon.ReadWriteRepo) ) (ret *StudentEventhorizonInitializer) {
    attendanceAggregateInitializer := NewAttendanceAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    courseAggregateInitializer := NewCourseAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    gradeAggregateInitializer := NewGradeAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    groupAggregateInitializer := NewGroupAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    schoolApplicationAggregateInitializer := NewSchoolApplicationAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
    schoolYearAggregateInitializer := NewSchoolYearAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
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









