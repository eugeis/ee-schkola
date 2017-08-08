package auth

import (
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
            
    return
    
}



