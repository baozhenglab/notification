package gochannel

import (
	"github.com/baozhenglab/notification/contract"
	"github.com/baozhenglab/notification/messages"
	slackbot "github.com/baozhenglab/slack-bot"
)

type SlackChannel struct{}

// Send the given notification.
func (SlackChannel) Send(service interface{}) contract.SendFunc {
	slackSv := service.(slackbot.SlackbotService)
	return func(notifiable interface{}, notifications interface{}) error {
		notificationNoti := notifications.(contract.SlackNotifice)
		message := notificationNoti.ToSlack(notifiable).(*messages.SlackMessage)
		return slackSv.SendMessage(message.GetPayload())
	}
}
