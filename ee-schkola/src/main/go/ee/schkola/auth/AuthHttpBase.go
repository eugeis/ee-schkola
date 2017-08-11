package auth

import (
    "github.com/eugeis/gee/net"
    "github.com/gorilla/mux"
    "net/http"
)
type AccountHttpQueryHandler struct {
}

func NewAccountHttpQueryHandler() (ret *AccountHttpQueryHandler) {
    ret = &AccountHttpQueryHandler{}
    return
}

func (o *AccountHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *AccountHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
}

func (o *AccountHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
}

func (o *AccountHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *AccountHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *AccountHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
}


type AccountHttpCommandHandler struct {
}

func NewAccountHttpCommandHandler() (ret *AccountHttpCommandHandler) {
    ret = &AccountHttpCommandHandler{}
    return
}

func (o *AccountHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
}

func (o *AccountHttpCommandHandler) Register(w http.ResponseWriter, r *http.Request)  {
}

func (o *AccountHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
}

func (o *AccountHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
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

func NewAccountRouter(PathPrefix string) (ret *AccountRouter) {
    
    ret.PathPrefix = ret.PathPrefix + "/" + "accounts"
    return
}

func (o *AccountRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindAccountAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindAccountById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountAccountById").HandlerFunc(o.QueryHandler.CountById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountAccountAll").HandlerFunc(o.QueryHandler.CountAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistAccountAll").HandlerFunc(o.QueryHandler.ExistAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistAccountById").HandlerFunc(o.QueryHandler.ExistById)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateAccount").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("RegisterAccount").HandlerFunc(o.CommandHandler.Register)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateAccount").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteAccount").HandlerFunc(o.CommandHandler.Delete)
    return
}


type AuthRouter struct {
    PathPrefix string
    AccountRouter *AccountRouter
    Router *mux.Router
}

func NewAuthRouter(PathPrefix string) (ret *AuthRouter) {
    
    ret.PathPrefix = ret.PathPrefix + "/" + "auth"
    return
}

func (o *AuthRouter) Setup(router *mux.Router) (ret error) {
    if ret = o.AccountRouter.Setup(router); ret != nil {
        return
    }
    return
}



