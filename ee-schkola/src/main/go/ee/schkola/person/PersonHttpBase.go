package person

import (
    "context"
    "github.com/eugeis/gee/eh"
    "github.com/eugeis/gee/net"
    "github.com/gorilla/mux"
    "github.com/looplab/eventhorizon"
    "net/http"
)
type ChurchHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *ChurchQueryRepository
}

func NewChurchHttpQueryHandler(queryRepository *ChurchQueryRepository) (ret *ChurchHttpQueryHandler) {
    httpQueryHandler := &eh.HttpQueryHandler{}
    ret = &ChurchHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *ChurchHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllChurch", w, r)
}

func (o *ChurchHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindByChurchId", w, r)
}

func (o *ChurchHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllChurch", w, r)
}

func (o *ChurchHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountByChurchId", w, r)
}

func (o *ChurchHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllChurch", w, r)
}

func (o *ChurchHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistByChurchId", w, r)
}


type ChurchHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewChurchHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *ChurchHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &ChurchHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *ChurchHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateChurch{Id: id}, w, r)
}

func (o *ChurchHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateChurch{Id: id}, w, r)
}

func (o *ChurchHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

func NewChurchRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus, 
                readRepos func (string, func () (ret interface{}) ) (ret eventhorizon.ReadWriteRepo) ) (ret *ChurchRouter) {
    pathPrefix = pathPrefix + "/" + "churches"
    modelFactory := func() interface{} { return NewChurch() }
    repo := readRepos(string(ChurchAggregateType), modelFactory)
    queryRepository := NewChurchQueryRepository(repo, context)
    queryHandler := NewChurchHttpQueryHandler(queryRepository)
    commandHandler := NewChurchHttpCommandHandler(context, commandBus)
    ret = &ChurchRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *ChurchRouter) Setup(router *mux.Router) (err error) {
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
    *eh.HttpQueryHandler
    QueryRepository *GraduationQueryRepository
}

func NewGraduationHttpQueryHandler(queryRepository *GraduationQueryRepository) (ret *GraduationHttpQueryHandler) {
    httpQueryHandler := &eh.HttpQueryHandler{}
    ret = &GraduationHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *GraduationHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllGraduation", w, r)
}

func (o *GraduationHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindByGraduationId", w, r)
}

func (o *GraduationHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllGraduation", w, r)
}

func (o *GraduationHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountByGraduationId", w, r)
}

func (o *GraduationHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllGraduation", w, r)
}

func (o *GraduationHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistByGraduationId", w, r)
}


type GraduationHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewGraduationHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *GraduationHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &GraduationHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *GraduationHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateGraduation{Id: id}, w, r)
}

func (o *GraduationHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateGraduation{Id: id}, w, r)
}

func (o *GraduationHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
                readRepos func (string, func () (ret interface{}) ) (ret eventhorizon.ReadWriteRepo) ) (ret *GraduationRouter) {
    pathPrefix = pathPrefix + "/" + "graduations"
    modelFactory := func() interface{} { return NewGraduation() }
    repo := readRepos(string(GraduationAggregateType), modelFactory)
    queryRepository := NewGraduationQueryRepository(repo, context)
    queryHandler := NewGraduationHttpQueryHandler(queryRepository)
    commandHandler := NewGraduationHttpCommandHandler(context, commandBus)
    ret = &GraduationRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *GraduationRouter) Setup(router *mux.Router) (err error) {
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
    *eh.HttpQueryHandler
    QueryRepository *ProfileQueryRepository
}

func NewProfileHttpQueryHandler(queryRepository *ProfileQueryRepository) (ret *ProfileHttpQueryHandler) {
    httpQueryHandler := &eh.HttpQueryHandler{}
    ret = &ProfileHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *ProfileHttpQueryHandler) FindByEmail(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    email := vars["email"]
    ret, err := o.QueryRepository.FindByEmail(email)
    o.HandleResult(ret, err, "FindByProfileEmail", w, r)
}

func (o *ProfileHttpQueryHandler) FindByPhone(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    phone := vars["phone"]
    ret, err := o.QueryRepository.FindByPhone(phone)
    o.HandleResult(ret, err, "FindByProfilePhone", w, r)
}

func (o *ProfileHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllProfile", w, r)
}

func (o *ProfileHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindByProfileId", w, r)
}

func (o *ProfileHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllProfile", w, r)
}

func (o *ProfileHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountByProfileId", w, r)
}

func (o *ProfileHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllProfile", w, r)
}

func (o *ProfileHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistByProfileId", w, r)
}


type ProfileHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewProfileHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandBus) (ret *ProfileHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &ProfileHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *ProfileHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateProfile{Id: id}, w, r)
}

func (o *ProfileHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateProfile{Id: id}, w, r)
}

func (o *ProfileHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

func NewProfileRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus, 
                readRepos func (string, func () (ret interface{}) ) (ret eventhorizon.ReadWriteRepo) ) (ret *ProfileRouter) {
    pathPrefix = pathPrefix + "/" + "profiles"
    modelFactory := func() interface{} { return NewProfile() }
    repo := readRepos(string(ProfileAggregateType), modelFactory)
    queryRepository := NewProfileQueryRepository(repo, context)
    queryHandler := NewProfileHttpQueryHandler(queryRepository)
    commandHandler := NewProfileHttpCommandHandler(context, commandBus)
    ret = &ProfileRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *ProfileRouter) Setup(router *mux.Router) (err error) {
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

func NewPersonRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandBus, 
                readRepos func (string, func () (ret interface{}) ) (ret eventhorizon.ReadWriteRepo) ) (ret *PersonRouter) {
    pathPrefix = pathPrefix + "/" + "person"
    churchRouter := NewChurchRouter(pathPrefix, context, commandBus, readRepos)
    graduationRouter := NewGraduationRouter(pathPrefix, context, commandBus, readRepos)
    profileRouter := NewProfileRouter(pathPrefix, context, commandBus, readRepos)
    ret = &PersonRouter{
        PathPrefix: pathPrefix,
        ChurchRouter: churchRouter,
        GraduationRouter: graduationRouter,
        ProfileRouter: profileRouter,
    }
    return
}

func (o *PersonRouter) Setup(router *mux.Router) (err error) {
    if err = o.ChurchRouter.Setup(router); err != nil {
        return
    }
    if err = o.GraduationRouter.Setup(router); err != nil {
        return
    }
    if err = o.ProfileRouter.Setup(router); err != nil {
        return
    }
    return
}









