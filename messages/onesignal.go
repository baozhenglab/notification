package messages

import "github.com/baozhenglab/onesignal"

//OneSignalMessage allows us to configure sending data request
type OneSignalMessage struct {
	Contents     string
	Headings     string
	Filters      interface{}
	Data         interface{}
	SmallIcon    string
	IOSSound     string
	AndroidSound string
	//Even though have AppID is export, it will be set when create notification
	//So, It only read
	AppID string
}

func (os OneSignalMessage) GetContent() map[string]string {
	return map[string]string{
		"en": os.Contents,
	}
}
func (os OneSignalMessage) GetHeading() map[string]string {
	return map[string]string{
		"en": os.Headings,
	}
}

func (os OneSignalMessage) SetAppID(appID string) OneSignalMessage {
	os.AppID = appID
	return os
}

//ToNotificationRequest convert to OneSignalMessage of onesignal
func (os OneSignalMessage) ToNotificationRequest() onesignal.NotificationRequest {
	return onesignal.NotificationRequest{
		Contents:     os.GetContent(),
		Headings:     os.GetHeading(),
		Filters:      os.Filters,
		Data:         os.Data,
		SmallIcon:    os.SmallIcon,
		IOSSound:     os.IOSSound,
		AndroidSound: os.AndroidSound,
		AppID:        os.AppID,
	}
}
