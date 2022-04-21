package steps

import (
	"fmt"
	"zarbat_test/internal/logging"
)

func IShouldListAtLeastNotification(amount int) error {
	notifications, err := NotificationPrimaryPort.ListNotifications()
	if err != nil {
		return fmt.Errorf("Could not list notifications.")
	}
	if len(notifications) < amount {
		return fmt.Errorf("Expected %d notifications, but got %d.", amount, len(notifications))
	}
	return nil
}

func IShouldViewTheLastNotification() error {
	notifications, err := NotificationPrimaryPort.ListNotifications()
	if err != nil {
		return fmt.Errorf("Could not list notifications.")
	}
	notification, err := NotificationPrimaryPort.ViewNotification(notifications[0].Sid)
	if err != nil {
		return fmt.Errorf("Expected one notification, but got 0.")
	}
	logging.Debug.Println("Log Level: ", notification.Log)
	return nil
}
