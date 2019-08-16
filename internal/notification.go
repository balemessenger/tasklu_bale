package internal

import (
	"errors"
	"taskulu/pkg"
	"taskulu/pkg/taskulu"
	"taskulu/pkg/taskulu/notification"
)

type NotificationService struct {
	log    *pkg.Logger
	client *taskulu.Client
}

func NewNotification(log *pkg.Logger, client *taskulu.Client) *NotificationService {
	return &NotificationService{
		log:    log,
		client: client,
	}
}

func (service *NotificationService) GetUnreadNotifications() ([]notification.NotificationsBody, error) {
	notif, err := service.client.GetNotifications(3)
	if err != nil {
		return nil, err
	}

	if notif.Status != "OK" {
		return nil, errors.New("Bad Request with status: " + notif.Status)
	}

	var results []notification.NotificationsBody

	for _, body := range notif.Data.Notifications {
		if !body.Seen {
			results = append(results, body)
		}
	}

	return results, nil

}

func (service *NotificationService) MarkSeen(notifications []notification.NotificationsBody) error {
	for _, n := range notifications {
		resp, err := service.client.MarkNotificationSeen(n.ID, 3)
		if err != nil {
			return err
		}
		if len(resp.Data) == 0 || resp.Data[0].Ids != n.ID {
			return errors.New("error in marking notification as read")
		}
		return nil
	}
	return nil
}

func (service *NotificationService) MarkAllSeen() error {
	_, err := service.client.MarkAllNotificationSeen(3)
	return err
}
