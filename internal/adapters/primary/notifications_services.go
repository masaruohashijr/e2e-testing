package primary

import (
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/notifications"
)

type port_notifications struct {
	driven notifications.SecondaryPort
}

func NewNotificationsService(driven notifications.SecondaryPort) notifications.PrimaryPort {
	return &port_notifications{
		driven,
	}
}

func (p *port_notifications) ViewNotification(notificationsSid string) (domains.Notification, error) {
	notifications, err := p.driven.ViewNotification(notificationsSid)
	return notifications, err
}

func (p *port_notifications) ListNotifications(from, to string) ([]domains.Notification, error) {
	notificationss, err := p.driven.ListNotifications(from, to)
	return notificationss, err
}

func (p *port_notifications) SendNotification(from, to, message string) error {
	err := p.driven.SendNotification(from, to, message)
	return err
}
