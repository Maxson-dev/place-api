package scheduledevent

import (
	"testing"
	"time"

	"github.com/Maxson-dev/place-api/internal/pkg/cerror"
	"github.com/stretchr/testify/assert"
)

func TestNew_Successful(t *testing.T) {
	t.Parallel()

	datetime := time.Now()
	eventType := SendNotification
	payload := []byte(`{"to": "example@example.com", "text": "Hello, world!"}`)

	event, err := New(datetime, string(eventType), payload)

	assert.NoError(t, err)
	assert.Equal(t, datetime, event.Datetime)
	assert.Equal(t, eventType, event.Payload.Type())
}

func TestNew_InvalidJSONPayload(t *testing.T) {
	t.Parallel()

	datetime := time.Now()
	eventType := SendNotification
	invalidPayload := []byte(`{"to": "example@example.com", "text": }`)

	event, err := New(datetime, string(eventType), invalidPayload)

	assert.Error(t, err)
	assert.Equal(t, cerror.ErrBadInput, err)
	assert.Equal(t, ScheduledEvent{}, event)
}

func TestNew_UnknownEventType(t *testing.T) {
	t.Parallel()

	datetime := time.Now()
	unknownEventType := "UNKNOWN"
	payload := []byte(`{"foo": "bar"}`)

	event, err := New(datetime, unknownEventType, payload)

	assert.Error(t, err)
	assert.Equal(t, cerror.ErrBadInput, err)
	assert.Equal(t, ScheduledEvent{}, event)
}
