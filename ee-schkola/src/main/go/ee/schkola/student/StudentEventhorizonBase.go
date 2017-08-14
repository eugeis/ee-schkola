package student

import (
    "errors"
    "fmt"
    "github.com/eugeis/gee/eh"
    "github.com/looplab/eventhorizon"
    "time"
)
type AttendanceCommandHandler struct {
    ConfirmHandler func (*ConfirmAttendance, *Attendance, eh.AggregateStoreEvent) error
    CancelHandler func (*CancelAttendance, *Attendance, eh.AggregateStoreEvent) error
    RegisterHandler func (*RegisterAttendance, *Attendance, eh.AggregateStoreEvent) error
    CreateHandler func (*CreateAttendance, *Attendance, eh.AggregateStoreEvent) error
    DeleteHandler func (*DeleteAttendance, *Attendance, eh.AggregateStoreEvent) error
    UpdateHandler func (*UpdateAttendance, *Attendance, eh.AggregateStoreEvent) error
}

func (o *AttendanceCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (ret error) {
    switch cmd.CommandType() {
    case ConfirmAttendanceCommand:
        ret = o.ConfirmHandler(cmd.(*ConfirmAttendance), entity.(*Attendance), store)
    case CancelAttendanceCommand:
        ret = o.CancelHandler(cmd.(*CancelAttendance), entity.(*Attendance), store)
    case RegisterAttendanceCommand:
        ret = o.RegisterHandler(cmd.(*RegisterAttendance), entity.(*Attendance), store)
    case CreateAttendanceCommand:
        ret = o.CreateHandler(cmd.(*CreateAttendance), entity.(*Attendance), store)
    case DeleteAttendanceCommand:
        ret = o.DeleteHandler(cmd.(*DeleteAttendance), entity.(*Attendance), store)
    case UpdateAttendanceCommand:
        ret = o.UpdateHandler(cmd.(*UpdateAttendance), entity.(*Attendance), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *AttendanceCommandHandler) SetupCommandHandler() (ret error) {
    if o.ConfirmHandler == nil {
        o.ConfirmHandler = func(command *ConfirmAttendance, entity *Attendance, store eh.AggregateStoreEvent) (ret error) {ret = eh.CommandHandlerNotImplemented(ConfirmAttendanceCommand)
            return
        }
    }
    
    if o.CancelHandler == nil {
        o.CancelHandler = func(command *CancelAttendance, entity *Attendance, store eh.AggregateStoreEvent) (ret error) {ret = eh.CommandHandlerNotImplemented(CancelAttendanceCommand)
            return
        }
    }
    
    if o.RegisterHandler == nil {
        o.RegisterHandler = func(command *RegisterAttendance, entity *Attendance, store eh.AggregateStoreEvent) (ret error) {ret = eh.CommandHandlerNotImplemented(RegisterAttendanceCommand)
            return
        }
    }
    
    if o.CreateHandler == nil {
        o.CreateHandler = func(command *CreateAttendance, entity *Attendance, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityAlreadyExists(entity.Id, AttendanceAggregateType)
            } else {
                store.StoreEvent(AttendanceCreatedEvent, &AttendanceCreated{
                    Id: command.Id,
                    Student: command.Student,
                    Date: command.Date,
                    Course: command.Course,
                    Hours: command.Hours,
                    State: command.State,
                    StateTrace: command.StateTrace,
                    Token: command.Token,
                    TokenTrace: command.TokenTrace,}, time.Now())
            }
            return
        }
    }
    
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteAttendance, entity *Attendance, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, AttendanceAggregateType)
            } else if entity.Id != command.Id {
                ret = eh.IdsDismatch(entity.Id, command.Id, AttendanceAggregateType)
            } else {
                store.StoreEvent(AttendanceDeletedEvent, &AttendanceDeleted{
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    
    if o.UpdateHandler == nil {
        o.UpdateHandler = func(command *UpdateAttendance, entity *Attendance, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, AttendanceAggregateType)
            } else if entity.Id != command.Id {
                ret = eh.IdsDismatch(entity.Id, command.Id, AttendanceAggregateType)
            } else {
                store.StoreEvent(AttendanceUpdatedEvent, &AttendanceUpdated{
                    Id: command.Id,
                    Student: command.Student,
                    Date: command.Date,
                    Course: command.Course,
                    Hours: command.Hours,
                    State: command.State,
                    StateTrace: command.StateTrace,
                    Token: command.Token,
                    TokenTrace: command.TokenTrace,}, time.Now())
            }
            return
        }
    }
    
    return
}


type AttendanceEventHandler struct {
    CreatedHandler func (*AttendanceCreated, *Attendance) error
    DeletedHandler func (*AttendanceDeleted, *Attendance) error
    UpdatedHandler func (*AttendanceUpdated, *Attendance) error
}

func (o *AttendanceEventHandler) Apply(event eventhorizon.Event, entity interface{}) (ret error) {
    switch event.EventType() {
    case AttendanceCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*AttendanceCreated), entity.(*Attendance))
    case AttendanceDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*AttendanceDeleted), entity.(*Attendance))
    case AttendanceUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*AttendanceUpdated), entity.(*Attendance))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *AttendanceEventHandler) SetupEventHandler() (ret error) {
    if o.CreatedHandler == nil {
        o.CreatedHandler = func(event *AttendanceCreated, entity *Attendance) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityAlreadyExists(entity.Id, AttendanceAggregateType)
            } else {
                entity.Id = event.Id
                entity.Student = event.Student
                entity.Date = event.Date
                entity.Course = event.Course
                entity.Hours = event.Hours
                entity.State = event.State
                entity.StateTrace = event.StateTrace
                entity.Token = event.Token
                entity.TokenTrace = event.TokenTrace
            }
            return
        }
    }
    
    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *AttendanceDeleted, entity *Attendance) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, AttendanceAggregateType)
            } else if entity.Id != event.Id {
                ret = eh.IdsDismatch(entity.Id, event.Id, AttendanceAggregateType)
            } else {
                entity.Student = nil
                entity.Date = nil
                entity.Course = nil
                entity.Hours = nil
                entity.State = nil
                entity.StateTrace = nil
                entity.Token = ""
                entity.TokenTrace = nil
                entity.SchkolaBase = nil
            }
            return
        }
    }
    
    if o.UpdatedHandler == nil {
        o.UpdatedHandler = func(event *AttendanceUpdated, entity *Attendance) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, AttendanceAggregateType)
            } else if entity.Id != event.Id {
                ret = eh.IdsDismatch(entity.Id, event.Id, AttendanceAggregateType)
            } else {
                entity.Student = event.Student
                entity.Date = event.Date
                entity.Course = event.Course
                entity.Hours = event.Hours
                entity.State = event.State
                entity.StateTrace = event.StateTrace
                entity.Token = event.Token
                entity.TokenTrace = event.TokenTrace
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
}



func NewAttendanceAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus) (ret *AttendanceAggregateInitializer) {
    
    commandHandler := &AttendanceCommandHandler{}
    eventHandler := &AttendanceEventHandler{}
	ret = &AttendanceAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(AttendanceAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(AttendanceAggregateType, id, commandHandler, eventHandler, NewAttendance())
        },
        AttendanceCommandTypes().Literals(), AttendanceEventTypes().Literals(),
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus),
        AttendanceCommandHandler: commandHandler,
        AttendanceEventHandler: eventHandler,
    }

    return
}


type CourseCommandHandler struct {
    CreateHandler func (*CreateCourse, *Course, eh.AggregateStoreEvent) error
    DeleteHandler func (*DeleteCourse, *Course, eh.AggregateStoreEvent) error
    UpdateHandler func (*UpdateCourse, *Course, eh.AggregateStoreEvent) error
}

func (o *CourseCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (ret error) {
    switch cmd.CommandType() {
    case CreateCourseCommand:
        ret = o.CreateHandler(cmd.(*CreateCourse), entity.(*Course), store)
    case DeleteCourseCommand:
        ret = o.DeleteHandler(cmd.(*DeleteCourse), entity.(*Course), store)
    case UpdateCourseCommand:
        ret = o.UpdateHandler(cmd.(*UpdateCourse), entity.(*Course), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *CourseCommandHandler) SetupCommandHandler() (ret error) {
    if o.CreateHandler == nil {
        o.CreateHandler = func(command *CreateCourse, entity *Course, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityAlreadyExists(entity.Id, CourseAggregateType)
            } else {
                store.StoreEvent(CourseCreatedEvent, &CourseCreated{
                    Id: command.Id,
                    Name: command.Name,
                    Begin: command.Begin,
                    End: command.End,
                    Teacher: command.Teacher,
                    SchoolYear: command.SchoolYear,
                    Fee: command.Fee,
                    Description: command.Description,}, time.Now())
            }
            return
        }
    }
    
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteCourse, entity *Course, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, CourseAggregateType)
            } else if entity.Id != command.Id {
                ret = eh.IdsDismatch(entity.Id, command.Id, CourseAggregateType)
            } else {
                store.StoreEvent(CourseDeletedEvent, &CourseDeleted{
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    
    if o.UpdateHandler == nil {
        o.UpdateHandler = func(command *UpdateCourse, entity *Course, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, CourseAggregateType)
            } else if entity.Id != command.Id {
                ret = eh.IdsDismatch(entity.Id, command.Id, CourseAggregateType)
            } else {
                store.StoreEvent(CourseUpdatedEvent, &CourseUpdated{
                    Id: command.Id,
                    Name: command.Name,
                    Begin: command.Begin,
                    End: command.End,
                    Teacher: command.Teacher,
                    SchoolYear: command.SchoolYear,
                    Fee: command.Fee,
                    Description: command.Description,}, time.Now())
            }
            return
        }
    }
    
    return
}


type CourseEventHandler struct {
    CreatedHandler func (*CourseCreated, *Course) error
    DeletedHandler func (*CourseDeleted, *Course) error
    UpdatedHandler func (*CourseUpdated, *Course) error
}

func (o *CourseEventHandler) Apply(event eventhorizon.Event, entity interface{}) (ret error) {
    switch event.EventType() {
    case CourseCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*CourseCreated), entity.(*Course))
    case CourseDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*CourseDeleted), entity.(*Course))
    case CourseUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*CourseUpdated), entity.(*Course))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *CourseEventHandler) SetupEventHandler() (ret error) {
    if o.CreatedHandler == nil {
        o.CreatedHandler = func(event *CourseCreated, entity *Course) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityAlreadyExists(entity.Id, CourseAggregateType)
            } else {
                entity.Id = event.Id
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
    
    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *CourseDeleted, entity *Course) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, CourseAggregateType)
            } else if entity.Id != event.Id {
                ret = eh.IdsDismatch(entity.Id, event.Id, CourseAggregateType)
            } else {
                entity.Name = ""
                entity.Begin = nil
                entity.End = nil
                entity.Teacher = nil
                entity.SchoolYear = nil
                entity.Fee = nil
                entity.Description = nil
                entity.SchkolaBase = nil
            }
            return
        }
    }
    
    if o.UpdatedHandler == nil {
        o.UpdatedHandler = func(event *CourseUpdated, entity *Course) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, CourseAggregateType)
            } else if entity.Id != event.Id {
                ret = eh.IdsDismatch(entity.Id, event.Id, CourseAggregateType)
            } else {
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
}



func NewCourseAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus) (ret *CourseAggregateInitializer) {
    
    commandHandler := &CourseCommandHandler{}
    eventHandler := &CourseEventHandler{}
	ret = &CourseAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(CourseAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(CourseAggregateType, id, commandHandler, eventHandler, NewCourse())
        },
        CourseCommandTypes().Literals(), CourseEventTypes().Literals(),
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus),
        CourseCommandHandler: commandHandler,
        CourseEventHandler: eventHandler,
    }

    return
}


type GradeCommandHandler struct {
    CreateHandler func (*CreateGrade, *Grade, eh.AggregateStoreEvent) error
    DeleteHandler func (*DeleteGrade, *Grade, eh.AggregateStoreEvent) error
    UpdateHandler func (*UpdateGrade, *Grade, eh.AggregateStoreEvent) error
}

func (o *GradeCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (ret error) {
    switch cmd.CommandType() {
    case CreateGradeCommand:
        ret = o.CreateHandler(cmd.(*CreateGrade), entity.(*Grade), store)
    case DeleteGradeCommand:
        ret = o.DeleteHandler(cmd.(*DeleteGrade), entity.(*Grade), store)
    case UpdateGradeCommand:
        ret = o.UpdateHandler(cmd.(*UpdateGrade), entity.(*Grade), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *GradeCommandHandler) SetupCommandHandler() (ret error) {
    if o.CreateHandler == nil {
        o.CreateHandler = func(command *CreateGrade, entity *Grade, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityAlreadyExists(entity.Id, GradeAggregateType)
            } else {
                store.StoreEvent(GradeCreatedEvent, &GradeCreated{
                    Id: command.Id,
                    Student: command.Student,
                    Course: command.Course,
                    Grade: command.Grade,
                    GradeTrace: command.GradeTrace,
                    Comment: command.Comment,}, time.Now())
            }
            return
        }
    }
    
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteGrade, entity *Grade, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, GradeAggregateType)
            } else if entity.Id != command.Id {
                ret = eh.IdsDismatch(entity.Id, command.Id, GradeAggregateType)
            } else {
                store.StoreEvent(GradeDeletedEvent, &GradeDeleted{
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    
    if o.UpdateHandler == nil {
        o.UpdateHandler = func(command *UpdateGrade, entity *Grade, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, GradeAggregateType)
            } else if entity.Id != command.Id {
                ret = eh.IdsDismatch(entity.Id, command.Id, GradeAggregateType)
            } else {
                store.StoreEvent(GradeUpdatedEvent, &GradeUpdated{
                    Id: command.Id,
                    Student: command.Student,
                    Course: command.Course,
                    Grade: command.Grade,
                    GradeTrace: command.GradeTrace,
                    Comment: command.Comment,}, time.Now())
            }
            return
        }
    }
    
    return
}


type GradeEventHandler struct {
    CreatedHandler func (*GradeCreated, *Grade) error
    DeletedHandler func (*GradeDeleted, *Grade) error
    UpdatedHandler func (*GradeUpdated, *Grade) error
}

func (o *GradeEventHandler) Apply(event eventhorizon.Event, entity interface{}) (ret error) {
    switch event.EventType() {
    case GradeCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*GradeCreated), entity.(*Grade))
    case GradeDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*GradeDeleted), entity.(*Grade))
    case GradeUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*GradeUpdated), entity.(*Grade))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *GradeEventHandler) SetupEventHandler() (ret error) {
    if o.CreatedHandler == nil {
        o.CreatedHandler = func(event *GradeCreated, entity *Grade) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityAlreadyExists(entity.Id, GradeAggregateType)
            } else {
                entity.Id = event.Id
                entity.Student = event.Student
                entity.Course = event.Course
                entity.Grade = event.Grade
                entity.GradeTrace = event.GradeTrace
                entity.Comment = event.Comment
            }
            return
        }
    }
    
    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *GradeDeleted, entity *Grade) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, GradeAggregateType)
            } else if entity.Id != event.Id {
                ret = eh.IdsDismatch(entity.Id, event.Id, GradeAggregateType)
            } else {
                entity.Student = nil
                entity.Course = nil
                entity.Grade = nil
                entity.GradeTrace = nil
                entity.Comment = ""
                entity.SchkolaBase = nil
            }
            return
        }
    }
    
    if o.UpdatedHandler == nil {
        o.UpdatedHandler = func(event *GradeUpdated, entity *Grade) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, GradeAggregateType)
            } else if entity.Id != event.Id {
                ret = eh.IdsDismatch(entity.Id, event.Id, GradeAggregateType)
            } else {
                entity.Student = event.Student
                entity.Course = event.Course
                entity.Grade = event.Grade
                entity.GradeTrace = event.GradeTrace
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
}



func NewGradeAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus) (ret *GradeAggregateInitializer) {
    
    commandHandler := &GradeCommandHandler{}
    eventHandler := &GradeEventHandler{}
	ret = &GradeAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(GradeAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(GradeAggregateType, id, commandHandler, eventHandler, NewGrade())
        },
        GradeCommandTypes().Literals(), GradeEventTypes().Literals(),
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus),
        GradeCommandHandler: commandHandler,
        GradeEventHandler: eventHandler,
    }

    return
}


type GroupCommandHandler struct {
    CreateHandler func (*CreateGroup, *Group, eh.AggregateStoreEvent) error
    DeleteHandler func (*DeleteGroup, *Group, eh.AggregateStoreEvent) error
    UpdateHandler func (*UpdateGroup, *Group, eh.AggregateStoreEvent) error
}

func (o *GroupCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (ret error) {
    switch cmd.CommandType() {
    case CreateGroupCommand:
        ret = o.CreateHandler(cmd.(*CreateGroup), entity.(*Group), store)
    case DeleteGroupCommand:
        ret = o.DeleteHandler(cmd.(*DeleteGroup), entity.(*Group), store)
    case UpdateGroupCommand:
        ret = o.UpdateHandler(cmd.(*UpdateGroup), entity.(*Group), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *GroupCommandHandler) SetupCommandHandler() (ret error) {
    if o.CreateHandler == nil {
        o.CreateHandler = func(command *CreateGroup, entity *Group, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityAlreadyExists(entity.Id, GroupAggregateType)
            } else {
                store.StoreEvent(GroupCreatedEvent, &GroupCreated{
                    Id: command.Id,
                    Name: command.Name,
                    Category: command.Category,
                    SchoolYear: command.SchoolYear,
                    Representative: command.Representative,
                    Students: command.Students,
                    Courses: command.Courses,}, time.Now())
            }
            return
        }
    }
    
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteGroup, entity *Group, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, GroupAggregateType)
            } else if entity.Id != command.Id {
                ret = eh.IdsDismatch(entity.Id, command.Id, GroupAggregateType)
            } else {
                store.StoreEvent(GroupDeletedEvent, &GroupDeleted{
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    
    if o.UpdateHandler == nil {
        o.UpdateHandler = func(command *UpdateGroup, entity *Group, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, GroupAggregateType)
            } else if entity.Id != command.Id {
                ret = eh.IdsDismatch(entity.Id, command.Id, GroupAggregateType)
            } else {
                store.StoreEvent(GroupUpdatedEvent, &GroupUpdated{
                    Id: command.Id,
                    Name: command.Name,
                    Category: command.Category,
                    SchoolYear: command.SchoolYear,
                    Representative: command.Representative,
                    Students: command.Students,
                    Courses: command.Courses,}, time.Now())
            }
            return
        }
    }
    
    return
}


type GroupEventHandler struct {
    CreatedHandler func (*GroupCreated, *Group) error
    DeletedHandler func (*GroupDeleted, *Group) error
    UpdatedHandler func (*GroupUpdated, *Group) error
}

func (o *GroupEventHandler) Apply(event eventhorizon.Event, entity interface{}) (ret error) {
    switch event.EventType() {
    case GroupCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*GroupCreated), entity.(*Group))
    case GroupDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*GroupDeleted), entity.(*Group))
    case GroupUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*GroupUpdated), entity.(*Group))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *GroupEventHandler) SetupEventHandler() (ret error) {
    if o.CreatedHandler == nil {
        o.CreatedHandler = func(event *GroupCreated, entity *Group) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityAlreadyExists(entity.Id, GroupAggregateType)
            } else {
                entity.Id = event.Id
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
    
    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *GroupDeleted, entity *Group) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, GroupAggregateType)
            } else if entity.Id != event.Id {
                ret = eh.IdsDismatch(entity.Id, event.Id, GroupAggregateType)
            } else {
                entity.Name = ""
                entity.Category = nil
                entity.SchoolYear = nil
                entity.Representative = nil
                entity.Students = nil
                entity.Courses = nil
                entity.SchkolaBase = nil
            }
            return
        }
    }
    
    if o.UpdatedHandler == nil {
        o.UpdatedHandler = func(event *GroupUpdated, entity *Group) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, GroupAggregateType)
            } else if entity.Id != event.Id {
                ret = eh.IdsDismatch(entity.Id, event.Id, GroupAggregateType)
            } else {
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
}



func NewGroupAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus) (ret *GroupAggregateInitializer) {
    
    commandHandler := &GroupCommandHandler{}
    eventHandler := &GroupEventHandler{}
	ret = &GroupAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(GroupAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(GroupAggregateType, id, commandHandler, eventHandler, NewGroup())
        },
        GroupCommandTypes().Literals(), GroupEventTypes().Literals(),
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus),
        GroupCommandHandler: commandHandler,
        GroupEventHandler: eventHandler,
    }

    return
}


type SchoolApplicationCommandHandler struct {
    CreateHandler func (*CreateSchoolApplication, *SchoolApplication, eh.AggregateStoreEvent) error
    DeleteHandler func (*DeleteSchoolApplication, *SchoolApplication, eh.AggregateStoreEvent) error
    UpdateHandler func (*UpdateSchoolApplication, *SchoolApplication, eh.AggregateStoreEvent) error
}

func (o *SchoolApplicationCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (ret error) {
    switch cmd.CommandType() {
    case CreateSchoolApplicationCommand:
        ret = o.CreateHandler(cmd.(*CreateSchoolApplication), entity.(*SchoolApplication), store)
    case DeleteSchoolApplicationCommand:
        ret = o.DeleteHandler(cmd.(*DeleteSchoolApplication), entity.(*SchoolApplication), store)
    case UpdateSchoolApplicationCommand:
        ret = o.UpdateHandler(cmd.(*UpdateSchoolApplication), entity.(*SchoolApplication), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *SchoolApplicationCommandHandler) SetupCommandHandler() (ret error) {
    if o.CreateHandler == nil {
        o.CreateHandler = func(command *CreateSchoolApplication, entity *SchoolApplication, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityAlreadyExists(entity.Id, SchoolApplicationAggregateType)
            } else {
                store.StoreEvent(SchoolApplicationCreatedEvent, &SchoolApplicationCreated{
                    Id: command.Id,
                    Profile: command.Profile,
                    RecommendationOf: command.RecommendationOf,
                    ChurchContactPerson: command.ChurchContactPerson,
                    ChurchContact: command.ChurchContact,
                    SchoolYear: command.SchoolYear,
                    Group: command.Group,}, time.Now())
            }
            return
        }
    }
    
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteSchoolApplication, entity *SchoolApplication, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, SchoolApplicationAggregateType)
            } else if entity.Id != command.Id {
                ret = eh.IdsDismatch(entity.Id, command.Id, SchoolApplicationAggregateType)
            } else {
                store.StoreEvent(SchoolApplicationDeletedEvent, &SchoolApplicationDeleted{
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    
    if o.UpdateHandler == nil {
        o.UpdateHandler = func(command *UpdateSchoolApplication, entity *SchoolApplication, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, SchoolApplicationAggregateType)
            } else if entity.Id != command.Id {
                ret = eh.IdsDismatch(entity.Id, command.Id, SchoolApplicationAggregateType)
            } else {
                store.StoreEvent(SchoolApplicationUpdatedEvent, &SchoolApplicationUpdated{
                    Id: command.Id,
                    Profile: command.Profile,
                    RecommendationOf: command.RecommendationOf,
                    ChurchContactPerson: command.ChurchContactPerson,
                    ChurchContact: command.ChurchContact,
                    SchoolYear: command.SchoolYear,
                    Group: command.Group,}, time.Now())
            }
            return
        }
    }
    
    return
}


type SchoolApplicationEventHandler struct {
    CreatedHandler func (*SchoolApplicationCreated, *SchoolApplication) error
    DeletedHandler func (*SchoolApplicationDeleted, *SchoolApplication) error
    UpdatedHandler func (*SchoolApplicationUpdated, *SchoolApplication) error
}

func (o *SchoolApplicationEventHandler) Apply(event eventhorizon.Event, entity interface{}) (ret error) {
    switch event.EventType() {
    case SchoolApplicationCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*SchoolApplicationCreated), entity.(*SchoolApplication))
    case SchoolApplicationDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*SchoolApplicationDeleted), entity.(*SchoolApplication))
    case SchoolApplicationUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*SchoolApplicationUpdated), entity.(*SchoolApplication))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *SchoolApplicationEventHandler) SetupEventHandler() (ret error) {
    if o.CreatedHandler == nil {
        o.CreatedHandler = func(event *SchoolApplicationCreated, entity *SchoolApplication) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityAlreadyExists(entity.Id, SchoolApplicationAggregateType)
            } else {
                entity.Id = event.Id
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
    
    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *SchoolApplicationDeleted, entity *SchoolApplication) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, SchoolApplicationAggregateType)
            } else if entity.Id != event.Id {
                ret = eh.IdsDismatch(entity.Id, event.Id, SchoolApplicationAggregateType)
            } else {
                entity.Profile = nil
                entity.RecommendationOf = nil
                entity.ChurchContactPerson = nil
                entity.ChurchContact = nil
                entity.SchoolYear = nil
                entity.Group = ""
                entity.SchkolaBase = nil
            }
            return
        }
    }
    
    if o.UpdatedHandler == nil {
        o.UpdatedHandler = func(event *SchoolApplicationUpdated, entity *SchoolApplication) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, SchoolApplicationAggregateType)
            } else if entity.Id != event.Id {
                ret = eh.IdsDismatch(entity.Id, event.Id, SchoolApplicationAggregateType)
            } else {
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
}



func NewSchoolApplicationAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus) (ret *SchoolApplicationAggregateInitializer) {
    
    commandHandler := &SchoolApplicationCommandHandler{}
    eventHandler := &SchoolApplicationEventHandler{}
	ret = &SchoolApplicationAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(SchoolApplicationAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(SchoolApplicationAggregateType, id, commandHandler, eventHandler, NewSchoolApplication())
        },
        SchoolApplicationCommandTypes().Literals(), SchoolApplicationEventTypes().Literals(),
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus),
        SchoolApplicationCommandHandler: commandHandler,
        SchoolApplicationEventHandler: eventHandler,
    }

    return
}


type SchoolYearCommandHandler struct {
    CreateHandler func (*CreateSchoolYear, *SchoolYear, eh.AggregateStoreEvent) error
    DeleteHandler func (*DeleteSchoolYear, *SchoolYear, eh.AggregateStoreEvent) error
    UpdateHandler func (*UpdateSchoolYear, *SchoolYear, eh.AggregateStoreEvent) error
}

func (o *SchoolYearCommandHandler) Execute(cmd eventhorizon.Command, entity interface{}, store eh.AggregateStoreEvent) (ret error) {
    switch cmd.CommandType() {
    case CreateSchoolYearCommand:
        ret = o.CreateHandler(cmd.(*CreateSchoolYear), entity.(*SchoolYear), store)
    case DeleteSchoolYearCommand:
        ret = o.DeleteHandler(cmd.(*DeleteSchoolYear), entity.(*SchoolYear), store)
    case UpdateSchoolYearCommand:
        ret = o.UpdateHandler(cmd.(*UpdateSchoolYear), entity.(*SchoolYear), store)
    default:
		ret = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *SchoolYearCommandHandler) SetupCommandHandler() (ret error) {
    if o.CreateHandler == nil {
        o.CreateHandler = func(command *CreateSchoolYear, entity *SchoolYear, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityAlreadyExists(entity.Id, SchoolYearAggregateType)
            } else {
                store.StoreEvent(SchoolYearCreatedEvent, &SchoolYearCreated{
                    Id: command.Id,
                    Name: command.Name,
                    Start: command.Start,
                    End: command.End,
                    Dates: command.Dates,}, time.Now())
            }
            return
        }
    }
    
    if o.DeleteHandler == nil {
        o.DeleteHandler = func(command *DeleteSchoolYear, entity *SchoolYear, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, SchoolYearAggregateType)
            } else if entity.Id != command.Id {
                ret = eh.IdsDismatch(entity.Id, command.Id, SchoolYearAggregateType)
            } else {
                store.StoreEvent(SchoolYearDeletedEvent, &SchoolYearDeleted{
                    Id: command.Id,}, time.Now())
            }
            return
        }
    }
    
    if o.UpdateHandler == nil {
        o.UpdateHandler = func(command *UpdateSchoolYear, entity *SchoolYear, store eh.AggregateStoreEvent) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, SchoolYearAggregateType)
            } else if entity.Id != command.Id {
                ret = eh.IdsDismatch(entity.Id, command.Id, SchoolYearAggregateType)
            } else {
                store.StoreEvent(SchoolYearUpdatedEvent, &SchoolYearUpdated{
                    Id: command.Id,
                    Name: command.Name,
                    Start: command.Start,
                    End: command.End,
                    Dates: command.Dates,}, time.Now())
            }
            return
        }
    }
    
    return
}


