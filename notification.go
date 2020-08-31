package notification

import (
	goservice "github.com/baozhenglab/go-sdk"
	"github.com/baozhenglab/notification/contract"
	gochannel "github.com/baozhenglab/notification/go-channel"
	"github.com/baozhenglab/onesignal"
	slackbot "github.com/baozhenglab/slack-bot"
)

type notificationService struct {
	sc goservice.ServiceContext
}

type NotificationService interface {
	Send(notifiables interface{}, notification contract.Notification)
}

func NewNotificationService(sc goservice.ServiceContext) NotificationService {
	return notificationService{sc}
}

func createTelegreamDriver() gochannel.TelegramChannel {
	return gochannel.TelegramChannel{}
}

// func createBroadcastDriver() gochannel.BroadcastChannel {
// 	return gochannel.BroadcastChannel{}
// }

func createMailDriver() gochannel.MailChannel {
	return gochannel.MailChannel{}
}

// func createDatabaseDriver() gochannel.DatabaseChannel {
// 	return gochannel.DatabaseChannel{}
// }
func createOneSignalDriver() gochannel.OneSignalChannel {
	return gochannel.OneSignalChannel{}
}

func createSlackDriver() gochannel.SlackChannel {
	return gochannel.SlackChannel{}
}

func getDriver(name string) contract.DriverChannel {
	switch name {
	case "telegram":
		return createTelegreamDriver()
	// case "broadcast":
	// 	return createBroadcastDriver()
	case "mail":
		return createMailDriver()
	// case "database":
	// 	return createDatabaseDriver()
	case "onesignal":
		return createOneSignalDriver()
	case "slack":
		return createSlackDriver()
	default:
		panic("no driver")
	}
}

func (noti notificationService) getService(driver string) interface{} {
	switch driver {
	case "telegram":
		return noti.sc.MustGet(telegram.KeyService)
	case "broadcast":
		return noti.sc.MustGet(gorabbitmq.KeyService)
	case "mail":
		return noti.sc.MustGet(mail.KeyMailService)
	case "database":
		return noti.sc.MustGet("mdb")
	case "onesignal":
		return noti.sc.MustGet(onesignal.KeyService)
	case "slack":
		return noti.sc.MustGet(slackbot.KeyService)
	default:
		panic("no driver")
	}
}

//Send the given notification immediately.
func (noti notificationService) sendNow(notifiables []interface{}, notification contract.Notification) {
	for _, notifiable := range notifiables {
		viaChannel := notification.Via()
		go func() {
			for _, channel := range viaChannel {
				driver := getDriver(channel)
				go driver.Send(noti.getService(channel))(notifiable, notification)
			}
		}()
	}
}

//Format the notifiables into a Collection / array if necessary.
func formatNotifiables(notifiables interface{}) []interface{} {
	if convert, ok := notifiables.([]interface{}); ok {
		return convert
	}
	return []interface{}{notifiables}
}

//Send the given notification to the given notifiable entities.
func (noti notificationService) Send(notifiables interface{}, notification contract.Notification) {
	notifiablesConvert := formatNotifiables(notifiables)
	noti.sendNow(notifiablesConvert, notification)
}
