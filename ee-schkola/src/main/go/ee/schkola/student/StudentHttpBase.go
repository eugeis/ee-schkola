package student

import (
    "github.com/eugeis/gee/net"
    "github.com/gorilla/mux"
    "net/http"
)
type AttendanceHttpQueryHandler struct {
}

func NewAttendanceHttpQueryHandler() (ret *AttendanceHttpQueryHandler) {
    ret = &AttendanceHttpQueryHandler{}
    return
}

func (o *AttendanceHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *AttendanceHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
}

func (o *AttendanceHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
}

func (o *AttendanceHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *AttendanceHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *AttendanceHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
}


type AttendanceHttpCommandHandler struct {
}

func NewAttendanceHttpCommandHandler() (ret *AttendanceHttpCommandHandler) {
    ret = &AttendanceHttpCommandHandler{}
    return
}

func (o *AttendanceHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
}

func (o *AttendanceHttpCommandHandler) Register(w http.ResponseWriter, r *http.Request)  {
}

func (o *AttendanceHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
}

func (o *AttendanceHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
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

func NewAttendanceRouter(PathPrefix string) (ret *AttendanceRouter) {
    
    ret.PathPrefix = ret.PathPrefix + "/" + "attendances"
    return
}

func (o *AttendanceRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindAttendanceAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindAttendanceById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountAttendanceById").HandlerFunc(o.QueryHandler.CountById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountAttendanceAll").HandlerFunc(o.QueryHandler.CountAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistAttendanceAll").HandlerFunc(o.QueryHandler.ExistAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistAttendanceById").HandlerFunc(o.QueryHandler.ExistById)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateAttendance").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("RegisterAttendance").HandlerFunc(o.CommandHandler.Register)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateAttendance").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteAttendance").HandlerFunc(o.CommandHandler.Delete)
    return
}


type CourseHttpQueryHandler struct {
}

func NewCourseHttpQueryHandler() (ret *CourseHttpQueryHandler) {
    ret = &CourseHttpQueryHandler{}
    return
}

func (o *CourseHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *CourseHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
}

func (o *CourseHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
}

func (o *CourseHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *CourseHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *CourseHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
}


type CourseHttpCommandHandler struct {
}

func NewCourseHttpCommandHandler() (ret *CourseHttpCommandHandler) {
    ret = &CourseHttpCommandHandler{}
    return
}

func (o *CourseHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
}

func (o *CourseHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
}

func (o *CourseHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
}


type CourseRouter struct {
    PathPrefix string
    QueryHandler *CourseHttpQueryHandler
    CommandHandler *CourseHttpCommandHandler
    Router *mux.Router
}

func NewCourseRouter(PathPrefix string) (ret *CourseRouter) {
    
    ret.PathPrefix = ret.PathPrefix + "/" + "courses"
    return
}

func (o *CourseRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindCourseAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindCourseById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountCourseById").HandlerFunc(o.QueryHandler.CountById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountCourseAll").HandlerFunc(o.QueryHandler.CountAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistCourseAll").HandlerFunc(o.QueryHandler.ExistAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistCourseById").HandlerFunc(o.QueryHandler.ExistById)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateCourse").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateCourse").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteCourse").HandlerFunc(o.CommandHandler.Delete)
    return
}


type GradeHttpQueryHandler struct {
}

func NewGradeHttpQueryHandler() (ret *GradeHttpQueryHandler) {
    ret = &GradeHttpQueryHandler{}
    return
}

func (o *GradeHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *GradeHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
}

func (o *GradeHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
}

func (o *GradeHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *GradeHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *GradeHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
}


type GradeHttpCommandHandler struct {
}

func NewGradeHttpCommandHandler() (ret *GradeHttpCommandHandler) {
    ret = &GradeHttpCommandHandler{}
    return
}

func (o *GradeHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
}

func (o *GradeHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
}

func (o *GradeHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
}


type GradeRouter struct {
    PathPrefix string
    QueryHandler *GradeHttpQueryHandler
    CommandHandler *GradeHttpCommandHandler
    Router *mux.Router
}

func NewGradeRouter(PathPrefix string) (ret *GradeRouter) {
    
    ret.PathPrefix = ret.PathPrefix + "/" + "grades"
    return
}

func (o *GradeRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindGradeAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindGradeById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountGradeById").HandlerFunc(o.QueryHandler.CountById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountGradeAll").HandlerFunc(o.QueryHandler.CountAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistGradeAll").HandlerFunc(o.QueryHandler.ExistAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistGradeById").HandlerFunc(o.QueryHandler.ExistById)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateGrade").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateGrade").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteGrade").HandlerFunc(o.CommandHandler.Delete)
    return
}


type GroupHttpQueryHandler struct {
}

func NewGroupHttpQueryHandler() (ret *GroupHttpQueryHandler) {
    ret = &GroupHttpQueryHandler{}
    return
}

