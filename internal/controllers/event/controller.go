package event

import (
	"encoding/json"
	"fmt"

	"github.com/Deadrafa/J.A.R.V.I.S/pkg/models"
)

type EventRouter struct {
	CalendarServiceURL string
}

type actionStatus struct {
	ActionStatus string `json:"action_status"`
}

func (r *EventRouter) RouteEvent(event *models.GigChatResp) error {
	var actStat actionStatus
	if err := json.Unmarshal([]byte(event.Choices[0].Message.Content), &actStat); err != nil {
		return fmt.Errorf("error in RouteEvent(): Unmarshal: %w", err)

	}

	switch actStat.ActionStatus {
	case "add":
		return r.sendToCalendarServiceCreate(event.Choices[0].Message.Content)
	case "delete":
		return r.sendToCalendarServiceDelete(event.Choices[0].Message.Content)
	default:
		return fmt.Errorf("unknown action_status: %s", actStat.ActionStatus)
	}

}
