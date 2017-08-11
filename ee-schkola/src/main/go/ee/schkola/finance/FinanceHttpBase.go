package finance

import (
    "github.com/eugeis/gee/net"
    "github.com/gorilla/mux"
    "net/http"
)
type ExpenseHttpQueryHandler struct {
}

func NewExpenseHttpQueryHandler() (ret *ExpenseHttpQueryHandler) {
    ret = &ExpenseHttpQueryHandler{}
    return
}

func (o *ExpenseHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *ExpenseHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
}

func (o *ExpenseHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
}

func (o *ExpenseHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *ExpenseHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *ExpenseHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
}


type ExpenseHttpCommandHandler struct {
}

func NewExpenseHttpCommandHandler() (ret *ExpenseHttpCommandHandler) {
    ret = &ExpenseHttpCommandHandler{}
    return
}

func (o *ExpenseHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
}

func (o *ExpenseHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
}

func (o *ExpenseHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
}


type ExpenseRouter struct {
    PathPrefix string
    QueryHandler *ExpenseHttpQueryHandler
    CommandHandler *ExpenseHttpCommandHandler
    Router *mux.Router
}

func NewExpenseRouter(PathPrefix string) (ret *ExpenseRouter) {
    
    ret.PathPrefix = ret.PathPrefix + "/" + "expenses"
    return
}

func (o *ExpenseRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindExpenseAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindExpenseById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountExpenseById").HandlerFunc(o.QueryHandler.CountById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountExpenseAll").HandlerFunc(o.QueryHandler.CountAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistExpenseAll").HandlerFunc(o.QueryHandler.ExistAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistExpenseById").HandlerFunc(o.QueryHandler.ExistById)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateExpense").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateExpense").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteExpense").HandlerFunc(o.CommandHandler.Delete)
    return
}


type ExpensePurposeHttpQueryHandler struct {
}

func NewExpensePurposeHttpQueryHandler() (ret *ExpensePurposeHttpQueryHandler) {
    ret = &ExpensePurposeHttpQueryHandler{}
    return
}

func (o *ExpensePurposeHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *ExpensePurposeHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
}

func (o *ExpensePurposeHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
}

func (o *ExpensePurposeHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *ExpensePurposeHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *ExpensePurposeHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
}


type ExpensePurposeHttpCommandHandler struct {
}

func NewExpensePurposeHttpCommandHandler() (ret *ExpensePurposeHttpCommandHandler) {
    ret = &ExpensePurposeHttpCommandHandler{}
    return
}

func (o *ExpensePurposeHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
}

func (o *ExpensePurposeHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
}

func (o *ExpensePurposeHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
}


type ExpensePurposeRouter struct {
    PathPrefix string
    QueryHandler *ExpensePurposeHttpQueryHandler
    CommandHandler *ExpensePurposeHttpCommandHandler
    Router *mux.Router
}

func NewExpensePurposeRouter(PathPrefix string) (ret *ExpensePurposeRouter) {
    
    ret.PathPrefix = ret.PathPrefix + "/" + "expensePurposes"
    return
}

func (o *ExpensePurposeRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindExpensePurposeAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindExpensePurposeById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountExpensePurposeById").HandlerFunc(o.QueryHandler.CountById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountExpensePurposeAll").HandlerFunc(o.QueryHandler.CountAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistExpensePurposeAll").HandlerFunc(o.QueryHandler.ExistAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistExpensePurposeById").HandlerFunc(o.QueryHandler.ExistById)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateExpensePurpose").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateExpensePurpose").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteExpensePurpose").HandlerFunc(o.CommandHandler.Delete)
    return
}


type FeeHttpQueryHandler struct {
}

func NewFeeHttpQueryHandler() (ret *FeeHttpQueryHandler) {
    ret = &FeeHttpQueryHandler{}
    return
}

func (o *FeeHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *FeeHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
}

func (o *FeeHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
}

func (o *FeeHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *FeeHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *FeeHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
}


type FeeHttpCommandHandler struct {
}

func NewFeeHttpCommandHandler() (ret *FeeHttpCommandHandler) {
    ret = &FeeHttpCommandHandler{}
    return
}

func (o *FeeHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
}

func (o *FeeHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
}

func (o *FeeHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
}


type FeeRouter struct {
    PathPrefix string
    QueryHandler *FeeHttpQueryHandler
    CommandHandler *FeeHttpCommandHandler
    Router *mux.Router
}

func NewFeeRouter(PathPrefix string) (ret *FeeRouter) {
    
    ret.PathPrefix = ret.PathPrefix + "/" + "fees"
    return
}

func (o *FeeRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindFeeAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindFeeById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountFeeById").HandlerFunc(o.QueryHandler.CountById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountFeeAll").HandlerFunc(o.QueryHandler.CountAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistFeeAll").HandlerFunc(o.QueryHandler.ExistAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistFeeById").HandlerFunc(o.QueryHandler.ExistById)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateFee").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateFee").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteFee").HandlerFunc(o.CommandHandler.Delete)
    return
}


type FeeKindHttpQueryHandler struct {
}

func NewFeeKindHttpQueryHandler() (ret *FeeKindHttpQueryHandler) {
    ret = &FeeKindHttpQueryHandler{}
    return
}

func (o *FeeKindHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *FeeKindHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
}

func (o *FeeKindHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
}

func (o *FeeKindHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *FeeKindHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *FeeKindHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
}


type FeeKindHttpCommandHandler struct {
}

func NewFeeKindHttpCommandHandler() (ret *FeeKindHttpCommandHandler) {
    ret = &FeeKindHttpCommandHandler{}
    return
}

func (o *FeeKindHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
}

func (o *FeeKindHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
}

func (o *FeeKindHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
}


type FeeKindRouter struct {
    PathPrefix string
    QueryHandler *FeeKindHttpQueryHandler
    CommandHandler *FeeKindHttpCommandHandler
    Router *mux.Router
}

func NewFeeKindRouter(PathPrefix string) (ret *FeeKindRouter) {
    
    ret.PathPrefix = ret.PathPrefix + "/" + "feeKinds"
    return
}

func (o *FeeKindRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindFeeKindAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindFeeKindById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountFeeKindById").HandlerFunc(o.QueryHandler.CountById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountFeeKindAll").HandlerFunc(o.QueryHandler.CountAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistFeeKindAll").HandlerFunc(o.QueryHandler.ExistAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistFeeKindById").HandlerFunc(o.QueryHandler.ExistById)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateFeeKind").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateFeeKind").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteFeeKind").HandlerFunc(o.CommandHandler.Delete)
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

func NewFinanceRouter(PathPrefix string) (ret *FinanceRouter) {
    
    ret.PathPrefix = ret.PathPrefix + "/" + "finance"
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



