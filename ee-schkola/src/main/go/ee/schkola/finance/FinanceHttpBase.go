package finance

import (
    "context"
    "github.com/eugeis/gee/eh"
    "github.com/eugeis/gee/net"
    "github.com/gorilla/mux"
    "github.com/looplab/eventhorizon"
    "net/http"
)
type ExpenseHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *ExpenseQueryRepository
}

func NewExpenseHttpQueryHandler(queryRepository *ExpenseQueryRepository) (ret *ExpenseHttpQueryHandler) {
    httpQueryHandler := &eh.HttpQueryHandler{}
    ret = &ExpenseHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *ExpenseHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllExpense", w, r)
}

func (o *ExpenseHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindByExpenseId", w, r)
}

func (o *ExpenseHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllExpense", w, r)
}

func (o *ExpenseHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountByExpenseId", w, r)
}

func (o *ExpenseHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllExpense", w, r)
}

func (o *ExpenseHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistByExpenseId", w, r)
}


type ExpenseHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewExpenseHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *ExpenseHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &ExpenseHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *ExpenseHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateExpense{Id: id}, w, r)
}

func (o *ExpenseHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateExpense{Id: id}, w, r)
}

func (o *ExpenseHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteExpense{Id: id}, w, r)
}


type ExpenseRouter struct {
    PathPrefix string
    QueryHandler *ExpenseHttpQueryHandler
    CommandHandler *ExpenseHttpCommandHandler
    Router *mux.Router
}

func NewExpenseRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus, 
                readRepos func (string) eventhorizon.ReadWriteRepo) (ret *ExpenseRouter) {
    pathPrefix = pathPrefix + "/" + "expenses"
    repo := readRepos(string(ExpenseAggregateType))
    queryRepository := NewExpenseQueryRepository(repo, context)
    queryHandler := NewExpenseHttpQueryHandler(queryRepository)
    commandHandler := NewExpenseHttpCommandHandler(context, commandBus)
    ret = &ExpenseRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *ExpenseRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountExpenseById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("CountExpenseAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistExpenseById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("ExistExpenseAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindExpenseById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("FindExpenseAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateExpense").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateExpense").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteExpense").HandlerFunc(o.CommandHandler.Delete)
    return
}


type ExpensePurposeHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *ExpensePurposeQueryRepository
}

func NewExpensePurposeHttpQueryHandler(queryRepository *ExpensePurposeQueryRepository) (ret *ExpensePurposeHttpQueryHandler) {
    httpQueryHandler := &eh.HttpQueryHandler{}
    ret = &ExpensePurposeHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *ExpensePurposeHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllExpensePurpose", w, r)
}

func (o *ExpensePurposeHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindByExpensePurposeId", w, r)
}

func (o *ExpensePurposeHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllExpensePurpose", w, r)
}

func (o *ExpensePurposeHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountByExpensePurposeId", w, r)
}

func (o *ExpensePurposeHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllExpensePurpose", w, r)
}

func (o *ExpensePurposeHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistByExpensePurposeId", w, r)
}


type ExpensePurposeHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewExpensePurposeHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *ExpensePurposeHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &ExpensePurposeHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *ExpensePurposeHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateExpensePurpose{Id: id}, w, r)
}

func (o *ExpensePurposeHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateExpensePurpose{Id: id}, w, r)
}

func (o *ExpensePurposeHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteExpensePurpose{Id: id}, w, r)
}


type ExpensePurposeRouter struct {
    PathPrefix string
    QueryHandler *ExpensePurposeHttpQueryHandler
    CommandHandler *ExpensePurposeHttpCommandHandler
    Router *mux.Router
}

func NewExpensePurposeRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus, 
                readRepos func (string) eventhorizon.ReadWriteRepo) (ret *ExpensePurposeRouter) {
    pathPrefix = pathPrefix + "/" + "expensePurposes"
    repo := readRepos(string(ExpensePurposeAggregateType))
    queryRepository := NewExpensePurposeQueryRepository(repo, context)
    queryHandler := NewExpensePurposeHttpQueryHandler(queryRepository)
    commandHandler := NewExpensePurposeHttpCommandHandler(context, commandBus)
    ret = &ExpensePurposeRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *ExpensePurposeRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountExpensePurposeById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("CountExpensePurposeAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistExpensePurposeById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("ExistExpensePurposeAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindExpensePurposeById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("FindExpensePurposeAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateExpensePurpose").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateExpensePurpose").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteExpensePurpose").HandlerFunc(o.CommandHandler.Delete)
    return
}


type FeeHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *FeeQueryRepository
}

func NewFeeHttpQueryHandler(queryRepository *FeeQueryRepository) (ret *FeeHttpQueryHandler) {
    httpQueryHandler := &eh.HttpQueryHandler{}
    ret = &FeeHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *FeeHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllFee", w, r)
}

func (o *FeeHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindByFeeId", w, r)
}

func (o *FeeHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllFee", w, r)
}

