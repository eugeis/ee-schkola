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



type AttendanceHttpCommandHandler struct {
}

func NewAttendanceHttpCommandHandler() (ret *AttendanceHttpCommandHandler) {
    ret = &AttendanceHttpCommandHandler{}
    return
}

func (o *AttendanceHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *AttendanceHttpCommandHandler) Register(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *AttendanceHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *AttendanceHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *AttendanceHttpCommandHandler) Confirm(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *AttendanceHttpCommandHandler) Cancel(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}



type AttendanceRouter struct {
    PathPrefix  string
    Router  *mux.Router
    QueryHandler  *AttendanceHttpQueryHandler
    CommandHandler  *AttendanceHttpCommandHandler
}

func NewAttendanceRouter(Router *mux.Router) (ret *AttendanceRouter) {
    ret = &AttendanceRouter{
        PathPrefix :"Attendances",
        Router :Router,
        QueryHandler :NewAttendanceHttpQueryHandler(),
        CommandHandler :NewAttendanceHttpCommandHandler(),
    }
    return
}

func (o *AttendanceRouter) Setup() (ret error) {
            
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateAttendance").HandlerFunc(o.CommandHandler.Create)
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("RegisterAttendance").HandlerFunc(o.CommandHandler.Register)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateAttendance").HandlerFunc(o.CommandHandler.Update)
    o.Router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteAttendance").HandlerFunc(o.CommandHandler.Delete)
    return
    
}



type CourseHttpQueryHandler struct {
}

func NewCourseHttpQueryHandler() (ret *CourseHttpQueryHandler) {
    ret = &CourseHttpQueryHandler{}
    return
}



type CourseHttpCommandHandler struct {
}

func NewCourseHttpCommandHandler() (ret *CourseHttpCommandHandler) {
    ret = &CourseHttpCommandHandler{}
    return
}

func (o *CourseHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *CourseHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *CourseHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}



type CourseRouter struct {
    PathPrefix  string
    Router  *mux.Router
    QueryHandler  *CourseHttpQueryHandler
    CommandHandler  *CourseHttpCommandHandler
}

func NewCourseRouter(Router *mux.Router) (ret *CourseRouter) {
    ret = &CourseRouter{
        PathPrefix :"Courses",
        Router :Router,
        QueryHandler :NewCourseHttpQueryHandler(),
        CommandHandler :NewCourseHttpCommandHandler(),
    }
    return
}

func (o *CourseRouter) Setup() (ret error) {
            
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateCourse").HandlerFunc(o.CommandHandler.Create)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateCourse").HandlerFunc(o.CommandHandler.Update)
    o.Router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteCourse").HandlerFunc(o.CommandHandler.Delete)
    return
    
}



type GradeHttpQueryHandler struct {
}

func NewGradeHttpQueryHandler() (ret *GradeHttpQueryHandler) {
    ret = &GradeHttpQueryHandler{}
    return
}



type GradeHttpCommandHandler struct {
}

func NewGradeHttpCommandHandler() (ret *GradeHttpCommandHandler) {
    ret = &GradeHttpCommandHandler{}
    return
}

func (o *GradeHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *GradeHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *GradeHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}



type GradeRouter struct {
    PathPrefix  string
    Router  *mux.Router
    QueryHandler  *GradeHttpQueryHandler
    CommandHandler  *GradeHttpCommandHandler
}

func NewGradeRouter(Router *mux.Router) (ret *GradeRouter) {
    ret = &GradeRouter{
        PathPrefix :"Grades",
        Router :Router,
        QueryHandler :NewGradeHttpQueryHandler(),
        CommandHandler :NewGradeHttpCommandHandler(),
    }
    return
}

func (o *GradeRouter) Setup() (ret error) {
            
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateGrade").HandlerFunc(o.CommandHandler.Create)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateGrade").HandlerFunc(o.CommandHandler.Update)
    o.Router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteGrade").HandlerFunc(o.CommandHandler.Delete)
    return
    
}



type GroupHttpQueryHandler struct {
}

func NewGroupHttpQueryHandler() (ret *GroupHttpQueryHandler) {
    ret = &GroupHttpQueryHandler{}
    return
}



type GroupHttpCommandHandler struct {
}

func NewGroupHttpCommandHandler() (ret *GroupHttpCommandHandler) {
    ret = &GroupHttpCommandHandler{}
    return
}

