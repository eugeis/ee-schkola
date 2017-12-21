package auth

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

func NewAccountHttpQueryHandler(queryRepository *QueryRepository) (ret *HttpQueryHandler) {
    httpQueryHandler := &eh.HttpQueryHandler{}
    ret = &HttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *HttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAll", w, r)
}

func (o *HttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindById", w, r)
}

func (o *HttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAll", w, r)
}

func (o *HttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountById", w, r)
}

func (o *HttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAll", w, r)
}

func (o *HttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistById", w, r)
}


type HttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewAccountHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandHandler) (ret *HttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &HttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *HttpCommandHandler) Create(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&Create{Id: id}, w, r)
}

func (o *HttpCommandHandler) Disable(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&Disable{Id: id}, w, r)
}

func (o *HttpCommandHandler) Enable(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&Enable{Id: id}, w, r)
}

func (o *HttpCommandHandler) Update(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&Update{Id: id}, w, r)
}

func (o *HttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&Delete{Id: id}, w, r)
}

func (o *HttpCommandHandler) SendEnabledConfirmation(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&SendEnabledConfirmation{Id: id}, w, r)
}

func (o *HttpCommandHandler) SendDisabledConfirmation(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&SendDisabledConfirmation{Id: id}, w, r)
}

func (o *HttpCommandHandler) Login(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&Login{Id: id}, w, r)
}

func (o *HttpCommandHandler) SendCreatedConfirmation(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&SendCreatedConfirmation{Id: id}, w, r)
}


type Router struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    QueryHandler *HttpQueryHandler `json:"queryHandler" eh:"optional"`
    CommandHandler *HttpCommandHandler `json:"commandHandler" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func New@@EMPTY@@(pathPrefix string, context context.Context, commandBus eventhorizon.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *Router) {
    pathPrefix = pathPrefix + "/" + "accounts"
    entityFactory := func() eventhorizon.Entity { return NewAccount() }
    repo := readRepos(string(AccountAggregateType), entityFactory)
    queryRepository := NewAccountQueryRepository(repo, context)
    queryHandler := NewAccountHttpQueryHandler(queryRepository)
    commandHandler := NewAccountHttpCommandHandler(context, commandBus)
    ret = &Router{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *Router) Setup(router *mux.Router) (err error) {
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("CountAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("ExistAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("FindAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
        Queries(net.Command, "login").
        Name("Login").HandlerFunc(o.CommandHandler.Login)
    router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
        Queries(net.Command, "sendCreatedConfirmation").
        Name("SendCreatedConfirmation").HandlerFunc(o.CommandHandler.SendCreatedConfirmation)
    router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
        Queries(net.Command, "sendEnabledConfirmation").
        Name("SendEnabledConfirmation").HandlerFunc(o.CommandHandler.SendEnabledConfirmation)
    router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
        Queries(net.Command, "sendDisabledConfirmation").
        Name("SendDisabledConfirmation").HandlerFunc(o.CommandHandler.SendDisabledConfirmation)
    router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("Create").HandlerFunc(o.CommandHandler.Create)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Queries(net.Command, "enable").
        Name("Enable").HandlerFunc(o.CommandHandler.Enable)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Queries(net.Command, "disable").
        Name("Disable").HandlerFunc(o.CommandHandler.Disable)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("Update").HandlerFunc(o.CommandHandler.Update)
    router.Methods(http.MethodDelete).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("Delete").HandlerFunc(o.CommandHandler.Delete)
    return
}


type AuthRouter struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    AccountRouter *Router `json:"accountRouter" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func New@@EMPTY@@(pathPrefix string, context context.Context, commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *AuthRouter) {
    pathPrefix = pathPrefix + "/" + "auth"
    accountRouter := New@@EMPTY@@(pathPrefix, context, commandBus, readRepos)
    ret = &AuthRouter{
        PathPrefix: pathPrefix,
        AccountRouter: accountRouter,
    }
    return
}

func (o *AuthRouter) Setup(router *mux.Router) (err error) {
    if err = o.AccountRouter.Setup(router); err != nil {
        return
    }
    return
}