type SchoolYearEventHandler struct {
    CreatedHandler func (*SchoolYearCreated, *SchoolYear) error
    DeletedHandler func (*SchoolYearDeleted, *SchoolYear) error
    UpdatedHandler func (*SchoolYearUpdated, *SchoolYear) error
}

func (o *SchoolYearEventHandler) Apply(event eventhorizon.Event, entity interface{}) (ret error) {
    switch event.EventType() {
    case SchoolYearCreatedEvent:
        ret = o.CreatedHandler(event.Data().(*SchoolYearCreated), entity.(*SchoolYear))
    case SchoolYearDeletedEvent:
        ret = o.DeletedHandler(event.Data().(*SchoolYearDeleted), entity.(*SchoolYear))
    case SchoolYearUpdatedEvent:
        ret = o.UpdatedHandler(event.Data().(*SchoolYearUpdated), entity.(*SchoolYear))
    default:
		ret = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *SchoolYearEventHandler) SetupEventHandler() (ret error) {
    if o.CreatedHandler == nil {
        o.CreatedHandler = func(event *SchoolYearCreated, entity *SchoolYear) (ret error) {
            if len(entity.Id) > 0 {
                ret = eh.EntityAlreadyExists(entity.Id, SchoolYearAggregateType)
            } else {
                entity.Id = event.Id
                entity.Name = event.Name
                entity.Start = event.Start
                entity.End = event.End
                entity.Dates = event.Dates
            }
            return
        }
    }
    
    if o.DeletedHandler == nil {
        o.DeletedHandler = func(event *SchoolYearDeleted, entity *SchoolYear) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, SchoolYearAggregateType)
            } else if entity.Id != event.Id {
                ret = eh.IdsDismatch(entity.Id, event.Id, SchoolYearAggregateType)
            } else {
                entity.Name = ""
                entity.Start = nil
                entity.End = nil
                entity.Dates = nil
                entity.SchkolaBase = nil
            }
            return
        }
    }
    
    if o.UpdatedHandler == nil {
        o.UpdatedHandler = func(event *SchoolYearUpdated, entity *SchoolYear) (ret error) {
            if len(entity.Id) == 0 {
                ret = eh.EntityNotExists(entity.Id, SchoolYearAggregateType)
            } else if entity.Id != event.Id {
                ret = eh.IdsDismatch(entity.Id, event.Id, SchoolYearAggregateType)
            } else {
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
}



func NewSchoolYearAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, eventPublisher eventhorizon.EventPublisher, 
                commandBus eventhorizon.CommandBus) (ret *SchoolYearAggregateInitializer) {
    
    commandHandler := &SchoolYearCommandHandler{}
    eventHandler := &SchoolYearEventHandler{}
	ret = &SchoolYearAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(SchoolYearAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(SchoolYearAggregateType, id, commandHandler, eventHandler, NewSchoolYear())
        },
        SchoolYearCommandTypes().Literals(), SchoolYearEventTypes().Literals(),
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, eventPublisher, commandBus),
        SchoolYearCommandHandler: commandHandler,
        SchoolYearEventHandler: eventHandler,
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
                commandBus eventhorizon.CommandBus) (ret *StudentEventhorizonInitializer) {
    ret = &StudentEventhorizonInitializer{
        eventStore: eventStore,
        eventBus: eventBus,
        eventPublisher: eventPublisher,
        commandBus: commandBus,
        AttendanceAggregateInitializer: NewAttendanceAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
        CourseAggregateInitializer: NewCourseAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
        GradeAggregateInitializer: NewGradeAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
        GroupAggregateInitializer: NewGroupAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
        SchoolApplicationAggregateInitializer: NewSchoolApplicationAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
        SchoolYearAggregateInitializer: NewSchoolYearAggregateInitializer(eventStore, eventBus, eventPublisher, commandBus),
    }
    return
}

func (o *StudentEventhorizonInitializer) Setup() (ret error) {
    
    if ret = o.AttendanceAggregateInitializer.Setup(); ret != nil {
        return
    }
    
    if ret = o.CourseAggregateInitializer.Setup(); ret != nil {
        return
    }
    
    if ret = o.GradeAggregateInitializer.Setup(); ret != nil {
        return
    }
    
    if ret = o.GroupAggregateInitializer.Setup(); ret != nil {
        return
    }
    
    if ret = o.SchoolApplicationAggregateInitializer.Setup(); ret != nil {
        return
    }
    
    if ret = o.SchoolYearAggregateInitializer.Setup(); ret != nil {
        return
    }

    return
}









