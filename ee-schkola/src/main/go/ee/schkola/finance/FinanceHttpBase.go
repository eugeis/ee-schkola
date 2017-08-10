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
}

func (o *ExpenseHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
}

func (o *ExpenseHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
}


type ExpenseRouter struct {
    PathPrefix string
    Router *mux.Router
    QueryHandler *ExpenseHttpQueryHandler
    CommandHandler *ExpenseHttpCommandHandler
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
}

func (o *ExpensePurposeHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
}

func (o *ExpensePurposeHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
}


type ExpensePurposeRouter struct {
    PathPrefix string
    Router *mux.Router
    QueryHandler *ExpensePurposeHttpQueryHandler
    CommandHandler *ExpensePurposeHttpCommandHandler
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
}

func (o *FeeHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
}

func (o *FeeHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
}


type FeeRouter struct {
    PathPrefix string
    Router *mux.Router
    QueryHandler *FeeHttpQueryHandler
    CommandHandler *FeeHttpCommandHandler
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
}

func (o *FeeKindHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
}

func (o *FeeKindHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
}


type FeeKindRouter struct {
    PathPrefix string
    Router *mux.Router
    QueryHandler *FeeKindHttpQueryHandler
    CommandHandler *FeeKindHttpCommandHandler
}

func (o *FeeKindRouter) Setup() (ret error) {
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateFeeKind").HandlerFunc(o.CommandHandler.Create)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateFeeKind").HandlerFunc(o.CommandHandler.Update)
    o.Router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteFeeKind").HandlerFunc(o.CommandHandler.Delete)
    return
}


type FinanceRouter struct {
    Router *mux.Router
    PathPrefix string
    ExpenseRouter *ExpenseRouter
    ExpensePurposeRouter *ExpensePurposeRouter
    FeeRouter *FeeRouter
    FeeKindRouter *FeeKindRouter
}

func (o *FinanceRouter) Setup() (ret error) {
    
    return
}



