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
}

func (o *BookHttpQueryHandler) FindByAuthor(w http.ResponseWriter, r *http.Request)  {
}

func (o *BookHttpQueryHandler) FindByPattern(w http.ResponseWriter, r *http.Request)  {
}


type BookHttpCommandHandler struct {
}

func NewBookHttpCommandHandler() (ret *BookHttpCommandHandler) {
    ret = &BookHttpCommandHandler{}
    return
}

func (o *BookHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
}

func (o *BookHttpCommandHandler) Register(w http.ResponseWriter, r *http.Request)  {
}

func (o *BookHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
}

func (o *BookHttpCommandHandler) Change(w http.ResponseWriter, r *http.Request)  {
}

func (o *BookHttpCommandHandler) ChangeLocation(w http.ResponseWriter, r *http.Request)  {
}

func (o *BookHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
}

func (o *BookHttpCommandHandler) Unregister(w http.ResponseWriter, r *http.Request)  {
}


type BookRouter struct {
    PathPrefix string
    QueryHandler *BookHttpQueryHandler
    CommandHandler *BookHttpCommandHandler
    Router *mux.Router
}

func NewBookRouter(PathPrefix string) (ret *BookRouter) {
    
    ret.PathPrefix = ret.PathPrefix + "/" + "books"
    return
}

func (o *BookRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindBookByTitle").HandlerFunc(o.QueryHandler.FindByTitle)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindBookByAuthor").HandlerFunc(o.QueryHandler.FindByAuthor)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindBookByPattern").HandlerFunc(o.QueryHandler.FindByPattern)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateBook").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("RegisterBook").HandlerFunc(o.CommandHandler.Register)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateBook").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("ChangeBook").HandlerFunc(o.CommandHandler.Change)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("ChangeBookLocation").HandlerFunc(o.CommandHandler.ChangeLocation)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteBook").HandlerFunc(o.CommandHandler.Delete)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("UnregisterBook").HandlerFunc(o.CommandHandler.Unregister)
    return
}


type LibraryRouter struct {
    PathPrefix string
    BookRouter *BookRouter
    Router *mux.Router
}

func NewLibraryRouter(PathPrefix string) (ret *LibraryRouter) {
    
    ret.PathPrefix = ret.PathPrefix + "/" + "library"
    return
}

func (o *LibraryRouter) Setup(router *mux.Router) (ret error) {
    if ret = o.BookRouter.Setup(router); ret != nil {
        return
    }
    return
}



