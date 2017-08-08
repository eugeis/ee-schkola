package library

import (
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
            
    return
    
}



