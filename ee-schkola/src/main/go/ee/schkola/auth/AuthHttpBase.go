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
    PathPrefix  string
    Router  *mux.Router
    QueryHandler  *AccountHttpQueryHandler
    CommandHandler  *AccountHttpCommandHandler
}

func NewAccountRouter(Router *mux.Router) (ret *AccountRouter) {
    ret = &AccountRouter{
        PathPrefix :"Accounts",
        Router :Router,
        QueryHandler :NewAccountHttpQueryHandler(),
        CommandHandler :NewAccountHttpCommandHandler(),
    }
    return
}

func (o *AccountRouter) Setup() (ret error) {
            
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateAccount").HandlerFunc(o.CommandHandler.Create)
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("RegisterAccount").HandlerFunc(o.CommandHandler.Register)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateAccount").HandlerFunc(o.CommandHandler.Update)
    o.Router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteAccount").HandlerFunc(o.CommandHandler.Delete)
    return
    
}



type AuthRouter struct {
    Router  *mux.Router
    PathPrefix  string
    AccountRouter  *AccountRouter
}

func (o *AuthRouter) Setup() (ret error) {
            
    return
}



