package person

import (
    "context"
    "fmt"
    "github.com/eugeis/gee/eh"
    "github.com/eugeis/gee/net"
    "github.com/gorilla/mux"
    "github.com/looplab/eventhorizon"
    "html"
    "net/http"
)
type ChurchHttpQueryHandler struct {
    QueryRepository *ChurchQueryRepository
}

func NewChurchHttpQueryHandler(queryRepository *ChurchQueryRepository) (ret *ChurchHttpQueryHandler) {
    ret = &ChurchHttpQueryHandler{
        QueryRepository: queryRepository,
    }
    return
}

func (o *ChurchHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    o.QueryRepository.FindAll()
    fmt.Fprintf(w, "Hello, %q from FindAllChurch", html.EscapeString(r.URL.Path))
}

func (o *ChurchHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindByChurchId", html.EscapeString(r.URL.Path))
}

func (o *ChurchHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountAllChurch", html.EscapeString(r.URL.Path))
}

func (o *ChurchHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountByChurchId", html.EscapeString(r.URL.Path))
}

func (o *ChurchHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistAllChurch", html.EscapeString(r.URL.Path))
}

func (o *ChurchHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistByChurchId", html.EscapeString(r.URL.Path))
}


type ChurchHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewChurchHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *ChurchHttpCommandHandler) {
    ret = &ChurchHttpCommandHandler{
        HttpCommandHandler: eh.NewHttpCommandHandler(context, commandBus),
    }
    return
}

func (o *ChurchHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateChurch{Id: id}, w, r)
}

func (o *ChurchHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateChurch{Id: id}, w, r)
}

func (o *ChurchHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteChurch{Id: id}, w, r)
}


type ChurchRouter struct {
    PathPrefix string
    QueryHandler *ChurchHttpQueryHandler
    CommandHandler *ChurchHttpCommandHandler
    Router *mux.Router
}

func NewChurchRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus, queryRepository *ChurchQueryRepository) (ret *ChurchRouter) {
    pathPrefix = pathPrefix + "/" + "churches"
    ret = &ChurchRouter{
        PathPrefix: pathPrefix,
        QueryHandler: NewChurchHttpQueryHandler(queryRepository),
        CommandHandler: NewChurchHttpCommandHandler(context, commandBus),
    }
    return
}

func (o *ChurchRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountChurchById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("CountChurchAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistChurchById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("ExistChurchAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindChurchById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("FindChurchAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateChurch").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateChurch").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteChurch").HandlerFunc(o.CommandHandler.Delete)
    return
}


type GraduationHttpQueryHandler struct {
    QueryRepository *GraduationQueryRepository
}

func NewGraduationHttpQueryHandler(queryRepository *GraduationQueryRepository) (ret *GraduationHttpQueryHandler) {
    ret = &GraduationHttpQueryHandler{
        QueryRepository: queryRepository,
    }
    return
}

func (o *GraduationHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindAllGraduation", html.EscapeString(r.URL.Path))
}

func (o *GraduationHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindByGraduationId", html.EscapeString(r.URL.Path))
}

func (o *GraduationHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountAllGraduation", html.EscapeString(r.URL.Path))
}

func (o *GraduationHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountByGraduationId", html.EscapeString(r.URL.Path))
}

func (o *GraduationHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistAllGraduation", html.EscapeString(r.URL.Path))
}

func (o *GraduationHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistByGraduationId", html.EscapeString(r.URL.Path))
}


type GraduationHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewGraduationHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *GraduationHttpCommandHandler) {
    ret = &GraduationHttpCommandHandler{
        HttpCommandHandler: eh.NewHttpCommandHandler(context, commandBus),
    }
    return
}

func (o *GraduationHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateGraduation{Id: id}, w, r)
}

func (o *GraduationHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateGraduation{Id: id}, w, r)
}

func (o *GraduationHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteGraduation{Id: id}, w, r)
}


type GraduationRouter struct {
    PathPrefix string
    QueryHandler *GraduationHttpQueryHandler
    CommandHandler *GraduationHttpCommandHandler
    Router *mux.Router
}

func NewGraduationRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus, 
                queryRepository *GraduationQueryRepository) (ret *GraduationRouter) {
    pathPrefix = pathPrefix + "/" + "graduations"
    ret = &GraduationRouter{
        PathPrefix: pathPrefix,
        QueryHandler: NewGraduationHttpQueryHandler(queryRepository),
        CommandHandler: NewGraduationHttpCommandHandler(context, commandBus),
    }
    return
}

