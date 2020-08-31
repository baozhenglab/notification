package gochannel

import (
	"fmt"

	"github.com/baozhenglab/notification/contract"
	"github.com/baozhenglab/onesignal"
)

//OneSignalChannel struct channel for a notifi with driver onesignal
type OneSignalChannel struct{}

//Send  notifitication to user
func (oc OneSignalChannel) Send(service interface{}) contract.SendFunc {
	onesignalSV := service.(onesignal.OneSignalService)
	return func(notifiable interface{}, notifications interface{}) error {
		notificationNoti := notifications.(contract.OneSignalNotifice)
		message := notificationNoti.ToOneSignal(notifiable)
		notificationReq := message.SetAppID(onesignalSV.GetAppID()).ToNotificationRequest()
		createRes, res, err := onesignalSV.SendNotification(notificationReq)
		fmt.Println(createRes, res, err)
		return err
	}

}
