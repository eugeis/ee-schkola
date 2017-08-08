package main

import (
	ehh "github.com/looplab/eventhorizon"
	commandbus "github.com/looplab/eventhorizon/commandbus/local"
	eventbus "github.com/looplab/eventhorizon/eventbus/local"
	eventstore "github.com/looplab/eventhorizon/eventstore/memory"
	eventpublisher "github.com/looplab/eventhorizon/publisher/local"
	"ee/schkola/person"
	"context"
	"log"
)

func main() {

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

	/*

	// Create the read repositories.
	invitationRepo := repo.NewRepo()
	guestListRepo := repo.NewRepo()

	// Setup the domain.
	eventID := ehh.NewUUID()

	domain.Setup(
		eventStore,
		eventBus,
		eventPublisher,
		commandBus,
		invitationRepo, guestListRepo,
		eventID,
	)
	*/

	// Set the namespace to use.
	ctx := ehh.NewContextWithNamespace(context.Background(), "simple")
	entityId1 := ehh.NewUUID()

	if err := commandBus.HandleCommand(ctx, &person.CreateChurch{Id: entityId1, Name: "BS"}); err != nil {
		log.Println("error:", err)
	}

}

/*
func StartChat() {
	fmt.Println("Starting application...")
	manager := schkola.NewClientManager()
	go manager.Start()
	http.HandleFunc("/ws", manager.Handler)
	http.ListenAndServe(":12345", nil)
}

func StartMailer() {
	mailer := schkola.Mailer{}
	mailer.SendRegistration()
}
*/