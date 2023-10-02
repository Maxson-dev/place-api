package scheduledevent

type SendNotificationPayload struct {
	To   string `json:"to"`
	Text string `json:"text"`
}

func (s SendNotificationPayload) Type() ScheduledEventType {
	return ScheduledEventTypeSendNotification
}

func (s SendNotificationPayload) IsScheduledEventPayload() {}