func (o *GraduationRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountGraduationById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("CountGraduationAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistGraduationById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("ExistGraduationAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindGraduationById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("FindGraduationAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateGraduation").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateGraduation").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteGraduation").HandlerFunc(o.CommandHandler.Delete)
    return
}


type ProfileHttpQueryHandler struct {
    QueryRepository *ProfileQueryRepository
}

func NewProfileHttpQueryHandler(queryRepository *ProfileQueryRepository) (ret *ProfileHttpQueryHandler) {
    ret = &ProfileHttpQueryHandler{
        QueryRepository: queryRepository,
    }
    return
}

func (o *ProfileHttpQueryHandler) FindByName(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindByProfileName", html.EscapeString(r.URL.Path))
}

func (o *ProfileHttpQueryHandler) FindByEmail(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindByProfileEmail", html.EscapeString(r.URL.Path))
}

func (o *ProfileHttpQueryHandler) FindByPhone(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindByProfilePhone", html.EscapeString(r.URL.Path))
}

func (o *ProfileHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindAllProfile", html.EscapeString(r.URL.Path))
}

func (o *ProfileHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from FindByProfileId", html.EscapeString(r.URL.Path))
}

func (o *ProfileHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountAllProfile", html.EscapeString(r.URL.Path))
}

func (o *ProfileHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from CountByProfileId", html.EscapeString(r.URL.Path))
}

func (o *ProfileHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistAllProfile", html.EscapeString(r.URL.Path))
}

func (o *ProfileHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ExistByProfileId", html.EscapeString(r.URL.Path))
}


type ProfileHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewProfileHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *ProfileHttpCommandHandler) {
    ret = &ProfileHttpCommandHandler{
        HttpCommandHandler: eh.NewHttpCommandHandler(context, commandBus),
    }
    return
}

func (o *ProfileHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateProfile{Id: id}, w, r)
}

func (o *ProfileHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateProfile{Id: id}, w, r)
}

func (o *ProfileHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteProfile{Id: id}, w, r)
}


type ProfileRouter struct {
    PathPrefix string
    QueryHandler *ProfileHttpQueryHandler
    CommandHandler *ProfileHttpCommandHandler
    Router *mux.Router
}

func NewProfileRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus, queryRepository *ProfileQueryRepository) (ret *ProfileRouter) {
    pathPrefix = pathPrefix + "/" + "profiles"
    ret = &ProfileRouter{
        PathPrefix: pathPrefix,
        QueryHandler: NewProfileHttpQueryHandler(queryRepository),
        CommandHandler: NewProfileHttpCommandHandler(context, commandBus),
    }
    return
}

func (o *ProfileRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountProfileById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("CountProfileAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistProfileById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("ExistProfileAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("FindProfileByName").HandlerFunc(o.QueryHandler.FindByName).
    Queries("name", "{name}")
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("FindProfileByEmail").HandlerFunc(o.QueryHandler.FindByEmail).
    Queries("email", "{email}")
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("FindProfileByPhone").HandlerFunc(o.QueryHandler.FindByPhone).
    Queries("phone", "{phone}")
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindProfileById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).
        Name("FindProfileAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateProfile").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateProfile").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteProfile").HandlerFunc(o.CommandHandler.Delete)
    return
}


type PersonRouter struct {
    PathPrefix string
    ChurchRouter *ChurchRouter
    GraduationRouter *GraduationRouter
    ProfileRouter *ProfileRouter
    Router *mux.Router
}

func NewPersonRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus) (ret *PersonRouter) {
    pathPrefix = pathPrefix + "/" + "person"
    ret = &PersonRouter{
        PathPrefix: pathPrefix,
        ChurchRouter: NewChurchRouter(pathPrefix, context, commandBus, queryRepository),
        GraduationRouter: NewGraduationRouter(pathPrefix, context, commandBus, queryRepository),
        ProfileRouter: NewProfileRouter(pathPrefix, context, commandBus, queryRepository),
    }
    return
}

func (o *PersonRouter) Setup(router *mux.Router) (ret error) {
    if ret = o.ChurchRouter.Setup(router); ret != nil {
        return
    }
    if ret = o.GraduationRouter.Setup(router); ret != nil {
        return
    }
    if ret = o.ProfileRouter.Setup(router); ret != nil {
        return
    }
    return
}









