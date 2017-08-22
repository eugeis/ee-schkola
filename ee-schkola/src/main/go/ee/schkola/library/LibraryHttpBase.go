package library

import (
    "context"
    "fmt"
    "github.com/eugeis/gee/eh"
    "github.com/eugeis/gee/net"
    "github.com/gorilla/mux"
    "github.com/looplab/eventhorizon"
    "html"
    "net/http"
)
type BookHttpQueryHandler struct {
}

func NewBookHttpQueryHandler() (ret *BookHttpQueryHandler) {
    ret = &BookHttpQueryHandler{}
    return
}

func (o *BookHttpQueryHandler) FindByTitle(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindByBookTitle", html.EscapeString(r.URL.Path))
}

func (o *BookHttpQueryHandler) FindByAuthor(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindByBookAuthor", html.EscapeString(r.URL.Path))
}

func (o *BookHttpQueryHandler) FindByPattern(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindByBookPattern", html.EscapeString(r.URL.Path))
}

func (o *BookHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindAllBook", html.EscapeString(r.URL.Path))
}

func (o *BookHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindByBookId", html.EscapeString(r.URL.Path))
}

func (o *BookHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountAllBook", html.EscapeString(r.URL.Path))
}

func (o *BookHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountByBookId", html.EscapeString(r.URL.Path))
}

func (o *BookHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistAllBook", html.EscapeString(r.URL.Path))
}

func (o *BookHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistByBookId", html.EscapeString(r.URL.Path))
}


type BookHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewBookHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *BookHttpCommandHandler) {
    ret = &BookHttpCommandHandler{
        HttpCommandHandler: eh.NewHttpCommandHandler(context, commandBus),
    }
    return
}

func (o *BookHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateBook{Id: id}, w, r)
}

func (o *BookHttpCommandHandler) ChangeLocation(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&ChangeBookLocation{Id: id}, w, r)
}

func (o *BookHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateBook{Id: id}, w, r)
}

func (o *BookHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteBook{Id: id}, w, r)
}


type BookRouter struct {
    PathPrefix string
    QueryHandler *BookHttpQueryHandler
    CommandHandler *BookHttpCommandHandler
    Router *mux.Router
}

func NewBookRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus) (ret *BookRouter) {
    pathPrefix = pathPrefix + "/" + "books"
    ret = &BookRouter{
        PathPrefix: pathPrefix,
        QueryHandler: NewBookHttpQueryHandler(),
        CommandHandler: NewBookHttpCommandHandler(context, commandBus),
    }
    return
}

func (o *BookRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountBookById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("CountBookAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistBookById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("ExistBookAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("FindBookByTitle").HandlerFunc(o.QueryHandler.FindByTitle).
    Queries("title", "{title}")
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("FindBookByAuthor").HandlerFunc(o.QueryHandler.FindByAuthor).
    Queries("author", "{author}")
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("FindBookByPattern").HandlerFunc(o.QueryHandler.FindByPattern).
    Queries("pattern", "{pattern}")
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindBookById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("FindBookAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateBook").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ChangeBookLocation").HandlerFunc(o.CommandHandler.ChangeLocation)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateBook").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteBook").HandlerFunc(o.CommandHandler.Delete)
    return
}


type LibraryRouter struct {
    PathPrefix string
    BookRouter *BookRouter
    Router *mux.Router
}

func NewLibraryRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus) (ret *LibraryRouter) {
    pathPrefix = pathPrefix + "/" + "library"
    ret = &LibraryRouter{
        PathPrefix: pathPrefix,
        BookRouter: NewBookRouter(pathPrefix, context, commandBus),
    }
    return
}

func (o *LibraryRouter) Setup(router *mux.Router) (ret error) {
    if ret = o.BookRouter.Setup(router); ret != nil {
        return
    }
    return
}









