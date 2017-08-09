package library

import (
    "github.com/eugeis/gee/net"
    "github.com/gorilla/mux"
    "net/http"
)

type BookHttpQueryHandler struct {
}

func NewBookHttpQueryHandler() (ret *BookHttpQueryHandler) {
    ret = &BookHttpQueryHandler{}
    return
}

func (o *BookHttpQueryHandler) FindByTitle(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *BookHttpQueryHandler) FindByAuthor(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *BookHttpQueryHandler) FindByPattern(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}



type BookHttpCommandHandler struct {
}

func NewBookHttpCommandHandler() (ret *BookHttpCommandHandler) {
    ret = &BookHttpCommandHandler{}
    return
}

func (o *BookHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *BookHttpCommandHandler) Register(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *BookHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *BookHttpCommandHandler) Change(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *BookHttpCommandHandler) ChangeLocation(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *BookHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *BookHttpCommandHandler) Unregister(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}



type BookRouter struct {
    Router  *mux.Router
    PathPrefix  string
    QueryHandler  *BookHttpQueryHandler
    CommandHandler  *BookHttpCommandHandler
}

func (o *BookRouter) Setup() (ret error) {
            
    o.Router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindByTitle").HandlerFunc(o.QueryHandler.FindByTitle), 
    o.Router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindByAuthor").HandlerFunc(o.QueryHandler.FindByAuthor), 
    o.Router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindByPattern").HandlerFunc(o.QueryHandler.FindByPattern)
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("Create").HandlerFunc(o.CommandHandler.Create)
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("Register").HandlerFunc(o.CommandHandler.Register)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("Update").HandlerFunc(o.CommandHandler.Update)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("Change").HandlerFunc(o.CommandHandler.Change)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("ChangeLocation").HandlerFunc(o.CommandHandler.ChangeLocation)
    o.Router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("Delete").HandlerFunc(o.CommandHandler.Delete)
    o.Router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("Unregister").HandlerFunc(o.CommandHandler.Unregister)
    return
    
}



