package messages

type SlackMessage struct {
	payload map[string]string
}

func (rn *SlackMessage) Create() *SlackMessage {
	rn.payload = map[string]string{}
	return rn
}

//Recipient's Chat ID.
func (tm *SlackMessage) To(to string) *SlackMessage {
	tm.payload["channel"] = to
	return tm
}

//Notification message (Supports Markdown).
func (tm *SlackMessage) Content(content string) *SlackMessage {
	tm.payload["text"] = content
	return tm
}

//Get body message for telegram channel
func (tm *SlackMessage) GetPayload() map[string]string {
	return tm.payload
}
