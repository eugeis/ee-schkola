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



type ExpenseHttpCommandHandler struct {
}

func NewExpenseHttpCommandHandler() (ret *ExpenseHttpCommandHandler) {
    ret = &ExpenseHttpCommandHandler{}
    return
}

func (o *ExpenseHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *ExpenseHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *ExpenseHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}



type ExpenseRouter struct {
    PathPrefix  string
    Router  *mux.Router
    QueryHandler  *ExpenseHttpQueryHandler
    CommandHandler  *ExpenseHttpCommandHandler
}

func NewExpenseRouter(Router *mux.Router) (ret *ExpenseRouter) {
    ret = &ExpenseRouter{
        PathPrefix :"Expenses",
        Router :Router,
        QueryHandler :NewExpenseHttpQueryHandler(),
        CommandHandler :NewExpenseHttpCommandHandler(),
    }
    return
}

func (o *ExpenseRouter) Setup() (ret error) {
            
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateExpense").HandlerFunc(o.CommandHandler.Create)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateExpense").HandlerFunc(o.CommandHandler.Update)
    o.Router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteExpense").HandlerFunc(o.CommandHandler.Delete)
    return
    
}



type ExpensePurposeHttpQueryHandler struct {
}

func NewExpensePurposeHttpQueryHandler() (ret *ExpensePurposeHttpQueryHandler) {
    ret = &ExpensePurposeHttpQueryHandler{}
    return
}



type ExpensePurposeHttpCommandHandler struct {
}

func NewExpensePurposeHttpCommandHandler() (ret *ExpensePurposeHttpCommandHandler) {
    ret = &ExpensePurposeHttpCommandHandler{}
    return
}

func (o *ExpensePurposeHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *ExpensePurposeHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *ExpensePurposeHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}



type ExpensePurposeRouter struct {
    PathPrefix  string
    Router  *mux.Router
    QueryHandler  *ExpensePurposeHttpQueryHandler
    CommandHandler  *ExpensePurposeHttpCommandHandler
}

func NewExpensePurposeRouter(Router *mux.Router) (ret *ExpensePurposeRouter) {
    ret = &ExpensePurposeRouter{
        PathPrefix :"ExpensePurposes",
        Router :Router,
        QueryHandler :NewExpensePurposeHttpQueryHandler(),
        CommandHandler :NewExpensePurposeHttpCommandHandler(),
    }
    return
}

func (o *ExpensePurposeRouter) Setup() (ret error) {
            
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateExpensePurpose").HandlerFunc(o.CommandHandler.Create)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateExpensePurpose").HandlerFunc(o.CommandHandler.Update)
    o.Router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteExpensePurpose").HandlerFunc(o.CommandHandler.Delete)
    return
    
}



type FeeHttpQueryHandler struct {
}

func NewFeeHttpQueryHandler() (ret *FeeHttpQueryHandler) {
    ret = &FeeHttpQueryHandler{}
    return
}



type FeeHttpCommandHandler struct {
}

func NewFeeHttpCommandHandler() (ret *FeeHttpCommandHandler) {
    ret = &FeeHttpCommandHandler{}
    return
}

func (o *FeeHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *FeeHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *FeeHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}



type FeeRouter struct {
    PathPrefix  string
    Router  *mux.Router
    QueryHandler  *FeeHttpQueryHandler
    CommandHandler  *FeeHttpCommandHandler
}

func NewFeeRouter(Router *mux.Router) (ret *FeeRouter) {
    ret = &FeeRouter{
        PathPrefix :"Fees",
        Router :Router,
        QueryHandler :NewFeeHttpQueryHandler(),
        CommandHandler :NewFeeHttpCommandHandler(),
    }
    return
}

func (o *FeeRouter) Setup() (ret error) {
            
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateFee").HandlerFunc(o.CommandHandler.Create)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateFee").HandlerFunc(o.CommandHandler.Update)
    o.Router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteFee").HandlerFunc(o.CommandHandler.Delete)
    return
    
}



type FeeKindHttpQueryHandler struct {
}

func NewFeeKindHttpQueryHandler() (ret *FeeKindHttpQueryHandler) {
    ret = &FeeKindHttpQueryHandler{}
    return
}



type FeeKindHttpCommandHandler struct {
}

func NewFeeKindHttpCommandHandler() (ret *FeeKindHttpCommandHandler) {
    ret = &FeeKindHttpCommandHandler{}
    return
}

func (o *FeeKindHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *FeeKindHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}

func (o *FeeKindHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
            
    return
    
}



type FeeKindRouter struct {
    PathPrefix  string
    Router  *mux.Router
    QueryHandler  *FeeKindHttpQueryHandler
    CommandHandler  *FeeKindHttpCommandHandler
}

func NewFeeKindRouter(Router *mux.Router) (ret *FeeKindRouter) {
    ret = &FeeKindRouter{
        PathPrefix :"FeeKinds",
        Router :Router,
        QueryHandler :NewFeeKindHttpQueryHandler(),
        CommandHandler :NewFeeKindHttpCommandHandler(),
    }
    return
}

func (o *FeeKindRouter) Setup() (ret error) {
            
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateFeeKind").HandlerFunc(o.CommandHandler.Create)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateFeeKind").HandlerFunc(o.CommandHandler.Update)
    o.Router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteFeeKind").HandlerFunc(o.CommandHandler.Delete)
    return
    
}



type FinanceRouter struct {
    Router  *mux.Router
    PathPrefix  string
    ExpenseRouter  *ExpenseRouter
    ExpensePurposeRouter  *ExpensePurposeRouter
    FeeRouter  *FeeRouter
    FeeKindRouter  *FeeKindRouter
}

func (o *FinanceRouter) Setup() (ret error) {
            
    return
}



