package gochannel

import (
	"fmt"

	"github.com/baozhenglab/mailer"
	"github.com/baozhenglab/notification/contract"
)

//MailChannel struct channel for a notifi with driver mail
type MailChannel struct{}

//Send mail notifitication to user
func (mc MailChannel) Send(service interface{}) contract.SendFunc {
	mailSv := service.(mailer.DriverContract)
	return func(notifiable interface{}, notifications interface{}) error {
		notificationNoti := notifications.(contract.MailNotifice)
		user := notifiable.(map[string]interface{})
		return mailSv.SendMail(fmt.Sprintf("%v", user["email"]), notificationNoti.ToMail(notifiable))
	}
}
