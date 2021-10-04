package notifications

import "zarbat_test/pkg/domains"

type PrimaryPort interface {
	ViewNotification(notificationSid string) (domains.Notification, error)
	ListNotifications() ([]domains.Notification, error)
}

type SecondaryPort interface {
	ViewNotification(notificationSid string) (domains.Notification, error)
	ListNotifications() ([]domains.Notification, error)
}