func (o *FeeHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountByFeeId", w, r)
}

func (o *FeeHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllFee", w, r)
}

func (o *FeeHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistByFeeId", w, r)
}


type FeeHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewFeeHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *FeeHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &FeeHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *FeeHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateFee{Id: id}, w, r)
}

func (o *FeeHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateFee{Id: id}, w, r)
}

func (o *FeeHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteFee{Id: id}, w, r)
}


type FeeRouter struct {
    PathPrefix string
    QueryHandler *FeeHttpQueryHandler
    CommandHandler *FeeHttpCommandHandler
    Router *mux.Router
}

func NewFeeRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus, 
                readRepos func (string) eventhorizon.ReadWriteRepo) (ret *FeeRouter) {
    pathPrefix = pathPrefix + "/" + "fees"
    repo := readRepos(string(FeeAggregateType))
    queryRepository := NewFeeQueryRepository(repo, context)
    queryHandler := NewFeeHttpQueryHandler(queryRepository)
    commandHandler := NewFeeHttpCommandHandler(context, commandBus)
    ret = &FeeRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *FeeRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountFeeById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("CountFeeAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistFeeById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("ExistFeeAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindFeeById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("FindFeeAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateFee").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateFee").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteFee").HandlerFunc(o.CommandHandler.Delete)
    return
}


type FeeKindHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *FeeKindQueryRepository
}

func NewFeeKindHttpQueryHandler(queryRepository *FeeKindQueryRepository) (ret *FeeKindHttpQueryHandler) {
    httpQueryHandler := &eh.HttpQueryHandler{}
    ret = &FeeKindHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *FeeKindHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllFeeKind", w, r)
}

func (o *FeeKindHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindByFeeKindId", w, r)
}

func (o *FeeKindHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllFeeKind", w, r)
}

func (o *FeeKindHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountByFeeKindId", w, r)
}

func (o *FeeKindHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllFeeKind", w, r)
}

func (o *FeeKindHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistByFeeKindId", w, r)
}


type FeeKindHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewFeeKindHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *FeeKindHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &FeeKindHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *FeeKindHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateFeeKind{Id: id}, w, r)
}

func (o *FeeKindHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateFeeKind{Id: id}, w, r)
}

func (o *FeeKindHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteFeeKind{Id: id}, w, r)
}


type FeeKindRouter struct {
    PathPrefix string
    QueryHandler *FeeKindHttpQueryHandler
    CommandHandler *FeeKindHttpCommandHandler
    Router *mux.Router
}

func NewFeeKindRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus, 
                readRepos func (string) eventhorizon.ReadWriteRepo) (ret *FeeKindRouter) {
    pathPrefix = pathPrefix + "/" + "feeKinds"
    repo := readRepos(string(FeeKindAggregateType))
    queryRepository := NewFeeKindQueryRepository(repo, context)
    queryHandler := NewFeeKindHttpQueryHandler(queryRepository)
    commandHandler := NewFeeKindHttpCommandHandler(context, commandBus)
    ret = &FeeKindRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *FeeKindRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountFeeKindById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("CountFeeKindAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistFeeKindById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("ExistFeeKindAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindFeeKindById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("FindFeeKindAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateFeeKind").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateFeeKind").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteFeeKind").HandlerFunc(o.CommandHandler.Delete)
    return
}


type FinanceRouter struct {
    PathPrefix string
    ExpenseRouter *ExpenseRouter
    ExpensePurposeRouter *ExpensePurposeRouter
    FeeRouter *FeeRouter
    FeeKindRouter *FeeKindRouter
    Router *mux.Router
}

func NewFinanceRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus, 
                readRepos func (string) eventhorizon.ReadWriteRepo) (ret *FinanceRouter) {
    pathPrefix = pathPrefix + "/" + "finance"
    expenseRouter := NewExpenseRouter(pathPrefix, context, commandBus, readRepos)
    expensePurposeRouter := NewExpensePurposeRouter(pathPrefix, context, commandBus, readRepos)
    feeRouter := NewFeeRouter(pathPrefix, context, commandBus, readRepos)
    feeKindRouter := NewFeeKindRouter(pathPrefix, context, commandBus, readRepos)
    ret = &FinanceRouter{
        PathPrefix: pathPrefix,
        ExpenseRouter: expenseRouter,
        ExpensePurposeRouter: expensePurposeRouter,
        FeeRouter: feeRouter,
        FeeKindRouter: feeKindRouter,
    }
    return
}

func (o *FinanceRouter) Setup(router *mux.Router) (ret error) {
    if ret = o.ExpenseRouter.Setup(router); ret != nil {
        return
    }
    if ret = o.ExpensePurposeRouter.Setup(router); ret != nil {
        return
    }
    if ret = o.FeeRouter.Setup(router); ret != nil {
        return
    }
    if ret = o.FeeKindRouter.Setup(router); ret != nil {
        return
    }
    return
}









