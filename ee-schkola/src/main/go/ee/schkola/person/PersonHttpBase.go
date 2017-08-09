package person

import (
    "github.com/eugeis/gee/net"
    "github.com/gorilla/mux"
    "net/http"
)

type ChurchHttpQueryHandler struct {
}

func NewChurchHttpQueryHandler() (ret *ChurchHttpQueryHandler) {
    ret = &ChurchHttpQueryHandler{}
    return
}



type ChurchHttpCommandHandler struct {
}

func NewChurchHttpCommandHandler() (ret *ChurchHttpCommandHandler) {
    ret = &ChurchHttpCommandHandler{}
    return
}

func (o *ChurchHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *ChurchHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *ChurchHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}



type ChurchRouter struct {
    Router  *mux.Router
    PathPrefix  string
    QueryHandler  *ChurchHttpQueryHandler
    CommandHandler  *ChurchHttpCommandHandler
}

func (o *ChurchRouter) Setup() (ret error) {
            
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("Create").HandlerFunc(o.CommandHandler.Create)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("Update").HandlerFunc(o.CommandHandler.Update)
    o.Router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("Delete").HandlerFunc(o.CommandHandler.Delete)
    return
    
}



type GraduationHttpQueryHandler struct {
}

func NewGraduationHttpQueryHandler() (ret *GraduationHttpQueryHandler) {
    ret = &GraduationHttpQueryHandler{}
    return
}



type GraduationHttpCommandHandler struct {
}

func NewGraduationHttpCommandHandler() (ret *GraduationHttpCommandHandler) {
    ret = &GraduationHttpCommandHandler{}
    return
}

func (o *GraduationHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *GraduationHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *GraduationHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}



type GraduationRouter struct {
    Router  *mux.Router
    PathPrefix  string
    QueryHandler  *GraduationHttpQueryHandler
    CommandHandler  *GraduationHttpCommandHandler
}

func (o *GraduationRouter) Setup() (ret error) {
            
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("Create").HandlerFunc(o.CommandHandler.Create)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("Update").HandlerFunc(o.CommandHandler.Update)
    o.Router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("Delete").HandlerFunc(o.CommandHandler.Delete)
    return
    
}



type ProfileHttpQueryHandler struct {
}

func NewProfileHttpQueryHandler() (ret *ProfileHttpQueryHandler) {
    ret = &ProfileHttpQueryHandler{}
    return
}

func (o *ProfileHttpQueryHandler) FindByName(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *ProfileHttpQueryHandler) FindByEmail(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *ProfileHttpQueryHandler) FindByPhone(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}



type ProfileHttpCommandHandler struct {
}

func NewProfileHttpCommandHandler() (ret *ProfileHttpCommandHandler) {
    ret = &ProfileHttpCommandHandler{}
    return
}

func (o *ProfileHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *ProfileHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *ProfileHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}



type ProfileRouter struct {
    Router  *mux.Router
    PathPrefix  string
    QueryHandler  *ProfileHttpQueryHandler
    CommandHandler  *ProfileHttpCommandHandler
}

func (o *ProfileRouter) Setup() (ret error) {
            
    o.Router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindByName").HandlerFunc(o.QueryHandler.FindByName), 
    o.Router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindByEmail").HandlerFunc(o.QueryHandler.FindByEmail), 
    o.Router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindByPhone").HandlerFunc(o.QueryHandler.FindByPhone)
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("Create").HandlerFunc(o.CommandHandler.Create)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("Update").HandlerFunc(o.CommandHandler.Update)
    o.Router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("Delete").HandlerFunc(o.CommandHandler.Delete)
    return
    
}



