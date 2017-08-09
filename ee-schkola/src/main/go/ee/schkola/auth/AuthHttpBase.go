package auth

import (
    "github.com/eugeis/gee/net"
    "github.com/gorilla/mux"
    "net/http"
)

type AccountHttpQueryHandler struct {
}

func NewAccountHttpQueryHandler() (ret *AccountHttpQueryHandler) {
    ret = &AccountHttpQueryHandler{}
    return
}



type AccountHttpCommandHandler struct {
}

func NewAccountHttpCommandHandler() (ret *AccountHttpCommandHandler) {
    ret = &AccountHttpCommandHandler{}
    return
}

func (o *AccountHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *AccountHttpCommandHandler) Register(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *AccountHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *AccountHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *AccountHttpCommandHandler) Enable(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *AccountHttpCommandHandler) Disable(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}



type AccountRouter struct {
    Router  *mux.Router
    PathPrefix  string
    QueryHandler  *AccountHttpQueryHandler
    CommandHandler  *AccountHttpCommandHandler
}

func (o *AccountRouter) Setup() (ret error) {
            
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("Create").HandlerFunc(o.CommandHandler.Create)
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("Register").HandlerFunc(o.CommandHandler.Register)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("Update").HandlerFunc(o.CommandHandler.Update)
    o.Router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("Delete").HandlerFunc(o.CommandHandler.Delete)
    return
    
}



