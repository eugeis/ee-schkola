package student

import (
    "context"
    "github.com/eugeis/gee/eh"
    "github.com/eugeis/gee/net"
    "github.com/gorilla/mux"
    "github.com/looplab/eventhorizon"
    "net/http"
)
type AttendanceHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *AttendanceQueryRepository
}

func NewAttendanceHttpQueryHandler(queryRepository *AttendanceQueryRepository) (ret *AttendanceHttpQueryHandler) {
    httpQueryHandler := &eh.HttpQueryHandler{}
    ret = &AttendanceHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *AttendanceHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllAttendance", w, r)
}

func (o *AttendanceHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindByAttendanceId", w, r)
}

func (o *AttendanceHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllAttendance", w, r)
}

func (o *AttendanceHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountByAttendanceId", w, r)
}

func (o *AttendanceHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllAttendance", w, r)
}

func (o *AttendanceHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistByAttendanceId", w, r)
}


type AttendanceHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewAttendanceHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *AttendanceHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &AttendanceHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *AttendanceHttpCommandHandler) Register(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&RegisterAttendance{Id: id}, w, r)
}

func (o *AttendanceHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateAttendance{Id: id}, w, r)
}

func (o *AttendanceHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateAttendance{Id: id}, w, r)
}

func (o *AttendanceHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteAttendance{Id: id}, w, r)
}

func (o *AttendanceHttpCommandHandler) Confirm(w http.ResponseWriter, r *http.Request)  {
}

func (o *AttendanceHttpCommandHandler) Cancel(w http.ResponseWriter, r *http.Request)  {
}


type AttendanceRouter struct {
    PathPrefix string
    QueryHandler *AttendanceHttpQueryHandler
    CommandHandler *AttendanceHttpCommandHandler
    Router *mux.Router
}

func NewAttendanceRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus, 
                readRepos func (string) eventhorizon.ReadWriteRepo) (ret *AttendanceRouter) {
    pathPrefix = pathPrefix + "/" + "attendances"
    queryRepository := NewAttendanceQueryRepository()
    queryHandler := NewAttendanceHttpQueryHandler(queryRepository)
    commandHandler := NewAttendanceHttpCommandHandler(context, commandBus)
    ret = &AttendanceRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *AttendanceRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountAttendanceById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("CountAttendanceAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistAttendanceById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("ExistAttendanceAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindAttendanceById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("FindAttendanceAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("RegisterAttendance").HandlerFunc(o.CommandHandler.Register)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateAttendance").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateAttendance").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteAttendance").HandlerFunc(o.CommandHandler.Delete)
    return
}


type CourseHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *CourseQueryRepository
}

func NewCourseHttpQueryHandler(queryRepository *CourseQueryRepository) (ret *CourseHttpQueryHandler) {
    httpQueryHandler := &eh.HttpQueryHandler{}
    ret = &CourseHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *CourseHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllCourse", w, r)
}

func (o *CourseHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindByCourseId", w, r)
}

func (o *CourseHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllCourse", w, r)
}

func (o *CourseHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountByCourseId", w, r)
}

func (o *CourseHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllCourse", w, r)
}

func (o *CourseHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistByCourseId", w, r)
}


type CourseHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewCourseHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *CourseHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &CourseHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *CourseHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateCourse{Id: id}, w, r)
}

func (o *CourseHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateCourse{Id: id}, w, r)
}

func (o *CourseHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteCourse{Id: id}, w, r)
}


type CourseRouter struct {
    PathPrefix string
    QueryHandler *CourseHttpQueryHandler
    CommandHandler *CourseHttpCommandHandler
    Router *mux.Router
}

func NewCourseRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus, 
                readRepos func (string) eventhorizon.ReadWriteRepo) (ret *CourseRouter) {
    pathPrefix = pathPrefix + "/" + "courses"
    queryRepository := NewCourseQueryRepository()
    queryHandler := NewCourseHttpQueryHandler(queryRepository)
    commandHandler := NewCourseHttpCommandHandler(context, commandBus)
    ret = &CourseRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *CourseRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountCourseById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("CountCourseAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistCourseById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("ExistCourseAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindCourseById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("FindCourseAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateCourse").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateCourse").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteCourse").HandlerFunc(o.CommandHandler.Delete)
    return
}


type GradeHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *GradeQueryRepository
}

func NewGradeHttpQueryHandler(queryRepository *GradeQueryRepository) (ret *GradeHttpQueryHandler) {
    httpQueryHandler := &eh.HttpQueryHandler{}
    ret = &GradeHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *GradeHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllGrade", w, r)
}

func (o *GradeHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindByGradeId", w, r)
}

func (o *GradeHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllGrade", w, r)
}

