package student

import (
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
    Router  *mux.Router
    PathPrefix  string
    QueryHandler  *AttendanceHttpQueryHandler
    CommandHandler  *AttendanceHttpCommandHandler
}

func (o *AttendanceRouter) Setup() (ret error) {
            
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
    Router  *mux.Router
    PathPrefix  string
    QueryHandler  *CourseHttpQueryHandler
    CommandHandler  *CourseHttpCommandHandler
}

func (o *CourseRouter) Setup() (ret error) {
            
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
    Router  *mux.Router
    PathPrefix  string
    QueryHandler  *GradeHttpQueryHandler
    CommandHandler  *GradeHttpCommandHandler
}

func (o *GradeRouter) Setup() (ret error) {
            
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
    Router  *mux.Router
    PathPrefix  string
    QueryHandler  *GroupHttpQueryHandler
    CommandHandler  *GroupHttpCommandHandler
}

func (o *GroupRouter) Setup() (ret error) {
            
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
    Router  *mux.Router
    PathPrefix  string
    QueryHandler  *SchoolApplicationHttpQueryHandler
    CommandHandler  *SchoolApplicationHttpCommandHandler
}

func (o *SchoolApplicationRouter) Setup() (ret error) {
            
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
    Router  *mux.Router
    PathPrefix  string
    QueryHandler  *SchoolYearHttpQueryHandler
    CommandHandler  *SchoolYearHttpCommandHandler
}

func (o *SchoolYearRouter) Setup() (ret error) {
            
    return
    
}



