package event

import (
	"bytes"
	"fmt"
	"net/http"
)

func (r *EventRouter) sendToCalendarServiceDelete(content string) error {
	req, err := http.NewRequest("DELETE", r.CalendarServiceURL+"/delete-event", bytes.NewBuffer([]byte(content)))
	if err != nil {
		return fmt.Errorf("error sendToCalendarServiceDelete: creating request: %w", err)

	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sendToCalendarServiceDelete: sending request: %w", err)

	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("error sendToCalendarServiceDelete: Status Code: %d", resp.StatusCode)

	}
	return nil
}
