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
    
    return
}

func (o *BookHttpQueryHandler) FindByAuthor(w http.ResponseWriter, r *http.Request)  {
    return
    
    return
}

func (o *BookHttpQueryHandler) FindByPattern(w http.ResponseWriter, r *http.Request)  {
    return
    
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
    
    return
}

func (o *BookHttpCommandHandler) Register(w http.ResponseWriter, r *http.Request)  {
    return
    
    return
}

func (o *BookHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    return
    
    return
}

func (o *BookHttpCommandHandler) Change(w http.ResponseWriter, r *http.Request)  {
    return
    
    return
}

func (o *BookHttpCommandHandler) ChangeLocation(w http.ResponseWriter, r *http.Request)  {
    return
    
    return
}

func (o *BookHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    return
    
    return
}

func (o *BookHttpCommandHandler) Unregister(w http.ResponseWriter, r *http.Request)  {
    return
    
    return
}


type BookRouter struct {
    PathPrefix string
    Router *mux.Router
    QueryHandler *BookHttpQueryHandler
    CommandHandler *BookHttpCommandHandler
}

func (o *BookRouter) Setup() (ret error) {
    o.Router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindBookByTitle").HandlerFunc(o.QueryHandler.FindByTitle)
    o.Router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindBookByAuthor").HandlerFunc(o.QueryHandler.FindByAuthor)
    o.Router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindBookByPattern").HandlerFunc(o.QueryHandler.FindByPattern)
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateBook").HandlerFunc(o.CommandHandler.Create)
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("RegisterBook").HandlerFunc(o.CommandHandler.Register)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateBook").HandlerFunc(o.CommandHandler.Update)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("ChangeBook").HandlerFunc(o.CommandHandler.Change)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("ChangeBookLocation").HandlerFunc(o.CommandHandler.ChangeLocation)
    o.Router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteBook").HandlerFunc(o.CommandHandler.Delete)
    o.Router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("UnregisterBook").HandlerFunc(o.CommandHandler.Unregister)
    return
    
    return
}


type LibraryRouter struct {
    Router *mux.Router
    PathPrefix string
    BookRouter *BookRouter
}

func (o *LibraryRouter) Setup() (ret error) {
    return
    return
}



