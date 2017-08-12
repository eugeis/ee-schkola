package auth

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/eugeis/gee/net"
    "github.com/gorilla/mux"
    "github.com/looplab/eventhorizon"
    "html"
    "net/http"
)
type AccountHttpQueryHandler struct {
}

func NewAccountHttpQueryHandler() (ret *AccountHttpQueryHandler) {
    ret = &AccountHttpQueryHandler{}
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
    context context.Context
    commandBus eventhorizon.CommandBus
}

func NewAccountHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *AccountHttpCommandHandler) {
    ret = &AccountHttpCommandHandler{
        context: context,
        commandBus: commandBus,
    }
    return
}

func (o *AccountHttpCommandHandler) Register(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := vars["id"]
    
    decoder := json.NewDecoder(r.Body)
    command := &RegisterAccount{}
    if err := decoder.Decode(command); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Can't decode body to command %v because of %v", command, err)
    }
    defer r.Body.Close()

    if err := o.commandBus.HandleCommand(o.context, command); err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		fmt.Fprintf(w, "Can't execute command %v because of %v", command, err)
		return
	}
    fmt.Fprintf(w, "id=%v, %q from AccountRegister", id, html.EscapeString(r.URL.Path))
}

func (o *AccountHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := vars["id"]
    
    decoder := json.NewDecoder(r.Body)
    command := &CreateAccount{}
    if err := decoder.Decode(command); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Can't decode body to command %v because of %v", command, err)
    }
    defer r.Body.Close()

    if err := o.commandBus.HandleCommand(o.context, command); err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		fmt.Fprintf(w, "Can't execute command %v because of %v", command, err)
		return
	}
    fmt.Fprintf(w, "id=%v, %q from AccountCreate", id, html.EscapeString(r.URL.Path))
}

func (o *AccountHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := vars["id"]
    
    decoder := json.NewDecoder(r.Body)
    command := &UpdateAccount{}
    if err := decoder.Decode(command); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Can't decode body to command %v because of %v", command, err)
    }
    defer r.Body.Close()

    if err := o.commandBus.HandleCommand(o.context, command); err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		fmt.Fprintf(w, "Can't execute command %v because of %v", command, err)
		return
	}
    fmt.Fprintf(w, "id=%v, %q from AccountUpdate", id, html.EscapeString(r.URL.Path))
}

func (o *AccountHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := vars["id"]
    
    decoder := json.NewDecoder(r.Body)
    command := &DeleteAccount{}
    if err := decoder.Decode(command); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Can't decode body to command %v because of %v", command, err)
    }
    defer r.Body.Close()

    if err := o.commandBus.HandleCommand(o.context, command); err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		fmt.Fprintf(w, "Can't execute command %v because of %v", command, err)
		return
	}
    fmt.Fprintf(w, "id=%v, %q from AccountDelete", id, html.EscapeString(r.URL.Path))
}

func (o *AccountHttpCommandHandler) Enable(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := vars["id"]
    
    fmt.Fprintf(w, "id=%v, %q from AccountEnable", id, html.EscapeString(r.URL.Path))
}

func (o *AccountHttpCommandHandler) Disable(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := vars["id"]
    
    fmt.Fprintf(w, "id=%v, %q from AccountDisable", id, html.EscapeString(r.URL.Path))
}


type AccountRouter struct {
    PathPrefix string
    QueryHandler *AccountHttpQueryHandler
    CommandHandler *AccountHttpCommandHandler
    Router *mux.Router
}

func NewAccountRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus) (ret *AccountRouter) {
    pathPrefix = pathPrefix + "/" + "accounts"
    ret = &AccountRouter{
        PathPrefix: pathPrefix,
        QueryHandler: NewAccountHttpQueryHandler(),
        CommandHandler: NewAccountHttpCommandHandler(context, commandBus),
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

func NewAuthRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus) (ret *AuthRouter) {
    pathPrefix = pathPrefix + "/" + "auth"
    ret = &AuthRouter{
        PathPrefix: pathPrefix,
        AccountRouter: NewAccountRouter(pathPrefix, context, commandBus),
    }
    return
}

func (o *AuthRouter) Setup(router *mux.Router) (ret error) {
    if ret = o.AccountRouter.Setup(router); ret != nil {
        return
    }
    return
}



