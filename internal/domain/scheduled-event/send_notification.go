package scheduledevent

type SendNotificationPayload struct {
	To   string `json:"to"`
	Text string `json:"text"`
}

func (s SendNotificationPayload) Type() Type {
	return SendNotification
}

func (s SendNotificationPayload) IsScheduledEventPayload() {}
