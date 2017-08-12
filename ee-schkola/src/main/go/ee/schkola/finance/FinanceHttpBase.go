package finance

import (
    "fmt"
    "github.com/eugeis/gee/net"
    "github.com/gorilla/mux"
    "html"
    "net/http"
)
type ExpenseHttpQueryHandler struct {
}

func NewExpenseHttpQueryHandler() (ret *ExpenseHttpQueryHandler) {
    ret = &ExpenseHttpQueryHandler{}
    return
}

func (o *ExpenseHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindAllExpense", html.EscapeString(r.URL.Path))
}

func (o *ExpenseHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindByExpenseId", html.EscapeString(r.URL.Path))
}

func (o *ExpenseHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountAllExpense", html.EscapeString(r.URL.Path))
}

func (o *ExpenseHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountByExpenseId", html.EscapeString(r.URL.Path))
}

func (o *ExpenseHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistAllExpense", html.EscapeString(r.URL.Path))
}

func (o *ExpenseHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistByExpenseId", html.EscapeString(r.URL.Path))
}


type ExpenseHttpCommandHandler struct {
}

func NewExpenseHttpCommandHandler() (ret *ExpenseHttpCommandHandler) {
    ret = &ExpenseHttpCommandHandler{}
    return
}

func (o *ExpenseHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExpenseCreate", html.EscapeString(r.URL.Path))
}

func (o *ExpenseHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExpenseUpdate", html.EscapeString(r.URL.Path))
}

func (o *ExpenseHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExpenseDelete", html.EscapeString(r.URL.Path))
}


type ExpenseRouter struct {
    PathPrefix string
    QueryHandler *ExpenseHttpQueryHandler
    CommandHandler *ExpenseHttpCommandHandler
    Router *mux.Router
}

