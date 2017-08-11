package person

import (
    "github.com/eugeis/gee/net"
    "github.com/gorilla/mux"
    "net/http"
)
type ChurchHttpQueryHandler struct {
}

func NewChurchHttpQueryHandler() (ret *ChurchHttpQueryHandler) {
    ret = &ChurchHttpQueryHandler{}
    return
}

func (o *ChurchHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *ChurchHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
}

func (o *ChurchHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
}

func (o *ChurchHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *ChurchHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *ChurchHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
}


type ChurchHttpCommandHandler struct {
}

func NewChurchHttpCommandHandler() (ret *ChurchHttpCommandHandler) {
    ret = &ChurchHttpCommandHandler{}
    return
}

func (o *ChurchHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
}

func (o *ChurchHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
}

func (o *ChurchHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
}


type ChurchRouter struct {
    PathPrefix string
    QueryHandler *ChurchHttpQueryHandler
    CommandHandler *ChurchHttpCommandHandler
    Router *mux.Router
}

func NewChurchRouter(PathPrefix string) (ret *ChurchRouter) {
    
    ret.PathPrefix = ret.PathPrefix + "/" + "churches"
    return
}

func (o *ChurchRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindChurchAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindChurchById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountChurchById").HandlerFunc(o.QueryHandler.CountById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountChurchAll").HandlerFunc(o.QueryHandler.CountAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistChurchAll").HandlerFunc(o.QueryHandler.ExistAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistChurchById").HandlerFunc(o.QueryHandler.ExistById)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateChurch").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateChurch").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteChurch").HandlerFunc(o.CommandHandler.Delete)
    return
}


type GraduationHttpQueryHandler struct {
}

func NewGraduationHttpQueryHandler() (ret *GraduationHttpQueryHandler) {
    ret = &GraduationHttpQueryHandler{}
    return
}

func (o *GraduationHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *GraduationHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request)  {
}

func (o *GraduationHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request)  {
}

func (o *GraduationHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *GraduationHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request)  {
}

func (o *GraduationHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request)  {
}


type GraduationHttpCommandHandler struct {
}

func NewGraduationHttpCommandHandler() (ret *GraduationHttpCommandHandler) {
    ret = &GraduationHttpCommandHandler{}
    return
}

func (o *GraduationHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
}

func (o *GraduationHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
}

func (o *GraduationHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
}


type GraduationRouter struct {
    PathPrefix string
    QueryHandler *GraduationHttpQueryHandler
    CommandHandler *GraduationHttpCommandHandler
    Router *mux.Router
}

func NewGraduationRouter(PathPrefix string) (ret *GraduationRouter) {
    
    ret.PathPrefix = ret.PathPrefix + "/" + "graduations"
    return
}

func (o *GraduationRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindGraduationAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindGraduationById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountGraduationById").HandlerFunc(o.QueryHandler.CountById)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("CountGraduationAll").HandlerFunc(o.QueryHandler.CountAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistGraduationAll").HandlerFunc(o.QueryHandler.ExistAll)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("ExistGraduationById").HandlerFunc(o.QueryHandler.ExistById)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateGraduation").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateGraduation").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteGraduation").HandlerFunc(o.CommandHandler.Delete)
    return
}


type ProfileHttpQueryHandler struct {
}

func NewProfileHttpQueryHandler() (ret *ProfileHttpQueryHandler) {
    ret = &ProfileHttpQueryHandler{}
    return
}

func (o *ProfileHttpQueryHandler) FindByName(w http.ResponseWriter, r *http.Request)  {
}

func (o *ProfileHttpQueryHandler) FindByEmail(w http.ResponseWriter, r *http.Request)  {
}

func (o *ProfileHttpQueryHandler) FindByPhone(w http.ResponseWriter, r *http.Request)  {
}


type ProfileHttpCommandHandler struct {
}

func NewProfileHttpCommandHandler() (ret *ProfileHttpCommandHandler) {
    ret = &ProfileHttpCommandHandler{}
    return
}

func (o *ProfileHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
}

func (o *ProfileHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
}

func (o *ProfileHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
}


type ProfileRouter struct {
    PathPrefix string
    QueryHandler *ProfileHttpQueryHandler
    CommandHandler *ProfileHttpCommandHandler
    Router *mux.Router
}

func NewProfileRouter(PathPrefix string) (ret *ProfileRouter) {
    
    ret.PathPrefix = ret.PathPrefix + "/" + "profiles"
    return
}

func (o *ProfileRouter) Setup(router *mux.Router) (ret error) {
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindProfileByName").HandlerFunc(o.QueryHandler.FindByName)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindProfileByEmail").HandlerFunc(o.QueryHandler.FindByEmail)
    router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindProfileByPhone").HandlerFunc(o.QueryHandler.FindByPhone)
    router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateProfile").HandlerFunc(o.CommandHandler.Create)
    router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateProfile").HandlerFunc(o.CommandHandler.Update)
    router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteProfile").HandlerFunc(o.CommandHandler.Delete)
    return
}


type PersonRouter struct {
    PathPrefix string
    ChurchRouter *ChurchRouter
    GraduationRouter *GraduationRouter
    ProfileRouter *ProfileRouter
    Router *mux.Router
}

func NewPersonRouter(PathPrefix string) (ret *PersonRouter) {
    
    ret.PathPrefix = ret.PathPrefix + "/" + "person"
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



