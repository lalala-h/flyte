package interfaces

import (
	"context"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

//go:generate mockery --name=Sender --output=../mocks --case=underscore --with-expecter

// Sender Defines the interface for sending cloudevents.
type Sender interface {
	// Send a cloud event to other services (AWS pub/sub, Kafka, Nats).
	Send(ctx context.Context, notificationType string, event cloudevents.Event) error
}
