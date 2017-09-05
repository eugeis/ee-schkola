package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"ee/schkola/person"
	"github.com/eugeis/gee/net"
	"github.com/looplab/eventhorizon"
	commandbus "github.com/looplab/eventhorizon/commandbus/local"
	eventbus "github.com/looplab/eventhorizon/eventbus/local"
	eventstore "github.com/looplab/eventhorizon/eventstore/mongodb"
	eventpublisher "github.com/looplab/eventhorizon/publisher/local"
	repo "github.com/looplab/eventhorizon/repo/mongodb"
	"context"
	"github.com/eugeis/gee/eh"
	"fmt"
	"github.com/eugeis/gee/lg"
	"ee/schkola/student"
	"ee/schkola/library"
	"ee/schkola/finance"
	"ee/schkola/auth"
	"github.com/rs/cors"
)

var log = lg.NewLogger("Schkola ")

func main() {
	log.Info("Server started")

	// Create the event store.
	eventStore := &eh.EventStoreDelegate{Factory:
	func() (ret eventhorizon.EventStore, err error) {
		return eventstore.NewEventStore("localhost", "schkola")
	}}

	// Create the event bus that distributes events.
	eventBus := eventbus.NewEventBus()
	eventPublisher := eventpublisher.NewEventPublisher()
	eventBus.SetPublisher(eventPublisher)

	// Create the command bus.
	commandBus := commandbus.NewCommandBus()

	repos := make(map[string]eventhorizon.ReadWriteRepo)
	readRepos := func(name string, factory func() interface{}) (ret eventhorizon.ReadWriteRepo) {
		if item, ok := repos[name]; !ok {
			ret = &eh.ReadWriteRepoDelegate{Factory: func() (ret eventhorizon.ReadWriteRepo, err error) {
				var retRepo *repo.Repo
				if retRepo, err = repo.NewRepo("localhost", "schkola", name); err == nil {
					retRepo.SetModel(factory)
					ret = retRepo
				}
				return
			}}
			repos[name] = ret
		} else {
			ret = item
		}
		return
	}

	authEngine := auth.NewAuthEventhorizonInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
	authEngine.Setup()

	financeEngine := finance.NewFinanceEventhorizonInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
	financeEngine.Setup()

	libraryEngine := library.NewLibraryEventhorizonInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
	libraryEngine.Setup()

	personEngine := person.NewPersonEventhorizonInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
	personEngine.Setup()

	studentEngine := student.NewStudentEventhorizonInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)
	studentEngine.Setup()

	router := mux.NewRouter().StrictSlash(true)

	context := eventhorizon.NewContextWithNamespace(context.Background(), "simple")

	authRouter := auth.NewAuthRouter("", context, commandBus, readRepos)
	authRouter.Setup(router)

	financeRouter := finance.NewFinanceRouter("", context, commandBus, readRepos)
	financeRouter.Setup(router)

	libraryRouter := library.NewLibraryRouter("", context, commandBus, readRepos)
	libraryRouter.Setup(router)

	personRouter := person.NewPersonRouter("", context, commandBus, readRepos)
	personRouter.Setup(router)

	studentRouter := student.NewStudentRouter("", context, commandBus, readRepos)
	studentRouter.Setup(router)

	router.Methods(net.GET).Path("/").Name("Index").HandlerFunc(Index)
	router.NotFoundHandler = http.HandlerFunc(NoFound)

	handler := cors.Default().Handler(router)

	log.Err("%v", http.ListenAndServe("127.0.0.1:8080", handler))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Schkola!")
}

func NoFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Schkola!")
}
