package messages

type TelegramMessage struct {
	payload map[string]string
}

func (tm *TelegramMessage) Create() *TelegramMessage {
	tm.payload = map[string]string{"parse_mode": "Markdown"}
	return tm
}

//Recipient's Chat ID.
func (tm *TelegramMessage) To(to string) *TelegramMessage {
	tm.payload["chat_id"] = to
	return tm
}

//Notification message (Supports Markdown).
func (tm *TelegramMessage) Content(content string) *TelegramMessage {
	tm.payload["text"] = content
	return tm
}

//Get body message for telegram channel
func (tm *TelegramMessage) GetPayload() map[string]string {
	return tm.payload
}
