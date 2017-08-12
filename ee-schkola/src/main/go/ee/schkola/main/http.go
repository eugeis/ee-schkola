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

	personEngine := person.NewPersonEventhorizonInitializer(eventStore, eventBus, eventPublisher, commandBus)

	personEngine.Setup()

	router := mux.NewRouter().StrictSlash(true)

	router.Methods(net.GET).Path("/").Name("Index").HandlerFunc(Index)

	context := eventhorizon.NewContextWithNamespace(context.Background(), "simple")

	personRouter := person.NewPersonRouter("", context, commandBus)
	personRouter.Setup(router)

	log.Err("%v", http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
