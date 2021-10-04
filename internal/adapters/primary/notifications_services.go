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

func (p *port_notifications) ListNotifications() ([]domains.Notification, error) {
	notificationss, err := p.driven.ListNotifications()
	return notificationss, err
}
