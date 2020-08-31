package gochannel

import (
	"encoding/json"

	"github.com/baozhenglab/notification/contract"
)

//DatabaseChannel struct channel for a notifi with driver database
type DatabaseChannel struct{}

//Send mail notifitication to user
// func (mc DatabaseChannel) Send(service interface{}) contract.SendFunc {
// 	db := service.(sdkmongo.MongoDBService)
// 	return func(notifiable interface{}, notifications interface{}) error {
// 		notificationNoti := notifications.(contract.DatabaseNotifice)
// 		user := notifiable.(map[string]interface{})
// 		if _, ok := user["id"]; !ok {
// 			log.Fatal("Requuired id of user")
// 		}
// 		if _, ok := user["type"]; !ok {
// 			log.Fatal("Required id of user")
// 		}
// 		notify := Notification{
// 			Type:           reflect.TypeOf(notifications).Name(),
// 			NotifiableType: fmt.Sprintf("%v", user["type"]),
// 			NotifiableID:   user["id"].(uint),
// 			Data:           notificationNoti.ToDatabase(notifiable),
// 		}
// 		bbytes, _ := bson.Marshal(notify)
// 		_, err := db.GetDatabase().Collection("notifications").InsertOne(context.Background(), bbytes)
// 		return err
// 	}
// }
func endCodeJsonData(notificationNoti contract.DatabaseNotifice, notifiable interface{}) string {
	if notificationNoti.ToDatabase(notifiable) == nil {
		return "{}"
	}
	jsonString, _ := json.Marshal(notificationNoti.ToDatabase(notifiable))
	return string(jsonString)
}
