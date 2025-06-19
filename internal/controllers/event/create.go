package event

import (
	"bytes"
	"fmt"
	"net/http"
)

func (r *EventRouter) sendToCalendarServiceCreate(content string) error {
	req, err := http.NewRequest("POST", r.CalendarServiceURL+"/create-event", bytes.NewBuffer([]byte(content)))
	if err != nil {
		return fmt.Errorf("error sendToCalendarServiceCreate: creating request: %w", err)

	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sendToCalendarServiceCreate: sending request: %w", err)

	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("error sendToCalendarServiceCreate: Status Code: %d", resp.StatusCode)

	}
	return nil
}
