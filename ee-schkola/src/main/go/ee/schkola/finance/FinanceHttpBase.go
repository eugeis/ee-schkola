package finance

import (
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
    Router  *mux.Router
    PathPrefix  string
    QueryHandler  *ExpenseHttpQueryHandler
    CommandHandler  *ExpenseHttpCommandHandler
}

func (o *ExpenseRouter) Setup() (ret error) {
            
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
    Router  *mux.Router
    PathPrefix  string
    QueryHandler  *ExpensePurposeHttpQueryHandler
    CommandHandler  *ExpensePurposeHttpCommandHandler
}

func (o *ExpensePurposeRouter) Setup() (ret error) {
            
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
    Router  *mux.Router
    PathPrefix  string
    QueryHandler  *FeeHttpQueryHandler
    CommandHandler  *FeeHttpCommandHandler
}

func (o *FeeRouter) Setup() (ret error) {
            
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
    Router  *mux.Router
    PathPrefix  string
    QueryHandler  *FeeKindHttpQueryHandler
    CommandHandler  *FeeKindHttpCommandHandler
}

func (o *FeeKindRouter) Setup() (ret error) {
            
    return
    
}



