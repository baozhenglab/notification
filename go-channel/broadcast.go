package gochannel

// //MailChannel struct channel for a notifi with driver mail
// type BroadcastChannel struct{}

// type BroadcastCollection struct {
// 	Event   string                 `json:"event"`
// 	Channel string                 `json:"channel"`
// 	Data    map[string]interface{} `json:"data"`
// }

// //Send mail notifitication to user
// func (bc BroadcastChannel) Send(service interface{}) contract.SendFunc {
// 	kafkaSV := service.(gorabbitmq.RabbitMQService)
// 	return func(notifiable interface{}, notifications interface{}) error {
// 		broadcast := messages.BroadcastMessage{notifiable, notifications}
// 		message := BroadcastCollection{
// 			Event:   broadcast.BroadcastAs(),
// 			Channel: broadcast.BroadcastOn(),
// 			Data:    broadcast.BroadcastWith(),
// 		}
// 		kafkaSV.Publish(message, "notifications")
// 		return nil
// 	}
// }