func (o *GroupHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *GroupHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
}

func (o *GroupHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
}

func (o *GroupHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *GroupHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *GroupHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
}


type GroupHttpCommandHandler struct {
}

func NewGroupHttpCommandHandler() (ret *GroupHttpCommandHandler) {
    ret = &GroupHttpCommandHandler{}
    return
}

func (o *GroupHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
}

func (o *GroupHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
}

func (o *GroupHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
}


type GroupRouter struct {
    PathPrefix string
    QueryHandler *GroupHttpQueryHandler
    CommandHandler *GroupHttpCommandHandler
    Router *mux.Router
}

func NewGroupRouter(PathPrefix string) (ret *GroupRouter) {
    
    ret.PathPrefix = ret.PathPrefix + "/" + "groups"
    return
}

func (o *GroupRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindGroupAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindGroupById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountGroupById").HandlerFunc(o.QueryHandler.CountById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountGroupAll").HandlerFunc(o.QueryHandler.CountAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistGroupAll").HandlerFunc(o.QueryHandler.ExistAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistGroupById").HandlerFunc(o.QueryHandler.ExistById)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateGroup").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateGroup").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteGroup").HandlerFunc(o.CommandHandler.Delete)
    return
}


type SchoolApplicationHttpQueryHandler struct {
}

func NewSchoolApplicationHttpQueryHandler() (ret *SchoolApplicationHttpQueryHandler) {
    ret = &SchoolApplicationHttpQueryHandler{}
    return
}

func (o *SchoolApplicationHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *SchoolApplicationHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
}

func (o *SchoolApplicationHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
}

func (o *SchoolApplicationHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *SchoolApplicationHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *SchoolApplicationHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
}


type SchoolApplicationHttpCommandHandler struct {
}

func NewSchoolApplicationHttpCommandHandler() (ret *SchoolApplicationHttpCommandHandler) {
    ret = &SchoolApplicationHttpCommandHandler{}
    return
}

func (o *SchoolApplicationHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
}

func (o *SchoolApplicationHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
}

func (o *SchoolApplicationHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
}


type SchoolApplicationRouter struct {
    PathPrefix string
    QueryHandler *SchoolApplicationHttpQueryHandler
    CommandHandler *SchoolApplicationHttpCommandHandler
    Router *mux.Router
}

func NewSchoolApplicationRouter(PathPrefix string) (ret *SchoolApplicationRouter) {
    
    ret.PathPrefix = ret.PathPrefix + "/" + "schoolApplications"
    return
}

func (o *SchoolApplicationRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindSchoolApplicationAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindSchoolApplicationById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountSchoolApplicationById").HandlerFunc(o.QueryHandler.CountById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountSchoolApplicationAll").HandlerFunc(o.QueryHandler.CountAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistSchoolApplicationAll").HandlerFunc(o.QueryHandler.ExistAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistSchoolApplicationById").HandlerFunc(o.QueryHandler.ExistById)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateSchoolApplication").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateSchoolApplication").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteSchoolApplication").HandlerFunc(o.CommandHandler.Delete)
    return
}


type SchoolYearHttpQueryHandler struct {
}

func NewSchoolYearHttpQueryHandler() (ret *SchoolYearHttpQueryHandler) {
    ret = &SchoolYearHttpQueryHandler{}
    return
}

func (o *SchoolYearHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *SchoolYearHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
}

func (o *SchoolYearHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
}

func (o *SchoolYearHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *SchoolYearHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *SchoolYearHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
}


type SchoolYearHttpCommandHandler struct {
}

func NewSchoolYearHttpCommandHandler() (ret *SchoolYearHttpCommandHandler) {
    ret = &SchoolYearHttpCommandHandler{}
    return
}

func (o *SchoolYearHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
}

func (o *SchoolYearHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
}

func (o *SchoolYearHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
}


type SchoolYearRouter struct {
    PathPrefix string
    QueryHandler *SchoolYearHttpQueryHandler
    CommandHandler *SchoolYearHttpCommandHandler
    Router *mux.Router
}

func NewSchoolYearRouter(PathPrefix string) (ret *SchoolYearRouter) {
    
    ret.PathPrefix = ret.PathPrefix + "/" + "schoolYears"
    return
}

func (o *SchoolYearRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindSchoolYearAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindSchoolYearById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountSchoolYearById").HandlerFunc(o.QueryHandler.CountById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountSchoolYearAll").HandlerFunc(o.QueryHandler.CountAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistSchoolYearAll").HandlerFunc(o.QueryHandler.ExistAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistSchoolYearById").HandlerFunc(o.QueryHandler.ExistById)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateSchoolYear").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateSchoolYear").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteSchoolYear").HandlerFunc(o.CommandHandler.Delete)
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

func NewStudentRouter(PathPrefix string) (ret *StudentRouter) {
    
    ret.PathPrefix = ret.PathPrefix + "/" + "student"
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



