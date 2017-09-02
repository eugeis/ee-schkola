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
	"encoding/json"
	"github.com/eugeis/gee/eh"
	"fmt"
	"github.com/eugeis/gee/lg"
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
	readRepos := func(name string) (ret eventhorizon.ReadWriteRepo) {
		if item, ok := repos[name]; !ok {
			repos[name] = &eh.ReadWriteRepoDelegate{Factory:
			func() (ret eventhorizon.ReadWriteRepo, err error) {
				return repo.NewRepo("localhost", "schkola", name)
			}}
		} else {
			ret = item
		}
		return
	}
	personEngine := person.NewPersonEventhorizonInitializer(eventStore, eventBus, eventPublisher, commandBus, readRepos)

	personEngine.Setup()

	router := mux.NewRouter().StrictSlash(true)

	context := eventhorizon.NewContextWithNamespace(context.Background(), "simple")

	personRouter := person.NewPersonRouter("", context, commandBus, readRepos)
	personRouter.Setup(router)

	//router.Methods(net.GET).Path("/").Name("Index").HandlerFunc(Index)

	router.Methods(net.GET).Path("/").Name("Index").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ret, err := personEngine.ChurchAggregateInitializer.ProjectorRepo.FindAll(context); err == nil {
			var js []byte
			if js, err = json.Marshal(ret); err == nil {
				w.Header().Set("Content-Type", "application/json")
				w.Write(js)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
		return
	})

	log.Err("%v", http.ListenAndServe("127.0.0.1:8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}