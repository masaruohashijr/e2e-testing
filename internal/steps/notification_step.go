package steps

import "fmt"

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

func IShouldViewTheLastNotificationFromTo(number string) error {
	return nil
}