func (o *GradeHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountByGradeId", w, r)
}

func (o *GradeHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllGrade", w, r)
}

func (o *GradeHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistByGradeId", w, r)
}


type GradeHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewGradeHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *GradeHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &GradeHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *GradeHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateGrade{Id: id}, w, r)
}

func (o *GradeHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateGrade{Id: id}, w, r)
}

func (o *GradeHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteGrade{Id: id}, w, r)
}


type GradeRouter struct {
    PathPrefix string
    QueryHandler *GradeHttpQueryHandler
    CommandHandler *GradeHttpCommandHandler
    Router *mux.Router
}

func NewGradeRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus, 
                readRepos func (string) eventhorizon.ReadWriteRepo) (ret *GradeRouter) {
    pathPrefix = pathPrefix + "/" + "grades"
    queryRepository := NewGradeQueryRepository()
    queryHandler := NewGradeHttpQueryHandler(queryRepository)
    commandHandler := NewGradeHttpCommandHandler(context, commandBus)
    ret = &GradeRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *GradeRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountGradeById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("CountGradeAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistGradeById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("ExistGradeAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindGradeById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("FindGradeAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateGrade").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateGrade").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteGrade").HandlerFunc(o.CommandHandler.Delete)
    return
}


type GroupHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *GroupQueryRepository
}

func NewGroupHttpQueryHandler(queryRepository *GroupQueryRepository) (ret *GroupHttpQueryHandler) {
    httpQueryHandler := &eh.HttpQueryHandler{}
    ret = &GroupHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *GroupHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllGroup", w, r)
}

func (o *GroupHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindByGroupId", w, r)
}

func (o *GroupHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllGroup", w, r)
}

func (o *GroupHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountByGroupId", w, r)
}

func (o *GroupHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllGroup", w, r)
}

func (o *GroupHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistByGroupId", w, r)
}


type GroupHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewGroupHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *GroupHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &GroupHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *GroupHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateGroup{Id: id}, w, r)
}

func (o *GroupHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateGroup{Id: id}, w, r)
}

func (o *GroupHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteGroup{Id: id}, w, r)
}


type GroupRouter struct {
    PathPrefix string
    QueryHandler *GroupHttpQueryHandler
    CommandHandler *GroupHttpCommandHandler
    Router *mux.Router
}

func NewGroupRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus, 
                readRepos func (string) eventhorizon.ReadWriteRepo) (ret *GroupRouter) {
    pathPrefix = pathPrefix + "/" + "groups"
    queryRepository := NewGroupQueryRepository()
    queryHandler := NewGroupHttpQueryHandler(queryRepository)
    commandHandler := NewGroupHttpCommandHandler(context, commandBus)
    ret = &GroupRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *GroupRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountGroupById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("CountGroupAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistGroupById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("ExistGroupAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindGroupById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("FindGroupAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateGroup").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateGroup").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteGroup").HandlerFunc(o.CommandHandler.Delete)
    return
}


type SchoolApplicationHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *SchoolApplicationQueryRepository
}

func NewSchoolApplicationHttpQueryHandler(queryRepository *SchoolApplicationQueryRepository) (ret *SchoolApplicationHttpQueryHandler) {
    httpQueryHandler := &eh.HttpQueryHandler{}
    ret = &SchoolApplicationHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *SchoolApplicationHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllSchoolApplication", w, r)
}

func (o *SchoolApplicationHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindBySchoolApplicationId", w, r)
}

func (o *SchoolApplicationHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllSchoolApplication", w, r)
}

func (o *SchoolApplicationHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountBySchoolApplicationId", w, r)
}

func (o *SchoolApplicationHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllSchoolApplication", w, r)
}

func (o *SchoolApplicationHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistBySchoolApplicationId", w, r)
}


type SchoolApplicationHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewSchoolApplicationHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *SchoolApplicationHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &SchoolApplicationHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *SchoolApplicationHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateSchoolApplication{Id: id}, w, r)
}

func (o *SchoolApplicationHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateSchoolApplication{Id: id}, w, r)
}

func (o *SchoolApplicationHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteSchoolApplication{Id: id}, w, r)
}


type SchoolApplicationRouter struct {
    PathPrefix string
    QueryHandler *SchoolApplicationHttpQueryHandler
    CommandHandler *SchoolApplicationHttpCommandHandler
    Router *mux.Router
}

func NewSchoolApplicationRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus, 
                readRepos func (string) eventhorizon.ReadWriteRepo) (ret *SchoolApplicationRouter) {
    pathPrefix = pathPrefix + "/" + "schoolApplications"
    queryRepository := NewSchoolApplicationQueryRepository()
    queryHandler := NewSchoolApplicationHttpQueryHandler(queryRepository)
    commandHandler := NewSchoolApplicationHttpCommandHandler(context, commandBus)
    ret = &SchoolApplicationRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *SchoolApplicationRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountSchoolApplicationById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("CountSchoolApplicationAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistSchoolApplicationById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("ExistSchoolApplicationAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindSchoolApplicationById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("FindSchoolApplicationAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateSchoolApplication").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateSchoolApplication").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteSchoolApplication").HandlerFunc(o.CommandHandler.Delete)
    return
}


type SchoolYearHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *SchoolYearQueryRepository
}

func NewSchoolYearHttpQueryHandler(queryRepository *SchoolYearQueryRepository) (ret *SchoolYearHttpQueryHandler) {
    httpQueryHandler := &eh.HttpQueryHandler{}
    ret = &SchoolYearHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *SchoolYearHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllSchoolYear", w, r)
}

func (o *SchoolYearHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindBySchoolYearId", w, r)
}

func (o *SchoolYearHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllSchoolYear", w, r)
}

func (o *SchoolYearHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountBySchoolYearId", w, r)
}

func (o *SchoolYearHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllSchoolYear", w, r)
}

func (o *SchoolYearHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistBySchoolYearId", w, r)
}


type SchoolYearHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewSchoolYearHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *SchoolYearHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &SchoolYearHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *SchoolYearHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateSchoolYear{Id: id}, w, r)
}

func (o *SchoolYearHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateSchoolYear{Id: id}, w, r)
}

func (o *SchoolYearHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteSchoolYear{Id: id}, w, r)
}


type SchoolYearRouter struct {
    PathPrefix string
    QueryHandler *SchoolYearHttpQueryHandler
    CommandHandler *SchoolYearHttpCommandHandler
    Router *mux.Router
}

func NewSchoolYearRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus, 
                readRepos func (string) eventhorizon.ReadWriteRepo) (ret *SchoolYearRouter) {
    pathPrefix = pathPrefix + "/" + "schoolYears"
    queryRepository := NewSchoolYearQueryRepository()
    queryHandler := NewSchoolYearHttpQueryHandler(queryRepository)
    commandHandler := NewSchoolYearHttpCommandHandler(context, commandBus)
    ret = &SchoolYearRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *SchoolYearRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountSchoolYearById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("CountSchoolYearAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistSchoolYearById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("ExistSchoolYearAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindSchoolYearById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("FindSchoolYearAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateSchoolYear").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateSchoolYear").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteSchoolYear").HandlerFunc(o.CommandHandler.Delete)
    return
}


type StudentRouter struct {
    PathPrefix string
    AttendanceRouter *AttendanceRouter
    CourseRouter *CourseRouter
    GradeRouter *GradeRouter
    GroupRouter *GroupRouter
    SchoolApplicationRouter *SchoolApplicationRouter
    SchoolYearRouter *SchoolYearRouter
    Router *mux.Router
}

func NewStudentRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus, 
                readRepos func (string) eventhorizon.ReadWriteRepo) (ret *StudentRouter) {
    pathPrefix = pathPrefix + "/" + "student"
    attendanceRouter := NewAttendanceRouter(pathPrefix, context, commandBus, readRepos)
    courseRouter := NewCourseRouter(pathPrefix, context, commandBus, readRepos)
    gradeRouter := NewGradeRouter(pathPrefix, context, commandBus, readRepos)
    groupRouter := NewGroupRouter(pathPrefix, context, commandBus, readRepos)
    schoolApplicationRouter := NewSchoolApplicationRouter(pathPrefix, context, commandBus, readRepos)
    schoolYearRouter := NewSchoolYearRouter(pathPrefix, context, commandBus, readRepos)
    ret = &StudentRouter{
        PathPrefix: pathPrefix,
        AttendanceRouter: attendanceRouter,
        CourseRouter: courseRouter,
        GradeRouter: gradeRouter,
        GroupRouter: groupRouter,
        SchoolApplicationRouter: schoolApplicationRouter,
        SchoolYearRouter: schoolYearRouter,
    }
    return
}

func (o *StudentRouter) Setup(router *mux.Router) (ret error) {
    if ret = o.AttendanceRouter.Setup(router); ret != nil {
        return
    }
    if ret = o.CourseRouter.Setup(router); ret != nil {
        return
    }
    if ret = o.GradeRouter.Setup(router); ret != nil {
        return
    }
    if ret = o.GroupRouter.Setup(router); ret != nil {
        return
    }
    if ret = o.SchoolApplicationRouter.Setup(router); ret != nil {
        return
    }
    if ret = o.SchoolYearRouter.Setup(router); ret != nil {
        return
    }
    return
}









