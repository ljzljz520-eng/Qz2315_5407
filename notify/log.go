package notify

import "independentjournal/domain"

type Delivery struct{ RecordID, Channel, Message string }

func BuildDelivery(r domain.Record, role string) Delivery {
	return Delivery{RecordID: r.ID, Channel: ChannelForRole(role), Message: DeliveryMessage(r)}
}
func IsDeliverable(d Delivery) bool { return d.RecordID != "" && d.Channel != "" && d.Message != "" }
func DeliveryKey(d Delivery) string { return d.RecordID + ":" + d.Channel }
