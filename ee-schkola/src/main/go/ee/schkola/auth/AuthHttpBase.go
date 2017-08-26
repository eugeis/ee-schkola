package auth

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
type AccountHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *AccountQueryRepository
}

func NewAccountHttpQueryHandler(queryRepository *AccountQueryRepository) (ret *AccountHttpQueryHandler) {
    httpQueryHandler := &eh.HttpQueryHandler{}
    ret = &AccountHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *AccountHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindAllAccount", html.EscapeString(r.URL.Path))
}

func (o *AccountHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindByAccountId", html.EscapeString(r.URL.Path))
}

func (o *AccountHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountAllAccount", html.EscapeString(r.URL.Path))
}

func (o *AccountHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountByAccountId", html.EscapeString(r.URL.Path))
}

func (o *AccountHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistAllAccount", html.EscapeString(r.URL.Path))
}

func (o *AccountHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistByAccountId", html.EscapeString(r.URL.Path))
}


type AccountHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewAccountHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *AccountHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &AccountHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *AccountHttpCommandHandler) Register(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&RegisterAccount{Id: id}, w, r)
}

func (o *AccountHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateAccount{Id: id}, w, r)
}

func (o *AccountHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateAccount{Id: id}, w, r)
}

func (o *AccountHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteAccount{Id: id}, w, r)
}

func (o *AccountHttpCommandHandler) Enable(w http.ResponseWriter, r *http.Request)  {
}

func (o *AccountHttpCommandHandler) Disable(w http.ResponseWriter, r *http.Request)  {
}


type AccountRouter struct {
    PathPrefix string
    QueryHandler *AccountHttpQueryHandler
    CommandHandler *AccountHttpCommandHandler
    Router *mux.Router
}

func NewAccountRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus, 
                readRepos func (string) eventhorizon.ReadWriteRepo) (ret *AccountRouter) {
    pathPrefix = pathPrefix + "/" + "accounts"
    queryRepository := NewAccountQueryRepository()
    queryHandler := NewAccountHttpQueryHandler(queryRepository)
    commandHandler := NewAccountHttpCommandHandler(context, commandBus)
    ret = &AccountRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *AccountRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountAccountById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("CountAccountAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistAccountById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("ExistAccountAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindAccountById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("FindAccountAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("RegisterAccount").HandlerFunc(o.CommandHandler.Register)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateAccount").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateAccount").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteAccount").HandlerFunc(o.CommandHandler.Delete)
    return
}


type AuthRouter struct {
    PathPrefix string
    AccountRouter *AccountRouter
    Router *mux.Router
}

func NewAuthRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus, 
                readRepos func (string) eventhorizon.ReadWriteRepo) (ret *AuthRouter) {
    pathPrefix = pathPrefix + "/" + "auth"
    accountRouter := NewAccountRouter(pathPrefix, context, commandBus, readRepos)
    ret = &AuthRouter{
        PathPrefix: pathPrefix,
        AccountRouter: accountRouter,
    }
    return
}

func (o *AuthRouter) Setup(router *mux.Router) (ret error) {
    if ret = o.AccountRouter.Setup(router); ret != nil {
        return
    }
    return
}









