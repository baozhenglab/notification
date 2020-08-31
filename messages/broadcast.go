package messages

import (
	"fmt"
)

type BroadcastMessage struct {
	Notifiable   interface{}
	Notification interface{}
}

type BroadcastNotifice interface {
	ToBroadcast(notifiable interface{}) map[string]interface{}
}

func (bm BroadcastMessage) BroadcastOn() string {
	user := bm.Notifiable.(map[string]interface{})
	return "private-NotificationBroadcast." + fmt.Sprintf("%v", user["id"])
}

func (bm BroadcastMessage) BroadcastAs() string {
	return `Illuminate\Notifications\Events\BroadcastNotificationCreated`
}

func (bm BroadcastMessage) BroadcastWith() map[string]interface{} {
	return (bm.Notification.(BroadcastNotifice)).ToBroadcast(bm.Notification)
}
