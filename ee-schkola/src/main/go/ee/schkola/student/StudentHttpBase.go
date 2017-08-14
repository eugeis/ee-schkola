package student

import (
    "context"
    "fmt"
    "github.com/eugeis/gee/eh"
    "github.com/eugeis/gee/net"
    "github.com/gorilla/mux"
    "github.com/looplab/eventhorizon"
    "html"
    "net/http"
)
type AttendanceHttpQueryHandler struct {
}

func NewAttendanceHttpQueryHandler() (ret *AttendanceHttpQueryHandler) {
    ret = &AttendanceHttpQueryHandler{}
    return
}

func (o *AttendanceHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindAllAttendance", html.EscapeString(r.URL.Path))
}

func (o *AttendanceHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindByAttendanceId", html.EscapeString(r.URL.Path))
}

func (o *AttendanceHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountAllAttendance", html.EscapeString(r.URL.Path))
}

func (o *AttendanceHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountByAttendanceId", html.EscapeString(r.URL.Path))
}

func (o *AttendanceHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistAllAttendance", html.EscapeString(r.URL.Path))
}

func (o *AttendanceHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistByAttendanceId", html.EscapeString(r.URL.Path))
}


type AttendanceHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewAttendanceHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *AttendanceHttpCommandHandler) {
    ret = &AttendanceHttpCommandHandler{
        HttpCommandHandler: eh.NewHttpCommandHandler(context, commandBus),
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

func NewAttendanceRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus) (ret *AttendanceRouter) {
    pathPrefix = pathPrefix + "/" + "attendances"
    ret = &AttendanceRouter{
        PathPrefix: pathPrefix,
        QueryHandler: NewAttendanceHttpQueryHandler(),
        CommandHandler: NewAttendanceHttpCommandHandler(context, commandBus),
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
}

func NewCourseHttpQueryHandler() (ret *CourseHttpQueryHandler) {
    ret = &CourseHttpQueryHandler{}
    return
}

func (o *CourseHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindAllCourse", html.EscapeString(r.URL.Path))
}

func (o *CourseHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindByCourseId", html.EscapeString(r.URL.Path))
}

func (o *CourseHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountAllCourse", html.EscapeString(r.URL.Path))
}

func (o *CourseHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountByCourseId", html.EscapeString(r.URL.Path))
}

func (o *CourseHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistAllCourse", html.EscapeString(r.URL.Path))
}

func (o *CourseHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistByCourseId", html.EscapeString(r.URL.Path))
}


type CourseHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewCourseHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *CourseHttpCommandHandler) {
    ret = &CourseHttpCommandHandler{
        HttpCommandHandler: eh.NewHttpCommandHandler(context, commandBus),
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

func NewCourseRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus) (ret *CourseRouter) {
    pathPrefix = pathPrefix + "/" + "courses"
    ret = &CourseRouter{
        PathPrefix: pathPrefix,
        QueryHandler: NewCourseHttpQueryHandler(),
        CommandHandler: NewCourseHttpCommandHandler(context, commandBus),
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
}

func NewGradeHttpQueryHandler() (ret *GradeHttpQueryHandler) {
    ret = &GradeHttpQueryHandler{}
    return
}

func (o *GradeHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindAllGrade", html.EscapeString(r.URL.Path))
}

func (o *GradeHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindByGradeId", html.EscapeString(r.URL.Path))
}

func (o *GradeHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountAllGrade", html.EscapeString(r.URL.Path))
}

func (o *GradeHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountByGradeId", html.EscapeString(r.URL.Path))
}

func (o *GradeHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistAllGrade", html.EscapeString(r.URL.Path))
}

func (o *GradeHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistByGradeId", html.EscapeString(r.URL.Path))
}


type GradeHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewGradeHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *GradeHttpCommandHandler) {
    ret = &GradeHttpCommandHandler{
        HttpCommandHandler: eh.NewHttpCommandHandler(context, commandBus),
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

func NewGradeRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus) (ret *GradeRouter) {
    pathPrefix = pathPrefix + "/" + "grades"
    ret = &GradeRouter{
        PathPrefix: pathPrefix,
        QueryHandler: NewGradeHttpQueryHandler(),
        CommandHandler: NewGradeHttpCommandHandler(context, commandBus),
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
}

func NewGroupHttpQueryHandler() (ret *GroupHttpQueryHandler) {
    ret = &GroupHttpQueryHandler{}
    return
}

func (o *GroupHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindAllGroup", html.EscapeString(r.URL.Path))
}

func (o *GroupHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindByGroupId", html.EscapeString(r.URL.Path))
}

func (o *GroupHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountAllGroup", html.EscapeString(r.URL.Path))
}

func (o *GroupHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountByGroupId", html.EscapeString(r.URL.Path))
}

func (o *GroupHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistAllGroup", html.EscapeString(r.URL.Path))
}

func (o *GroupHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistByGroupId", html.EscapeString(r.URL.Path))
}


type GroupHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewGroupHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *GroupHttpCommandHandler) {
    ret = &GroupHttpCommandHandler{
        HttpCommandHandler: eh.NewHttpCommandHandler(context, commandBus),
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

func NewGroupRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus) (ret *GroupRouter) {
    pathPrefix = pathPrefix + "/" + "groups"
    ret = &GroupRouter{
        PathPrefix: pathPrefix,
        QueryHandler: NewGroupHttpQueryHandler(),
        CommandHandler: NewGroupHttpCommandHandler(context, commandBus),
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
}

func NewSchoolApplicationHttpQueryHandler() (ret *SchoolApplicationHttpQueryHandler) {
    ret = &SchoolApplicationHttpQueryHandler{}
    return
}

func (o *SchoolApplicationHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindAllSchoolApplication", html.EscapeString(r.URL.Path))
}

func (o *SchoolApplicationHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindBySchoolApplicationId", html.EscapeString(r.URL.Path))
}

func (o *SchoolApplicationHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountAllSchoolApplication", html.EscapeString(r.URL.Path))
}

func (o *SchoolApplicationHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountBySchoolApplicationId", html.EscapeString(r.URL.Path))
}

func (o *SchoolApplicationHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistAllSchoolApplication", html.EscapeString(r.URL.Path))
}

func (o *SchoolApplicationHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistBySchoolApplicationId", html.EscapeString(r.URL.Path))
}


type SchoolApplicationHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewSchoolApplicationHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *SchoolApplicationHttpCommandHandler) {
    ret = &SchoolApplicationHttpCommandHandler{
        HttpCommandHandler: eh.NewHttpCommandHandler(context, commandBus),
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

func NewSchoolApplicationRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus) (ret *SchoolApplicationRouter) {
    pathPrefix = pathPrefix + "/" + "schoolApplications"
    ret = &SchoolApplicationRouter{
        PathPrefix: pathPrefix,
        QueryHandler: NewSchoolApplicationHttpQueryHandler(),
        CommandHandler: NewSchoolApplicationHttpCommandHandler(context, commandBus),
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
}

func NewSchoolYearHttpQueryHandler() (ret *SchoolYearHttpQueryHandler) {
    ret = &SchoolYearHttpQueryHandler{}
    return
}

func (o *SchoolYearHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindAllSchoolYear", html.EscapeString(r.URL.Path))
}

func (o *SchoolYearHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindBySchoolYearId", html.EscapeString(r.URL.Path))
}

func (o *SchoolYearHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountAllSchoolYear", html.EscapeString(r.URL.Path))
}

func (o *SchoolYearHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountBySchoolYearId", html.EscapeString(r.URL.Path))
}

func (o *SchoolYearHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistAllSchoolYear", html.EscapeString(r.URL.Path))
}

func (o *SchoolYearHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistBySchoolYearId", html.EscapeString(r.URL.Path))
}


type SchoolYearHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewSchoolYearHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *SchoolYearHttpCommandHandler) {
    ret = &SchoolYearHttpCommandHandler{
        HttpCommandHandler: eh.NewHttpCommandHandler(context, commandBus),
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

func NewSchoolYearRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus) (ret *SchoolYearRouter) {
    pathPrefix = pathPrefix + "/" + "schoolYears"
    ret = &SchoolYearRouter{
        PathPrefix: pathPrefix,
        QueryHandler: NewSchoolYearHttpQueryHandler(),
        CommandHandler: NewSchoolYearHttpCommandHandler(context, commandBus),
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

func NewStudentRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus) (ret *StudentRouter) {
    pathPrefix = pathPrefix + "/" + "student"
    ret = &StudentRouter{
        PathPrefix: pathPrefix,
        AttendanceRouter: NewAttendanceRouter(pathPrefix, context, commandBus),
        CourseRouter: NewCourseRouter(pathPrefix, context, commandBus),
        GradeRouter: NewGradeRouter(pathPrefix, context, commandBus),
        GroupRouter: NewGroupRouter(pathPrefix, context, commandBus),
        SchoolApplicationRouter: NewSchoolApplicationRouter(pathPrefix, context, commandBus),
        SchoolYearRouter: NewSchoolYearRouter(pathPrefix, context, commandBus),
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