func (o *GroupHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *GroupHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *GroupHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}



type GroupRouter struct {
    PathPrefix  string
    Router  *mux.Router
    QueryHandler  *GroupHttpQueryHandler
    CommandHandler  *GroupHttpCommandHandler
}

func NewGroupRouter(Router *mux.Router) (ret *GroupRouter) {
    ret = &GroupRouter{
        PathPrefix :"Groups",
        Router :Router,
        QueryHandler :NewGroupHttpQueryHandler(),
        CommandHandler :NewGroupHttpCommandHandler(),
    }
    return
}

func (o *GroupRouter) Setup() (ret error) {
            
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateGroup").HandlerFunc(o.CommandHandler.Create)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateGroup").HandlerFunc(o.CommandHandler.Update)
    o.Router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteGroup").HandlerFunc(o.CommandHandler.Delete)
    return
    
}



type SchoolApplicationHttpQueryHandler struct {
}

func NewSchoolApplicationHttpQueryHandler() (ret *SchoolApplicationHttpQueryHandler) {
    ret = &SchoolApplicationHttpQueryHandler{}
    return
}



type SchoolApplicationHttpCommandHandler struct {
}

func NewSchoolApplicationHttpCommandHandler() (ret *SchoolApplicationHttpCommandHandler) {
    ret = &SchoolApplicationHttpCommandHandler{}
    return
}

func (o *SchoolApplicationHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *SchoolApplicationHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *SchoolApplicationHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}



type SchoolApplicationRouter struct {
    PathPrefix  string
    Router  *mux.Router
    QueryHandler  *SchoolApplicationHttpQueryHandler
    CommandHandler  *SchoolApplicationHttpCommandHandler
}

func NewSchoolApplicationRouter(Router *mux.Router) (ret *SchoolApplicationRouter) {
    ret = &SchoolApplicationRouter{
        PathPrefix :"SchoolApplications",
        Router :Router,
        QueryHandler :NewSchoolApplicationHttpQueryHandler(),
        CommandHandler :NewSchoolApplicationHttpCommandHandler(),
    }
    return
}

func (o *SchoolApplicationRouter) Setup() (ret error) {
            
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateSchoolApplication").HandlerFunc(o.CommandHandler.Create)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateSchoolApplication").HandlerFunc(o.CommandHandler.Update)
    o.Router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteSchoolApplication").HandlerFunc(o.CommandHandler.Delete)
    return
    
}



type SchoolYearHttpQueryHandler struct {
}

func NewSchoolYearHttpQueryHandler() (ret *SchoolYearHttpQueryHandler) {
    ret = &SchoolYearHttpQueryHandler{}
    return
}



type SchoolYearHttpCommandHandler struct {
}

func NewSchoolYearHttpCommandHandler() (ret *SchoolYearHttpCommandHandler) {
    ret = &SchoolYearHttpCommandHandler{}
    return
}

func (o *SchoolYearHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *SchoolYearHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *SchoolYearHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}



type SchoolYearRouter struct {
    PathPrefix  string
    Router  *mux.Router
    QueryHandler  *SchoolYearHttpQueryHandler
    CommandHandler  *SchoolYearHttpCommandHandler
}

func NewSchoolYearRouter(Router *mux.Router) (ret *SchoolYearRouter) {
    ret = &SchoolYearRouter{
        PathPrefix :"SchoolYears",
        Router :Router,
        QueryHandler :NewSchoolYearHttpQueryHandler(),
        CommandHandler :NewSchoolYearHttpCommandHandler(),
    }
    return
}

func (o *SchoolYearRouter) Setup() (ret error) {
            
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateSchoolYear").HandlerFunc(o.CommandHandler.Create)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateSchoolYear").HandlerFunc(o.CommandHandler.Update)
    o.Router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteSchoolYear").HandlerFunc(o.CommandHandler.Delete)
    return
    
}



type StudentRouter struct {
    Router  *mux.Router
    PathPrefix  string
    AttendanceRouter  *AttendanceRouter
    CourseRouter  *CourseRouter
    GradeRouter  *GradeRouter
    GroupRouter  *GroupRouter
    SchoolApplicationRouter  *SchoolApplicationRouter
    SchoolYearRouter  *SchoolYearRouter
}

func (o *StudentRouter) Setup() (ret error) {
            
    return
}