func NewExpenseRouter(pathPrefix string) (ret *ExpenseRouter) {
    pathPrefix = pathPrefix + "/" + "expenses"
    ret = &ExpenseRouter{
        PathPrefix: pathPrefix,
        QueryHandler: NewExpenseHttpQueryHandler(),
        CommandHandler: NewExpenseHttpCommandHandler(),
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
}

func NewExpensePurposeHttpQueryHandler() (ret *ExpensePurposeHttpQueryHandler) {
    ret = &ExpensePurposeHttpQueryHandler{}
    return
}

func (o *ExpensePurposeHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindAllExpensePurpose", html.EscapeString(r.URL.Path))
}

func (o *ExpensePurposeHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindByExpensePurposeId", html.EscapeString(r.URL.Path))
}

func (o *ExpensePurposeHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountAllExpensePurpose", html.EscapeString(r.URL.Path))
}

func (o *ExpensePurposeHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountByExpensePurposeId", html.EscapeString(r.URL.Path))
}

func (o *ExpensePurposeHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistAllExpensePurpose", html.EscapeString(r.URL.Path))
}

func (o *ExpensePurposeHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistByExpensePurposeId", html.EscapeString(r.URL.Path))
}


type ExpensePurposeHttpCommandHandler struct {
}

func NewExpensePurposeHttpCommandHandler() (ret *ExpensePurposeHttpCommandHandler) {
    ret = &ExpensePurposeHttpCommandHandler{}
    return
}

func (o *ExpensePurposeHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExpensePurposeCreate", html.EscapeString(r.URL.Path))
}

func (o *ExpensePurposeHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExpensePurposeUpdate", html.EscapeString(r.URL.Path))
}

func (o *ExpensePurposeHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExpensePurposeDelete", html.EscapeString(r.URL.Path))
}


type ExpensePurposeRouter struct {
    PathPrefix string
    QueryHandler *ExpensePurposeHttpQueryHandler
    CommandHandler *ExpensePurposeHttpCommandHandler
    Router *mux.Router
}

func NewExpensePurposeRouter(pathPrefix string) (ret *ExpensePurposeRouter) {
    pathPrefix = pathPrefix + "/" + "expensePurposes"
    ret = &ExpensePurposeRouter{
        PathPrefix: pathPrefix,
        QueryHandler: NewExpensePurposeHttpQueryHandler(),
        CommandHandler: NewExpensePurposeHttpCommandHandler(),
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
}

func NewFeeHttpQueryHandler() (ret *FeeHttpQueryHandler) {
    ret = &FeeHttpQueryHandler{}
    return
}

func (o *FeeHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindAllFee", html.EscapeString(r.URL.Path))
}

func (o *FeeHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindByFeeId", html.EscapeString(r.URL.Path))
}

func (o *FeeHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountAllFee", html.EscapeString(r.URL.Path))
}

func (o *FeeHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountByFeeId", html.EscapeString(r.URL.Path))
}

func (o *FeeHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistAllFee", html.EscapeString(r.URL.Path))
}

func (o *FeeHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistByFeeId", html.EscapeString(r.URL.Path))
}


type FeeHttpCommandHandler struct {
}

func NewFeeHttpCommandHandler() (ret *FeeHttpCommandHandler) {
    ret = &FeeHttpCommandHandler{}
    return
}

func (o *FeeHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FeeCreate", html.EscapeString(r.URL.Path))
}

func (o *FeeHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FeeUpdate", html.EscapeString(r.URL.Path))
}

func (o *FeeHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FeeDelete", html.EscapeString(r.URL.Path))
}


type FeeRouter struct {
    PathPrefix string
    QueryHandler *FeeHttpQueryHandler
    CommandHandler *FeeHttpCommandHandler
    Router *mux.Router
}

func NewFeeRouter(pathPrefix string) (ret *FeeRouter) {
    pathPrefix = pathPrefix + "/" + "fees"
    ret = &FeeRouter{
        PathPrefix: pathPrefix,
        QueryHandler: NewFeeHttpQueryHandler(),
        CommandHandler: NewFeeHttpCommandHandler(),
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
}

func NewFeeKindHttpQueryHandler() (ret *FeeKindHttpQueryHandler) {
    ret = &FeeKindHttpQueryHandler{}
    return
}

func (o *FeeKindHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindAllFeeKind", html.EscapeString(r.URL.Path))
}

func (o *FeeKindHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindByFeeKindId", html.EscapeString(r.URL.Path))
}

func (o *FeeKindHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountAllFeeKind", html.EscapeString(r.URL.Path))
}

func (o *FeeKindHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountByFeeKindId", html.EscapeString(r.URL.Path))
}

func (o *FeeKindHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistAllFeeKind", html.EscapeString(r.URL.Path))
}

func (o *FeeKindHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistByFeeKindId", html.EscapeString(r.URL.Path))
}


type FeeKindHttpCommandHandler struct {
}

func NewFeeKindHttpCommandHandler() (ret *FeeKindHttpCommandHandler) {
    ret = &FeeKindHttpCommandHandler{}
    return
}

func (o *FeeKindHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FeeKindCreate", html.EscapeString(r.URL.Path))
}

func (o *FeeKindHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FeeKindUpdate", html.EscapeString(r.URL.Path))
}

func (o *FeeKindHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FeeKindDelete", html.EscapeString(r.URL.Path))
}


type FeeKindRouter struct {
    PathPrefix string
    QueryHandler *FeeKindHttpQueryHandler
    CommandHandler *FeeKindHttpCommandHandler
    Router *mux.Router
}

func NewFeeKindRouter(pathPrefix string) (ret *FeeKindRouter) {
    pathPrefix = pathPrefix + "/" + "feeKinds"
    ret = &FeeKindRouter{
        PathPrefix: pathPrefix,
        QueryHandler: NewFeeKindHttpQueryHandler(),
        CommandHandler: NewFeeKindHttpCommandHandler(),
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

func NewFinanceRouter(pathPrefix string) (ret *FinanceRouter) {
    pathPrefix = pathPrefix + "/" + "finance"
    ret = &FinanceRouter{
        PathPrefix: pathPrefix,
        ExpenseRouter: NewExpenseRouter(pathPrefix),
        ExpensePurposeRouter: NewExpensePurposeRouter(pathPrefix),
        FeeRouter: NewFeeRouter(pathPrefix),
        FeeKindRouter: NewFeeKindRouter(pathPrefix),
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



