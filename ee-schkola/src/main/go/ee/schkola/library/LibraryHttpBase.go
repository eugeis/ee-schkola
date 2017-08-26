package library

import (
    "context"
    "github.com/eugeis/gee/eh"
    "github.com/eugeis/gee/net"
    "github.com/gorilla/mux"
    "github.com/looplab/eventhorizon"
    "net/http"
)
type BookHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *BookQueryRepository
}

func NewBookHttpQueryHandler(queryRepository *BookQueryRepository) (ret *BookHttpQueryHandler) {
    httpQueryHandler := &eh.HttpQueryHandler{}
    ret = &BookHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *BookHttpQueryHandler) FindByTitle(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    title := vars["title"]
    ret, err := o.QueryRepository.FindByTitle(title)
    o.HandleResult(ret, err, "FindByBookTitle", w, r)
}

func (o *BookHttpQueryHandler) FindByAuthor(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    author := vars["author"]
    ret, err := o.QueryRepository.FindByAuthor(author)
    o.HandleResult(ret, err, "FindByBookAuthor", w, r)
}

func (o *BookHttpQueryHandler) FindByPattern(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    pattern := vars["pattern"]
    ret, err := o.QueryRepository.FindByPattern(pattern)
    o.HandleResult(ret, err, "FindByBookPattern", w, r)
}

func (o *BookHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllBook", w, r)
}

func (o *BookHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindByBookId", w, r)
}

func (o *BookHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllBook", w, r)
}

func (o *BookHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountByBookId", w, r)
}

func (o *BookHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllBook", w, r)
}

func (o *BookHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistByBookId", w, r)
}


type BookHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewBookHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *BookHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &BookHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
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

func NewBookRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus, 
                readRepos func (string) eventhorizon.ReadWriteRepo) (ret *BookRouter) {
    pathPrefix = pathPrefix + "/" + "books"
    queryRepository := NewBookQueryRepository()
    queryHandler := NewBookHttpQueryHandler(queryRepository)
    commandHandler := NewBookHttpCommandHandler(context, commandBus)
    ret = &BookRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
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

func NewLibraryRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus, 
                readRepos func (string) eventhorizon.ReadWriteRepo) (ret *LibraryRouter) {
    pathPrefix = pathPrefix + "/" + "library"
    bookRouter := NewBookRouter(pathPrefix, context, commandBus, readRepos)
    ret = &LibraryRouter{
        PathPrefix: pathPrefix,
        BookRouter: bookRouter,
    }
    return
}

func (o *LibraryRouter) Setup(router *mux.Router) (ret error) {
    if ret = o.BookRouter.Setup(router); ret != nil {
        return
    }
    return
}









