package main

import (
	"github.com/eugeis/gee/lg"
	"net/http"
	"github.com/gorilla/mux"
	"ee/schkola/person"
	"fmt"
	"github.com/eugeis/gee/net"
	"github.com/looplab/eventhorizon"
	commandbus "github.com/looplab/eventhorizon/commandbus/local"
	eventbus "github.com/looplab/eventhorizon/eventbus/local"
	eventstore "github.com/looplab/eventhorizon/eventstore/memory"
	eventpublisher "github.com/looplab/eventhorizon/publisher/local"
	repo "github.com/looplab/eventhorizon/repo/memory"
	"context"
)

var log = lg.NewLogger("Schkola ")

func main() {
	log.Info("Server started")

	// Create the event store.
	eventStore := eventstore.NewEventStore()

	// Create the event bus that distributes events.
	eventBus := eventbus.NewEventBus()
	eventPublisher := eventpublisher.NewEventPublisher()
	eventBus.SetPublisher(eventPublisher)

	// Create the command bus.
	commandBus := commandbus.NewCommandBus()

	repos := make(map[string]eventhorizon.ReadWriteRepo)
	readRepos := func(name string) (ret eventhorizon.ReadWriteRepo) {
		if item, ok := repos[name]; !ok {
			ret = repo.NewRepo()
			repos[name] = ret
		} else {
			ret = item
		}
		return
	}
	personEngine := person.NewPersonEventhorizonInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)

	personEngine.Setup()

	router := mux.NewRouter().StrictSlash(true)

	router.Methods(net.GET).Path("/").Name("Index").HandlerFunc(Index)

	context := eventhorizon.NewContextWithNamespace(context.Background(), "simple")

	personRouter := person.NewPersonRouter("", context, commandBus)
	personRouter.Setup(router)

	log.Err("%v", http.ListenAndServe("127.0.0.1:8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
