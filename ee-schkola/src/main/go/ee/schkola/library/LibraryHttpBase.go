package library

import (
    "context"
    "github.com/eugeis/gee/eh"
    "github.com/eugeis/gee/net"
    "github.com/gorilla/mux"
    "github.com/looplab/eventhorizon"
    "github.com/looplab/eventhorizon/commandhandler/bus"
    "net/http"
)
type HttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *QueryRepository `json:"queryRepository" eh:"optional"`
}

func NewBookHttpQueryHandler(queryRepository *QueryRepository) (ret *HttpQueryHandler) {
    httpQueryHandler := &eh.HttpQueryHandler{}
    ret = &HttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *HttpQueryHandler) FindByTitle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    title := vars["title"]
    ret, err := o.QueryRepository.FindByTitle(title)
    o.HandleResult(ret, err, "FindByBookTitle", w, r)
}

func (o *HttpQueryHandler) FindByPattern(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    pattern := vars["pattern"]
    ret, err := o.QueryRepository.FindByPattern(pattern)
    o.HandleResult(ret, err, "FindByBookPattern", w, r)
}

func (o *HttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllBook", w, r)
}

func (o *HttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindByBookId", w, r)
}

func (o *HttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllBook", w, r)
}

func (o *HttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountByBookId", w, r)
}

func (o *HttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllBook", w, r)
}

func (o *HttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistByBookId", w, r)
}


type HttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewBookHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandHandler) (ret *HttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &HttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *HttpCommandHandler) Create(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateBook{Id: id}, w, r)
}

func (o *HttpCommandHandler) ChangeLocation(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&ChangeBookLocation{Id: id}, w, r)
}

func (o *HttpCommandHandler) Update(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateBook{Id: id}, w, r)
}

func (o *HttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteBook{Id: id}, w, r)
}


type Router struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    QueryHandler *HttpQueryHandler `json:"queryHandler" eh:"optional"`
    CommandHandler *HttpCommandHandler `json:"commandHandler" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func NewBookRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *Router) {
    pathPrefix = pathPrefix + "/" + "books"
    entityFactory := func() eventhorizon.Entity { return NewBook() }
    repo := readRepos(string(BookAggregateType), entityFactory)
    queryRepository := NewBookQueryRepository(repo, context)
    queryHandler := NewBookHttpQueryHandler(queryRepository)
    commandHandler := NewBookHttpCommandHandler(context, commandBus)
    ret = &Router{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *Router) Setup(router *mux.Router) (err error) {
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountBookById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("CountBookAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistBookById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("ExistBookAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("FindBookByTitle").HandlerFunc(o.QueryHandler.FindByTitle).
    Queries("title", "{title}")
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("FindBookByPattern").HandlerFunc(o.QueryHandler.FindByPattern).
    Queries("pattern", "{pattern}")
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindBookById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("FindBookAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateBook").HandlerFunc(o.CommandHandler.Create)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Queries(net.Command, "changeLocation").
        Name("ChangeBookLocation").HandlerFunc(o.CommandHandler.ChangeLocation)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateBook").HandlerFunc(o.CommandHandler.Update)
    router.Methods(http.MethodDelete).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteBook").HandlerFunc(o.CommandHandler.Delete)
    return
}


type LibraryRouter struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    BookRouter *Router `json:"bookRouter" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func NewLibraryRouter(pathPrefix string, context context.Context, commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *LibraryRouter) {
    pathPrefix = pathPrefix + "/" + "library"
    bookRouter := NewBookRouter(pathPrefix, context, commandBus, readRepos)
    ret = &LibraryRouter{
        PathPrefix: pathPrefix,
        BookRouter: bookRouter,
    }
    return
}

func (o *LibraryRouter) Setup(router *mux.Router) (err error) {
    if err = o.BookRouter.Setup(router); err != nil {
        return
    }
    return
}









