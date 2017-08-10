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


type ChurchHttpCommandHandler struct {
}

func NewChurchHttpCommandHandler() (ret *ChurchHttpCommandHandler) {
    ret = &ChurchHttpCommandHandler{}
    return
}

func (o *ChurchHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    return
    
    return
}

func (o *ChurchHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    return
    
    return
}

func (o *ChurchHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    return
    
    return
}


type ChurchRouter struct {
    PathPrefix string
    Router *mux.Router
    QueryHandler *ChurchHttpQueryHandler
    CommandHandler *ChurchHttpCommandHandler
}

func (o *ChurchRouter) Setup() (ret error) {
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateChurch").HandlerFunc(o.CommandHandler.Create)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateChurch").HandlerFunc(o.CommandHandler.Update)
    o.Router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteChurch").HandlerFunc(o.CommandHandler.Delete)
    return
    
    return
}


type GraduationHttpQueryHandler struct {
}

func NewGraduationHttpQueryHandler() (ret *GraduationHttpQueryHandler) {
    ret = &GraduationHttpQueryHandler{}
    return
}


type GraduationHttpCommandHandler struct {
}

func NewGraduationHttpCommandHandler() (ret *GraduationHttpCommandHandler) {
    ret = &GraduationHttpCommandHandler{}
    return
}

func (o *GraduationHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    return
    
    return
}

func (o *GraduationHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    return
    
    return
}

func (o *GraduationHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    return
    
    return
}


type GraduationRouter struct {
    PathPrefix string
    Router *mux.Router
    QueryHandler *GraduationHttpQueryHandler
    CommandHandler *GraduationHttpCommandHandler
}

func (o *GraduationRouter) Setup() (ret error) {
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateGraduation").HandlerFunc(o.CommandHandler.Create)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateGraduation").HandlerFunc(o.CommandHandler.Update)
    o.Router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteGraduation").HandlerFunc(o.CommandHandler.Delete)
    return
    
    return
}


type ProfileHttpQueryHandler struct {
}

func NewProfileHttpQueryHandler() (ret *ProfileHttpQueryHandler) {
    ret = &ProfileHttpQueryHandler{}
    return
}

func (o *ProfileHttpQueryHandler) FindByName(w http.ResponseWriter, r *http.Request)  {
    return
    
    return
}

func (o *ProfileHttpQueryHandler) FindByEmail(w http.ResponseWriter, r *http.Request)  {
    return
    
    return
}

func (o *ProfileHttpQueryHandler) FindByPhone(w http.ResponseWriter, r *http.Request)  {
    return
    
    return
}


type ProfileHttpCommandHandler struct {
}

func NewProfileHttpCommandHandler() (ret *ProfileHttpCommandHandler) {
    ret = &ProfileHttpCommandHandler{}
    return
}

func (o *ProfileHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request)  {
    return
    
    return
}

func (o *ProfileHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request)  {
    return
    
    return
}

func (o *ProfileHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request)  {
    return
    
    return
}


type ProfileRouter struct {
    PathPrefix string
    Router *mux.Router
    QueryHandler *ProfileHttpQueryHandler
    CommandHandler *ProfileHttpCommandHandler
}

func (o *ProfileRouter) Setup() (ret error) {
    o.Router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindProfileByName").HandlerFunc(o.QueryHandler.FindByName)
    o.Router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindProfileByEmail").HandlerFunc(o.QueryHandler.FindByEmail)
    o.Router.Methods(net.GET).PathPrefix(o.PathPrefix).Name("FindProfileByPhone").HandlerFunc(o.QueryHandler.FindByPhone)
    o.Router.Methods(net.POST).PathPrefix(o.PathPrefix).Name("CreateProfile").HandlerFunc(o.CommandHandler.Create)
    o.Router.Methods(net.PUT).PathPrefix(o.PathPrefix).Name("UpdateProfile").HandlerFunc(o.CommandHandler.Update)
    o.Router.Methods(net.DELETE).PathPrefix(o.PathPrefix).Name("DeleteProfile").HandlerFunc(o.CommandHandler.Delete)
    return
    
    return
}


type PersonRouter struct {
    Router *mux.Router
    PathPrefix string
    ChurchRouter *ChurchRouter
    GraduationRouter *GraduationRouter
    ProfileRouter *ProfileRouter
}

func (o *PersonRouter) Setup() (ret error) {
    return
    return
}



