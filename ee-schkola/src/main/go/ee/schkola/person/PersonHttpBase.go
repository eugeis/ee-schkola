package person

import (
    "fmt"
    "github.com/eugeis/gee/net"
    "github.com/gorilla/mux"
    "html"
    "net/http"
)
type ChurchHttpQueryHandler struct {
}

func NewChurchHttpQueryHandler() (ret *ChurchHttpQueryHandler) {
    ret = &ChurchHttpQueryHandler{}
    return
}

func (o *ChurchHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
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
}

func NewChurchHttpCommandHandler() (ret *ChurchHttpCommandHandler) {
    ret = &ChurchHttpCommandHandler{}
    return
}

func (o *ChurchHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ChurchCreate", html.EscapeString(r.URL.Path))
}

func (o *ChurchHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ChurchUpdate", html.EscapeString(r.URL.Path))
}

func (o *ChurchHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ChurchDelete", html.EscapeString(r.URL.Path))
}


type ChurchRouter struct {
    PathPrefix string
    QueryHandler *ChurchHttpQueryHandler
    CommandHandler *ChurchHttpCommandHandler
    Router *mux.Router
}

func NewChurchRouter(pathPrefix string) (ret *ChurchRouter) {
    pathPrefix = pathPrefix + "/" + "churches"
    ret = &ChurchRouter{
        PathPrefix: pathPrefix,
        QueryHandler: NewChurchHttpQueryHandler(),
        CommandHandler: NewChurchHttpCommandHandler(),
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
}

func NewGraduationHttpQueryHandler() (ret *GraduationHttpQueryHandler) {
    ret = &GraduationHttpQueryHandler{}
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
}

func NewGraduationHttpCommandHandler() (ret *GraduationHttpCommandHandler) {
    ret = &GraduationHttpCommandHandler{}
    return
}

func (o *GraduationHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from GraduationCreate", html.EscapeString(r.URL.Path))
}

func (o *GraduationHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from GraduationUpdate", html.EscapeString(r.URL.Path))
}

func (o *GraduationHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from GraduationDelete", html.EscapeString(r.URL.Path))
}


type GraduationRouter struct {
    PathPrefix string
    QueryHandler *GraduationHttpQueryHandler
    CommandHandler *GraduationHttpCommandHandler
    Router *mux.Router
}

func NewGraduationRouter(pathPrefix string) (ret *GraduationRouter) {
    pathPrefix = pathPrefix + "/" + "graduations"
    ret = &GraduationRouter{
        PathPrefix: pathPrefix,
        QueryHandler: NewGraduationHttpQueryHandler(),
        CommandHandler: NewGraduationHttpCommandHandler(),
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
}

func NewProfileHttpQueryHandler() (ret *ProfileHttpQueryHandler) {
    ret = &ProfileHttpQueryHandler{}
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
}

func NewProfileHttpCommandHandler() (ret *ProfileHttpCommandHandler) {
    ret = &ProfileHttpCommandHandler{}
    return
}

func (o *ProfileHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ProfileCreate", html.EscapeString(r.URL.Path))
}

func (o *ProfileHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ProfileUpdate", html.EscapeString(r.URL.Path))
}

func (o *ProfileHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello, %q from ProfileDelete", html.EscapeString(r.URL.Path))
}


type ProfileRouter struct {
    PathPrefix string
    QueryHandler *ProfileHttpQueryHandler
    CommandHandler *ProfileHttpCommandHandler
    Router *mux.Router
}

func NewProfileRouter(pathPrefix string) (ret *ProfileRouter) {
    pathPrefix = pathPrefix + "/" + "profiles"
    ret = &ProfileRouter{
        PathPrefix: pathPrefix,
        QueryHandler: NewProfileHttpQueryHandler(),
        CommandHandler: NewProfileHttpCommandHandler(),
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

func NewPersonRouter(pathPrefix string) (ret *PersonRouter) {
    pathPrefix = pathPrefix + "/" + "person"
    ret = &PersonRouter{
        PathPrefix: pathPrefix,
        ChurchRouter: NewChurchRouter(pathPrefix),
        GraduationRouter: NewGraduationRouter(pathPrefix),
        ProfileRouter: NewProfileRouter(pathPrefix),
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



