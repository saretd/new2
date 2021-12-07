package sender

import (
	"github.com/ozonmp/trv-airport-api/internal/model"
)

//EventSender sends Event
type EventSender interface {
	Send(airport *model.AirportEvent) error
}
