package gochannel

import (
	"github.com/baozhenglab/notification/contract"
	"github.com/baozhenglab/notification/messages"
	"github.com/baozhenglab/telegram"
)

type TelegramChannel struct{}

// Send the given notification.
func (tc TelegramChannel) Send(service interface{}) contract.SendFunc {
	telegramSv := service.(telegram.TelegramService)
	return func(notifiable interface{}, notifications interface{}) error {
		notificationNoti := notifications.(contract.TelegramNotifice)
		content := notificationNoti.ToTelegram(notifiable)
		message := new(messages.TelegramMessage)
		to := telegramSv.GetUserName()
		message.Create().To(to).Content(content)
		return telegramSv.SendMessage(message.GetPayload())
	}
}
