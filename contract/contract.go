package contract

import (
	"github.com/baozhenglab/mailer"
	"github.com/baozhenglab/notification/messages"
)

type TelegramNotifice interface {
	ToTelegram(notifiable interface{}) string
}

type SlackNotifice interface {
	ToSlack(notifiable interface{}) interface{}
}

// MailNotifice interface required method config ToMail return Mailable
type MailNotifice interface {
	ToMail(notifiable interface{}) mailer.Mailable
}

// OneSignalNotifice interface required method config ToOneSignal return OneSignalMessage
type OneSignalNotifice interface {
	ToOneSignal(notifiable interface{}) messages.OneSignalMessage
}

// DatabaseNotifice interface required method config ToMail return interface
type DatabaseNotifice interface {
	ToDatabase(notifiable interface{}) interface{}
}

type BroadcastNotifice interface {
	ToBroadcast(notifiable interface{}) map[string]interface{}
}

type DriverChannel interface {
	Send(service interface{}) SendFunc
}

type ShouldBroadcast interface {
	/**
	 * Get the channels the event should broadcast on.
	 */
	BroadcastOn() string

	/**
	 * The event's broadcast name.
	 *
	 * @return string
	 */
	BroadcastAs() string

	/**
	 * Get the data to broadcast.
	 *
	 */
	BroadcastWith() map[string]interface{}
}

type Notification interface {
	Via() []string
}

type SendFunc func(interface{}, interface{}) error
