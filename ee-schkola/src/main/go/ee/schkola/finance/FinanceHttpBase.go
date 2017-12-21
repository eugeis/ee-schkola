package finance

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

func NewExpenseHttpQueryHandler(queryRepository *QueryRepository) (ret *HttpQueryHandler) {
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

func NewExpenseHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandHandler) (ret *HttpCommandHandler) {
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


type Router struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    QueryHandler *HttpQueryHandler `json:"queryHandler" eh:"optional"`
    CommandHandler *HttpCommandHandler `json:"commandHandler" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func New@@EMPTY@@(pathPrefix string, context context.Context, commandBus eventhorizon.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *Router) {
    pathPrefix = pathPrefix + "/" + "expenses"
    entityFactory := func() eventhorizon.Entity { return NewExpense() }
    repo := readRepos(string(ExpenseAggregateType), entityFactory)
    queryRepository := NewExpenseQueryRepository(repo, context)
    queryHandler := NewExpenseHttpQueryHandler(queryRepository)
    commandHandler := NewExpenseHttpCommandHandler(context, commandBus)
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
        Name("Create").HandlerFunc(o.CommandHandler.Create)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("Update").HandlerFunc(o.CommandHandler.Update)
    router.Methods(http.MethodDelete).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("Delete").HandlerFunc(o.CommandHandler.Delete)
    return
}


type HttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *QueryRepository `json:"queryRepository" eh:"optional"`
}

func NewExpensePurposeHttpQueryHandler(queryRepository *QueryRepository) (ret *HttpQueryHandler) {
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

func NewExpensePurposeHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandHandler) (ret *HttpCommandHandler) {
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


type Router struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    QueryHandler *HttpQueryHandler `json:"queryHandler" eh:"optional"`
    CommandHandler *HttpCommandHandler `json:"commandHandler" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func New@@EMPTY@@(pathPrefix string, context context.Context, commandBus eventhorizon.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *Router) {
    pathPrefix = pathPrefix + "/" + "expensePurposes"
    entityFactory := func() eventhorizon.Entity { return NewExpensePurpose() }
    repo := readRepos(string(ExpensePurposeAggregateType), entityFactory)
    queryRepository := NewExpensePurposeQueryRepository(repo, context)
    queryHandler := NewExpensePurposeHttpQueryHandler(queryRepository)
    commandHandler := NewExpensePurposeHttpCommandHandler(context, commandBus)
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
        Name("Create").HandlerFunc(o.CommandHandler.Create)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("Update").HandlerFunc(o.CommandHandler.Update)
    router.Methods(http.MethodDelete).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("Delete").HandlerFunc(o.CommandHandler.Delete)
    return
}


type HttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *QueryRepository `json:"queryRepository" eh:"optional"`
}

func NewFeeHttpQueryHandler(queryRepository *QueryRepository) (ret *HttpQueryHandler) {
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

func NewFeeHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandHandler) (ret *HttpCommandHandler) {
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


type Router struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    QueryHandler *HttpQueryHandler `json:"queryHandler" eh:"optional"`
    CommandHandler *HttpCommandHandler `json:"commandHandler" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func New@@EMPTY@@(pathPrefix string, context context.Context, commandBus eventhorizon.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *Router) {
    pathPrefix = pathPrefix + "/" + "fees"
    entityFactory := func() eventhorizon.Entity { return NewFee() }
    repo := readRepos(string(FeeAggregateType), entityFactory)
    queryRepository := NewFeeQueryRepository(repo, context)
    queryHandler := NewFeeHttpQueryHandler(queryRepository)
    commandHandler := NewFeeHttpCommandHandler(context, commandBus)
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
        Name("Create").HandlerFunc(o.CommandHandler.Create)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("Update").HandlerFunc(o.CommandHandler.Update)
    router.Methods(http.MethodDelete).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("Delete").HandlerFunc(o.CommandHandler.Delete)
    return
}


type HttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *QueryRepository `json:"queryRepository" eh:"optional"`
}

func NewFeeKindHttpQueryHandler(queryRepository *QueryRepository) (ret *HttpQueryHandler) {
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

func NewFeeKindHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandHandler) (ret *HttpCommandHandler) {
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


type Router struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    QueryHandler *HttpQueryHandler `json:"queryHandler" eh:"optional"`
    CommandHandler *HttpCommandHandler `json:"commandHandler" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func New@@EMPTY@@(pathPrefix string, context context.Context, commandBus eventhorizon.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *Router) {
    pathPrefix = pathPrefix + "/" + "feeKinds"
    entityFactory := func() eventhorizon.Entity { return NewFeeKind() }
    repo := readRepos(string(FeeKindAggregateType), entityFactory)
    queryRepository := NewFeeKindQueryRepository(repo, context)
    queryHandler := NewFeeKindHttpQueryHandler(queryRepository)
    commandHandler := NewFeeKindHttpCommandHandler(context, commandBus)
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
        Name("Create").HandlerFunc(o.CommandHandler.Create)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("Update").HandlerFunc(o.CommandHandler.Update)
    router.Methods(http.MethodDelete).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("Delete").HandlerFunc(o.CommandHandler.Delete)
    return
}


type FinanceRouter struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    ExpenseRouter *Router `json:"expenseRouter" eh:"optional"`
    ExpensePurposeRouter *Router `json:"expensePurposeRouter" eh:"optional"`
    FeeRouter *Router `json:"feeRouter" eh:"optional"`
    FeeKindRouter *Router `json:"feeKindRouter" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func New@@EMPTY@@(pathPrefix string, context context.Context, commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *FinanceRouter) {
    pathPrefix = pathPrefix + "/" + "finance"
    expenseRouter := New@@EMPTY@@(pathPrefix, context, commandBus, readRepos)
    expensePurposeRouter := New@@EMPTY@@(pathPrefix, context, commandBus, readRepos)
    feeRouter := New@@EMPTY@@(pathPrefix, context, commandBus, readRepos)
    feeKindRouter := New@@EMPTY@@(pathPrefix, context, commandBus, readRepos)
    ret = &FinanceRouter{
        PathPrefix: pathPrefix,
        ExpenseRouter: expenseRouter,
        ExpensePurposeRouter: expensePurposeRouter,
        FeeRouter: feeRouter,
        FeeKindRouter: feeKindRouter,
    }
    return
}

func (o *FinanceRouter) Setup(router *mux.Router) (err error) {
    if err = o.ExpenseRouter.Setup(router); err != nil {
        return
    }
    if err = o.ExpensePurposeRouter.Setup(router); err != nil {
        return
    }
    if err = o.FeeRouter.Setup(router); err != nil {
        return
    }
    if err = o.FeeKindRouter.Setup(router); err != nil {
        return
    }
    return
}









